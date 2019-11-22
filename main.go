package main

import (
	datasql "bag/dataserver"
	"bag/ipserver"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	ip, iperr := ipserver.ExternalIP()
	if iperr != nil {
		fmt.Println(iperr)
	}
	fmt.Println("当前本机IP地址为")
	fmt.Println(ip.String())
	err := datasql.OpenDB()
	if err != nil {
		fmt.Println("数据库链接失败，请检查")
		time.Sleep(time.Second * 10)
		panic("失败")
	} else {
		r := gin.Default()
		r.POST("/queryinfo", datasql.Querybaginfo)
		r.POST("/submitdata", datasql.Submitdata)
		r.Run(":9000")
	}

}
