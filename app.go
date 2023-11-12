package main

import (
	"changeme/internal/define"
	"changeme/internal/service"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// ConnectionList
func (a *App) ConnectionList() H {
	conn, err := service.ConnectionList()
	if err != nil {
		return M{
			"code": 500,
			"msg":  err.Error(),
		}
	}

	return M{
		"code": 200,
		"msg":  "success",
		"data": conn,
	}
}

// CreateConnection
func (a *App) CreateConnection(conn *define.Connection) H {
	err := service.CreateConnection(conn)
	if err != nil {
		return M{
			"code": 500,
			"msg":  err.Error(),
		}
	}

	return M{
		"code": 200,
		"msg":  "create success",
	}

}
