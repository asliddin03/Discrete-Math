package relation

// Relation представляет математическое отношение
type Relation struct {
	Elements []string
	Pairs    [][2]string
}

// New создаем новое отношение
func New(elements []string, pairs [][2]string) Relation {
	return Relation{
		Elements: elements,
		Pairs:    pairs,
	}
}

// ContainsPair проверяет, содержит ли отношение указанную пару
func (r Relation) ContainsPair(a, b string) bool {
	for _, pair := range r.Pairs {
		if pair[0] == a && pair[1] == b {
			return true
		}
	}
	return false
}
