package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/alnshine/go-grpc-api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

const (
	base int = 10
	size int = 32
)

func FineOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), base, size)
	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}
