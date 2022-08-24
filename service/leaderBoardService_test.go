package service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mageBATestCase/model/db"
	"mageBATestCase/model/entity"
	"testing"
)

func TestFind(t *testing.T) {
	client, ctx := db.GetConnection()
	userCollection := client.Database("mage").Collection("user")

	var userList []entity.User
	cursor, err := userCollection.Find(ctx, bson.M{"username": user.Username})
	if err != nil {
		return entity.User{}, err
	}
	if err = cursor.All(ctx, &userList); err != nil {
		return entity.User{}, err
	}
	if len(userList) != 1 {
		return entity.User{}, errors.New("there are multiple userList with the same username")
	}
	if !isSamePassword(user.Password, userList[0].Password) {
		return entity.User{}, errors.New("password is not correct")
	}

	defer client.Disconnect(ctx)
	return userList[0], nil
}

func (u LeaderboardService) FindAll() ([]entity.Player, error) {
	client, ctx := db.GetConnection()
	leaderboardCollection := client.Database("mage").Collection("leaderboard")

	opts := options.Find().SetSort(bson.D{{"score", -1}})
	cursor, err := leaderboardCollection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, errors.New("user not found" + err.Error())
	}
	var leaderBoardList []entity.Player
	if err = cursor.All(ctx, &leaderBoardList); err != nil {
		return nil, errors.New("user not found" + err.Error())
	}
	defer client.Disconnect(ctx)
	return leaderBoardList, nil
}

func (u LeaderboardService) ValidateUserList(playerList []entity.Player) error {

	for _, value := range playerList {
		var objID, err = primitive.ObjectIDFromHex(value.UserID)
		if err != nil {
			return errors.New("`user_id` is not valid or missing, error: " + err.Error())
		}
		_, err = Userservice{}.FindByUserID(objID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u LeaderboardService) PutResults(playerList []entity.Player) {
	client, ctx := db.GetConnection()
	leaderboardCollection := client.Database("mage").Collection("leaderboard")
	for _, value := range playerList {
		filter := bson.D{{"user_id", value.UserID}}
		update := bson.D{{"$inc", bson.D{{"score", value.Score}}}}
		opts := options.Update().SetUpsert(true)
		result, err := leaderboardCollection.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			panic(err)
		}
	}
	defer client.Disconnect(ctx)

}
