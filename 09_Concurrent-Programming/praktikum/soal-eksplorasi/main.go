package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Product struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func main() {
	apiURL := "https://fakestoreapi.com/products"
	numWorkers := 5 // Jumlah goroutine yang akan dijalankan

	// Membuat channel untuk mengirim data produk
	productChannel := make(chan Product)

	// Membuat Wait Group untuk menunggu semua goroutine selesai
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Membuat goroutine untuk mengambil data produk
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			fetchProducts(apiURL, productChannel)
		}()
	}

	// Menutup channel setelah semua goroutine selesai
	go func() {
		wg.Wait()
		close(productChannel)
	}()

	// Menerima dan mencetak data produk dari channel
	fmt.Println("products data")
	fmt.Println("===")
	for product := range productChannel {
		fmt.Printf("title: %s\nprice: %.2f\ncategory: %s\n===\n", product.Title, product.Price, product.Category)
	}
}

func fetchProducts(apiURL string, productChannel chan<- Product) {
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Gagal mengambil data:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Gagal mengambil data, status code:", resp.StatusCode)
		return
	}

	var products []Product
	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		fmt.Println("Gagal mengurai data:", err)
		return
	}

	// Mengirim data produk ke channel
	for _, product := range products {
		productChannel <- product
	}
}
