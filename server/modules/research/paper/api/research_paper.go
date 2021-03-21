package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/modules/research/paper/model"
	"gin-vue-admin/modules/research/paper/model/dto"
	"gin-vue-admin/modules/research/paper/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Create(c *gin.Context) {
	var R dto.ResearchPaperDto
	_ = c.ShouldBindJSON(&R)
	//if err := utils.Verify(R, utils.RegisterVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	user := &model.ResearchPaper{Title: R.Title}
	err, userReturn := service.Create(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败", zap.Any("err", err))
		response.FailWithDetailed(dto.ResearchPaperResponse{ResearchPaper: userReturn}, "保存失败", c)
	} else {
		response.OkWithDetailed(dto.ResearchPaperResponse{ResearchPaper: userReturn}, "保存成功", c)
	}
}

func Update(c *gin.Context) {
	var user model.ResearchPaper
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, ReqUser := service.Update(user); err != nil {
		global.GVA_LOG.Error("设置失败", zap.Any("err", err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "设置成功", c)
	}
}

func Get(c *gin.Context) {
	var detail model.ResearchPaper
	_ = c.ShouldBindQuery(&detail)
	if err := utils.Verify(detail, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	println(detail.ID)
	if err, resysDictionaryDetail := service.Get(detail.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"detail": resysDictionaryDetail}, "查询成功", c)
	}
}

func Delete(c *gin.Context) {
	var detail model.ResearchPaper
	_ = c.ShouldBindJSON(&detail)
	if err := service.Delete(detail); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetUserInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
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