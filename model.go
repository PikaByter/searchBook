package main

type DirInfo struct {
	Name string `json:"Name"`
	Size int `json:"Size"`
	FileList []string `json:"FileList"`
	DirList []DirInfo `json:"DirList"`
}

type DirInfoAndPath struct {
	DirInfo
	path string
}
