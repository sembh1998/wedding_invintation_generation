package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sembh1998/wedding_invitation_generation/cmd/bootstrap"
	"github.com/sembh1998/wedding_invitation_generation/internal/core/services/guestssrv"
	"github.com/sembh1998/wedding_invitation_generation/internal/core/services/userssrv"
	"github.com/sembh1998/wedding_invitation_generation/internal/handlers/frontend"
	"github.com/sembh1998/wedding_invitation_generation/internal/handlers/guesthdl"
	"github.com/sembh1998/wedding_invitation_generation/internal/handlers/tokenrequired"
	"github.com/sembh1998/wedding_invitation_generation/internal/handlers/userhdl"
	"github.com/sembh1998/wedding_invitation_generation/internal/repositories/guestsrepo"
	"github.com/sembh1998/wedding_invitation_generation/internal/repositories/usersrepo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	//host 150.136.101.161 port 3306 user root password firevivaldixdzzz
	format   = "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	host     = "150.136.101.161"
	port     = 3306
	user     = "root"
	password = "firevivaldixdzzz"
	dbname   = "wedding"
)

func main() {

	if err := bootstrap.ExtractEmbeddedFile(); err != nil {
		panic(err)
	}

	if err := bootstrap.ExtractAssetsFolder(); err != nil {
		panic(err)
	}

	connstring := fmt.Sprintf(format, user, password, host, port, dbname)
	mysqlgorm, err := gorm.Open(mysql.Open(connstring), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	userRepository := usersrepo.NewUserMysql(mysqlgorm)
	userService := userssrv.New(userRepository)
	userHandler := userhdl.New(userService)

	guestRepository := guestsrepo.NewGuestMysql(mysqlgorm)
	guestService := guestssrv.New(guestRepository)
	guestHandler := guesthdl.New(guestService)

	frontendHandler := frontend.New(guestService, userService)

	router := gin.Default()

	// Add a middleware to handle 404 responses
	router.NoRoute(func(c *gin.Context) {
		// Select a random joke
		randomIndex := rand.Intn(len(bootstrap.Jokes))
		randomJoke := bootstrap.Jokes[randomIndex]

		c.String(http.StatusNotFound, randomJoke)
	})

	router.GET("/", frontendHandler.LoginHTMX)
	router.POST("/login", frontendHandler.Login)
	router.GET("/crudguests", tokenrequired.TokenRequiredMiddleware(), frontendHandler.CrudGuestsHTMX)
	//post add-guest
	router.POST("/add-guest", tokenrequired.TokenRequiredMiddleware(), frontendHandler.AddGuest)
	//delete delete-guest
	router.DELETE("/delete-guest/:id", tokenrequired.TokenRequiredMiddleware(), frontendHandler.DeleteGuest)
	//get guest
	router.GET("/guest/:id", frontendHandler.FetchGuestHTMX)

	//post attend
	router.POST("/guest/:id/attend", frontendHandler.AttendConfirmation)

	//gift-preferences
	router.GET("/gift-preferences", frontendHandler.GiftPreferencesHTMX)

	//assets
	router.StaticFS("/assets", http.Dir("assets"))

	//api v1
	v1 := router.Group("/api/v1")

	v1.POST("/login", userHandler.Login)

	v1.POST("/guest", tokenrequired.BearerTokenRequiredMiddleware(), guestHandler.CreateGuest)
	v1.GET("/guest", tokenrequired.BearerTokenRequiredMiddleware(), guestHandler.FindGuests)
	v1.PUT("/guest/:id", tokenrequired.BearerTokenRequiredMiddleware(), guestHandler.UpdateGuest)
	v1.DELETE("/guest/:id", tokenrequired.BearerTokenRequiredMiddleware(), guestHandler.DeleteGuest)

	v1.GET("/guest/:id", guestHandler.FetchGuest)
	//AttendConfirmation
	v1.PUT("/guest/:id/attend", guestHandler.AttendConfirmation)

	router.Run(":8080")
}
