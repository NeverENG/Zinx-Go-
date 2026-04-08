package util

import (
	"encoding/json"
	"fmt"
	"os"
	"zinx/src/ziface"
)

type GlobalObj struct {
	TcpServer ziface.IServer

	Version        string
	Host           string
	Port           int
	Name           string
	MaxConn        int
	MaxPackageSize int
}

var G *GlobalObj

func (G *GlobalObj) Reload() {
	data, err := os.ReadFile("../util/conf/zinx.json")
	if err != nil {
		fmt.Println("[GlobalObj.Reload]" + err.Error())
		panic(err)
	}
	err = json.Unmarshal(data, &G)
	fmt.Println("[ERROR]：全局变量初始化失败")
}

func init() {
	G = &GlobalObj{
		Name:           "zinx v0.4",
		Version:        "v0.4",
		Host:           "0.0.0.0",
		Port:           8080,
		MaxConn:        1024,
		MaxPackageSize: 1024,
	}
	G.Reload()
}
