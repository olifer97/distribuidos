package main

import (
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/7574-sistemas-distribuidos/docker-compose-init/client/common"
)

type ConfigJSON struct {
    CLI_ID string `json:"CLI_ID"`
    CLI_SERVER_ADDRESS string `json:"CLI_SERVER_ADDRESS"`
    CLI_LOOP_LAPSE string `json:"CLI_LOOP_LAPSE"`
    CLI_LOOP_PERIOD string `json:"CLI_LOOP_PERIOD"`
}

// InitConfig Function that uses viper library to parse env variables or config file. 
// Returns ClientConfig struct.
// If some of the variables cannot be parsed, an error is returned
func InitConfig() (*common.ClientConfig, error) {
	v := viper.New()

	v.SetConfigName("client") 
	v.AddConfigPath("./config")  
	
	var configuration ConfigJSON

	v.AutomaticEnv()

	err_read := v.ReadInConfig()
	err_unmarshal := v.Unmarshal(&configuration)
	if err_read != nil && err_unmarshal != nil  {
		return nil, errors.Wrapf(err_read, "Could not parse config variables")
	}

	var id string
	if v.IsSet("CLI_ID") {
		id = v.GetString("CLI_ID")
	} else {
		id = configuration.CLI_ID
	}

	var server_address string
	if v.IsSet("CLI_SERVER_ADDRESS") {
		server_address = v.GetString("CLI_SERVER_ADDRESS")
	} else {
		server_address = configuration.CLI_SERVER_ADDRESS
	}

	var loop_lapse string
	if v.IsSet("CLI_LOOP_LAPSE") {
		loop_lapse = v.GetString("CLI_LOOP_LAPSE")
	} else {
		loop_lapse = configuration.CLI_LOOP_LAPSE
	}

	duration_loop_lapse, err := time.ParseDuration(loop_lapse)

	// Parse time.Duration variables and return an error
	// if those variables cannot be parsed
	if err != nil {
		return nil, errors.Wrapf(err, "Could not parse CLI_LOOP_LAPSE config var as time.Duration.")
	}

	var loop_period string
	if v.IsSet("CLI_LOOP_PERIOD") {
		loop_period = v.GetString("CLI_LOOP_PERIOD")
	} else {
		loop_period = configuration.CLI_LOOP_PERIOD
	}

	duration_loop_period, err := time.ParseDuration(loop_period)

	if err != nil {
		return nil, errors.Wrapf(err, "Could not parse CLI_LOOP_PERIOD config var as time.Duration.")
	}

	return &common.ClientConfig{
		ServerAddress: server_address,
		ID:            id,
		LoopLapse:     duration_loop_lapse,
		LoopPeriod:    duration_loop_period,
	}, nil
}

func main() {
	clientConfig, err := InitConfig()
	if err != nil {
		log.Fatalf("%s", err)
	}

	client := common.NewClient(*clientConfig)
	client.StartClientLoop()
}
