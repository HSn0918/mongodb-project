package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/book"
	"github.com/flipped-aurora/gin-vue-admin/server/router/college"
	"github.com/flipped-aurora/gin-vue-admin/server/router/ehcart"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	College college.RouterGroup
	Book    book.RouterGroup
	Ehcart  ehcart.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
