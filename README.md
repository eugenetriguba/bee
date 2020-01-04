# Bee

> Bee Programming Language

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