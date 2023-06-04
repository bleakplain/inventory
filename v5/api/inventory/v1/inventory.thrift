namespace go inventoryv1

service InventoryService {
  Inventory CreateInventory(1: CreateInventoryRequest request),
  Inventory UpdateInventory(1: UpdateInventoryRequest request),
  void DeleteInventory(1: DeleteInventoryRequest request),
  Inventory GetInventory(1: GetInventoryRequest request),
  ListInventoriesResponse ListInventories(1: ListInventoriesRequest request),
}

struct Inventory {
  1: i64 id,
  2: string sku,
  3: i64 warehouse_id,
  4: string channel,
  5: i32 quantity,
}

struct CreateInventoryRequest {
  1: string sku,
  2: i64 warehouse_id,
  3: string channel,
  4: i32 quantity,
}

struct UpdateInventoryRequest {
  1: i64 id,
  2: string sku,
  3: i64 warehouse_id,
  4: string channel,
  5: i32 quantity,
}

struct DeleteInventoryRequest {
  1: i64 id,
}

struct GetInventoryRequest {
  1: i64 id,
}

struct ListInventoriesRequest {
  1: string sku,
  2: i64 warehouse_id,
  3: string channel,
}

struct ListInventoriesResponse {
  1: list<Inventory> inventories,
}
