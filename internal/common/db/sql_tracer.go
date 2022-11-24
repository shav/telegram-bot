package db

import (
	"context"
	"database/sql/driver"
	"strings"

	"github.com/ngrok/sqlmw"
)

// SqlTracer собирает трейсы sql-запросов.
type SqlTracer struct {
	sqlmw.NullInterceptor
	// Трейс db-запросв.
	trace *strings.Builder
	// Порядковый номер последнего залогированного запроса.
	queryIndex int
}

// NewSqlTracer создаёт новый трейсер sql-запросов.
func NewSqlTracer() *SqlTracer {
	return &SqlTracer{
		trace: &strings.Builder{},
	}
}

// GetTrace возвращает текущий трейс ранее выполненных sql-запросов.
func (s *SqlTracer) GetTrace() string {
	return s.trace.String()
}

// traceQuery записывает sql-запрос в трейс.
func (s *SqlTracer) traceQuery(query string) {
	s.queryIndex++
	if s.queryIndex > 1 {
		s.trace.WriteString("\n")
	}
	s.trace.WriteString(query)
}

// #region Interceptor

func (s *SqlTracer) ConnExecContext(ctx context.Context, conn driver.ExecerContext, query string, args []driver.NamedValue) (driver.Result, error) {
	s.traceQuery(query)
	return s.NullInterceptor.ConnExecContext(ctx, conn, query, args)
}

func (s *SqlTracer) ConnQueryContext(ctx context.Context, conn driver.QueryerContext, query string, args []driver.NamedValue) (context.Context, driver.Rows, error) {
	s.traceQuery(query)
	return s.NullInterceptor.ConnQueryContext(ctx, conn, query, args)
}

func (s *SqlTracer) StmtExecContext(ctx context.Context, stmt driver.StmtExecContext, query string, args []driver.NamedValue) (driver.Result, error) {
	s.traceQuery(query)
	return s.NullInterceptor.StmtExecContext(ctx, stmt, query, args)
}

func (s *SqlTracer) StmtQueryContext(ctx context.Context, stmt driver.StmtQueryContext, query string, args []driver.NamedValue) (context.Context, driver.Rows, error) {
	s.traceQuery(query)
	return s.NullInterceptor.StmtQueryContext(ctx, stmt, query, args)
}

// Reset сбрасывает состояние трейсера и очищает ранее записанный трейс.
func (s *SqlTracer) Reset() {
	s.trace = &strings.Builder{}
	s.queryIndex = 0
}

// #endregion
