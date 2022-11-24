package maps

// Equal проверяет две карты на равенство.
func Equal[TKey comparable, TValue comparable](map1, map2 map[TKey]TValue) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		if value2, ok := map2[key]; !ok || value1 != value2 {
			return false
		}
	}

	return true
}
