package main
import ("fmt")

// arrayni key kalit so'zi yordamida yozish
// var array_name = [length]datatype{values} agar length ma'lum bo'lmasa ... dan foydalaniladi

// := belgi bilan yozish
// array_name := [length]datatype{values}
func main() {
	var arr1 = []int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	var cars = [7]string{"Volvo", "BMW", "Ford", "Mazda"}
	cars[0] = "Volvo updated"
	fmt.Println(cars)

	fmt.Println(len(arr1))
	fmt.Println(arr2)
}
