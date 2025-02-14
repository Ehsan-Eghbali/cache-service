package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

// ایجاد یک cache با زمان انقضای پیش‌فرض 5 دقیقه و پاکسازی هر 10 دقیقه
var c *cache.Cache

func init() {
	c = cache.New(5*time.Minute, 10*time.Minute)
}

// تابع شبیه‌سازی شده برای واکشی داده (مثلاً از دیتابیس یا سرویس دیگر)
func fetchData(key string) string {
	// شبیه‌سازی تاخیر در واکشی داده
	time.Sleep(100 * time.Millisecond)
	return "value for " + key
}

// هندلر برای دریافت داده
func getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing key parameter", http.StatusBadRequest)
		return
	}

	// بررسی وجود کلید در cache
	if cached, found := c.Get(key); found {
		response := map[string]string{
			"key":    key,
			"value":  cached.(string),
			"source": "cache",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// واکشی داده از منبع اصلی در صورت عدم وجود در cache
	value := fetchData(key)
	// ذخیره داده در cache
	c.Set(key, value, cache.DefaultExpiration)
	response := map[string]string{
		"key":    key,
		"value":  value,
		"source": "database",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/data", getHandler)
	log.Println("Starting in-memory cache microservice on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
