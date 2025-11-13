package relation

// FindEquivalenceClasses находит классы эквивалентности
func FindEquivalenceClasses(rel Relation) [][]string {
	visited := make(map[string]bool)
	var classes [][]string

	for _, elem := range rel.Elements {
		if !visited[elem] {
			class := findEquivalenceClass(rel, elem, visited)
			classes = append(classes, class)
		}
	}

	return classes
}

func findEquivalenceClass(rel Relation, start string, visited map[string]bool) []string {
	var class []string
	var queue []string

	queue = append(queue, start)
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		class = append(class, current)

		for _, elem := range rel.Elements {
			if !visited[elem] && areEquivalent(rel, current, elem) {
				visited[elem] = true
				queue = append(queue, elem)
			}
		}
	}

	return class
}

func areEquivalent(rel Relation, a, b string) bool {
	return rel.ContainsPair(a, b) && rel.ContainsPair(b, a)
}

// FindMinMaxElements находит минимальные и максимальные элементы
func FindMinMaxElements(rel Relation) ([]string, []string) {
	var minElements, maxElements []string

	for _, elem := range rel.Elements {
		isMin := true
		isMax := true

		for _, other := range rel.Elements {
			if elem != other {
				// Если есть элемент, который меньше текущего
				if rel.ContainsPair(other, elem) {
					isMin = false
				}
				// Если есть элемент, который больше текущего
				if rel.ContainsPair(elem, other) {
					isMax = false
				}
			}
		}

		if isMin {
			minElements = append(minElements, elem)
		}
		if isMax {
			maxElements = append(maxElements, elem)
		}
	}

	return minElements, maxElements
}
