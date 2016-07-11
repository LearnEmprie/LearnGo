package mystruct

import (
	"math"
	"fmt"
	"strconv"
	"runtime"
)

//对象的字段（属性）不应该由 2 个或 2 个以上的不同线程在同一时间去改变。如果在程序发生这种情况，为了安全并发访问，可以使用包 sync（参考第 9.3 节）中的方法。
//一个struct包含的两个类型中有同一个名字的函数,会冲突,但儿子struct有同名的就不冲突了
//当定义了一个有很多方法的类型时，十之八九你会使用 String() 方法来定制类型的字符串形式的输出，换句话说：一种可阅读性和打印性的输出
//类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
//类型 T 的可调用方法集包含接受者为 T 的所有方法
//类型 T 的可调用方法集不包含接受者为 *T 的方法

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (NP *NamedPoint) Abs() float64 {
	return NP.Point.Abs()*100
}

type NamedPoint struct {
	Point
	name string
}

func List(){
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs(),n.Point.Abs())

	C := CameraPhone{}

	C.Call()
	C.TakeAPicture()


	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()

	two1 := &TwoInts{} // 当你定义指针和实例的时候
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)

	var dog [100000]TwoInts

	for i := 0 ; i < 100000 ; i ++ {
		dog[i].b = 1000000
		dog[i].a = 10000322
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)

	for i := 0 ; i < 100000 ; i ++ {
		dog[i].b = 1000000
		dog[i].a = 10000322

	}
	fmt.Println(dog[99999])

	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)

	var i Simpler = new(Simple)
	i.Set(1233)
	fmt.Println(i.Get())

	//类型断言
	if v , ok := i.(*Simple) ; ok {
		fmt.Println(v.Get())
	}

	switch t := i.(type) {
	case *Simple:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

	//是否实现了接口
	if k , ok := i.(Simpler) ; ok {
		fmt.Println(k.Get())
	}

}

type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}
//多重继承不能出现同一个 函数名
type CameraPhone struct {
	Camera
	Phone
}

type Pig struct  {}

func (Pig)Magic()  {
	fmt.Println("pig magic")
}

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
	Pig
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

type TwoInts struct {
	a int
	b int
}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

/* method to determine the value of a stock position */
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

/* method to determine the value of a car */
func (c car) getValue() float32 {
	return c.price
}

/* contract that defines different things that have value */
type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

type Simpler interface {
	Get() int
	Set(i int)
}

type specialString string

var  whatIsThis specialString = "hello"

func TypeSwitch() {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type", v)
		case int:
			fmt.Printf("any %v is an int type", v)
		case float32:
			fmt.Printf("any %v is a float32 type", v)
		case string:
			fmt.Printf("any %v is a string type", v)
		case specialString:
			fmt.Printf("any %v is a special String!", v)
		default:
			fmt.Println("unknown type!")
		}
	}
	testFunc(whatIsThis)
}


type Simple struct  {
	num int
}

func (self Simple)Get() int {
	return self.num
}

func (self *Simple)Set(i int){
	self.num = i
}