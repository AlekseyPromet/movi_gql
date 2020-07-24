package graph

import (
	"context"

	"github.com/prometalex/graph/model"
)

type Resolver struct{}

func (r *Resolver) CreateVideo(ctx context.Context, input NewVideo) (model.Video, error) {
	newVideo := model.Video{}
}
