package main

import (
	"math/rand"
	"runtime"
	"testing"
)

// Генерация данных
const size = 1000000

var sliceData []int
var mapData map[int]int
var arrayData [size]int

func init() {
	sliceData = make([]int, size)
	mapData = make(map[int]int, size)
	for i := 0; i < size; i++ {
		value := rand.Intn(size)
		sliceData[i] = value
		mapData[i] = value
		arrayData[i] = value
	}
}

// Бенчмарк тест для массива
func BenchmarkArrayAccess(b *testing.B) {
	defer runtime.GC()

	index := rand.Intn(size)
	for i := 0; i < b.N; i++ {
		_ = arrayData[index]
	}
}

// Бенчмарк тест для слайса
func BenchmarkSliceAccess(b *testing.B) {
	defer runtime.GC()

	index := rand.Intn(size)
	for i := 0; i < b.N; i++ {
		_ = sliceData[index]
	}
}

// Бенчмарк тест для мапы
func BenchmarkMapAccess(b *testing.B) {
	defer runtime.GC()

	key := rand.Intn(size)
	for i := 0; i < b.N; i++ {
		_ = mapData[key]
	}
}
