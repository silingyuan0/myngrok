package client

import (
	"fmt"
	"ngrok/client/mvc"
	//	"ngrok/client/views/term"
	//	"ngrok/client/views/web"
	"ngrok/log"
	//	"ngrok/proto"
	"ngrok/util"
	"sync"
)

type command interface{}

type cmdQuit struct {
	// 退出时显示的信息
	message string
}

type cmdPlayRequest struct {
	// the tunnel to play this request over
	tunnel mvc.Tunnel

	// the bytes of the request to issue
	payload []byte
}

// The MVC Controller
type Controller struct {
	// 控制器log输出
	log.Logger

	// 模型通过广播信道发送更新
	updates *util.Broadcast

	// 模型接口
	model mvc.Model

	// 视图
	views []mvc.View

	//结构体内部发送命令到控制器
	cmds chan command

	// 结构体内部同步对象的访问状态
	state chan mvc.State

	//配置文件
	config *Configuration
}

// public interface
func NewController() *Controller {
	ctl := &Controller{
		Logger:  log.NewPrefixLogger("controller"),
		updates: util.NewBroadcast(),
		cmds:    make(chan command),
		views:   make([]mvc.View, 0),
		state:   make(chan mvc.State),
	}

	return ctl
}

func (ctl *Controller) State() mvc.State {
	return <-ctl.state
}

func (ctl *Controller) Update(state mvc.State) {
	ctl.updates.In() <- state
}

func (ctl *Controller) Updates() *util.Broadcast {
	return ctl.updates
}

func (ctl *Controller) Shutdown(message string) {
	ctl.cmds <- cmdQuit{message: message}
}

func (ctl *Controller) PlayRequest(tunnel mvc.Tunnel, payload []byte) {
	ctl.cmds <- cmdPlayRequest{tunnel: tunnel, payload: payload}
}

func (ctl *Controller) Go(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				err := util.MakePanicTrace(r)
				ctl.Error(err)
				ctl.Shutdown(err)
			}
		}()

		fn()
	}()
}

// 私有函数
func (ctl *Controller) doShutdown() {
	ctl.Info("Shutting down")

	var wg sync.WaitGroup

	// wait for all of the views, plus the model
	wg.Add(len(ctl.views) + 1)

	for _, v := range ctl.views {
		vClosure := v
		ctl.Go(func() {
			vClosure.Shutdown()
			wg.Done()
		})
	}

	ctl.Go(func() {
		ctl.model.Shutdown()
		wg.Done()
	})

	wg.Wait()
}

func (ctl *Controller) AddView(v mvc.View) {
	ctl.views = append(ctl.views, v)
}

func (ctl *Controller) GetWebInspectAddr() string {
	return ctl.config.InspectAddr
}

func (ctl *Controller) SetupModel(config *Configuration) *ClientModel {
	model := newClientModel(config, ctl)
	ctl.model = model
	return model
}

func (ctl *Controller) GetModel() *ClientModel {
	return ctl.model.(*ClientModel)
}

func (ctl *Controller) Run(config *Configuration) {
	//保存配置文件
	ctl.config = config
	//定义客户端模型结构体
	var model *ClientModel

	if ctl.model == nil {
		model = ctl.SetupModel(config) //根据配置文件建立模型
	} else {
		model = ctl.model.(*ClientModel)
	}

	// init the model
	var state mvc.State = model

	// init web ui
	//	var webView *web.WebView
	//	if config.InspectAddr != "disabled" {
	//	webView = web.NewWebView(ctl, config.InspectAddr)
	//		ctl.AddView(webView)
	//	}

	// init term ui
	//	var termView *term.TermView
	//	if config.LogTo != "stdout" {
	//		termView = term.NewTermView(ctl)
	//		ctl.AddView(termView)
	//	}

	//	for _, protocol := range model.GetProtocols() {
	//		switch p := protocol.(type) {
	//		case *proto.Http:
	//			if termView != nil {
	//				ctl.AddView(termView.NewHttpView(p))
	//			}

	//			if webView != nil {
	//				ctl.AddView(webView.NewHttpView(p))
	//			}
	//		default:
	//		}
	//	}

	ctl.Go(func() { autoUpdate(state, config.AuthToken) })
	ctl.Go(ctl.model.Run)

	updates := ctl.updates.Reg()
	defer ctl.updates.UnReg(updates)

	done := make(chan int)
	for {
		select {
		case obj := <-ctl.cmds:
			switch cmd := obj.(type) {
			case cmdQuit:
				msg := cmd.message
				go func() {
					ctl.doShutdown()
					fmt.Println(msg)
					done <- 1
				}()

			case cmdPlayRequest:
				ctl.Go(func() { ctl.model.PlayRequest(cmd.tunnel, cmd.payload) })
			}

		case obj := <-updates:
			state = obj.(mvc.State)

		case ctl.state <- state:
		case <-done:
			return
		}
	}
}
