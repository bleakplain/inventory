package data

type Inventory struct {
	ID          uint64 `gorm:"primary_key"`
	Sku         string
	WarehouseID uint64
	Channel     string
	Quantity    int
}


func (d *Data) UpdateInventory(inventory *Inventory) error {
	return d.db.Save(inventory).Error
}

func (d *Data) GetInventory(id uint64) (*Inventory, error) {
	var inventory Inventory
	err := d.db.First(&inventory, id).Error
	return &inventory, err
}

func (d *Data) ListInventories(sku string, warehouseID uint64, channel string) ([]*Inventory, error) {
	var inventories []*Inventory
	query := d.db.Where("sku = ? AND warehouse_id = ? AND channel = ?", sku, warehouseID, channel)
	err := query.Find(&inventories).Error
	return inventories, err
}
