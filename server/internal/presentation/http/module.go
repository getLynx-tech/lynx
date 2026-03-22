package http

import (
	"github.com/getLynx-tech/lynx/internal/presentation/http/handler"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"api",
	handler.Module,
	fx.Provide(NewServer),
	fx.Invoke(RegisterServer),
)
