package infrastructure

import (
	"github.com/google/wire"
	"github.com/yourusername/yourproject/internal/domain"
)

func InitializeInventoryService(repo domain.InventoryRepository) domain.InventoryService {
	return domain.NewInventoryService(repo)
}

var InventorySet = wire.NewSet(
	domain.NewInMemoryInventoryRepository,
	InitializeInventoryService,
)
