package global

import "os"

var (
	Token  = os.Getenv("TOKEN")
	Port   = os.Getenv("PORT")
)

const (
	portConf = "3000"
)

func GetPort() string {
	if Port == "" {
		Port = portConf
	}
	return Port
}
