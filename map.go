package gofn

// MapEqual compares contents of 2 map
func MapEqual[K comparable, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

// MapEqualPred compares contents of 2 map
func MapEqualPred[K comparable, V any](m1, m2 map[K]V, equalFunc func(v1, v2 V) bool) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || !equalFunc(v1, v2) {
			return false
		}
	}
	return true
}

// MapContainKeys tests if a map contains one or more keys
func MapContainKeys[K comparable, V any](m map[K]V, keys ...K) bool {
	for _, k := range keys {
		if _, exists := m[k]; !exists {
			return false
		}
	}
	return true
}

// MapContainValues tests if a map contains one or more values (complexity is O(n))
// If you often need to check existence of map value, consider using bi-map data structure
func MapContainValues[K comparable, V comparable](m map[K]V, values ...V) bool {
	for _, v := range values {
		found := false
		for _, x := range m {
			if x == v {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// MapKeys gets map keys as slice
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// MapValues gets map values as slice
func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// MapEntries returns a slice of map entries as Tuple2 type
func MapEntries[K comparable, V any](m map[K]V) []*Tuple2[K, V] {
	items := make([]*Tuple2[K, V], len(m))
	i := 0
	for k, v := range m {
		items[i] = &Tuple2[K, V]{k, v}
		i++
	}
	return items
}

// MapUpdate merges map content with another map
// Not change the target map, only change the source map
func MapUpdate[K comparable, V any](m1, m2 map[K]V) map[K]V {
	if m1 == nil {
		m1 = make(map[K]V, len(m2))
	}
	if m2 == nil {
		return m1
	}
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

// MapUpdateExistingOnly update map existing items with another map
// Not change the target map, only change the source map
func MapUpdateExistingOnly[K comparable, V any](m1, m2 map[K]V) map[K]V {
	if m1 == nil {
		return make(map[K]V, 0)
	}
	if m2 == nil {
		return m1
	}
	for k, v := range m2 {
		if _, existing := m1[k]; existing {
			m1[k] = v
		}
	}
	return m1
}

// MapUpdateNewOnly update map with another map and not override the existing values
// Not change the target map, only change the source map
func MapUpdateNewOnly[K comparable, V any](m1, m2 map[K]V) map[K]V {
	if m1 == nil {
		return MapUpdate(make(map[K]V, len(m2)), m2)
	}
	if m2 == nil {
		return m1
	}
	for k, v := range m2 {
		if _, existing := m1[k]; !existing {
			m1[k] = v
		}
	}
	return m1
}

// MapGet gets the value for the key, if not exist, returns the default one
func MapGet[K comparable, V any](m map[K]V, k K, defaultVal V) V {
	if val, ok := m[k]; ok {
		return val
	}
	return defaultVal
}

// MapPop deletes and returns the value of the key if exists, returns the default one if not
func MapPop[K comparable, V any](m map[K]V, k K, defaultVal V) V {
	if val, ok := m[k]; ok {
		delete(m, k)
		return val
	}
	return defaultVal
}

func MapSetDefault[K comparable, V any](m map[K]V, k K, defaultVal V) V {
	if val, ok := m[k]; ok {
		return val
	}
	if m != nil {
		m[k] = defaultVal
	}
	return defaultVal
}

func MapUnionKeys[K comparable, V any](m1 map[K]V, ms ...map[K]V) []K {
	keys := MapKeys(m1)
	for _, m := range ms {
		keys = Union(keys, MapKeys(m))
	}
	return keys
}

func MapIntersectionKeys[K comparable, V any](m1 map[K]V, ms ...map[K]V) []K {
	keys := MapKeys(m1)
	for _, m := range ms {
		keys = Intersection(keys, MapKeys(m))
	}
	return keys
}

func MapDifferenceKeys[K comparable, V any](m1, m2 map[K]V) ([]K, []K) {
	return Difference(MapKeys(m1), MapKeys(m2))
}

func MapCopy[K comparable, V any](m map[K]V, onlyKeys ...K) map[K]V {
	// Copy the whole map (this is shallow copy)
	if len(onlyKeys) == 0 {
		ret := make(map[K]V, len(m))
		for k, v := range m {
			ret[k] = v
		}
		return ret
	}

	// Copy only keys in the list
	ret := make(map[K]V, len(onlyKeys))
	for _, k := range onlyKeys {
		v, ok := m[k]
		if ok {
			ret[k] = v
		}
	}
	return ret
}

func MapCopyExcludeKeys[K comparable, V any](m map[K]V, excludedKeys ...K) map[K]V {
	// Copy the whole map (this is shallow copy)
	if len(excludedKeys) == 0 {
		ret := make(map[K]V, len(m))
		for k, v := range m {
			ret[k] = v
		}
		return ret
	}

	excludedMap := make(map[K]struct{}, len(excludedKeys))
	for _, k := range excludedKeys {
		excludedMap[k] = struct{}{}
	}

	// Copy only keys not in the list
	ret := make(map[K]V, len(m))
	for k, v := range m {
		_, ok := excludedMap[k]
		if !ok {
			ret[k] = v
		}
	}
	return ret
}
