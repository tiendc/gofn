![status](https://github.com/tiendc/gofn/actions/workflows/master.yml/badge.svg)

# gofn - Utility functions for Go 1.18+

This is a collection of generics utility functions for Go 1.18+.

## Why generics?

This library avoids using `reflect` on determined types, such as a slice of `int` or a map of `string` and `struct`.
With common determined types, using generics will bring us more performance than using `reflect`.

## Installation

```shell
go get github.com/tiendc/gofn
```

## Usage


### Functions for slices

---

#### gofn.Equal

Returns true if 2 slices have the same size and all elements of them equal in the current order.
This function is equivalent to `reflect.DeepEqual()`, but faster (for how much faster, see Benchmark section).

```go
gofn.Equal([]int{1, 2, 3}, []int{1, 2, 3}) // true
gofn.Equal([]int{1, 2, 3}, []int{3, 2, 1}) // false

// Use gofn.EqualPred for custom equal comparison
gofn.EqualPred([]string{"one", "TWO"}, []string{"ONE", "two"}, func(elem1, elem2 string) bool {
    return strings.ToLower(elem1) == strings.ToLower(elem2)
}) // true
```

#### gofn.ContentEqual

Returns true if 2 slices equal regardless of the order of elements in the slices.

```go
gofn.ContentEqual([]int{1, 2, 3}, []int{2, 1, 3})       // true
gofn.ContentEqual([]int{1, 2, 2, 3}, []int{2, 1, 3})    // true
gofn.ContentEqual([]int{1, 2, 2, 3}, []int{1, 1, 2, 3}) // true

// Use gofn.ContentEqualPred for custom key function
gofn.ContentEqualPred([]string{"one", "TWO"}, []string{"two", "ONE"}, func(elem string) string {
    return strings.ToLower(elem)
}) // true
```

#### gofn.Concat

Concatenates two or more slices.

```go
gofn.Concat([]int{1}, []int{2}, []int{2, 3}) // []int{1, 2, 2, 3}
```

#### gofn.Contain

Returns true if a slice contains a value.

```go
gofn.Contain([]int{1, 2, 3}, 2) // true
gofn.Contain([]int{1, 2, 3}, 0) // false

// Use gofn.ContainPred for custom function
gofn.ContainPred([]string{"one", "TWO"}, func(elem string) bool {
    return strings.ToLower(elem) == "two"
}) // true
```

#### gofn.ContainAll

Returns true if a slice contains all given values.

```go
gofn.ContainAll([]int{1, 2, 3, 4, 5}, 2, 4, 3) // true
gofn.ContainAll([]int{1, 2, 3, 4, 5}, 2, 7)    // false
```

#### gofn.ContainAny

Returns true if a slice contains any of the given values.

```go
gofn.ContainAny([]int{1, 2, 3, 4, 5}, 2, 4, 7) // true
gofn.ContainAny([]int{1, 2, 3, 4, 5}, 7, 8, 9) // false
```

#### gofn.IndexOf

Returns the index of an element in a slice if found, -1 otherwise.

```go
gofn.IndexOf([]int{1, 2, 3}, 4) // -1
gofn.IndexOf([]int{1, 2, 3}, 2) // 1

// Use gofn.IndexOfPred for custom function
gofn.IndexOfPred([]string{"one", "TWO"}, func(elem string) bool {
    return strings.ToLower(elem) == "two"
}) // 1
```

#### gofn.LastIndexOf

Returns the last index of an element in a slice if found, -1 otherwise.

```go
gofn.LastIndexOf([]int{1, 2, 3}, 4)    // -1
gofn.LastIndexOf([]int{1, 2, 1, 3}, 1) // 2
```

#### gofn.RemoveAt

Removes an element at the specified index.

```go
s := []int{1, 2, 3}
gofn.RemoveAt(&s, 1)    // s == []int{1, 3}
```

#### gofn.FastRemoveAt

Removes an element at the specified index by swapping it with the last element of the slice.
This function is fast as it doesn't cause copying of slice content.

```go
s := []int{1, 2, 3}
gofn.RemoveAt(&s, 0)    // s == []int{3, 2}
```

#### gofn.Remove

Removes a value from a slice.

```go
s := []int{1, 2, 3}
gofn.Remove(&s, 1)    // s == []int{2, 3}
```

#### gofn.FastRemove

Removes a value from a slice by swapping it with the last element of the slice.

```go
s := []int{1, 2, 3}
gofn.FastRemove(&s, 1)    // s == []int{3, 2}
```

#### gofn.RemoveLastOf

Removes last occurrence of a value from a slice.

```go
s := []int{1, 2, 1, 3}
gofn.RemoveLastOf(&s, 1)    // s == []int{1, 2, 3}
```

#### gofn.FastRemoveLastOf

Removes last occurrence of a value from a slice by swapping it with the last element of the slice.

```go
s := []int{1, 2, 1, 3, 4}
gofn.FastRemoveLastOf(&s, 1)    // s == []int{1, 2, 4, 3}
```

#### gofn.RemoveAll

Removes all occurrences of a value from a slice.

```go
s := []int{1, 2, 1, 3, 1}
gofn.RemoveAll(&s, 1)    // s == []int{2, 3}
```

#### gofn.Replace

Replaces first occurrence of a value with another value.

```go
gofn.Replace([]int{1, 2, 1, 3, 1}, 1, 11)    // []int{11, 2, 1, 3, 1}
```

#### gofn.ReplaceAll

Replaces all occurrences of a value with another value.

```go
gofn.ReplaceAll([]int{1, 2, 1, 3, 1}, 1, 11)    // []int{11, 2, 11, 3, 11}
```

#### gofn.Fill

Fills a slice with specified value.

```go
s := make([]int, 5)
gofn.Fill(s, 1)  // s == []int{1, 1, 1, 1, 1}

s2 := s[2:4]
gofn.Fill(s2, 1) // s2 == []int{1, 1}, s == []int{0, 0, 1, 1, 0}
```

#### gofn.CountValue

Returns the number of occurrences of a value in a slice.

```go
gofn.CountValue([]int{1, 2, 3}, 4)    // 0
gofn.CountValue([]int{1, 2, 3, 2}, 2) // 2
```

#### gofn.ContainSlice

Returns true if a slice contains another slice.

```go
gofn.ContainSlice([]int{1, 2, 3, 4, 5}, []int{2, 3, 4}) // true
gofn.ContainSlice([]int{1, 2, 3, 4, 5}, []int{2, 4})    // false
```

#### gofn.IndexOfSlice

Finds the first occurrence of a sub-slice in a slice.

```go
gofn.IndexOfSlice([]int{1, 2, 3, 4, 5}, []int{2, 3, 4}) // 1
gofn.IndexOfSlice([]int{1, 2, 3, 4, 5}, []int{2, 4})    // -1
```

#### gofn.LastIndexOfSlice

Finds the last occurrence of a sub-slice in a slice.

```go
gofn.LastIndexOfSlice([]int{1, 2, 3, 1, 2, 3, 4}, []int{1, 2, 3}) // 3
gofn.LastIndexOfSlice([]int{1, 2, 3, 4, 5}, []int{2, 4})          // -1
```

#### gofn.GetFirst

Returns the first element of slice if slice is not empty, otherwise return the default value.

```go
gofn.GetFirst([]int{1, 2, 3}, 4) // 1
gofn.GetFirst([]int{}, 11)       // 11
```

#### gofn.GetLast

Returns the last element of slice if slice is not empty, otherwise return the default value.

```go
gofn.GetLast([]int{1, 2, 3}, 4) // 3
gofn.GetLast([]int{}, 11)       // 11
```

### Functions for maps

---

#### gofn.MapEqual

Returns true if 2 maps equal.
This function is equivalent to `reflect.DeepEqual()`, but faster (for how much faster, see Benchmark section).

```go
gofn.MapEqual(map[int]string{1: "one", 2: "two"}, map[int]string{2: "two", 1: "one"}) // true
gofn.MapEqual(map[int]string{1: "one", 2: "two"}, map[int]string{1: "one"})           // false
```

#### gofn.MapContainKeys

Returns true if a map contains one or more keys.

```go
gofn.MapContainKeys(map[int]int{1: 11, 2: 22}, 1)    // true
gofn.MapContainKeys(map[int]int{1: 11, 2: 22}, 1, 2) // true
gofn.MapContainKeys(map[int]int{1: 11, 2: 22}, 1, 3) // false
```

#### gofn.MapContainValues

Returns true if a map contains one or more values.

```go
gofn.MapContainValues(map[int]int{1: 11, 2: 22}, 11)     // true
gofn.MapContainValues(map[int]int{1: 11, 2: 22}, 11, 22) // true
gofn.MapContainValues(map[int]int{1: 11, 2: 22}, 11, 33) // false
```

#### gofn.MapKeys

Returns a slice of keys of a map.

```go
gofn.MapKeys(map[int]int{1: 11, 2: 22}) // []int{1, 2} (note: values may be in different order)
```

#### gofn.MapValues

Returns a slice of values of a map.

```go
gofn.MapValues(map[int]int{1: 11, 2: 22, 3: 22}) // []int{11, 22, 22} (note: values may be in different order)
```

#### gofn.MapEntries

Returns a slice of entries (key, value) of a map.

```go
gofn.MapEntries(map[int]int{1: 11, 2: 22}) // []*Tuple2[int,int]{{1,11}, {2,22}} (note: values may be in different order)
```

#### gofn.MapUpdate

Update a map content with another map.

```go
s := map[int]int{1: 11, 2: 22}
gofn.MapUpdate(s, map[int]int{3: 33}) // s == map[int]int{1: 11, 2: 22, 3: 33}
```

#### gofn.MapGet

Retrieves value of a key, returns the default value if not exists.

```go
gofn.MapGet(map[int]int{1: 11, 2: 22}, 1, 0) // 11 (found)
gofn.MapGet(map[int]int{1: 11, 2: 22}, 3, 0) // 0 (not found)
```

#### gofn.MapPop

Removes entry from a map and returns the current value if found.

```go
gofn.MapPop(map[int]int{1: 11, 2: 22}, 1, 0) // 11 (found)
gofn.MapPop(map[int]int{1: 11, 2: 22}, 3, 0) // 0 (not found)
```

#### gofn.MapSetDefault

Sets default value for a key and returns the current value.

```go
gofn.MapSetDefault(map[int]int{1: 11, 2: 22}, 1, 0) // 11 (no value added to the map)
gofn.MapSetDefault(map[int]int{1: 11, 2: 22}, 3, 0) // 0 (entry [3, 0] is added to the map)
```

### Transformation functions

---

#### gofn.Filter / FilterXX

Filter a slice with condition.

```go
gofn.Filter([]int{1, 2, 3, 4}, func (i int) bool {
    return i % 2 == 0
}) // []int{2, 4}


gofn.FilterLT([]int{1, 2, 3, 4}, 3)        // []int{1, 2}
gofn.FilterLTE([]int{1, 2, 3, 4}, 3)       // []int{1, 2, 3}
gofn.FilterGT([]int{1, 2, 3, 4}, 3)        // []int{4}
gofn.FilterGTE([]int{1, 2, 3, 4}, 3)       // []int{3, 4}
gofn.FilterNE([]int{1, 2, 3, 4}, 3)        // []int{1, 2, 4}
gofn.FilterIN([]int{1, 2, 3, 4}, 3, 2, 7)  // []int{2, 3}
gofn.FilterNIN([]int{1, 2, 3, 4}, 3, 2, 7) // []int{1, 4}
```

#### gofn.ToSet

Returns a slice of unique values for the given slice.

```go
gofn.ToSet([]int{1, 2, 3, 1, 2})        // []int{1, 2, 3}
gofn.ToSet([]string{"one", "2", "one"}) // []string{"one", "2"}

// Use gofn.ToSetPred for custom key function
ToSetPred([]string{"one", "TWO", "two", "One"}, func(elem string) string {
    return strings.ToLower(elem)
}) // []string{"one", "TWO"}
```

#### gofn.MapSlice

Transforms a slice to a slice.

```go
gofn.MapSlice([]int{1, 2, 3}, func (i int) string {
    return fmt.Sprintf("%d", i)
}) // []string{"1", "2", "3"}


// Use gofn.MapSliceEx for combining both transformation and filter
gofn.MapSliceEx([]int{1, 2, 3, -1}, func (i int) (string, bool) {
    if i < 0 {
        return "", false
    }
    return fmt.Sprintf("%d", i), true
}) // []string{"1", "2", "3"}
```

#### gofn.MapSliceToMap

Transforms a slice to a map.

```go
gofn.MapSliceToMap([]int{1, 2, 3}, func (i int) (int, string) {
    return i, fmt.Sprintf("%d", i)
}) // map[int]string{1: "1", 2: "2", 3: "3"}

// Use gofn.MapSliceToMapEx for combining both transformation and filter
gofn.MapSliceToMapEx([]int{1, 2, 3, -1}, func (i int) (int, string, bool) {
    if i < 0 {
        return 0, "", false
    }
    return i, fmt.Sprintf("%d", i), true
}) // map[int]string{1: "1", 2: "2", 3: "3"}
```

#### gofn.Chunk / ChunkByPieces

Splits a slice into chunks.

```go
gofn.Chunk([]int{1, 2, 3, 4, 5}, 2)         // [][]int{[]int{1, 2}, []int{3, 4}, []int{5}}
gofn.ChunkByPieces([]int{1, 2, 3, 4, 5}, 2) // [][]int{[]int{1, 2, 3}, []int{4, 5}}
```

#### gofn.Reverse

Reverses a slice.

```go
gofn.Reverse([]int{1, 2, 3}) // []int{3, 2, 1}

s1 := []int{1, 2, 3}
s2 := gofn.ReverseCopy(s1)  // s1 == []int{1, 2, 3}, s2 == []int{3, 2, 1}
```

#### gofn.Flatten

Flattens multi-dimensional slice.

```go
gofn.Flatten([]int{1, 2, 3}, []int{4, 5})                    // []int{1, 2, 3, 4, 5}
gofn.Flatten3([][]int{{1, 2}, {3, 4}, [][]int{{5, 6}, {7}}}) // []int{1, 2, 3, 4, 5, 6, 7}
```

#### gofn.Zip

Combines values from multiple slices by each position.

```go
gofn.Zip([]int{1, 2, 3}, []int{4, 5})               // []*Tuple2{{1, 4), {2, 5}}
gofn.Zip3([]int{1, 2, 3}, []int{4, 5}, []int{6, 7}) // []*Tuple3{{1, 4, 6), {2, 5, 7}}
```

### Conversion functions

---

#### gofn.ToIntfSlice

Converts a slice of any type to a slice of interfaces.

```go
gofn.ToIntfSlice([]int{1, 2, 3})         // []interface{}{1, 2, 3}
gofn.ToIntfSlice([]string{"foo", "bar"}) // []interface{}{"foo", "bar"}
```

#### gofn.ToStringSlice

Converts a slice of string-approximate type to a slice of strings.

```go
type XType string
gofn.ToStringSlice([]XType{XType("foo"), XType("bar")}) // []string{"foo", "bar"}
```

#### gofn.ToNumberSlice

Converts a slice of number type to a slice of specified number type.

```go
gofn.ToNumberSlice[int]([]int8{1, 2, 3})    // []int{1, 2, 3}
gofn.ToNumberSlice[float32]([]int{1, 2, 3}) // []float32{1.0, 2.0, 3.0}

type XType int
gofn.ToNumberSlice[int]([]XType{XType(1), XType(2)}) // []int{1, 2}
```

#### gofn.ToSlice

Returns a slice for individual values.

```go
gofn.ToSlice(1, 2, 3) // []int{1, 2, 3}
```

### Common functions

---

#### gofn.ForEach

Iterates through each element of a slice.

```go
gofn.ForEach([]int{1, 2, 3}, func (i, v int) {
    fmt.Printf("%d ", v)
}) // prints 1 2 3
```

#### gofn.ForEachReverse

Iterates through each element of a slice from the end.

```go
gofn.ForEachReverse([]int{1, 2, 3}, func (i, v int) {
    fmt.Printf("%d ", v)
}) // prints 3 2 1
```

#### gofn.If

Provides a similar function to C ternary operator (cond ? expr_a : expr_b).

```go
gofn.If(direction > 0, 1, -1) // 1 if direction > 0, -1 if direction <= 0
```

NOTE: don't use this function this way:

```go
firstElem := gofn.If(len(slice) > 0, slice[0], defaultVal)
```

Unlike C ternary operator (cond ? expr_a : expr_b), both expressions are evaluated by gofn.If().
So if the slice is empty, the above statement will cause panic when the code tries
to evaluate the first expression.


#### gofn.Must

gofn.Must accepts any function that returns 2 values and the 2nd value is error type.
gofn.Must returns the 1st value if the 2nd is `nil`, otherwise it panics.

```go
func CalculateAmount() (int, error) {}

amount := gofn.Must(CalculateAmount()) // panic on error, otherwise returns the amount
```

#### gofn.New

Creates a new variable and return the address of it. Very helpful in unit testing.

```go
func f(ptr *int) {}

// Normally we need to declare a var before accessing its address
val := 10
f(&val)

// With using gofn.New
f(gofn.New(10))
```

### Sorting functions

---

#### gofn.Sort / SortDesc

Convenient wrapper of the built-in `sort.Slice`.

```go
gofn.Sort([]int{1, 3, 2})         // []int{1, 2, 3}
gofn.SortDesc([]int{1, 3, 2})     // []int{3, 2, 1}
gofn.IsSorted([]int{1, 3, 2})     // false
gofn.IsSortedDesc([]int{3, 2, 1}) // true
```

### Union / Intersection / Difference functions

---

#### gofn.Union

Finds all unique values from multiple slices.

```go
gofn.Union([]int{1, 3, 2}, []int{1, 2, 2, 4}) // []int{1, 3, 2, 4}
```

#### gofn.Intersection

Finds all unique shared values from multiple slices.

```go
gofn.Intersection([]int{1, 3, 2}, []int{1, 2, 2, 4}) // []int{1, 2}
```

#### gofn.Difference

Finds all different values from 2 slices.

```go
left, right := gofn.Difference([]int{1, 3, 2}, []int{2, 2, 4}) // left == []int{1, 3}, right == []int{4}
```

### Other functions

---

#### gofn.Sum / SumAs

Calculates `sum` value of a slice.

```go
gofn.Sum([]int{1, 2, 3})            // 6
gofn.SumAs[int]([]int8{50, 60, 70}) // 180 (Sum() will fail as the result overflows int8)
```

#### gofn.Product / ProductAs

Calculates `product` value of a slice.

```go
gofn.Product([]int{1, 2, 3})         // 6
gofn.ProductAs[int]([]int8{5, 6, 7}) // 210 (Product() will fail as the result overflows int8)
```

#### gofn.Reduce / ReduceEx

Reduces a slice to a value.

```go
gofn.Reduce([]int{1, 2, 3}, func (accumulator int, currentValue int) int {
    return accumulator + currentValue
}) // 6
```

#### gofn.Min / Max

Find minimum/maximum value in a slice.

```go
gofn.Min(1, 2, 3, -1)    // -1
gofn.Max(1, 2, 3, -1)    // 3
gofn.MinMax(1, 2, 3, -1) // -1, 3
```

#### gofn.MinTime / MaxTime

Find minimum/maximum time value in a slice.

```go
t0 := time.Time{}
t1 := time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC)
t2 := time.Date(2000, time.December, 2, 0, 0, 0, 0, time.UTC)
gofn.MinTime(t0, t1, t2) // t0
gofn.MinTime(t1, t2)     // t1
gofn.MaxTime(t0, t1, t2) // t2
```

#### gofn.Abs

Calculates absolute value of an integer.

```go
gofn.Abs(-123)          // int64(123)
gofn.Abs(123)           // int64(123)
gofn.Abs(math.MinInt64) // math.MinInt64 (special case)
```

#### gofn.RandString

Calculates absolute value of an integer.

```go
gofn.RandString(10)                     // a random string has 10 characters (default of alphabets and digits)
gofn.RandStringEx(10, []rune("01234"))  // a random string has 10 characters (only 0-4)
```

## Benchmark

#### gofn.Equal vs gofn.ContentEqual vs reflect.DeepEqual
___

```
Benchmark_Slice_Equal/StructSlice/Equal
Benchmark_Slice_Equal/StructSlice/Equal-8         	510845715	         2.047 ns/op
Benchmark_Slice_Equal/StructSlice/ContentEqual
Benchmark_Slice_Equal/StructSlice/ContentEqual-8  	583167950	         2.061 ns/op
Benchmark_Slice_Equal/StructSlice/DeepEqual
Benchmark_Slice_Equal/StructSlice/DeepEqual-8     	15403771	         79.19 ns/op

Benchmark_Slice_Equal/IntSlice/Equal
Benchmark_Slice_Equal/IntSlice/Equal-8            	589706185	         2.087 ns/op
Benchmark_Slice_Equal/IntSlice/ContentEqual
Benchmark_Slice_Equal/IntSlice/ContentEqual-8     	523120755	         2.194 ns/op
Benchmark_Slice_Equal/IntSlice/DeepEqual
Benchmark_Slice_Equal/IntSlice/DeepEqual-8        	15243183	         77.93 ns/op
```

## Contributing

- You are welcome to make pull requests for new functions and bug fixes.

## Authors

- Dao Cong Tien (tiendc/tiendc.vn)
