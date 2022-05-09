package comment

import (
	"context"
	"golang-projects-a/pkg/core/adapter/commentadapter"
	"golang-projects-a/pkg/core/service"
)

type DeleteManyReq struct {
	OrganizationId uint64
}

func (s Service) DeleteMany(ctx context.Context, req DeleteManyReq) (serviceErr service.Error) {
	err := s.CommentRepo.DeleteMany(ctx, commentadapter.RepoFilter{
		OrganizationId: req.OrganizationId,
	})
	if err != nil {
		return service.ErrDatasourceAccess("delete list comment query error")
	}

	return nil
}
