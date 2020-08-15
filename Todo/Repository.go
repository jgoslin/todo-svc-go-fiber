package Todo

import (
	"context"
	"errors"
	"fmt"
	"go-fiber-todos/Database"
	"go-fiber-todos/Errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

var mgi *Database.MongoInstance
var todoCollection *mongo.Collection

func init() {
	mgi = new(Database.MongoInstance)
	err := mgi.Connect()
	if err != nil {
		// Panic when cannot connect to database
		panic(err)
	}
	todoCollection = mgi.Client.Database("Sample").Collection("todos")
}

func getTodosFromDB() []*Todo {
	var todoResults []*Todo
	cursor, err := todoCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for cursor.Next(mgi.CTX) {
		var todo Todo
		err = cursor.Decode(&todo)
		if err == nil {
			todoResults = append(todoResults, &todo)
		} else {
			fmt.Println(err)
		}
	}
	return todoResults
}

func getTodoFromDB(id string) (*Todo, error) {
	var todo Todo
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objectId}}
	err := todoCollection.FindOne(context.TODO(), filter).Decode(&todo)
	if err == nil {
		return &todo, nil
	} else {
		log.Println(err)
		return nil, errors.New(Errors.NOTFOUNDERROR)
	}
}

func insertTodoIntoDB(todo Todo) error {
	todo.Id = primitive.NewObjectID()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	if todo.Name != "" {
		insertResults, err := todoCollection.InsertOne(context.TODO(), todo)
		if err == nil {
			log.Println(insertResults.InsertedID)
			return nil
		} else {
			return errors.New(Errors.DBPROCESSINGERROR)
		}
	}
	return errors.New(Errors.BADREQUESTERROR)
}

func deleteTodoFromDB(id string) error {
	objectId, objErr := primitive.ObjectIDFromHex(id)
	if objErr == nil {
		filter := bson.D{{"_id", objectId}}
		delResult, err := todoCollection.DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Println(err)
			return errors.New(Errors.DBPROCESSINGERROR)
		} else {
			if delResult.DeletedCount == 0 {
				return errors.New(Errors.NOTFOUNDERROR)
			} else {
				return nil
			}
		}
	} else {
		log.Println(objErr)
		return errors.New(Errors.BADREQUESTERROR)
	}
}

func updateTodoInDB(id string, todo Todo) error {
	var currentTodo Todo
	objectId, objErr := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objectId}}
	err := todoCollection.FindOne(context.TODO(), filter).Decode(&currentTodo)
	if objErr == nil {
		if err == nil {
			if todo.Name != currentTodo.Name && todo.Name != "" {
				_, err := todoCollection.UpdateOne(context.TODO(), filter, bson.M{
					"$set": bson.M{
						"name":       todo.Name,
						"updated_at": time.Now()}})
				if err != nil {
					log.Println(err)
					return errors.New(Errors.DBPROCESSINGERROR)
				}
			}
			if todo.Completed != currentTodo.Completed {
				_, err := todoCollection.UpdateOne(context.TODO(), filter, bson.M{
					"$set": bson.M{
						"completed":  todo.Completed,
						"updated_at": time.Now()}})
				if err != nil {
					log.Println(err)
					return errors.New(Errors.DBPROCESSINGERROR)
				}
			}
		} else {
			log.Println(err)
			return errors.New(Errors.NOTFOUNDERROR)
		}
	} else {
		log.Println(objErr)
		return errors.New(Errors.BADREQUESTERROR)
	}
	return nil
}
