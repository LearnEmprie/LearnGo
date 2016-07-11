package mystring

import (
	"strings"
	"fmt"
	"strconv"
)

func List()  {


	fmt.Println(strings.HasPrefix("pssss","ps"))
	fmt.Println(strings.HasSuffix("pssss","s"))
	//直接能索引字符串
	fmt.Println(string("ddd"[0]))
	fmt.Println(strings.Contains("pppxxxdddll","xd"))

	fmt.Println(strings.Index("lksdksl","s"))
	fmt.Println(strings.LastIndex("lksdksl","s"))

	//中文是三个字节
	s:="Go编程"
	fmt.Println(len(s))
	//rune 能操作 任何字符 byte 不支持中文的操作
	fmt.Println(len(string(rune('编'))))

	dog := "1狗"
	//rune 就是4字节 能支持任何文字
	pig := []byte(dog)
	cat := []rune(dog)
	fmt.Println(string(pig[0]),string(cat[1]))

	//混合的怎么用??
	fmt.Println(strings.IndexAny("我lllpl","我"))
	//替换字符串
	fmt.Println(strings.Replace("sfdsfsd", "s" , "+" , -1))

	var str string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))

	newS := strings.Repeat("Dog!", 3)
	fmt.Printf("The new repeated string is: %s\n", newS)


	var orig string = "Hey, how are you George?"
	var lower string
	var upper string

	fmt.Printf("The original string is: %s\n", orig)
	lower = strings.ToLower(orig)
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The uppercase string is: %s\n", upper)



	fmt.Println(strings.TrimSpace("   xxx   "))
	fmt.Println(strings.Trim("xxxpppxsxxx","xx"))
	fmt.Println(strings.TrimLeft("xxxpppxsxxx","xx"))
	fmt.Println(strings.TrimRight("xxxpppxsxxx","xx"))


	stri := "The quick brown fox jumps over the lazy dog"
	sl  := strings.Fields(stri)
	fmt.Printf("Splitted in slice: %v\n", sl)

	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)

	str3 := strings.Join(sl2,";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)

	fmt.Println(strconv.Itoa(111))
	fmt.Println(strconv.Atoi("111"))

	fmt.Println(strconv.ParseInt("11",10,64))
	fmt.Println(strconv.FormatInt(11,10))

}