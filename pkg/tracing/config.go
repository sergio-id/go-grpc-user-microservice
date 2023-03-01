package tracing

// Config is jaeger config
type Config struct {
	ServiceName string `yaml:"serviceName" env:"JAEGER_SERVICE_NAME" env-default:"go-service"`
	HostPort    string `yaml:"hostPort" env:"JAEGER_HOST_PORT" env-default:"localhost:6831"`
	Enable      bool   `yaml:"enable" env:"JAEGER_ENABLE" env-default:"true"`
	LogSpans    bool   `yaml:"logSpans" env:"JAEGER_LOG_SPANS" env-default:"false"`
}
