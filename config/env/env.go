package env

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	color "github.com/logrusorgru/aurora"
)

var (
	GinMode           string
	LogLevel          string
	AwsAccessKeyId    string
	AwsSecretKeyId    string
	AwsUserPoolId     string
	AwsUserPoolRegion string
	AwsClientId       string
)

func Load() {
	var now = time.Now().Format("02/01/2006 15:04:05")

	if err := godotenv.Load(); err != nil {
		fmt.Println(color.Yellow(fmt.Sprintf("%s -> Running application without .env file", now)))
	}

	setEnvVars()
}

func setEnvVars() {
	if GinMode = os.Getenv("GIN_MODE"); GinMode == "" {
		fmt.Println(color.Yellow("Application Starts Without value for Env Var GIN_MODE"))
	}

	if LogLevel = os.Getenv("LOG_LEVEL"); GinMode == "" {
		fmt.Println(color.Yellow("Application Starts Without value for Env Var LOG_LEVEL"))
	}

	if AwsAccessKeyId = os.Getenv("AWS_ACCESS_KEY_ID"); AwsAccessKeyId == "" {
		fmt.Println(color.Yellow("Application Starts Without value for Env Var AWS_ACCESS_KEY_ID"))
	}

	if AwsSecretKeyId = os.Getenv("AWS_SECRET_KEY_ID"); AwsSecretKeyId == "" {
		fmt.Println(color.Yellow("Application Starts Without value for Env Var AWS_SECRET_KEY_ID"))
	}

	if AwsUserPoolId = os.Getenv("AWS_USER_POOL_ID"); AwsUserPoolId == "" {
		fmt.Println(color.Yellow("Application Starts Without value for Env Var AWS_USER_POOL_ID"))
	}

	if AwsUserPoolRegion = os.Getenv("AWS_USER_POOL_REGION"); AwsUserPoolRegion == "" {
		fmt.Println(color.Yellow("Application Starts Without value for Env Var AWS_USER_POOL_REGION"))
	}

	if AwsClientId = os.Getenv("AWS_CLIENT_ID"); AwsClientId == "" {
		fmt.Println(color.Yellow("Application Starts Without value for Env Var AWS_CLIENT_ID"))
	}
}
