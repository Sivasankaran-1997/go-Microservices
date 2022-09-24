package services

import (
	"buy/domain"
	"buy/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/mongo"
)

func BuyTicket(buy domain.Buy) (*mongo.InsertOneResult, *utils.Resterr) {
	if err := buy.Vaildate(); err != nil {
		return nil, err
	}
	guid := xid.New()
	buy.BuyID = guid.String()

	client := &http.Client{}

	getTrainapi := os.Getenv("GETATRAINAPI")
	trainAPI, _ := http.Get(getTrainapi + buy.TrainNo)

	defer trainAPI.Body.Close()

	if trainAPI.StatusCode == http.StatusOK {
		var ticket domain.Ticket
		json.NewDecoder(trainAPI.Body).Decode(&ticket)

		tick, _ := strconv.Atoi(ticket.TrainAvaliable)
		bu, _ := strconv.Atoi(buy.UserTicket)

		if tick >= bu {
			avalticket := tick - bu
			ats := strconv.Itoa(avalticket)
			postBody, _ := json.Marshal(map[string]string{
				"trainavaliable": ats,
			})
			responseBody := bytes.NewBuffer(postBody)

			postupdateapi := os.Getenv("POSTUPDATETRAINAPI")
			updateAPI, _ := http.NewRequest("POST", postupdateapi+buy.TrainNo, responseBody)
			resp, _ := client.Do(updateAPI)

			defer resp.Body.Close()

		} else {
			restErr := utils.InternalErr("Ticket is Not Avaliable")
			return nil, restErr
		}
	} else {
		restErr := utils.NotFound("Train Services is Not Applicable")
		return nil, restErr

	}

	insertNo, restErr := buy.BuyTicket()
	if restErr != nil {
		return nil, restErr
	}
	return insertNo, nil
}

func GetUserBuy(buyid string) (*domain.Buy, *utils.Resterr) {

	result := &domain.Buy{BuyID: buyid}

	if err := result.GetUserBuy(); err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteUserBuy(buyid string) (*mongo.DeleteResult, *utils.Resterr) {

	result := &domain.Buy{BuyID: buyid}

	if err := result.GetUserBuy(); err != nil {
		return nil, err
	}

	client := &http.Client{}

	getTrainapi := os.Getenv("GETATRAINAPI")
	trainAPI, _ := http.Get(getTrainapi + result.TrainNo)

	defer trainAPI.Body.Close()

	if trainAPI.StatusCode == http.StatusOK {
		var ticket domain.Ticket
		json.NewDecoder(trainAPI.Body).Decode(&ticket)

		tick, _ := strconv.Atoi(ticket.TrainAvaliable)
		bu, _ := strconv.Atoi(result.UserTicket)

		if bu > 0 {
			avalticket := tick + bu
			ats := strconv.Itoa(avalticket)
			postBody, _ := json.Marshal(map[string]string{
				"trainavaliable": ats,
			})
			responseBody := bytes.NewBuffer(postBody)

			postupdateapi := os.Getenv("POSTUPDATETRAINAPI")
			updateAPI, _ := http.NewRequest("POST", postupdateapi+result.TrainNo, responseBody)
			resp, _ := client.Do(updateAPI)

			defer resp.Body.Close()

		} else {
			restErr := utils.InternalErr("User Not Buy Ticket")
			return nil, restErr
		}
	} else {
		restErr := utils.NotFound("Train Services is Not Applicable")
		return nil, restErr

	}

	deleteID, err := result.Delete()

	if err != nil {
		return nil, err
	}

	return deleteID, nil
}

func GetAllUserTickets(userid string) ([]domain.Buy, *utils.Resterr) {

	result := domain.Buy{UserID: userid}

	res, err := result.GetAllUserTicket()

	if err != nil {
		return nil, err
	}

	return res, nil
}
