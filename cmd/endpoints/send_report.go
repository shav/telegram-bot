package cmd_endpoints

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/shav/telegram-bot/internal/modules/core/transport/grpc/api/generated"
	"github.com/shav/telegram-bot/internal/modules/core/transport/grpc/server"
	"github.com/shav/telegram-bot/internal/observability/logger"
)

// messageSender позволяет отправлять сообщения в какой-либо мессенджер.
type messageSender interface {
	// SendMessage отправляет пользователю userId сообщение с текстом text.
	SendMessage(ctx context.Context, userId int64, text string) error
}

// Конечная точка для отправки пользователям отчётов.
type SendReportEndpoint struct {
	// Отправитель сообщений.
	messageSender messageSender
}

// NewSendReport создаёт конечную точку для отправки пользователям отчётов.
func NewSendReport(messageSender messageSender) *SendReportEndpoint {
	return &SendReportEndpoint{
		messageSender: messageSender,
	}
}

// Listen слушает на портах httpPort и grpcPort запросы к конечной точке для отправки пользователям отчётов.
func (e *SendReportEndpoint) Listen(ctx context.Context, grpcPort int, httpPort int) error {
	logger.Info(ctx, fmt.Sprintf("Start listening grpc on port %d...", grpcPort))
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Start listening grpc on port %d failed", grpcPort))
	}

	server := grpc.NewServer()
	sendReportServer, err := core_transport_grpc.NewSendReportServer(e.messageSender)
	if err != nil {
		return errors.Wrap(err, "Create SendReportServer failed")
	}
	api.RegisterReportSenderServer(server, sendReportServer)

	rmux := runtime.NewServeMux()
	mux := http.NewServeMux()
	mux.Handle("/", rmux)
	{
		sendReportServer, err := core_transport_grpc.NewSendReportServer(e.messageSender)
		if err != nil {
			return errors.Wrap(err, "Create SendReportServer failed")
		}
		logger.Info(ctx, "Register report send grpc handler")
		err = api.RegisterReportSenderHandlerServer(ctx, rmux, sendReportServer)
		if err != nil {
			return errors.Wrap(err, "Register send report grpc handler failed")
		}
	}

	var runGrpc = func() {
		reflection.Register(server)
		if err := server.Serve(grpcListener); err != nil {
			logger.Error(ctx, "Serve send report handler via grpc failed", logger.Fields.Error(err))
		}
	}
	go runGrpc()

	logger.Info(ctx, fmt.Sprintf("Start listening http on port %d...", httpPort))
	httpListener, err := net.Listen("tcp", fmt.Sprintf(":%d", httpPort))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Start listening http on port %d failed", httpPort))
	}
	err = http.Serve(httpListener, mux)
	if err != nil {
		return errors.Wrap(err, "Serve send report handler via http failed")
	}

	return nil
}
