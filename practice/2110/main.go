package main

import "fmt"

func task1(base, height float64) float64 {
	// "### 1. **Функция вычисления площади треугольника**\n",
	// "Напишите функцию `triangleArea(base float64, height float64) float64`,
	// которая принимает длину основания и высоту треугольника и возвращает его площадь.\n",
	var area float64 = base * height / 2
	return area
}

func task2(arr []int) []int {
	// "### 2. **Сортировка массива по возрастанию**\n",
	//     "Напишите функцию `sortArray(arr []int) []int`, которая принимает
	// массив целых чисел и возвращает его, отсортированный по возрастанию,
	// используя любой метод сортировки.\n"
	res := make([]int, len(arr))
	copy(res, arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {

			if res[j] > res[j+1] {
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}

	return res
}

func task3(n int) int {
	// "### 3. **Сумма квадратов чётных чисел**\n",
	// "Напишите функцию `sumOfSquares(n int) int`, которая принимает целое число `n`
	// и возвращает сумму квадратов всех чётных чисел от 1 до `n`.
	// Используйте цикл для вычисления.\n",

	var res int
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			continue
		}
		res += i * i
	}
	return res

}

func task4(s string) bool {
	// "### 4. **Проверка палиндрома**\n",
	// "Напишите функцию `isPalindrome(s string) bool`, которая проверяет, является ли
	// строка палиндромом (читается одинаково слева направо и справа налево), и возвращает
	//`true` или `false`.\n",
	for i := 0; i <= len(s)-1; i++ {
		//fmt.Println(string(s[i]))
		if string(s[i]) == string(s[len(s)-i-1]) {
			continue
		} else {
			return false
		}
	}
	return true
}

func task5(n int) bool {
	// "### 5. **Проверка числа на простоту**\n",
	//     "Напишите функцию `isPrime(n int) bool`, которая принимает целое число и возвращает
	// `true`, если число простое, и `false` в противном случае. Для проверки используйте цикл.\n",

	// for i := 2; i <= n; i++ {
	// 	for j:=2
	// }
	for i := 1; i <= 9; i++ {
		if n == i || i == 1 {
			continue
		}
		if n%i == 0 {
			return false
		}
	}
	return true
}

func task6(limit int) []int {
	// "### 6. **Генерация последовательности простых чисел**\n",
	//     "Напишите функцию `generatePrimes(limit int) []int`, которая принимает целое число
	//`limit` и возвращает массив всех простых чисел меньше или равных этому значению.\n",
	res := make([]int, 0)
	res = append(res, 1)

	for i := 2; i <= limit; i++ {
		if task5(i) {
			res = append(res, i)
		}
	}
	return res
}

func task7(n int) string {
	// "### 7. **Перевод числа в двоичную систему**\n",
	// "Напишите функцию `toBinary(n int) string`, которая принимает целое число и
	//  возвращает строку, представляющую число в двоичной системе счисления.\n",
	var res string

	for n > 0 {
		remainder := n % 2
		res = fmt.Sprintf("%d%s", remainder, res)
		n = n / 2
	}

	return res
}

func task8(arr []int) int {
	// "### 8. **Нахождение максимального элемента в массиве**\n",
	// "Напишите функцию `findMax(arr []int) int`, которая принимает массив целых чисел
	// и возвращает максимальное число в массиве.\n",
	var res int = arr[0]

	for _, el := range arr {
		if el > res {
			res = el
		}
	}

	return res
}

func task9(a, b int) int {
	// "### 9. **Функция вычисления НОД (наибольший общий делитель)**\n",
	//     "Напишите функцию `gcd(a int, b int) int`, которая принимает два целых числа и возвращает их наибольший общий делитель, используя алгоритм Евклида.\n",
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func task10(arr []int) int {
	// "### 10. **Сумма элементов массива**\n",
	// "Напишите функцию `sumArray(arr []int) int`, которая принимает массив целых чисел
	// и возвращает их сумму.\n",
	var res int
	for _, el := range arr {
		res += el
	}
	return res
}

func main() {
	//mas := [...]int{1, 4, 0, 2, 3}
	// fmt.Println(task2(mas[:]))
	//fmt.Println(task3(5))
	//fmt.Println(task4("asdksa"))
	//fmt.Println(task5(15))
	//fmt.Println(task6(15))
	//fmt.Println(task7(10))
	//fmt.Println(task8(mas[:]))
	//fmt.Println(task10(mas[:]))
}
