package main

import (
	"discrete-math-lab2/internal/fileio"
	"discrete-math-lab2/internal/relation"
	"fmt"
	"log"
)

func main() {
	inputFile := "input.txt"
	outputFile := "output.txt"

	fmt.Printf("Чтение отношения из файла: %s\n", inputFile)

	// Чтение отношения из файла
	rel, err := fileio.ReadRelationFromFile(inputFile)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}

	fmt.Printf("Успешно прочитано отношение:\n")
	fmt.Printf("  Элементы: %v\n", rel.Elements)
	fmt.Printf("  Количество пар: %d\n", len(rel.Pairs))

	// Анализ свойств
	fmt.Println("Анализ свойств отношения...")
	properties := relation.CheckProperties(rel)

	// Запись результатов в файл
	fmt.Printf("Запись результатов в файл: %s\n", outputFile)
	err = fileio.WriteResults(outputFile, rel, properties)
	if err != nil {
		log.Fatalf("Ошибка при записи файла: %v", err)
	}
	fmt.Println("Анализ завершен! Результаты сохранены в", outputFile)

	// Вывод кратких результатов в консоль
	fmt.Println("\nКРАТКИЕ РЕЗУЛЬТАТЫ:")
	fmt.Println("===================")

	if properties["equivalence"] {
		classes := relation.FindEquivalenceClasses(rel)
		fmt.Printf("✓ Отношение эквивалентности (классов: %d)\n", len(classes))
	}

	if properties["partialOrder"] {
		minElements, maxElements := relation.FindMinMaxElements(rel)
		fmt.Printf("✓ Отношение порядка (min: %v, max: %v)\n", minElements, maxElements)
	}
}
