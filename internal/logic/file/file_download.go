package file

import (
	"context"
	"path"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-file/ent/file"
	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadFileLogic {
	return &DownloadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadFileLogic) DownloadFile(req *types.UUIDPathReq) (filePath string, err error) {
	if err != nil {
		return "", dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	if file.Status == 1 {
		logx.Infow("public download", logx.Field("fileName", file.Name),
			logx.Field("userId", l.ctx.Value("userId").(string)),
			logx.Field("filePath", file.Path))
		return path.Join(l.svcCtx.Config.UploadConf.PublicStorePath, file.Path), nil
	} else {
		logx.Infow("private download", logx.Field("fileName", file.Name),
			logx.Field("userId", l.ctx.Value("userId").(string)),
			logx.Field("filePath", file.Path))
		return path.Join(l.svcCtx.Config.UploadConf.PrivateStorePath, file.Path), nil
	}
}
