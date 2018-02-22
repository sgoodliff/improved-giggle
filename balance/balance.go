package balance

import (
	"math/rand"
	"strconv"

	log "github.com/sirupsen/logrus"
)

//PushBalance sends it back to the client
func PushBalance(userid int, balance int) {
	log.Debug("Pushing " + strconv.Itoa(balance) + " to " + strconv.Itoa(userid))
}

//GetBalance just returns a random number for now
func GetBalance(userid int) int {
	var balance int
	balance = rand.Intn(100)
	log.Debug("Balance for " + strconv.Itoa(userid) + " = " + strconv.Itoa(balance))
	return balance
}
