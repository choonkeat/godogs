# Getting started

type `make`

```
$ make
godog
Feature: Addition

  Scenario: valid add          # features/addition.feature:3
    Given user enter "4, 4, +" # addition_test.go:6 -> userEnter
      TODO: write pending definition
    Then result is 8           # addition_test.go:10 -> resultIs

1 scenarios (1 pending)
2 steps (1 pending, 1 skipped)
156.678µs
```
