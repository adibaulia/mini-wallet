package config

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type (
	Configuration struct {
		Host        string `mapstructure:"host"`
		Port        string `mapstructure:"port"`
		Username    string `mapstructure:"username"`
		Password    string `mapstructure:"password"`
		DBName      string `mapstructure:"dbName"`
		ServicePort string `mapstructure:"servicePort"`
	}
	Conn struct {
		PostgreCon *sql.DB
	}
)

var (
	connection *Conn
	Conf       *Configuration
)

func getEnvConfig() *Configuration {

	envFlag := flag.String("env", "dev", "state the active profile, default staging")
	flag.Parse()

	env := *envFlag

	viper.AddConfigPath("./config")
	viper.SetConfigName(fmt.Sprintf("config.%v", env))
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Cannot load configuration file: %v", err)
	}

	conf := &Configuration{}
	err = viper.Unmarshal(conf)
	if err != nil {
		log.Fatalf("Cannot unmarshall config: %v", err)
	}
	return conf
}

func GetConnection() *Conn {
	if connection != nil {
		return connection
	}
	Conf = getEnvConfig()
	connStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		Conf.Host, Conf.Port, Conf.Username, Conf.Password, Conf.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error opening postgres connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database!")

	connection = &Conn{
		PostgreCon: db,
	}
	return connection
}
