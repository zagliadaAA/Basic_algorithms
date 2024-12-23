package map_algorithms

import (
	"fmt"
	"regexp"
	"strings"
)

// FreqAnalysis частотный анализ. Принимает текст (строку) и возвращает мапу map[string]int,
// где ключ — это слово, а значение — частота его появления.
func FreqAnalysis(input string) map[string]int {
	freqMap := make(map[string]int)

	words := regexp.MustCompile("[\n\t*, %_.:-]").Split(input, -1)

	for _, word := range words {
		if word != "" {
			word = strings.TrimSpace(word)
			word = strings.ToLower(word)
			freqMap[word]++
		}
	}

	return freqMap
}

// SearchUniqElem поиск уникального значения. Принимает массив целых чисел, в котором все числа,
// кроме одного, имеют дубликаты. Используй мапу, чтобы найти и вернуть уникальное число.
func SearchUniqElem(mass []int) int {
	numMap := make(map[int]int)

	for _, value := range mass {
		numMap[value]++
	}

	for key, value := range numMap {
		if value == 1 {
			return key
		}
	}

	return -1
}

// KeyMaxVal Ключи с максимальным значением. Дана мапа map[string]int,
// в которой могут быть разные ключи с одинаковыми значениями. Напиши функцию,
// которая возвращает массив всех ключей с максимальным значением.
func KeyMaxVal(inputMap map[string]int) []string {
	maxValue := 0
	keyMaxValSlice := make([]string, 0)

	for key, value := range inputMap {
		if value > maxValue {
			maxValue = value
			keyMaxValSlice = []string{key}
		} else if value == maxValue {
			keyMaxValSlice = append(keyMaxValSlice, key)
		}
	}

	//либо 2 цикла
	/*for _, value := range inputMap {
		if value > maxValue {
			maxValue = value
		}
	}

	for key, value := range inputMap {
		if value == maxValue {
			keyMaxValSlice = append(keyMaxValSlice, key)
		}
	}*/

	return keyMaxValSlice
}

// MapConcatenation Слияние мап с выбором максимального значения: Даны две мапы map[string]int.
// Напиши функцию, которая объединяет их, оставляя для каждого ключа только максимальное значение.
func MapConcatenation(map1 map[string]int, map2 map[string]int) map[string]int {
	for key, value := range map2 {
		if _, exists := map1[key]; !exists {
			map1[key] = value
		} else if map1[key] < value {
			map1[key] = value
		}
	}

	return map1
}

// GroupByPrefix Группировка по префиксу: Напиши функцию, которая принимает массив строк и группирует их
// по префиксу заданной длины. Результат — мапа map[string][]string, где ключи — это уникальные префиксы,
// а значения — массивы слов с этим префиксом.
func GroupByPrefix(mass []string, prefixLen int) map[string][]string {
	groupByPrefixMap := make(map[string][]string)

	for i, val := range mass {
		if len(val) >= prefixLen {
			prefix := val[0:prefixLen]
			groupByPrefixMap[prefix] = append(groupByPrefixMap[prefix], mass[i])
		}
	}

	return groupByPrefixMap
}

// FindSumUniq Нахождение суммы пар с уникальными значениями: Дана мапа map[int]int, где ключи и значения — целые числа.
// Напиши функцию, которая находит все пары ключ-значение, такие, что их сумма уникальна.
// Верни результат в виде мапы, где ключ — это сумма, а значение — массив пар [ключ, значение].
func FindSumUniq(inputMap map[int]int) map[int][][]int {
	resultMap := make(map[int][][]int)

	for key, value := range inputMap {
		pair := []int{key, value}
		resultMap[key+value] = append(resultMap[key+value], pair)
	}

	//удаление !уникальных
	for key, value := range resultMap {
		if len(value) > 1 {
			delete(resultMap, key)
		}
	}

	return resultMap
}

type MyMap map[string]int

// UpdateStorageMap Кассовый учёт: Дана мапа map[string]int, где ключи — это названия товаров, а значения — их количество на складе.
// Напиши функцию, которая принимает список покупок (массив строк) и обновляет мапу склада,
// уменьшая количество каждого купленного товара. Если товара не хватает на складе, выдай ошибку или
// выведи сообщение об отсутствии.
func (s *MyMap) UpdateStorageMap(buySlice []string) error {
	buyMap := make(map[string]int)

	for _, value := range buySlice {
		buyMap[value]++
	}

	for key, value := range buyMap {
		if _, ok := (*s)[key]; !ok {
			return fmt.Errorf("товара %s не существует", key)
		} else if (*s)[key] < value {
			return fmt.Errorf("товара %s нет в нужном количестве", key)
		} else {
			(*s)[key] -= value
		}
	}

	return nil
}
