package grpc

import (
	"context"

	"github.com/yourusername/inventory-service/api/inventory/v1"
	"github.com/yourusername/inventory-service/internal/application"
	"google.golang.org/grpc"
)

type InventoryServer struct {
	v1.UnimplementedInventoryServiceServer
	appSvc *application.InventoryService
}

func NewInventoryServer(appSvc *application.InventoryService) *InventoryServer {
	return &InventoryServer{
		appSvc: appSvc,
	}
}

func (s *InventoryServer) GetInventory(ctx context.Context, req *v1.GetInventoryRequest) (*v1.Inventory, error) {
	return s.appSvc.GetInventory(ctx, req)
}

func (s *InventoryServer) UpdateInventory(ctx context.Context, req *v1.UpdateInventoryRequest) (*v1.Inventory, error) {
	return s.appSvc.UpdateInventory(ctx, req)
}

func (s *InventoryServer) ListInventories(ctx context.Context, _ *v1.Empty) (*v1.ListInventoriesResponse, error) {
	return s.appSvc.ListInventories(ctx)
}

func RegisterInventoryServer(s *grpc.Server, appSvc *application.InventoryService) {
	v1.RegisterInventoryServiceServer(s, NewInventoryServer(appSvc))
}
