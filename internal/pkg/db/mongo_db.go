package db


import (
    "context"
    "fmt"
    "log"

	"github.com/samuael/Project/Weg/internal/pkg/entity"
    // "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongodb function to connect mongodb 
func  ConnectMongodb() *mongo.Database {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client  , era :=  mongo.Connect(context.TODO()  , clientOption )
	if era != nil {
		log.Println(era)
		return nil
	}
	era = client.Ping(context.TODO()  , nil )
	if era != nil {
		log.Println("Error WHILE PINGING "  , era )
		return nil 
	}
	fmt.Println("DB Connected ...\nDB : Mongo ")
	return client.Database(entity.DBName)
}