package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	size := 10
	result := generateRandomElements(size)

	// Проверяем длину среза
	if len(result) != size {
		t.Errorf("Ожидаемая длина %d, получена %d", size, len(result))
	}

	// Проверяем, что все элементы в допустимом диапазоне [1, 1000000]
	for _, num := range result {
		if num < 1 || num > 1000000 {
			t.Errorf("Число %d выходит за допустимый диапазон [1, 1000000]", num)
		}
	}
}

func TestGenerateRandomElements_ZeroSize(t *testing.T) {
	size := 0
	result := generateRandomElements(size)

	if len(result) != 0 {
		t.Errorf("Для размера 0 ожидался пустой срез, но получена длина %d", len(result))
	}
}

func TestMaximum_EmptySlice(t *testing.T) {
	data := []int{}
	result := maximum(data)
	
	if result != 0 {
		t.Errorf("Для пустого слайса ожидался 0, но получено %d", result)
	}
}

func TestMaximum_SingleElement(t *testing.T) {
	data := []int{42}
	result := maximum(data)
	
	if result != 42 {
		t.Errorf("Для слайса [42] ожидался 42, но получено %d", result)
	}
}

func TestMaximum_AllSameElements(t *testing.T) {
	data := []int{5, 5, 5, 5}
	result := maximum(data)
	
	if result != 5 {
		t.Errorf("Для слайса [5,5,5,5] ожидался 5, но получено %d", result)
	}
}

// TestMaximum_PositiveNumbers проверяет слайс только с положительными числами
func TestMaximum(t *testing.T) {
	data := []int{1, 100, 50, 75, 25}
	result := maximum(data)
	
	expected := 100
	if result != expected {
		t.Errorf("Для слайса %v ожидался %d, но получено %d", data, expected, result)
	}
}
