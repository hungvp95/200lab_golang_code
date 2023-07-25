package ginupload

import (
	"fmt"
	"food-delivery-200lab/common"
	"food-delivery-200lab/component"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImageSimple(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header, err := ctx.FormFile("imageFile")
		if err != nil {
			panic(err)
		}

		if err := ctx.SaveUploadedFile(header, fmt.Sprintf("static/%s", header.Filename)); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(common.Image{
			Url: fmt.Sprintf("http://localhost:8080/static/%s", header.Filename),
		}))
	}
}
