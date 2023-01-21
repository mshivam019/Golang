package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

)


var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

type PostPayload struct {
	Name string `json:"name"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	var s1 string
	var flag int


	// post := map[string]string{
	// 	"rahul": "jkghvkjhf",
	// }
	// r := gin.Default() //application will be configured on port 8080
	// pingFunc := func(c *gin.Context) {

	// 	key := map[string]any{
	// 		"ping": "pingu",
	// 	}
	// 	c.Keys = key
	// 	res := gin.H{
	// 		"message":      "pong",
	// 		"contextvalue": c.Keys["ping"],
	// 	}
	// 	c.JSON(http.StatusUnauthorized, res)
	// }
	// pongFunc := func(c *gin.Context) {
	// 	res := gin.H{
	// 		"message":      "pong",
	// 		"contextvalue": c.Keys["ping"],
	// 	}
	// 	log.Println(c.Keys["ping"])
	// 	c.JSON(http.StatusUnauthorized, res)
	// }
	// dingFunc := func(c *gin.Context) {
	// 	res := gin.H{
	// 		"message": "pong",
	// 	}
	// 	c.JSON(http.StatusUnauthorized, res)
	// }
	// dongFunc := func(c *gin.Context) {
	// 	res := gin.H{
	// 		"message": "pong",
	// 	}
	// 	c.JSON(http.StatusUnauthorized, res)
	// }
	// r.GET("/ping", pingFunc) //http://localhost:8080/ping
	// r.GET("/pong", pongFunc)
	// r.GET("/ding", dingFunc)
	// r.GET("/dong", dongFunc)

	post := map[string]string{
		"rahul": "Present",
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
	getPost := func(c *gin.Context) {
		res := gin.H{
			"message": post,
		}
		c.JSON(http.StatusUnauthorized, res)
	}

	r.GET("/post", getPost)
	createPost := func(c *gin.Context) {

		var p1 PostPayload
		err := c.ShouldBindJSON(&p1)
		if err != nil {
			log.Print("got error", err)
		}
		post[p1.Name] = "Present"
		res := gin.H{
			"message": p1,
		}
		c.JSON(http.StatusOK, res)
	}
	r.POST("/post", createPost)
	deletePost := func(c *gin.Context) {

		//key := c.Params.ByName("id")
		key := c.Query("id")
		//to delete the key value pair from a map via key
		delete(post, key)
		res := gin.H{
			"message": key + " Deleted",
		}
		c.JSON(http.StatusOK, res)
	}
	r.DELETE("/post", deletePost)
	updatePostById := func(c *gin.Context) {
		id := c.Query("id")

		post[id] = "absent"
		res := gin.H{
			"message": id + " marked absent",
		}
		c.JSON(http.StatusOK, res)
	}
	r.PATCH("/post", updatePostById) //data will be sent via query
	updateValuebyId := func(c *gin.Context) {
		var p1 PostPayload
		val := c.Query("val")
		err := c.ShouldBindJSON(&p1)
		if err != nil {
			log.Print("got error", err)
		}
		post[p1.Name] = val
		res := gin.H{
			"message": p1.Name + " updated",
		}
		c.JSON(http.StatusOK, res)
	}
	getPostById := func(c *gin.Context) {

		id := c.Query("id")

		val := post[id]
		res := gin.H{
			"message": id + " " + val,
		}
		c.JSON(http.StatusOK, res)
	}
	isogramc := func(c *gin.Context) {
		s := NewSet()
		flag=0
		word := c.Query("word")

		for _, character := range word {
			if s.Contains(string(character)) {
				s1 = " is not an isogram"
				flag = 1
			} else {
				s.Add(string(character))
			}
		}
		if flag == 0 {
			s1 = " is an isogram"
		}
		res := gin.H{
			"message": word+s1,
		}
		maps.Clear(s)
		c.JSON(http.StatusOK, res)
	}
	r.GET("/isogram", isogramc)
	r.PUT("/post", updateValuebyId) //data will be sent via payload and query
	r.GET("/postbyid", getPostById)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run(fmt.Sprintf(":%s", port))
}

//get all posts localhost:8080/post
/// delete post by id:-
// url: http://localhost:8080/post?id=rahul query
// url: http://localhost:8080/post/rahul param
// request method: delete
/* add post POST
url: localhost:8080/post
and json as
 {
	   "name":"shivam"
  }
*/
// patch  value changes to absent
// url: localhost:8080/post?id=rahul
// get post by id:-
// url: http://localhost:8080/postbyid?id=rahul
// request method: get

//http://localhost:8080/post
/* put call to change the value of post to anything passed in query
url: localhost:8080/post?val=check
and json as
{
    "name": "shivam"
}
*/
