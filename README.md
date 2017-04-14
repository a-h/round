## Rounding

Why!? Because Go's `math` package doesn't include a `Round()` function.

All of the implementations I saw lacked unit tests or comments to make it easy to understand, so I made one and tested it.

There are two functions:

* ToEven
  * Default in Python 3 and C#
  * -3.5 = -4
  * -2.5 = -2
  * -1.5 = -2
  * 1.5 = 2
  * 2.5 = 2
  * 3.5 = 4
* AwayFromZero
  * Default in Python 2
  * -3.5 = -4
  * -2.5 = -3
  * -1.5 = -2
  * 1.5 = 2
  * 2.5 = 3
  * 3.5 = 4

# How to use

```
go get github.com/a-h/round
```

```
import "github.com/a-h/round"

func main() {
    fmt.Println(round.ToEven(float64(-3.5), 0)) // -4
    fmt.Println(round.ToEven(float64(-2.5), 0)) // -2
    fmt.Println(round.ToEven(float64(-1.5), 0)) // -2
    fmt.Println(round.ToEven(float64(-1.49), 0)) // -1
}
```

# Reference

Learn about Rounding Methods.
 * https://www.mathsisfun.com/numbers/rounding-methods.html

Issue on Github asking for a float function.
 * https://github.com/golang/go/issues/4594

Discussion on Go forum discussing how to Round, and not really reaching consensus.
 * https://groups.google.com/forum/#!topic/golang-nuts/ITZV08gAugI/discussion
