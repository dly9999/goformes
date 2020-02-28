package Demotest

import (
	"goformescloud/mesnewcloud/pkg"

	//"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/****
**************
声明struct首字符必须大写否则 不会序列化

**************


*/

type Test struct {
	Name string
	Mps  string
}

type Testlist []Test

type RtDemo struct {
	code  int
	msg   string
	data  Testlist
	total int
}

// @Tags Demo
// @Summary Demo
// @Produce  json
// @Param name query string true "name"
// @Resource TestDemo
// @Success 200 {object} RtDemo
// @Failure 400 {object} RtDemo
// @Router /test/t [Get]
func TestDemo(c *gin.Context) {
	//	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	var tl Testlist
	var dt = Test{}
	dt.Name = c.Query("name")
	code := pkg.SUCCESS
	//dt.Mps = append(dt.Mps, make) //注意分配内存
	if dt.Name != "" {

		dt.Mps = "hello gin"

	} else {
		code = pkg.INVALID_PARAMS
	}
	tl = append(tl, dt)
	tl = append(tl, dt)
	fmt.Println(dt)
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   pkg.GetMsg(code),
		"data":  tl,
		"total": len(tl),
	})
}
