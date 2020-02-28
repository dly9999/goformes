package ProOrdtemp

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

// @Tags ProOrdtemp
// @Summary ProOrdtemp
// @Produce  json
// @Resource ProOrdtemp
// @Success 200 {object}  model.ProOrder
// @Failure 400 {object}  model.ProOrder
// @Router /ProOrdtemp/t [Get]
func GetOrdertemp(c *gin.Context) {
	tenantId := c.Query("tenantId")
	Isbulk := c.Query("Isbulk")
	orderCode := c.Query("orderCode")

	customer := c.Query("customer")
	sort := c.Query("sort")
	deliveryDateFrom := c.Query("deliveryDateFrom")
	deliveryDateTo := c.Query("deliveryDateTo")
	currentPage, _ := com.StrTo(c.Query("currentPage")).Int()
	pageSize, _ := com.StrTo(c.Query("pageSize")).Int()
	valid := validation.Validation{}
	valid.Required(tenantId, "tenantId").Message("租户编号不能为空")
	if !valid.HasErrors() {
		p := model.QueryMultiProOrder(Db.ReDb().Sql, tenantId, Isbulk, orderCode, customer, sort, deliveryDateFrom, deliveryDateTo, pageSize, currentPage)
		fmt.Println("-----------")
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
