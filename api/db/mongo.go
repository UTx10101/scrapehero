package db

import (
	// builtin
	"log"

	// vendored
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DBClient *mongo.Client

func InitMongo() error {
	var err error = nil
	
	var mongoHost = viper.GetString("mongo.host")
	var mongoPort = viper.GetString("mongo.port")
	var mongoUsername = viper.GetString("mongo.username")
	var mongoPassword = viper.GetString("mongo.password")
	
	clientOptions := options.Client().ApplyURI("mongodb://"+mongoUsername+":"+"mongoPassword"+"@"+mongoHost+":"+mongoPort)
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        return err
    }
    DBClient = client
	return err
}
