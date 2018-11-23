package webapp

import (
	"net/http"
	"pusher/blackjack/internal/chatter"
	"pusher/blackjack/internal/game"
	"pusher/blackjack/internal/notifier"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer(chatter *chatter.Chatter, notifier *notifier.Notifier, games *game.Games) {
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "PUT", "POST", "DELETE"}
	r.Use(cors.New(corsConfig))
	r.POST("/chatkit/auth", func(c *gin.Context) {
		userID := c.Query("user_id")

		authRes, err := chatter.Authenticate(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			c.JSON(http.StatusOK, authRes.TokenResponse())
		}
	})

	r.GET("/games/:id", func(c *gin.Context) {
		game := games.GetGame(c.Param("id"))
		c.JSON(http.StatusOK, game)
	})

	r.PUT("/games/:id/:player", func(c *gin.Context) {
		game := games.GetGame(c.Param("id"))
		player := c.Param("player")
		seat, _ := strconv.ParseUint(c.PostForm("seat"), 10, 16)

		_, err := game.JoinTable(uint16(seat), player)
		if err == nil {
			notifier.Notify(game.ID)
			c.JSON(http.StatusOK, game)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
	})

	r.DELETE("/games/:id/:player", func(c *gin.Context) {
		game := games.GetGame(c.Param("id"))
		player := c.Param("player")

		game.LeaveTable(player)
		notifier.Notify(game.ID)
		c.JSON(http.StatusOK, game)
	})

	r.PUT("/games/:id/:player/bet", func(c *gin.Context) {
		game := games.GetGame(c.Param("id"))
		player := c.Param("player")
		amount, _ := strconv.ParseUint(c.PostForm("amount"), 10, 16)

		game.Bet(player, uint16(amount))
		notifier.Notify(game.ID)
		c.JSON(http.StatusOK, game)
	})

	r.POST("/games/:id/:player/action", func(c *gin.Context) {
		game := games.GetGame(c.Param("id"))
		player := c.Param("player")
		action := c.PostForm("action")

		var err error
		if action == "hit" {
			err = game.Hit(player)
		} else if action == "stick" {
			err = game.Stick(player)
		}

		if err == nil {
			notifier.Notify(game.ID)
			c.JSON(http.StatusOK, game)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
	})
	r.Run()
}
