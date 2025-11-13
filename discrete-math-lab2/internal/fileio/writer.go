package fileio

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"discrete-math-lab2/internal/relation"
)

// WriteResults записывает результаты анализа в файл
func WriteResults(filename string, rel relation.Relation, properties relation.Properties) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Вспомогательная функция для записи
	write := func(format string, args ...interface{}) error {
		_, err := fmt.Fprintf(writer, format, args...)
		return err
	}

	// Записываем информацию об отношении
	if err := write("АНАЛИЗ ОТНОШЕНИЯ\n"); err != nil {
		return err
	}
	if err := write("================\n\n"); err != nil {
		return err
	}

	if err := write("Множество-носитель: %s\n", strings.Join(rel.Elements, ", ")); err != nil {
		return err
	}
	if err := write("Пары отношения:\n"); err != nil {
		return err
	}
	for _, pair := range rel.Pairs {
		if err := write("  (%s, %s)\n", pair[0], pair[1]); err != nil {
			return err
		}
	}
	if err := write("\n"); err != nil {
		return err
	}

	// Записываем свойства
	if err := write("СВОЙСТВА ОТНОШЕНИЯ:\n"); err != nil {
		return err
	}
	if err := write("-------------------\n"); err != nil {
		return err
	}

	propertyNames := map[string]string{
		"reflexive":     "Рефлексивность",
		"irreflexive":   "Иррефлексивность",
		"symmetric":     "Симметричность",
		"antisymmetric": "Антисимметричность",
		"asymmetric":    "Асимметричность",
		"transitive":    "Транзитивность",
		"connected":     "Связность",
		"equivalence":   "Эквивалентность",
		"partialOrder":  "Частичный порядок",
	}

	order := []string{
		"reflexive", "irreflexive", "symmetric", "antisymmetric",
		"asymmetric", "transitive", "connected", "equivalence", "partialOrder",
	}

	for _, prop := range order {
		sign := "-"
		if properties[prop] {
			sign = "+"
		}
		if err := write("%s: %s\n", propertyNames[prop], sign); err != nil {
			return err
		}
	}
	if err := write("\n"); err != nil {
		return err
	}

	// Если отношение эквивалентности
	if properties["equivalence"] {
		if err := write("ОТНОШЕНИЕ ЭКВИВАЛЕНТНОСТИ\n"); err != nil {
			return err
		}
		if err := write("-------------------------\n"); err != nil {
			return err
		}
		classes := relation.FindEquivalenceClasses(rel)
		if err := write("Классы эквивалентности (%d):\n", len(classes)); err != nil {
			return err
		}
		for i, class := range classes {
			if err := write("  Класс %d: {%s}\n", i+1, strings.Join(class, ", ")); err != nil {
				return err
			}
		}
		if err := write("Индекс разбиения: %d\n\n", len(classes)); err != nil {
			return err
		}
	}

	// Если отношение порядка
	if properties["partialOrder"] {
		if err := write("ОТНОШЕНИЕ ПОРЯДКА\n"); err != nil {
			return err
		}
		if err := write("------------------\n"); err != nil {
			return err
		}
		minElements, maxElements := relation.FindMinMaxElements(rel)
		if err := write("Минимальные элементы: {%s}\n", strings.Join(minElements, ", ")); err != nil {
			return err
		}
		if err := write("Максимальные элементы: {%s}\n", strings.Join(maxElements, ", ")); err != nil {
			return err
		}
	}

	return nil
}
