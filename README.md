# Go Exercises

## Overview

This repo contains programming exercises in Go language.

- [Project Euler](https://projecteuler.net/)
- [Advent of Code](https://adventofcode.com/)

## Verification

```
gotest ./...
```

### Check test coverage

```
gotest ./... -cover
gotest ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```
