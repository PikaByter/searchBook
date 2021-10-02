package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var paths []string

func file_get_contents(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func loadData(){
	var d DirInfo
	loadDirInfo(&d)
	d=d.DirList[0]
	pathQ:=CreateQueue(999)
	pathQ.Push(DirInfoAndPath{d,""})
	paths=getPaths(pathQ)
}

func searchBook(a string)[]string{
	var res []string
	for _,val:=range paths{
		if strings.Contains(val,a){
			res=append(res,val)
		}
	}
	return res
}

func getPaths(pathQ Queue)[]string{
	var res []string
	for {
		if pathQ.Empty() {
			break
		}
		dNow := pathQ.Pop()
		if len(dNow.FileList) != 0 {
			suf := dNow.path + "/" + dNow.Name + "/"
			for _, val := range dNow.FileList {
				fileName := strings.Split(val, "|")
				res = append(res, suf+fileName[0])
			}
		}
		if len(dNow.DirList) != 0 {
			for _, val := range dNow.DirList {
				pathQ.Push(DirInfoAndPath{
					val,
					dNow.path + "/" + dNow.Name,
				})
			}
		}
	}
	return res
}

func loadDirInfo(d *DirInfo){
	var content []byte
	var err error

	content, err = file_get_contents("电子书.json")
	if err != nil {
		fmt.Println("open file error: " + err.Error())
		panic(err)
	}
	err = json.Unmarshal([]byte(content), &d)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		panic(err)
	}
}