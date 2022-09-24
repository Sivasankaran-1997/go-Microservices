package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"users/domain"
	"users/utils"

	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user domain.User) (*mongo.InsertOneResult, *utils.Resterr) {
	if err := user.Vaildate(); err != nil {
		return nil, err
	}
	guid := xid.New()
	user.UserID = guid.String()
	pass := utils.HashPasswordMD5(user.UserPassword)
	user.UserPassword = pass
	insertNo, restErr := user.Create()
	if restErr != nil {
		return nil, restErr
	}
	return insertNo, nil
}

func GetUser(email string, password string) (*domain.User, *utils.Resterr) {

	result := &domain.User{UserEmail: email}

	if err := result.FindUser(); err != nil {
		return nil, err
	}

	pass := result.UserPassword
	passcheck := utils.CheckHash(pass, password)

	if !passcheck {
		return nil, utils.BadRequest("Invalid Password")
	}
	result.UserPassword = ""
	return result, nil
}

func DeleteUser(email string, password string) (*mongo.DeleteResult, *utils.Resterr) {

	result := &domain.User{UserEmail: email}

	if err := result.FindUser(); err != nil {
		return nil, err
	}

	pass := result.UserPassword
	passcheck := utils.CheckHash(pass, password)

	if !passcheck {
		return nil, utils.BadRequest("Invalid Password")
	}

	deleteID, err := result.Delete()

	if err != nil {
		return nil, err
	}

	return deleteID, nil
}

func UpdateUser(isParital bool, user domain.User, email string, password string) (*mongo.UpdateResult, *utils.Resterr) {

	result := &domain.User{UserEmail: email}

	if err := result.FindUser(); err != nil {
		return nil, err
	}

	pass := result.UserPassword
	passcheck := utils.CheckHash(pass, password)

	if !passcheck {
		return nil, utils.BadRequest("Invalid Password")
	}

	if isParital {
		if user.UserName != "" {
			result.UserName = user.UserName
		}

		if user.UserEmail != "" {
			result.UserEmail = user.UserEmail
		}
		if user.UserPhone != "" {
			result.UserPhone = user.UserPhone
		}
	} else {
		result.UserName = user.UserName
		result.UserEmail = user.UserEmail
		result.UserPhone = user.UserPhone
	}

	updateID, err := result.Update()
	if err != nil {
		return nil, err
	}

	return updateID, nil

}

func UserBuyTicket(buy domain.Buy, email string, password string) (*domain.Buy, *utils.Resterr) {

	result := &domain.User{UserEmail: email}

	if err := result.FindUser(); err != nil {
		return nil, err
	}

	pass := result.UserPassword
	passcheck := utils.CheckHash(pass, password)

	if !passcheck {
		return nil, utils.BadRequest("Invalid Password")
	}

	buy.UserID = result.UserID

	client := &http.Client{}
	postuserbuyeapi := os.Getenv("USERBUYTICKET")

	postBody, _ := json.Marshal(map[string]string{
		"userid":     buy.UserID,
		"trainno":    buy.TrainNo,
		"userticket": buy.UserTicket,
	})

	responseBody := bytes.NewBuffer(postBody)

	userbuyAPI, _ := http.NewRequest("POST", postuserbuyeapi, responseBody)
	resp, _ := client.Do(userbuyAPI)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		json.NewDecoder(resp.Body).Decode(&buy)
		return &domain.Buy{UserID: buy.UserID, TrainNo: buy.TrainNo, UserTicket: buy.UserTicket}, nil
	} else {
		return nil, utils.InternalErr("can't User buy the ticket.")
	}

}

func UserCancel(email string, password string, userticket string) *utils.Resterr {

	result := &domain.User{UserEmail: email}

	if err := result.FindUser(); err != nil {
		return err
	}

	pass := result.UserPassword
	passcheck := utils.CheckHash(pass, password)

	if !passcheck {
		return utils.BadRequest("Invalid Password")
	}

	client := &http.Client{}
	postusercancelapi := os.Getenv("USERCANCELTICKET")

	usercancelAPI, _ := http.NewRequest("DELETE", postusercancelapi+userticket, nil)
	resp, _ := client.Do(usercancelAPI)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	} else {
		return utils.InternalErr("can't User Delete the ticket.")
	}

}
