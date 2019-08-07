# Equation

Simple library which can solve complex Math expression for you!


## Operators and Constants

* Operators and their priority, [here](operators/operators.go)
* Constants and their value, [here](constants/constants.go)

## Examples

```go
package main

import (
	"fmt"

	"github.com/amovah/equation"
	"github.com/amovah/equation/constants"
	"github.com/amovah/equation/operators"
)

func main() {
	fmt.Println(equation.Solve(
		"2+(-5)+3*3-8^2",
		operators.Defaults(),
		constants.Defaults(),
	))
}
```

---

```go
package main

import (
	"fmt"

	"github.com/amovah/equation"
	"github.com/amovah/equation/constants"
	"github.com/amovah/equation/operators"
)

func main() {
	ops := operators.Add(
		operators.Defaults(), // use default operators, it is not necessary, but recommended
		operators.Operator{
			Symbol: "@",
			Operation: func(nums ...float64) float64 {
				return nums[0] * 2
			},
			Priority: 3,
		},
	)

	consts := constants.Add(
		constants.Defaults(), // use default constants, it is not necessary, but recommended
		constants.Constant{
			Symbol: "foo",
			Value:  10,
		},
	)

	fmt.Println(equation.Solve(
		"@2+@(2)+foo",
		ops,
		consts,
	))
}

```