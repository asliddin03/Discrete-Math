package fileio

import (
	"bufio"
	"os"
	"strings"

	"discrete-math-lab2/internal/relation"
)

// ReadRelationFromFile читает отношение из файла
func ReadRelationFromFile(filename string) (relation.Relation, error) {
	file, err := os.Open(filename)
	if err != nil {
		return relation.Relation{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var elements []string
	var pairs [][2]string

	// Чтение элементов множества-носителя
	if scanner.Scan() {
		line := scanner.Text()
		elements = strings.Fields(line)
	}

	// Чтение пар отношения
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 {
			pairs = append(pairs, [2]string{parts[0], parts[1]})
		}
	}

	if err := scanner.Err(); err != nil {
		return relation.Relation{}, err
	}

	return relation.New(elements, pairs), nil
}
