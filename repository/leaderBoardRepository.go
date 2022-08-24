package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mageBATestCase/model/db"
	"mageBATestCase/model/entity"
)

type LeaderBoardRepository interface {
	FindAllLeaderBoard() ([]entity.Player, error)
	PutLeaderBoard(value entity.Player, leaderboardCollection *mongo.Collection, ctx context.Context) error
}

func PutLeaderBoard(value entity.Player) error {
	client, ctx := db.GetConnection()
	leaderboardCollection := client.Database("mage").Collection("leaderboard")
	filter := bson.D{{"user_id", value.UserID}}
	update := bson.D{{"$inc", bson.D{{"score", value.Score}}}}
	opts := options.Update().SetUpsert(true)
	_, err := leaderboardCollection.UpdateOne(ctx, filter, update, opts)
	defer client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func FindAllLeaderBoard() ([]entity.Player, error) {
	client, ctx := db.GetConnection()
	leaderboardCollection := client.Database("mage").Collection("leaderboard")

	opts := options.Find().SetSort(bson.D{{"score", -1}})
	cursor, err := leaderboardCollection.Find(ctx, bson.M{}, opts)
	defer client.Disconnect(ctx)
	if err != nil {
		return nil, errors.New("user not found" + err.Error())
	}
	var leaderBoardList []entity.Player
	if err = cursor.All(ctx, &leaderBoardList); err != nil {
		return nil, errors.New("user not found" + err.Error())
	}
	return leaderBoardList, nil
}
