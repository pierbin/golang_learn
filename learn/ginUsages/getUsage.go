package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ResponseParams struct {
	Name   string `json:"name" uri:"name"`
	Age    int    `json:"age" uri:"age"`
	Gender string `json:"gender" uri:"gender"`
}

// r means router
func main() {
	r := gin.Default()
	jsonResponse(r)
	jsonResponseParamAndQuery(r)
	jsonResponseFullPath(r)
	jsonResponseBind(r)
	jsonResponseBindUri(r)
	getResponseString(r)

	r.Run(":8888") //All requests will access port:8888.
}

//Get Request: http://localhost:8888/ping
func jsonResponse(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin framework base usage.",
		})
	})
}

/*
	Get Response: http://localhost:8888/get/123?user=davy&pwd=admin
	c.Param is used to get arguments before "?".
	c.DefaultQuery will get key value first, if it does not have, it will use the default one.
	c.Query is used to get arguments after "?".
*/
func jsonResponseParamAndQuery(router *gin.Engine) {
	// This handler will match /user/{id} but will not match /user/ or /user
	router.GET("/get/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.DefaultQuery("user", "jeff")
		pwd := c.Query("pwd")

		//c.JSON return a json.
		c.JSON(http.StatusOK, gin.H{
			"message": "hell gyy",
			"id":      id,
			"user":    user,
			"pwd":     pwd,
		})
	})
}

// "c.String" will output string, but it always download a file.
func getResponseString(router *gin.Engine) {
	// However, this one will match /user/{id}/ and also /user/{id}/send
	// If no other routers match /user/{id}, it will redirect to /user/{id}/
	router.GET("/user/:id/*action", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		action := c.Param("action")
		message := "User ID is " + string(id) + ", action is " + action

		c.String(http.StatusOK, message)
	})
}

func jsonResponseFullPath(router *gin.Engine) {
	// For each matched request Context will hold the route definition
	router.GET("/user/groups", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": c.FullPath(),
		})
	})
}

func jsonResponseBind(router *gin.Engine) {
	router.GET("/testBind", func(c *gin.Context) {
		p := ResponseParams{}
		err := c.ShouldBind(&p) //Here if it uses c.ShouldBindJSON(), it will have an error. So change it to c.ShouldBind() fix it.
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "Something wrong",
				"data": gin.H{},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "success",
				"data": p,
			})
		}
	})
}

//Get request: http://localhost:8888/getParamsFromUrl/davy/40/male
func jsonResponseBindUri(router *gin.Engine) {
	router.GET("/getParamsFromUrl/:name/:age/:gender", func(c *gin.Context) {
		p := ResponseParams{}
		err := c.ShouldBind(&p) ////Here if it uses c.ShouldBindUri(), it will have an error. So change it to c.ShouldBindUri() fix it.
		p.Name = c.Param("name")
		p.Age, _ = strconv.Atoi(c.Param("age"))
		p.Gender = c.Param("gender")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "Something wrong!",
				"data": gin.H{},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "success",
				"data": p,
			})
		}
	})
}
