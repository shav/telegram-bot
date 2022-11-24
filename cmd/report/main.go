package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/shav/telegram-bot/cmd"
	"github.com/shav/telegram-bot/cmd/endpoints"
	"github.com/shav/telegram-bot/cmd/report/settings"
	"github.com/shav/telegram-bot/internal/config"
	"github.com/shav/telegram-bot/internal/modules"
	"github.com/shav/telegram-bot/internal/modules/finances"
	"github.com/shav/telegram-bot/internal/observability/logger"
	"github.com/shav/telegram-bot/internal/observability/metrics"
	"github.com/shav/telegram-bot/internal/observability/tracing"
)

// Модули приложения.
var allModules = []modules.Module{finances.NewModule()}

// Тип конфига приложения.
var configType = flag.String("configType", string(config.File), fmt.Sprintf("config type: %s or %s", config.File, config.Env))

// Путь к конфиг-файлу приложения.
var configFile = flag.String("configPath", report_service.DefaultSettings.ConfigFile, "path to config file")

func main() {
	flag.Parse()

	cfg, err := config.New(config.Type(*configType), report_service.DefaultSettings.ServiceName, *configFile)
	if err != nil {
		log.Fatal("Config init failed: ", err)
	}

	zapLogger, err := logger.NewZapLogger(cfg.LogMode())
	if err != nil {
		log.Fatal("Logger init failed: ", err)
	}
	err = logger.Init(zapLogger, tracing.AddTraceIdToLog)
	if err != nil {
		log.Fatal("Logger init failed: ", err)
	}
	defer logger.Stop()

	ctx, contextCancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	err = tracing.Init(cfg.ServiceName(), cfg.TraceSampling())
	defer tracing.Stop()
	if err != nil {
		logger.Error(ctx, "Tracing init failed", logger.Fields.Error(err))
	}

	metrics.Init()
	go func() {
		metricsEndpoint := cmd_endpoints.NewMetrics()
		metricsPort := cfg.MetricsPort(report_service.DefaultSettings.MetricsPort)
		err := metricsEndpoint.Listen(metricsPort)
		if err != nil {
			logger.Error(ctx, "Metrics endpoint init failed", logger.Fields.Error(err))
		}
	}()

	initArgs := modules.ModuleInitArgs{
		AppName:     cmd.AppName,
		ServiceName: report_service.DefaultSettings.ServiceName,
		ConfigType:  config.Type(*configType),
		ConfigFile:  *configFile,
		Config:      cfg,
		Ctx:         ctx,
	}
	err = initModules(initArgs)
	defer stop(ctx, contextCancel)
	if err != nil {
		// Ошибка логируется в initModules
		return
	}

	<-ctx.Done()
}

// initModules инициализирует все модули приложения.
func initModules(args modules.ModuleInitArgs) error {
	for _, module := range allModules {
		moduleName := logger.Fields.String("moduleName", module.GetName())
		logger.Info(args.Ctx, "Module {moduleName} is initializing...", moduleName)
		err := module.Init(args)
		if err != nil {
			logger.Error(args.Ctx, "Module {moduleName} init failed", moduleName, logger.Fields.Error(err))
			return err
		}
		err = module.InitMessageQueueHandlers(args)
		if err != nil {
			logger.Error(args.Ctx, "Module {moduleName} init message queue handlers failed", moduleName, logger.Fields.Error(err))
			return err
		}
		logger.Info(args.Ctx, "Module {moduleName} has been initialized.", moduleName)
	}
	return nil
}

// stop завершает работу приложения.
func stop(ctx context.Context, contextCancel context.CancelFunc) {
	logger.Info(ctx, "Application is stopping...")
	contextCancel()

	for _, module := range allModules {
		moduleName := logger.Fields.String("moduleName", module.GetName())
		logger.Info(ctx, "Module {moduleName} is stopping...", moduleName)
		err := module.Stop()
		if err != nil {
			logger.Error(ctx, "Module {moduleName} stopping failed", moduleName, logger.Fields.Error(err))
		}
		logger.Info(ctx, "Module {moduleName} has been stopped", moduleName)
	}

	logger.Info(ctx, "Application has been stopped...")
}
