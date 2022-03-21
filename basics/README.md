# Basics

## Packages
Every Go program is made up of packages.

Programs start running in package `main`.

This program is using the packages with import paths `"fmt"` and `"math/rand"`.

By convention, the package name is the same as the last element of the import path. For instance, the `"math/rand"` package comprises files that begin with the statement package rand.

Note: The environment in which these programs are executed is deterministic, so each time you run the example program `rand.Intn` will return the same number.

(To see a different number, seed the number generator; see [rand.Seed](https://pkg.go.dev/math/rand#Seed). Time is constant in the playground, so you will need to use something else as the seed.)

## Imports

This code groups the imports into a parenthesized, "factored" import statement.

You can also write multiple import statements, like:

```
import "fmt"
import "math"
```

But it is good style to use the factored import statement.

## Exported names

In Go, a name is exported if it begins with a capital letter. When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

## Functions

A function can take zero or more arguments. Type comes *after* the variable name. See the article on [Go's declaration syntax](https://blog.golang.org/gos-declaration-syntax).

When two or more consercutive named function parameters share a type, you can omit the type from all but the last.

For example,

```
x int, y int
```

Can be shortened to
```
x, y int
```

## Multiple Results

A function can return any number of results. The `swap` function below returns two strings.

```
func swap(x, y string) (string, string) {
	return y, x
}
```

## Named Return Values

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A `return` statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as they can harm readability.

## Variables

The `var` statement declares a list of variables; as in function argument lists, the type is last.

A `var` statement can be at package or function level.

## Variables with Initializers

A `var` declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted. The variable will take the type of the initializer.

## Short Variable Declarations

Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.

Outside a function, every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.

## Basic Types

Go's basic types are

```
bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32 - Unicode code point
float32 float64
complex64 complex128
```

Variable declarations may be "factored" into blocks, like we do with import statements.

```
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

The `int`, `uint`, and `unintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use `int` unless you have a specific reason to use a sized or unsigned integer type.

## Zero Values

Variables declared without an explicit intitial value are given their *zero value*.

The zero value is:
- `0` for numeric types
- `false` for the boolean type
- `""` for strings

## Type conversions

The expression `T(v)` converts the value `v` to the type `T`.

Some numeric conversions:
```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

Or, simply:
```
i := 42
f := float64(i)
u := uint(f)
```

Unlike in C, in Go assignment between items of different types requires an explicit conversion.

## Type Inference

When delcaring a variable without specifying an explicit type (either by using the `:=` syntax or `var =` expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:

```
var i int
j := i // j is an int
```

But when the right hand side contains an untyped numeric constant, the new variable may be an `int`, `float64`, or `complex128` depending on the precision of the constant:

```
i := 42           // int
f := 3.142        // float
g := 0.867 + 0.5i // complex128
```

## Constants

Constants are declared like variables, but with the `const` keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the `:=` syntax.

## Numeric Constants

Numeric constants are high-precision *values*

An untyped constant takes the type needed by its context.
