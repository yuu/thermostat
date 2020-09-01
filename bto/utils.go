package bto

import (
        "context"
        "log"
        "time"
)

type ThermostatController interface {
        On() (*WriteResponse, error)
        Off() (*WriteResponse, error)
        // Mode() (*WriteResponse, error)
        Up() (*WriteResponse, error)
        Down() (*WriteResponse, error)
}

type Config struct {
        Aircon struct {
                On  []uint32 `toml:"on"`
                Off []uint32 `toml:"off"`
                Up  []uint32 `toml:"up"`
                Down []uint32 `toml:"down"`
        } `toml:"aircon"`
}

type iRThermostat struct {
        client IRServiceClient
        conf Config
}

func NewThermostatController(client IRServiceClient, conf Config) ThermostatController {
	return &iRThermostat{client, conf}
}

func (ir *iRThermostat) On() (*WriteResponse, error) {
        return ir.write(ir.conf.Aircon.On)
}

func (ir *iRThermostat) Off() (*WriteResponse, error) {
        return ir.write(ir.conf.Aircon.Off)
}

func (ir *iRThermostat) Up() (*WriteResponse, error) {
        return ir.write(ir.conf.Aircon.Up)
}

func (ir *iRThermostat) Down() (*WriteResponse, error) {
        return ir.write(ir.conf.Aircon.Down)
}

func (ir * iRThermostat) write(data []uint32) (*WriteResponse, error) {
        ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
        defer cancel()

        req := WriteRequest{
                Frequency: 38000,
                Data:      data,
        }
        res, err := ir.client.Write(ctx, &req)
        if err != nil {
                log.Fatalf("could not write: %v", err)
                return nil, err
        }

        return res, nil
}
