package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	postJsonResponse(r)
	multiParamsFormResponse(r)
	queryAndFormPostResponse(r)

	r.Run(":8888") //All requests will access port:8888.
}

/*
	Post response: http://localhost:8888/post
	c.DefaultPostForm will get key value first, if it doesn't have, it will use the default value.
*/
func postJsonResponse(router *gin.Engine) {
	router.POST("/post", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "jeff")
		pwd := c.DefaultPostForm("pwd", "pwd")
		c.JSON(http.StatusOK, gin.H{
			"message": "hell gyy",
			"user":    user,
			"pwd":     pwd,
		})
	})
}

//Post response: http://localhost:8888/form_post
func multiParamsFormResponse(router *gin.Engine) {
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
}

/*
	Post response: http://localhost:8888/query_form_post
	c.Query is used to get arguments before "?".
	c.PostForm is used to get arguments from post form.
 */
func queryAndFormPostResponse(router *gin.Engine) {
	router.POST("/query_form_post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		//During output Using fmt.Printf does not need convert string to int.
		//Using fmt.Printf, it won't have an output on browser, but it will have the output on command line.
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)

		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})
}
