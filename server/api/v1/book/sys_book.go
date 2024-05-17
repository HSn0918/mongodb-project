package book

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/book"
	bookReq "github.com/flipped-aurora/gin-vue-admin/server/model/book/request"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysBookApi struct {
}

var sysBookService = service.ServiceGroupApp.BookServiceGroup.SysBookService

func (SysBookApi *SysBookApi) UpdateBookById(c *gin.Context) {
	var req book.UpdateBookReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	resp, err := sysBookService.UpdateBookById(req)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithDetailed(resp, "更新成功", c)
}
func (sysBookApi *SysBookApi) EChart(c *gin.Context) {
	var eChart *book.EChart
	eChart, err := sysBookService.EChart()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(eChart, "查询成功", c)
}

func (sysBookApi *SysBookApi) InsertBook(c *gin.Context) {
	var insertReq book.InsertBookReq
	err := c.ShouldBindJSON(&insertReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var insertResp *book.InsertBookResp
	insertResp, err = sysBookService.InsertBook(insertReq)
	if !insertResp.Flag {
		response.OkWithDetailed(insertResp, "插入失败", c)
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(insertResp, "插入成功", c)
}

func (sycBookApi *SysBookApi) FindBookByTeacherName(c *gin.Context) {
	var req book.FindBookReq
	c.ShouldBindJSON(&req)
	var books []*book.FindBookResp
	books, err := sysBookService.FindBookByTeacherName(req.Name)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(books, "查询成功", c)
}

// CreateSysBook 创建教材
// @Tags SysBook
// @Summary 创建教材
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body book.SysBook true "创建教材"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysBook/createSysBook [post]
func (sysBookApi *SysBookApi) CreateSysBook(c *gin.Context) {
	var sysBook book.SysBook
	err := c.ShouldBindJSON(&sysBook)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysBookService.CreateSysBook(&sysBook); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysBook 删除教材
// @Tags SysBook
// @Summary 删除教材
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body book.SysBook true "删除教材"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysBook/deleteSysBook [delete]
func (sysBookApi *SysBookApi) DeleteSysBook(c *gin.Context) {
	ID := c.Query("ID")
	if err := sysBookService.DeleteSysBook(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysBookByIds 批量删除教材
// @Tags SysBook
// @Summary 批量删除教材
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysBook/deleteSysBookByIds [delete]
func (sysBookApi *SysBookApi) DeleteSysBookByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := sysBookService.DeleteSysBookByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysBook 更新教材
// @Tags SysBook
// @Summary 更新教材
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body book.SysBook true "更新教材"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysBook/updateSysBook [put]
func (sysBookApi *SysBookApi) UpdateSysBook(c *gin.Context) {
	var sysBook book.SysBook
	err := c.ShouldBindJSON(&sysBook)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysBookService.UpdateSysBook(sysBook); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysBook 用id查询教材
// @Tags SysBook
// @Summary 用id查询教材
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query book.SysBook true "用id查询教材"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysBook/findSysBook [get]
func (sysBookApi *SysBookApi) FindSysBook(c *gin.Context) {
	ID := c.Query("ID")
	if resysBook, err := sysBookService.GetSysBook(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysBook": resysBook}, c)
	}
}

// GetSysBookList 分页获取教材列表
// @Tags SysBook
// @Summary 分页获取教材列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bookReq.SysBookSearch true "分页获取教材列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBook/getSysBookList [get]
func (sysBookApi *SysBookApi) GetSysBookList(c *gin.Context) {
	var pageInfo bookReq.SysBookSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysBookService.GetSysBookInfoList(pageInfo); err != nil {
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

// GetSysBookPublic 不需要鉴权的教材接口
// @Tags SysBook
// @Summary 不需要鉴权的教材接口
// @accept application/json
// @Produce application/json
// @Param data query bookReq.SysBookSearch true "分页获取教材列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBook/getSysBookPublic [get]
func (sysBookApi *SysBookApi) GetSysBookPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的教材接口信息",
	}, "获取成功", c)
}
