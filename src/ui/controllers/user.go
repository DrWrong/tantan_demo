// user related contoller layer
// Author: drwrong<yuhangchaney@gmail.com>

package controllers

import (
	"github.com/DrWrong/monica/core"
	"io/ioutil"

	"ui/models"
)

// get user list contoller
func userListContoller(c *core.Context) {
	responses, err := models.GetAllUsers()
	if err != nil {
		c.ErrorResponse(err)
	} else {
		c.RenderJson(responses)
	}
}

// create user contoller
func userCreateContoller(c *core.Context) {
	data, err := ioutil.ReadAll(c.Body)
	defer c.Body.Close()
	if err != nil {
		c.ErrorResponse(err)
		return
	}
	res, err := models.CreateUser(data)
	if err != nil {
		c.ErrorResponse(err)
	} else {
		c.RenderJson(res)
	}
}

// entern contoller
func UserController(c *core.Context) {
	if c.Method == "POST" {
		userCreateContoller(c)
	} else {
		userListContoller(c)
	}
}
