package main

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/olivere/elastic"
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
// 4. ElasticSearch

const host = "http://10.196.102.145:9200"

type Employee struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Comment   string   `json:"comment"`
	Interests []string `json:"interests"`
}

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintEmployee(result *elastic.SearchResult, err error) {
	if err != nil {
		panic(err)
	}

	var t Employee
	for _, item := range result.Each(reflect.TypeOf(t)) {
		e := item.(Employee)
		fmt.Println(e)
	}
}

func TestElasticSearch() {
	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false))
	PanicError(err)

	version, err := client.ElasticsearchVersion(host)
	PanicError(err)
	fmt.Println("Version:", version)

	ids := []string{}

	// 1. 创建数据
	e1 := Employee{"Jame", 20, "I am a boy.", []string{"music", "book"}}
	resp, err := client.Index().Index("bankdata").Type("employee").BodyJson(e1).Do(context.Background())
	PanicError(err)
	ids = append(ids, resp.Id)
	fmt.Println("New inserted id:", resp.Id)

	e2 := `{"name": "John", "age": 45, "comment":"I like sleeping", "interests":["Sleep", "book"]}`
	resp, err = client.Index().Index("bankdata").Type("employee").BodyJson(e2).Do(context.Background())
	PanicError(err)
	ids = append(ids, resp.Id)
	fmt.Println("New inserted id:", resp.Id)

	e3 := `{"name": "John", "age": 20, "comment":"I like swimming", "interests":["swim", "book"]}`
	resp, err = client.Index().Index("bankdata").Type("employee").BodyJson(e3).Do(context.Background())
	PanicError(err)
	ids = append(ids, resp.Id)
	fmt.Println("New inserted id:", resp.Id)

	// 2. 修改数据
	_, err = client.Update().Index("bankdata").Type("employee").Id(ids[1]).Doc(map[string]interface{}{"age": 88}).Do(context.Background())
	PanicError(err)

	// 3. 查找数据
	result, err := client.Get().Index("bankdata").Type("employee").Id(ids[0]).Do(context.Background())
	PanicError(err)
	if result.Found {
		fmt.Println(string(result.Source))
	}

	// 4. 查询数据
	// 取所有数据
	r, err := client.Search("bankdata").Type("employee").Do(context.Background())
	fmt.Println("All data:")
	PrintEmployee(r, err)

	q := elastic.NewQueryStringQuery("name: John")
	r, err = client.Search("bankdata").Type("employee").Query(q).Do(context.Background())
	fmt.Println("John data:")
	PrintEmployee(r, err)

	// 年龄大于三十的
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("name", "John"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	r, err = client.Search("bankdata").Type("employee").Query(boolQ).Do(context.Background())
	fmt.Println("John && age > 30:")
	PrintEmployee(r, err)

	// comment中含有swimming, sleeping
	matchQ := elastic.NewMatchPhraseQuery("comment", "swimming sleeping")
	r, err = client.Search("bankdata").Type("employee").Query(matchQ).Do(context.Background())
	fmt.Println("Contain swimming, sleeping:")
	PrintEmployee(r, err)

	// 分页
	r, err = client.Search("bankdata").Type("employee").Size(1).From(2).Do(context.Background())
	fmt.Println("Paging:")
	PrintEmployee(r, err)

	// 5. 删除数据
	_, err = client.Delete().Index("bankdata").Type("employee").Id(ids[1]).Do(context.Background())
	PanicError(err)
}

// -----------------------------------------------------------

func main() {
	testRedisConnPool()
}
