package controllers

import (
	"craw/models"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func tracefile(str_content string)  {
	fd,_:=os.OpenFile("a.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	fd_time:=time.Now().Format("2006-01-02 15:04:05");
	fd_content:=strings.Join([]string{"======",fd_time,"=====",str_content,"\n"},"")
	buf:=[]byte(fd_content)
	fd.Write(buf)
	fd.Close()
}
func s(){
	for us:=range ch{
		username:=us.username
		userid:=us.userid
		ojid:=us.ojid
		CrawlPOJ(username,userid,ojid)
	}

}
func CrawlPOJ(username string,userid int ,ojid int)  {
	// Request the HTML page.

	top:=""
	end:=false
	count:=0
	for{
		if end==true{
			break
		}
		end=true
		url:= "http://poj.org/status?user_id="+username+"&top="+top
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		time.Sleep(1*time.Second)
		if resp.StatusCode != 200 {
			fmt.Println("resp.StatusCode!=200:",resp.StatusCode,err)
			return
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err!=nil{
			fmt.Println(err)
			return
		}
		tracefile(doc.Text())
		var serviceOK=false
		doc.Find("table.a").Each(func(i int, s *goquery.Selection){
			serviceOK=true
			s.Find("tr").Each(func(j int ,ss *goquery.Selection){
				if j>0{
					var memory int
					var problemid string
					var runid string
					var Result string
					var time int
					var language string
					var codelength string
					var submittime string
					var ojuser string
					status:=1
					ss.Find("td").Each(func(k int, selection *goquery.Selection) {
					//	fmt.Println(k,selection.Text())
						if k==0{
							end=false
							top=selection.Text()
							runid=selection.Text()
						}
						if k==1{
						ojuser=selection.Text()
						}
						if k==2{

							if selection.Text()!=""{
								problemid=selection.Text()
							}
						}
						if k==3{
								Result=selection.Text()
								if Result=="Accepted"{
									status=0
								}
						}
						if k==4{
							if selection.Text()!=""&&len(selection.Text())>1{
								memory,err=strconv.Atoi(selection.Text()[0:len(selection.Text())-1])
								if err!=nil{
									memory=0
								}
							}else{
								memory=0
							}
						}

						if k==5{
							if selection.Text()!=""&&len(selection.Text())>2{
								time,err=strconv.Atoi(selection.Text()[0:len(selection.Text())-2])
								if err!=nil{
									time=0
								}
							}else{
								time=0
							}
						}
						if k==6{
							language=selection.Text()
						}
						if k==7{
							codelength=selection.Text()
						}
						if k==8{
								submittime=selection.Text()
						}


					})
					if models.FindVJrunid(ojid,runid)!=1{
						r,err:=models.AddVJUserProblem(userid,ojid,problemid,memory,time,language,submittime,strconv.Itoa(status),Result,ojuser,codelength,runid)
						if r==false&&err!=nil{
							panic(err)
						}
						count++
					}else{
						end=true
					}

				}

			})
		})
		if !serviceOK{
			fmt.Println("Service Error")
		}
		//time.Sleep(time.Second)
	}

fmt.Println(username,count)


}