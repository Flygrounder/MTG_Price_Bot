package scenario

import (
	"errors"

	"gitlab.com/flygrounder/go-mtg-vk/internal/cardsinfo"
)

type testInfoFetcher struct{}

func (t *testInfoFetcher) GetPrices(name string) ([]cardsinfo.ScgCardPrice, error) {
	if name == "good" || name == "uncached" {
		return nil, nil
	}
	return nil, errors.New("test")
}

func (t *testInfoFetcher) GetNameByCardId(_ string, _ string) string {
	return "good"
}

func (t *testInfoFetcher) GetOriginalName(name string) string {
	if name == "good" || name == "bad" || name == "uncached" {
		return name
	}
	return ""
}
