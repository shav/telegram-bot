package core_transport_grpc

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/shav/telegram-bot/internal/modules/core/domain/models"
	"github.com/shav/telegram-bot/internal/modules/core/transport/grpc/api/generated"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// ReportSendClient представлет из себя клиента для отправки отчётов пользователю.
type ReportSendClient struct {
	// grpc-pеализация клиента.
	client api.ReportSenderClient
}

// NewReportSendClient создаёт клиента для отправки отчётов пользователю.
func NewReportSendClient(conn grpc.ClientConnInterface) (*ReportSendClient, error) {
	if conn == nil {
		return nil, errors.New("New NewReportSendClient: connection is not assigned")
	}

	return &ReportSendClient{
		client: api.NewReportSenderClient(conn),
	}, nil
}

// SendReport отправляет пользователю userId отчёт report.
func (c *ReportSendClient) SendReport(ctx context.Context, userId int64, report core_models.Report) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "ReportSendClient.SendReport")
	defer span.Finish()

	request := &api.SendReportRequest{
		UserId: userId,
		Report: &api.Report{
			Title:   report.Title,
			Content: report.Content,
		},
	}
	_, err := c.client.SendReport(ctx, request)

	if err != nil {
		tracing.SetError(span)
	}
	return err
}
