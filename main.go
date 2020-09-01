package main

import (
	"log"
        "flag"
        "fmt"
        "os"

	"github.com/gin-gonic/gin"
        "github.com/pelletier/go-toml"
	"google.golang.org/grpc"

	"thermostat.org/bto"
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

        file, err := os.Open(*irPath)
        if err != nil {
                log.Fatalf("open error: %v, file: %v", err, *irPath)
                return
        }
        decoder := toml.NewDecoder(file)
        var conf bto.Config
        if err := decoder.Decode(&conf); err != nil {
                log.Fatalf("toml parser error: %v", err)
                return
        }

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
