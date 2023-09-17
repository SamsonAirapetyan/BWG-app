package handler

import (
	"github.com/SamsonAirapetyan/BWG-test"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllListRequest struct {
	Data []BWG_test.Answer `json:"data"`
}

func (h *Handler) Add_into_wallet(c *gin.Context) {
	var input BWG_test.Request

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Problem with request")
		return
	}

	err := h.services.Transactions.AddSum(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) Take_from_wallet(c *gin.Context) {
	var input BWG_test.Request

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Problem with request")
		return
	}

	err := h.services.Transactions.TakeOff(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) GetAll(c *gin.Context) {
	list, err := h.services.Transactions.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllListRequest{
		Data: list,
	})
}
