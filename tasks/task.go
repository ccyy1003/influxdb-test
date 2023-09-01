package tasks

import (
	"fmt"
	"influxdb-test/common"
	"strings"

	"github.com/influxdata/influxdb/client/v2"
)

type TaskInterface interface {
	Start() common.TestRes
	Prepare()
	DbName() string
}

type Task struct {
	TaskName string
	CmdList  map[string]string
}

// prepare test cmd list for some task
func (t *Task) Prepare() {}

func (t *Task) DbName() string { return "mydb" }

func (t *Task) Start() common.TestRes {
	var clnt = common.HttpClnt.Client

	var tr common.TestRes
	tr.TaskName = t.TaskName
	tr.TotalCnt = len(t.CmdList)

	var q client.Query

	tr.PassCnt = 0
	for syntax, cmd := range t.CmdList {
		q = client.NewQuery(cmd, t.DbName(), "ns")

		if response, err := clnt.Query(q); err == nil && response.Error() == nil {
			tr.SupportedSyntax = append(tr.SupportedSyntax, fmt.Sprintf("[ %s ] %s ", syntax, cmd))
			tr.PassCnt++
		} else {
			// record errinfo
			format_cmd := strings.ReplaceAll(cmd, "\\", "")
			info := fmt.Sprintf("[ %s ] %s :", syntax, format_cmd)
			if err != nil {
				info += " " + err.Error() + " "
			}
			if response.Error() != nil {
				info += " " + response.Error().Error() + " "
			}

			tr.ErrInfos = append(tr.ErrInfos, info)
		}

	}

	return tr
}
