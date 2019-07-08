package main

import (
	"github.com/qianlidongfeng/loger"
)

type Global struct{
	cfg Config
	log loger.Loger
}

var (
	global Global
)