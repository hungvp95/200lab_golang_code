package middleware

import (
	"food-delivery-200lab/common"
	"food-delivery-200lab/component"
	"github.com/gin-gonic/gin"
)

func Recover(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(appErr)
				}

				appErr := common.ErrInternal(err.(error))
				ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(appErr)
			}
		}()

		ctx.Next()
	}
}
