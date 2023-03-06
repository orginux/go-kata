# CodeWars Katas
This repository contains my solutions to various katas from CodeWars. Each solution is written in Go and is accompanied by tests to validate the implementation.

![image](https://www.codewars.com/users/orginux/badges/large)

## How to Use This Repository
Each kata is stored in its own directory, named according to the kata's title. Within each directory, you'll find the following files:

- `kata.go`: This file contains the implementation of the solution.
- `kata_test.go`: This file contains the test suite for the solution.

Each kata is linked to its corresponding page on CodeWars, where you can find a description of the challenge.

### Add new kata
```bash
make kata NAME="<kata name>" LINK=<link to kata>
```


## Tests
To test solutions, run the following command:
```bash
make test
```

## License
This repository is licensed under the MIT License, which means you are free to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the code, subject to the conditions in the LICENSE file.
