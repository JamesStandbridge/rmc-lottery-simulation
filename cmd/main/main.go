package main


import (
	"fmt"
	"github.com/jamesstandbridge/rmc-lottery-simulation/pkg/models"
	"github.com/jamesstandbridge/rmc-lottery-simulation/pkg/utils"
)



func main() {
	ownerWallet := models.Wallet{Address: "master", Balance: 100}
	
	nfts := utils.GenerateNFTs([]int{5, 5, 2}, ownerWallet)
	//wallets := utils.GenerateWallets(len(nfts))

	lottery := models.Lottery{
		MaxDay: len(nfts[0].Caracteristics), 
		CurrentDay: 1, 
		Tickets: nfts, 
		TicketPrice: 2.5,
		Owner: ownerWallet,
	}

	fmt.Println(lottery)
}