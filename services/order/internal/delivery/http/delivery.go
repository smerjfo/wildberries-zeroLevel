package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"l0/services/order/internal/useCase"
)

func init() {
	viper.SetDefault("HTTP_PORT", 8080)
	viper.SetDefault("HTTP_HOST", "127.0.0.1")
}

type Delivery struct {
	ucOrder useCase.Order
	router  *gin.Engine
}

func New(ucOrder useCase.Order) *Delivery {
	var d = &Delivery{ucOrder: ucOrder}
	gin.SetMode(gin.DebugMode)
	d.router = gin.New()
	d.router.LoadHTMLGlob("./services/order/internal/templates/*")
	d.router.Group("/orders").GET("/:id", d.ReadOrderByID)
	return d
}

func (d *Delivery) Run() error {
	logrus.Infoln("service started successfully on http port: %d", viper.GetUint("HTTP_PORT"))
	return d.router.Run(fmt.Sprintf("%s:%d", viper.GetString("HTTP_HOST"), viper.GetUint("HTTP_PORT")))
}
