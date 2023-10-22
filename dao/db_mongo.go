package dao

import (
	"context"
	"kallme/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client   *mongo.Client
	database *mongo.Database
)

func InitMongoDB() {
	var err error
	clientOption := options.Client().ApplyURI("mongodb://" + config.Config.MongoDB.Host + ":" + config.Config.MongoDB.Port)

	client, err = mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		config.Log.Error("Fail to connect to MongoDB, error: ", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		config.Log.Error("Fail to ping to MongoDB, error: ", err)
	}

	database = client.Database(config.Config.MongoDB.Database)
}

func CheckDataExist(CollectionName string, Key string, Value string) (bool, error) {
	collection := database.Collection(CollectionName)
	filter := bson.M{Key: Value}
	err := collection.FindOne(context.Background(), filter).Err()
	config.Log.Error("CheckDataExist: ", err)
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func GetDataByKey(CollectionName string, Key string, Value string) (*mongo.Cursor, error) {
	collection := database.Collection(CollectionName)
	filter := bson.M{Key: Value}
	return collection.Find(context.Background(), filter)
}

func InsertData(CollectionName string, data interface{}) (*mongo.InsertOneResult, error) {
	collection := database.Collection(CollectionName)
	return collection.InsertOne(context.Background(), data)
}

func UpdateData(CollectionName string, Key string, Value string, data interface{}) (*mongo.UpdateResult, error) {
	collection := database.Collection(CollectionName)
	filter := bson.M{Key: Value}
	return collection.UpdateOne(context.Background(), filter, bson.M{"$set": data})
}
