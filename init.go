package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/qianlidongfeng/toolbox"
)

func init(){
	G=Global{}
	G.Init()
	InitTables()
}

func InitTables(){
	InitUserTable()
}

func InitUserTable(){
	fileds:=map[string]string{
		"id":`int(11) UNSIGNED NOT NULL AUTO_INCREMENT`,
		"name":`varchar(32) NOT NULL`,
		"password":`varchar(64) NOT NULL`,
		"grop":`varchar(16) NOT NULL`,
		"ltime":`datetime(0) NOT NULL`,
	}
	index:=map[string]toolbox.MysqlIndex{
			"name":toolbox.MysqlIndex{Typ:"unique",Name:"column",Method:"BTREE"},
	}
	if err:=toolbox.CheckAndFixTable(G.db,"users",fileds,index);err!=nil{
		G.log.Fatal(err)
	}
}