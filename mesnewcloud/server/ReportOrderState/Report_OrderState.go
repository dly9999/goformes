package Report_OrderState

import (
	"goformescloud/mesnewcloud/model"
	"goformescloud/mesnewcloud/pkg"

	//"encoding/json"
	"fmt"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"

	"github.com/gin-gonic/gin"
)

// @Tags report
// @Summary Query Order Brush Card List
// @Produce  json
// @Success 200 {object}  model.ProOrder
// @Failure 400 {object}  model.ProOrder
// @Router /report/orderstateinfo [Get]
// @Param tenantId query string true "tenantId"
// @Param orderCode query string  true "orderCode"
// @Param orderSort query string true  "orderSort"
// @Param currentPage query string  true "currentPage"
// @Param pageSize query string  true "pageSize"
func Method_OrderStateInfo(c *gin.Context) {

	tenantId := c.Query("tenantId")
	orderCode := c.Query("orderCode")
	OrderSort := c.Query("orderSort")
	currentPage, _ := com.StrTo(c.Query("currentPage")).Int()
	pageSize, _ := com.StrTo(c.Query("pageSize")).Int()
	if currentPage < 1 {
		currentPage = 1
	}
	if pageSize < 10 {
		pageSize = 10
	} else if pageSize > 100 {
		pageSize = 100
	}
	valid := validation.Validation{}
	valid.Required(tenantId, "tenantId").Message("租户编号不能为空")
	valid.Required(orderCode, "orderCode").Message("订单编号不能为空")

	code := pkg.SUCCESS
	data := model.Method_QueryOrderStateList(tenantId, orderCode, OrderSort, pageSize, currentPage)
	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  pkg.GetMsg(code),
		"data": data,
	})
}
