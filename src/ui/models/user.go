// model layer for the application
// @Author DrWrong<yuhangchaney@gmail.com>
package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
)

// db layer orm for user
type User struct {
	Id   int64
	Name string
}

// define the response of get User
type UserResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (u *User) toResponse() *UserResponse {
	return &UserResponse{
		Id:   fmt.Sprintf("%d", u.Id),
		Name: u.Name,
		Type: "user",
	}
}

// TODO: Validate user this can be performed by beego
// But in this demo we leave it
func (u *User) validate() error {
	return nil
}

// get all users for list
// in deployed environment this should be cached to reduce the db query
// however in this demo we just query db directly
func GetAllUsers() ([]*UserResponse, error) {
	o := orm.NewOrm()
	var users []*User
	if _, err := o.QueryTable("user").All(&users); err != nil {
		return nil, err
	}

	userResponses := make([]*UserResponse, 0, len(users))
	for _, user := range users {
		userResponses = append(
			userResponses, user.toResponse())
	}
	return userResponses, nil
}

// create user
func CreateUser(data []byte) (*UserResponse, error) {
	user := new(User)
	if err := json.Unmarshal(data, user); err != nil {
		return nil, err
	}

	if err := user.validate(); err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	if _, err := o.Insert(user); err != nil {
		return nil, err
	}
	return user.toResponse(), nil

}

func init() {
	orm.RegisterModel(new(User))
}
