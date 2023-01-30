package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/jamesstandbridge/rmc-lottery-simulation/pkg/models"
	"github.com/jamesstandbridge/rmc-lottery-simulation/pkg/utils"
)

const NUMBER_OF_WALLETS = 3000

var blockchain models.Blockchain

func main() {
	//initialization
	rand.Seed(time.Now().UnixNano())
	protocol := models.Wallet{Address: "protocol", Balance: 0}
	owner := models.Wallet{Address: "owner", Balance: 0}
	blockchain = models.Blockchain{Wallets: utils.GenerateWallets(NUMBER_OF_WALLETS)}

	/**	LOTTERY	1 **/
	collection := models.Collection{Caracteristics: []int{5, 5, 6, 5, 2}, GoldDay: 3, Owner: &protocol}

	nfts := utils.GenerateCollectionNfts(&collection)
	blockchain.Nfts = append(blockchain.Nfts, nfts...)

	lottery := models.Lottery{
		TicketPrice: 2.5,
		Owner:       &protocol,
		Collection:  &collection,
		Tickets:     nfts,
	}

	for i := 0; i < len(blockchain.Wallets); i++ {
		lottery.BuyTicket(blockchain.Wallets[i], i)
	}

	lottery.Play()

	//print owner balance
	println("Protocol balance before lottery: ", lottery.Owner.Balance)
	println("Owner balance before lottery: ", owner.Balance)

	go blockchain.Run()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called : %s!", r.URL.Path[1:])
}
