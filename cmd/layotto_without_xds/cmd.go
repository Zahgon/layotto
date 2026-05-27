package main

import (
	"runtime"

	v2 "mosn.io/mosn/pkg/config/v2"
	"mosn.io/mosn/pkg/metrics"
	"mosn.io/mosn/pkg/mosn"
	"mosn.io/mosn/pkg/stagemanager"

	// Actuator

	"github.com/urfave/cli"
)

var (
	flagToMosnLogLevel = map[string]string{
		"trace":    "TRACE",
		"debug":    "DEBUG",
		"info":     "INFO",
		"warning":  "WARN",
		"error":    "ERROR",
		"critical": "FATAL",
		"off":      "OFF",
	}

	cmdStart = cli.Command{
		Name:  "start",
		Usage: "start runtime. For example:  ./layotto start -c configs/config_standalone.json",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "config, c",
				Usage:  "Load configuration from `FILE`",
				EnvVar: "RUNTIME_CONFIG",
				Value:  "configs/config.json",
			}, cli.StringFlag{
				Name:   "feature-gates, f",
				Usage:  "config feature gates",
				EnvVar: "FEATURE_GATES",
			}, cli.StringFlag{
				Name:   "log-level, l",
				Usage:  "mosn log level, trace|debug|info|warning|error|critical|off",
				EnvVar: "LOG_LEVEL",
			}, cli.StringFlag{
				Name:   "logging-level, ll",
				Usage:  "layotto log level, trace|debug|info|warn|error|fatal",
				EnvVar: "LOGGING_LEVEL",
			}, cli.StringFlag{
				Name:   "logging-path, lp",
				Usage:  "layotto log file path, default ./",
				EnvVar: "LOGGING_PATH",
			},
		},
		Action: func(c *cli.Context) error {
			app := mosn.NewMosn()
			stm := stagemanager.InitStageManager(c, c.String("config"), app)

			// if needs featuregate init in parameter stage or init stage
			// append a new stage and called featuregate.ExecuteInitFunc(keys...)
			// parameter parsed registered
			stm.AppendParamsParsedStage(ExtensionsRegister)

			stm.AppendParamsParsedStage(DefaultParamsParsed)

			// init Stage
			stm.AppendInitStage(mosn.DefaultInitStage)
			stm.AppendInitStage(func(_ *v2.MOSNConfig) {
				// set version and go version
				metrics.SetVersion(GitVersion)
				metrics.SetGoVersion(runtime.Version())
			})
			// pre-startup
			stm.AppendPreStartStage(mosn.DefaultPreStartStage) // called finally stage by default
			// startup
			stm.AppendStartStage(mosn.DefaultStartStage)
			// after-startup
			stm.AppendAfterStartStage(SetActuatorAfterStart)
			// execute all stages
			stm.RunAll()
			return nil
		},
	}

	cmdStop = cli.Command{
		Name:  "stop",
		Usage: "stop mosn proxy",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "config, c",
				Usage:  "load configuration from `FILE`",
				EnvVar: "MOSN_CONFIG",
				Value:  "configs/mosn_config.json",
			},
		},
		Action: func(c *cli.Context) (err error) {
			app := mosn.NewMosn()
			stm := stagemanager.InitStageManager(c, c.String("config"), app)
			stm.AppendInitStage(mosn.InitDefaultPath)
			return stm.StopMosnProcess()
		},
	}

	cmdReload = cli.Command{
		Name:  "reload",
		Usage: "reconfiguration",
		Action: func(c *cli.Context) error {
			return nil
		},
	}
)

func SetActuatorAfterStart(_ stagemanager.Application) {
	_ = "STUB: not implemented"
	// register component actuator
	return
}

// set started

func DefaultParamsParsed(c *cli.Context) {
	_ = "STUB: not implemented"
	// log level control
	return
}

// log level control

// set feature gates

// ExtensionsRegister for register mosn rpc extensions
func ExtensionsRegister(_ *cli.Context) {
	_ = "STUB: not implemented"
	// 1. tracer driver register
	// Q: What is a tracer driver ?
	// A: MOSN implement a group of trace drivers, but only a configured driver will be loaded.
	//
	//	A tracer driver can create different tracer by different protocol.
	//	When MOSN receive a request stream, MOSN will try to start a tracer according to the request protocol.
	//	For more details,see https://mosn.io/blog/posts/skywalking-support/
	return
}

// 2. xprotocol action register
// RegisterXProtocolAction is MOSN's xprotocol framework's extensions.
// when a xprotocol implementation (defined by api.XProtocolCodec) registered, the registered action will be called.

// 3. register protocols that are used by layotto.
// RegisterXProtocolCodec add a new xprotocol implementation, which is a wrapper for protocol register

// 4. register tracer

// register buffer logger
