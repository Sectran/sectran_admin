package file

import (
	"context"
	"net/http"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewDownloadFileLogic(r *http.Request, ctx context.Context, svcCtx *svc.ServiceContext) *DownloadFileLogic {
	return &DownloadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *DownloadFileLogic) DownloadFile(req *types.UUIDPathReq) (filePath string, err error) {
	path := l.r.Header.Get("relativePath")
	if path == "" {
		return "", errorx.NewCodeInvalidArgumentError("relativePath.missing")
	}

	return path, nil
}
