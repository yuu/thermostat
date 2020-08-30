package bto

import (
        "context"
        "log"
        "time"
)

type ThermostatController interface {
        // On()
        // Off()
        // Mode()
        Up()
        // Down()
}

type iRThermostat struct {
        client IRServiceClient
}

func NewThermostatController(client IRServiceClient) ThermostatController {
	return &iRThermostat{client}
}

func (ir *iRThermostat) Up() {
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()

        req := WriteRequest{
                Frequency: 38000,
                Data:      []uint32{0x00, 0x00},
        }
        _, err := ir.client.Write(ctx, &req)
        if err != nil {
                log.Fatalf("could not write: %v", err)
        }
}
