package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Leaderboard struct {
	User  primitive.ObjectID `bson:"user,omitempty"`
	Score int                `bson:"score,omitempty"`
}
