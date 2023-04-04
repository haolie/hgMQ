package sys

import (
	"context"
	"fmt"
	"time"

	"github.com/haolie/goUtil/logUtil"
)

const (
	con_file_name = "sys"
)

type LoadFun func() []error

var (
	perLoadMap = make(map[string]LoadFun, 4)

	loafMap = make(map[string]LoadFun, 4)

	completeMap = make(map[string]LoadFun, 4)

	unLoadMap = make(map[string]LoadFun, 4)

	cancelFun func()
)

func handleErrs(errList []error) {
	for _, err := range errList {
		logUtil.ErrLog("handleErrs", err)
	}
}

func Start() (ctx context.Context, success bool) {
	if run(perLoadMap) == false {
		logUtil.ErrLog("start", fmt.Errorf("run perLoad fail"))
		return
	}

	if run(loafMap) == false {
		logUtil.ErrLog("start", fmt.Errorf("run load fail"))
		return
	}

	if run(completeMap) == false {
		logUtil.ErrLog("start", fmt.Errorf("run complete fail"))
		return
	}

	logUtil.InfoLog("系统初始化", "sys", "start", "sys run success")

	ctx, cancelFun = context.WithCancel(context.Background())
	time.AfterFunc(time.Second*5, func() {
		cancelFun()
	})

	return
}

func run(fnMap map[string]LoadFun) bool {
	var list []error
	for _, fn := range fnMap {
		errs := fn()
		if len(errs) > 0 {
			list = append(list, errs...)
		}
	}

	if len(list) > 0 {
		handleErrs(list)
		return false
	}

	return true
}

func RegisterPerLoad(key string, fun LoadFun) {
	_, exists := perLoadMap[key]
	if exists {
		panic(fmt.Errorf("%s key=%s 重复注册", "RegisterPerLoad", key))
	}

	perLoadMap[key] = fun
}

func RegisterLoad(key string, fun LoadFun) {
	_, exists := loafMap[key]
	if exists {
		panic(fmt.Errorf("%s key=%s 重复注册", "RegisterLoad", key))
	}

	loafMap[key] = fun
}

func RegisterComplete(key string, fun LoadFun) {
	_, exists := completeMap[key]
	if exists {
		panic(fmt.Errorf("%s key=%s 重复注册", "RegisterComplete", key))
	}

	completeMap[key] = fun
}

func RegisterUnLoad(key string, fun LoadFun) {
	_, exists := unLoadMap[key]
	if exists {
		panic(fmt.Errorf("%s key=%s 重复注册", "RegisterUnLoad", key))
	}

	unLoadMap[key] = fun
}
