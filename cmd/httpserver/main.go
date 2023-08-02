package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sembh1998/wedding_invitation_generation/internal/core/services/userssrv"
	"github.com/sembh1998/wedding_invitation_generation/internal/handlers/userhdl"
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

	connstring := fmt.Sprintf(format, user, password, host, port, dbname)
	mysqlgorm, err := gorm.Open(mysql.Open(connstring), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	userRepository := usersrepo.NewUserMysql(mysqlgorm)
	userService := userssrv.New(userRepository)
	userHandler := userhdl.New(userService)

	router := gin.Default()

	router.POST("/login", userHandler.Login)

	router.Run(":8080")
}
