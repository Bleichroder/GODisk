# GODisk

	A openstack based , golang implemented web server and storage server.

## Deployment

	As GODisk is a Front-Back project, it is a complex job to illuminate the deployment. Please read the deployment.docx


## Infrastructure

#### OS:

	* Ubuntu Desktop 16.04 LTS

#### Platform:

	* Openstack: Version Undetermined
	* Client: openstack/golang-client

#### Language:

	* Golang: latest
	* IDE: liteide latest

#### Framework(choice):

	* Beego
	* revel
	* gin
	* echo

#### Database:

	* MongoDB: Version Undetermined
	* mgo: Version Undetermined
	* MariaDB server: 10.0.x
	* Driver for database/sql: go-sql-driver/mysql V1.2 stable

#### Web:

	* jQuery: 2.2.4 (3.x cannot work with Bootstrap)
	* Bootstrap: 3.3.6
	* Fine Uploader: 5.10.1 (core mode)

## Contributing

	* If you are not a member yet, please consider joining as a contributor.
	* If you have questions or comments, please contact lcx061125@163.com 

## Development Environment

#### Golang

##### The Go Programming Language  [http://golangtc.com/download](http://golangtc.com/download)

* Installation:  
	1. 	Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go.
	`sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`
	2.  Add /usr/local/go/bin to the PATH environment variable.
	`export PATH=$PATH:/usr/local/go/bin`
	`export GOROOT=/usr/local/go` 
	3.  Add you workspace dir into GOPATH.
	`export GOPATH=/home/jason/workspace/go`
	4. 	**Attention: You would better add step 2&3 into /etc/profile**
* Test: go version
		
##### Liteide  [http://golangtc.com/download/liteide](http://golangtc.com/download/liteide)

* Installation:
	1.  Download the archive and extract it into /usr/local.
	`sudo tar -C /usr/local -xzf liteidex$VERSION.$OS-$ARCH.tar.bz2`
	2.  Launch liteide.
	`/usr/local/liteide/bin/liteide`

#### Text Editor

##### Sublime Text 3  [http://www.sublimetext.com/3](http://www.sublimetext.com/3)

* Registration: [https://fatesinger.com/?s=sublime](https://fatesinger.com/?s=sublime)
* Package:
	* Package control: [https://packagecontrol.io/](https://packagecontrol.io/)
		* Emmet
		* JSFormat
		* jQuery
		* Bracket Highlighter
		* Markdown Preview
