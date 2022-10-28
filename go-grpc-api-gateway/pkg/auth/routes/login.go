package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/theninjacoder-uz/go-grpc-api-gateway/pkg/auth/pb"
	"net/http"
)

type LoginRequestBody struct {
	Email    string `json:"emiail"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	body := LoginRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
