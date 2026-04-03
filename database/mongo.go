package database

import (
    "context"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
    clientOptions := options.Client().ApplyURI("mongodb+srv://kididrtina:ivkelicaloshmilica@hikari-agency.4r1keoy.mongodb.net/?appName=hikari-agency")
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    DB = client.Database("hikari-agency")
    log.Println("Connected to MongoDB")
}