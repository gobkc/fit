package conf

type Conf struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Debug       bool   `json:"debug"`
	RestAddr    string `json:"rest_addr"`
	Dsn         string `json:"dsn"`
	Email       Email  `json:"email"`
	Cors        Cors   `json:"cors"`
	MaxIdle     int    `json:"max_idle"`
	MaxConn     int    `json:"max_conn"`
	MaxLeftTime int    `json:"max_left_time"`
	JwtSalt     string `json:"jwt_salt"`
}

type Cors struct {
	Enabled          bool     `json:"enabled"`
	MaxAge           int      `json:"max_age"`
	AllowedOrigins   []string `json:"allowed_origins"`
	AllowedMethods   []string `json:"allowed_methods"`
	AllowedHeaders   []string `json:"allowed_headers"`
	AllowCredentials bool     `json:"allow_credentials"`
}

type Email struct {
	Imap string `json:"imap"`
	Smtp string `json:"smtp"`
	User string `json:"user"`
	Pass string `json:"pass"`
}
