package domain

import (
	"context"
	"ticket/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (ticket *Ticket) AddTrain() (*mongo.InsertOneResult, *utils.Resterr) {
	ticketD := db.Collection("ticket")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	trainnoCount, _ := ticketD.CountDocuments(ctx, bson.M{"trainno": ticket.TrainNO})

	defer cancel()
	if trainnoCount > 0 {
		return nil, utils.BadRequest("Train No Already Register")
	}

	result, err := ticketD.InsertOne(ctx, ticket)

	if err != nil {
		restErr := utils.InternalErr("can't insert Train to the database.")
		return nil, restErr
	}

	return result, nil
}

func (ticket *Ticket) GetTrains() *utils.Resterr {

	ticketD := db.Collection("ticket")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	filter := bson.M{"trainno": ticket.TrainNO}
	err := ticketD.FindOne(ctx, filter).Decode(&ticket)

	defer cancel()

	if err != nil {
		return utils.NotFound("Train is Not Found")
	}

	return nil

}

func (ticket *Ticket) GetAllTrain() ([]Ticket, *utils.Resterr) {

	ticketD := db.Collection("ticket")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	cur, errs := ticketD.Find(ctx, bson.M{})

	if errs != nil {
		return nil, utils.NotFound("Trains is Not Found")
	}

	var tickets []Ticket
	defer cancel()

	if err := cur.All(context.TODO(), &tickets); err != nil {
		return nil, utils.NotFound("Trains is Not Found")
	}

	return tickets, nil

}

func (ticket *Ticket) Delete() (*mongo.DeleteResult, *utils.Resterr) {

	ticketD := db.Collection("ticket")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	filter := bson.M{"trainno": ticket.TrainNO}
	result, err := ticketD.DeleteOne(ctx, filter)
	defer cancel()
	if result.DeletedCount == 0 {
		return nil, utils.BadRequest("Train Record Not Found")
	}

	if err != nil {
		return nil, utils.NotFound("Train No is Not Found")
	}

	return result, nil

}

func (ticket *Ticket) Update() (*mongo.UpdateResult, *utils.Resterr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	ticketD := db.Collection("ticket")

	filter := bson.M{"trainno": ticket.TrainNO}

	trainnoCount, _ := ticketD.CountDocuments(ctx, bson.M{"trainno": ticket.TrainNO})

	if trainnoCount == 0 {
		return nil, utils.NotFound("Train Not Found")
	}

	updateValue := bson.M{"$set": bson.M{"trainavaliable": ticket.TrainAvaliable}}

	opts := options.Update().SetUpsert(true)

	result, err := ticketD.UpdateOne(ctx, filter, updateValue, opts)

	defer cancel()

	if result.ModifiedCount == 0 {
		return nil, utils.BadRequest("Ticket not modified")
	}

	if err != nil {
		return nil, utils.InternalErr("Data not Updated")
	}

	return result, nil

}
