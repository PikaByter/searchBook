package main

import (
	"fmt"
	"net/http"
)
import "github.com/gin-gonic/gin"
//写成server,Over
//todo:部署到服务器
//todo:导出阿里云盘分享链接：https://github.com/wxy1343/aliyunpan
//todo:使用qq实现便捷查询
//todo：使用信息检索的方式，分词后加快检索

func handelGet(c *gin.Context){
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	c.HTML(http.StatusOK, "searchBook.html", gin.H{})
}

func handelResultPost(c *gin.Context){
	value:=c.PostForm("search")
	fmt.Println("获得检索请求",value)
	res:=searchBook(value)
	fmt.Println("结果为：",res)
	c.HTML(http.StatusOK, "result.html", gin.H{
		"data": res,
	})
}
func main() {
	loadData()
	g:=gin.Default()
	g.LoadHTMLGlob("templates/*")
	g.GET("/", handelGet)
	g.POST("/result",handelResultPost)
	g.Run(":8293")
	//a:="计算机网络"
	//loadData()
	//res:=searchBook(a)
	//for _,val:=range res{
	//	fmt.Println(val)
	//}

}
