package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
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
// 1. gorm crud

func testGormCrud() {
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

	// 5. Delete data.
	db.Delete(&product)                               // 如果有deleteAt字段，软删除
	db.Unscoped().Where(`code = "B"`).Find(&products) // 查看软删除字段
	fmt.Println(products)

	db.Unscoped().Delete(&product) // 真正删除
	db.Unscoped().Where(`code = "B"`).Find(&products)
	fmt.Println(products)
}

// -----------------------------------------------------------
// 2. redis

func testRedis() {
	conn, err := redis.Dial("tcp", "10.196.102.145:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	r, err := conn.Do("set", "key1", "value1")
	r, err = conn.Do("get", "key1")
	s, _ := redis.String(r, err)
	fmt.Println(s)

	r, err = conn.Do("rpush", "key2", 11, 2, 3, 4)
	r, err = conn.Do("lrange", "key2", 0, -1)
	ss, _ := redis.Ints(r, err)
	fmt.Println(ss)

	// 设置超时时间
	_, err = conn.Do("expire", "key2", 5)
	time.Sleep(time.Second * 5)
	r, err = conn.Do("lrange", "key2", 0, -1)
	ss, _ = redis.Ints(r, err)
	fmt.Println(ss)
}

// -----------------------------------------------------------
// 3. Redis connection pool

func testRedisConnPool() {
	pool := &redis.Pool{
		// 最大活动连接数
		MaxActive: 100,

		// 最大闲置连接数
		MaxIdle: 20,

		//闲置连接的超时时间
		IdleTimeout: time.Second * 100,

		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "10.196.102.145:6379")
			return conn, err
		},
	}
	defer pool.Close()

	for i := 0; i < 10; i++ {
		go func() {
			conn := pool.Get()
			defer conn.Close()

			r, err := conn.Do("set", "key"+strconv.Itoa(i), i)
			s, _ := redis.String(r, err)
			fmt.Println(s)
		}()
	}

	time.Sleep(3 * time.Second)
}

// -----------------------------------------------------------

func main() {
	testRedisConnPool()
}
