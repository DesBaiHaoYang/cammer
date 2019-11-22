package datasql

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

//Baginfo 包的属性
type Baginfo struct {
	Bagid    string `form:"bagid"`
	Imgface  string `form:"imgface"`
	Imgbag   string `form:"imgbag"`
	Phonenum string `form:"phonenum"`
	// ID      string `json:id`
}

// Querybag 请求体属性
type Querybag struct {
	Bagid    string `form:"bagid"`
	Phonenum string `form:"phonenum"`
}

var db *sql.DB

func Querybaginfo(c *gin.Context) {

	var p Querybag
	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"statecode": "1001",
			"message":   "原因是" + err.Error(),
			"title":     "服务器错误",
		})
	} else {
		bagid := p.Bagid
		phonenum := p.Phonenum
		var b Baginfo
		b, err = queryDB(bagid, phonenum)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"statecode": "1001",
				"title":     "提示",
				"message":   "未查到该号码的信息",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"statecode": "200",
			"imgface":   b.Imgface,
			"imgbag":    b.Imgbag,
		})
	}
}
func Submitdata(c *gin.Context) {
	var p Baginfo
	err := c.BindJSON(&p)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"err_no":  "4001",
			"message": "request's params wrong!,err:" + err.Error(),
		})

	} else {
		bagid := p.Bagid
		imgface := p.Imgface
		imgbag := p.Imgbag
		phonenum := p.Phonenum
		isexits, err := insertDB(bagid, phonenum, imgface, imgbag)
		if isexits {
			if err == nil {
				c.JSON(http.StatusOK, gin.H{
					"statecode": "1001",
					"message":   "原因是编号已存在",
					"title":     "数据存储失败",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"statecode": "1001",
					"message":   "原因是" + err.Error(),
					"title":     "数据存储失败",
				})
			}
			return
		}
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"statecode": "1001",
				"message":   "原因是" + err.Error(),
				"title":     "数据存储失败",
			})
		} else {
			returncode := bagid
			if returncode == "" {
				returncode = phonenum
			}
			c.JSON(http.StatusOK, gin.H{
				"statecode": "200",
				"message":   "取包编号为:" + returncode,
				"title":     "入库成功",
			})
		}
	}
}
func insertDB(bagid string, phonenum string, imgface string, imgbag string) (iseits bool, ierr error) {
	_, err := queryDB(bagid, phonenum)
	iseits = false
	if err == nil {
		iseits = true
		return iseits, err
	} else {
		var cretatetime = time.Now().Format("2006-1-2 15:04:05")
		iseits = false
		selstr := "INSERT INTO `baginfo` (`bagid`,`phonenum`,`imgface`,`imgbag`,`createtime`)VALUES('" + bagid + "', '" + phonenum + "', '" + imgface + "', '" + imgbag + "', '" + cretatetime + "');"
		_, err := db.Exec(selstr)
		if err != nil {
			fmt.Println("插入数据错误:" + err.Error())
		}
		return iseits, err
	}
}

//QueryDB 查询数据库
func queryDB(bagid string, phonenum string) (bags Baginfo, errs error) {
	var selstr string
	if bagid == "" {
		selstr = "SELECT imgface,imgbag FROM baginfo where phonenum = '" + phonenum + "';"
	} else {
		selstr = "SELECT imgface,imgbag FROM baginfo where bagid = '" + bagid + "';"
	}

	var baginfos Baginfo
	err := db.QueryRow(selstr).Scan(&baginfos.Imgface, &baginfos.Imgbag)
	if err != nil {
		fmt.Println("我是错误" + err.Error())
	}
	return baginfos, err
}

//OpenDB 打开数据库
func OpenDB() (err error) {
	fmt.Println("开始连接数据库")
	//数据库信息
	dsn := "root:root@tcp(127.0.0.1:3306)/test"
	//链接数据库
	db, err = sql.Open("mysql", dsn) //open不会校验用户名和密码是否正确 知识检测数据源格式是否正确 只有格式不正确的时候才会报错
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	fmt.Println("数据库链接成功")
	return nil
}
