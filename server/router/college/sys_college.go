package college

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysCollegeRouter struct {
}

// InitSysCollegeRouter 初始化 学院 路由信息
func (s *SysCollegeRouter) InitSysCollegeRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sysCollegeRouter := Router.Group("sysCollege").Use(middleware.OperationRecord())
	sysCollegeRouterWithoutRecord := Router.Group("sysCollege")
	sysCollegeRouterWithoutAuth := PublicRouter.Group("sysCollege")

	var sysCollegeApi = v1.ApiGroupApp.CollegeApiGroup.SysCollegeApi
	{
		sysCollegeRouter.POST("createSysCollege", sysCollegeApi.CreateSysCollege)             // 新建学院
		sysCollegeRouter.DELETE("deleteSysCollege", sysCollegeApi.DeleteSysCollege)           // 删除学院
		sysCollegeRouter.DELETE("deleteSysCollegeByIds", sysCollegeApi.DeleteSysCollegeByIds) // 批量删除学院
		sysCollegeRouter.PUT("updateSysCollege", sysCollegeApi.UpdateSysCollege)              // 更新学院
	}
	{
		sysCollegeRouterWithoutRecord.GET("findSysCollege", sysCollegeApi.FindSysCollege)       // 根据ID获取学院
		sysCollegeRouterWithoutRecord.GET("getSysCollegeList", sysCollegeApi.GetSysCollegeList) // 获取学院列表
	}
	{
		sysCollegeRouterWithoutAuth.GET("findAllcollegeName", sysCollegeApi.FindAllCollegeName)
		sysCollegeRouterWithoutAuth.GET("getSysCollegePublic", sysCollegeApi.GetSysCollegePublic) // 获取学院列表
		sysCollegeRouterWithoutAuth.POST("findCollegeDean", sysCollegeApi.FindDean)
		sysCollegeRouterWithoutAuth.POST("findAllTeacher", sysCollegeApi.FindAllTeacher)
	}
}
