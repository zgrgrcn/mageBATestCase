package service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"mageBATestCase/model/dto"
	"mageBATestCase/model/entity"
	"mageBATestCase/repository"
)

type UserService struct{}

func (u UserService) FindByUserID(userId primitive.ObjectID) (entity.User, error) {
	users, err := repository.FindByUserID(userId)
	if err != nil {
		return entity.User{}, err
	}
	if len(users) == 0 {
		return entity.User{}, errors.New("user not found with the given id: " + userId.String())
	}
	return users[0], nil
}

func (u UserService) FindByUserName(userName string) (entity.User, error) {
	users, err := repository.FindByUserName(userName)
	if err != nil {
		return entity.User{}, err
	}
	return users[0], nil
}

func (u UserService) Create(userPayload *dto.UserRequest) (entity.User, error) {
	dbUser := entity.User{}
	userPayload.Password, _ = encryptPassword(userPayload.Password)
	result, err := repository.CrateUser(userPayload)
	if err != nil {
		return entity.User{}, err
	}
	dbUser.ID = result.InsertedID.(primitive.ObjectID)
	dbUser.Username = userPayload.Username
	dbUser.Password = userPayload.Password

	return dbUser, nil
}

func (u UserService) Find(user *dto.UserRequest) (entity.User, error) {
	users, err := repository.FindUser(user)
	if err != nil {
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
