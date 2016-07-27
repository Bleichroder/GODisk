package main

import (
	"database/sql"
)

var GODiskDB *sql.DB

type indexController struct{}
type registerController struct{}
type loginController struct{}

type userLoginInfo struct {
	name     string
	password string
}

type userRegistryInfo struct {
	name     string
	password string
	confirm  string
	authcode string
}

type Inode struct {
	Path      string
	FileName  string
	FileType  string
	FileSize  int
	ModTime   string
	MD5       string
	StorageID string
}

type InodeToTemplate struct {
	TypeClass string
	FileName  string
	FileSize  string
	ModTime   string
}

type TaskToTemplate struct {
	TypeClass string
	FileName  string
	Progress  string
	TranSpeed string
}
