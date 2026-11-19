package main

import (
	"fmt"
	// "sort"
	"strings"
)

func main() {
	// // 1
	// fmt.Println("Задание 1")

	// var numbers []int

	// for {
	// 	var input int
	// 	fmt.Println("Введите число (или 'stop' для завершения цикла):")

	// 	_, err := fmt.Scan(&input)
	// 	if err != nil {
	// 		break
	// 	} else {
	// 		numbers = append(numbers, input)
	// 	}
	// }

	// sort.Ints(numbers)

	// counter := 0
	// for _, value := range numbers {
	// 	counter += value
	// }

	// fmt.Println("Сумма всех чисел:", counter)
	// var q string
	// fmt.Scanln(&q)
	// // 2
	// fmt.Println("Задание 2")

	// var values []int
	// for {
	// 	var inpt int
	// 	fmt.Println("Введите число (или 'stop' для завершения):")

	// 	_, err := fmt.Scan(&inpt)
	// 	if err != nil {
	// 		break
	// 	} else {
	// 		values = append(values, inpt)
	// 	}
	// }
	// sort.Ints(values)

	// evenNumbers := make([]int, 0)
	// for _, val := range values {
	// 	if val%2 == 0 {
	// 		evenNumbers = append(evenNumbers, val)
	// 	}
	// }
	// sort.Ints(evenNumbers)
	// fmt.Println("Слайс с четными числами:", evenNumbers)
	// fmt.Println("Исходный слайс:", values)

	// var s string
	// fmt.Scanln(&s)
	// // 3
	// fmt.Println("Задание 3")

	// var data []int
	// for {
	// 	var inptt int
	// 	fmt.Println("Введите число (или 'stop' для завершения):")

	// 	_, err := fmt.Scan(&inptt)
	// 	if err != nil {
	// 		break
	// 	} else {
	// 		data = append(data, inptt)
	// 	}
	// }

	// sort.Ints(data)
	// fmt.Println("Слайс до удаления числа с индексом 2:", data)
	// fmt.Println("Число с индексом 2:", data[2])

	// data = append(data[:2], data[3:]...)
	// fmt.Println("Слайс после удаления числа с индексом 2:", data)

	// var w string
	// fmt.Scanln(&w)
	// // 4
	// fmt.Println("Задание 4")

	// var temps []int
	// for {
	// 	var inptts int
	// 	fmt.Println("Введите число (или 'stop' для завершения):")

	// 	_, err := fmt.Scan(&inptts)
	// 	if err != nil {
	// 		break
	// 	} else {
	// 		temps = append(temps, inptts)
	// 	}
	// }
	// fmt.Println("Слайс температур:", temps)

	// minX := 0
	// maxX := 0
	// for index_x, temp := range temps {
	// 	if temp < temps[minX] {
	// 		minX = index_x
	// 	}
	// 	if temp > temps[maxX] {
	// 		maxX = index_x
	// 	}
	// }
	// fmt.Println("Минимальная температура:", temps[minX])
	// fmt.Println("Максимальная температура:", temps[maxX])

	// var e string
	// fmt.Scanln(&e)

	// // 5
	// fmt.Println("Задание 5")

	// var words []string
	// for {
	// 	var inpts string
	// 	fmt.Println("Введите строку (или 'stop' для завершения):")

	// 	fmt.Scanln(&inpts)
	// 	if inpts == strings.ToLower("STOP") {
	// 		break
	// 	} else {
	// 		words = append(words, inpts)
	// 	}
	// }
	// fmt.Println("Слайс слов:", words)

	// reversed := make([]string, len(words))

	// for i := 0; i < len(words); i++ {
	// 	reversed[len(words)-1-i] = words[i]
	// 	words = reversed
	// }

	// fmt.Println(words)

	// var g string
	// fmt.Scanln(&g)
	// // 6
	// fmt.Println("Задание 6")

	// var nums []int
	// for {
	// 	var intps int
	// 	fmt.Println("Введите число (или 'stop' для завершения):")

	// 	_, err := fmt.Scan(&intps)
	// 	if err != nil {
	// 		break
	// 	} else {
	// 		nums = append(nums, intps)
	// 	}
	// }

	// sorted := true
	// for i := 0; i < len(nums)-1; i++ {
	// 	if nums[i] > nums[i+1] {
	// 		sorted = false
	// 		break
	// 	}
	// }

	// fmt.Println(sorted)

	// var h string
	// fmt.Scanln(&h)
	// 7
	fmt.Println("Задание 7")

	var registrationList []string
	var myWishList []string

	for {
		var intpsss string
		fmt.Println("Введите подарок, который хотите вы (или 'stop' для завершения):")

		fmt.Scan(&intpsss)
		if intpsss == strings.ToLower("STOP") {
			break
		} else {
			myWishList = append(myWishList, intpsss)
		}
	}

	for i := 0; i < len(myWishList); i++ {
		fmt.Printf("Подарок, который вы хотите №%d: %s\n", i+1, myWishList[i])
	}

	var friendsWishList []string
	for {
		var inpss string
		fmt.Println("Введите подарок, который хочет ваш друг(или 'stop' для завершения):")

		fmt.Scan(&inpss)
		if inpss == strings.ToLower("STOP") {
			break
		} else {
			friendsWishList = append(friendsWishList, inpss)
		}
	}

	for i := 0; i < len(friendsWishList); i++ {
		fmt.Printf("Подарок, который хочет ваш друг №%d: %s\n", i+1, friendsWishList[i])
	}

	for i := 0; i < len(myWishList); i++ {
		registrationList = append(registrationList, myWishList[i])
	}

	for i := 0; i < len(friendsWishList); i++ {
		registrationList = append(registrationList, friendsWishList[i])
	}

	fmt.Println("Объеденённый виш-лист: ", registrationList)

	var b string
	fmt.Scanln(&b)

	// 8
	fmt.Println("Задание 8")

	var toyList []string

	for {
		var itpsss string
		fmt.Println("Введите название любой игрушки (или 'stop' для завершения):")

		fmt.Scan(&itpsss)
		if itpsss == strings.ToLower("STOP") {
			break
		} else {
			toyList = append(toyList, itpsss)
		}
	}

	fmt.Println("Слайс с игрушками:", toyList)

	testToyList := make([]string, 3)
	copy(testToyList, toyList)

	fmt.Println("Копия первого слайса с игрушками, до замены:", testToyList)

	testToyList[1] = "Boat"

	fmt.Println("Оригинальный слайс с игрушками:", toyList)
	fmt.Println("Копия первого слайса с игрушками, почле замены второго элемента:", testToyList)
}
