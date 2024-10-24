package file

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sectran_admin/ent"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"strings"

	"github.com/duke-git/lancet/fileutil"
	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogic(r *http.Request, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadLogic) Upload() (resp *types.UploadResp, err error) {
	err = l.r.ParseMultipartForm(l.svcCtx.Config.UploadConf.MaxVideoSize)
	if err != nil {
		logx.Error("fail to parse the multipart form")
		return nil, errorx.NewCodeInvalidArgumentError("file.parseFormFailed")
	}

	file, handler, err := l.r.FormFile("file")
	if err != nil {
		logx.Error("the value of file cannot be found")
		return nil, errorx.NewCodeInvalidArgumentError("file.parseFormFailed")
	}
	defer file.Close()

	// judge if the file size is over max size
	// 判断文件大小是否超过设定值
	fileType := strings.Split(handler.Header.Get("Content-Type"), "/")[0]
	if fileType != "image" && fileType != "video" && fileType != "audio" {
		fileType = "other"
	}

	domain := l.ctx.Value("request_domain").((*ent.User))
	err = CheckOverSize(l.ctx, l.svcCtx, fileType, handler.Size)
	if err != nil {
		logx.Errorw("the file is over size", logx.Field("type", fileType),
			logx.Field("userId", domain.ID), logx.Field("size", handler.Size),
			logx.Field("fileName", handler.Filename))
		return nil, err
	}

	service := l.r.Header.Get("serviceName")
	if service == "" {
		return nil, errorx.NewCodeInvalidArgumentError("service.missing")
	}

	//服务名称/upload/用户名称
	uploadDir := filepath.Join(service, "upload", domain.Name)
	if !fileutil.IsExist(uploadDir) {
		err = fileutil.CreateDir(uploadDir + "/")
		if err != nil {
			logx.Errorw("failed to create directory for storing public files", logx.Field("path", uploadDir))
			return nil, errorx.NewCodeError(errorcode.Internal,
				l.svcCtx.Trans.Trans(l.ctx, i18n.Failed))
		}
	}

	filePath := filepath.Join(uploadDir, handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		logx.Error("fail to create the file on server")
		return nil, errorx.NewCodeInternalError("file.createFailed")
	}
	defer dst.Close()

	// 将上传的文件内容保存到目标文件
	_, err = io.Copy(dst, file)
	if err != nil {
		logx.Error("fail to save the file to server")
		return nil, errorx.NewCodeInternalError("file.saveFailed")
	}

	return &types.UploadResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)},
		Data:         types.UploadInfo{Name: handler.Filename, Url: ""},
	}, nil
}
