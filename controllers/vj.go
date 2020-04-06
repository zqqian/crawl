package controllers

import (
	"craw/models"
	"crypto/tls"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"strconv"
	"time"
)
type VJresult struct{
	Memory interface{}
	Access interface{}
	StatusType interface{}
	Runtime interface{}
	Language interface{}
	StatusCanonical interface{}
	UserName interface{}
	UserId interface{}
	LanguageCanonical interface{}
	UDebugUrl interface{}
	Processing interface{}
	RunId int
	Time int64
	Oj interface{}
	ProblemId interface{}
	SourceLength interface{}
	ProbNum interface{}
	Status interface{}

}
type Vj struct{
	Data []VJresult
}
type SubmitResult struct {
	Data []struct {
		Memory int `json:"memory"`
		Access int `json:"access"`
		StatusType int `json:"statusType"`
		Runtime int `json:"runtime"`
		Language string `json:"language"`
		StatusCanonical string `json:"statusCanonical"`
		UserName string `json:"userName"`
		UserID int `json:"userId"`
		LanguageCanonical string `json:"languageCanonical"`
		UDebugURL string `json:"uDebugUrl"`
		Processing bool `json:"processing"`
		RunID int `json:"runId"`
		Time int64 `json:"time"`
		Oj string `json:"oj"`
		ProblemID int `json:"problemId"`
		SourceLength int `json:"sourceLength"`
		ProbNum string `json:"probNum"`
		Status string `json:"status"`
	} `json:"data"`
	RecordsTotal int `json:"recordsTotal"`
	RecordsFiltered int `json:"recordsFiltered"`
	Draw int `json:"draw"`
}

func CrawlVJ(username string,userid int,vjid int)  {
	var res SubmitResult
	start:=0
	end:=false
	for{
		if end{
			break
		}
		req:=httplib.Post("https://vjudge.net/status/data/").SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		req.Header("Accept-Encoding","gzip,deflate,sdch")
		req.Header("User-Agent","Mozilla/5.0 (SDUFE_Acm_Bot/1.0; Sorry for trouble caused)")
		req.Param("start",strconv.Itoa(start))
		req.Param("length","20")
		req.Param("un",username)
		req.Param("OJId","All")
		req.Param("probNum","")
		req.Param("res","0")
		req.Param("language","")
		req.Param("orderBy","run_id")
		fmt.Println(req.String())
		req.ToJSON(&res)
		fmt.Println(res)
		fmt.Println(username)
		if len(res.Data)<20{
			end=true
		}else{
			start+=20
		}
		for j:=0;j< len(res.Data);j++{
			if models.FindVJrunid(vjid,strconv.Itoa(res.Data[j].RunID))==1{
				end=true
				break
			}
			tm := time.Unix(res.Data[j].Time/1000, 0)
			models.AddVJUserProblem(userid,vjid,res.Data[j].Oj+res.Data[j].ProbNum,res.Data[j].Memory,res.Data[j].Runtime,res.Data[j].Language,tm.Format("2006-01-02 15:04:05"),strconv.Itoa(res.Data[j].StatusType),res.Data[j].Status,res.Data[j].UserName,strconv.Itoa(res.Data[j].SourceLength),strconv.Itoa(res.Data[j].RunID))
		}

		time.Sleep(time.Second)
	}

}

