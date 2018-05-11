package main

import "go-middleware-example/app"

type appServer struct {
	app.Server
}

func newAppServer(s app.Server) *appServer {
	return &appServer{s}
}

func main() {
	s := app.Server{}
	app := newAppServer(s)
	app.Routes()
	app.Run()
}
