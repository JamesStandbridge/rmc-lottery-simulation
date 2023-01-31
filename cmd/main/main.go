package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"

	"github.com/ispringteam/eventbus"
	"github.com/jamesstandbridge/rmc-lottery-simulation/pkg/models"
	"github.com/jamesstandbridge/rmc-lottery-simulation/pkg/utils"
)

const POPULATION = 1000

var blockchain models.Blockchain

func main() {
	//initialization
	rand.Seed(time.Now().UnixNano())

	chans := models.BlockchainChanels{
		TestChan: eventbus.New(),
	}

	//http server
	router := gin.New()

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())

		s.Join("blockchain")

		go func() {
			chans.TestChan.Subscribe("event.transaction", func(e eventbus.Event) {
				server.BroadcastToRoom("/", "blockchain", "blockchain", e.EventID())
			})
		}()

		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
	}))

	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	router.StaticFS("/public", http.Dir("../asset"))

	// end http server

	protocolWallet := models.Wallet{Address: "protocol", Balance: 0}

	blockchain = models.Blockchain{
		Wallets: append(utils.GenerateWallets(POPULATION), &protocolWallet),
	}

	go blockchain.Run(chans)
	go protocolWallet.Run(chans)

	for i := 0; i < len(blockchain.Wallets); i++ {
		go blockchain.Wallets[i].Run(chans)
	}

	if err := router.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}

	defer server.Close()
}
