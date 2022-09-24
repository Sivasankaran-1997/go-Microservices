package domain

import (
	"strings"
	"ticket/utils"
)

type Ticket struct {
	TrainID          string `json:trainid"`
	TrainNO          string `json:"trainno"`
	TrainName        string `json:trainname bson:"trainname,omitempty"`
	TrainSource      string `json:trainsource bson:"trainsource,omitempty"`
	TrainDestination string `json:traindestination bson:"traindestination,omitempty"`
	TrainAvaliable   string `json:trainavaliable bson:"trainavaliable,omitempty"`
}

func (ticket *Ticket) Vaildate() *utils.Resterr {
	if strings.TrimSpace(ticket.TrainNO) == "" {
		return utils.BadRequest("Train No Required")
	}

	if strings.TrimSpace(ticket.TrainName) == "" {
		return utils.BadRequest("Train Name Required")
	}

	if strings.TrimSpace(ticket.TrainSource) == "" {
		return utils.BadRequest("Train Source Required")
	}

	if strings.TrimSpace(ticket.TrainDestination) == "" {
		return utils.BadRequest("Train Destination Required")
	}

	if strings.TrimSpace(ticket.TrainAvaliable) == "" {
		return utils.BadRequest("Train Avaliable Required")
	}

	return nil
}
