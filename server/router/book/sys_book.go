package book

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysBookRouter struct {
}

// InitSysBookRouter 初始化 教材 路由信息
func (s *SysBookRouter) InitSysBookRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sysBookRouter := Router.Group("sysBook").Use(middleware.OperationRecord())
	sysBookRouterWithoutRecord := Router.Group("sysBook")
	sysBookRouterWithoutAuth := PublicRouter.Group("sysBook")

	var sysBookApi = v1.ApiGroupApp.BookApiGroup.SysBookApi
	{
		sysBookRouter.POST("createSysBook", sysBookApi.CreateSysBook)             // 新建教材
		sysBookRouter.DELETE("deleteSysBook", sysBookApi.DeleteSysBook)           // 删除教材
		sysBookRouter.DELETE("deleteSysBookByIds", sysBookApi.DeleteSysBookByIds) // 批量删除教材
		sysBookRouter.PUT("updateSysBook", sysBookApi.UpdateSysBook)              // 更新教材
	}
	{
		sysBookRouterWithoutRecord.GET("findSysBook", sysBookApi.FindSysBook)       // 根据ID获取教材
		sysBookRouterWithoutRecord.GET("getSysBookList", sysBookApi.GetSysBookList) // 获取教材列表
	}
	{
		sysBookRouterWithoutAuth.GET("getSysBookPublic", sysBookApi.GetSysBookPublic) // 获取教材列表
		sysBookRouterWithoutAuth.POST("FindBookByTeacherName", sysBookApi.FindBookByTeacherName)
		sysBookRouterWithoutAuth.POST("InsertBook", sysBookApi.InsertBook)
		sysBookRouterWithoutAuth.POST("UpdateBook", sysBookApi.UpdateBookById)
		sysBookRouterWithoutAuth.GET("echart", sysBookApi.EChart)
	}
}
