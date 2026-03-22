package http

import (
	"github.com/getLynx-tech/lynx/internal/presentation/http/handler"
	"github.com/getLynx-tech/lynx/internal/presentation/http/middleware"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"api",
	middleware.Module,
	handler.Module,
	fx.Provide(NewServer),
	fx.Invoke(RegisterServer),
)
