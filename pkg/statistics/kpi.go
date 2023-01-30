package statistics

import (
	"github.com/jamesstandbridge/rmc-lottery-simulation/pkg/models"
)

func CountNfTypes(nfts []*models.NFT) []int {
	count := make([]int, 4)
	for i := 0; i < len(nfts); i++ {
		count[nfts[i].Type]++
	}
	return count
}
