module github.com/yuu/thermostat

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/pelletier/go-toml v1.8.0 // indirect
	google.golang.org/grpc v1.31.1
	thermostat.org/bto v0.0.0-00010101000000-000000000000
	thermostat.org/defaults v0.0.0-00010101000000-000000000000 // indirect
)

replace thermostat.org/bto => ./bto

replace thermostat.org/defaults => ./defaults
