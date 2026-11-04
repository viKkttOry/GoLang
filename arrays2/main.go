package main

import "fmt"

func main() {
	// 1
	fmt.Println("Задание 1")

	// 1 часть задачи.

	// декларация массивов.
	var numbers1 [5]int
	var numbers2 [5]int
	var numbers3 [5]int

	// декларация и инизиализация элементов 1 массива.
	for i := 0; i < len(numbers1); i++ {
		fmt.Printf("Введите элемент №%d массива:\n ", i+1)
		fmt.Scan(&numbers1[i])
	}

	// декларация и инизиализация элементов 2 массива.
	for j := 0; j < len(numbers2); j++ {
		fmt.Printf("Введите элемент №%d массива:\n ", j+1)
		fmt.Scan(&numbers2[j])
	}

	// декларация и инизиализация элементов 3 массива.
	for k := 0; k < len(numbers3); k++ {
		fmt.Printf("Введите элемент №%d массива:\n ", k+1)
		fmt.Scan(&numbers3[k])
	}

	// декларация и инизиализация общего двумерного массива.
	resultNumbers := [3][5]int{
		numbers1,
		numbers2,
		numbers3,
	}

	// декларация и инициализация переменных, отвечающих за минимальные и максимальные значения.
	minX := 0
	minY := 0

	maxX := 0
	maxY := 0

	// 2 часть задачи.

	// нахождение индексов и их вывод.
	for index_y, value := range resultNumbers {
		for index_x, val := range value {
			if val < resultNumbers[minY][minX] {
				minY = index_y
				minX = index_x
			}
			if val > resultNumbers[maxY][maxX] {
				maxY = index_y
				maxX = index_x
			}
		}
	}

	fmt.Printf("Row: %d, col: %d\n", maxY, maxX)

	// 2
	fmt.Println("Задание 2")

	// 1 часть.

	// Декларация и инициализация элементов двумерного массива.
	point := "."
	star := "*"
	var snowflake [11][11]string

	// 2 часть.

	// Заполнение  снежинки.
	for index_x, value := range snowflake {
		for index_y := range value {
			if index_x == index_y {
				snowflake[index_x][index_y] = star
			} else if index_x == len(snowflake)/2 {
				snowflake[index_x][index_y] = star
			} else if index_y == len(snowflake[index_x])/2 {
				snowflake[index_x][index_y] = star
			} else if index_x+index_y == (len(snowflake) - 1) {
				snowflake[index_x][index_y] = star
			} else {
				snowflake[index_x][index_y] = point
			}
		}
	}

	// 3 часть.

	// Вывод снежинки.
	for i := 0; i < len(snowflake); i++ {
		for j := 0; j < len(snowflake[i]); j++ {
			fmt.Print(snowflake[i][j], "")
		}
		fmt.Println()
	}

	// 3.
	fmt.Println("Задачние 3")

	// 1 Часть.

	// Декларация и инициализация переменных и двумерного массива.
	dot := "."
	starr := "*"
	var chessDesk [8][8]string

	// 2 Часть.

	// Заполнение шахматной доски.
	for index_x, value := range chessDesk {
		for index_y := range value {
			if (index_x+index_y)%2 == 0 {
				chessDesk[index_x][index_y] = dot
			} else {
				chessDesk[index_x][index_y] = starr
			}
		}
	}

	// Часть 3. вывод шахматной доски.
	for i := 0; i < len(chessDesk); i++ {
		for j := 0; j < len(chessDesk[i]); j++ {
			fmt.Print(chessDesk[i][j], "")
		}
		fmt.Println()
	}

	// 4
	fmt.Println("Задание 4")

	// 1 часть.

	// Декларация и инициализация массивов и двумерного массива.

	var numbersArr [4][4]int

	counter := 1
	for i, row := range numbersArr {
		for j := range row {
			fmt.Printf("Введите число:%d\n", counter)
			counter++
			fmt.Scan(&numbersArr[i][j])
		}
	}

	// 2 Часть.

	// Вывод массива
	var i int
	fmt.Println("Введите первое число: ")
	fmt.Scan(&i)

	var j int
	fmt.Println("Введите второе число: ")
	fmt.Scan(&j)

	var arr [len(numbersArr[0])]int
	arr = numbersArr[i-1]
	numbersArr[i-1] = numbersArr[j-1]
	numbersArr[j-1] = arr
	fmt.Println(numbersArr)

	// 5
	fmt.Println("Задание 5")

	var numbersAr [4][4]int

	counterr := 1
	for i, row := range numbersAr {
		for j := range row {
			fmt.Printf("Введите число:%d\n", counterr)
			counterr++
			fmt.Scan(&numbersAr[i][j])
		}
	}

	var ii int

	fmt.Println("Введите первое число: ")
	fmt.Scan(&ii)

	var jj int

	fmt.Println("Введите второе число: ")
	fmt.Scan(&jj)

	for i := range numbersAr {
		numbersAr[i][ii-1], numbersAr[i][jj-1] = numbersAr[i][jj-1], numbersAr[i][ii-1]
	}

	fmt.Println(numbersAr)
}
