[![Go Version][gover-img]][gover] [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![GoReport][rpt-img]][rpt]

# gofn - Utility functions for Go 1.18+

This is a collection of generics utility functions for Go 1.18+.

## Why generics?

This library avoids using `reflect` on determined types, such as a slice of `int` or a map of `string` and `struct`.
With common determined types, using generics will bring back more performance than using `reflect`.

## Functionalities

`gofn` consists of useful and convenient functions for most common needs when working on slices, maps, transformation, conversion, and so on.

**NOTE**: This lib does not provide deep-copy function as that will require using `reflect`. If you need the function, try [go-deepcopy](https://github.com/tiendc/go-deepcopy).

## Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Functions for slices](#functions-for-slices)
  - [Functions for maps](#functions-for-maps)
  - [Functions for strings](#functions-for-strings)
  - [Functions for numbers](#functions-for-numbers)
  - [Transformation functions](#transformation-functions)
  - [Conversion functions](#conversion-functions)
  - [Common functions](#common-functions)
  - [Specific algo functions](#specific-algo-functions)
  - [Other functions](#other-functions)
- [Benchmarks](#benchmarks)
- [Contributing](#contributing)
- [Authors](#authors)
- [License](#license)

## Installation

```shell
go get github.com/tiendc/gofn
```

## Usage

### Functions for slices

---

#### Equal

Returns true if 2 slices have the same size and all elements of them equal in the current order.
This function is equivalent to `reflect.DeepEqual()`, but faster (for how much faster, see Benchmark section).

```go
Equal([]int{1, 2, 3}, []int{1, 2, 3}) // true
Equal([]int{1, 2, 3}, []int{3, 2, 1}) // false

// Use EqualPred for custom equal comparison
EqualPred([]string{"one", "TWO"}, []string{"ONE", "two"}, strings.EqualFold) // true
```

#### ContentEqual

Returns true if 2 slices have the same size and contents equal regardless of the order of elements.

```go
ContentEqual([]int{1, 2, 3}, []int{2, 1, 3})       // true
ContentEqual([]int{1, 2, 2, 3}, []int{2, 1, 3})    // false
ContentEqual([]int{1, 2, 2, 3}, []int{2, 1, 2, 3}) // true
ContentEqual([]int{1, 2, 2, 3}, []int{1, 1, 2, 3}) // false

// Use ContentEqualPred for custom key function
ContentEqualPred([]string{"one", "TWO"}, []string{"two", "ONE"}, strings.ToLower) // true
```

#### Concat

Concatenates two or more slices.

```go
Concat([]int{1}, []int{2}, []int{2, 3}) // []int{1, 2, 2, 3}
```

#### Contain

Returns true if a slice contains a value.

```go
Contain([]int{1, 2, 3}, 2) // true
Contain([]int{1, 2, 3}, 0) // false

// Use ContainPred for custom function
ContainPred([]string{"one", "TWO"}, func(elem string) bool {
    return strings.ToLower(elem) == "two"
}) // true
```

#### ContainAll

Returns true if a slice contains all given values.

```go
ContainAll([]int{1, 2, 3, 4, 5}, 2, 4, 3) // true
ContainAll([]int{1, 2, 3, 4, 5}, 2, 7)    // false
```

#### ContainAny

Returns true if a slice contains any of the given values.

```go
ContainAny([]int{1, 2, 3, 4, 5}, 2, 4, 7) // true
ContainAny([]int{1, 2, 3, 4, 5}, 7, 8, 9) // false
```

#### IsUnique

Returns true if a slice contains unique values.

```go
IsUnique([]int{1, 2, 3}) // true
IsUnique([]int{1, 2, 1}) // false

// Use IsUniquePred for custom function
IsUniquePred([]string{"one", "ONE"}, strings.ToLower) // false
```

#### IndexOf

Finds the index of a value in a slice, returns -1 if not found.

```go
IndexOf([]int{1, 2, 3}, 4) // -1
IndexOf([]int{1, 2, 3}, 2) // 1

// Use IndexOfPred for custom function
IndexOfPred([]string{"one", "TWO"}, func(elem string) bool {
    return strings.ToLower(elem) == "two"
}) // 1
```

#### LastIndexOf

Finds the last index of an element in a slice, returns -1 if not found.

```go
LastIndexOf([]int{1, 2, 3}, 4)    // -1
LastIndexOf([]int{1, 2, 1, 3}, 1) // 2
```

#### RemoveAt

Removes element at the specified index.

```go
s := []int{1, 2, 3}
RemoveAt(&s, 1) // s == []int{1, 3}
```

#### FastRemoveAt

Removes element at the specified index by swapping it with the last element of the slice.
This function is fast as it doesn't cause copying of slice content.

```go
s := []int{1, 2, 3, 4}
FastRemoveAt(&s, 1) // s == []int{1, 4, 3} (2 and 4 are exchanged)
```

#### Remove

Removes a value from a slice.

```go
s := []int{1, 2, 3}
Remove(&s, 1) // s == []int{2, 3}
```

#### FastRemove

Removes a value from a slice by swapping it with the last element of the slice.

```go
s := []int{1, 2, 3, 4}
FastRemove(&s, 2) // s == []int{1, 4, 3} (2 and 4 are exchanged)
```

#### RemoveLastOf

Removes last occurrence of a value from a slice.

```go
s := []int{1, 2, 1, 3}
RemoveLastOf(&s, 1) // s == []int{1, 2, 3}
```

#### FastRemoveLastOf

Removes last occurrence of a value from a slice by swapping it with the last element of the slice.

```go
s := []int{1, 2, 1, 3, 4}
FastRemoveLastOf(&s, 1) // s == []int{1, 2, 4, 3} (1 and 4 are exchanged)
```

#### RemoveAll

Removes all occurrences of a value from a slice.

```go
s := []int{1, 2, 1, 3, 1}
RemoveAll(&s, 1) // s == []int{2, 3}
```

#### Replace

Replaces first occurrence of a value with another value.

```go
Replace([]int{1, 2, 1, 3, 1}, 1, 11) // []int{11, 2, 1, 3, 1}
```

#### ReplaceAll

Replaces all occurrences of a value with another value.

```go
ReplaceAll([]int{1, 2, 1, 3, 1}, 1, 11) // []int{11, 2, 11, 3, 11}
```

#### Fill

Fills a slice with specified value.

```go
s := make([]int, 5)
Fill(s, 1)  // s == []int{1, 1, 1, 1, 1}

s2 := s[2:4]
Fill(s2, 1) // s2 == []int{1, 1}, s == []int{0, 0, 1, 1, 0}
```

#### CountValue

Counts the number of occurrences of a value in a slice.

```go
CountValue([]int{1, 2, 3}, 4)    // 0
CountValue([]int{1, 2, 3, 2}, 2) // 2
```

#### ContainSlice

Returns true if a slice contains another slice.

```go
ContainSlice([]int{1, 2, 3, 4, 5}, []int{2, 3, 4}) // true
ContainSlice([]int{1, 2, 3, 4, 5}, []int{2, 4})    // false
```

#### IndexOfSlice

Finds the first occurrence of a sub-slice in a slice.

```go
IndexOfSlice([]int{1, 2, 3, 4, 5}, []int{2, 3, 4}) // 1
IndexOfSlice([]int{1, 2, 3, 4, 5}, []int{2, 4})    // -1
```

#### LastIndexOfSlice

Finds the last occurrence of a sub-slice in a slice.

```go
LastIndexOfSlice([]int{1, 2, 3, 1, 2, 3, 4}, []int{1, 2, 3}) // 3
LastIndexOfSlice([]int{1, 2, 3, 4, 5}, []int{2, 4})          // -1
```

#### GetFirst

Returns the first element of slice if it is not empty, otherwise return the default value.

```go
GetFirst([]int{1, 2, 3}, 4) // 1
GetFirst([]int{}, 11)       // 11
```

#### GetLast

Returns the last element of slice if it is not empty, otherwise return the default value.

```go
GetLast([]int{1, 2, 3}, 4) // 3
GetLast([]int{}, 11)       // 11
```

### Functions for maps

---

#### MapEqual

Returns true if 2 maps equal.
This function is equivalent to `reflect.DeepEqual()`, but faster (for how much faster, see Benchmark section).

```go
MapEqual(map[int]string{1: "one", 2: "two"}, map[int]string{2: "two", 1: "one"}) // true
MapEqual(map[int]string{1: "one", 2: "two"}, map[int]string{1: "one", 2: "TWO"}) // false
```

#### MapContainKeys

Returns true if a map contains all given keys.

```go
MapContainKeys(map[int]int{1: 11, 2: 22}, 1)    // true
MapContainKeys(map[int]int{1: 11, 2: 22}, 1, 2) // true
MapContainKeys(map[int]int{1: 11, 2: 22}, 1, 3) // false
```

#### MapContainValues

Returns true if a map contains all given values.

```go
MapContainValues(map[int]int{1: 11, 2: 22}, 11)     // true
MapContainValues(map[int]int{1: 11, 2: 22}, 11, 22) // true
MapContainValues(map[int]int{1: 11, 2: 22}, 11, 33) // false
```

#### MapKeys

Gets all keys of a map.

```go
MapKeys(map[int]int{1: 11, 2: 22}) // []int{1, 2} (note: values may be in different order)
```

#### MapValues

Gets all values of a map.

```go
MapValues(map[int]int{1: 11, 2: 22, 3: 22}) // []int{11, 22, 22} (note: values may be in different order)
```

#### MapEntries

Gets all entries (key, value) of a map.

```go
MapEntries(map[int]int{1: 11, 2: 22}) // []*Tuple2[int,int]{{1,11}, {2,22}} (note: values may be in different order)
```

#### MapUpdate/MapUpdateXX

Updates a map content with another map.

```go
s := map[int]int{1: 11, 2: 22}
MapUpdate(s, map[int]int{1: 111, 3: 33})             // s == map[int]int{1: 111, 2: 22, 3: 33}
MapUpdateExistingOnly(s, map[int]int{2: 222, 3: 33}) // s == map[int]int{1: 11, 2: 222}
MapUpdateNewOnly(s, map[int]int{2: 222, 3: 33})      // s == map[int]int{1: 11, 2: 22, 3: 33}
```

#### MapGet

Retrieves map value for a key, returns the default value if not exist.

```go
MapGet(map[int]int{1: 11, 2: 22}, 1, 0) // 11 (found)
MapGet(map[int]int{1: 11, 2: 22}, 3, 0) // 0 (not found)
```

#### MapPop

Removes entry from a map and returns the current value if found.

```go
MapPop(map[int]int{1: 11, 2: 22}, 1, 0) // 11 (found)
MapPop(map[int]int{1: 11, 2: 22}, 3, 0) // 0 (not found)
```

#### MapSetDefault

Sets default value for a key and returns the current value.

```go
MapSetDefault(map[int]int{1: 11, 2: 22}, 1, 0) // 11 (no value added to the map)
MapSetDefault(map[int]int{1: 11, 2: 22}, 3, 0) // 0 (entry [3, 0] is added to the map)
```

### Functions for strings

---

#### MultilineString

Removes all leading spaces from every line in the given string. This function is useful to declare a string with a neat multiline-style.

```go
func DoSomething() {
	// Commonly you may use this style to create multiline string in Go (which looks ugly)
	s := `
line-1 abc xyz
line-2 abc xyz
`
	// Use this function
	s := MultilineString(
		`line-1 abc xyz
		line-2 abc xyz`
	)
}
```

#### LinesTrim/LinesTrimSpace

Removes all certain leading and trailing characters from every line in the given string.

```go
LinesTrimSpace("  line-1  \n  line-2  ")      // "line-1\nline-2"
LinesTrim("a line-1 b \n a line-2 ab", " ba") // "line-1\nline2"
```

#### LinesTrimLeft/LinesTrimLeftSpace

Removes all certain leading characters from every line in the given string.

```go
LinesTrimLeftSpace("  line-1  \n  line-2  ")      // "line-1  \nline-2  "
LinesTrimLeft("ab line-1  \n a line-2 ab", " ba") // "line-1  \nline2 ab"
```

#### LinesTrimRight/LinesTrimRightSpace

Removes all certain trailing characters from every line in the given string.

```go
LinesTrimRightSpace("  line-1  \n  line-2  ")    // "  line-1\nline-2"
LinesTrimRight("line-1 b \n a line-2 ab", " ba") // "line-1\n a line2"
```

### Functions for numbers

---

#### ParseInt/ParseIntXX

Parses integer using **strconv.ParseInt** then converts the value to a specific type.

```go
ParseInt[int16]("111")            // int16(111)
ParseInt[int8]("128")             // strconv.ErrRange

// Return default value on failure
ParseIntDef("200", 10)            // int(200)
ParseIntDef("200", int8(10))      // int8(10)

// Parse integer with specific base
ParseIntEx[int8]("eeff1234", 16)  // strconv.ErrRange
ParseIntEx[int]("eeff1234", 16)   // int value for "eeff1234"

// Parse string containing commas
ParseInt[int]("1,234,567")        // strconv.ErrSyntax
ParseIntUngroup[int]("1,234,567") // int(1234567)
```

- **NOTE**: There are also **ParseUint** for unsigned integers and **ParseFloat** for floating numbers.

#### FormatInt/FormatIntXX

Formats an integer.

```go
FormatInt(123)            // "123"

// Format number with specific format string (use fmt.Sprintf)
FormatIntEx(123, "%05d")  // "00123"

// Format number with decimal grouping
FormatIntGroup(1234567)   // 1,234,567
```

- **NOTE**: There are also **FormatUint** for unsigned integers and **FormatFloat** for floating numbers.

### Transformation functions

---

#### Filter / FilterXX

Filters a slice with condition.

```go
Filter([]int{1, 2, 3, 4}, func (i int) bool {
    return i % 2 == 0
}) // []int{2, 4}


FilterLT([]int{1, 2, 3, 4}, 3)        // []int{1, 2}
FilterLTE([]int{1, 2, 3, 4}, 3)       // []int{1, 2, 3}
FilterGT([]int{1, 2, 3, 4}, 3)        // []int{4}
FilterGTE([]int{1, 2, 3, 4}, 3)       // []int{3, 4}
FilterNE([]int{1, 2, 3, 4}, 3)        // []int{1, 2, 4}
FilterIN([]int{1, 2, 3, 4}, 3, 2, 7)  // []int{2, 3}
FilterNIN([]int{1, 2, 3, 4}, 3, 2, 7) // []int{1, 4}
FilterLIKE([]string{"*Abc*", "*abc*", "abc*", "*abc"}, "Abc")  // []string{"*Abc*"}
FilterILIKE([]string{"*Abc*", "*abc*", "abc*", "*abc"}, "Abc") // []string{"*Abc*", "*abc*", "abc*", "*abc"}
```

#### ToSet

Calculates unique values of a slice.

```go
ToSet([]int{1, 2, 3, 1, 2})        // []int{1, 2, 3}
ToSet([]string{"one", "2", "one"}) // []string{"one", "2"}

// Use ToSetPred for custom key function
ToSetPred([]string{"one", "TWO", "two", "One"}, strings.ToLower) // []string{"one", "TWO"}
```

#### MapSlice/MapSliceEx

Transforms a slice to a slice.

```go
MapSlice([]string{"1", "2 ", " 3"}, strings.TrimSpace)   // []string{"1", "2", "3"}

// Use MapSliceEx to transform with error handling
MapSliceEx([]string{"1","2","3"}, gofn.ParseInt[int])    // []int{1, 2, 3}
MapSliceEx([]string{"1","x","3"}, gofn.ParseInt[int])    // strconv.ErrSyntax
MapSliceEx([]string{"1","200","3"}, gofn.ParseInt[int8]) // strconv.ErrRange
```

#### MapSliceToMap/MapSliceToMapEx

Transforms a slice to a map.

```go
MapSliceToMap([]int{1, 2, 3}, func (i int) (int, string) {
    return i, fmt.Sprintf("%d", i)
}) // map[int]string{1: "1", 2: "2", 3: "3"}

// Use MapSliceToMapEx to transform with error handling
MapSliceToMapEx([]string{"1","300","3"}, func (s string) (string, int, bool) {
    v, e := gofn.ParseInt[int8](s)
    return s, v, e
}) // strconv.ErrRange
```

#### Chunk / ChunkByPieces

Splits slice content into chunks.

```go
Chunk([]int{1, 2, 3, 4, 5}, 2)         // [][]int{[]int{1, 2}, []int{3, 4}, []int{5}}
ChunkByPieces([]int{1, 2, 3, 4, 5}, 2) // [][]int{[]int{1, 2, 3}, []int{4, 5}}
```

#### Reverse

Reverses slice content.

```go
Reverse([]int{1, 2, 3}) // []int{3, 2, 1}

s1 := []int{1, 2, 3}
s2 := ReverseCopy(s1)  // s1 == []int{1, 2, 3}, s2 == []int{3, 2, 1}
```

#### Flatten

Flattens multi-dimension slice.

```go
Flatten([]int{1, 2, 3}, []int{4, 5})                    // []int{1, 2, 3, 4, 5}
Flatten3([][]int{{1, 2}, {3, 4}, [][]int{{5, 6}, {7}}}) // []int{1, 2, 3, 4, 5, 6, 7}
```

#### Zip

Combines values from multiple slices by each position.

```go
Zip([]int{1, 2, 3}, []int{4, 5})                              // []*Tuple2{{1, 4), {2, 5}}
Zip3([]int{1, 2, 3}, []string{"4", "5"}, []float32{6.0, 7.0}) // []*Tuple3{{1, "4", 6.0), {2, "5", 7.0}}
```

### Conversion functions

---

#### ToIntfSlice

Converts a slice of any type to a slice of interfaces.

```go
ToIntfSlice([]int{1, 2, 3})         // []interface{}{1, 2, 3}
ToIntfSlice([]string{"foo", "bar"}) // []interface{}{"foo", "bar"}
```

#### ToStringSlice

Converts a slice of string-approximate type to a slice of strings.

```go
type XType string
ToStringSlice([]XType{XType("foo"), XType("bar")}) // []string{"foo", "bar"}
```

#### ToNumberSlice

Converts a slice of number type to a slice of specified number type.

```go
ToNumberSlice[int]([]int8{1, 2, 3})    // []int{1, 2, 3}
ToNumberSlice[float32]([]int{1, 2, 3}) // []float32{1.0, 2.0, 3.0}

type XType int
ToNumberSlice[int]([]XType{XType(1), XType(2)}) // []int{1, 2}
```

#### ToSlice

Creates a slice for individual values.

```go
ToSlice(1, 2, 3) // []int{1, 2, 3}
```

### Common functions

---

#### ForEach

Iterates over slice content.

```go
ForEach([]int{1, 2, 3}, func (i, v int) {
    fmt.Printf("%d ", v)
}) // prints 1 2 3
```

#### ForEachReverse

Iterates over slice content from the end.

```go
ForEachReverse([]int{1, 2, 3}, func (i, v int) {
    fmt.Printf("%d ", v)
}) // prints 3 2 1
```

#### All

Returns `true` if all given values are evaluated `true`.

```go
All(1, "1", 0.5) // true
All(1, "1", 0.0) // false
All(1, "", -1)   // false
All()            // true
```

#### Any

Returns `true` if any of the given values is evaluated `true`.

```go
Any(1, "", 0.5)  // true
Any(1, "1", 0.0) // true
Any(0, "", 0.0)  // false
Any()            // false
```

#### MustN (N is from 1 to 6)

MustN functions accept a number of arguments with the last one is of `error` type.
MustN functions return the first N-1 arguments if the error is `nil`, otherwise they panic.

```go
func CalculateAmount() (int, error) {}
amount := Must(CalculateAmount()) // panic on error, otherwise returns the amount

func CalculateData() (int, string, float64, error) {}
v1, v2, v3 := Must4(CalculateData()) // panic on error, otherwise returns the 3 first values
```

### Sorting functions

---

#### Sort / SortDesc

Convenient wrapper of the built-in `sort.Slice`.

```go
Sort([]int{1, 3, 2})         // []int{1, 2, 3}
SortDesc([]int{1, 3, 2})     // []int{3, 2, 1}
IsSorted([]int{1, 3, 2})     // false
IsSortedDesc([]int{3, 2, 1}) // true
```

### Specific Algo functions

---

#### Union

Finds all unique values from multiple slices.

```go
Union([]int{1, 3, 2}, []int{1, 2, 2, 4}) // []int{1, 3, 2, 4}
```

#### Intersection

Finds all unique shared values from multiple slices.

```go
Intersection([]int{1, 3, 2}, []int{1, 2, 2, 4}) // []int{1, 2}
```

#### Difference

Finds all different values from 2 slices.

```go
left, right := Difference([]int{1, 3, 2}, []int{2, 2, 4}) // left == []int{1, 3}, right == []int{4}
```

#### Sum / SumAs

Calculates `sum` value of slice elements.

```go
Sum([]int{1, 2, 3})            // 6
SumAs[int]([]int8{50, 60, 70}) // 180 (Sum() will fail as the result overflows int8)
```

#### Product / ProductAs

Calculates `product` value of slice elements.

```go
Product([]int{1, 2, 3})         // 6
ProductAs[int]([]int8{5, 6, 7}) // 210 (Product() will fail as the result overflows int8)
```

#### Reduce / ReduceEx

Reduces a slice to a value.

```go
Reduce([]int{1, 2, 3}, func (accumulator int, currentValue int) int {
    return accumulator + currentValue
}) // 6
```

#### Min / Max

Finds minimum/maximum value in a slice.

```go
Min(1, 2, 3, -1)    // -1
Max(1, 2, 3, -1)    // 3
MinMax(1, 2, 3, -1) // -1, 3
```

#### MinTime / MaxTime

Finds minimum/maximum time value in a slice.

```go
t0 := time.Time{}
t1 := time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC)
t2 := time.Date(2000, time.December, 2, 0, 0, 0, 0, time.UTC)
MinTime(t0, t1, t2) // t0
MinTime(t1, t2)     // t1
MaxTime(t0, t1, t2) // t2
```

#### Abs

Calculates absolute value of an integer.

```go
Abs(-123)          // int64(123)
Abs(123)           // int64(123)
Abs(math.MinInt64) // math.MinInt64 (special case)
```

#### RandString

Generates a random string.

```go
RandString(10)                     // a random string has 10 characters (default of alphabets and digits)
RandStringEx(10, []rune("01234"))  // a random string has 10 characters (only 0-4)
```

### Other functions

---

#### New

Creates a new variable and return the address of it. Very helpful in unit testing.

```go
func f(ptr *int) {}

// Normally we need to declare a var before accessing its address
val := 10
f(&val)

// With using New
f(New(10))
```

#### Head

Takes the first argument.

```go
Head(1, "2", 1.0, 3) // 1
```

#### Tail

Takes the last argument.

```go
Tail[string](true, "2", 1.0, "3") // "3"
```


#### FirstTrue

Returns the first "true" value in the given arguments if found.
True values are considered:
  - not zero values (0, empty string, false, nil, ...)
  - not empty containers (slice, array, map, channel)
  - NOTE: non-nil pointers are considered "true" (even they point to zero objects)

```go
FirstTrue(0, 0, -1, 2, 3)                           // -1
FirstTrue("", "", " ", "b")                         // " "
FirstTrue([]int{}, nil, []int{1}, []int{2, 3})      // []int{1}
FirstTrue[interface{}](nil, 0, 0.0, "", struct{}{}) // nil (the first zero value)
```

## Benchmarks

#### Equal vs ContentEqual vs reflect.DeepEqual
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

- Dao Cong Tien ([tiendc](https://github.com/tiendc))

## License

- [MIT License](LICENSE)

[doc-img]: https://pkg.go.dev/badge/github.com/tiendc/gofn
[doc]: https://pkg.go.dev/github.com/tiendc/gofn
[gover-img]: https://img.shields.io/badge/Go-%3E%3D%201.18-blue
[gover]: https://img.shields.io/badge/Go-%3E%3D%201.18-blue
[ci-img]: https://github.com/tiendc/gofn/actions/workflows/go.yml/badge.svg
[ci]: https://github.com/tiendc/gofn/actions/workflows/go.yml
[cov-img]: https://codecov.io/gh/tiendc/gofn/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/tiendc/gofn
[rpt-img]: https://goreportcard.com/badge/github.com/tiendc/gofn
[rpt]: https://goreportcard.com/report/github.com/tiendc/gofn
