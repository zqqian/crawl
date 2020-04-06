package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)
type VjUserList struct{
	Userid string
	UserVJName string
}
var o orm.Ormer
func init()  {
	orm.RegisterDriver("mysql",orm.DRMySQL)
//	orm.RegisterDataBase("default","mysql","root:root@/stepbystep?charset=utf8&loc=Asia%2FShanghai")
	orm.RegisterDataBase("default", "mysql", "root:Zqqian123@tcp(bj-cynosdbmysql-grp-eqivy326.sql.tencentcdb.com:22627)/stepbystep?charset=utf8&loc=Asia%2FShanghai")

	o=orm.NewOrm()

}
func add(sql string)(bool,error){
	res, err := o.Raw(sql).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		//fmt.Println("mysql row affected nums: ", num)
		if num==1{
			return true,nil
		}else{
			return false,errors.New("Add failed")
		}
	}else{
		return false,err
	}
}
func FindVJid(ojname string) int {
	sql:="SELECT id FROM `ojlist` WHERE ojname=?"
	var lists []orm.ParamsList
	num, err := o.Raw(sql,ojname).ValuesList(&lists)
	if err == nil && num > 0 {
		id, _ := strconv.Atoi(lists[0][0].(string))
		if id > 0 {
			return id
		} else {
			return -1
		}
	} else {
		return -1
	}
}
func FindVJrunid(vjid int,runid string) int {
	sql:="SELECT id FROM `userproblem` WHERE oj ="+strconv.Itoa(vjid)+" and runid="+runid+""
	var lists []orm.ParamsList
	num, err := o.Raw(sql).ValuesList(&lists)
	if err == nil && num > 0 {
		id, _ := strconv.Atoi(lists[0][0].(string))
		if id > 0 {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}
func GetVjUserList(vjid int)[]VjUserList{
	sql:="SELECT * FROM `user-oj` WHERE `oj` = ?\n"
	var lists []orm.ParamsList
	num, err := o.Raw(sql, vjid).ValuesList(&lists)
	var l []VjUserList
	if err == nil && num > 0 {
		for i := 0; i < len(lists); i++ {
			var c VjUserList
			c.Userid = lists[i][1].(string)
			c.UserVJName = lists[i][3].(string)

			l = append(l, c)
		}

	}
	return l

}
func AddVJUserProblem(user int, oj int, problem string, memory int, runtime int, language string,submittime string,statustype string,status string ,ojusername string, codelength string,runid string)  (r bool,err error){
	sql := "INSERT INTO `userproblem` (`id`, `user`, `oj`, `problem`, `memory`, `runtime`, `language`, `submittime`, `statustype`, `status`, `crawtime`, `ojusername`, `codelength`,`runid`) VALUES (NULL, '" + strconv.Itoa(user) + "', " + strconv.Itoa(oj) + ", '" + problem + "', '" + strconv.Itoa(memory) + "', '" + strconv.Itoa(runtime) + "', '" + language + "', '"+submittime+"', '"+statustype+"', '"+status+"', CURRENT_TIMESTAMP, '"+ojusername+"', '"+codelength+"','"+runid+"');"
	//fmt.Println(sql)
	return add(sql)
}