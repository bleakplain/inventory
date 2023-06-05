package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/bleakplain/inventory/internal/conf"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	// create a new tracer provider with Prometheus exporter
	exporter, err := prometheus.NewExporter(prometheus.Options{})
	if err != nil {
		log.Fatalf("failed to create Prometheus exporter: %v", err)
	}
	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))

	// set global tracer provider
	otel.SetTracerProvider(tp)

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.TracerProvider(tp),
	)
}

func main() {
	flag.Parse()

	// create a new logger
	logger := log.NewStdLogger(os.Stdout)
	log := log.NewHelper(logger)

	// open log file
	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer file.Close()

	// add file logger to multi logger
	multiLogger := log.NewMultiLogger()
	multiLogger.Add(logger)
	multiLogger.Add(log.NewStdLogger(file))

	// update logger to use multi logger
	log = log.NewHelper(multiLogger)

	c := config.New(
		config.WithSource(
			file.NewSource("config.yaml"),
		),
	)
	if err := c.Load(); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		log.Fatalf("error", err)
		os.Exit(1)
	}

	app, cleanup, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		log.Fatalf("initApp error", err)
		os.Exit(1)
	}
	defer cleanup()

	// start and wait for stop signal
	go func() {
		if err := app.Run(); err != nil {
			log.Fatalf("app run error", err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	if err := app.Stop(); err != nil {
		log.Fatalf("error", err)
		os.Exit(1)
	}
}
