package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/yourusername/inventory-service/api/inventory/v1"
	"github.com/yourusername/inventory-service/internal/conf"
	"github.com/yourusername/inventory-service/internal/service"
)

type grpcServer struct {
	v1.UnimplementedInventoryServer
	*service.InventoryService
}

func newGRPCServer(svc *service.InventoryService, logger log.Logger) *grpc.Server {
	grpcSrv := grpc.NewServer(
		grpc.Address(conf.Conf.GRPC.Addr),
		grpc.Logger(logger),
	)
	v1.RegisterInventoryServer(grpcSrv, &grpcServer{InventoryService: svc})
	return grpcSrv
}

func (s *grpcServer) GetInventory(ctx context.Context, req *v1.GetInventoryRequest) (*v1.GetInventoryReply, error) {
	inventory, err := s.InventoryService.GetInventory(ctx, req.Sku, req.WarehouseID, req.Channel)
	if err != nil {
		return nil, err
	}
	return &v1.GetInventoryReply{
		Sku:         inventory.Sku,
		WarehouseID: inventory.WarehouseID,
		Channel:     inventory.Channel,
		Quantity:    inventory.Quantity,
	}, nil
}

func (s *grpcServer) UpdateInventory(ctx context.Context, req *v1.UpdateInventoryRequest) (*v1.UpdateInventoryReply, error) {
	err := s.InventoryService.UpdateInventory(ctx, req.Sku, req.WarehouseID, req.Channel, req.Quantity)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateInventoryReply{
		Success: true,
	}, nil
}
