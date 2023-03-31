package products

import "github.com/nanwp/api-sederhana/models/category"

type ProductResponse struct {
	ID       int                       `json:"id"`
	SKU      string                    `json:"sku"`
	Name     string                    `json:"nama_produk"`
	Stock    int                       `json:"stock"`
	Price    int                       `json:"price"`
	Image    string                    `json:"image"`
	Category category.CategoryResponse `json:"kategori"`
}

type ProductCreateResponse struct {
	ID         int    `json:"id"`
	SKU        string `json:"sku"`
	Name       string `json:"nama_produk"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
	Image      string `json:"image"`
	CategoryId int    `json:"kategori_id"`
}

type ProductUpdateResponse struct {
	SKU        string `json:"sku"`
	Name       string `json:"nama_produk"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
	Image      string `json:"image"`
	CategoryId int    `json:"kategori_id"`
}
