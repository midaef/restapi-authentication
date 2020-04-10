package server

import (
	"io/ioutil"
	"log"

	"github.com/buger/jsonparser"
)

type Configuration struct {
	ServerHost string
	Port       string
	LogLevel   string
	Host       string
	DBname     string
	DBport     uint16
	User       string
	Password   string
}

func NewConfig(path string) *Configuration {
	file, error := ioutil.ReadFile(path)
	if error != nil {
		log.Fatal("Error open file: ", error)
	}
	serverHost, _ := jsonparser.GetString(file, "server", "server_host")
	port, _ := jsonparser.GetString(file, "server", "server_port")
	logLevel, _ := jsonparser.GetString(file, "server", "log_level")
	host, _ := jsonparser.GetString(file, "database", "host")
	dbname, _ := jsonparser.GetString(file, "database", "dbname")
	dbport, _ := jsonparser.GetInt(file, "database", "db_port")
	user, _ := jsonparser.GetString(file, "database", "user")
	password, _ := jsonparser.GetString(file, "database", "password")
	return &Configuration{
		ServerHost: serverHost,
		Port:       port,
		LogLevel:   logLevel,
		Host:       host,
		DBname:     dbname,
		DBport:     uint16(dbport),
		User:       user,
		Password:   password,
	}
}
