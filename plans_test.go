package main

import "github.com/cucumber/godog"

func iEat(arg1 int) error {
        return godog.ErrPending
}

func thereAreGodogs(arg1 int) error {
        return godog.ErrPending
}

func thereShouldBeRemaining(arg1 int) error {
        return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Given(`^there are (\d+) godogs$`, thereAreGodogs)
	ctx.When(`^I eat (\d+)$`, iEat)
	ctx.Then(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
}