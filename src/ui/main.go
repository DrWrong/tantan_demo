package main

import (
	"github.com/DrWrong/monica"
	"github.com/DrWrong/monica/core"

	"ui/controllers"
)

// customized config this is a callback function for the monica framework
// it will be invoked after the monica is inited
func customizedConfig() {
	monica.InitPostgres()
	routerConfigure()
}

// route config
func routerConfigure() {
	core.Handle(`^/users$`, controllers.UserController)
	core.Handle(`^/users/(?P<user_id>\d+)/relationships$`, controllers.RelationshipListController)
	core.Handle(`^/users/(?P<user_id>\d+)/relationships/(?P<other_user_id>\d+)`, controllers.RelationshipOperationController)
}

func main() {
	monica.BootStrap(customizedConfig)
}
