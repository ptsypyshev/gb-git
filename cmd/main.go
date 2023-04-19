package main

import (
	"github.com/ptsypyshev/gbgit/internal/app"
)

func main() {
	server := app.NewApp()
	server.Run()
}
