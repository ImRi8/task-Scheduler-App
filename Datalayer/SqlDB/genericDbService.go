package SqlDB

import "gofr.dev/pkg/gofr"

type GenericDbService interface {
	ShadowRowInEntity(Id int64, ctx *gofr.Context) error
	UnShadowRowInEntity(Id int64, ctx *gofr.Context) error
	CreateRowInEntity(ctx *gofr.Context, obj interface{}) (interface{}, error)
	UpdateRowInEntity(ctx *gofr.Context, obj interface{}) (interface{}, error)
	GetEntityById(Id int64, ctx *gofr.Context) (interface{}, error)
}
