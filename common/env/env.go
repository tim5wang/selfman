package env

import "os"

const (
	Env_dev  = "dev"
	Env_live = "live"
)

func Env() string {
	e := os.Getenv("env")
	if e == "" {
		return Env_dev
	}
	return e
}
