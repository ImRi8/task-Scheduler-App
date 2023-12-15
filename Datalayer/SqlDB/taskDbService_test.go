package SqlDB

import (
	"Task-scheduler-App/Datalayer/Entity"
	"context"
	"gofr.dev/pkg/gofr"
	"testing"
)

func intializeCtx(t *testing.T) *gofr.Context {
	app := gofr.New()
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	return ctx
}

func TestTaskDbService_GetEntityById(t *testing.T) {
	ctx := intializeCtx(t)
	taskDbService := TaskDbService{}
	test := []struct {
		id       int64
		response interface{}
	}{
		{1, nil},
		{2, nil},
		{3, Entity.Task{
			ID: 1,
		}}}

	for _, tc := range test {
		resp, err := taskDbService.GetEntityById(tc.id, ctx)
		if err != nil {
			ctx.Logger.Error("Error Occured")
		}
		if resp == tc.response {
			ctx.Logger.Error("Error Occured")
		}
	}

}
