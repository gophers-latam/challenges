package global

import (
	"log"
	"os"
	"strings"

	chg "github.com/gophers-latam/challenges/http"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const Prefix = ".go"

// get from real environment variables
var (
	Token  = os.Getenv("TOKEN")
	Port   = os.Getenv("PORT")
	DbHost = os.Getenv("DBHOST")
	DbName = os.Getenv("DBNAME")
	DbUser = os.Getenv("DBUSER")
	DbPass = os.Getenv("DBPASS")
)

// override if empty environment variables
type Config struct {
	Token  string
	Port   string
	DbHost string
	DbName string
	DbUser string
	DbPass string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file")
	}

	logger, err := zap.NewProduction()
	_ = logger.Sync()
	if err != nil {
		panic(err.Error())
	}

	zap.ReplaceGlobals(logger)

	for country, data := range chg.TimeZones {
		chg.FlagToCountry[strings.ToLower(data.Flag)] = country
	}
}

// handle envars setting
func GetConfig() Config {
	if Token == "" {
		Token = os.Getenv("DEV_TOKEN")
	}
	if Port == "" {
		Port = os.Getenv("DEV_PORT")
	}
	if DbHost == "" {
		DbHost = os.Getenv("DEV_DBHOST")
	}
	if DbName == "" {
		DbName = os.Getenv("DEV_DBNAME")
	}
	if DbUser == "" {
		DbUser = os.Getenv("DEV_DBUSER")
	}
	if DbPass == "" {
		DbPass = os.Getenv("DEV_DBPASS")
	}

	return Config{
		Token:  Token,
		Port:   Port,
		DbHost: DbHost,
		DbName: DbName,
		DbUser: DbUser,
		DbPass: DbPass,
	}
}
