# Bee

> Bee Programming Language

[![GoDoc](https://godoc.org/github.com/eugenetriguba/bee?status.svg)](https://godoc.org/github.com/eugenetriguba/bee)
[![Go Report Card](https://goreportcard.com/badge/github.com/eugenetriguba/bee)](https://goreportcard.com/report/github.com/eugenetriguba/bee)
[![codebeat badge](https://codebeat.co/badges/5405b9cd-2a4b-4dab-afe4-17a8866035b9)](https://codebeat.co/projects/github-com-eugenetriguba-bee-master)

Bee is an interpreted programming language with a JavaScript-like syntax. 
The interpreter is written as a tree-walking interpreter.

It looks like this.
```
let five = 5; 
let ten = 10;

let add = fn(x, y) { 
  x + y; 
};

let result = add(five, ten);
```
