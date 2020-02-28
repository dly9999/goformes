package QualityDefectQuery

import (
	"fmt"
	"goformescloud/mesnewcloud/db"
	"goformescloud/mesnewcloud/model"

	"goformescloud/mesnewcloud/pkg"

	"net/http"

	"github.com/Unknwon/com"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func GetOrderDefect(c *gin.Context) {
	tenantId := c.Query("tenantId")
	groupName := c.Query("groupName")
	sort := c.Query("sort")
	outQualityName := c.Query("outQualityName")
	currentPage, _ := com.StrTo(c.Query("currentPage")).Int()
	pageSize, _ := com.StrTo(c.Query("pageSize")).Int()
	valid := validation.Validation{}
	valid.Required(tenantId, "tenantId").Message("租户编号不能为空")
	if !valid.HasErrors() {
		p := model.QueryMultiOrderDefectOne(Db.ReDb().Sql, tenantId, groupName, sort, outQualityName, pageSize, currentPage)
		fmt.Println("-------")
		code := pkg.SUCCESS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  pkg.GetMsg(code),
			"data": p,
		})
	} else {
		code := pkg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  pkg.GetMsg(code),
			"data": "",
		})
	}

}
