package main

import (
	"fmt"
	"time"
)

type User struct {
	Name string
	Age  uint
}

func PrintInfo(pr User) {
	fmt.Println("Name: ", pr.Name)
	fmt.Println("Age: ", pr.Age)
}

func add(x, y int) int {
	return x + y
}

func main() {
	var name string = "tuschy"                       // string
	age := 22                                        //int
	var o uint = 1                                   //not -1
	const pi float32 = 3.14                          //const and float
	pi1 := 3.1428                                    // declar :=
	var isStudent bool = true                        //boolean
	arr := [2]int{1, 2}                              //array
	slice := []string{"Tuschy", "Miko"}              //slice
	slice1 := append(slice, "Yumi", "Nriny")         //append
	map1 := map[int]string{1: "Nrinee", 2: "Tuschy"} //Map , map[keytype]valuetype{key: value}

	fmt.Println("string: ", name)
	fmt.Println("Int: ", age)
	fmt.Println("uint not minus: ", o)
	fmt.Println("Const and float : ", pi)
	fmt.Println("Declar ':=' and float : ", pi1)
	fmt.Println("boolean: ", isStudent)
	fmt.Println("Array: ", arr)
	fmt.Println("slice: ", slice)
	fmt.Println("slice append: ", slice1)
	fmt.Println("Map Show all: ", map1)   // show all
	fmt.Println("Map getByID: ", map1[1]) // show by id

	x, y := 10, 5
	fmt.Printf("function ADD %d + %d = %d \n", x, y, add(x, y))

	fmt.Println("-----Struct and Function-------")

	var user User

	user.Name = "Tuschykk"
	user.Age = 22

	PrintInfo(user)

	fmt.Println("------Goroutines and Channels------")

	// fmt.Println("ซื้อแว่น ที่ เซเว่น")
	// fmt.Println("ซื้อนาฬิกา ที่ เซ็นทรัล")
	// go fmt.Println("ซื้อผลไม้ ที่ สยามพารากอน")
	// fmt.Println("ซื้อรถ ที่ ศูนย์ Toyota")
	// time.Sleep(1 * time.Second)
	// fmt.Println("------------")

	// function read more at file goroutins.go

	start := time.Now()     // เริ่มจับเวลาในการ Run
	ch := make(chan string) // สร้างท่อ Channel เอาไว้ส่งข้อมูล

	go BuyGlassesAtSevenEleven() //put Go
	go BuyWatchAtCentral()       // put Go
	go SendToMisterA(ch)         // ส่งท่อ ch เข้าไปใน Function

	BuyFruitAtSiamParagon()
	BuyCarAtToyota()
	messageFromMisterB := <-ch // ค่าจากท่อ Channel จะออกตรงนี้ รอจนกว่า SendToMisterA() จะทำเสร็จ

	if messageFromMisterB == "กำลังส่งของให้ นาย A" {
		fmt.Println("นาย A ได้รับของแล้ว")
		fmt.Println("ใช้เวลาในการ Run ทั้งสิ้น : ", time.Since(start), " วินาที")
	}

}
