package domain

import (
	"buy/utils"
	"strings"
)

type Buy struct {
	BuyID      string `json:"buyid"`
	UserID     string `json:"userid"`
	TrainNo    string `json:trainno bson:"trainno,omitempty"`
	UserTicket string `json:userticket bson:"userticket,omitempty"`
}

type Ticket struct {
	TrainID          string `json:trainid"`
	TrainNO          string `json:"trainno"`
	TrainName        string `json:trainname bson:"trainname,omitempty"`
	TrainSource      string `json:trainsource bson:"trainsource,omitempty"`
	TrainDestination string `json:traindestination bson:"traindestination,omitempty"`
	TrainAvaliable   string `json:trainavaliable bson:"trainavaliable,omitempty"`
}

func (buy *Buy) Vaildate() *utils.Resterr {
	if strings.TrimSpace(buy.UserID) == "" {
		return utils.BadRequest("UserID No Required")
	}

	if strings.TrimSpace(buy.TrainNo) == "" {
		return utils.BadRequest("Train No Required")
	}

	if strings.TrimSpace(buy.UserTicket) == "" {
		return utils.BadRequest("User Ticket Required")
	}

	return nil
}
