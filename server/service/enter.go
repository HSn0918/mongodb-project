package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/book"
	"github.com/flipped-aurora/gin-vue-admin/server/service/college"
	"github.com/flipped-aurora/gin-vue-admin/server/service/ehcart"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	CollegeServiceGroup college.ServiceGroup
	BookServiceGroup    book.ServiceGroup
	EhcartServiceGroup  ehcart.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
