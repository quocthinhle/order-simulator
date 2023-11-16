package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	drivingadapter "order-saga-user/adapter/in"
	drivenadapter "order-saga-user/adapter/out/persistence"
	appcommon "order-saga-user/common"
	domainuserservice "order-saga-user/internal/domain/user_auth/use_case"
	"os"
)

func main() {
	httpServer := gin.Default()

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timeZone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot initialize database connection")
	}

	appCtx := appcommon.NewAppCtx(db)

	api := httpServer.Group("api")
	v1 := api.Group("v1")
	authV1 := v1.Group("auth")

	createUserPersistenceAdapter := drivenadapter.NewCreateUserAdapter(appCtx)
	registerUseCase := domainuserservice.NewUserRegisterUseCase(createUserPersistenceAdapter)
	registerHttpAdapter := drivingadapter.NewRegisterHttpAdapter(registerUseCase)

	getUserAdapter := drivenadapter.NewGetUserAdapter(appCtx)
	loginUseCase := domainuserservice.NewUserLoginUseCase(getUserAdapter)
	loginHttpAdapter := drivingadapter.NewLoginHttpAdapter(loginUseCase)

	authV1.POST("register", registerHttpAdapter.HandleRegister())
	authV1.POST("login", loginHttpAdapter.HandleLogin())

	httpServer.Run()
}
