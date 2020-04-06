package main

import (
	"craw/controllers"
	"fmt"
	"time"
)

func e()  {
	fmt.Println("a")
	time.Sleep(3*time.Second)
	fmt.Println("b")
}
func d()  {
	 e()
}
func c(){

	 e()
}
func main() {
//	controllers.GetUserList()
controllers.C()
time.Sleep(100*time.Second)
	//controllers.CrawlPOJ("sdufe20161846316",111,2)
}

