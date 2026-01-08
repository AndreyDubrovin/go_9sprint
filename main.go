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
	if size == 0 {
		return nil
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = r.Int()
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
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
	if len(data) == 0 {
		return 0
	}
	sliceNumbers := make([]int, CHUNKS)
	var wg sync.WaitGroup
	sliceLen := len(data)            // длинна
	oneChankLen := sliceLen / CHUNKS // 12500000
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			startHalf := oneChankLen * i
			lastHalf := oneChankLen * (i + 1)
			sliceHalf := data[startHalf:lastHalf]
			sliceNumbers[i] = maximum(sliceHalf)
		}(i)

	}
	wg.Wait()
	return maximum(sliceNumbers)
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
