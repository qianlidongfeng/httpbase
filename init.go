package main

import (
	"github.com/qianlidongfeng/loger"
	"github.com/qianlidongfeng/toolbox"
	"os"
)

func init(){
	global=Global{}
	configFile,err:=toolbox.GetConfigFile()
	if err != nil{
		panic(err)
	}
	err=toolbox.LoadConfig(configFile,&global.cfg)
	if err != nil{
		panic(err)
	}
	if global.cfg.Debug == false{
		toolbox.RediRectOutPutToLog()
	}
	global.log,err=loger.NewLoger(global.cfg.Log)
	if err != nil{
		panic(err)
	}
	apppath,err:=toolbox.AppDir()
	if err != nil{
		global.log.Fatal(err)
	}
	if err:=os.Chdir(apppath);err!=nil{
		global.log.Fatal(err)
	}
}