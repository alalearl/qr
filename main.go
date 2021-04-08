package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Generate struct {
	Value string `form:"value" json:"value" xml:"value"  binding:"required"`
}

func main() {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())


  v1 := router.Group("/v1")
  {
   v1.GET("/generate/:value", QRGetGeneratorHandler)
	 v1.POST("/generate", QRPostGeneratorHandler)

  }
  router.Run()
}

func QRGetGeneratorHandler(context *gin.Context) {
	value := context.Params.ByName("value")
	qrBase64 := generate(value)
	context.JSON(200, gin.H{"value": qrBase64})
}

func QRPostGeneratorHandler(context *gin.Context) {
	var data Generate
	if err := context.ShouldBindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	value := data.Value
	qrBase64 := generate(value)
	context.JSON(200, gin.H{"value": qrBase64})
}