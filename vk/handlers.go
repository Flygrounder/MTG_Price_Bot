package vk

import (
	"github.com/flygrounder/mtg-price-vk/cardsinfo"
	"github.com/gin-gonic/gin"
	"net/http"
)

const CARDSLIMIT = 8

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func HandleMessage(c *gin.Context) {
	var req MessageRequest
	c.BindJSON(&req)
	if (req.Type == "confirmation") && (req.GroupId == GROUPID) {
		c.String(http.StatusOK, CONFIRMATION_STRING)
		return
	}
	defer c.String(http.StatusOK, "ok")
	if req.Secret != SECRET_KEY {
		return
	}
	cardName := cardsinfo.GetOriginalName(req.Object.Body)
	if cardName == "" {
		Message(req.Object.UserId, "Карта не найдена")
	} else {
		prices, _ := cardsinfo.GetSCGPrices(cardName)
		elements := min(CARDSLIMIT, len(prices))
		prices = prices[:elements]
		priceInfo := cardsinfo.FormatCardPrices(cardName, prices)
		Message(req.Object.UserId, priceInfo)
	}
}
