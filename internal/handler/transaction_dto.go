package handler

type CheckoutItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

type CheckoutRequest struct {
	Items []CheckoutItemRequest `json:"items" binding:"required,dive,required"`
	Total float64               `json:"total" binding:"required"`
}

type ReportResponse struct {
	TotalRevenue   float64     `json:"total_revenue"`
	TotalTransaksi int64       `json:"total_transaksi"`
	ProdukTerlaris *TopProduct `json:"produk_terlaris,omitempty"`
}

type TopProduct struct {
	Nama       string `json:"nama"`
	QtyTerjual int64  `json:"qty_terjual"`
}
