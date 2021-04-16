package config

type Config struct {
	Server Server
}

type Server struct {
	Http struct {
		Port string
	}
}

type DB struct {
	DSN     string
	Dialect string
}

type Log struct {
	Mode string
	Type string
}
