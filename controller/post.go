package controller

import (
	"fmt"
	"gin/basic/dao"
	"gin/basic/db"
	"gin/basic/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// acting as db for as
var post = map[string]string{
	"Shivam": "Present",
}
// @Summary Get a greeting message
// @Description Get a greeting message from the server
// @ID get-greeting
// @Produce json
// @Success 200 {string} string "Greeting message"
// @Router / [get]
func HelloWorld(c *gin.Context) {

	res := gin.H{
		"message": "Hello there!",
	}
	c.JSON(http.StatusOK, res)
}
// @Summary Get all posts
// @Description Get a list of all posts
// @ID get-all-posts
// @Param token header string true "token"
// @Produce json
// @Success 200 {array} model.Post "List of posts"
// @Router /post [get]
func GetPost(c *gin.Context) {

	res := gin.H{
		"message": "",
	}
	listOfPosts, err := dao.GetPosts(db.GetDBConn())
	if err != nil {
		log.Print("got error", err)
		res = gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
	} else {
		res = gin.H{
			"message": listOfPosts,
		}
		c.JSON(http.StatusOK, res)
	}
}
// @Summary Delete a post by ID
// @Description Delete a post by providing its ID
// @ID delete-post
// @Param id path string true "Post ID"
// @Param token header string true "token"
// @Produce json
// @Success 200 {string} string "Post deleted"
// @Success 404 {string} string "Post not found"
// @Router /post/{id} [delete]
func DeletePostByID(c *gin.Context) {

	key := c.Params.ByName("id")
	res := gin.H{
		"message": "",
	}
	post, err := dao.GetPostByID(db.GetDBConn(), key)
	log.Println(fmt.Sprint(post.Id), "post.Id, key", len(fmt.Sprint(post.Id)), len(key), fmt.Sprint(post.Id) == key)
	if err != nil {
		log.Print("got error", err)
		res = gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
	} else if fmt.Sprint(post.Id) == key { //means post exist
		post, err := dao.DeletePostById(db.GetDBConn(), key)
		if err != nil {

			log.Print("got error", err)
			res = gin.H{
				"message": err.Error(),
			}
			c.JSON(http.StatusInternalServerError, res)
		} else {
			res = gin.H{
				"post deleted": post,
			}
			c.JSON(http.StatusOK, res)
		}
	} else {
		res = gin.H{
			"post exists": false,
		}
		c.JSON(http.StatusNotFound, res)
	}
	//out=fmt.Sprint(`hey my name is %s. I am from %s`, "rahul", "here")
}

// @Summary Create a new post
// @Description Create a new post with title and content
// @ID create-post
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param post body model.Post true "Post data"
// @Success 200 {string} string "Post created"
// @Router /post [post]
func CreatePost(c *gin.Context) {
	//log.Println("*********", c, "*********")
	var p1 model.Post

	err := c.ShouldBindJSON(&p1)
	if err != nil {
		log.Print("got error", err)
	}
	res := gin.H{
		"message": "",
	}
	_, err = dao.CreatePost(db.GetDBConn(), &p1)
	if err != nil {
		log.Print("got error", err)
		res = gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
	} else {
		res = gin.H{
			"message": "Done",
		}
		c.JSON(http.StatusOK, res)
	}

}
// @Summary Get a post by ID
// @Description Get a post by providing its ID
// @ID get-post-by-id
// @Param id path string true "Post ID"
// @Param token header string true "token"
// @Produce json
// @Success 200 {string} string "Post details"
// @Router /post/{id} [get]
func GetPostById(c *gin.Context) {
	key := c.Params.ByName("id")
	res := gin.H{
		"message": "",
	}
	post, err := dao.GetPostByID(db.GetDBConn(), key)
	if err != nil {
		log.Print("got error", err)
		res = gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
	} else {
		res = gin.H{
			"message": post,
		}
		c.JSON(http.StatusOK, res)
	}
}
// @Summary Update a post
// @Description Update a post by providing updated data
// @ID update-post
// @Accept json
// @Produce json
// @Param post body model.Post true "Updated post data"
// @Param token header string true "token"
// @Success 200 {string} string "Post updated"
// @Router /post [put]
func UpdatePost(c *gin.Context) {

	// key := c.Params.ByName("id")
	res := gin.H{
		"message": "",
	}
	var p1 model.Post

	err := c.ShouldBindJSON(&p1)
	if err != nil {
		log.Print("got error", err)
	}
	_, err = dao.UpdatePost(db.GetDBConn(), &p1)
	if err != nil {
		log.Print("got error", err)
		res = gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
	} else {
		res = gin.H{
			"message": "Done",
		}
		c.JSON(http.StatusOK, res)
	}
	// post, err := dao.GetPostByID(db.GetDBConn(), key)
	// log.Println(fmt.Sprint(post.Id), "post.Id, key", len(fmt.Sprint(post.Id)), len(key), fmt.Sprint(post.Id) == key)
	// if err != nil {
	// 	log.Print("got error", err)
	// 	res = gin.H{
	// 		"message": err.Error(),
	// 	}
	// 	c.JSON(http.StatusInternalServerError, res)
	// } else if fmt.Sprint(post.Id) == key { //means post exist
	// 	post, err := dao.DeletePostById(db.GetDBConn(), key)
	// 	if err != nil {

	// 		log.Print("got error", err)
	// 		res = gin.H{
	// 			"message": err.Error(),
	// 		}
	// 		c.JSON(http.StatusInternalServerError, res)
	// 	} else {
	// 		res = gin.H{
	// 			"post deleted": post,
	// 		}
	// 		c.JSON(http.StatusOK, res)
	// 	}
	// } else {
	// 	res = gin.H{
	// 		"post exists": false,
	// 	}
	// 	c.JSON(http.StatusNotFound, res)
	// }
	//out=fmt.Sprint(`hey my name is %s. I am from %s`, "rahul", "here")
}
