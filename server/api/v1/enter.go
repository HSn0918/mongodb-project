package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/book"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/college"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/ehcart"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	CollegeApiGroup college.ApiGroup
	BookApiGroup    book.ApiGroup
	EhcartApiGroup  ehcart.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
