package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/yuu/thermostat/bto"
	cl "github.com/yuu/thermostat/controller"
)

const (
	address = "localhost:50051"
)

// /status
// /targetHeatingCoolingState/{INT_VALUE__0_TO_3}
// /targetTemperature/{INT_VALUE}
// /targetRelativeHumidity/{FLOAT_VALUE}
func routing(r *gin.Engine, ctrl *cl.Controller) {
	r.GET("/status", ctrl.Status)
	r.GET("/targetHeatingCoolingState/:id", ctrl.TargetHeatingCoolingState)
	r.GET("/targetTemperature/:id", ctrl.TargetTemperature)
	r.GET("/targetRelativeHumidity/:id", ctrl.TargetRelativeHumidity)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := bto.NewIRServiceClient(conn)
        irclient := bto.NewThermostatController(client)
	ctrl := cl.Controller{
		Client: irclient,
		State: cl.Status{
			CurrentHeatingCoolingState: 0,
			CurrentTemperature:         0,
			CurrentRelativeHumidity:    0,
			TargetHeatingCoolingState:  0,
			TargetTemperature:          0,
			TargetRelativeHumidity:     0,
		},
	}
	r := gin.Default()
	routing(r, &ctrl)

	r.Run(":3000")
}
