package controllers

import (
	"craw/models"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
var ch = make(chan struct {
	username string
	userid   int
	ojid     int
})
func C(){

	go GetUserList("poj")
	go s()
	time.Sleep(100*time.Second)
}
func GetUserList(ojname string)  {
	vjid:=models.FindVJid(ojname)
	if vjid==-1{
		//-------------------
		return
	}
	UserList:=models.GetVjUserList(vjid)
	for i:=0;i< len(UserList);i++{
		id,err:=strconv.Atoi(UserList[i].Userid)
		if(err==nil){
			if ojname=="vjudge"{
				CrawlVJ(UserList[i].UserVJName,id,vjid)
			}else if ojname =="poj"{
				ch<- struct {
					username string
					userid   int
					ojid     int
				}{username: UserList[i].UserVJName, userid: id, ojid: vjid}
				//time.Sleep(3*time.Second)
						//CrawlPOJ(UserList[i].UserVJName,id,vjid)
			}

		}
	}
}