package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = mapp[string]LoginDetails{
	"alex": {
		AuthToen: "123ABC"
		Username: "alex"
	},
	"ajasonex": {
		AuthToen: "456DEF"
		Username: "jason"
	},
	"marie": {
		AuthToen: "789GHI"
		Username: "marie"
	},
}

var mockCoinsDetails = mapp[string]CoinDetails{
	"alex": {
		Coins: 100
		Username: "alex"
	},
	"ajasonex": {
		Coins: 200
		Username: "jason"
	},
	"marie": {
		Coins: 300
		Username: "marie"
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Se)
}