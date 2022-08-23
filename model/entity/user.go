package entity

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mage/utils"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
}

func (user *User) GetJwtToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Username": string(user.Username),
	})
	secretKey := utils.GetEnvVariable("TOKEN_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}
