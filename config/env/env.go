package env

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	color "github.com/logrusorgru/aurora"
)

var (
	GinMode  string
	LogLevel string
)

func Load() {
	var now = time.Now().Format("02/01/2006 15:04:05")

	if err := godotenv.Load(); err != nil {
		fmt.Println(color.Yellow(fmt.Sprintf("%s -> Running application without .env file", now)))
	}

	GinMode = os.Getenv("GIN_MODE")
	LogLevel = os.Getenv("LOG_LEVEL")
}
