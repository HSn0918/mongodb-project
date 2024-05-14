import service from '@/utils/request'


export const getEchart = ()=>{
  return service({
    url:"/sysBook/echart",
    method:"get"
  })
}
export const insertBook = (data) => {
  return service({
    url:"/sysBook/InsertBook",
    method:"post",
    data
  })
}

export const getBooksList= (data)=>{
  return service({
    url:"/sysBook/FindBookByTeacherName",
    method:"post",
    data
  })
}
// @Tags SysBook
// @Summary 创建教材
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBook true "创建教材"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysBook/createSysBook [post]
export const createSysBook = (data) => {
  return service({
    url: '/sysBook/createSysBook',
    method: 'post',
    data
  })
}

// @Tags SysBook
// @Summary 删除教材
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBook true "删除教材"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysBook/deleteSysBook [delete]
export const deleteSysBook = (params) => {
  return service({
    url: '/sysBook/deleteSysBook',
    method: 'delete',
    params
  })
}

// @Tags SysBook
// @Summary 批量删除教材
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除教材"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysBook/deleteSysBook [delete]
export const deleteSysBookByIds = (params) => {
  return service({
    url: '/sysBook/deleteSysBookByIds',
    method: 'delete',
    params
  })
}

// @Tags SysBook
// @Summary 更新教材
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBook true "更新教材"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysBook/updateSysBook [put]
export const updateSysBook = (data) => {
  return service({
    url: '/sysBook/updateSysBook',
    method: 'put',
    data
  })
}

// @Tags SysBook
// @Summary 用id查询教材
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysBook true "用id查询教材"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysBook/findSysBook [get]
export const findSysBook = (params) => {
  return service({
    url: '/sysBook/findSysBook',
    method: 'get',
    params
  })
}

// @Tags SysBook
// @Summary 分页获取教材列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取教材列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysBook/getSysBookList [get]
export const getSysBookList = (params) => {
  return service({
    url: '/sysBook/getSysBookList',
    method: 'get',
    params
  })
}
