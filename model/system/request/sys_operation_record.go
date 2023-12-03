package request

import (
	"github.com/Sectran/sectran_admin/model/common/request"
	"github.com/Sectran/sectran_admin/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
