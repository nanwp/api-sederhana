package products

type ProductCreate struct {
	SKU        string `json:"sku" binding:"required"`
	Name       string `json:"nama_produk" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	Image      string `json:"image" binding:"required"`
	CategoryId int    `json:"kategori_id" binding:"required"`
}
