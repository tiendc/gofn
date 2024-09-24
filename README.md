
[![Go Version][gover-img]][gover] [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![GoReport][rpt-img]][rpt]

# gofn - Utility functions for Go 1.18+

This is a collection of generics utility functions for Go 1.18+.

## Functionalities

`gofn` consists of useful and convenient functions for most common needs when working on slices, maps, structs, transformation, conversion, and so on.

Try related libs:
- [go-deepcopy](https://github.com/tiendc/go-deepcopy): true deep copy function

## Installation

```shell
go get github.com/tiendc/gofn
```

## Function list

**Slice**
  - [Equal / EqualBy](#equal--equalby)
  - [ContentEqual / ContentEqualBy](#contentequal--contentequalby)
  - [Sort / IsSorted](#sort--issorted)
  - [RemoveAt](#removeat)
  - [FastRemoveAt](#fastremoveat)
  - [Remove](#remove)
  - [FastRemove](#fastremove)
  - [RemoveLastOf](#removelastof)
  - [FastRemoveLastOf](#fastremovelastof)
  - [RemoveAll](#removeall)
  - [Replace / ReplaceN / ReplaceAll ](#replace--replacen--replaceall)
  - [Concat](#concat)
  - [Fill](#fill)
  - [SliceByRange](#slicebyrange)
  - [First / FirstOr](#first--firstor)
  - [Last / LastOr](#last--lastor)

**Slice searching**
  - [Contain](#contain--containby)
  - [ContainAll](#containall)
  - [ContainAny](#containany)
  - [Find](#find)
  - [FindLast](#findlast)
  - [IndexOf / IndexOfBy](#indexof--indexofby)
  - [LastIndexOf / LastIndexOfBy](#lastindexof--lastindexofby)
  - [CountValue / CountValueBy](#countvalue--countvalueby)

**Slice filtering**
  - [Filter](#filter)
  - [FilterNE](#filter)
  - [FilterLT / FilterLTE](#filter)
  - [FilterGT / FilterGTE](#filter)
  - [FilterGT / FilterGTE](#filter)
  - [FilterIN / FilterNIN](#filter)
  - [FilterLIKE / FilterILIKE](#filter)

**Slice iteration**
  - [ForEach / ForEachReverse](#foreach--foreachreverse)
  - [Iter / IterReverse](#iter--iterreverse)

**Slice uniqueness**
  - [IsUnique / IsUniqueBy](#isunique--isuniqueby)
  - [FindUniques / FindUniquesBy](#finduniques--finduniquesby)
  - [FindDuplicates / FindDuplicatesBy](#findduplicates--findduplicatesby)
  - [ToSet / ToSetBy / ToSetByReverse ](#toset--tosetby--tosetbyreverse)

**Slice transformation**
  - [MapSlice / MapSliceEx](#mapslice--mapsliceex)
  - [MapSliceToMap / MapSliceToMapEx](#mapslicetomap--mapslicetomapex)
  - [MapSliceToMapKeys](#mapslicetomapkeys)

**Slice algo**
  - [Compact](#compact)
  - [Drop](#drop)
  - [Reverse / ReverseCopy](#reverse--reversecopy)
  - [Shuffle](#shuffle)
  - [Chunk / ChunkByPieces](#chunk--chunkbypieces)
  - [Union / UnionBy](#union--unionby)
  - [Intersection / IntersectionBy](#intersection--intersectionby)
  - [Difference / DifferenceBy](#difference--differenceby)
  - [Reduce / ReduceEx](#reduce--reduceex)
  - [ReduceReverse / ReduceReverseEx](#reducereverse--reducereverseex)
  - [Partition / PartitionN](#partition--partitionn)
  - [Flatten / Flatten3](#flatten--flatten3)
  - [Zip / Zip\<N\>](#zip--zipn)

**Slice conversion**
  - [ToIfaceSlice](#toifaceslice)
  - [ToStringSlice](#tostringslice)
  - [ToNumberSlice](#tonumberslice)
  - [ToSlice](#toslice)
  - [ToPtrSlice](#toptrslice)

**Slice subset**
  - [ContainSlice](#containslice)
  - [IndexOfSlice](#indexofslice)
  - [LastIndexOfSlice](#lastindexofslice)
  - [SubSlice](#subslice)

**Map**
  - [MapEqual / MapEqualBy](#mapequal--mapequalby)
  - [MapContainKeys](#mapcontainkeys)
  - [MapContainValues](#mapcontainvalues)
  - [MapKeys](#mapkeys)
  - [MapValues](#mapvalues)
  - [MapEntries](#mapentries)
  - [MapGet](#mapget)
  - [MapPop](#mappop)
  - [MapSetDefault](#mapsetdefault)
  - [MapUpdate / MapUpdateExistingOnly / MapUpdateNewOnly](#mapupdate--mapupdateexistingonly--mapupdatenewonly)
  - [MapCopy](#mapcopy)
  - [MapPick](#mappick)
  - [MapOmit](#mapomit)
  - [MapOmitCopy](#mapomitcopy)

**Struct**
  - [StructToMap](#structtomap)
  - [ParseTag / ParseTagOf / ParseTagsOf](#parsetag--parsetagof--parsetagsof)

**String**
  - [RuneLength](#runelength)
  - [RandString / RandStringEx](#randstring--randstringex)
  - [StringJoin / StringJoinEx / StringJoinBy](#stringjoin--stringjoinex--stringjoinby)
  - [StringLexJoin / StringLexJoinEx](#stringlexjoin--stringlexjoinex)
  - [StringWrap / StringUnwrap](#stringwrap--stringunwrap)

**Number**
  - [ParseInt / ParseUint / ParseFloat](#parseint--parseuint--parsefloat)
  - [FormatInt / FormatUint / FormatFloat](#formatint--formatuint--formatfloat)
  - [NumberFmtGroup](#numberfmtgroup)

**Time**
  - [MinTime / MaxTime / MinMaxTime](#mintime--maxtime--minmaxtime)
  - [ExecDuration / ExecDurationN](#execduration--execdurationn)

**Math**
  - [All](#all)
  - [Any](#any)
  - [Abs](#abs)
  - [Clamp](#clamp)
  - [Min / Max / MinMax](#min--max--minmax)
  - [Sum / SumAs](#sum--sumas)
  - [Product / ProductAs](#product--productas)

**Concurrency**
  - [ExecTasks / ExecTasksEx](#exectasks--exectasksex)
  - [ExecTaskFunc / ExecTaskFuncEx](#exectaskfunc--exectaskfuncex)

**Function**
  - [Bind\<N\>Arg\<M\>Ret ](#bindnargmret)

**Randomization**
  - [RandChoice](#randchoice)
  - [RandChoiceMaker](#randchoicemaker)

**Utility**
  - [FirstNonEmpty](#firstnonempty)
  - [Coalesce](#coalesce)
  - [If](#if)
  - [Must\<N\>](#mustn)
  - [ToPtr](#toptr)
  - [Head](#head)
  - [Tail](#tail)


### Slice
---
 
#### Equal / EqualBy

Returns true if 2 slices have the same size and all elements of them equal in the current order.

```go
Equal([]int{1, 2, 3}, []int{1, 2, 3}) // true
Equal([]int{1, 2, 3}, []int{3, 2, 1}) // false

// Use EqualBy for custom equal comparison
EqualBy([]string{"one", "TWO"}, []string{"ONE", "two"}, strings.EqualFold) // true
```
 
#### ContentEqual / ContentEqualBy

Returns true if 2 slices have the same size and contents equal regardless of the order of elements.

```go
ContentEqual([]int{1, 2, 3}, []int{2, 1, 3})       // true
ContentEqual([]int{1, 2, 2, 3}, []int{2, 1, 3})    // false
ContentEqual([]int{1, 2, 2, 3}, []int{2, 1, 2, 3}) // true
ContentEqual([]int{1, 2, 2, 3}, []int{1, 1, 2, 3}) // false

// Use ContentEqualBy for custom key function
ContentEqualBy([]string{"one", "TWO"}, []string{"two", "ONE"}, strings.ToLower) // true
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
 
#### Sort / IsSorted
 
Convenient wrappers of the built-in `sort.Slice`.

```go
Sort([]int{1, 3, 2})         // []int{1, 2, 3}
SortDesc([]int{1, 3, 2})     // []int{3, 2, 1}
IsSorted([]int{1, 3, 2})     // false
IsSortedDesc([]int{3, 2, 1}) // true
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

#### Replace / ReplaceN / ReplaceAll

Replaces a value with another value.

```go
// Replaces first occurrence of the value
Replace([]int{1, 2, 1, 3, 1}, 1, 11)     // []int{11, 2, 1, 3, 1}
// Replaces first n occurrences of the value (use -1 to replace all)
ReplaceN([]int{1, 2, 1, 3, 1}, 1, 11, 2) // []int{11, 2, 11, 3, 1}
// Replaces all occurrences of the value
ReplaceAll([]int{1, 2, 1, 3, 1}, 1, 11)  // []int{11, 2, 11, 3, 11}
```

#### Concat
 
Concatenates two or more slices.

```go
Concat([]int{1}, []int{2}, []int{2, 3}) // []int{1, 2, 2, 3}
```

#### Fill

Fills a slice with specified value.

```go
s := make([]int, 5)
Fill(s, 1)  // s == []int{1, 1, 1, 1, 1}

s2 := s[2:4]
Fill(s2, 1) // s2 == []int{1, 1}, s == []int{0, 0, 1, 1, 0}
```

#### First / FirstOr

Returns the first element of slice.

```go
First([]int{1, 2, 3})      // 1, true
First([]string{})          // "", false

// Return the default value if slice is empty
FirstOr([]int{1, 2, 3}, 4) // 1
FirstOr([]int{}, 11)       // 11
```

#### Last / LastOr

Returns the last element of slice.

```go
Last([]int{1, 2, 3})       // 3, true
Last([]string{})           // "", false

// Return the default value if slice is empty
LastOr([]int{1, 2, 3}, 4)  // 3
LastOr([]int{}, 11)        // 11
```

#### SliceByRange

Generates a slice for the given range.

```go
s := SliceByRange(0, 5, 1)         // []int{0, 1, 2, 3, 4}
s := SliceByRange(0.0, 5, 2)       // []float64{0, 2, 4}
s := SliceByRange(int32(5), 0, -2) // []int32{5, 3, 1}
```
 
### Slice searching
---

#### Contain / ContainBy

Returns true if a slice contains a value.

```go
Contain([]int{1, 2, 3}, 2) // true
Contain([]int{1, 2, 3}, 0) // false

// Use ContainBy for custom function
ContainBy([]string{"one", "TWO"}, func(elem string) bool {
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
 
#### Find

Finds a value in a slice.

```go
v, found := Find([]string{"one", "TWO"}, func(elem string) bool {
    return strings.ToLower(elem) == "two"
}) // v == "TWO", found == true
```

#### FindLast

Finds a value in a slice from the end.

```go
v, found := FindLast([]string{"one", "TWO", "ONe"}, func(elem string) bool {
    return strings.ToLower(elem) == "one"
}) // v == "ONe", found == true
```

#### IndexOf / IndexOfBy

Finds the index of a value in a slice, returns -1 if not found.

```go
IndexOf([]int{1, 2, 3}, 4) // -1
IndexOf([]int{1, 2, 3}, 2) // 1

// Use IndexOfBy for custom function
IndexOfBy([]string{"one", "TWO"}, func(elem string) bool {
    return strings.ToLower(elem) == "two"
}) // 1
```

#### LastIndexOf / LastIndexOfBy

Finds the last index of an element in a slice, returns -1 if not found.

```go
LastIndexOf([]int{1, 2, 3}, 4)    // -1
LastIndexOf([]int{1, 2, 1, 3}, 1) // 2
```
 
#### CountValue / CountValueBy

Counts the number of occurrences of a value in a slice.

```go
CountValue([]int{1, 2, 3}, 4)    // 0
CountValue([]int{1, 2, 3, 2}, 2) // 2
```

### Slice filtering
---

#### Filter

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
 
### Slice iteration
---

#### ForEach / ForEachReverse

Iterates over slice content.

```go
ForEach([]int{1, 2, 3}, func (i, v int) {
    fmt.Printf("%d ", v)
}) // prints 1 2 3

ForEachReverse([]int{1, 2, 3}, func (i, v int) {
    fmt.Printf("%d ", v)
}) // prints 3 2 1

// ForEachPtr can be faster if you iterate over a slice of big structs
ForEachPtr([]BigStruct{...}, func (i, v *BigStruct) { ... })
ForEachPtrReverse([]BigStruct{...}, func (i, v *BigStruct) { ... })
```

#### Iter / IterReverse

Iterates over one or multiple slices with ability to stop.

```go
Iter(func (i, v int) bool {
    fmt.Printf("%d ", v)
    return i < 3
}, []int{1, 2, 3}, []int{4, 5}) // prints 1 2 3 4

IterReverse(func (i, v int) bool {
    fmt.Printf("%d ", v)
    return true
}, []int{1, 2, 3}, []int{4, 5}) // prints 5 4 3 2 1

// IterPtr can be faster if you iterate over a slice of big structs
IterPtr(func (i, v *BigStruct) bool { ... }, []BigStruct{...})
IterPtrReverse(func (i, v *BigStruct) bool { ... }, []BigStruct{...})
```

### Slice uniqueness
---

#### IsUnique / IsUniqueBy

Returns true if a slice contains unique values.

```go
IsUnique([]int{1, 2, 3}) // true
IsUnique([]int{1, 2, 1}) // false

// Use IsUniqueBy for custom function
IsUniqueBy([]string{"one", "ONE"}, strings.ToLower) // false
```

#### FindUniques / FindUniquesBy

Finds all unique elements in the given slice. The order of elements in the result is the same
as they appear in the input.

```go
FindUniques([]int{1, 2, 3, 2})  // []int{1, 3}
FindUniques([]int{1, 2, 2, 1})  // []int{}

// FindUniquesBy supports custom key function
FindUniquesBy([]string{"one", "ONE", "Two"}, func (s string) string {
    return strings.ToLower(s)
}) // []string{"Two"}
```

#### FindDuplicates / FindDuplicatesBy

Finds all elements which are duplicated in the given slice. The order of elements in the result is the same
as they appear in the input.

```go
FindDuplicates([]int{1, 2, 3, 2})  // []int{2}
FindDuplicates([]int{1, 2, 3})     // []int{}

// FindDuplicatesBy supports custom key function
FindDuplicatesBy([]string{"one", "ONE", "Two"}, func (s string) string {
    return strings.ToLower(s)
}) // []string{"one"}
```

#### ToSet / ToSetBy / ToSetByReverse

Calculates unique values of a slice.

```go
ToSet([]int{1, 2, 3, 1, 2})        // []int{1, 2, 3}
ToSet([]string{"one", "2", "one"}) // []string{"one", "2"}

// Use ToSetBy for custom key function
ToSetBy([]string{"one", "TWO", "two", "One"}, strings.ToLower) // []string{"one", "TWO"}


// Use ToSetByReverse for iterating items from the end of the list
ToSetByReverse([]string{"one", "TWO", "two", "One"}, strings.ToLower) // []string{"One", "two"}
```

### Slice transformation
---

#### MapSlice / MapSliceEx

Transforms a slice to a slice.

```go
MapSlice([]string{"1", "2 ", " 3"}, strings.TrimSpace)   // []string{"1", "2", "3"}

// Use MapSliceEx to transform with error handling
MapSliceEx([]string{"1","2","3"}, gofn.ParseInt[int])    // []int{1, 2, 3}
MapSliceEx([]string{"1","x","3"}, gofn.ParseInt[int])    // strconv.ErrSyntax
MapSliceEx([]string{"1","200","3"}, gofn.ParseInt[int8]) // strconv.ErrRange
```

#### MapSliceToMap / MapSliceToMapEx

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

#### MapSliceToMapKeys

Transforms a slice to a map with using slice items as map keys.

```go
MapSliceToMapKeys([]int{1, 2, 3, 2}, "x")     // map[int]string{1: "x", 2: "x", 3: "x"}
MapSliceToMapKeys([]int{1, 2, 1}, struct{}{}) // map[int]struct{}{1: struct{}{}, 2: struct{}{}}
```

### Slice algo
---

#### Compact

Compacts a slice by removing zero elements.  Not change the source slice.

```go
s := Compact([]int{1, 0, 3}) // s == []int{1, 3}
```

#### Reverse / ReverseCopy

Reverses slice content.

```go
Reverse([]int{1, 2, 3}) // []int{3, 2, 1}

s1 := []int{1, 2, 3}
s2 := ReverseCopy(s1)  // s1 == []int{1, 2, 3}, s2 == []int{3, 2, 1}
```

#### Drop

Returns a new slice with dropping values in the specified list.
NOTE: this function is just a call to `FilterNIN()`.

```go
Drop([]int{1, 2, 3, 4}, 3, 1) // []int{2, 4}
```

#### Shuffle

Shuffle items of a slice. Not change the source slice.

```go
s := Shuffle([]int{1, 2, 3}) // s is a new slice with random items of the input
```

#### Chunk / ChunkByPieces

Splits slice content into chunks.

```go
Chunk([]int{1, 2, 3, 4, 5}, 2)         // [][]int{[]int{1, 2}, []int{3, 4}, []int{5}}
ChunkByPieces([]int{1, 2, 3, 4, 5}, 2) // [][]int{[]int{1, 2, 3}, []int{4, 5}}
```

#### Union / UnionBy

Finds all unique values from multiple slices.

```go
Union([]int{1, 3, 2}, []int{1, 2, 2, 4}) // []int{1, 3, 2, 4}
```

#### Intersection / IntersectionBy

Finds all unique shared values from multiple slices.

```go
Intersection([]int{1, 3, 2}, []int{1, 2, 2, 4}) // []int{1, 2}
```

#### Difference / DifferenceBy

Finds all different values from 2 slices.

```go
left, right := Difference([]int{1, 3, 2}, []int{2, 2, 4}) // left == []int{1, 3}, right == []int{4}
```
 
#### Reduce / ReduceEx

Reduces a slice to a value.

```go
Reduce([]int{1, 2, 3}, func (accumulator int, currentValue int) int {
    return accumulator + currentValue
}) // 6

ReduceEx([]int{1, 2, 3}, func (accumulator int, currentValue int, i index) int {
    return accumulator + currentValue
}, 10) // 16
```

#### ReduceReverse / ReduceReverseEx

Reduces a slice to a value with iterating from the end.

```go
ReduceReverse([]int{1, 2, 3}, func (accumulator int, currentValue int) int {
    return accumulator + currentValue
}) // 6

ReduceReverseEx([]int{1, 2, 3}, func (accumulator int, currentValue int, i index) int {
    return accumulator + currentValue
}, 10) // 16
```

#### Partition / PartitionN

Splits a slice into multiple partitions.

```go
// Splits a slice into 2 partitions
p0, p1 := Partition([]int{1, 2, 3, 4, 5}, func (v int, index int) bool {return v%2==0})) // p0 == []int{2, 4}, p1 == []int{1, 3, 5}

// Splits a slice into 3 partitions
p := PartitionN([]int{1, 2, 3, 4, 5}, 3, func (v int, index int) int {return v%3})) // p == [[3], [1, 4], [2, 5]]
```

#### Flatten / Flatten3

Flattens multi-dimension slice.

```go
Flatten([]int{1, 2, 3}, []int{4, 5})                    // []int{1, 2, 3, 4, 5}
Flatten3([][]int{{1, 2}, {3, 4}, [][]int{{5, 6}, {7}}}) // []int{1, 2, 3, 4, 5, 6, 7}
```

#### Zip / Zip\<N\>

Combines values from multiple slices by each position (`N` is from 3 to 5).

```go
Zip([]int{1, 2, 3}, []int{4, 5})                              // []*Tuple2{{1, 4), {2, 5}}
Zip3([]int{1, 2, 3}, []string{"4", "5"}, []float32{6.0, 7.0}) // []*Tuple3{{1, "4", 6.0), {2, "5", 7.0}}
```

### Slice conversion
---

#### ToIfaceSlice

Converts a slice of any type to a slice of interfaces.

```go
ToIfaceSlice([]int{1, 2, 3})         // []any{1, 2, 3}
ToIfaceSlice([]string{"foo", "bar"}) // []any{"foo", "bar"}
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

#### ToPtrSlice

Creates a slice of pointers point to the given elements.

```go
s1 := []int{1, 2, 3}
s2 := ToPtrSlice(s1) // []*int{&s1[0], &s1[1], &s1[2]}
```

### Slice subset
---

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

#### SubSlice

Returns sub slice of a slice in range [start, end). `end` param is exclusive. This function doesn't raise error.
Passing negative numbers for `start` and `end` to get items from the end of the slice.

```go
SubSlice([]int{1, 2, 3}, 0, 2)   // []{1, 2}
SubSlice([]int{1, 2, 3}, -1, -2) // []{3}
SubSlice([]int{1, 2, 3}, -1, -3) // []{2, 3}
```

### Map 
---
 
#### MapEqual / MapEqualBy

Returns true if 2 maps equal.

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

#### MapUpdate / MapUpdateExistingOnly / MapUpdateNewOnly

Updates a map content with another map.

```go
s := map[int]int{1: 11, 2: 22}
MapUpdate(s, map[int]int{1: 111, 3: 33})             // s == map[int]int{1: 111, 2: 22, 3: 33}
MapUpdateExistingOnly(s, map[int]int{2: 222, 3: 33}) // s == map[int]int{1: 11, 2: 222}
MapUpdateNewOnly(s, map[int]int{2: 222, 3: 33})      // s == map[int]int{1: 11, 2: 22, 3: 33}
```

#### MapCopy

Copies a map.

```go
MapCopy(map[int]int{1: 11, 2: 22, 3: 33}) // map[int]int{1: 11, 2: 22, 3: 33}
```

#### MapPick

Copies map content for specified keys only.

```go
MapPick(map[int]int{1: 11, 2: 22, 3: 33}, 2, 3, 2) // map[int]int{2: 22, 3: 33}
```

#### MapOmit

Omits keys from a map.

```go
m := map[int]int{1: 11, 2: 22, 3: 33}
MapOmit(m, 2, 3, 4) // m == map[int]int{1: 11}
```

#### MapOmitCopy

Copies map content with omitting specified keys.

```go
MapOmitCopy(map[int]int{1: 11, 2: 22, 3: 33}, 2, 3, 4) // map[int]int{1: 11}
```

### Struct
---

#### StructToMap

Converts struct contents to a map. This function is a shortcut to [rflutil.StructToMap](https://github.com/tiendc/go-rflutil#structtomap).

#### ParseTag / ParseTagOf / ParseTagsOf

Parses struct tags. These functions are shortcuts to [rflutil.ParseTag](https://github.com/tiendc/go-rflutil#parsetag--parsetagof--parsetagsof).

### String
---

#### RuneLength

Alias of `utf8.RuneCountInString`.

```go
len("café")        // 5
RuneLength("café") // 4
```

#### RandString / RandStringEx

Generates a random string.

```go
RandString(10)                         // Generates a string of 10 characters from alphabets and digits
RandStringEx(10, []rune("0123456789")) // Generates a string of 10 characters from the specified ones
```

#### StringJoin / StringJoinEx / StringJoinBy

Joins a slice of any element type.

```go
s := StringJoin([]int{1,2,3}, ", ") // s == "1, 2, 3"

type Struct struct {
    I int
    S string
}
s := StringJoinBy([]Struct{{I:1, s:"a"}, {I:2, s:"b"}}, ", ", func (v Struct) string {
    return fmt.Sprintf("%d:%s", v.I, v.S)
}) // s == "1:a, 2:b"
```

#### StringLexJoin / StringLexJoinEx

Joins a slice of any element type in lexical manner.

```go
StringLexJoin([]int{1,2,3}, ", ", " and ")              // return "1, 2 and 3"

// Use a custom format string
StringLexJoinEx([]int64{254, 255}, ", ", " or ", "%#x") // returns "0xfe or 0xff"
```

#### StringWrap / StringUnwrap

Wraps/Unwraps a string with the given tokens.

```go
StringWrap("abc", "*")            // "*abc*"
StringWrapLR("abc", "[", "]")     // "[abc]"

StringUnwrap("*abc*", "*")        // "abc"
StringUnwrapLR("[abc]", "[", "]") // "abc"
```

### Number
---

#### ParseInt / ParseUint / ParseFloat

Parses a number using **strconv.ParseInt** then converts the value to a specific type.

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

#### FormatInt / FormatUint / FormatFloat

Formats a number value.

```go
FormatInt(123)            // "123"

// Format number with specific format string (use fmt.Sprintf)
FormatIntEx(123, "%05d")  // "00123"

// Format number with decimal grouping
FormatIntGroup(1234567)   // "1,234,567"
```

- **NOTE**: There are also **FormatUint** for unsigned integers and **FormatFloat** for floating numbers.

#### NumberFmtGroup

Groups digits of a number. Input number is of type string.

```go
NumberFmtGroup("1234567", '.', ',')         // "1,234,567"
NumberFmtGroup("1234567", ',', '.')         // "1.234.567"
NumberFmtGroup("1234567.12345", '.', ',')   // "1,234,567.12345"
```

### Concurrency
---

#### ExecTasks / ExecTasksEx

Executes tasks concurrently with ease. This function provides a convenient way for one of the most
popular use case in practical.

```go
// In case you want to store the task results into a shared variable,
// make sure you use enough synchronization
var task1Result any
var task2Result []any

// Allow spending maximum 10s to finish all the tasks
ctx := context.WithTimeout(context.Background(), 10 * time.Second)

err := ExecTasks(ctx, 0 /* max concurrent tasks */,
    // Task 1st:
    func(ctx context.Context) (err error) {
        task1Result, err = getDataFromDB()
        return err
    },
    // Task 2nd:
    func(ctx context.Context) (err error) {
        for i:=0; i<10; i++ {
            if err := ctx.Err(); err != nil {
                return err
            }
            task2Result = append(task2Result, <some data>)
            return nil
        }
    },
)
if err != nil {
    // one or more tasks failed
}
```

#### ExecTaskFunc / ExecTaskFuncEx

Executes a task function on every target objects concurrently. This function is similar to `ExecTasks()`, but it
takes only one function and a list of target objects.

```go
var mu sync.Mutex
var evens []int
var odds []any

taskFunc := func(ctx context.Context, v int) error {
    mu.Lock()
    if v%2 == 0 {
        evens = append(evens, v)
    } else {
        odds = append(odds, v)
    }
    mu.Unlock()
    return nil
}

err := ExecTaskFunc(ctx, 0 /* max concurrent tasks */, taskFunc, 1, 2, 3, 4, 5)
if err != nil {
    // one or more tasks failed
}

// Result is: evens has [2, 4], odds has [1, 3, 5] (with undetermined order of items)
```

### Function
---

#### Bind\<N\>Arg\<M\>Ret

Fully binds a function with returning a function which requires no argument.

```go
func myCalc(a1 int, a2 string) error { ... }
myQuickCalc := Bind2Arg1Ret(myCalc, 100, "hello")
err := myQuickCalc()
```

### Randomization

**NOTE**: Should not use these functions for crypto purpose.

---

#### RandChoice

Picks up an item randomly from a list of items.

```go
val, valid := RandChoice(1, 2, 3) // valid == true and `val` is one of the input items
val, valid := RandChoice[int]()   // valid == false and `val` is int(0)
```

#### RandChoiceMaker

Provides a method to pick up items randomly from a list of items without duplication of choice.

```go
choiceMaker := NewRandChoiceMaker([]int{1, 2, 3})
for choiceMaker.HasNext() {
    randItem, _ := choiceMaker.Next()
}
// OR
for {
    randItem, valid := choiceMaker.Next()
    if !valid {
        break
    }
}
```

### Time
---

#### MinTime / MaxTime / MinMaxTime

Finds minimum/maximum time value in a slice.

```go
t0 := time.Time{}
t1 := time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC)
t2 := time.Date(2001, time.December, 1, 0, 0, 0, 0, time.UTC)
MinTime(t0, t1, t2)     // t0
MinTime(t1, t2)         // t1
MaxTime(t0, t1, t2)     // t2
MinMaxTime(t0, t1, t2)  // t0, t2
```

#### ExecDuration / ExecDurationN

Measures time executing a function.

```go
duration := ExecDuration(func() { // do somthing })
// The given function returns a value
outVal, duration := ExecDuration1(func() string { return "hello" })              // outVal == "hello"
// The given function returns 2 values
outVal1, err, duration := ExecDuration2(func() (int, error) { return 123, nil }) // outVal1 == 123, err == nil
```

### Math 
---

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

#### Abs

Calculates absolute value of an integer.

```go
Abs(-123)          // int64(123)
Abs(123)           // int64(123)
Abs(math.MinInt64) // math.MinInt64 (special case)
```
 
#### Clamp

Clamps a value within a range (lower and upper bounds are inclusive).

```go
Clamp(3, 10, 20)  // 10
Clamp(30, 10, 20) // 20
Clamp(15, 10, 20) // 15
```

#### Min / Max / MinMax

Finds minimum/maximum value in a slice.

```go
Min(1, 2, 3, -1)             // -1
MinIn([]int{1, 2, 3, -1})    // -1
MinInBy([]string{"a", "B"}, func(a, b int) bool {
    return strings.ToLower(a) < strings.ToLower(b)
}) // "a"

Max(1, 2, 3, -1)             // 3
MaxIn([]int{1, 2, 3, -1})    // 3
MaxInBy([]string{"a", "B"}, func(a, b int) bool {
    return strings.ToLower(a) < strings.ToLower(b)
}) // "B"

MinMax(1, 2, 3, -1)          // -1, 3
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

### Utility
---

#### FirstNonEmpty

Returns the first non-empty value in the given arguments if found, otherwise returns the zero value.
This function uses `reflection`. You can connsider using `Coalesce` for generic types.

Values considered "non-empty" are:
  - not empty values (0, empty string, false, nil, ...)
  - not empty containers (slice, array, map, channel)
  - not pointers that point to zero/empty values

```go
FirstNonEmpty(0, -1, 2, 3)                          // -1
FirstNonEmpty("", " ", "b")                         // " "
FirstNonEmpty([]int{}, nil, []int{1}, []int{2, 3})  // []int{1}
FirstNonEmpty([]int{}, nil, &[]int{}, []int{2, 3})  // []int{2, 3}
FirstNonEmpty[any](nil, 0, 0.0, "", struct{}{})     // nil (the first zero value)
```

#### Coalesce

Returns the first non-zero argument. Input type must be comparable.

```go
Coalesce(0, -1, 2)         // -1
Coalesce("", " ", "s")     // " "
Coalesce[*int](ptr1, ptr2) // the first non-nil pointer
```

#### If

A convenient function works like C ternary operator.

**NOTE**: This function is deprecated as it may cause the program to crash on misuses due to both passing
expressions are evaluated regardless of the condition.
For example: `firstItem := If(len(slice) > 0, slice[0], defaultVal)` will crash if `slice` is empty as
the expression `slice[0]` is always evaluated. Use it at your own risk.

```go
val := If(x > 100, val1, val2) // If x > 100, val == val1, otherwise val == val2
```

#### Must\<N\>

Must\<N\> ( N is from 1 to 6) functions accept a number of arguments with the last one is of `error` type.
Must\<N\> functions return the first N-1 arguments if the error is `nil`, otherwise they panic.

```go
func CalculateAmount() (int, error) {}
amount := Must(CalculateAmount()) // panic on error, otherwise returns the amount

func CalculateData() (int, string, float64, error) {}
v1, v2, v3 := Must4(CalculateData()) // panic on error, otherwise returns the 3 first values
```

#### ToPtr

Returns a pointer to the input argument.

```go
func aFunc(ptr *int) {}

// Use ToPtr to pass a value inline
aFunc(ToPtr(10))
```

#### Head

Takes the first argument.

```go
Head(1, "2", 1.0, 3)   // 1
```

#### Tail

Takes the last argument.

```go
Tail[string](true, "2", 1.0, "3") // "3"
``` 

## Benchmarks

#### Equal vs ContentEqual vs reflect.DeepEqual
___

```
Benchmark_Slice_Equal/StructSlice/Equal
Benchmark_Slice_Equal/StructSlice/Equal-8           510845715           2.047 ns/op
Benchmark_Slice_Equal/StructSlice/ContentEqual
Benchmark_Slice_Equal/StructSlice/ContentEqual-8    583167950           2.061 ns/op
Benchmark_Slice_Equal/StructSlice/DeepEqual
Benchmark_Slice_Equal/StructSlice/DeepEqual-8       15403771            79.19 ns/op

Benchmark_Slice_Equal/IntSlice/Equal
Benchmark_Slice_Equal/IntSlice/Equal-8              589706185           2.087 ns/op
Benchmark_Slice_Equal/IntSlice/ContentEqual
Benchmark_Slice_Equal/IntSlice/ContentEqual-8       523120755           2.194 ns/op
Benchmark_Slice_Equal/IntSlice/DeepEqual
Benchmark_Slice_Equal/IntSlice/DeepEqual-8          15243183            77.93 ns/op
```

## Contributing

- You are welcome to make pull requests for new functions and bug fixes.

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
