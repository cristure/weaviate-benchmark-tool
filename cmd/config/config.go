package config

var (
	GlobalConfig Config
)

type Config struct {
	Host   string `long:"host" short:"h" description:"Host to connect to" default:"localhost"`
	Scheme string `long:"scheme" short:"s" description:"Scheme to connect to" default:"http"`
}
