package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price float32
}

func (Product) TableName() string {
	return "T_Product"
}

// -----------------------------------------------------------

func testCRUD() {
	// db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("Failed to connect DB!")
	}
	defer db.Close()

	// 1. Create table
	if err := db.AutoMigrate(&Product{}).Error; err != nil {
		panic(err)
	}

	// 2. Insert data.
	p := &Product{Code: "A", Price: 1}
	if err := db.Create(p).Error; err != nil {
		panic(err)
	}

	fmt.Println("Affected rows: ", db.RowsAffected)
	fmt.Println("Last insert id: ", p.ID)

	db.Create(&Product{Code: "B", Price: 2})
	db.Create(&Product{Code: "C", Price: 3})

	// 3. Query data.
	var product Product
	// 获取第一条记录，按主键排序
	db.First(&product, 1) // 查询id为3的数据
	fmt.Println(product)

	product.ID = 0                      // product里面的id不为空时，会作为匹配条件
	db.First(&product, "code = ?", "B") // 查询code为B的数据
	fmt.Println(product)

	var products []Product
	db.Find(&products) // 查询所有记录
	fmt.Println(products)

	// 4. Update data.
	db.Model(&product).Update("price", 200) // 更新product这个数据，当前product为Code等于B的数据。

	product.ID = 0
	db.First(&product, "code = ?", "B") // 查询code为B的数据
	fmt.Println(product)

	// 	// 更新 - 更新product的price为2000
	// 	db.Model(&product).Update("Price", 2000)
}

// -----------------------------------------------------------

func main() {
	testCRUD()
}

// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

// func main() {
// 	// db, err := gorm.Open("sqlite3", ":memory:")
// 	db, err := gorm.Open("sqlite3", "test.db")

// 	if err != nil {
// 		panic("连接数据库失败")
// 	}
// 	defer db.Close()

// 	// 自动迁移模式
// 	db.AutoMigrate(&Product{})

// 	// 创建
// 	db.Create(&Product{Code: "L1212", Price: 1000})

// 	// 读取
// 	var product Product
// 	db.First(&product, 1)                   // 查询id为1的product
// 	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

// 	fmt.Println(product)

// 	// 更新 - 更新product的price为2000
// 	db.Model(&product).Update("Price", 2000)

// 	// 删除 - 删除product
// 	db.Delete(&product)
// }
