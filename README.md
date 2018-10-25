<h1 align="center">
  <a href="#">
    <img src="./assets/logo.svg" alt="logo" width="100px">
  </a>
  <br>
  ape
</h1>
<h4 align="center">A dynamic, general-purpose programming language</h4>
<br>

## Why?

This is a recreational project, so the main goal is having fun, studying language design and exploring various subjects related to it.

## Features

Here a few snippets documenting the feature set of the ape programming language.

#### Variables and types

```
let age = 1;
let name = "Ape";
let result = 10 * (20 / 2);
```

#### Arrays

```
let array = [1, 2, 3, 4, 5];
array[0] // => 1
```

#### Maps (Hashes)

```
let person = {"name": "John", "age": 25};
person["name"] // => "John"
```

#### Functions

```
let sum = fn(a, b) { return a + b; };
sum(1, 2);
```

#### Functions (implicit return)

```
let prod = fn(a, b) { a * b; };
prod(1, 2);
```

#### Functions (real-world example)

```
let fibonacci = fn(n) {
  if (n == 0) {
    0
  } else {
    if (n == 1) {
      1
    } else {
      fibonacci(n - 1) + fibonacci(n - 2);
    }
  }
};
```

#### Functions (higher-order)

```
let twice = fn(f, x) {
  return f(f(x));
};

let increment = fn(x) {
  return x + 1;
};

twice(increment, 5); // => 7
```

## Status

Currently the [lexer](./src/lexer), [ast](./src/ast), [parser](./src/parser) and an [interpreter](./src/interpreter) are implemented. A full-fledged [compiler](./src/compiler) is currently in the works.

## Editor Support

Currently the only editor supporting Ape is the one I am using for developing it. To add support to VSCode for Ape, install the official [Ape Lang](https://marketplace.visualstudio.com/items?itemName=klintmane.ape-lang) extension.

## Requirements

This project requires zero external dependencies, except the Go language compiler if you ever want to build it.

## Contributing

As this language is still being actively designed and developed, contribution would not be practical. That said, fixes and improvements are always welcome.
