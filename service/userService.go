package service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"mageBATestCase/model/db"
	"mageBATestCase/model/dto"
	"mageBATestCase/model/entity"
)

type Userservice struct{}

func (u Userservice) FindByUserID(userId primitive.ObjectID) (entity.User, error) {
	//TODO: check here
	client, ctx := db.GetConnection()
	userCollection := client.Database("mage").Collection("user")
	var users []entity.User
	cursor, err := userCollection.Find(ctx, bson.M{"_id": userId})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &users); err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	if len(users) == 0 {
		return entity.User{}, errors.New("user not found with the given id: " + userId.String())
	}
	return users[0], nil
}

func (u Userservice) FindByUserName(userName string) (entity.User, error) {
	client, ctx := db.GetConnection()
	userCollection := client.Database("mage").Collection("user")
	var users []entity.User
	cursor, err := userCollection.Find(ctx, bson.M{"username": userName})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &users); err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	return users[0], nil
}

func (u Userservice) Create(userPayload *dto.UserRequest) (entity.User, error) {
	dbUser := entity.User{}
	userPayload.Password, _ = encryptPassword(userPayload.Password)
	client, ctx := db.GetConnection()
	userCollection := client.Database("mage").Collection("user")

	result, err := userCollection.InsertOne(ctx, userPayload)
	if err != nil {
		return entity.User{}, err
	}
	dbUser.ID = result.InsertedID.(primitive.ObjectID)
	dbUser.Username = userPayload.Username
	dbUser.Password = userPayload.Password
	defer client.Disconnect(ctx)
	return dbUser, nil
}

func (u Userservice) Find(user *dto.UserRequest) (entity.User, error) {
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
	if len(users) > 1 {
		return entity.User{}, errors.New("there are multiple users with the same username")
	}
	if len(users) == 0 {
		return entity.User{}, errors.New("there is no users with the same username")
	}

	if !isSamePassword(user.Password, users[0].Password) {
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
	return err == nil || password == hash //no error or password could be already hashed
}
