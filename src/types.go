package main

type Configuration struct {
	Environment string                `json:"environment"`
	Database    DatabaseConfiguration `json:"db"`
	Ssl         SslConfiguration      `json:"ssl"`
}

type DatabaseConfiguration struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type WebServiceConfiguration struct {
	Ip   string           `json:"ip"`
	Port int              `json:"port"`
	Ssl  SslConfiguration `json:"ssl_configuration"`
}

type SslConfiguration struct {
	Force           bool   `json:"foce_ssl"`
	CertificatePath string `json:"certificate_path"`
	CertificateKey  string `json:"certificate_key"`
}
