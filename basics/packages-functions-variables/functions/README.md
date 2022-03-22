# Basics - packages, functions, and variables

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