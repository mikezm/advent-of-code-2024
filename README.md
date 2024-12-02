# advent-of-code-go-template
A simple template for completing [Advent Of Code](https://adventofcode.com) challenges in Golang.

## Usage

The template uses an interface to provide the functions `A()` and `B()` for each day's challenge. 
These can be imported into the `main.go` file as show. 

A `Makefile` is provided to facilitate the generation and cleanup of new days.

Once generated, the day can be imported into the [main.go](main.go) file on line 8 and added to the map on line 18.

I recommend updating forking the repo and updating your [go.mod](go.mod) module to something like:
`github.com/<your-gitgub>/advent-of-code-<year>`

## Generating a new Day's Challenge

Generate a Day 1 challenge

```shell
make gen d=1
```

This creates a new folder for the provided day, as well as a go file and an empty input.txt file. It is assumed you will
copy the contents of the challenge's input file manually. 

## Removing a Day's challenge

```shell
make clean d=1
```

This removes the day's folder and all it's contents. 

## Executing a day's challenge

```shell
go run main.go 1 a
```

```shell
go run main.go 1 B
```
 
