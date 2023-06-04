namespace go inventoryv1

struct GetInventoryRequest {
  1: string sku,
  2: i64 warehouse_id,
  3: string channel,
}

struct UpdateInventoryRequest {
  1: string sku,
  2: i64 warehouse_id,
  3: string channel,
  4: i32 quantity,
}

struct Inventory {
  1: string sku,
  2: i64 warehouse_id,
  3: string channel,
  4: i32 quantity,
}

struct ListInventoriesResponse {
  1: list<Inventory> inventories,
}

service InventoryService {
  Inventory GetInventory(1: GetInventoryRequest request),
  Inventory UpdateInventory(1: UpdateInventoryRequest request),
  ListInventoriesResponse ListInventories(),
}
