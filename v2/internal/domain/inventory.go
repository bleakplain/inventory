package domain

type Inventory struct {
	SKU         string
	WarehouseID int64
	Channel     string
	Quantity    int32
}

type InventoryRepository interface {
	GetInventory(sku string, warehouseID int64, channel string) (*Inventory, error)
	UpdateInventory(sku string, warehouseID int64, channel string, quantity int32) (*Inventory, error)
	ListInventories() ([]*Inventory, error)
}

type InventoryService interface {
	GetInventory(sku string, warehouseID int64, channel string) (*Inventory, error)
	UpdateInventory(sku string, warehouseID int64, channel string, quantity int32) (*Inventory, error)
	ListInventories() ([]*Inventory, error)
}
