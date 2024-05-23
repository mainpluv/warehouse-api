package model

type Warehouse struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    IsAvailable bool   `json:"is_available"`
}

type Product struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Size        string `json:"size"`
    Code        string `json:"code"`
    Quantity    int    `json:"quantity"`
    WarehouseID int    `json:"warehouse_id"`
}
