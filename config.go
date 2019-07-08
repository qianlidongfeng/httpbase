package main

import (
	"github.com/qianlidongfeng/httpserver"
	"github.com/qianlidongfeng/loger"
	"github.com/qianlidongfeng/toolbox"
)

type Config struct{
	Debug bool
	Httpserver httpserver.Config
	Log loger.Config
	DB toolbox.MySqlConfig
}
