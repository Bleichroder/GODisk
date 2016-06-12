package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

const PCWebDir = "static"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer

	cmd := exec.Command("ls", "-l", "/home/jason/workspace/git/GODisk/other/web_template/")
	cmd.Stdout = &buffer
	cmd.Run()
	count := stringHandler(buffer.String(), w)
	fmt.Fprintf(w, "Total %d \n", count)
}

func stringHandler(s string, w http.ResponseWriter) int {

	var filetype string
	var filecounter int

	for index, content := range s {
		if content == '\n' {
			indexNextStart := index + 1
			if indexNextStart < len(s) {
				filecounter++
				filename := stringSep(s[indexNextStart:])
				if s[indexNextStart] == 'd' {
					filetype = "folder"
				} else {
					filetype = "file"
				}
				fmt.Fprintf(w, "%s\t\t%s\n", filetype, filename)
			}
		}
	}
	return filecounter
}

func stringSep(s string) (name string) {
	end := 0
	for index, content := range s {
		if content == '\n' {
			end = index
			break
		}
	}
	for a := end; a > 0; a-- {
		if s[a] == ' ' {
			b := a + 1
			name = s[b:end]
			break
		}
	}
	return name
}

func main() {
	http.Handle("/css/", http.FileServer(http.Dir(PCWebDir)))
	http.Handle("/img/", http.FileServer(http.Dir(PCWebDir)))
	http.Handle("/js/", http.FileServer(http.Dir(PCWebDir)))
	http.Handle("/fonts/", http.FileServer(http.Dir(PCWebDir)))
	http.Handle("/font-awesome/", http.FileServer(http.Dir(PCWebDir)))

	//http.HandleFunc("/", homeHandler)
	http.Handle("/", http.FileServer(http.Dir(PCWebDir)))
	//http.HandleFunc("/login/", logInHandler)
	//http.HandleFunc("/signup/", signUpHandler)
	http.HandleFunc("/index/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
