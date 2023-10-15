package middleware

import (
  "log"
  "net/http"
  "os"
  "fmt"
  "encoding/json"
  "context"
  "github.com/joho/godotenv"
  "github.com/gorilla/mux"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

var collection *mongo.Collection

func init () {
  loadTheEnv()
  createDBInstance()
}

func loadTheEnv() {
  err := godotenv.Load(".env")
  if err 0
}

func GetAllTasks() {

}

func CreateTask() {

}

func TaskComplete() {

}

func UndoTask() {

}

func DeleteTask() {

}

func DeleteAllTasks() {

}


