# Project Euler Solutions

Solutions to Project Euler problems implemented in Go.

## Goals

- Solve Project Euler problems using Go
- Practice algorithms and problem solving
- Analyze time and space complexity
- Benchmark solutions where useful
- Improve solutions through refactoring and optimization

## Problems

| Problem | Solution                        | Benchmark                           |
| ------- | ------------------------------- | ----------------------------------- |
| 1       | [Problem 1](problem1/main.go)   | -                                   |
| 2       | [Problem 2](problem2/main.go)   | -                                   |
| 3       | [Problem 3](problem3/main.go)   | -                                   |
| 4       | [Problem 4](problem4/main.go)   | [Benchmark](problem4/main_test.go)  |
| 5       | [Problem 5](problem5/main.go)   | [Benchmark](problem5/main_test.go)  |
| 6       | [Problem 6](problem6/main.go)   | [Benchmark](problem6/main_test.go)  |
| 7       | [Problem 7](problem7/main.go)   | [Benchmark](problem7/main_test.go)  |
| 8       | [Problem 8](problem8/main.go)   | [Benchmark](problem8/main_test.go)  |
| 9       | [Problem 9](problem9/main.go)   | [Benchmark](problem9/main_test.go)  |
| 10      | [Problem 10](problem10/main.go) | [Benchmark](problem10/main_test.go) |

More problems will be added progressively.

## Language

- Go

## Project Structure

```text
.
├── problem1/
│   └── main.go
├── problem2/
│   └── main.go
├── ...
├── problem10/
│   ├── main.go
│   └── main_test.go
├── go.mod
├── LICENSE
├── .gitignore
├── .gitattributes
├── new-problem.sh
└── README.md

```

## Running a Solution

From the repository root:

```bash
go run ./problem10
```

Replace `problem10` with the problem you want to run.

## Running Benchmarks

For problems with benchmarks:

```bash
go test ./problem10 -bench=.
```

For all tests and benchmarks:

```bash
go test ./... -bench=.
```

## Approach

Solutions are written with an emphasis on:

- Clear and idiomatic Go
- Algorithmic efficiency
- Time and space complexity
- Benchmarking
- Incremental optimization

The goal is not only to solve each problem, but also to understand why a solution performs the way it does.

## Disclaimer

This repository contains my own solutions to Project Euler problems.

Project Euler problem statements and related materials remain the property of Project Euler. Problem statements are not reproduced in this repository.

## License

The source code in this repository is licensed under the MIT License.
See [LICENSE](LICENSE) for details.
