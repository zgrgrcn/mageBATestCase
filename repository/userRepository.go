package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mageBATestCase/model/db"
	"mageBATestCase/model/dto"
	"mageBATestCase/model/entity"
)

type UserRepository interface {
	CrateUser(userPayload *dto.UserRequest) (*mongo.InsertOneResult, error)
	FindUser(user *dto.UserRequest) ([]entity.User, error)
	FindByUserID(userId primitive.ObjectID) ([]entity.User, error)
	FindByUserName(userName string) ([]entity.User, error)
}

func CrateUser(userPayload *dto.UserRequest) (*mongo.InsertOneResult, error) {
	client, ctx := db.GetConnection()
	userCollection := client.Database("mage").Collection("user")

	result, err := userCollection.InsertOne(ctx, userPayload)
	defer client.Disconnect(ctx)
	return result, err
}

func FindUser(user *dto.UserRequest) ([]entity.User, error) {
	client, ctx := db.GetConnection()
	userCollection := client.Database("mage").Collection("user")
	cursor, err := userCollection.Find(ctx, bson.M{"username": user.Username})
	defer client.Disconnect(ctx)
	if err != nil {
		return nil, err
	}
	var users []entity.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func FindByUserID(userId primitive.ObjectID) ([]entity.User, error) {
	client, ctx := db.GetConnection()
	userCollection := client.Database("mage").Collection("user")
	var users []entity.User
	cursor, err := userCollection.Find(ctx, bson.M{"_id": userId})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	return users, nil
}

func FindByUserName(userName string) ([]entity.User, error) {
	client, ctx := db.GetConnection()
	userCollection := client.Database("mage").Collection("user")
	var users []entity.User
	cursor, err := userCollection.Find(ctx, bson.M{"username": userName})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	return users, nil
}
