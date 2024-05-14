import service from '@/utils/request'



export const findAllTeacher = (data)=>{
  return service({
    url:'/sysCollege/findAllTeacher',
    method:'post',
    data
  })

}


export const findCollegeDean = (data) => {
  return service({
    url: '/sysCollege/findCollegeDean',
    method: 'post',
    data
  })
}


export  const findAllcollegeName = () => {
  return service({
    url: '/sysCollege/findAllcollegeName',
    method: 'get',
  })
}
// @Tags SysCollege
// @Summary 创建学院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysCollege true "创建学院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysCollege/createSysCollege [post]
export const createSysCollege = (data) => {
  return service({
    url: '/sysCollege/createSysCollege',
    method: 'post',
    data
  })
}

// @Tags SysCollege
// @Summary 删除学院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysCollege true "删除学院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysCollege/deleteSysCollege [delete]
export const deleteSysCollege = (params) => {
  return service({
    url: '/sysCollege/deleteSysCollege',
    method: 'delete',
    params
  })
}

// @Tags SysCollege
// @Summary 批量删除学院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除学院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysCollege/deleteSysCollege [delete]
export const deleteSysCollegeByIds = (params) => {
  return service({
    url: '/sysCollege/deleteSysCollegeByIds',
    method: 'delete',
    params
  })
}

// @Tags SysCollege
// @Summary 更新学院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysCollege true "更新学院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysCollege/updateSysCollege [put]
export const updateSysCollege = (data) => {
  return service({
    url: '/sysCollege/updateSysCollege',
    method: 'put',
    data
  })
}

// @Tags SysCollege
// @Summary 用id查询学院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysCollege true "用id查询学院"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysCollege/findSysCollege [get]
export const findSysCollege = (params) => {
  return service({
    url: '/sysCollege/findSysCollege',
    method: 'get',
    params
  })
}

// @Tags SysCollege
// @Summary 分页获取学院列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取学院列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysCollege/getSysCollegeList [get]
export const getSysCollegeList = (params) => {
  return service({
    url: '/sysCollege/getSysCollegeList',
    method: 'get',
    params
  })
}
