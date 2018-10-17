package balance

import (
	"math/rand"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func balanceHomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/home" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "html/home.html")
}

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
