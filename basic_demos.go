package main

import (
	"bufio"
	"crypto/md5"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

//---------------------------------------
// 1. struct

type Student struct {
	Name string
	Age  int
}

func New(name string, age int) *Student {
	return &Student{
		Name: name,
		Age:  age,
	}
}

func (student *Student) SetName(name string) {
	student.Name = name
}

func testStruct() {
	s1 := New("A", 123)
	fmt.Println(s1.Name, s1.Age)

	s1.SetName("A1")
	fmt.Println(s1.Name, s1.Age)

	s2 := Student{"s2", 3}
	fmt.Println(&s1.Age, " ", &s2.Age)

	s3 := &Student{"A3", 44}
	fmt.Println(s3)
}

//---------------------------------------
// 2. map

func selectionSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

func testMap() {
	m1 := make(map[int32]string)
	m1[1] = "A"
	m1[2] = "B"
	fmt.Println(m1[1])

	fmt.Println(m1[3]) // ""访问不存在的键值，返回该类型默认值

	_, ok := m1[3]
	if !ok {
		fmt.Println("No")
	}

	// 通过slice排序来输出有序键值
	m2 := make(map[int]string)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		m2[r.Intn(100)] = strconv.Itoa(i) + "Abc"
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	fmt.Println("In order")

	s1 := []int{}
	for k := range m2 {
		s1 = append(s1, k)
	}

	selectionSort(s1)
	for _, k := range s1 {
		fmt.Println(k, m2[k])
	}
}

//---------------------------------------
// 3. ！slice传值

// 传值：传的是slice的值
func changeSlice(a []int) {
	a[0] = 3
	a = append(a, 2, 3, 4, 5, 6)
	fmt.Println(a) // [3 2 2 3 4 5 6]
}

// 传址：传的是slice的地址
func changeSlice2(a *[]int) {
	(*a)[0] = 5
	(*a) = append(*a, 4, 5, 6, 7, 8)
}

func testSlice() {
	a := []int{1, 2}
	changeSlice(a)
	fmt.Println(a) // [3 2]
	changeSlice2(&a)
	fmt.Println(a) // [5 2 4 5 6 7 8]
}

//---------------------------------------
// 4. ！字符串遍历

func testRangeFor() {
	s := "123阿斯蒂"

	// go语言中的字符串实际上是类型为byte的只读切片,所以输出都是字节流的的值
	// 49 50 51 38463 26031 33922
	for _, c := range s {
		fmt.Print(c, " ")
	}
	fmt.Println()

	// 49 50 51 233 152 191 230 150 175 232 146 130
	for _, c := range []byte(s) {
		fmt.Print(c, " ")
	}
	fmt.Println()

	// 49 50 51 38463 26031 33922
	for _, c := range []rune(s) {
		fmt.Print(c, " ")
	}
	fmt.Println()

	// 1 2 3 阿 斯 蒂
	for _, c := range s {
		fmt.Print(string(c), " ")
	}
	fmt.Println()

	// 49 50 51 233 152 191 230 150 175 232 146 130
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
}

//---------------------------------------
// 5. 多返回值

func swapInt(a int, b int) (int, int) {
	a, b = b, a
	return a, b
}

func testFunction() {
	a := 1
	b := 2
	a, b = swapInt(a, b)
	fmt.Println(a, " ", b)
}

//---------------------------------------
// 6. switch语句

func testSwitch() {
	for {
		x := 0
		_, ok := fmt.Scanln(&x)
		if ok != nil {
			break
		}
		// 默认break
		switch x {
		case 1, 2, 3:
			fmt.Println("A")
		case 5:
			fmt.Println("B")
		case 6:
			fmt.Println("fallthrough")
			fallthrough
		default:
			fmt.Println("default")
		}
	}
}

//---------------------------------------
// 7. 函数类型

type greeting func() string

func testFunc(f greeting) {
	fmt.Println(f())
}

func testFunc1() {
	testFunc(func() string {
		return "Hello"
	})
}

//---------------------------------------
// 8. Goroutine

func TimeAdd(i, j int) {
	time.Sleep(time.Second)
	fmt.Println(i + j)
}

func testAdd() {
	for i := 0; i < 10; i++ {
		go TimeAdd(i, i)
	}
	time.Sleep(time.Second * 3)
}

//---------------------------------------
// 9. slice

func testSlice1() {
	a1 := [3]int{1, 2, 3} // 数组[3]int
	fmt.Printf("%T\n", a1)

	a2 := [...]int{1, 2, 3, 4} // 数组[4]int
	fmt.Printf("%T\n", a2)

	s1 := []int{1, 2, 3, 4} // 切片[]int
	fmt.Printf("%T\n", s1)

	var s2 []int // 切片[]int
	fmt.Printf("%T\n", s2)

	s3 := a1[:] // 切片[]int
	fmt.Printf("%T\n", s3)

	// 合并两个切片
	/// ...的作用
	// 第一个用法主要是用于函数有多个不定参数的情况，可以接受多个不确定数量的参数。
	// 第二个用法是slice可以被打散进行传递。
	s4 := make([]int, 3, 4)
	s4 = append(s4, s1...)
	fmt.Println(s4) // [0 0 0 1 2 3 4]

	// 容量不够时，容量以翻倍的策略扩容，找到一片新的连续内存，将原有元素拷贝过去
	var s5 []int

	s5 = append(s5, 1)
	// slice address: 0xc0000046e0, Len:1, cap: 1, data address:0xc000012288
	fmt.Printf("slice address: %p, Len:%d, cap: %d, data address:%p\n", &s5, len(s5), cap(s5), &s5[0])

	s5 = append(s5, 1)
	// slice address: 0xc0000046e0, Len:2, cap: 2, data address:0xc0000122a0
	fmt.Printf("slice address: %p, Len:%d, cap: %d, data address:%p\n", &s5, len(s5), cap(s5), &s5[0])

	s5 = append(s5, 1)
	// slice address: 0xc0000046e0, Len:3, cap: 4, data address:0xc00000a680
	fmt.Printf("slice address: %p, Len:%d, cap: %d, data address:%p\n", &s5, len(s5), cap(s5), &s5[0])

	s5 = append(s5, 1)
	// slice address: 0xc0000046e0, Len:4, cap: 4, data address:0xc00000a680
	fmt.Printf("slice address: %p, Len:%d, cap: %d, data address:%p\n", &s5, len(s5), cap(s5), &s5[0])
}

//---------------------------------------
// 10. Defer,panic

// defer 传入的函数不是在退出代码块的作用域时执行的，它只会在当前函数和方法返回之前被调用
// 函数调用panic 时会立刻停止执行函数的其他代码，并在执行结束后在当前 Goroutine 中递归执行调用方的延迟函数调用 defer
func testDefer() {
	for i := 0; i < 2; i++ {
		fmt.Println(i)
		defer fmt.Println("defer1: ", i)
	}
	fmt.Println("end1")

	for i := 0; i < 2; i++ {
		fmt.Println(i)
		defer fmt.Println("defer2: ", i)
	}
	panic("")
	fmt.Println("end2")

	// Output:
	// 0
	// 1
	// end1
	// 0
	// 1
	// defer2:  1
	// defer2:  0
	// defer1:  1
	// defer1:  0
	// panic:
}

//---------------------------------------
// 11. MD5

func testMd5() {
	m := md5.New()
	m.Write([]byte("12345"))
	r := m.Sum([]byte(""))
	fmt.Printf("%x\n", r)
}

//---------------------------------------
// 12. Http

func testHttp() {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!"))
		})
		http.ListenAndServe("127.0.0.1:8080", nil)
	}()

	time.Sleep(time.Second)

	res, _ := http.Get("http://127.0.0.1:8080")
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

//---------------------------------------
// 13. Interface

type Animal1 interface {
	fly()
}

type Animal2 interface {
	fly()
	run()
}

type Bird struct {
}

func (bird *Bird) fly() {
	fmt.Println("fly")
}

func (bird *Bird) run() {
	fmt.Println("run")
}

func testInterface() {
	var a1 Animal1
	a1 = new(Bird)
	a1.fly()

	var a2 Animal2
	a2 = new(Bird)
	a2.fly()
	a2.run()
}

//---------------------------------------
// 14. Json

type People struct {
	Name string `json:"peopleName"`
	Age  int
}

func testJson() {
	a := [5]int{1, 2, 3, 4}
	s, _ := json.Marshal(a)
	fmt.Println(string(s))

	m := make(map[string]string)
	m["A"] = "1"
	m["B"] = "2"
	m["C"] = "3"
	s, _ = json.Marshal(m)
	fmt.Println(string(s))

	m1 := make(map[string]interface{})
	m1["Name"] = "AAA"
	m1["Age"] = 15
	m1["Die"] = true
	s, _ = json.MarshalIndent(m1, "", "  ")
	fmt.Println(string(s))

	p := People{
		"Abc",
		12,
	}

	s, _ = json.MarshalIndent(p, "", "  ")
	fmt.Println(string(s))

	var j interface{}
	json.Unmarshal(s, &j)
	fmt.Printf("%v\n", j)

	var p1 People
	json.Unmarshal(s, &p1)
	fmt.Println(p1)

	// 读写json文件
	file, _ := os.OpenFile("./test.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	encoder := json.NewEncoder(file)
	encoder.Encode(p)
	file.Close()

	var p2 People
	file, _ = os.Open("./test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&p2)
	fmt.Println(p2)
	file.Close()

	os.Remove("./test.json")
}

//---------------------------------------
// 15. Channel

func testChannel() {
	// ch := make(chan int)
	ch := make(chan int, 5)

	go func(ch chan int) {
		for i := 0; i < 20; i++ {
			ch <- i
			fmt.Println("Send: ", i)
		}
	}(ch)

	go func(ch chan int) {
		for i := 0; i < 20; i++ {
			n := <-ch
			fmt.Println("Receive: ", n)
		}

	}(ch)

	time.Sleep(time.Second)
}

//---------------------------------------
// 16. Channel

func createWoker(name string, ch chan int) func() {
	return func() {
		for true {
			i := <-ch
			time.Sleep(time.Millisecond * 10)
			fmt.Println(name, ": ", i)
		}
	}
}

func testChannel2() {
	ch := make(chan int, 100)
	for i := 0; i < 4; i++ {
		go createWoker(strconv.Itoa(i), ch)()
	}

	for i := 0; i < 100; i++ {
		ch <- i
	}
	time.Sleep(time.Second * 2)
}

//---------------------------------------
// 17. Select关键字

func testSelect1() {
	ch := make(chan int)
	go func() {
		ch <- 10
	}()

	time.Sleep(10 * time.Millisecond)

	select {
	case <-ch:
		fmt.Println("Read data")
	default:
		fmt.Println("default")
	}
}

func testSelect2() {
	ch := make(chan int)
	timeout := make(chan bool)
	go func() {
		time.Sleep(time.Second * 3)
		timeout <- true
	}()

	select {
	case <-ch:
		fmt.Println("Read data")
	case <-timeout:
		fmt.Println("timeout")
	}
}

func testSelect3() {
	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println("Read data")
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")
	}
}

//---------------------------------------
// 18. iota

// Go没有enum类型，用iota来自增
const (
	c1 = iota
	c2
	c3
)

func testIota() {
	fmt.Println(c1, c2, c3)
}

//---------------------------------------
// 19. math

func testMath() {
	fmt.Println(math.Pow(2, 4)) //16

	// Round：四舍五入
	fmt.Println(math.Round(2.4)) // 2
	fmt.Println(math.Round(2.5)) //3

	// Ceil: 天花板
	fmt.Println(math.Ceil(2.4)) //3
	fmt.Println(math.Ceil(2.5)) //3

	// Floor: 地板
	fmt.Println(math.Floor(2.4)) //2
	fmt.Println(math.Floor(2.5)) //2
}

//---------------------------------------
// 20. random

func testRandom() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(101)

	for {
		var input string
		fmt.Scan(&input)

		if input == "exit" {
			break
		}

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		if number == n {
			fmt.Println("Win")
			break
		} else if number > n {
			fmt.Println("Big")
		} else {
			fmt.Println("Small")
		}
	}
}

//---------------------------------------
// 21. 不定长参数

// 不定长参数, 类型为切片，并且不定长参数只能位于参数列表的末尾
func variableParams(c int, names ...string) {
	fmt.Printf("%T\n", names) // []string

	for _, n := range names {
		fmt.Println(n)
	}
}

func testFunc2() {
	variableParams(2, "A", "V", "C")
}

//---------------------------------------
// 22. closure

// 闭包的作用是保存函数的状态，类似C++中的仿函数
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func testClosure() {
	counter := createCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

//---------------------------------------
// 23. nil

func testNil() {
	// nil是不能比较的
	// 如果两个比较的nil值之一是一个接口值, 而另一个不是, 假设它们是可比较的, 则比较结果总是 false。
	// 原因是在进行比较之前, 接口值将转换为接口值的类型。转换后的接口值具有具体的动态类型, 但其他接口值没有。这就是为什么比较结果总是错误的。
	// fmt.Println(nil == nil)  // invalid operation: nil == nil (operator == not defined on nil)

	// 按照Go语言规范，任何类型在未初始化时都对应一个零值：布尔类型是false，整型是0,
	// 字符串是""，而指针、函数、interface、slice、channel和map的零值都是nil。
	var p *int
	fmt.Println(p)        // <nil>
	fmt.Println(p == nil) // true

	var s []int
	fmt.Println(s)        // []
	fmt.Println(s == nil) // true
}

//---------------------------------------
// 24. strings

func testStrings() {
	fmt.Println(strings.Contains("Abc", "Ab")) // true
	fmt.Println(strings.Contains("Abc啊", "啊")) // true

	fmt.Println(strings.Index("ABC", "BC"))     // 1
	fmt.Println(strings.IndexAny("ABC", "CBA")) // 0

	fmt.Println(strings.ToUpper("abc"))
	fmt.Println(strings.ToLower("ABC"))
	fmt.Println(strings.Title("hello world")) // Hello World

	fmt.Println(strings.TrimSpace(" \n\ta\t\n ")) // a
	fmt.Println(strings.Trim("!!a;bc;;", "!;"))   // a;bc

	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("b", "a")) //1
	fmt.Println(strings.Compare("a", "a")) //0

	fmt.Println(strings.Split("a;b;c", ";"))  // [a b c]
	fmt.Println(strings.Split("a;b;c;", ";")) // [a b c ]

	fmt.Println(strings.SplitAfter("a;b;c", ";"))  // [a; b; c]
	fmt.Println(strings.SplitAfter("a;b;c;", ";")) // [a; b; c; ]

	fmt.Println(strings.ReplaceAll("AbAcAd", "A", "HH")) // HHbHHcHHd

	fmt.Println(strings.Join([]string{"A", "B", "C", "D"}, ";")) // A;B;C;D

	// %U: Unicode格式：U+1234，等同于 "U+%04X"
	// %c:相应Unicode码点所表示的字符
	fmt.Printf("%c\n", 'a') // a
	fmt.Printf("%c\n", '啊') // 啊
	fmt.Printf("%U\n", 'a') // U+0061
	fmt.Printf("%U\n", '啊') // U+554A
}

//---------------------------------------
// 24. flag

func testFlag() {
	fmt.Println(os.Args)

	name := flag.String("name", "", "Input Name")
	age := flag.Int("age", 0, "Input Age")
	check := flag.Bool("check", false, "Input check")

	flag.Parse()

	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*check)
}

//---------------------------------------
// 25. 继承

type Animal struct {
	name string
}

func (animal *Animal) SetName(name string) {
	animal.name = name
}

func (animal Animal) GetName() string {
	return animal.name
}

func (animal Animal) Eat() {
	fmt.Println("Eating")
}

type Fish struct {
	Animal // 持有父类，相当于继承父类
	fins   int
}

func (fish *Fish) Swim() {
	fmt.Println("Swimming")
}

func testInherit() {
	f := &Fish{Animal{"fish"}, 4}
	f.Eat()
	f.Swim()

	f.SetName("fish2")
	fmt.Println(f.GetName())

	fmt.Println(f.name)
	fmt.Println(f.fins)
}

//---------------------------------------
// 26. 多态

// 接口Window
type Window interface {
	Create(title string) bool
	PrintTitle()
}

// 类型Panel
type Panel struct {
	title string
}

func (panel *Panel) Create(title string) bool {
	panel.title = title
	fmt.Println("Create Panel!")
	return true
}

func (panel *Panel) PrintTitle() {
	fmt.Println("Print panel title:", panel.title)
}

// 类型Control
type Control struct {
	title string
}

func (control *Control) Create(title string) bool {
	control.title = title
	fmt.Println("Create Control!")
	return true
}

func (control Control) PrintTitle() {
	fmt.Println("Print control title:", control.title)
}

// 类型ComboBox
// 继承Control类，同时实现了Window接口
type ComboBox struct {
	Control
}

func (cb *ComboBox) Create(title string) bool {
	cb.title = title
	fmt.Println("Create ComboBox!")
	return true
}

func (cb ComboBox) PrintTitle() {
	fmt.Println("Print ComboBox title:", cb.title)
}

func (cb ComboBox) GetSelection() {
	fmt.Println("Select 1")
}

func testPolymorphism() {
	windows := []Window{}
	windows = append(windows, new(Panel))
	windows = append(windows, new(Control))
	windows = append(windows, new(ComboBox))

	// 多态
	for _, w := range windows {
		w.Create("123")
	}

	for _, w := range windows {
		w.PrintTitle()
	}

	// 断言类型的语法：x.(T)，这里x表示一个接口的类型，T表示一个类型（也可为接口类型）。
	// 一个类型断言检查一个接口对象x的动态类型是否和断言的类型T匹配。

	// Type switch
	for _, w := range windows {
		switch w.(type) {
		case *Panel:
			fmt.Println("I am Panel")
		case *Control:
			fmt.Println("I am Control")
		case *ComboBox:
			fmt.Println("I am ComboBox")
		}
	}

	// Type assertion
	for _, w := range windows {
		if cb, ok := w.(*ComboBox); ok {
			cb.GetSelection()
		}
	}

	// 接口实例可以接收子类指针
	// 父类实例指针不能接收子类指针
	// 也就是说定义了接口才能实现多态
	cb := new(ComboBox)
	// var c *Control
	// c = cb  // cannot use cb (type *ComboBox) as type *Control in assignment
	var w Window
	w = cb
	w.PrintTitle()
}

//---------------------------------------
// 27. File IO

func testFileIO() {
	data := `int main() {
		return 0;
}`

	// 1. ioutil.WriteFile：一次覆盖写整个数据。
	// perm: 如果过创建文件，赋给perm这个权限。生成的文件的权限是-rw-rw-rw-
	ioutil.WriteFile("./test.txt", []byte(data), 0666)

	file, _ := os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0x666)

	// 2. bufio.NewWriter: 使用缓冲区写
	// bufio模块通过对io模块的封装，提供了数据缓冲功能，能够一定程度减少大块数据读写带来的开销。
	// 实际上在bufio各个组件内部都维护了一个缓冲区，数据读写操作都直接通过缓存区进行。
	// 当发起一次读写操作时，会首先尝试从缓冲区获取数据；只有当缓冲区没有数据时，才会从数据源获取数据更新缓冲。
	writer := bufio.NewWriter(file)
	writer.WriteString(" /* 注释")
	writer.Flush()

	// 3. file.Write: 直接写
	file.Write([]byte("*/"))
	file.Close()

	// 4. ioutil.ReadFile：一次读出所有数据
	bytes, _ := ioutil.ReadFile("./test.txt")
	fmt.Println(string(bytes))

	// 5. os.Open: 只读打开
	file, _ = os.Open("./test.txt")

	// 6. file.Read: 直接读
	byte3 := make([]byte, 0)
	bytes2 := make([]byte, 8)
	for {
		count, err := file.Read(bytes2)
		if err == io.EOF {
			break
		}

		byte3 = append(byte3, bytes2[:count]...)
	}
	fmt.Println(string(byte3))
	file.Close()

	file, _ = os.OpenFile("./test.txt", os.O_RDONLY, 0x666)

	// 7. bufio.NewReader：缓冲区读
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Println(str) // 读完最后一行，则err == io.EOF
		if err == io.EOF {
			break
		}
	}
	file.Close()

	file, _ = os.OpenFile("./test.txt", os.O_RDONLY, 0x666)

	// 7. bufio.NewScanner：缓冲区读
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text()) // 逐行读
	}

	file.Close()

	os.Remove("./test.txt")
}

//---------------------------------------
// 28. OS

func testOS() {
	wd, _ := os.Getwd()
	fmt.Println(wd)

	fmt.Println(os.Getenv("path"))
	fmt.Println(os.Environ())

	h, _ := os.Hostname()
	fmt.Println(h)

	fmt.Println(os.TempDir())

	fileinfo, err := os.Stat("./main.go")
	if err == nil {
		fmt.Println(fileinfo.IsDir())
		fmt.Println(fileinfo.Mode())    // -rw-rw-rw-
		fmt.Println(fileinfo.ModTime()) // 2020-06-02 00:42:35.2746246 +0800 CST
		fmt.Println(fileinfo.Name())    // main.go
		fmt.Println(fileinfo.Size())
	}
}

//---------------------------------------
// 29. time

func testTime() {
	now := time.Now()
	fmt.Println(now.Year())  // 2020
	fmt.Println(now.Month()) // June
	fmt.Println(now.Day())   // 2 月中的第几天

	fmt.Println(now.Date())    // 2020 June 2
	fmt.Println(now.YearDay()) // 154
	fmt.Println(now.Weekday()) // Tuesday

	fmt.Println(now.Hour())       // 0
	fmt.Println(now.Minute())     // 50
	fmt.Println(now.Second())     // 32
	fmt.Println(now.Nanosecond()) // 81211400

	d := time.Date(2020, time.June, 1, 12, 12, 12, 0, time.Now().Location())
	fmt.Println(d)

	// 2020-06-02 01:00:27.0953827 +0800 CST m=+0.004996701
	fmt.Println(now)

	duration, _ := time.ParseDuration("24h0m0s")
	// 2020-06-03 01:00:27.0953827 +0800 CST m=+86400.004996701
	fmt.Println(now.Add(duration))

	// Sub()是求时间差，Add()才是增加或者减少时间，用正负控制
	// 2020-06-01 01:00:27.0953827 +0800 CST m=-86399.995003299
	fmt.Println(now.Add(duration * -1))

	// -12h48m15.0953827s
	fmt.Println(d.Sub(now))
}

//---------------------------------------
// 30. panic&error

type InvalidParamError struct {
	invalidValue int
}

func (err InvalidParamError) Error() string {
	return fmt.Sprintf("Invalid param: %d.", err.invalidValue)
}

func validateA(a int) error {
	if a == 0 {
		panic("Can't be zero")
	}

	if a == 1 || a == 2 || a == 3 || a == 4 {
		return &InvalidParamError{a}
	}

	if a == 7 {
		panic(&InvalidParamError{a})
	}

	return nil
}

func testPanic() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("End1")
		}
	}()

	if err := validateA(2); err != nil {
		fmt.Println(err)
	}

	// validateA(0)
	validateA(7)

	fmt.Println("End2")
}

//---------------------------------------
// 31. 反射

type RParent struct {
	Name string
	Age  int
}

type RSon struct {
	RParent
	School string
}

func (s RSon) CanSwim() bool {
	return true
}

func (s *RSon) Set(name string, age int, school string) {
	s.Name = name
	s.Age = age
	s.School = school
}

func testReflect() {
	var object interface{}
	object = RSon{RParent{"AAA", 14}, "BBCC"}

	// 1. reflect.TypeOf: 返回类型信息

	objType := reflect.TypeOf(object)
	fmt.Println(objType)        // main.RSon
	fmt.Println(objType.Name()) // RSon (类型名称)
	fmt.Println(objType.Kind()) // struct (原始类型)

	fmt.Println(objType.NumField())  // 2 (属性个数，RParent算一个整体属性)
	fmt.Println(objType.NumMethod()) // 1 (方法个数)

	// 遍历属性名称和类型
	// RParent main.RParent
	// School string
	for i := 0; i < objType.NumField(); i++ {
		filed := objType.Field(i)
		fmt.Println(filed.Name, filed.Type)
	}

	// 遍历方法名称和类型
	// Swim func(main.RSon)
	for i := 0; i < objType.NumMethod(); i++ {
		method := objType.Method(i)
		fmt.Println(method.Name, method.Type)
	}

	// 获取父类属性，(0,1)代表RSon的第零个属性中的第一个属性
	fmt.Println(objType.FieldByIndex([]int{0, 1}).Name) // Age

	// 2. reflect.ValueOf: 返回值信息

	objValue := reflect.ValueOf(object)
	fmt.Println(objValue) // {{AAA 14} BBCC}

	// 遍历属性的值
	// {AAA 14}
	// BBCC
	for i := 0; i < objValue.NumField(); i++ {
		fmt.Println(objValue.Field(i).Interface())
	}

	// 获取父类属性的值
	fmt.Println(objValue.FieldByIndex([]int{0, 0}).Interface()) // AAA

	// 该类型不能修改值，指针类型才可以
	// panic: reflect: reflect.flag.mustBeAssignable using unaddressable value
	// objValue.FieldByName("School").SetString("123")

	// 3. 修改属性值

	var objectp interface{}
	r := &RSon{RParent{"AAA", 14}, "BBCC"}
	objectp = r

	objpValue := reflect.ValueOf(objectp)
	fmt.Println(objpValue.Kind()) // ptr

	// 获取指针指向的对象
	elem := objpValue.Elem()

	// 遍历属性的值
	// {AAA 14}
	// BBCC
	for i := 0; i < elem.NumField(); i++ {
		fmt.Println(elem.Field(i).Interface())
	}

	// 修改属性
	elem.FieldByName("School").SetString("1234")
	fmt.Println(r) // &{{AAA 14} 1234}

	// 调用方法
	m1 := objpValue.MethodByName("CanSwim")
	r1 := m1.Call([]reflect.Value{})
	fmt.Println(r1[0].Interface()) // true

	m2 := objpValue.MethodByName("Set")
	m2.Call([]reflect.Value{reflect.ValueOf("DDD"), reflect.ValueOf(123), reflect.ValueOf("CCC")})
	fmt.Println(r) // &{{DDD 123} CCC}
}

//---------------------------------------
// 32. regex

func testRegex() {
	r, err := http.Get("https://www.haomagujia.com/")
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(r.Body)
	html := string(data)
	fmt.Print(html)

	reg := regexp.MustCompile(`1[3-9]\d{9}`)
	s := reg.FindAllStringSubmatch(html, -1)
	for _, x := range s {
		fmt.Println(x[0])
	}
}

//---------------------------------------
// 33. csv

func testCSV() {
	records := [][]string{
		{"Id", "Name", "Age"},
		{"1", "A\"\"", "11"},
		{"2", "B D", "12"},
		{"3", "C,M", "13"},
	}

	w := csv.NewWriter(os.Stdout)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			panic(err)
		}
	}

	w.Flush()

	stringBuilder := strings.Builder{}
	w2 := csv.NewWriter(&stringBuilder)
	w2.WriteAll(records)
	fmt.Println(stringBuilder.String())

	r := csv.NewReader(strings.NewReader(stringBuilder.String()))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(record)
	}
}

//---------------------------------------
// 34. unsafe

func testUnsafe() {
	var a1 int32
	fmt.Println(unsafe.Sizeof(a1)) // 4

	type T struct {
		a byte  // 1
		b int32 // 4
		c int64 // 8
	}

	// 内存对齐
	a2 := T{1, 4, 8}
	fmt.Println(unsafe.Sizeof(a2)) // 16

	// 查看对齐方式
	fmt.Println(unsafe.Alignof(a2))   // 8
	fmt.Println(unsafe.Alignof(a2.a)) // 1
	fmt.Println(unsafe.Alignof(a2.b)) // 4
	fmt.Println(unsafe.Alignof(a2.c)) // 8

	a3 := new(T)
	fmt.Println(unsafe.Sizeof(a3)) // 8
}

//---------------------------------------

func main() {
	testUnsafe()
}
