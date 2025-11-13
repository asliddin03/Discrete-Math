package relation

// Properties представляет свойства отношения
type Properties map[string]bool

// CheckProperties проверяет все свойства отношения
func CheckProperties(rel Relation) Properties {
	properties := make(Properties)

	properties["reflexive"] = isReflexive(rel)
	properties["irreflexive"] = isIrreflexive(rel)
	properties["symmetric"] = isSymmetric(rel)
	properties["antisymmetric"] = isAntisymmetric(rel)
	properties["asymmetric"] = isAsymmetric(rel)
	properties["transitive"] = isTransitive(rel)
	properties["connected"] = isConnected(rel)

	// Отношение эквивалентности: рефлексивное, симметричное, транзитивное
	properties["equivalence"] = properties["reflexive"] && properties["symmetric"] && properties["transitive"]

	// Отношение частичного порядка: рефлексивное, антисимметричное, транзитивное
	properties["partialOrder"] = properties["reflexive"] && properties["antisymmetric"] && properties["transitive"]

	return properties
}

func isReflexive(rel Relation) bool {
	for _, elem := range rel.Elements {
		if !rel.ContainsPair(elem, elem) {
			return false
		}
	}
	return true
}

func isIrreflexive(rel Relation) bool {
	for _, elem := range rel.Elements {
		if rel.ContainsPair(elem, elem) {
			return false
		}
	}
	return true
}

func isSymmetric(rel Relation) bool {
	for _, pair := range rel.Pairs {
		if !rel.ContainsPair(pair[1], pair[0]) {
			return false
		}
	}
	return true
}

func isAntisymmetric(rel Relation) bool {
	for _, pair := range rel.Pairs {
		if pair[0] != pair[1] && rel.ContainsPair(pair[1], pair[0]) {
			return false
		}
	}
	return true
}

func isAsymmetric(rel Relation) bool {
	for _, pair := range rel.Pairs {
		if rel.ContainsPair(pair[1], pair[0]) {
			return false
		}
	}
	return true
}

func isTransitive(rel Relation) bool {
	for _, pair1 := range rel.Pairs {
		for _, pair2 := range rel.Pairs {
			if pair1[1] == pair2[0] {
				if !rel.ContainsPair(pair1[0], pair2[1]) {
					return false
				}
			}
		}
	}
	return true
}

func isConnected(rel Relation) bool {
	for i, a := range rel.Elements {
		for j, b := range rel.Elements {
			if i != j {
				if !rel.ContainsPair(a, b) && !rel.ContainsPair(b, a) {
					return false
				}
			}
		}
	}
	return true
}
