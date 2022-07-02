package controllers

import (
	"main/models"
	"main/utils"
	"net/http"
)

type Req struct {
	Name         string
	Participants []string
}

func CreateBill(_ http.ResponseWriter, r *http.Request) {
	body := utils.GetRequestBody[Req](r.Body)
	billName := body.Name
	participants := body.Participants
	models.CreateBill(billName, participants)
}
