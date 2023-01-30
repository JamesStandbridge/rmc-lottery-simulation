package utils

import (
	"fmt"

	"github.com/jamesstandbridge/rmc-lottery-simulation/pkg/models"
)

const minWalletBalance = 10
const maxWalletBalance = 20

func GenerateWallets(size int) []*models.Wallet {
	var wallets []*models.Wallet
	for i := 1; i <= size; i++ {
		balance := randFloat(minWalletBalance, maxWalletBalance)
		wallet := models.Wallet{Address: fmt.Sprint("0x", i), Balance: balance}
		wallets = append(wallets, &wallet)
	}
	return wallets
}

func GenerateCollectionNfts(collection *models.Collection) []*models.NFT {
	//delcare an empty array of NFT
	var nfts []*models.NFT

	totalCombinaisons := ArrayProduct(collection.Caracteristics)

	currCombinaison := ArrFill(make([]int, len(collection.Caracteristics)), 1)

	for i := 0; i < totalCombinaisons; i++ {
		nft := models.NFT{
			Caracteristics: append([]int(nil), currCombinaison...),
			Owner:          collection.Owner,
			Type:           0,
			IsActive:       true,
			Collection:     collection,
		}
		nfts = append(nfts, &nft)

		for j := len(collection.Caracteristics) - 1; j >= 0; j-- {
			if currCombinaison[j] < collection.Caracteristics[j] {
				currCombinaison[j]++
				break
			} else {
				currCombinaison[j] = 1
			}
		}
	}

	return nfts
}
