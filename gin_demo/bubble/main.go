package main

import (
	"net/http"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gin-gonic/gin"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMysql() (err error) {

	dsn := "root:123456@tcp(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {

	//创建数据库
	//sql:CREATE DATABASE bubble;
	//连接数据库
	err := initMysql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	//模型绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()

	r.Static("/static", "static")

	//告诉Gin框架去哪里找到模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")

	{
		//待办事项
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			//从前端页面填入代办事项 点击提交，会发请求到这里
			//1.从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			//存入数据库
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				//c.JSON(http.StatusOK, todo)
				c.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":  "success",
					"data": todo,
				})
			}

		})
		//查看所有
		v1Group.GET("/todo", func(c *gin.Context) {

			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})

			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		//指定查看
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "invaild id!!!"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			}
			c.BindJSON(&todo)
			DB.Save(todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				//c.JSON(http.StatusOK, todo)
				c.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":  "success",
					"data": todo,
				})
			}

		})
		//删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "invaild id!!!"})
				return
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}

		})

	}

	r.Run()
}
