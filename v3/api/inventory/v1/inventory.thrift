namespace go inventoryv1

service Inventory {
  InventoryResponse GetInventory(1:GetInventoryRequest request),
  InventoryResponse UpdateInventory(1:UpdateInventoryRequest request),
}

struct GetInventoryRequest {
  1: string sku,
}

struct UpdateInventoryRequest {
  1: string sku,
  2: i32 warehouse_id,
  3: string channel,
  4: i32 quantity,
}

struct InventoryResponse {
  1: string sku,
  2: i32 warehouse_id,
  3: string channel,
  4: i32 quantity,
}
