package probes

// Config is a probes config
type Config struct {
	ReadinessPath        string `yaml:"readinessPath" env:"PROBES_READINESS_PATH" env-default:"/ready"`
	LivenessPath         string `yaml:"livenessPath" env:"PROBES_LIVENESS_PATH" env-default:"/live"`
	Port                 string `yaml:"port" env:"PROBES_PORT" env-default:"8080"`
	PrometheusPath       string `yaml:"prometheusPath" env:"PROBES_PROMETHEUS_PATH" env-default:"/metrics"`
	PrometheusPort       string `yaml:"prometheusPort" env:"PROBES_PROMETHEUS_PORT" env-default:"8081"`
	CheckIntervalSeconds int    `yaml:"checkIntervalSeconds" env:"PROBES_CHECK_INTERVAL_SECONDS" env-default:"5"`
}
