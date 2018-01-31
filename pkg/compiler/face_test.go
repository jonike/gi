package compiler

import (
	"fmt"
	"testing"

	cv "github.com/glycerine/goconvey/convey"
)

var _ = fmt.Printf
var _ = testing.T{}
var _ = cv.So

func Test100Interfaces(t *testing.T) {

	/*
		       0) declare a var as an interface

			        a) one-value conversion:
			             as := any.(Stringer)

			    b)    two-value conversion check:

			   type Stringer interface {
			      String() string
			   }
			    if v, ok := any.(Stringer); ok {
			        return v.String()
			    }

			    c) type switch:

			   func ToString(any interface{}) string {

			    switch v := any.(type) {
			    case int:
			        return strconv.Itoa(v)
			    case float:
			        return strconv.Ftoa(v, 'g', -1)
			    }
			    return "???"
			    }

			                d) assignment /compile time check:

			       var s Stringer = &MyType{}

	*/

	cv.Convey(`declare an interface`, t, func() {

		code := `
type Counter interface {
   Next() int
}
type S struct {
   v int
}
func (s *S) Next() int {
  s.v++
  return s.v
}
var c Counter = &S{}
a := c.Next()
b := c.Next()

type ByTen struct {
   v int
}
func (s *ByTen) Next() int {
   s.v += 10
   return s.v
}
bt := &ByTen{}
c = bt
d := c.Next()
e := c.Next()
	`
		vm, err := NewLuaVmWithPrelude(nil)
		panicOn(err)
		defer vm.Close()
		inc := NewIncrState(vm, nil)

		translation := inc.Tr([]byte(code))
		fmt.Printf("\n translation='%s'\n", translation)

		//cv.So(string(translation), cv.ShouldMatchModuloWhiteSpace, ``)

		// and verify that it happens correctly
		LuaRunAndReport(vm, string(translation))

		LuaMustInt64(vm, "a", 1)
		LuaMustInt64(vm, "b", 2)

		LuaMustInt64(vm, "d", 10)
		LuaMustInt64(vm, "e", 20)

	})
}
