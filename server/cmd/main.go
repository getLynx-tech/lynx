package main

import (
	"github.com/getLynx-tech/lynx/internal/presentation/http"
	"go.uber.org/fx"
)

// @title Lynx API
// @version 1.0
func main() {
	fx.New(
		http.Module,
	).Run()
}
