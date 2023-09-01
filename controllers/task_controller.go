package controllers

import (
	"influxdb-test/common"
	"influxdb-test/data"
	"influxdb-test/tasks"
	"strings"
)

type TaskController struct {
	BaseAPIView
}

func DoTask(t tasks.TaskInterface) common.TestRes {
	data.Init()
	t.Prepare()
	return t.Start()
}

// GetTaskList ...
func (c *TaskController) CQTask() {

	if err := c.initRequest(); err != nil {
		c.RespError(401, err.Error())
		return
	}
	//c.Ctx.Output.Body([]byte("hello world"))
	//return

	c.RespOK(DoTask(&tasks.CQTask{}))
}

func (c *TaskController) FuncTask() {
	if err := c.initRequest(); err != nil {
		c.RespError(401, err.Error())
		return
	}

	q := tasks.FuncTask{}
	c.RespOK(DoTask(&q))
}

func (c *TaskController) HintTask() {
	if err := c.initRequest(); err != nil {
		c.RespError(401, err.Error())
		return
	}

	q := tasks.HintTask{}
	c.RespOK(DoTask(&q))
}

func (c *TaskController) InterfaceTask() {
	if err := c.initRequest(); err != nil {
		c.RespError(401, err.Error())
		return
	}

	q := tasks.InterfaceTask{}
	c.RespOK(DoTask(&q))
}

func (c *TaskController) MathOptTask() {
	if err := c.initRequest(); err != nil {
		c.RespError(401, err.Error())
		return
	}

	q := tasks.MathOptTask{}
	c.RespOK(DoTask(&q))
}

func (c *TaskController) MgDbTask() {
	if err := c.initRequest(); err != nil {
		c.RespError(401, err.Error())
		return
	}

	q := tasks.MgDbTask{}
	c.RespOK(DoTask(&q))
}

func (c *TaskController) QueryTask() {
	if err := c.initRequest(); err != nil {
		c.RespError(401, err.Error())
		return
	}

	q := tasks.QueryTask{}
	c.RespOK(DoTask(&q))
}

func (c *TaskController) ShowTask() {
	if err := c.initRequest(); err != nil {
		c.RespError(401, err.Error())
		return
	}

	q := tasks.ShowTask{}
	c.RespOK(DoTask(&q))
}

func (c *TaskController) AllTask() {
	if err := c.initRequest(); err != nil {
		c.RespError(401, err.Error())
		return
	}
	var trs []common.TestRes
	trs = append(trs, DoTask(&tasks.CQTask{}))
	trs = append(trs, DoTask(&tasks.FuncTask{}))
	trs = append(trs, DoTask(&tasks.HintTask{}))
	trs = append(trs, DoTask(&tasks.InterfaceTask{}))
	trs = append(trs, DoTask(&tasks.MathOptTask{}))
	trs = append(trs, DoTask(&tasks.MgDbTask{}))
	trs = append(trs, DoTask(&tasks.QueryTask{}))
	trs = append(trs, DoTask(&tasks.ShowTask{}))

	for i := 0; i < len(trs); i++ {
		for j := 0; j < len(trs[i].SupportedSyntax); j++ {
			id := strings.Index(trs[i].SupportedSyntax[j], "]")
			if id != -1 {
				trs[i].SupportedSyntax[j] = trs[i].SupportedSyntax[j][:id+1]
			}
		}
	}
	c.RespOK(trs)
}
