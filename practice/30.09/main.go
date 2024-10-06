package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func task11() {
	//		Сумма цифр числа:
	//	   Напишите программу, которая принимает целое число и вычисляет сумму его цифр.
	//
	// Вход: `1234`
	// Выход: `10` (1 + 2 + 3 + 4)

	fmt.Println("Программа, которая принимает целое число и вычисляет сумму его цифр (только 4 цифры!!!)):")
	var nums int
	var sum int
	fmt.Scanln(&nums)
	var n1 int = nums / 1000
	var n2 int = (nums / 100) % 10
	var n3 int = (nums / 10) % 10
	var n4 int = nums % 10
	sum = n1 + n2 + n3 + n4
	fmt.Println(sum)
}

func task12() {
	//  2. Преобразование температуры:
	//     Напишите программу, которая преобразует температуру из градусов Цельсия в Фаренгейты и обратно.
	//     Вход: `25 (Celsius)`
	//     Выход: `77 (Fahrenheit)`

	fmt.Println("Преобразование температуры из градусов Цельсия в Фаренгейты. Введите температуру в цельсиях:")
	var Celsius float32
	fmt.Scanln(&Celsius)
	var Fahrenheit float32 = Celsius*1.8 + 32
	fmt.Println(Fahrenheit, "(Fahrenheit)")
}

func task13() {
	//		Удвоение каждого элемента массива:
	//		Напишите программу, которая принимает массив чисел и возвращает новый массив, где каждое число удвоено.
	//	 Вход: `[1, 2, 3, 4]`
	//	 Выход: `[2, 4, 6, 8]`
	fmt.Println("Программа, которая принимает массив чисел и возвращает новый массив, где каждое число удвоено (только 4 элемента):")
	var el1, el2, el3, el4 string
	fmt.Scanln(&el1, &el2, &el3, &el4)

	val1, _ := strconv.Atoi(el1[1 : len(el1)-1])
	val2, _ := strconv.Atoi(el2[0 : len(el2)-1])
	val3, _ := strconv.Atoi(el3[0 : len(el3)-1])
	val4, _ := strconv.Atoi((el4[0 : len(el4)-1]))

	fmt.Printf("[%d, %d, %d, %d]\n", val1*2, val2*2, val3*2, val4*2)
}

func task14() {
	// Напишите программу, которая принимает несколько строк и объединяет их в одну строку через пробел.

	fmt.Println(" напишите две строки в виде массива")
	var wrd1, wrd2 string
	fmt.Scanln(&wrd1, &wrd2)

	w1 := wrd1[2 : len(wrd1)-2]
	w2 := wrd2[1 : len(wrd2)-2]
	fmt.Printf("%s %s \n", w1, w2)
}

func task15() {
	// Напишите программу, которая вычисляет расстояние между двумя точками в 2D пространстве.
	// (x1=1, y1=1), (x2=4, y2=5)

	fmt.Println("Вычисление расстояния между двумя точками в 2D пространстве.")
	var x1, y1, x2, y2 string
	fmt.Scanln(&x1, &y1, &x2, &y2)

	valx1, _ := strconv.Atoi(x1[4 : len(x1)-1])
	valx2, _ := strconv.Atoi(x2[4 : len(x1)-1])
	valy1, _ := strconv.Atoi(y1[3 : len(y1)-2])
	valy2, _ := strconv.Atoi(y2[3 : len(y1)-2])

	res := math.Pow(math.Pow(float64(valx2-valx1), 2)+math.Pow(float64(valy2-valy1), 2), 1/2.0)
	fmt.Println(res)
}

func task21() {
	//	Напишите программу, которая проверяет, является ли введенное число четным или нечетным.
	//
	// Вход: `4`
	// Выход: `Четное`
	fmt.Println("Введите число для проверки на четность:")
	var num int
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("Четное")
		return
	}
	fmt.Println("Нечетное")
}

func task22() {
	// Напишите программу, которая проверяет, является ли введенный год високосным.
	// Вход: `2020`
	// Выход: `Високосный`

	fmt.Println("Введите год для проверки является ли введенный год високосным:")
	var year int
	fmt.Scan(&year)

	if year%4 == 0 {
		fmt.Println("Високосный")
		return
	}
	fmt.Println("Не високосный")
}

func task23() {
	// Напишите программу, которая принимает три числа и выводит наибольшее из них.
	// Вход: `4, 9, 7`
	// Выход: `9`
	fmt.Println("программу, которая принимает три числа и выводит наибольшее из них:")
	var num1, num2, num3 string
	fmt.Scanln(&num1, &num2, &num3)

	n1, _ := strconv.Atoi(num1[0 : len(num1)-1])
	n2, _ := strconv.Atoi(num2[0 : len(num1)-1])
	n3, _ := strconv.Atoi(num3)

	if n1 > n2 {
		if n1 > n3 {
			fmt.Println(n1)
			return
		}
		fmt.Println(n3)
	} else {
		if n2 > n3 {
			fmt.Println(n2)
			return
		}
		fmt.Println(n3)
	}
}

func task24() {
	// Напишите программу, которая принимает возраст человека и выводит, к какой возрастной группе он относится (ребенок, подросток, взрослый, пожилой. В комментариях указать возрастные рамки).
	//ребенок - <14, подросток - [14;17], взрослый - [18:65], пожилой - 65+
	fmt.Println("Введите возраст:")
	var age int
	fmt.Scan(&age)
	if age < 14 {
		fmt.Println("Ребенок")
	} else if age >= 14 && age <= 17 {
		fmt.Println("Подросток")
	} else if age >= 18 && age < 65 {
		fmt.Println("Взрослый")
	} else if age >= 65 {
		fmt.Println("Пожилой")
	}
}

func task25() {
	//		Проверка делимости на 3 и 5:
	//		Напишите программу, которая проверяет, делится ли число одновременно на 3 и 5.
	//	 Вход: `15`
	//	 Выход: `Делится`
	fmt.Println("Введите число для проверки делимости на 3 и 5:")
	var num int
	fmt.Scanln(&num)
	if num%5 == 0 && num%3 == 0 {
		fmt.Println("Делится")
		return
	}
	fmt.Println("Не делится")
}

func task31() {
	// 	Факториал числа:
	//    Напишите программу, которая вычисляет факториал числа.
	// Вход: `5`
	// Выход: `120` (5! = 5 × 4 × 3 × 2 × 1)
	fmt.Println("Введите число для вычисления факториала:")
	var num int
	fmt.Scan(&num)
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	fmt.Println(res)
}

func task32() {
	//		Числа Фибоначчи:
	//		Напишите программу, которая выводит первые `n` чисел Фибоначчи.
	//	 Вход: `n = 7`
	//	 Выход: `0, 1, 1, 2, 3, 5, 8`
	fmt.Println("программу, которая выводит первые n чисел Фибоначчи. Введите n:")
	var n int
	fmt.Scan(&n)
	var res = make([]int, n)
	res[0] = 0
	res[1] = 1
	for i := 2; i < n; i++ {
		res[i] = res[i-1] + res[i-2]
	}

	for i, el := range res {
		fmt.Printf("%d", el)
		if i < n-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

func task33(arr []int) []int {
	//		Реверс массива:
	//	   Напишите программу, которая переворачивает массив чисел.
	//
	// Вход: `[1, 2, 3, 4, 5]`
	// Выход: `[5, 4, 3, 2, 1]`

	var res = make([]int, len(arr))
	j := 0
	for i := len(arr) - 1; i >= 0; i-- {
		res[j] = arr[i]
		j++
	}
	fmt.Println(res)
	return res
}

func task34() {
	//     Поиск простых чисел:
	//     Напишите программу, которая выводит все простые числа до заданного числа.
	//
	// Вход: `n = 20`
	// Выход: `2, 3, 5, 7, 11, 13, 17, 19`

	var n int
	fmt.Scanln(&n)

	for i := 0; i <= n; i++ {
		var num = i
		if big.NewInt(int64(num)).ProbablyPrime(0) {
			fmt.Printf("%d, ", i)
		}
	}
	fmt.Println()
}

func task35(arr []int) {
	//	Напишите программу, которая вычисляет сумму всех чисел в массиве.
	//
	// Вход: `[1, 2, 3, 4, 5]`
	// Выход: `15`
	var res int
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	fmt.Println(res)
}
func main() {
}
