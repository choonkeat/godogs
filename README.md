# Getting started

type `make`


```
$ make
godog
Feature: eat godogs
  In order to be happy
  As a hungry gopher
  I need to be able to eat godogs

  Scenario: Eat 5 out of 12          # features/godogs.feature:6
    Given there are 12 godogs        # godogs_test.go:6 -> thereAreGodogs
      TODO: write pending definition
    When I eat 5                     # godogs_test.go:10 -> iEat
    Then there should be 7 remaining # godogs_test.go:14 -> thereShouldBeRemaining

1 scenarios (1 pending)
3 steps (1 pending, 2 skipped)
245.822µs
```
