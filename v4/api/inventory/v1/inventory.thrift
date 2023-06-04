namespace go inventoryv1

service Inventory {
  GetInventoryResponse GetInventory(1:GetInventoryRequest request),
  UpdateInventoryResponse UpdateInventory(1:UpdateInventoryRequest request),
}

struct GetInventoryRequest {
  1: string sku,
  2: i64 warehouse_id,
  3: string channel,
}

struct GetInventoryResponse {
  1: Inventory inventory,
}

struct UpdateInventoryRequest {
  1: string sku,
  2: i64 warehouse_id,
  3: string channel,
  4: i32 new_quantity,
}

struct UpdateInventoryResponse {
  1: bool success,
}

struct Inventory {
  1: string sku,
  2: i64 warehouse_id,
  3: string channel,
  4: i32 quantity,
}
