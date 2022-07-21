package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	//var (
	//	u8  uint8     = 255
	//	i8  int8      = 127
	//	f32 float32   = 0.2
	//	c64 complex64 = -5 + 12i
	//)
	//
	//x := 1 + 1
	//fmt.Println(x)
	//fmt.Println(1+1, 2+2)
	//fmt.Println(u8, i8, f32, c64)
	//fmt.Printf("%T\n", u8)
	//fmt.Printf("%T\n", i8)
	//fmt.Printf("%T\n", f32)
	//fmt.Printf("%T\n", c64)

	/// increment, decrement
	//x := 0
	//fmt.Println(x)
	//x++
	//fmt.Println(x)
	//x--
	//fmt.Println(x)
	//x += 1
	//fmt.Println(x)
	//x -= 1
	//fmt.Println(x)

	/// cast
	//var x int = 1
	//xx := float64(x)
	//fmt.Printf("%T, %v, %f\n", xx, xx, xx)
	//
	//var y float64 = 1.5
	//yy := int(y)
	//fmt.Printf("%T, %v, %d\n", yy, yy, yy)
	//
	//var str string = "123"
	//num, _ := strconv.Atoi(str)
	//fmt.Printf("%T, %v", num, num)

	/// array
	//var a [2]int
	//a[0] = 100
	//a[1] = 200
	//fmt.Println(a)
	//
	//var b [2]int = [2]int{100, 200}
	//fmt.Println(b)
	//
	//c := []int{1, 2, 3, 4, 5, 6}
	//c = append(c, 7, 8, 9)
	//fmt.Println(c)
	//fmt.Println(c[1:3])
	//fmt.Println(c[:3])
	//fmt.Println(c[3:])
	//
	//board := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}
	//fmt.Println(board)

	/// slice make cap
	//n := make([]int, 3, 5)
	//fmt.Printf("len=%d, cap=%d, value=%v\n", len(n), cap(n), n)
	//n = append(n, 0)
	//log1(n)
	//n = append(n, 1, 2, 3, 4, 5)
	//log1(n)
	//
	//a := make([]int, 3)
	//log1(a)
	//b := make([]int, 0) // empty
	//log1(b)
	//var c []int // nil
	//log1(c)
	//d := []int{} // empty
	//fmt.Printf("%t\n", b == nil)
	//fmt.Printf("%t\n", c == nil)
	//fmt.Printf("%t\n", d == nil)
	//
	//c = make([]int, 5)
	//for i := 0; i < 6; i++ {
	//	c = append(c, 1)
	//	fmt.Println(c)
	//}
	//fmt.Println(c)
	//log1(c)
	//
	//var c1 []int
	//c1 = make([]int, 0, 5)
	//for i := 0; i < 6; i++ {
	//	c1 = append(c1, 1)
	//	fmt.Println(c1)
	//}
	//fmt.Println(c1)
	//log1(c1)

	/// Map
	//m := map[string]int{"apple": 100, "banana": 200}
	//fmt.Println(m)
	//fmt.Println(m["apple"])
	//
	//m["banana"] = 300
	//fmt.Println(m)
	//
	//m["new"] = 500
	//fmt.Println(m)
	//
	//fmt.Println(m["nothing"])
	//
	//v, exists := m["apple"]
	//fmt.Println(v, exists)
	//
	//v2, exists2 := m["nothing"]
	//fmt.Println(v2, exists2)
	//
	//m2 := make(map[string]int)
	//fmt.Println(m2)
	//m2["pc"] = 5000
	//fmt.Println(m2)
	//
	//var m3 map[string]int
	////var m3 map[string]int = map[string]int{}
	//if m3 == nil {
	//	fmt.Println("m3 is nil")
	//	fmt.Println(m3)
	//} else {
	//	fmt.Println("m3 is not nil")
	//	m3["pc"] = 5000
	//	fmt.Println(m3)
	//}

	/// byte
	//b := []byte{72, 73}
	//fmt.Println(b)
	//fmt.Println(string(b))
	//
	//c := []byte("HI")
	//fmt.Println(c)
	//fmt.Println(string(c))

	/// function
	//ra, rs := calc(10, 20)
	//fmt.Println(ra, rs)
	//
	//r3 := cal(100, 2)
	//fmt.Println(r3)
	//
	//f := func(x int) {
	//	fmt.Println("inner func", x)
	//}

	//f(1)
	//
	//func(x int) {
	//	fmt.Println("inner func", x)
	//}(2)

	/// closure
	//counter := incrementGenerator()
	//fmt.Println(counter())
	//fmt.Println(counter())
	//fmt.Println(counter())
	//
	//c1 := circleArea(3.14)
	//fmt.Println(c1(2))
	//fmt.Println(c1(3))
	//
	//c2 := circleArea(3)
	//fmt.Println(c2(2))
	//fmt.Println(c2(3))

	/// variable arguments
	//foo()
	//foo(10, 20)
	//foo(10, 20, 30)
	//
	//s := []int{1, 2, 3}
	//foo(s...) //展開が必要

	///演習
	/////Q1
	//f := 1.11
	//fmt.Println(int(f))
	/////Q2
	//s := []int{1, 2, 5, 6, 2, 3, 1}
	//fmt.Println(s[2:4]) //5,6
	/////Q3
	//m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	//fmt.Printf("%T %v", m, m)

	///if文
	//result := by2(10)
	//if result == "ok" {
	//	fmt.Println("great")
	//}
	//fmt.Println(result)
	//
	//if result2 := by2(10); result2 == "ok" {
	//	fmt.Println("great 2")
	//	fmt.Println(result2)
	//}

	/// for文
	//// 標準
	//for i := 0; i < 10; i++ {
	//	if i == 3 {
	//		fmt.Println("continue")
	//		continue
	//	}
	//	if i > 5 {
	//		fmt.Println("break")
	//		break
	//	}
	//	fmt.Println(i)
	//}
	//
	//// 初期化を外だし
	//sum := 1
	//for sum < 10 {
	//	sum += sum
	//	fmt.Println(sum)
	//}
	//
	////無限ループ
	//for {
	//	fmt.Println("hello")
	//}

	/// range
	//l := []string{"python", "go", "java"}
	//
	//for i := 0; i < len(l); i++ {
	//	fmt.Println(i, l[i])
	//}
	//
	//for i, v := range l {
	//	fmt.Println(i, v)
	//}
	//
	//for _, v := range l {
	//	fmt.Println(v)
	//}
	//
	//m := map[string]int{"apple": 100, "banana": 200}
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}
	//
	//for _, v := range m {
	//	fmt.Println(v)
	//}
	//
	//for k := range m {
	//	fmt.Println(k)
	//}

	/// Switch
	//os1 := getOsName()
	//switch os1 {
	//case "mac":
	//	fmt.Println("Mac")
	//case "windows":
	//	fmt.Println("Windows")
	//default:
	//	fmt.Println("default")
	//}
	//
	//switch os2 := getOsName(); os2 {
	//case "mac":
	//	fmt.Println("Mac")
	//case "windows":
	//	fmt.Println("Windows")
	//}
	//
	//h := time.Now()
	//switch {
	//case h.Hour() < 12:
	//	fmt.Println("Morning")
	//case h.Hour() < 17:
	//	fmt.Println("Afternoon")
	//case h.Hour() < 24:
	//	fmt.Println("Night")
	//}

	/// Defer
	//fmt.Println("run")
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	//fmt.Println("success")
	//
	//file, _ := os.Open("./lesson.go")
	//defer file.Close()
	//data := make([]byte, 100)
	//file.Read(data)
	//fmt.Println(string(data))

	/// Logging
	//LoggingSettings("test.log")
	//_, err := os.Open("foo")
	//if err != nil {
	//	log.Fatalln("Fatal Error", err)
	//}
	//log.Println("logging")
	//log.Printf("%T, %v", "test", "test")
	//
	//log.Fatalf("%T, %v", "fatal", "fatal")
	//log.Fatalln("fatal")
	//
	//fmt.Println("no display log")

	/// Error Handling
	//file, err := os.Open("./lesson.go")
	//if err != nil {
	//	log.Fatalln("Error Open", err)
	//}
	//defer file.Close()
	//data := make([]byte, 100)
	//count, err := file.Read(data) // 初期化が1つでも可能なら:=が使える. errは最初のやつの使い回し
	//if err != nil {
	//	log.Fatalln("Error")
	//}
	//fmt.Println(count, string(data))
	//
	//if err = os.Chdir("test"); err != nil {
	//	log.Fatalln("Error Chdir", err)
	//}

	/// Panic & Recover
	//thirdPartyConnectDB := func() {
	//	panic("Unable to connect database")
	//}
	//
	//save := func() {
	//	defer func() {
	//		s := recover()
	//		fmt.Println(s)
	//	}()
	//	thirdPartyConnectDB()
	//}
	//
	//save()
	//fmt.Println("is OK?")

	///演習
	////Q1
	//l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	//sort.Ints(l)
	//fmt.Println(l[0])
	//
	////Q2
	//m := map[string]int{
	//	"apple":  200,
	//	"banana": 300,
	//	"grapes": 150,
	//	"orange": 80,
	//	"papaya": 500,
	//	"kiwi":   90,
	//}
	//sum := 0
	//for _, v := range m {
	//	sum += v
	//}
	//fmt.Println(sum)

	/// Pointer
	//var n int = 100
	//one(&n)
	//fmt.Println(n, &n)

	/// Make & New
	/// ポインタの時のみnew.そしてvarでの初期化が必要
	//s := make([]int, 0)
	//fmt.Printf("%T\n", s)
	//
	//m := make(map[string]int)
	//fmt.Printf("%T\n", m)
	//
	//ch := make(chan int)
	//fmt.Printf("%T\n", ch)
	//
	//var p *int = new(int)
	//fmt.Printf("%T\n", p)
	//
	//var st = new(struct{})
	//fmt.Printf("%T\n", st)

	/// Struct
	//v := Vertex{x: 1, y: 2}
	//fmt.Println(v)
	//fmt.Println(v.x, v.y)
	//v.x = 100
	//fmt.Println(v.x, v.y)
	//
	//v2 := Vertex{x: 1}
	//fmt.Println(v2)
	//
	//v3 := Vertex{1, 2, "test"}
	//fmt.Println(v3)
	//
	//v4 := Vertex{}
	//fmt.Printf("%T, %v\n", v4, v4)
	//
	//var v5 Vertex
	//fmt.Printf("%T, %v\n", v5, v5)
	//
	//v6 := new(Vertex)
	//fmt.Printf("%T, %v\n", v6, v6)
	//
	//v7 := &Vertex{}
	//fmt.Printf("%T, %v\n", v7, v7)
	//
	//v21 := Vertex{1, 2, "test"}
	//changeVertex(v21)
	//fmt.Println(v21)
	//
	//v22 := &Vertex{1, 2, "test"}
	//changeVertex2(v22)
	//fmt.Println(*v22, v22)
	//
	//v23 := Vertex{1, 2, "test"}
	//changeVertex2(&v23)
	//fmt.Println(v23, &v23)

	///演習3
	////Q1
	//var i1 int = 10
	//var p *int
	//p = &i1
	//var j1 int
	//j1 = *p
	//fmt.Println(j1) // 10
	//
	////Q2
	//var i2 int = 100
	//var j2 int = 200
	//var p1 *int
	//var p2 *int
	//p1 = &i2
	//p2 = &j2
	//i2 = *p1 + *p2  //100 + 200
	//p2 = p1         // &i2
	//j2 = *p2 + i2   // i2 + i2
	//fmt.Println(j2) // 600

	/// Pointer Receiver & Value Receiver
	//v := Vertex{x: 3, y: 4} //labelなしの時は、全てのfieldに対して値の設定が必要
	//fmt.Println(Area(v))
	//v.Scale(10)
	//fmt.Println(v.Area())

	/// Constructor
	//v := New(3, 4) // GoLangでは、ConstructorをNewで定義する御作法があるので、それに倣う
	//v.Scale(10)
	//fmt.Printf("%T, %v", v, v.Area())

	/// Embedded
	//v := New3D(3, 4, 5)
	//v.Scale3D(10)
	//fmt.Printf("%T, %v", v, v.Area3D())

	/// non struct method
	//myInt := MyInt(10)
	//fmt.Println(myInt.Double())

	/// Interface & Duck typing
	//var mike Human = &Person{"Mike"}
	//var x Human = &Person{"x"}
	////var dog Dog = Dog{"dog"}
	//DriveCar(mike)
	//DriveCar(x)
	////DriveCar(dog) // error Say()が定義されていないので、Humanに適合していない

	/// Type assertion & switch type
	//do1(10)
	//do1("Mike")
	//do1(true)
	//do2(10)
	//do2("Mike")
	//do2(true)

	/// Stringer
	//p := Person{"Mike", 22}
	//fmt.Println(p)
	//fmt.Println(p.String())

	/// Custom Error
	//if err := myFunc(); err != nil {
	//	fmt.Println(err)
	//}

	/// 演習
	//// Q1
	//v := Vertex{3, 4, ""}
	//fmt.Println(v.Plus())
	//// Q2
	//fmt.Println(v)

	/// Goroutine
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go goroutine("world", &wg)
	//normal("hello")
	//wg.Wait()

	/// Channel
	//s := []int{1, 2, 3, 4, 5}
	//c := make(chan int)
	//go goroutine1(s, c)
	//go goroutine1(s, c)
	//fmt.Println("before x := <- c")
	//x := <-c
	//fmt.Println("after x := <- c")
	//fmt.Println("before y := <- c")
	//y := <-c
	//fmt.Println("after y := <- c")
	//fmt.Println(x)
	//fmt.Println(y)

	/// Buffered Channels
	//ch := make(chan int, 2)
	//ch <- 100
	//fmt.Printf("ch count: %d, address: %v\n", len(ch), ch)
	//ch <- 200
	//fmt.Printf("ch count: %d, address: %v\n", len(ch), ch)
	//// closeしないとrangeで取得できない
	//close(ch)
	//
	//for c := range ch {
	//	fmt.Println(c)
	//}

	/// Range & Close
	//s := []int{1, 2, 3, 4, 5}
	//c := make(chan int, 5)
	//go goroutine1close(s, c)
	//fmt.Println("before for i := range c")
	//for i := range c {
	//	fmt.Println("before println")
	//	fmt.Println(i)
	//	fmt.Println(len(c))
	//	fmt.Println("after println")
	//}
	//fmt.Println("after for i := range c")

	/// Producer & Consumer
	//var wg sync.WaitGroup
	//ch := make(chan int)
	//
	//// Producer
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go producer(ch, i)
	//}
	//
	//// Consumer
	//go consumer(ch, &wg)
	//wg.Wait()
	//close(ch)

	/// fan-out fan-in
	//first := make(chan int)
	//second := make(chan int)
	//third := make(chan int)
	//
	//go producer1(first)
	//go multi2(first, second)
	//go multi4(second, third)
	//
	//for i := range third {
	//	fmt.Println(i)
	//}

	/// channel select
	//c1 := make(chan string)
	//c2 := make(chan string)
	//go goroutine21(c1)
	//go goroutine22(c2)
	//
	//for {
	//	select {
	//	case msg1 := <-c1:
	//		fmt.Println(msg1)
	//	case msg2 := <-c2:
	//		fmt.Println(msg2)
	//	}
	//}

	/// Default select & for break
	//	tick := time.Tick(200 * time.Millisecond)
	//	boom := time.After(1000 * time.Millisecond)
	//loop:
	//	for {
	//		select {
	//		case <-tick:
	//			fmt.Println("tick.")
	//		case <-boom:
	//			fmt.Println("BOOM!")
	//			// returnでfuncを抜けても良し、break for labelでfor文を抜けても良し
	//			break loop
	//			//return
	//		default:
	//			fmt.Println("    .")
	//			time.Sleep(100 * time.Millisecond)
	//		}
	//	}
	//	fmt.Println("end.")

	/// Mutex
	////同一の変数を複数のgoroutineから操作する時に競合するので、競合しないようにMutexでlock/unlockしてやる
	//c := Counter{v: make(map[string]int)}
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c.Int("key")
	//	}
	//}()
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c.Int("key")
	//	}
	//}()
	//time.Sleep(1 * time.Second)
	//fmt.Println(c.Value("key"))

	///演習
	//Q1
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string)
	go goroutine31(words, c)
	for w := range c {
		fmt.Println(w)
	}
}

func log1(slice []int) {
	fmt.Printf("len=%d, cap=%d, value=%v\n", len(slice), cap(slice), slice)
}

func calc(x, y int) (int, int) {
	fmt.Println("calc function")
	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	return x + y, x - y
}
func cal(price, item int) (result int) {
	fmt.Printf("before result %d\n", result)
	result = price * item
	fmt.Printf("after result %d\n", result)
	return
}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x += 1
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func getOsName() string {
	return "mac"
}

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // logに書き込むものを指定(書き込まれる順番は固定)
	log.SetOutput(multiLogFile)
}

func one(x *int) {
	*x = 1
}

type Vertex struct {
	x, y int
	s    string
}

func changeVertex(v Vertex) {
	v.x = 1000
}

func changeVertex2(v *Vertex) {
	v.x = 1000
}

func Area(v Vertex) int {
	return v.x * v.y
}

func (v Vertex) Area() int {
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func New(x, y int) *Vertex {
	return &Vertex{x: x, y: y}
}

type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New3D(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x: x, y: y}, z}
}

type MyInt int

func (i MyInt) Double() int {
	fmt.Printf("%T, %v\n", i, i)
	return int(i * 2)
}

type Human interface {
	Say() string
}

type Person struct {
	Name string
	Age  int
}

type Dog struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func do1(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

func do2(i interface{}) {
	if v, ok := i.(int); ok {
		fmt.Println(v * 3)
	} else if s, ok := i.(string); ok {
		fmt.Println(s + "!!")
	} else {
		fmt.Printf("I don't know %T too\n", i)
	}
}

func (p Person) String() string {
	return "My name is " + p.Name
	//return fmt.Sprintf("My name is %v", p.Name)
}

func myFunc() error {
	// Something wrong
	ok := false
	if ok {
		return nil
	}
	err := &UserNotFound{"Mike"}
	return err
}

type UserNotFound struct {
	UserName string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.UserName)
}

func (v Vertex) Plus() int {
	return v.x + v.y
}

func (v Vertex) String() string {
	return fmt.Sprintf("X is %v! Y is %v!", v.x, v.y)
}

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		//time.Sleep(200 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(200 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutine1(s []int, c chan int) {
	sum := 0
	for i, v := range s {
		sum += v
		fmt.Printf("count: %v\n", i)
		time.Sleep(1 * time.Second)
	}
	c <- sum
}

func goroutine1close(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		fmt.Println("before c <- sum")
		c <- sum
		fmt.Println("after c <- sum")
		time.Sleep(1 * time.Second)
	}
	close(c)
}

func producer(ch chan int, i int) {
	ch <- i
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process index: ", i)
		wg.Done()
	}
}

func producer1(first chan<- int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

// <-で向き先が固定できるので宣言に含めるのが良さそう
// <-chan output only, chan<- input only
func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func multi4(second <-chan int, third chan<- int) {
	defer close(third)
	for i := range second {
		third <- i * 4
	}
}

func goroutine21(ch chan<- string) {
	for {
		ch <- "packet from 1"
		time.Sleep(3500 * time.Millisecond)
	}
}

func goroutine22(ch chan<- string) {
	for {
		ch <- "packet from 2"
		time.Sleep(1 * time.Second)
	}
}

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Int(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func goroutine31(s []string, c chan<- string) {
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}
