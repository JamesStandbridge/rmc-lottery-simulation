package models

//create a new model for the lottery. 
//A lottery is the main class of this application, we can create only one lottery at a time.

type Lottery struct {
	MaxDay int
	CurrentDay int
	Tickets []NFT
	TicketPrice float64
}

//create interface for the lottery
type LotteryInterface interface {
	//the lottery can be initialized
	Initialize(maxDay int)
	//the lottery can be played
	Play()
}

//initialize the lottery
func (lottery *Lottery) Initialize(maxDay int, caracteristicsMapping []int) {
	lottery.MaxDay = maxDay
	lottery.CurrentDay = 0
}