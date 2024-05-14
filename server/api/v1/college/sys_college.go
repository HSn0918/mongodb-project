package college

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/college"
	collegeReq "github.com/flipped-aurora/gin-vue-admin/server/model/college/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysCollegeApi struct {
}

var sysCollegeService = service.ServiceGroupApp.CollegeServiceGroup.SysCollegeService

func (sysCollegeApi *SysCollegeApi) FindAllTeacher(c *gin.Context) {
	var req college.TeachersReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var teachers []*college.Teacher
	teachers, err = sysCollegeService.FindAllTeacher(req.Name)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(teachers, c)
}

func (sysCollegeApi *SysCollegeApi) FindDean(c *gin.Context) {
	var req college.DeanReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var dean *college.Dean

	dean, err = sysCollegeService.FindDean(req.Name)
	response.OkWithData(dean, c)
}
func (sysCollegeApi *SysCollegeApi) FindAllCollegeName(c *gin.Context) {
	var college []string
	college, err := sysCollegeService.FindAllCollegeName()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(college, "查询成功", c)
}

// CreateSysCollege 创建学院
// @Tags SysCollege
// @Summary 创建学院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body college.SysCollege true "创建学院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysCollege/createSysCollege [post]
func (sysCollegeApi *SysCollegeApi) CreateSysCollege(c *gin.Context) {
	var sysCollege college.SysCollege
	err := c.ShouldBindJSON(&sysCollege)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysCollegeService.CreateSysCollege(&sysCollege); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysCollege 删除学院
// @Tags SysCollege
// @Summary 删除学院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body college.SysCollege true "删除学院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysCollege/deleteSysCollege [delete]
func (sysCollegeApi *SysCollegeApi) DeleteSysCollege(c *gin.Context) {
	ID := c.Query("ID")
	if err := sysCollegeService.DeleteSysCollege(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysCollegeByIds 批量删除学院
// @Tags SysCollege
// @Summary 批量删除学院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysCollege/deleteSysCollegeByIds [delete]
func (sysCollegeApi *SysCollegeApi) DeleteSysCollegeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := sysCollegeService.DeleteSysCollegeByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysCollege 更新学院
// @Tags SysCollege
// @Summary 更新学院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body college.SysCollege true "更新学院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysCollege/updateSysCollege [put]
func (sysCollegeApi *SysCollegeApi) UpdateSysCollege(c *gin.Context) {
	var sysCollege college.SysCollege
	err := c.ShouldBindJSON(&sysCollege)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysCollegeService.UpdateSysCollege(sysCollege); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysCollege 用id查询学院
// @Tags SysCollege
// @Summary 用id查询学院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query college.SysCollege true "用id查询学院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysCollege/findSysCollege [get]
func (sysCollegeApi *SysCollegeApi) FindSysCollege(c *gin.Context) {
	ID := c.Query("ID")
	if resysCollege, err := sysCollegeService.GetSysCollege(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysCollege": resysCollege}, c)
	}
}

// GetSysCollegeList 分页获取学院列表
// @Tags SysCollege
// @Summary 分页获取学院列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query collegeReq.SysCollegeSearch true "分页获取学院列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysCollege/getSysCollegeList [get]
func (sysCollegeApi *SysCollegeApi) GetSysCollegeList(c *gin.Context) {
	var pageInfo collegeReq.SysCollegeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysCollegeService.GetSysCollegeInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetSysCollegePublic 不需要鉴权的学院接口
// @Tags SysCollege
// @Summary 不需要鉴权的学院接口
// @accept application/json
// @Produce application/json
// @Param data query collegeReq.SysCollegeSearch true "分页获取学院列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysCollege/getSysCollegePublic [get]
func (sysCollegeApi *SysCollegeApi) GetSysCollegePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的学院接口信息",
	}, "获取成功", c)
}
