package uploadbusiness

import (
	"bytes"
	"context"
	"fmt"
	"food-delivery-200lab/common"
	"food-delivery-200lab/component/uploadprovider"
	"food-delivery-200lab/module/upload/uploadmodel"
	"image"
	"io"
	"log"
	"path/filepath"
	"time"
)

type CreateImageStorage interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBusiness struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStorage
}

func NewUploadBusiness(provider uploadprovider.UploadProvider, imgStore CreateImageStorage) *uploadBusiness {
	return &uploadBusiness{provider: provider, imgStore: imgStore}
}

func (business *uploadBusiness) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	width, height, err := getImageDimension(fileBytes)
	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := business.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, uploadmodel.ErrCanNotSaveFile(err)
	}

	img.Width = width
	img.Height = height
	img.Extension = fileExt
	//img.CloudName = "s3" // should be set in provider

	//if err := uploadbusiness.imgStore.CreateImage(ctx, img); err != nil {
	//	return nil, uploadmodel.ErrCanNotSaveFile(err)
	//}

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
