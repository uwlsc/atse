package lib

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`

	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`

	MaxMultipartMemory int64  `mapstructure:"MAX_MULTIPART_MEMORY"`
	StorageBucketName  string `mapstructure:"STORAGE_BUCKET_NAME"`
	MailClientID       string `mapstructure:"MAIL_CLIENT_ID"`
	MailClientSecret   string `mapstructure:"MAIL_CLIENT_SECRET"`
	MailTokenType      string `mapstructure:"MAIL_TOKEN_TYPE"`

	AWSRegion          string `mapstructure:"AWS_REGION"`
	AWSAccessKey       string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AWSLocationIndex   string `mapstructure:"AWS_LOCATION_INDEX"`

	PoolID       string `mapstructure:"COGNITO_POOL_ID"`
	ClientID     string `mapstructure:"COGNITO_CLIENT_ID"`
	S3BucketName string `mapstructure:"AWS_S3_BUCKET_NAME"`
	AWSPINPOINT  string `mapstructure:"AWS_PINPOINT"`

	AdminEmail    string `mapstructure:"ADMIN_EMAIL"`
	AdminPassword string `mapstructure:"ADMIN_PASSWORD"`
}

var globalEnv = Env{
	MaxMultipartMemory: 10 << 20, // 10 MB
}

func GetEnv() Env {
	return globalEnv
}

func NewEnv(configFile string) func() Env {
	return func() Env {
		if configFile == "" {
			viper.SetConfigFile(".env")
		} else {
			viper.SetConfigFile(configFile)
		}

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("cannot read configuration", err)
		}

		err = viper.Unmarshal(&globalEnv)
		if err != nil {
			log.Fatal("environment cant be loaded: ", err)
		}

		log.Printf("%#v \n", &globalEnv)
		return globalEnv
	}
}
