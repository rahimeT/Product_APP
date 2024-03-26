package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

func GetConnectionPool( context context.Context, config *Config) *pgxpool.Pool{
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s pool_max_conns=%s pool_max_conn_idle_time=%s", 
	config.Host, 
	config.Port, 
	config.User, 
	config.Password,
	config.DBName, 
	config.MaxConnections, 
	config.MaxConnectionIdleTime)
	
	connConfig, parseConfigErr := pgxpool.ParseConfig(connString)
	if parseConfigErr != nil {
		panic(parseConfigErr)
	}

	conn, err := pgxpool.ConnectConfig(context, connConfig)
	if err != nil {
		log.Error("Error connecting to the database: ", err)
		panic(err)
	}

	return conn
}