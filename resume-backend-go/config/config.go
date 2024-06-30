package config

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBClient *mongo.Client

type Config struct {
    MongoDB struct {
        URI string `mapstructure:"uri"`
    } `mapstructure:"mongodb"`
}

var AppConfig Config

func LoadConfig(env string) {
    viper.SetConfigName("config." + env)
    viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
    viper.AddConfigPath(".")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }

    if err := viper.Unmarshal(&AppConfig); err != nil {
        log.Fatalf("Unable to decode into struct, %v", err)
    }
}

func ConnectMongoDB() {
    uri := AppConfig.MongoDB.URI
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")
    MongoDBClient = client
}
