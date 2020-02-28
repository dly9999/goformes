package EmployeeOutput

import (
	"encoding/json"
	"fmt"
	"goformescloud/mesnewcloud/db"
	"goformescloud/mesnewcloud/logprint"
	"goformescloud/mesnewcloud/model"
	"goformescloud/mesnewcloud/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostEmployeeOutPut(c *gin.Context) {
	conbuffer := make([]byte, 1024)
	/*var rtf RtProForms*/
	code := pkg.SUCCESS
	total, _ := c.Request.Body.Read(conbuffer)
	//	var dt Produc_Form.Cutforms
	var employee model.EmployeeOutPutParameter

	if er := json.Unmarshal(conbuffer[:total], &employee); er != nil {
		logprint.Logprint("cutjsonerr:", er)
		code = pkg.INVALID_PARAMS
	}
	fmt.Println("employee", employee)
	if employee.TenantId != "" {
		/*cuts, er := model.QueryMultiEmployeeOutPut(Db.ReDb().Sql, employee)
		if er != nil {
			code = pkg.ERROR
		}
		rtf.Data = cuts*/
		p := model.QueryMultiEmployeeOutPut(Db.ReDb().Sql, employee)
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  pkg.GetMsg(code),
			"data": p,
		})

	} else {
		code = pkg.NOTENANTID
	}

}
