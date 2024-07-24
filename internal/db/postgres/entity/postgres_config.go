package entity

// PostgresConfig описыает конфиг посгри.
type PostgresConfig struct {
	Host     string
	Port     int
	DBName   string
	UserName string
	Password string
}
