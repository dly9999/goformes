package router

import (
	//	"mesnewcloud/config"
	"goformescloud/mesnewcloud/server/Demotest"
	"goformescloud/mesnewcloud/server/ProOrdtemp"
	"goformescloud/mesnewcloud/server/QualityDefectQuery"
	"goformescloud/mesnewcloud/server/ReportBrushrecard"

	"goformescloud/mesnewcloud/server/EmployeeOutput"
	"goformescloud/mesnewcloud/server/ProcfuctForms"
	"goformescloud/mesnewcloud/server/ReportOrderState"
	"goformescloud/mesnewcloud/server/ReportReworkRecord"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())   //日志中间件
	r.Use(gin.Recovery()) //恢复中间件
	gin.SetMode(gin.DebugMode)
	test := r.Group("/test")
	{
		test.GET("/t", Demotest.TestDemo)
	}
	test2 := r.Group("/ProOrdtemp")
	{
		test2.GET("/GetOrdertemp", ProOrdtemp.GetOrdertemp)
	}
	test3 := r.Group("/report")
	{
		test3.GET("/brushcardrecord", Report_BrushCard.Method_QueryBrushCardList)
		test3.GET("/orderstateinfo", Report_OrderState.Method_OrderStateInfo)
		test3.GET("/reworkrecord", Report_ReworkRecord.Method_ReworkRecord)
	}
	prdcform := r.Group("/prdcform")
	{
		prdcform.POST("/cut", ProcfuctForms.Post_Cuforms)
		prdcform.POST("/sew", ProcfuctForms.Post_Sewforms)
	}
	test4 := r.Group("/QualityDefectQuery")
	{
		test4.GET("/GetOrderDefect", QualityDefectQuery.GetOrderDefect)
	}
	test5 := r.Group("/EmployeeOutput")
	{
		test5.POST("/PostEmployeeOutPut", EmployeeOutput.PostEmployeeOutPut)
	}
	return r
}
