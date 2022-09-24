package domain

import (
	"strings"
	"users/utils"
)

type User struct {
	UserID       string `json:userid"`
	UserName     string `json:username bson:"username,omitempty"`
	UserEmail    string `json:useremail bson:"useremail,omitempty"`
	UserPassword string `json:userpassword bson:"userpassword,omitempty"`
	UserPhone    string `json:userphone bson:"userphone,omitempty"`
}

type Ticket struct {
	TrainNO          string `json:"trainno"`
	TrainName        string `json:trainname bson:"trainname,omitempty"`
	TrainSource      string `json:trainsource bson:"trainsource,omitempty"`
	TrainDestination string `json:traindestination bson:"traindestination,omitempty"`
	TrainAvaliable   string `json:trainavaliable bson:"trainavaliable,omitempty"`
}

type Buy struct {
	BuyID      string `json:"buyid"`
	UserID     string `json:"userid"`
	TrainNo    string `json:trainno bson:"trainno,omitempty"`
	UserTicket string `json:userticket bson:"userticket,omitempty"`
}

func (user *User) Vaildate() *utils.Resterr {
	if strings.TrimSpace(user.UserName) == "" {
		return utils.BadRequest("User Name Required")
	}

	if strings.TrimSpace(user.UserEmail) == "" {
		return utils.BadRequest("User Email  Required")
	}

	if strings.TrimSpace(user.UserPassword) == "" {
		return utils.BadRequest("User PassWord Required")
	}

	if strings.TrimSpace(user.UserPhone) == "" {
		return utils.BadRequest("User Phone Required")
	}

	return nil
}
