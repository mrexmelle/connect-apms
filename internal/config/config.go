package config

import (
	"context"
	"fmt"
	"os"

	"github.com/go-chi/jwtauth"
	"github.com/mrexmelle/connect-apms/internal/idp"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Db             *mongo.Database
	IdpClient      *idp.Client
	TokenAuth      *jwtauth.JWTAuth
	JwtValidMinute int
	Port           int
}

func New(
	plainConfigName string,
	configType string,
	configPaths []string,
) (Config, error) {
	viper.SetConfigName(ComposeConfigName(plainConfigName))
	viper.SetConfigType(configType)
	for _, cp := range configPaths {
		viper.AddConfigPath(cp)
	}
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	db, err := CreateDb()
	if err != nil {
		return Config{}, err
	}

	idp := CreateIdpClient()

	jwta := CreateTokenAuth("app.security.jwt.secret")
	jwtvm := viper.GetInt("app.security.jwt.valid-minute")
	port := viper.GetInt("app.server.port")

	return Config{
		Db:             db,
		IdpClient:      idp,
		TokenAuth:      jwta,
		JwtValidMinute: jwtvm,
		Port:           port,
	}, nil
}

func ComposeConfigName(plainConfigName string) string {
	profile := os.Getenv("APP_PROFILE")
	if profile == "" {
		profile = "local"
	}
	return fmt.Sprintf("%s-%s", plainConfigName, profile)
}

func CreateDb() (*mongo.Database, error) {
	var dbName = viper.GetString("app.datasource.dbName")
	var dns = fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s",
		viper.GetString("app.datasource.user"),
		viper.GetString("app.datasource.password"),
		viper.GetString("app.datasource.host"),
		viper.GetString("app.datasource.port"),
		dbName,
	)

	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(dns),
	)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	return client.Database(dbName), nil
}

func CreateTokenAuth(configKey string) *jwtauth.JWTAuth {
	return jwtauth.New(
		"HS256",
		[]byte(viper.GetString(configKey)),
		nil,
	)
}

func CreateIdpClient() *idp.Client {
	return idp.NewClient(
		viper.GetString("app.idp.protocol"),
		viper.GetString("app.idp.host"),
		viper.GetInt("app.idp.port"),
	)
}
