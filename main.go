package main

import (
	"fmt"
	"sort"
)

func main() {
	mass := []int{1, 2, 6, 3, 4, 6, 3, 3, 2, 8, 2, 7, 4, 7, 8, 9, 7, 5, 3, 1, 8, 3, 3, 6, 9}
	massCon := []int{5, 2, 8, 2, 5}

	fmt.Println("Сумма чисел в массиве:", sum(mass[:]))

	minNum, maxNum := minMax(mass)
	fmt.Printf("Минимальное число: %d, максимальное число: %d\n", minNum, maxNum)

	fmt.Println("Реверс массива:", rev(mass))

	fmt.Printf("Среднее значение: %.1f\n", averageValue(mass))

	evenNum, oddNum := evenOdd(mass) // четное, нечетное
	fmt.Printf("Кол-во четных чисел: %d, кол-во нечетных чисел: %d\n", evenNum, oddNum)

	fmt.Println("Объединенный массив:", masConcatenation(mass, massCon))

	fmt.Println("Массив 1 палиндром:", palindrome(mass))
	fmt.Println("Массив 2 палиндром:", palindrome(massCon))

	fmt.Println("Массив без дубликатов:", dub(mass))

	sort.Ints(mass)
	fmt.Println(mass)
	fmt.Println("Бинарный поиск, индекс:", binSearch(mass, 3))

}

// Sum 1) Сумма элементов массива: напишите функцию, которая принимает массив целых чисел и возвращает
//
//	сумму всех его элементов.
func sum(mass []int) int {
	sum := 0

	for _, num := range mass {
		sum += num
	}

	return sum
}

// minMax 2) Поиск минимума и максимума: Создайте функцию, которая находит минимальное и максимальное
//
//	значение в массиве.
func minMax(mass []int) (int, int) {
	var minNum, maxNum = mass[0], mass[0]

	for _, num := range mass {
		if num < minNum {
			minNum = num
		}
		if num > maxNum {
			maxNum = num
		}
	}

	return minNum, maxNum
}

// Rev 3) Реверс массива: Реализуйте функцию, которая принимает массив и возвращает его в обратном порядке.
func rev(mass []int) []int {
	revSlice := make([]int, len(mass), cap(mass))

	for i, _ := range mass {
		revSlice[i] = mass[len(mass)-i-1]
	}

	return revSlice
}

// averageValue 4) Поиск среднего значения: Напишите функцию, которая вычисляет среднее значение элементов
//
//	массива и возвращает его.
func averageValue(mass []int) float32 {
	sum := 0

	for _, num := range mass {
		sum += num
	}

	return float32(sum) / float32(len(mass))
	//либо сразу
	//return float32(sum(mass)) / float32(len(mass))
}

// evenOdd 5) Подсчет четных и нечетных чисел: Реализуйте функцию, которая считает количество
//
//	четных и нечетных чисел в массиве и возвращает их количество.
func evenOdd(mass []int) (int, int) {
	evenNum, oddNum := 0, 0

	for _, num := range mass {
		if num%2 == 0 {
			evenNum++
		} else {
			oddNum++
		}
	}

	return evenNum, oddNum
}

// asConcatenation 6) Объединение массивов: Напишите функцию, которая принимает два массива и возвращает
//
//	новый массив, состоящий из элементов обоих массивов.
func masConcatenation(mass1 []int, mass2 []int) []int {
	conSlice := make([]int, len(mass1)+len(mass2), cap(mass1)+cap(mass2))

	for i := 0; i < len(mass1); i++ {
		conSlice[i] = mass1[i]
	}

	for i := 0; i < len(mass2); i++ {
		conSlice[i+len(mass1)] = mass2[i]
	}

	//либо с append
	//conSlice := append(mass1, mass2...)

	return conSlice
}

// Palindrome 7) Проверка на палиндром: Создайте функцию, которая проверяет, является ли массив
//
//	палиндромом (то есть элементы массива читаются одинаково слева направо и справа налево).
func palindrome(mass []int) bool {
	for i := 0; i < len(mass); i++ {
		if mass[i] == mass[len(mass)-1-i] {
			continue
		} else {
			return false
		}
	}

	return true
}

// Dub 8) Удаление дубликатов: Напишите функцию, которая удаляет дубликаты из массива и
//
//	возвращает новый массив с уникальными значениями.
func dub(mass []int) []int {
	nonDubSlice := make([]int, 0, cap(mass))
	flag := false

	for _, massVal := range mass {
		for _, nonVal := range nonDubSlice {
			if nonVal == massVal {
				flag = true
				break
			} else {
				flag = false
			}
		}

		if flag == false {
			nonDubSlice = append(nonDubSlice, massVal)
		}
	}

	return nonDubSlice
}

// binSearch 9) Поиск индекса элемента: Реализуйте функцию, которая принимает массив и значение,
//
//	а затем возвращает индекс первого вхождения этого значения в массиве.
//	Если элемента нет, возвращает -1.  (Алгоритм бинарного поиска)
func binSearch(mass []int, val int) int {
	left := 0
	right := len(mass)

	for left < right {
		middleVal := (left + right) / 2
		if mass[middleVal] == val {
			for mass[middleVal] == val { //нужно, чтобы найти именно первое вхождение
				save := middleVal
				middleVal--
				if mass[middleVal] == val {
					continue
				} else if mass[middleVal] != val {
					return save
				}
			}
		} else if mass[middleVal] > val {
			right = middleVal - 1
		} else if mass[middleVal] < val {
			left = middleVal + 1
		} else if left == right {
			return -1
		}
	}

	return -1
}
