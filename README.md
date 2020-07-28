# Bee

> Bee Programming Language

[![GoDoc](https://godoc.org/github.com/eugenetriguba/bee?status.svg)](https://godoc.org/github.com/eugenetriguba/bee)
[![Go Report Card](https://goreportcard.com/badge/github.com/eugenetriguba/bee)](https://goreportcard.com/report/github.com/eugenetriguba/bee)
[![codebeat badge](https://codebeat.co/badges/5405b9cd-2a4b-4dab-afe4-17a8866035b9)](https://codebeat.co/projects/github-com-eugenetriguba-bee-master)

Bee is an interpreted programming language with a JavaScript-like syntax.

It looks like this.

```
let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);

let trueOrFalse = if (10 > 5) { true } else { false };

(fn(x) { return x }(5) + 10 ) * 10;
```

The language only has two statements: `let` and `return`.
Everything else is considered a expression.

# REPL

Bee comes with a simple REPL program we can use to test it out.

```
$ go run main.go

Hello eugene! Welcome to the Bee programming language!
Feel free to type in commands. Enter 'exit' to leave the repl
>> let testVar = 5;
{Type:LET Literal:let}
{Type:IDENT Literal:testVar}
{Type:= Literal:=}
{Type:INT Literal:5}
{Type:; Literal:;}
>> exit
```

For now, it just tokenizes the input.
