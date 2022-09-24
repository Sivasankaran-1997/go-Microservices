package services

import (
	"ticket/domain"
	"ticket/utils"

	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddTrains(ticket domain.Ticket) (*mongo.InsertOneResult, *utils.Resterr) {
	if err := ticket.Vaildate(); err != nil {
		return nil, err
	}
	guid := xid.New()
	ticket.TrainID = guid.String()
	insertNo, restErr := ticket.AddTrain()
	if restErr != nil {
		return nil, restErr
	}
	return insertNo, nil
}

func GetTrain(trainno string) (*domain.Ticket, *utils.Resterr) {

	result := &domain.Ticket{TrainNO: trainno}

	if err := result.GetTrains(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetAllTrain() ([]domain.Ticket, *utils.Resterr) {

	var ticket domain.Ticket
	result, resterr := ticket.GetAllTrain()

	if resterr != nil {
		return nil, resterr
	}

	return result, nil
}

func DeleteTrain(trainno string) (*mongo.DeleteResult, *utils.Resterr) {

	result := &domain.Ticket{TrainNO: trainno}

	if err := result.GetTrains(); err != nil {
		return nil, err
	}

	deleteID, err := result.Delete()

	if err != nil {
		return nil, err
	}

	return deleteID, nil
}

func UpdateTrain(isParital bool, ticket domain.Ticket, trainno string) (*mongo.UpdateResult, *utils.Resterr) {

	result := &domain.Ticket{TrainNO: trainno}

	if err := result.GetTrains(); err != nil {
		return nil, err
	}

	if isParital {
		if ticket.TrainName != "" {
			result.TrainName = ticket.TrainName
		}

		if ticket.TrainSource != "" {
			result.TrainSource = ticket.TrainSource
		}

		if ticket.TrainDestination != "" {
			result.TrainDestination = ticket.TrainDestination
		}

		if ticket.TrainAvaliable != "" {
			result.TrainAvaliable = ticket.TrainAvaliable
		}
	} else {
		result.TrainName = ticket.TrainName
		result.TrainSource = ticket.TrainSource
		result.TrainDestination = ticket.TrainDestination
		result.TrainAvaliable = ticket.TrainAvaliable
	}

	updateID, err := result.Update()
	if err != nil {
		return nil, err
	}

	return updateID, nil

}
