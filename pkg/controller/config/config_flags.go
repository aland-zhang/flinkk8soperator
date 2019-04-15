// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"fmt"

	"github.com/spf13/pflag"
)

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "resyncPeriod"), "10s", "Determines the resync period for all watchers.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "limitNamespace"), "", "Namespaces to watch for this propeller")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "metricsPrefix"), "flinkk8soperator", "Prefix for metrics propagated to prometheus")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "prof-port"), "10254", "Profiler port")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "ingressUrlFormat"), *new(string), "")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "useKubectlProxy"), *new(bool), "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "ProxyPort"), "8001", "The port at which flink cluster runs locally")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "containerNameFormat"), *new(string), "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "statemachineStalenessDuration"), "10m", "Duration for statemachine staleness.")
	return cmdFlags
}