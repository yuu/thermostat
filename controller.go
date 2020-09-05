package main

import (
        "os"
        "log"
        "net/http"

	"github.com/gin-gonic/gin"

	"thermostat.org/bto"
        "thermostat.org/defaults"
)

const (
        MODE_OFF  = 0
        MODE_HEAT = 1
        MODE_COOL = 2
        MODE_AUTO = 3
)

type Status struct {
        handler defaults.Defaults

        CurrentHeatingCoolingState int `json:"CurrentHeatingCoolingState"`
        CurrentTemperature         float64 `json:"CurrentTemperature"`
        CurrentRelativeHumidity    float64 `json:"CurrentRelativeHumidity"`

        TargetHeatingCoolingState  int `json:"TargetHeatingCoolingState"`
        TargetTemperature          float64 `json:"TargetTemperature"`
        TargetRelativeHumidity     float64 `json:"TargetRelativeHumidity"`
}

func (s *Status) save() {
        s.handler.Save(s)
}

func (s *Status) load() {
        s.handler.Load(s)
}

type Controller struct {
	irClient bto.ThermostatController
        state Status
}

type intParam struct {
        ID int `uri:"id"`
}
type floatParam struct {
        ID float64 `uri:"id"`
}

func NewController(irCtrl bto.ThermostatController) *Controller {
        home := os.Getenv("HOME")
        var state Status
        handler := defaults.New(home + "/.config/io.flutia.thermostat.toml")
        state.handler = handler
        state.load()

        return &Controller{irCtrl, state}
}

func (c *Controller) Status(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, c.state)
}

func (c *Controller) TargetHeatingCoolingState(ctx *gin.Context) {
        var param intParam
        if err := ctx.ShouldBindUri(&param); err != nil {
                ctx.JSON(http.StatusBadRequest, gin.H{"msg": err})
                return
        }

        c.state.CurrentHeatingCoolingState = param.ID
        c.state.TargetHeatingCoolingState = param.ID
        c.state.save()

        log.Printf("newMode: %v", param.ID)
}

func (c *Controller) TargetTemperature(ctx *gin.Context) {
        var param floatParam
        if err := ctx.ShouldBindUri(&param); err != nil {
                ctx.JSON(http.StatusBadRequest, gin.H{"msg": err})
                return
        }

        c.state.CurrentTemperature = param.ID
        c.state.TargetTemperature = param.ID
        c.state.save()

        log.Printf("newTemp: %v", param.ID)
}

func (c *Controller) TargetRelativeHumidity(ctx *gin.Context) {
        var param floatParam
        if err := ctx.ShouldBindUri(&param); err != nil {
                ctx.JSON(http.StatusBadRequest, gin.H{"msg": err})
                return
        }

        c.state.CurrentRelativeHumidity = param.ID
        c.state.TargetRelativeHumidity = param.ID
        c.state.save()

        log.Printf("newHumi: %v", param.ID)
}
