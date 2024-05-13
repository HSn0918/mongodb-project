package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/hsn0918/mongodb_project/req"
	"github.com/hsn0918/mongodb_project/resp"
	"github.com/hsn0918/mongodb_project/svc"
)

var SvcContext *svc.ServiceContext

func init() {
	SvcContext = svc.NexServiceContext()
}
func CollegeTeacher(c *gin.Context) {
	var req *req.CollegeTeacherReq
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		resp.FailWithDetailed("参数错误", err.Error(), c)
		return
	}
	// 调用svcContext
	collegeTeacher, err := SvcContext.CollegeModel.FindAllCollegeTeacherID(req.College)
	if err != nil {
		resp.FailWithDetailed("查询失败", err.Error(), c)
		return
	}
	var teacherIDs []string
	for _, v := range *collegeTeacher {
		teacherIDs = append(teacherIDs, v.TeacherID)
	}
	teacher, err := SvcContext.TeacherModel.FindOneByTeacherIds(teacherIDs)
	if err != nil {
		resp.FailWithDetailed("查询失败", err.Error(), c)
		return
	}
	var respTeacher []resp.TeacherResp
	for _, v := range *teacher {
		respTeacher = append(respTeacher, resp.TeacherResp{
			TeacherID: v.TeacherID,
			Name:      v.Name,
			Contact:   v.Contact,
		})
	}
	resp.OkWithData(respTeacher, c)
}
func CollegeName(c *gin.Context) {
	collegeName, err := SvcContext.CollegeModel.FindAllCollegeName()
	if err != nil {
		resp.FailWithDetailed("查询失败", err.Error(), c)
		return
	}
	resp.OkWithData(collegeName, c)
}
func Dean(c *gin.Context) {
	var req *req.DeanReq
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		resp.FailWithDetailed("参数错误", err.Error(), c)
		return
	}
	// 调用svcContext
	dean, err := SvcContext.CollegeModel.FindDean(req.College)
	if err != nil {
		resp.FailWithDetailed("查询失败", err.Error(), c)
		return
	}
	resp.OkWithData(dean, c)
}
