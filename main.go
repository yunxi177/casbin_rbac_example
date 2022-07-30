package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	a, _ := gormadapter.NewAdapter("mysql", "username:password@tcp(host:3306)/", "your policy table name") // Your driver and data source.
	text := `
		[request_definition]
		r = sub, dom, obj, act
		
		[policy_definition]
		p = sub, dom, obj, act
		
		[role_definition]
		g = _, _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act
		`
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(err)
	}
	e, _ := casbin.NewEnforcer(m, a)
	// 初始化 policy
	e.LoadPolicy()
	e.AddPolicy("alice", "domain1", "/alice/data1", "read")
	e.AddPolicy("alice", "domain1", "/alice/datadata1", "write")
	e.AddPolicy("user1", "domain1", "/alice/user/1", "read")
	e.AddPolicy("user2", "domain1", "/alice/user/1", "write")
	e.AddGroupingPolicy("user2", "user1", "domain1")
	e.AddGroupingPolicy("user1", "alice", "domain1")

	e.SavePolicy()
	//验证
	e.LoadPolicy()
	success, _ := e.Enforce("user2", "domain1", "/alice/user/1", "write")
	fmt.Println(success)
}
