//go:generate minimock -i messageSender -o ./mocks/ -s ".go"

package core_transport_grpc

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/shav/telegram-bot/internal/modules/core/transport/grpc/api/generated"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// messageSender позволяет отправлять сообщения в какой-либо мессенджер.
type messageSender interface {
	// SendMessage отправляет пользователю userId сообщение с текстом text.
	SendMessage(ctx context.Context, userId int64, text string) error
}

// SendReportServer принимает запросы на отправку отчётов пользователям.
type SendReportServer struct {
	api.UnimplementedReportSenderServer
	// Отправитель сообщений пользователю.
	messageSender messageSender
}

// NewSendReportServer создаёт сервер для обработки запросов на отправку отчётов пользователям.
func NewSendReportServer(messageSender messageSender) (*SendReportServer, error) {
	if messageSender == nil {
		return nil, errors.New("New SendReportServer: message sender is not assigned")
	}

	return &SendReportServer{
		messageSender: messageSender,
	}, nil
}

// SendReport отправляет отчёт пользователю.
func (s *SendReportServer) SendReport(ctx context.Context, request *api.SendReportRequest) (*emptypb.Empty, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "SendReportServer.SendReport")
	defer span.Finish()

	report := fmt.Sprintf("%s\n%s", request.Report.Title, request.Report.Content)
	err := s.messageSender.SendMessage(ctx, request.UserId, report)

	if err != nil {
		tracing.SetError(span)
	}
	return &emptypb.Empty{}, err
}
