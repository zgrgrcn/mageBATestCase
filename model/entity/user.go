package entity

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mageBATestCase/util"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
	Token    string             `bson:"token,omitempty"`
}

func (user *User) GetJwtToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Username": string(user.Username),
	})
	secretKey := util.GetEnvVariable("TOKEN_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

// FOR INDEX AND UNIQUE FOR USERNAME
//_, err := userCollection.Indexes().CreateOne(
//	context.Background(),
//	mongo.IndexModel{
//		Keys:    bson.D{{Key: "username", Value: 1}},
//		Options: options.Index().SetUnique(true),
//	},
//)
