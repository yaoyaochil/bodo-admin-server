package request

import (
	"github.com/wangrui19970405/wu-shi-admin/server/model/common/request"
	"github.com/wangrui19970405/wu-shi-admin/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
