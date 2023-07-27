package uploadprovider

import (
	"context"
	"food-delivery-200lab/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
