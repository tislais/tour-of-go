# Basics - Flow Control Statements

## Switch

A `switch` statement is a shorter way to write a sequence if `if - else` statements. It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the `break` statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

## Switch evaluation order

Switch cases evaluate cases from top to bottom, stopping when a case succeeds. 

For example, the below code does not call `f` if `i==0`
```
switch i {
    case 0:
    case f():
}
```

**Note:** Time in the Go playground always appears to start at 2009-11-10 23:00:00 UTC, a value whose significance is left as an exercise for the reader.

## Switch with no condition

Switch without a condition is the same as `switch true`.

This construct can be a clean way to write long if-then-else chains.