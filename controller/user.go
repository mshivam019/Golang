package controller

import (
	"gin/basic/dao"
	"gin/basic/db"
	"gin/basic/model"
	"gin/basic/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)
// @Summary Sign up a new user
// @Description Register a new user
// @ID sign-up
// @Accept json
// @Produce json
// @Param user body model.User true "User details"
// @Success 200 {object} model.User
// @Router /signup [post]
func SignUp(c *gin.Context) {
	//parse user data from body
	//call save user form dao
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	err = dao.SaveUser(db.GetDBConn(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, user)
}
// @Summary Log in a user
// @Description Log in a user and get an authentication token
// @ID log-in
// @Accept json
// @Produce json
// @Param user body model.User true "User details"
// @Success 200 {string} string
// @Router /login [post]
func LogIn(c *gin.Context) {
	/*
		1. parse user data from body
		2. match user details from db - dao.login
		3. generate token and send
	*/

	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	err = dao.LogIn(db.GetDBConn(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	//now you verified alreday that user is g[d i.e, exist

	token, err := utils.GenerateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, token)

}
// @Summary Log out a user
// @Description Log out a user and invalidate the token
// @ID log-out
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {string} string
// @Router /logout [post]
func LogOut(c *gin.Context) {
    // Extract the token from the Authorization header
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        c.JSON(http.StatusBadRequest, "Authorization header missing")
        return
    }

    // Remove "Bearer " prefix from the token
    userToken := authHeader[7:]

    // Invalidate the token
    utils.InvalidateToken(userToken)

    c.JSON(http.StatusOK, "Logged out successfully")
}

