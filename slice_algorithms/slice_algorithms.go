package slice_algorithms

// Sum 1) Сумма элементов массива: напишите функцию, которая принимает массив целых чисел и возвращает
//
//	сумму всех его элементов.
func Sum(mass []int) int {
	sum := 0

	for _, num := range mass {
		sum += num
	}

	return sum
}

// MinMax 2) Поиск минимума и максимума: Создайте функцию, которая находит минимальное и максимальное
//
//	значение в массиве.
func MinMax(mass []int) (int, int) {
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
func Rev(mass []int) []int {
	revSlice := make([]int, len(mass), cap(mass))

	for i, _ := range mass {
		revSlice[i] = mass[len(mass)-i-1]
	}

	return revSlice
}

// AverageValue 4) Поиск среднего значения: Напишите функцию, которая вычисляет среднее значение элементов
//
//	массива и возвращает его.
func AverageValue(mass []int) float32 {
	sum := 0

	for _, num := range mass {
		sum += num
	}

	return float32(sum) / float32(len(mass))
	//либо сразу
	//return float32(sum(mass)) / float32(len(mass))
}

// EvenOdd 5) Подсчет четных и нечетных чисел: Реализуйте функцию, которая считает количество
//
//	четных и нечетных чисел в массиве и возвращает их количество.
func EvenOdd(mass []int) (int, int) {
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

// MasConcatenation AsConcatenation 6) Объединение массивов: Напишите функцию, которая принимает два массива и возвращает
//
//	новый массив, состоящий из элементов обоих массивов.
func MasConcatenation(mass1 []int, mass2 []int) []int {
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
func Palindrome(mass []int) bool {
	for i := 0; i < len(mass)/2; i++ {
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
func Dub(mass []int) []int {
	supMap := make(map[int]int)
	nonDubSlice := make([]int, 0)

	for _, num := range mass {
		supMap[num]++
	}

	for key, _ := range supMap {
		nonDubSlice = append(nonDubSlice, key)
	}

	return nonDubSlice
}

// BinSearch 9) Поиск индекса элемента: Реализуйте функцию, которая принимает массив и значение,
//
//	а затем возвращает индекс первого вхождения этого значения в массиве.
//	Если элемента нет, возвращает -1.  (Алгоритм бинарного поиска)
func BinSearch(mass []int, val int) int {
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
