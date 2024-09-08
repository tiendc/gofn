package gofn

// MapEqual compares contents of 2 maps
func MapEqual[K comparable, V comparable, M ~map[K]V](m1, m2 M) bool {
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

// MapEqualBy compares contents of 2 maps
func MapEqualBy[K comparable, V any, M ~map[K]V](m1, m2 M, equalCmp func(v1, v2 V) bool) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || !equalCmp(v1, v2) {
			return false
		}
	}
	return true
}

// Deprecated: use MapEqualBy instead
func MapEqualPred[K comparable, V any, M ~map[K]V](m1, m2 M, equalCmp func(v1, v2 V) bool) bool {
	return MapEqualBy(m1, m2, equalCmp)
}

// MapContainKeys tests if a map contains one or more keys
func MapContainKeys[K comparable, V any, M ~map[K]V](m M, keys ...K) bool {
	for _, k := range keys {
		if _, exists := m[k]; !exists {
			return false
		}
	}
	return true
}

// MapContainValues tests if a map contains one or more values (complexity is O(n)).
// If you often need to check existence of map value, consider using bi-map data structure.
func MapContainValues[K comparable, V comparable, M ~map[K]V](m M, values ...V) bool {
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
func MapKeys[K comparable, V any, M ~map[K]V](m M) []K {
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
func MapEntries[K comparable, V any, M ~map[K]V](m M) []*Tuple2[K, V] {
	items := make([]*Tuple2[K, V], len(m))
	i := 0
	for k, v := range m {
		items[i] = &Tuple2[K, V]{k, v}
		i++
	}
	return items
}

// MapUpdate merges map content with another map.
// Not change the target map, only change the source map.
func MapUpdate[K comparable, V any, M ~map[K]V](m1, m2 M) M {
	if m1 == nil {
		m1 = make(M, len(m2))
	}
	if m2 == nil {
		return m1
	}
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

// MapUpdateExistingOnly update map existing items with another map.
// Not change the target map, only change the source map.
func MapUpdateExistingOnly[K comparable, V any, M ~map[K]V](m1, m2 M) M {
	if m1 == nil {
		return make(M, 0)
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

// MapUpdateNewOnly update map with another map and not override the existing values.
// Not change the target map, only change the source map.
func MapUpdateNewOnly[K comparable, V any, M ~map[K]V](m1, m2 M) M {
	if m1 == nil {
		return MapUpdate(make(M, len(m2)), m2)
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
func MapGet[K comparable, V any, M ~map[K]V](m M, k K, defaultVal V) V {
	if val, ok := m[k]; ok {
		return val
	}
	return defaultVal
}

// MapPop deletes and returns the value of the key if exists, returns the default one if not
func MapPop[K comparable, V any, M ~map[K]V](m M, k K, defaultVal V) V {
	if val, ok := m[k]; ok {
		delete(m, k)
		return val
	}
	return defaultVal
}

// MapSetDefault sets default value for a key and returns the current value
func MapSetDefault[K comparable, V any, M ~map[K]V](m M, k K, defaultVal V) V {
	if m == nil {
		var zero V
		return zero
	}
	if val, ok := m[k]; ok {
		return val
	}
	m[k] = defaultVal
	return defaultVal
}

// MapUnionKeys returns a list of unique keys that are collected from multiple maps
func MapUnionKeys[K comparable, V any, M ~map[K]V](m1 M, ms ...M) []K {
	keys := MapKeys(m1)
	for _, m := range ms {
		keys = Union(keys, MapKeys(m))
	}
	return keys
}

// MapIntersectionKeys returns a list of unique keys that exist in all maps
func MapIntersectionKeys[K comparable, V any, M ~map[K]V](m1 M, ms ...M) []K {
	keys := MapKeys(m1)
	for _, m := range ms {
		keys = Intersection(keys, MapKeys(m))
	}
	return keys
}

// MapDifferenceKeys returns 2 lists of keys that are differences of 2 maps
func MapDifferenceKeys[K comparable, V any, M ~map[K]V](m1, m2 M) ([]K, []K) {
	return Difference(MapKeys(m1), MapKeys(m2))
}

// MapCopy returns a copied a map
func MapCopy[K comparable, V any, M ~map[K]V](m M) M {
	ret := make(M, len(m))
	for k, v := range m {
		ret[k] = v
	}
	return ret
}

// MapPick returns a new map with picking up the specified keys only
func MapPick[K comparable, V any, M ~map[K]V](m M, keys ...K) M {
	ret := make(M, len(keys))
	for _, k := range keys {
		v, ok := m[k]
		if ok {
			ret[k] = v
		}
	}
	return ret
}

// MapOmit omits keys from a map
func MapOmit[K comparable, V any, M ~map[K]V](m M, keys ...K) {
	for _, k := range keys {
		delete(m, k)
	}
}

// MapOmitCopy returns a new map with omitting the specified keys
func MapOmitCopy[K comparable, V any, M ~map[K]V](m M, keys ...K) M {
	omitKeys := MapSliceToMapKeys(keys, struct{}{})

	// Copy only keys not in the list
	ret := make(M, len(m))
	for k, v := range m {
		_, ok := omitKeys[k]
		if !ok {
			ret[k] = v
		}
	}
	return ret
}

// MapCopyExcludeKeys returns a new map with omitting the specified keys.
// Deprecated: Use MapOmit instead.
func MapCopyExcludeKeys[K comparable, V any, M ~map[K]V](m M, keys ...K) M {
	return MapOmitCopy(m, keys...)
}
