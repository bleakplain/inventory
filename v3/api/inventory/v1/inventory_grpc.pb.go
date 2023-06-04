package v1

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"google.golang.org/grpc"

	"github.com/yourusername/inventory-service/internal/service"
)

type inventoryGRPCServer struct {
	invSvc service.InventoryService
}

func (s *inventoryGRPCServer) GetInventory(ctx context.Context, req *GetInventoryRequest) (*GetInventoryResponse, error) {
	return s.invSvc.GetInventory(ctx, req)
}

func (s *inventoryGRPCServer) UpdateInventory(ctx context.Context, req *UpdateInventoryRequest) (*UpdateInventoryResponse, error) {
	return s.invSvc.UpdateInventory(ctx, req)
}

func RegisterInventoryGRPCServer(s *grpc.Server, invSvc service.InventoryService) {
	server := &inventoryGRPCServer{
		invSvc: invSvc,
	}
	RegisterInventoryServiceServer(s, server)
}

func init() {
	middleware.RegisterUnaryServerInterceptor("grpc", middleware.ChainUnaryServer(
		middleware.UnaryServerInterceptor,
	))
}
