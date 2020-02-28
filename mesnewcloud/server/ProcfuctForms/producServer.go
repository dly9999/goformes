package ProcfuctForms

import (
	"encoding/json"
	"fmt"
	"goformescloud/mesnewcloud/logprint"
	"goformescloud/mesnewcloud/model/Produc_Form"
	"goformescloud/mesnewcloud/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RtProForms struct {
	Code int
	Msg  string
	Data Produc_Form.Cutforms
}

type RtForms struct {
	Code int
	Msg  string
	Data Produc_Form.Forms
}

func Post_Cuforms(c *gin.Context) {
	conbuffer := make([]byte, 1024)
	var rtf RtProForms
	code := pkg.SUCCESS
	total, _ := c.Request.Body.Read(conbuffer)
	//	var dt Produc_Form.Cutforms
	var condtion Produc_Form.Cut_Condition
	if er := json.Unmarshal(conbuffer[:total], &condtion); er != nil {
		logprint.Logprint("cutjsonerr:", er)
		code = pkg.INVALID_PARAMS
	}
	fmt.Println("condtion", condtion)
	if condtion.TenantId != "" {
		cuts, er := Produc_Form.Dbcutforms(condtion)
		if er != nil {
			code = pkg.ERROR
		}
		rtf.Data = cuts
	} else {
		code = pkg.NOTENANTID
	}
	rtf.Code = code
	rtf.Msg = pkg.GetMsg(code)
	//	fmt.Println("---------------->", rtf)
	c.JSON(http.StatusOK, gin.H{
		"code": rtf.Code,
		"msg":  rtf.Msg,
		"data": rtf.Data,
	})
}
func Post_Sewforms(c *gin.Context) {
	conbuffer := make([]byte, 1024)
	var rtf RtForms
	code := pkg.SUCCESS
	total, _ := c.Request.Body.Read(conbuffer)
	//	var dt Produc_Form.Cutforms
	var condtion Produc_Form.Sew_Condition
	if er := json.Unmarshal(conbuffer[:total], &condtion); er != nil {
		logprint.Logprint("cutjsonerr:", er)
		code = pkg.INVALID_PARAMS
	}
	fmt.Println("condtion", condtion)
	if condtion.TenantId != "" {
		cuts, er := Produc_Form.Dbsewforms(condtion)
		if er != nil {
			code = pkg.ERROR
		}
		rtf.Data = cuts
	} else {
		code = pkg.NOTENANTID
	}
	rtf.Code = code
	rtf.Msg = pkg.GetMsg(code)
	//	fmt.Println("---------------->", rtf)
	c.JSON(http.StatusOK, gin.H{
		"code": rtf.Code,
		"msg":  rtf.Msg,
		"data": rtf.Data,
	})
}
