package postgresql

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	MaxConnections string
	MaxConnectionIdleTime string
}