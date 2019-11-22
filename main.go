package main

import (
	datasql "bag/dataserver"
	"bag/ipserver"
	"fmt"

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
	r := gin.Default()
	r.POST("/queryinfo", datasql.Querybaginfo)
	r.POST("/submitdata", datasql.Submitdata)
	r.GET("/", awstt)
	r.Run(":81")

}
func awstt(c *gin.Context) {
	c.Writer.WriteString("hellobhy")

}
