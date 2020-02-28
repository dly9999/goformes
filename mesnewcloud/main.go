package main

/*"ginweb/gin-blog/routers"
"net/http"*/
import (
	"fmt"
	"goformescloud/mesnewcloud/config"
	_ "goformescloud/mesnewcloud/db"
	_ "goformescloud/mesnewcloud/docs"
	"goformescloud/mesnewcloud/router"
	"net/http"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	r := router.InitRouter()

	fmt.Println(config.Getmode())
	_, _, _, ht := config.Getconfig()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", ht.Http_Port),
		Handler:        r,
		ReadTimeout:    ht.Read_Timeout,
		WriteTimeout:   ht.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.ListenAndServe()
}
