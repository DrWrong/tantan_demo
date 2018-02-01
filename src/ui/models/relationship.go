// relationship relataed models
// @Author DrWrong<yuhangchaney@gmail.com>
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

const (
	// relation ship state const use 0 to stand for invalid
	undifined RelationState = iota
	disliked
	liked
	matched
)

// a state enum to string translate map
var stateTranslateMap map[RelationState]string = map[RelationState]string{
	1: "disliked",
	2: "liked",
	3: "matched",
}

// here we make a alias of int8 to add a to string function
// usually this is atomic down in `thrift` layer in deployed mode.
type RelationState int8

func (r RelationState) toString() string {
	return stateTranslateMap[r]
}

// db layer orm for relationship
type Relationship struct {
	Id    int64
	Uid1  int64
	Uid2  int64
	State RelationState
}

// relationship response
type RelationshipResponse struct {
	UserId string `json:"user_id"`
	State  string `json:"state"`
	Type   string `json:"type"`
}

// translate a relationship to response
func (r *Relationship) toResponse() *RelationshipResponse {
	return &RelationshipResponse{
		UserId: fmt.Sprintf("%d", r.Uid2),
		State:  r.State.toString(),
		Type:   "relationship",
	}
}

// insert into db or upate on conflict
func (r *Relationship) insertOrUpdate(o orm.Ormer) error {
	_, err := o.Raw(
		`INSERT INTO relationship(uid1, uid2, state)
		 Values (?, ?, ?)
		 ON CONFLICT(uid1, uid2) DO UPDATE set state = ?`,
		r.Uid1,
		r.Uid2,
		r.State, r.State).Exec()

	return err
}

// list all the relationship related to uid1
func ListRelationshipForUser(uid int64) ([]*RelationshipResponse, error) {
	o := orm.NewOrm()
	var relationships []*Relationship
	if _, err := o.QueryTable("relationship").Filter("uid1", uid).All(&relationships); err != nil {
		return nil, err
	}

	res := make([]*RelationshipResponse, 0, len(relationships))

	for _, r := range relationships {
		res = append(res, r.toResponse())
	}
	return res, nil
}

type RelationOperateForm struct {
	UserId      int64
	OtherUserId int64
	State       string `json:"state"`
}

// verify user input data
// this method will auto invoked by the c.Bind
func (form *RelationOperateForm) Valid(v *validation.Validation) {
	// verify state
	if form.State != "liked" && form.State != "disliked" {
		v.SetError("State", "not proper state")
	}

	// verify two userid
	o := orm.NewOrm()
	num, _ := o.QueryTable("user").Filter("id__in", form.UserId, form.OtherUserId).Count()
	if num != 2 {
		v.SetError("UserId", "wrong userid")
		v.SetError("OtherUserId", "wrong other userid")
	}
}

func (form *RelationOperateForm) Operate() (res *RelationshipResponse, err error) {
	var r *Relationship
	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		return nil, err
	}
	// leared from docker
	// defer a functon to process rollback
	rollback := true
	defer func() {
		if rollback {
			o.Rollback()
		} else {
			o.Commit()
		}
	}()

	switch form.State {
	case "liked":
		r, err = form.doLiked(o)
	case "disliked":
		r, err = form.doDisLiked(o)
	}
	if err != nil {
		return nil, err
	}
	if err := r.insertOrUpdate(o); err != nil {
		return nil, err
	}

	rollback = false
	return r.toResponse(), nil
}

func (form *RelationOperateForm) doLiked(o orm.Ormer) (*Relationship, error) {
	// this is a trick well performed in mysql howerver I don't know wheather it will work in pg
	// in mysql transactions all the locks will only be rellease at the end of the transaction. so I don't have to lock any row mannuly, just update will ensure to data consistency

	// instead of select the rows first, I use the update method and examine effected row to judge wheather relationship had been built

	// fist we examine wheather uid2 have alerady liked uid1 if so the state sholud be set to matched
	num, err := o.QueryTable("relationship").Filter(
		"uid1", form.OtherUserId).Filter(
		"uid2", form.UserId).Filter(
		"state", liked).Update(
		orm.Params{
			"state": matched,
		})
	if err != nil {
		return nil, err
	}

	// if she like me then we married
	// else I could only like her
	state := liked
	if num > 0 {
		state = matched
	}

	r := &Relationship{
		Uid1:  form.UserId,
		Uid2:  form.OtherUserId,
		State: state,
	}


	return r, nil
}

func (form *RelationOperateForm) doDisLiked(o orm.Ormer) (*Relationship, error) {
	// broken up
	_, err := o.QueryTable("relationship").Filter(
		"uid1", form.OtherUserId).Filter(
		"uid2", form.UserId).Filter(
		"state", matched).Update(
		orm.Params{
			"state": liked,
		})
	if err != nil {
		return nil, err
	}

	r := &Relationship{
		Uid1:  form.UserId,
		Uid2:  form.OtherUserId,
		State: disliked,
	}
	return r, nil


}

func init() {
	orm.RegisterModel(new(Relationship))
}
