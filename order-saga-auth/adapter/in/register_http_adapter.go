package drivingadapter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	appcommon "order-saga-user/common"
	drivingport "order-saga-user/internal/port/in"
	domainusermodel "order-saga-user/internal/port/in/model"
)

type RegisterHttpAdapter struct {
	useCase drivingport.RegisterUseCase
}

func NewRegisterHttpAdapter(useCase drivingport.RegisterUseCase) *RegisterHttpAdapter {
	return &RegisterHttpAdapter{
		useCase: useCase,
	}
}

func (adapter *RegisterHttpAdapter) HandleRegister() gin.HandlerFunc {
	return func(context *gin.Context) {
		var body domainusermodel.UserRegisterDto

		if err := context.ShouldBind(&body); err != nil {
			context.JSON(http.StatusBadRequest, "")
			return
		}

		credentials, err := adapter.useCase.HandleUserRegister(body.ToUserDomainModel())

		if err != nil {
			context.JSON(http.StatusUnauthorized, appcommon.ErrorCannotCreateEntity("users", err))
			return
		}

		context.JSON(http.StatusOK, credentials)
		return
	}
}
