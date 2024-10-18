package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func task1(num int) int {
	// "1. **Проверка на простоту**  \n",
	// "Напишите функцию, которая проверяет, является ли переданное число простым.
	// Ваша программа должна использовать циклы для проверки делителей, и если число
	// не является простым, выводить первый найденный делитель.\n",

	for j := 2; j < num; j++ {
		if num%j == 0 {
			return j // делитель не простого числа
		}
	}

	return num // простое число
}

func task2(n1, n2 int) int {
	// 		"2. **Наибольший общий делитель (НОД)**  \n",
	//      "   Напишите программу для нахождения наибольшего общего делителя (НОД)
	// двух чисел с использованием алгоритма Евклида. Используйте цикл `for` для
	// вычислений.\n",
	if n1 < n2 {
		n1, n2 = n2, n1
	}
	for n2 != 0 {
		t := n2
		n2 = n1 % n2
		n1 = t
	}

	return n1
}

func task3(nums []int) {
	// "3. **Сортировка пузырьком**  \n",
	//
	//	"   Реализуйте сортировку пузырьком для списка целых чисел. Программа должна выполнять сортировку на месте и выводить каждый шаг изменения массива.\n",
	//	"\n",

	for {
		ready := true
		for i := 0; i < len(nums)-1; i++ {

			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				ready = false
				fmt.Println(nums)
			}

		}
		if ready {
			break
		}
	}
	fmt.Println(nums)
}

func task4() {
	// "4. **Таблица умножения в формате матрицы**  \n",
	//	"   Напишите программу, которая выводит таблицу умножения
	// в формате матрицы 10x10. Используйте циклы для генерации строк и столбцов.\n",
	for i := 1; i <= 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, " ")
		}
		fmt.Println()
	}
}

func task6(num int) {
	// "6. **Обратные числа**  \n",
	// "   Напишите программу, которая принимает целое число и выводит его в обратном
	// порядке. Например, для числа 12345 программа должна вывести 54321.
	// Используйте цикл для обработки цифр числа.\n"
	var val string = strconv.Itoa(num)
	//fmt.Println(val[0:1])

	for i := len(val); i > 0; i-- {
		fmt.Print(val[i-1 : i])
	}
	fmt.Println()
}

func task7(level int) {
	// "7. **Треугольник Паскаля**  \n",
	// "   Напишите программу, которая выводит треугольник Паскаля до заданного уровня.
	// Для этого используйте цикл и массивы для хранения предыдущих значений строки треугольника.\n",
	arr := make([][]int, level)

	for i := 0; i < level; i++ {
		arr[i] = make([]int, i+1)
		arr[i][0] = 1
		arr[i][i] = 1

		for j := 1; j < i; j++ {
			arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
		}
	}

	for i := 0; i < level; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}

}

func task8(num int) bool {
	// "8. **Число палиндром**  \n",
	// "   Напишите программу, которая проверяет, является ли число палиндромом (одинаково читается слева направо и справа налево). Не используйте строки для решения этой задачи — работайте только с числами.\n",
	if num < 0 {
		return false
	}

	original := num
	reversed := 0

	for num > 0 {
		rem := num % 10
		reversed = reversed*10 + rem
		num /= 10
	}

	return original == reversed
}

func task9(arr []int) (int, int) {
	// **Нахождение максимума и минимума в массиве**  \n",
	//     "   Напишите функцию, которая принимает массив целых чисел и возвращает одновременно максимальный и минимальный элемент с использованием одного прохода по массиву.\n",
	min := arr[0]
	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		} else if arr[i] <= min {
			min = arr[i]
		}
	}

	return min, max
}

func task10() {
	// "10. **Игра \"Угадай число\"**  \n",
	// "   Напишите программу, которая загадывает случайное число от 1 до 100, а пользователь пытается его угадать. Программа должна давать подсказки \"больше\" или \"меньше\" после каждой попытки. Реализуйте ограничение на количество попыток.\n",
	val := rand.Intn(100) // число
	//fmt.Println(val)
	try := 5 // попытки
	fmt.Printf("Отгадайте число, у вас %d попыток\n", try)
	var num int
	for val != num {
		if try == 0 {
			fmt.Printf("Гаме овер. Угадываемое число - %d \n", val)
			return
		}
		fmt.Scanln(&num)
		if num == val {
			fmt.Println("Верно!")
			return
		}
		diff := num - val
		if diff > 0 {
			fmt.Print("Меньше. ")
		} else {
			fmt.Print("Больше. ")

		}

		try--
		fmt.Printf("Осталось попыток: %d \n", try)

	}
}

func task11(n1, n2 int) {
	// "11. **Числа Армстронга**  \n",
	// "   Напишите программу, которая проверяет, является ли число числом Армстронга (число равно сумме своих цифр, возведённых в степень, равную количеству цифр числа). Например, 153 = 1³ + 5³ + 3³.\n",

	for i := n1; i < n2; i++ {
		if isArmstrong(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func isArmstrong(num int) bool {
	numsInNum := int(math.Floor(math.Log10(float64(num)))) + 1
	sum := 0
	i := num
	for i > 0 {
		n := i % 10
		sum += int(math.Pow(float64(n), float64(numsInNum)))
		i /= 10
	}
	return sum == num
}

func task12(str string) {
	// "12. **Подсчет слов в строке**  \n",
	//     "   Напишите программу, которая принимает строку и выводит количество уникальных слов в ней. Используйте `map` для хранения слов и их количества.\n",
	//     "\n",
	words := make(map[string]int)

	val := strings.Fields(str)

	for _, el := range val {
		words[el]++
	}
	//fmt.Println(words)
	for wrd, count := range words {
		if count == 1 {
			fmt.Println(wrd)
		}
	}
}

func task14(num int) int {
	// "14. **Цифровой корень числа**  \n",
	// "   Напишите программу, которая вычисляет цифровой корень числа. Цифровой корень — это рекурсивная сумма цифр числа, пока не останется только одна цифра. Например, цифровой корень числа 9875 равен 2, потому что 9+8+7+5=29 → 2+9=11 → 1+1=2.\n",
	res := 1 + (num-1)%9
	fmt.Println(res)
	return res
}

func task15(num int) string {
	// "15. **Римские цифры**  \n",
	// "   Напишите функцию, которая преобразует арабское число (например, 1994) в римское (например, \"MCMXCIV\"). Программа должна использовать циклы и условные операторы для создания римской записи.\n"

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""

	for i := 0; num > 0; i++ {

		for num >= values[i] {
			res += romans[i]
			num -= values[i]
		}
	}

	fmt.Println(res)
	return res
}

func main() {

}
