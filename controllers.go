package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/**********************************************************************************
 Index path controller
**********************************************************************************/

func (this *indexController) IndexAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/html/index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
	log.Println("Index HTML transmition complete!")
}

func (this *indexController) IndexGetpathAction(w http.ResponseWriter, r *http.Request) {

	// Split Path
	pathInfo := r.FormValue("destPath")
	path := strings.Split(strings.Trim(pathInfo, "/"), "/")
	pathLength := len(path)
	log.Println("Path request input: ", path)

	// Render Home Path Template
	var tempHome PathMiddle
	tempHome.AbsolutePath = "/"
	tempHome.RelativePath = "Home"
	err := template.Must(template.New("pathhome").Parse(PathTemplateMiddle)).Execute(w, tempHome)
	if err != nil {
		log.Println(err)
	}

	if path[0] != "" {
		// Render Middle Path Template
		tempAb := "/"
		tempMiddle := make([]PathMiddle, pathLength)
		for i := 0; i < pathLength-1; i++ {
			tempAb += path[i] + "/"
			tempMiddle[i].AbsolutePath = tempAb
			tempMiddle[i].RelativePath = path[i]
			err := template.Must(template.New("pathmiddle").Parse(PathTemplateMiddle)).Execute(w, tempMiddle[i])
			if err != nil {
				log.Println(err)
			}
		}
		// Render Current Path Template
		var tempCurrent PathCurrent
		tempCurrent.Path = path[pathLength-1]
		err := template.Must(template.New("pathcurrent").Parse(PathTemplateCurrent)).Execute(w, tempCurrent)
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("Path information request all responded.")
}

func (this *indexController) IndexGetinodeAction(w http.ResponseWriter, r *http.Request) {

	// Open Database
	GODiskDB, err := dbInit()
	if err != nil {
		log.Println(err)
	}
	defer GODiskDB.Close()

	// Query all file information
	path := r.FormValue("destPath")
	ret := inodeQuery(GODiskDB, "user_jason", path)
	inodeList := template.Must(template.New("inodelist").Parse(InodeTemplate))
	for _, inode := range ret {
		var temp InodeToTemplate
		temp.FileName = inode.FileName
		temp.FileSize = strconv.Itoa(inode.FileSize)
		switch inode.FileType {
		case "folder":
			temp.TypeClass = "folder"
		case "text":
			temp.TypeClass = "file-text-o"
		case "code":
			temp.TypeClass = "file-code-o"
		case "archive":
			temp.TypeClass = "file-archive-o"
		case "audio":
			temp.TypeClass = "file-audio-o"
		case "image":
			temp.TypeClass = "file-image-o"
		case "video":
			temp.TypeClass = "file-video-o"
		case "pdf":
			temp.TypeClass = "file-pdf-o"
		case "word":
			temp.TypeClass = "file-word-o"
		case "excel":
			temp.TypeClass = "file-excel-o"
		case "powerpoint":
			temp.TypeClass = "file-powerpoint-o"
		default:
			temp.TypeClass = "file-o"
		}
		temp.ModTime = inode.ModTime
		err = inodeList.Execute(w, temp)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("File information request all responded.")
}

func (this *indexController) TaskAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/html/task.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
	log.Println("Task HTML transmition complete!")
}

func (this *indexController) SettingAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/html/setting.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
	log.Println("Setting HTML transmition complete!")
}

func (this *indexController) SkinconfigAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/html/skin-config.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
	log.Println("Skinconfig HTML transmition complete!")
}

/**********************************************************************************
 Register path controller
**********************************************************************************/

func (this *registerController) IndexAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/html/register.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
	log.Println("Register HTML transmition complete!")
}

func (this *registerController) SubmitAction(w http.ResponseWriter, r *http.Request) {

	type Result struct {
		Ret int
		Log string
	}

	var registryResult *Result
	var userInfo userRegistryInfo

	// Parameter Transformation
	userInfo.name = r.FormValue("register_username")
	userInfo.password = r.FormValue("register_password")
	userInfo.confirm = r.FormValue("register_confirm")
	userInfo.authcode = r.FormValue("register_authcode")
	log.Println("User registry request: {" + userInfo.name + "}{" + userInfo.password + "}{" + userInfo.confirm + "}{" + userInfo.authcode + "}")

	// Open Database
	GODiskDB, err := dbInit()
	if err != nil {
		log.Println(err)
	}
	defer GODiskDB.Close()

	// Register service
	ret := registerService(GODiskDB, &userInfo)
	switch ret {
	case 0:
		registryResult = &Result{0, "Registry success."}
	case 1:
		registryResult = &Result{1, "Wrong authority code."}
	case 2:
		registryResult = &Result{2, "Username already exists."}
	default:
		log.Println("GODisk server inner error!")
		registryResult = &Result{3, "Inner Error."}
	}

	// Response
	b, err := json.Marshal(registryResult)
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("Response message have sent.")
		w.Write(b)
	}
}

/**********************************************************************************
 Login path controller
**********************************************************************************/

func (this *loginController) IndexAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/html/login.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
	log.Println("Login HTML transmition complete!")
}

func (this *loginController) SubmitAction(w http.ResponseWriter, r *http.Request) {

	type Result struct {
		Ret int
		Log string
	}

	var loginResult *Result
	var userInfo userLoginInfo

	// Parameter transformation
	userInfo.name = r.FormValue("login_username")
	userInfo.password = r.FormValue("login_password")
	log.Println("User login request: {" + userInfo.name + "}{" + userInfo.password + "}")

	// Open Database
	GODiskDB, err := dbInit()
	if err != nil {
		log.Println(err)
	}
	defer GODiskDB.Close()

	// Login service
	ret := loginService(GODiskDB, &userInfo)
	switch ret {
	case 0:
		loginResult = &Result{0, "Login success."}
		cookie := http.Cookie{Name: "username", Value: userInfo.name, Path: "/"}
		http.SetCookie(w, &cookie)
	case 1:
		loginResult = &Result{1, "Username not exist."}
	case 2:
		loginResult = &Result{2, "Password not match."}
	default:
		log.Println("GODisk server inner error!")
		loginResult = &Result{3, "Inner error."}
	}

	// Response
	b, err := json.Marshal(loginResult)
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("Response message of login handler have sent.")
		w.Write(b)
	}
}

/**********************************************************************************
 Not Found path controller
**********************************************************************************/
func NotFoundAction(w http.ResponseWriter) {
	t, err := template.ParseFiles("static/html/404.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
	log.Println("404 HTML transmition complete!")
}
