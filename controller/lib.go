package controller

import (
        "log"
        "net/http"

	"github.com/gin-gonic/gin"

	"github.com/yuu/thermostat/bto"
)

const (
        MODE_OFF  = 0
        MODE_HEAT = 1
        MODE_COOL = 2
        MODE_AUTO = 3
)

type Status struct {
        CurrentHeatingCoolingState int `json:"CurrentHeatingCoolingState"`
        CurrentTemperature         float64 `json:"CurrentTemperature"`
        CurrentRelativeHumidity    float64 `json:"CurrentRelativeHumidity"`

        TargetHeatingCoolingState  int `json:"TargetHeatingCoolingState"`
        TargetTemperature          float64 `json:"TargetTemperature"`
        TargetRelativeHumidity     float64 `json:"TargetRelativeHumidity"`
}

type Controller struct {
	Client bto.ThermostatController
        State Status
}

// func(c *gin.Context) {


//         c.JSON(200, gin.H{"message": "pong", "code": res.GetCode()})
// }
func (c *Controller) Status(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, c.State)
}

func (c *Controller) TargetHeatingCoolingState(ctx *gin.Context) {
	newMode := ctx.Param("id")
        log.Printf("newMode: %v", newMode)
}

func (c *Controller) TargetTemperature(ctx *gin.Context) {
	newTemp := ctx.Param("id")
        log.Printf("newTemp: %v", newTemp)
}

func (c *Controller) TargetRelativeHumidity(ctx *gin.Context) {
	newHumi := ctx.Param("id")
        log.Printf("newHumi: %v", newHumi)
}
