package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/model"
	"net/http"
	"go-gin-demo/service"
	"fmt"
	"go-gin-demo/util"
	log "github.com/jeanphorn/log4go"
	"go-gin-demo/middleware"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	personService := &service.PersonServiceImpl{}
	//ResponseUtil := util.ResponseUtil{}
	JsonUtil := util.JsonUtil{}
	LogMiddleware := middleware.LogMiddleware{}
	log.LoadConfiguration("./logconf.json")



	router.Use(LogMiddleware.Logger())
	router.GET("/ping/:name/*action", func(context *gin.Context) {

		context.JSON(200, gin.H{
			"message": "嘿嘿嘿，" + context.Param("name") + context.Param("action"),
		})
	})

	router.POST("/getPersonByName", func(context *gin.Context) {

		var personVO model.PersonVO
		if err := context.BindJSON(&personVO); err != nil {
			context.JSON(http.StatusBadRequest, ResponseUtil.CreateFailResponse(nil))
		}
		defer func() {
			if err := recover(); err != nil {
				log.Error("error", err)
			}
		}()

		log.Info("input param:%s", JsonUtil.ToJson(personVO))
		fmt.Println(personVO.Name)
		personBO := personService.GetPersonInfo(personVO.Name)

		log.Info("output data :%s", JsonUtil.ToJson(personBO))
		//panic("it is a error")
		if &personBO != nil {
			context.JSON(200, ResponseUtil.CreateSuccessResponse(personBO))
			return
		}

		context.JSON(500, ResponseUtil.CreateFailResponse(nil))
	})
	router.Run(":8088")
}
