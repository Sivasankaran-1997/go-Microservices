package domain

import (
	"buy/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

func (buy *Buy) BuyTicket() (*mongo.InsertOneResult, *utils.Resterr) {
	buyD := db.Collection("buy")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	result, err := buyD.InsertOne(ctx, buy)
	defer cancel()

	if err != nil {
		restErr := utils.InternalErr("can't insert Train to the database.")
		return nil, restErr
	}

	return result, nil
}

func (buy *Buy) GetUserBuy() *utils.Resterr {

	buyD := db.Collection("buy")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	filter := bson.M{"buyid": buy.BuyID}
	err := buyD.FindOne(ctx, filter).Decode(&buy)

	defer cancel()

	if err != nil {
		return utils.NotFound("User Buy is Not Found")
	}

	return nil

}

func (buy *Buy) Delete() (*mongo.DeleteResult, *utils.Resterr) {

	buyD := db.Collection("buy")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	filter := bson.M{"buyid": buy.BuyID}
	result, err := buyD.DeleteOne(ctx, filter)
	defer cancel()
	if result.DeletedCount == 0 {
		return nil, utils.BadRequest("User Buy Not Found")
	}

	if err != nil {
		return nil, utils.NotFound("User Buy is Not Found")
	}

	return result, nil

}

func (buy *Buy) GetAllUserTicket() ([]Buy, *utils.Resterr) {

	buyD := db.Collection("buy")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	filter := bson.M{"userid": buy.UserID}

	cur, errs := buyD.Find(ctx, filter)

	if errs != nil {
		return nil, utils.NotFound("Users Tickets is Not Found")
	}

	var buys []Buy
	defer cancel()

	if err := cur.All(context.TODO(), &buys); err != nil {
		return nil, utils.NotFound("Users Ticket is Not Found")
	}

	return buys, nil
}
