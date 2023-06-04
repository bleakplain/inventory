package server

import (
	"context"
	"net"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/yourusername/inventory-service/api/inventory/v1"
	"github.com/yourusername/inventory-service/internal/conf"
	"github.com/yourusername/inventory-service/internal/service"
)

type ThriftServer struct {
	s *thrift.TSimpleServer
}

func NewThriftServer(c *conf.ServerConfig, inventoryService *service.InventoryService, logger log.Logger) *ThriftServer {
	processor := v1.NewInventoryServiceProcessor(inventoryService)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()

	addr := c.Thrift.Addr
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Log(log.LevelError, "msg", "failed to listen thrift address", "addr", addr, "error", err)
		return nil
	}

	transport := thrift.NewTFramedTransportFactory(transportFactory)
	serverTransport := thrift.NewTServerSocketFromListenerTimeout(listener, 0)
	s := thrift.NewTSimpleServer4(processor, serverTransport, transport, protocolFactory)

	return &ThriftServer{
		s: s,
	}
}

func (s *ThriftServer) Start(ctx context.Context) error {
	return s.s.Serve()
}

func (s *ThriftServer) Stop(ctx context.Context) error {
	s.s.Stop()
	return nil
}
