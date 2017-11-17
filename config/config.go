package config

var DB string
var DictPath string
var Port int

func ConfigInit() {
    DB = "cuihovah:7576583asd@localhost:27017"
    DictPath = "/usr/local/gopath/src/github.com/huichen/sego/data/dictionary.txt"
    Port = 8023
}
