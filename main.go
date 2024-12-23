package main

import (
	"algorithms/map_algorithms"
	"algorithms/slice_algorithms"
	"fmt"
	"sort"
)

func main() {
	mass := []int{1, 2, 6, 9, 4, 6, 8, 4, 1, 8}
	massCon := []int{5, 2, 7, 2, 5}

	fmt.Println("Слайсы------------------------------------------------")
	fmt.Println("Сумма чисел в массиве:", slice_algorithms.Sum(mass[:]))

	minNum, maxNum := slice_algorithms.MinMax(mass)
	fmt.Printf("Минимальное число: %d, максимальное число: %d\n", minNum, maxNum)

	fmt.Println("Реверс массива:", slice_algorithms.Rev(mass))

	fmt.Printf("Среднее значение: %.1f\n", slice_algorithms.AverageValue(mass))

	evenNum, oddNum := slice_algorithms.EvenOdd(mass) // четное, нечетное
	fmt.Printf("Кол-во четных чисел: %d, кол-во нечетных чисел: %d\n", evenNum, oddNum)

	fmt.Println("Объединенный массив:", slice_algorithms.MasConcatenation(mass, massCon))

	fmt.Println("Массив 1 палиндром:", slice_algorithms.Palindrome(mass))
	fmt.Println("Массив 2 палиндром:", slice_algorithms.Palindrome(massCon))

	fmt.Println("Массив без дубликатов:", slice_algorithms.Dub(mass))

	sort.Ints(mass)
	fmt.Println("Бинарный поиск, индекс:", slice_algorithms.BinSearch(mass, 4))

	fmt.Println("Мапы--------------------------------------------------")
	inputString := "Ночь, улица, фонарь, аптека,\n" +
		"Бессмысленный и тусклый свет.\n" +
		"Живи ещё хоть четверть века -\n" +
		"Всё будет так. Исхода нет.\n" +
		"Умрёшь - начнёшь опять сначала\n" +
		"И повторится всё, как встарь:\n" +
		"Ночь, ледяная рябь канала,\n" +
		"Аптека, улица, фонарь."
	fmt.Println("Частотный анализ:", map_algorithms.FreqAnalysis(inputString))

	massDub := []int{1, 1, 1, 9, 9, 8, 4, 8, 7, 7, 7, 7, 5, 5}
	fmt.Println("Поиск уникального числа. Число:", map_algorithms.SearchUniqElem(massDub))

	maxValMap := map[string]int{
		"Кошка":   2,
		"Птичка":  3,
		"Собачка": 2,
		"Жираф":   3,
		"Пони":    2,
	}
	fmt.Println("Ключи с максимальными значениями. Ключи:", map_algorithms.KeyMaxVal(maxValMap))

	mapForCon := map[string]int{
		"Кенгуру": 2,
		"Птичка":  2,
		"Ежик":    2,
	}
	fmt.Println("Объединение мап. Результирующая мапа:", map_algorithms.MapConcatenation(maxValMap, mapForCon))

	sliceForGroupByPrefix := []string{"бабахнуть", "бабочка", "бабник", "сажа", "саженец", "навага", "навар", "шакал"}
	fmt.Println("Группировка по префиксу:", map_algorithms.GroupByPrefix(sliceForGroupByPrefix, 3))

	findSumUniq := map[int]int{
		2: 2,
		1: 3,
		5: 3,
		4: 4,
		3: 3,
		6: 4,
	}
	fmt.Println("Нахождение суммы пар с уникальными значениями:", map_algorithms.FindSumUniq(findSumUniq))

	storageMap := map_algorithms.MyMap{
		"банан":    5,
		"яблоко":   10,
		"вишня":    0,
		"мандарин": 15,
		"виноград": 2,
	}

	buySlice := []string{"банан", "банан", "виноград", "виноград", "мандарин"}
	err := storageMap.UpdateStorageMap(buySlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Кассовый учет:", storageMap)
	}

}
