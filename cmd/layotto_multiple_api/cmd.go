package main

import (
	"runtime"
	"time"

	"mosn.io/mosn/istio/istio1106"
	v2 "mosn.io/mosn/pkg/config/v2"
	"mosn.io/mosn/pkg/log"
	"mosn.io/mosn/pkg/metrics"
	"mosn.io/mosn/pkg/mosn"
	"mosn.io/mosn/pkg/server"
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
				Name:   "service-cluster, s",
				Usage:  "sidecar service cluster",
				EnvVar: "SERVICE_CLUSTER",
			}, cli.StringFlag{
				Name:   "service-node, n",
				Usage:  "sidecar service node",
				EnvVar: "SERVICE_NODE",
			}, cli.StringFlag{
				Name:   "service-type, p",
				Usage:  "sidecar service type",
				EnvVar: "SERVICE_TYPE",
			}, cli.StringSliceFlag{
				Name:   "service-meta, sm",
				Usage:  "sidecar service metadata",
				EnvVar: "SERVICE_META",
			}, cli.StringSliceFlag{
				Name:   "service-lables, sl",
				Usage:  "sidecar service metadata labels",
				EnvVar: "SERVICE_LAB",
			}, cli.StringSliceFlag{
				Name:   "cluster-domain, domain",
				Usage:  "sidecar service metadata labels",
				EnvVar: "CLUSTER_DOMAIN",
			}, cli.StringFlag{
				Name:   "feature-gates, f",
				Usage:  "config feature gates",
				EnvVar: "FEATURE_GATES",
			}, cli.StringFlag{
				Name:   "pod-namespace, pns",
				Usage:  "mosn pod namespaces",
				EnvVar: "POD_NAMESPACE",
			}, cli.StringFlag{
				Name:   "pod-name, pn",
				Usage:  "mosn pod name",
				EnvVar: "POD_NAME",
			}, cli.StringFlag{
				Name:   "pod-ip, pi",
				Usage:  "mosn pod ip",
				EnvVar: "POD_IP",
			}, cli.StringFlag{
				Name:   "log-level, l",
				Usage:  "mosn log level, trace|debug|info|warning|error|critical|off",
				EnvVar: "LOG_LEVEL",
			}, cli.StringFlag{
				Name:  "log-format, lf",
				Usage: "mosn log format, currently useless",
			}, cli.StringSliceFlag{
				Name:  "component-log-level, lc",
				Usage: "mosn component format, currently useless",
			}, cli.StringFlag{
				Name:  "local-address-ip-version",
				Usage: "ip version, v4 or v6, currently useless",
			}, cli.IntFlag{
				Name:  "restart-epoch",
				Usage: "epoch to restart, align to Istio startup params, currently useless",
			}, cli.IntFlag{
				Name:  "drain-time-s",
				Usage: "seconds to drain connections, default 600 seconds",
				Value: 600,
			}, cli.StringFlag{
				Name:  "parent-shutdown-time-s",
				Usage: "parent shutdown time seconds, align to Istio startup params, currently useless",
			}, cli.IntFlag{
				Name:  "max-obj-name-len",
				Usage: "object name limit, align to Istio startup params, currently useless",
			}, cli.IntFlag{
				Name:  "concurrency",
				Usage: "concurrency, align to Istio startup params, currently useless",
			}, cli.IntFlag{
				Name:  "log-format-prefix-with-location",
				Usage: "log-format-prefix-with-location, align to Istio startup params, currently useless",
			}, cli.IntFlag{
				Name:  "bootstrap-version",
				Usage: "API version to parse the bootstrap config as (e.g. 3). If unset, all known versions will be attempted",
			}, cli.StringFlag{
				Name:  "drain-strategy",
				Usage: "immediate",
			}, cli.BoolTFlag{
				Name:  "disable-hot-restart",
				Usage: "disable-hot-restart",
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
			stm.AppendInitStage(func(cfg *v2.MOSNConfig) {
				drainTime := c.Int("drain-time-s")
				server.SetDrainTime(time.Duration(drainTime) * time.Second)
				// istio parameters
				serviceCluster := c.String("service-cluster")
				serviceNode := c.String("service-node")
				serviceType := c.String("service-type")
				serviceMeta := c.StringSlice("service-meta")
				metaLabels := c.StringSlice("service-lables")
				clusterDomain := c.String("cluster-domain")
				podName := c.String("pod-name")
				podNamespace := c.String("pod-namespace")
				podIp := c.String("pod-ip")

				if serviceNode != "" {
					istio1106.InitXdsInfo(cfg, serviceCluster, serviceNode, serviceMeta, metaLabels)
				} else {
					if istio1106.IsApplicationNodeType(serviceType) {
						sn := podName + "." + podNamespace
						serviceNode = serviceType + "~" + podIp + "~" + sn + "~" + clusterDomain
						istio1106.InitXdsInfo(cfg, serviceCluster, serviceNode, serviceMeta, metaLabels)
					} else {
						log.StartLogger.Infof("[mosn] [start] xds service type is not router/sidecar, use config only")
						istio1106.InitXdsInfo(cfg, "", "", nil, nil)
					}
				}
			})
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
