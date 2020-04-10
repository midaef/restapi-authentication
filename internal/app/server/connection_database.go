package server

import (
	"github.com/jackc/pgx"
	"go.uber.org/zap"
)

func NewConnection() (conn *pgx.Conn) {
	config := NewConfig("resources/configs/config.json")
	var runtimeParams map[string]string
	runtimeParams = make(map[string]string)
	runtimeParams["application_name"] = "todo"
	connConfig := pgx.ConnConfig{
		User:              config.User,
		Password:          config.Password,
		Host:              config.Host,
		Port:              config.DBport,
		Database:          config.DBname,
		TLSConfig:         nil,
		UseFallbackTLS:    false,
		FallbackTLSConfig: nil,
		RuntimeParams:     runtimeParams,
	}
	conn, err := pgx.Connect(connConfig)
	loggerConnection, _ := zap.NewProduction()
	if err != nil {
		loggerConnection.Error("Connection error", zap.Error(err))
	} else if config.LogLevel == "debug" {
		loggerConnection.Debug("Connection DB",
			zap.String("user", config.User),
			zap.String("host", config.Host),
			zap.String("port", string(config.DBport)),
			zap.String("dbname", config.DBname))
	}
	return conn
}
