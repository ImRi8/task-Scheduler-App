package SqlDB

import "gofr.dev/pkg/gofr"

type GenericDbService interface {
	shadowRowInEntity(Id int64, ctx *gofr.Context) error
	unShadowRowInEntity(Id int64, ctx *gofr.Context) error
	createRowInEntity(ctx *gofr.Context, obj interface{}) (interface{}, error)
	updateRowInEntity(ctx *gofr.Context, obj interface{}) (interface{}, error)
	getEntityById(Id int64, ctx *gofr.Context) (interface{}, error)
}
