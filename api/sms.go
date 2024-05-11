package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/pkg/output"
	"github.com/provider-go/pkg/sms"
	"github.com/provider-go/pkg/sms/typesms"
	"github.com/provider-go/pkg/util"
	"github.com/provider-go/sms/global"
	"github.com/provider-go/sms/models"
	"strconv"
)

func SendCodeByAli(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	phone := output.ParamToString(json["phone"])
	code := strconv.Itoa(util.GeFourRandInt())
	// 添加发送记录
	err := models.CreateSMSLog("code", "ali", phone, code)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	}
	// 使用第方法短信通道
	accessKeyId, err := global.SMCC.GetConfig("sms.ali.AccessKeyId")
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "获取 sms.ali.AccessKeyId 错误~")
	}
	accessKeySecret, err := global.SMCC.GetConfig("sms.ali.AccessKeySecret")
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "获取 sms.ali.AccessKeySecret 错误~")
	}
	endpoint, err := global.SMCC.GetConfig("sms.ali.Endpoint")
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "获取 sms.ali.Endpoint 错误~")
	}
	signName, err := global.SMCC.GetConfig("sms.ali.SignName")
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "获取 sms.ali.SignName 错误~")
	}
	templateCode, err := global.SMCC.GetConfig("sms.ali.TemplateCode")
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "获取 sms.ali.TemplateCode 错误~")
	}
	cfg := typesms.ConfigSMS{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Endpoint:        endpoint,
		SignName:        signName,
		TemplateCode:    templateCode,
	}
	instance := sms.NewSMS("ali", cfg)
	err = instance.Send(phone, code)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "Send 错误~")
	}
	// 验证码保存redis 10分钟
	global.Cache.Set(phone, code, 600)
	output.ReturnSuccessResponse(ctx, nil)
}

func SendCodeBySandbox(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	phone := output.ParamToString(json["phone"])
	code := "6666"
	// 添加发送记录
	err := models.CreateSMSLog("code", "sandbox", phone, code)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	}
	// 验证码保存redis 10分钟
	global.Cache.Set(phone, code, 600)

	output.ReturnSuccessResponse(ctx, nil)
}

func LogList(ctx *gin.Context) {
	pageSize := output.ParamToInt(ctx.Query("pageSize"))
	pageNum := output.ParamToInt(ctx.Query("pageNum"))

	list, total, err := models.ListSMSLog(pageSize, pageNum)

	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		res := make(map[string]interface{})
		res["records"] = list
		res["total"] = total
		output.ReturnSuccessResponse(ctx, res)
	}
}
