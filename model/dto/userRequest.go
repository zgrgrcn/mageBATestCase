package dto

type UserRequest struct {
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
}
