package main

import (
	"api-key-service/controller"
	"api-key-service/entity"
	"api-key-service/repository"
	"api-key-service/service"

	"github.com/gin-gonic/gin"
)

var (
	keyRepository repository.ApiKeyRepository = repository.NewRepository()
	keyService    service.ApiKeyService       = service.NewService(keyRepository)
	keyController controller.ApiKeyController = controller.NewController(&keyService)
)

func main() {
	server := gin.Default()

	// Key generation endpoint
	server.POST("keys", func(ctx *gin.Context) {
		var request entity.GenerateKeyRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.AbortWithError(400, err)
			return
		}

		key, err := keyController.GenerateKey(request)
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}

		ctx.JSON(200, key)
	})

	// Owner keys endpoint
	server.GET("keys", func(ctx *gin.Context) {
		var request entity.FindKeysRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.AbortWithError(400, err)
			return
		}

		keys, err := keyController.FindKeys(request)
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}

		ctx.JSON(200, keys)
	})

	// Sample authentication endpoint
	server.GET("auth", func(ctx *gin.Context) {
		var request entity.AuthenticateRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.AbortWithError(400, err)
			return
		}

		success, err := keyController.Authenticate(request)
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}

		if success {
			ctx.Writer.WriteHeader(200)
		} else {
			ctx.Writer.WriteHeader(403)
		}
	})

	server.Run(":8080")
}
