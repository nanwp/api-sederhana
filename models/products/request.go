package products

type ProductCreate struct {
	SKU        string `json:"sku" binding:"required"`
	Name       string `json:"nama_produk" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	Image      string `json:"image" binding:"required"`
	CategoryId int    `json:"kategori_id" binding:"required"`
}

type ProductUpdate struct {
	SKU        string `json:"sku"`
	Name       string `json:"nama_produk"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
	Image      string `json:"image"`
	CategoryId int    `json:"kategori_id"`
}
