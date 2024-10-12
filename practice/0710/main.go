package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

func task11() {
	// 	1. Перевод чисел из одной системы счисления в другую
	//    Напишите программу, которая принимает на вход число в произвольной системе счисления (от 2 до 36) и переводит его в другую систему счисления.
	//    *Input:* число, исходная система, конечная система.
	//    *Output:* число в конечной системе.

	var num, s1, s2 string
	fmt.Scanln(&num, &s1, &s2)

	n := num[0 : len(num)-1]
	s1val, _ := strconv.Atoi(s1[0 : len(s1)-1])
	s2val, _ := strconv.Atoi(s2)

	if s1val < 2 || s1val > 36 {
		fmt.Println("исходная система счисления может принимать число от 2 до 36")
		return
	}
	i, err := strconv.ParseInt(n, s1val, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(strconv.FormatInt(i, s2val))

}

func task12(a, b, c float64) (complex128, complex128) {
	//		Решение квадратного уравнения
	//	   Реализуйте функцию для нахождения корней квадратного уравнения вида \( ax^2 + bx + c = 0 \). Учтите случай комплексных корней.
	// *Input:* коэффициенты \(a, b, c \).
	// *Output:* корни уравнения.

	d := b*b - 4*a*c

	k1 := complex(0, 0)
	k2 := complex(0, 0)

	if d > 0 {
		k1 = complex((-b+math.Sqrt(d))/(2*a), 0)
		k2 = complex((-b-math.Sqrt(d))/(2*a), 0)
	} else if d == 0 {
		k1 = complex(-b/(2*a), 0)
		k2 = k1
	} else {
		realPart := -b / (2 * a)
		imPart := math.Sqrt(-d) / (2 * a)
		k1 = complex(realPart, imPart)
		k2 = complex(realPart, -imPart)
	}

	fmt.Printf("Корни уравнения: %.2f; %.2f\n", k1, k2)
	return k1, k2
}

func task13(arr []int) []int {
	//  3. Сортировка чисел по модулю
	//     Дан массив чисел. Напишите функцию, которая отсортирует его элементы по возрастанию их абсолютных значений.
	//     *Input:* массив чисел.
	//     *Output:* отсортированный массив.
	sort.Slice(arr, func(i, j int) bool {
		if math.Abs(float64(arr[i])) == math.Abs(float64(arr[j])) {
			return arr[i] < arr[j]
		}
		return math.Abs(float64(arr[i])) < math.Abs(float64(arr[j]))
	})
	fmt.Println(arr)
	return arr
}

func task14(arr1 []int, arr2 []int) []int {

	//  4. Слияние двух отсортированных массивов
	//     Напишите функцию, которая объединяет два отсортированных массива в один, сохраняя их упорядоченность.
	//     *Input:* два отсортированных массива.
	//     *Output:* отсортированный массив.
	var s []int
	s = append(s, arr1...)
	s = append(s, arr2...)
	fmt.Println(s)
	return s
}

func task15(str, substr string) int {
	//  5. Нахождение подстроки в строке без использования встроенных функций
	//     Реализуйте функцию, которая находит первую позицию вхождения одной строки в другую. Если подстрока не найдена, вернуть -1.
	//     *Input:* две строки.
	//     *Output:* индекс первого вхождения или -1.
	n, m := len(str), len(substr)
	for i := 0; i <= n-m; i++ {
		if str[i:i+m] == substr {
			fmt.Println(i)
			return i
		}
	}
	fmt.Println(-1)
	return -1
}

func task21() {
	//  1. Калькулятор с расширенными операциями
	//     Напишите программу, которая выполняет различные математические операции (+, -, *, /, ^, %), заданные пользователем. Реализуйте обработку ошибок, связанных с делением на ноль или недопустимой операцией.
	//     *Input:* два числа и оператор.
	//     *Output:* результат операции.
	var a, operation, b string
	fmt.Scanln(&a, &operation, &b)

	num1, _ := strconv.Atoi(a)
	num2, _ := strconv.Atoi(b)

	switch operation {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num1 == 0 || num2 == 0 {
			fmt.Println("Деление невозможно")
			return
		}
		fmt.Println(num1 / num2)
	case "^":
		fmt.Println(math.Pow(float64(num1), float64(num2)))
	case "%":
		fmt.Println(num1 % num2)
	default:
		fmt.Println("Операция не определена")
	}
}

func task22() {
	// 2. Проверка палиндрома
	// Реализуйте функцию, которая проверяет, является ли строка палиндромом (игнорируя пробелы, знаки препинания и регистр).
	// *Input:* строка.
	// *Output:* true/false.
	var str string
	var res bool
	fmt.Scanln(&str)
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[len(str)-1-i] {
			res = true
			continue
		} else {
			res = false
			break
		}
	}
	fmt.Println(res)
}

func task23() {
	// 3. Нахождение пересечения трех отрезков
	// Даны три отрезка на числовой оси (их начальные и конечные точки). Нужно определить, существует ли область пересечения всех трех отрезков.
	// *Input:* три пары чисел, задающих отрезки.
	// *Output:* true/false.
	var x1, y1, x2, y2, x3, y3 int
	fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3)
	minX := max(x1, x2, x3)
	minY := min(y1, y2, y3)

	if minX <= minY {
		fmt.Println(true)
		return
	}
	fmt.Println(false)
}

func task24() {
	// 	4. Выбор самого длинного слова в предложении
	//    Напишите программу, которая находит самое длинное слово в предложении, игнорируя знаки препинания.
	//    *Input:* строка.
	//    *Output:* самое длинное слово.
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	arr := strings.Split(str, " ")
	var symbols = [...]string{".", ",", ";", "?", "!"}
	maxLenghtWord := 0
	wordLenght := 0
	indexEl := 0
	for i := 0; i < len(arr)-1; i++ {
		el := arr[i]
		lastChar := el[len(el)-1:]
		wordLenght = len(el)

		for j := 0; j < len(symbols)-1; j++ {
			if symbols[j] == lastChar {
				wordLenght = len(el) - 1
			}
		}
		if maxLenghtWord < wordLenght {
			maxLenghtWord = wordLenght
			indexEl = i
		}
	}
	fmt.Println(maxLenghtWord)
	fmt.Println(arr[indexEl])
}

func task25() {
	//  5. Проверка высокосного года
	//     Реализуйте функцию, которая проверяет, является ли введенный год високосным по правилам григорианского календаря.
	//     *Input:* год.
	//     *Output:* true/false.
	var year int
	fmt.Scan(&year)
	if year%4 == 0 {
		fmt.Println(true)
		return
	}
	fmt.Println(false)
}

func task31() {
	//  1. Числа Фибоначчи до определенного числа
	//     Напишите программу, которая выводит все числа Фибоначчи, не превышающие заданного значения.
	//     *Input:* целое число \( n \).
	//     *Output:* последовательность чисел Фибоначчи.
	var n int
	fmt.Scan(&n)

	var fib = make([]int, n)
	fib[0] = 1
	fib[1] = 2
	val := 0
	for i := 1; i < n; i++ {
		val = fib[i] + fib[i-1]
		if val > n {
			break
		}
		fib[i+1] = val

	}
	for _, el := range fib {
		if el == 0 {
			break
		}
		fmt.Print(el, " ")
	}
	fmt.Println()
}

func task32() {
	//  2. Определение простых чисел в диапазоне
	//     Реализуйте функцию, которая принимает два числа и выводит все простые числа в диапазоне между ними.
	//     *Input:* два числа.
	//     *Output:* список простых чисел.
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	for i := a; i <= b; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			fmt.Print(i, " ")
		}
	}

	fmt.Println()
}

func task33(n1, n2 int) {
	//  3. Числа Армстронга в заданном диапазоне
	//     Напишите программу, которая выводит все числа Армстронга в заданном диапазоне.
	// Число Армстронга – это число, равное сумме своих цифр, возведенных в степень, равную количеству цифр числа.
	//     *Input:* два числа, задающие диапазон.
	//     *Output:* список чисел Армстронга.

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

func task34() {
	//  4. Реверс строки
	//     Напишите программу, которая принимает строку и возвращает ее в обратном порядке, не используя встроенные функции для реверса строк.
	//     *Input:* строка.
	//     *Output:* строка в обратном порядке.

	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	for i := len(str) - 3; i >= 1; i-- {
		fmt.Printf("%c", str[i])
	}
	fmt.Println()
}

func task35(n1, n2 int) int {
	//		Нахождение наибольшего общего делителя (НОД)
	//	   Реализуйте алгоритм Евклида для нахождения наибольшего общего делителя двух чисел с использованием цикла.
	//	   *Input:* два числа.
	//	   *Output:* НОД.
	if n1 < n2 {
		n1, n2 = n2, n1
	}
	for n2 != 0 {
		t := n2
		n2 = n1 % n2
		n1 = t
	}

	fmt.Println(n1)
	return n1
}

func main() {

}
