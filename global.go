package main

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/qianlidongfeng/loger"
	"github.com/qianlidongfeng/toolbox"
	"html/template"
	"os"
	"runtime"
)

var (
	G Global
)

type Global struct{
	cfg Config
	log loger.Loger
	db *sql.DB
	DB *Database
	slash string
	middleware map[string]echo.MiddlewareFunc
	renderer echo.Renderer
}

func (this *Global) Init(){
	configFile,err:=toolbox.GetConfigFile()
	if err != nil{
		panic(err)
	}
	err=toolbox.LoadConfig(configFile,&this.cfg)
	if err != nil{
		panic(err)
	}
	if this.cfg.Debug == false{
		toolbox.RediRectOutPutToLog()
	}
	this.log,err=loger.NewLoger(this.cfg.Log)
	if err != nil{
		panic(err)
	}
	this.db,err=toolbox.InitMysql(this.cfg.DB)
	if err != nil{
		this.log.Fatal(err)
	}
	this.DB=NewDatabase(this.db)
	if appdir,err:=toolbox.AppDir();err != nil{
		this.log.Fatal(err)
	}else if err:=os.Chdir(appdir);err!=nil{
		this.log.Fatal(err)
	}
	if runtime.GOOS == "windows"{
		this.slash="\\"
	}else{
		this.slash="/"
	}
	this.initMiddleWare()
}


func (this *Global) initMiddleWare(){
	mw:=Mw{}
	this.middleware=make(map[string]echo.MiddlewareFunc)
	this.middleware["recover"]=mw.recover(this.log)
	this.middleware["loginJwtAuth"]=mw.loginJwtAuth([]byte(this.cfg.Secret.LoginJwt))
}


func (this *Global) initRenderer(){
	renderer := &Renderer{}
	renderer.templates=template.Must(template.ParseGlob("template"+G.slash+"*.html"))
	this.renderer=renderer
}