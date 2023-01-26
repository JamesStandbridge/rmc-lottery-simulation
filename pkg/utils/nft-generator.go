package utils

import (
	"github.com/jamesstandbridge/rmc-lottery-simulation/pkg/models"
)

//create a function which take two parameters :
// the first one is an array of integers. Each index of the array represent a caracteristic of the NFT
// and the value of the index, the number of different values for this caracteristic
// the second parameter is the Wallet owner of all future NFTs
// the function returns an array of NFTs
// Each NFT is a unique combinaison of caracteristics and we should have as many NFTs as possible

func GenerateNFTs(caracteristics []int, owner models.Wallet) []models.NFT {
	//delcare an empty array of NFT
	var nfts []models.NFT
	
	totalCombinaisons := ArrayProduct(caracteristics)

	currCombinaison := ArrFill(make([]int, len(caracteristics)), 1)

	for i := 0; i < totalCombinaisons; i++ {
		print(currCombinaison)
		nft := models.NFT{Caracteristics: currCombinaison, Owner: owner}
		nfts = append(nfts, nft)
		//generate next combinaison
		for j := len(caracteristics) - 1; j >= 0; j-- {
			if currCombinaison[j] < caracteristics[j] {
				currCombinaison[j]++
				break
			} else {
				currCombinaison[j] = 1
			}
		}
	}

	return nfts
}
