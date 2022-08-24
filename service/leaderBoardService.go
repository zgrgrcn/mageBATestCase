package service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mageBATestCase/model/entity"
	"mageBATestCase/repository"
)

type LeaderboardService struct{}

func (u LeaderboardService) FindAll() ([]entity.Player, error) {
	leaderBoardList, err := repository.FindAllLeaderBoard()
	if err != nil {
		return nil, err
	}
	return leaderBoardList, nil
}

func (u LeaderboardService) ValidateUserList(playerList []entity.Player) error {

	for _, value := range playerList {
		var objID, err = primitive.ObjectIDFromHex(value.UserID)
		if err != nil {
			return errors.New("`user_id` is not valid or missing, error: " + err.Error())
		}
		_, err = UserService{}.FindByUserID(objID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u LeaderboardService) PutResults(playerList []entity.Player) error {
	for _, value := range playerList {
		err := repository.PutLeaderBoard(value)
		if err != nil {
			return err
		}
	}
	return nil
}
