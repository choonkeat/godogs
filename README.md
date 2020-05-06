# Getting started

type `make`

```
$ make
godog
Feature: Addition

  Scenario: valid add          # features/addition.feature:3
    Given user enter "4, 4, +"
    Then result is 8

1 scenarios (1 undefined)
2 steps (2 undefined)
133.646µs

You can implement step definitions for undefined steps with these snippets:

func userEnter(arg1 string) error {
	return godog.ErrPending
}

func resultIs(arg1 int) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^user enter "([^"]*)"$`, userEnter)
	s.Step(`^result is (\d+)$`, resultIs)
}
```

godog will detect missing functions and print some code snippets.
