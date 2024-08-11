package gofn

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_StringJoin(t *testing.T) {
	assert.Equal(t, "", StringJoin[int]([]int(nil), ","))
	assert.Equal(t, "1", StringJoin[int]([]int{1}, ","))
	assert.Equal(t, "1,2,3", StringJoin[int64]([]int64{1, 2, 3}, ","))
	// Slice has nil element
	assert.Equal(t, "1,null,3", StringJoin[any]([]any{1, nil, "3"}, ","))
	// Slice of fmt.Stringer elements
	assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC, 2020-12-01 00:00:00 +0000 UTC", StringJoin[any]([]any{
		time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2020, time.November, 31, 0, 0, 0, 0, time.UTC),
	}, ", "))
}

func Test_StringJoinPred(t *testing.T) {
	assert.Equal(t, "", StringJoinPred[int]([]int(nil), ",", func(v int) string {
		return fmt.Sprintf("%d", v)
	}))
	assert.Equal(t, "1", StringJoinPred[int]([]int{1}, ",", func(v int) string {
		return fmt.Sprintf("%d", v)
	}))
	assert.Equal(t, "1,2,3", StringJoinPred[int64]([]int64{1, 2, 3}, ",", func(v int64) string {
		return fmt.Sprintf("%d", v)
	}))
	// Slice has nil element
	assert.Equal(t, "1,nil,3", StringJoinPred[any]([]any{1, nil, "3"}, ",", func(v any) string {
		if v == nil {
			return "nil"
		}
		return fmt.Sprintf("%v", v)
	}))
}

func Test_LinesTrimLeft(t *testing.T) {
	assert.Equal(t, "", LinesTrimLeftSpace(""))
	assert.Equal(t, "line-1\nline-2", LinesTrimLeftSpace(`line-1
		line-2`))
	assert.Equal(t, "\nline-1\nline-2", LinesTrimLeftSpace(`
		line-1
		line-2`))
	assert.Equal(t, "\nline-1\nline-2\n", LinesTrimLeftSpace(`
		line-1
		line-2
	`))
	assert.Equal(t, "\nx  line-1\nline-2", LinesTrimLeftSpace(`
		  x  line-1
		line-2`))
	// Unicode
	assert.Equal(t, "\nâ  line-1\nline-2 â", LinesTrimLeftSpace(`
		  â  line-1
		line-2 â`))
	// Extra func test
	assert.Equal(t, "\nê b â  line-1\nê line-2 â ê ", LinesTrimLeft(`
		 a ê b â  line-1
		  b aê line-2 â ê `, "a b\t")) // ascii cutset
	assert.Equal(t, "\nâ  line-1\nline-2 â ê", LinesTrimLeft(`
		 ê â  line-1
		 ê line-2 â ê`, "a b\tê")) // unicode cutset
}

func Test_LinesTrimRight(t *testing.T) {
	assert.Equal(t, "", LinesTrimRightSpace(""))
	assert.Equal(t, "line-1\nline-2", LinesTrimRightSpace(`line-1
line-2`))
	assert.Equal(t, "\nline-1\nline-2\n", LinesTrimRightSpace(`
line-1   
line-2
	`))
	assert.Equal(t, "\nline-1\nline-2\n", LinesTrimRightSpace(`
line-1
line-2
	`))
	assert.Equal(t, "\nx  line-1  x\nline-2", LinesTrimRightSpace(`
x  line-1  x  
line-2`))
	// Unicode
	assert.Equal(t, "\nâ  line-1  â\nline-2", LinesTrimRightSpace(`
â  line-1  â  
line-2`))
	// Extra func test
	assert.Equal(t, "\nâ  line-1  â  ê\nline-2  ê", LinesTrimRight(`
â  line-1  â  ê  
line-2  ê  `, "a b\t")) // ascii cutset
	assert.Equal(t, "\n â  line-1  â\nline-2", LinesTrimRight(`
 â  line-1  â  ê  
line-2  ê  `, "a b\tê")) // unicode cutset
}

func Test_LinesTrim(t *testing.T) {
	assert.Equal(t, "", LinesTrimSpace(""))
	assert.Equal(t, "line-1\nline-2", LinesTrimSpace(`line-1
line-2`))
	assert.Equal(t, "\nline-1\nline-2\n", LinesTrimSpace(`
line-1   
   line-2
	`))
	assert.Equal(t, "\nline-1\nline-2\n", LinesTrimSpace(`
   line-1
   line-2
	`))
	assert.Equal(t, "\nx  line-1  x\nline-2", LinesTrimSpace(`
   x  line-1  x  
   line-2   `))
	// Unicode
	assert.Equal(t, "\nâ  line-1  â\nline-2", LinesTrimSpace(`
  â  line-1  â  
  line-2  `))
	// Extra func test
	assert.Equal(t, "\nâ  line-1  â\nline-2", LinesTrim(`
    â  line-1  â  ê  
    line-2  ê  `, "a b\tê"))
}
