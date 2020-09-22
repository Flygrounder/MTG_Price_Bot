package tests

import (
	"fmt"
	"github.com/flygrounder/go-mtg-vk/cardsinfo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	prices, err := cardsinfo.GetPrices("Black lotus")
	fmt.Println(prices)
	assert.Nil(t, err)
	assert.NotEmpty(t, prices)
}
