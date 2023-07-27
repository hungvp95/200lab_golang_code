package ginupload

import (
	"fmt"
	"food-delivery-200lab/common"
	"food-delivery-200lab/component"
	"food-delivery-200lab/module/upload/uploadbusiness"
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

func UploadImageFile(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("imageFile")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := ctx.DefaultPostForm("folder", "images")

		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		business := uploadbusiness.NewUploadBusiness(appCtx.UploadProvider(), nil)
		img, err := business.Upload(ctx.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
