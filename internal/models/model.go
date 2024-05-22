package models

type Warehouse struct {
	ID          int
	Name        string
	IsAvailable bool
}

type Item struct {
	ID   int
	Name string
	Size string
	Code string
}

type Inventory struct {
	WarehouseID int
	ItemID      int
	Quantity    int
}

type ReserveRequest struct {
	Codes       []string `json:"codes"`
	WarehouseID int      `json:"warehouse_id"`
}

type ReleaseRequest struct {
	Codes       []string `json:"codes"`
	WarehouseID int      `json:"warehouse_id"`
}

type RemainingRequest struct {
	WarehouseID int `json:"warehouse_id"`
}
