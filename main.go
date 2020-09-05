package main

import (
	"log"
        "flag"
        "fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"thermostat.org/bto"
        "thermostat.org/defaults"
)

// /status
// /targetHeatingCoolingState/{INT_VALUE__0_TO_3}
// /targetTemperature/{INT_VALUE}
// /targetRelativeHumidity/{FLOAT_VALUE}
func routing(r *gin.Engine, ctrl *Controller) {
	r.GET("/status", ctrl.Status)
	r.GET("/targetHeatingCoolingState/:id", ctrl.TargetHeatingCoolingState)
	r.GET("/targetTemperature/:id", ctrl.TargetTemperature)
	r.GET("/targetRelativeHumidity/:id", ctrl.TargetRelativeHumidity)
}

func main() {
        var (
                host = flag.String("host", "localhost", "host address")
                port = flag.Int("port", 50051, "port number")
                irPath = flag.String("ir", "", "ir data path")
        )
        flag.Parse()

        var conf bto.Config
        irDataHander := defaults.New(*irPath)
        irDataHander.Load(&conf)

        address := fmt.Sprintf("%v:%d", *host, *port)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := bto.NewIRServiceClient(conn)
        irclient := bto.NewThermostatController(client, conf)
	ctrl := Controller{
		Client: irclient,
		State: Status{
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
