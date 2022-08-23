package service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"mageBATestCase/model/db"
	"mageBATestCase/model/entity"
)

type Userservice struct{}

func (u Userservice) Create(userPayload *entity.User) (entity.User, error) {
	userPayload.Password, _ = encryptPassword(userPayload.Password)
	client, ctx := db.GetConnection()
	userCollection := client.Database("mage").Collection("user")

	result, err := userCollection.InsertOne(ctx, userPayload)
	if err != nil {
		return entity.User{}, err
	}
	userPayload.ID = result.InsertedID.(primitive.ObjectID)
	defer client.Disconnect(ctx)
	return *userPayload, nil
}

func (u Userservice) Find(user *entity.User) (entity.User, error) {
	client, ctx := db.GetConnection()
	userCollection := client.Database("mage").Collection("user")

	var users []entity.User
	cursor, err := userCollection.Find(ctx, bson.M{"username": user.Username})
	if err != nil {
		return entity.User{}, err
	}
	if err = cursor.All(ctx, &users); err != nil {
		return entity.User{}, err
	}
	if len(users) != 1 {
		return entity.User{}, errors.New("there are multiple users with the same username")
	}

	if isSamePassword(user.Password, users[0].Password) {
		return entity.User{}, errors.New("password is not correct")
	}

	defer client.Disconnect(ctx)
	return users[0], nil
}

func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func isSamePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}
