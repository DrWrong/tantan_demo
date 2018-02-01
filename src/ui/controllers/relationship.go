// relationship related contoller layer
// Author: drwrong<yuhangchaney@gmail.com>
package controllers

import (
	"github.com/DrWrong/monica/core"
	"io/ioutil"
	"encoding/json"

	"ui/models"
)

// list all the relationships
func RelationshipListController(c *core.Context) {
	uid, _ := c.QueryInt("user_id")
	res, err := models.ListRelationshipForUser(int64(uid))
	if err != nil {
		c.ErrorResponse(err)
	}
	c.RenderJson(res)
}


// relationship operation
func RelationshipOperationController(c *core.Context) {
	data, err := ioutil.ReadAll(c.Body)
	if err != nil {
		c.ErrorResponse(err)
	}
	defer c.Body.Close()
	form := new(models.RelationOperateForm)
	if err := json.Unmarshal(data, form); err != nil {
		c.ErrorResponse(err)
	}
	if err := c.Bind(form); err != nil {
		c.ErrorResponse(err)
	}
	res, err := form.Operate()
	if err != nil {
		c.ErrorResponse(err)
	} else {
		c.RenderJson(res)
	}
}
