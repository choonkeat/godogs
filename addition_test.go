package godogs

import "github.com/cucumber/godog"

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
