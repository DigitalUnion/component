package dukafka

type Config struct {
	ApiVersionRequest bool   `json:"api_version_request" yaml:"api_version_request"`
	Acks              string `json:"acks" yaml:"acks"`
	TimeoutMs         int    `json:"timeout_ms" yaml:"timeout_ms"`
	Hosts             string `json:"hosts" yaml:"hosts"`
	GroupId           string `json:"group_id" yaml:"group_id"`
	OffsetReset       string `json:"offset_reset" yaml:"offset_reset"`
	Topic             string `json:"topic" yaml:"topic"`
	SecurityProtocol  string `json:"security_protocol" yaml:"security_protocol"`
	SslCaLocation     string `json:"ssl_ca_location" yaml:"ssl_ca_location"`
	SaslUsername      string `json:"sasl_username" yaml:"sasl_username"`
	SaslPassword      string `json:"sasl_password" yaml:"sasl_password"`
	SaslMechanism     string `json:"sasl_mechanism" yaml:"sasl_mechanism"`
}
