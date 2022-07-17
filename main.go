package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	"google.golang.org/grpc/codes"
)

func main() {
	r := gin.New()

	r.GET("/", Healthy)
	r.POST("/", GiteeInterceptor)

	if err := r.Run(":80"); err != nil {
		log.Fatal(err)
	}
}

func Healthy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func GiteeInterceptor(ctx *gin.Context) {
	req := &v1alpha1.InterceptorRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, &v1alpha1.InterceptorResponse{
			Status: v1alpha1.Status{
				Code:    codes.Unavailable,
				Message: err.Error(),
			},
		})
		return
	}

	token, ok := req.Header["X-Xiaolatiao-Token"]
	if !ok || len(token) == 0 || token[0] != "xiaolatao" {
		ctx.JSON(http.StatusServiceUnavailable, &v1alpha1.InterceptorResponse{
			Status: v1alpha1.Status{
				Code:    codes.Unavailable,
				Message: "Token不正确",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, &v1alpha1.InterceptorResponse{
		Continue: true,
		Status: v1alpha1.Status{
			Code: codes.OK,
		},
	})
}
