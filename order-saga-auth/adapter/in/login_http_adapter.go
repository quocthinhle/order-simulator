package drivingadapter

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"net/http"
	appcommon "order-saga-user/common"
	drivingport "order-saga-user/internal/port/in"
	domainusermodel "order-saga-user/internal/port/in/model"
	"time"
)

type LoginHttpAdapter struct {
	loginUseCase drivingport.LoginUseCase
	AppContext   *appcommon.AppContext
}

func NewLoginHttpAdapter(useCase drivingport.LoginUseCase) *LoginHttpAdapter {
	return &LoginHttpAdapter{
		loginUseCase: useCase,
	}
}

func (adapter *LoginHttpAdapter) HandleLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		var loginData domainusermodel.UserLoginDto

		if err := context.ShouldBind(&loginData); err != nil {
			context.JSON(http.StatusBadRequest, "")
			return
		}

		user, err := adapter.loginUseCase.HandleLoginWithUsername(loginData)

		if err != nil {
			context.JSON(http.StatusUnauthorized, "")
			return
		}

		userViewPayload := (&domainusermodel.UserViewPayload{}).FromDomainModel(user)

		jwt, err := appcommon.NewJwt().Sign(appcommon.SignOption{
			ExpiredIn: time.Hour * 24,
			Alg:       jwt2.SigningMethodRS256,
			Payload:   structs.Map(userViewPayload),
		})

		fmt.Println(err)

		context.JSON(http.StatusOK, jwt)
		return
	}
}
