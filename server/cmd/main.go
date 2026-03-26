package main

import (
	"github.com/getLynx-tech/lynx/internal/application"
	"github.com/getLynx-tech/lynx/internal/conf"
	"github.com/getLynx-tech/lynx/internal/domain/repository"
	"github.com/getLynx-tech/lynx/internal/domain/service"
	"github.com/getLynx-tech/lynx/internal/infrastructure/postgres"
	"github.com/getLynx-tech/lynx/internal/presentation/http"
	"go.uber.org/fx"
)

// @title Lynx API
// @version 1.0
func main() {
	fx.New(
		conf.Module,
		postgres.Module,
		repository.Module,
		service.Module,
		application.Module,
		http.Module,
	).Run()
}
