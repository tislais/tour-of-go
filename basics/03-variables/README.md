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