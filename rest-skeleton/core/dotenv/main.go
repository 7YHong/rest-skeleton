package dotenv

import (
	"github.com/mix-go/dotenv"
)

func init() {
	// Env
	if err := dotenv.Load(".env"); err != nil {
		panic(err)
	}
}
