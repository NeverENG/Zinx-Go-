package ziface

type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(msg uint32, router IRouter)
}
