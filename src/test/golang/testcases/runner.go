package testcases

import (
	"sort"
	"strings"
	"time"

	"github.com/starter-go/application"
	"github.com/starter-go/vlog"
)

const (
	OrderBase = iota

	OrderAuthSignUp

	OrderAuthSendVerificationMail
	OrderAuthSendVerificationSMS

	OrderAuthSignIn
	OrderSessionsCurrent
)

// TestRunnableInfo ...
type TestRunnableInfo struct {
	application.Life

	Name  string
	Index int
}

// TestRunnable ...
type TestRunnable interface {
	GetRunnableInfo() *TestRunnableInfo
}

// TestRunnableResult ...
type TestRunnableResult struct {
	Runnable *TestRunnableInfo
	Error    error
	Skipped  bool
}

////////////////////////////////////////////////////////////////////////////////

// TestRunner ...
type TestRunner struct {

	//starter:component

	AppContext   application.Context //starter:inject("context")
	RunnableList []TestRunnable      //starter:inject(".")

	units []*TestRunnableInfo
}

// Life ...
func (inst *TestRunner) Life() *application.Life {
	return &application.Life{
		OnStartPost: inst.start,
	}
}

func (inst *TestRunner) start() error {
	go func() {
		time.Sleep(time.Second * 3)
		err := inst.run()
		if err != nil {
			vlog.Error(err.Error())
		}
	}()
	return nil
}

func (inst *TestRunner) loadUnits() []*TestRunnableInfo {
	src := inst.RunnableList
	dst := make([]*TestRunnableInfo, 0)
	for _, item1 := range src {
		if item1 == nil {
			continue
		}
		info := item1.GetRunnableInfo()
		if info == nil {
			continue
		}
		dst = append(dst, info)
	}
	return dst
}

func (inst *TestRunner) getUnits() []*TestRunnableInfo {
	list := inst.units
	if list == nil {
		list = inst.loadUnits()
		inst.units = list
		inst.sort()
	}
	return list
}

func (inst *TestRunner) run() error {

	list := inst.getUnits()
	results := make([]*TestRunnableResult, 0)
	bar := strings.Repeat("--------", 10)

	for idx, item := range list {
		res := &TestRunnableResult{
			Runnable: item,
		}
		item.Index = idx
		if inst.isItemEnabled(item) {
			vlog.Info(bar)
			err := inst.invokeItem(item)
			vlog.Info(bar)
			res.Error = err
		} else {
			res.Skipped = true
		}
		results = append(results, res)
	}

	// print results
	for _, res := range results {
		inst.printResult(res)
	}

	return nil
}

func (inst *TestRunner) printResult(res *TestRunnableResult) {

	idx := res.Runnable.Index
	name := res.Runnable.Name
	vlog.Info("testcase[%d]: %s", idx, name)

	err := res.Error
	if err != nil {
		vlog.Error(" ... [error]: %s", err.Error())
		return
	}

	if res.Skipped {
		vlog.Info(" ... [skipped]")
		return
	}

	vlog.Info(" ... [OK]")
}

func (inst *TestRunner) Swap(i1, i2 int) {
	list := inst.units
	list[i1], list[i2] = list[i2], list[i1]
}
func (inst *TestRunner) Less(i1, i2 int) bool {
	o1 := inst.units[i1]
	o2 := inst.units[i2]
	return o1.Order < o2.Order
}
func (inst *TestRunner) Len() int {
	return len(inst.units)
}
func (inst *TestRunner) sort() {
	sort.Sort(inst)
}

func (inst *TestRunner) isItemEnabled(item *TestRunnableInfo) bool {
	name := item.Name
	pName := "tester.enable." + name
	props := inst.AppContext.GetProperties()
	getter := props.Getter()
	b := getter.GetBool(pName)
	str := getter.GetString(pName)
	if str == "" {
		vlog.Warn("no property with name: " + pName)
	}
	return b
}

func (inst *TestRunner) invokeItem(item *TestRunnableInfo) error {

	vlog.Info("run testcase[%d]: %s", item.Index, item.Name)

	fnlist := make([]func() error, 0)
	fnlist = append(fnlist, item.OnCreate)
	fnlist = append(fnlist, item.OnStartPre)
	fnlist = append(fnlist, item.OnStart)
	fnlist = append(fnlist, item.OnStartPost)
	fnlist = append(fnlist, item.OnLoop)
	fnlist = append(fnlist, item.OnStopPre)
	fnlist = append(fnlist, item.OnStop)
	fnlist = append(fnlist, item.OnStopPost)
	fnlist = append(fnlist, item.OnDestroy)

	for _, step := range fnlist {
		if step == nil {
			continue
		}
		err := step()
		if err != nil {
			return err
		}
	}
	return nil
}
