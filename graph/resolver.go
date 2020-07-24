package graph

import (
	"context"

	"github.com/AlekseyPromet/graph/model"
)

type Resolver struct{}

func (r *Resolver) CreateVideo(ctx context.Context, input NewVideo) (model.Video, error) {
	newVideo := model.Video{}
}
