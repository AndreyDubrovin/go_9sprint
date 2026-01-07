package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		num := r.Intn(1000000) + 1
		slice[i] = num
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	maxNumber := 0
	for _, n := range data {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}

// // maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	sliceNumbers := make([]int, CHUNKS)
	maxNumber := 0
	var wg sync.WaitGroup
	sliceLen := len(data)            // длинна
	oneChankLen := sliceLen / CHUNKS // 12500000
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			maxNumber := 0
			startHalf := oneChankLen * i
			lastHalf := oneChankLen * (i + 1)
			sliceHalf := data[startHalf:lastHalf]
			for _, n := range sliceHalf {
				if n > maxNumber {
					maxNumber = n
				}
			}
			sliceNumbers[i] = maxNumber
		}(i)

	}
	wg.Wait()
	for _, n := range sliceNumbers {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	sliceNumbers := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(sliceNumbers)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(sliceNumbers)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
