## Lunar Plan Advisor

[![Test](https://github.com/romietis/lunar-plan-advisor/actions/workflows/test.yaml/badge.svg)](https://github.com/romietis/lunar-plan-advisor/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/romietis/lunar-plan-advisor/v3)](https://goreportcard.com/report/github.com/romietis/lunar-plan-advisor/v3)
[![codecov](https://codecov.io/github/romietis/lunar-plan-advisor/graph/badge.svg?token=7TL3J6GCYT)](https://codecov.io/github/romietis/lunar-plan-advisor)
[![Go Reference](https://pkg.go.dev/badge/github.com/romietis/lunar-plan-advisor/v3.svg)](https://pkg.go.dev/github.com/romietis/lunar-plan-advisor/v3)
[![GitHub Release](https://img.shields.io/github/v/release/romietis/lunar-plan-advisor)](https://github.com/romietis/lunar-plan-advisor/releases)


*Currently only for Lunar Denmark*

This is a simple web application that helps you to choose the best Lunar plan based on
your savings. It is built with [Gin](https://gin-gonic.com/) and basic HTML/CSS/JavaScript.

## How to run

### Web-Server

```bash
go run cmd/web/main.go
```

or run with Docker

```bash
docker build -t lunar-plan-advisor .
docker run -p 8080:8080 lunar-plan-advisor
```
Your application will be available at localhost:8080 and 0.0.0.0:8080

### Tests

Run unit tests
```bash
go test -v -shuffle=on -coverprofile=cover.out ./internal/endpoints/... ./advisor/...
```

View coverage report
```bash
go tool cover -html=cover.out
```

Run BDD tests
```bash
go test -v ./internal/bdd/...
```

## API
Now exposing API endpoint with query parameter `balance`

```bash
curl https://lunar-plan-advisor.calmground-6bcda4d8.northeurope.azurecontainerapps.io/plans?balance=100000
```

## Background

With Lunar you receive positive interest rate on active balance of your accounts.

Lunar has an
[interest rate calculator](https://www.lunar.app/en/personal/positive-interest-rate-lunar)
to help you see what you can earn with different amounts and plans.
But it doesn't tell you which plan is most profitable based on *your* savings.

*Lunar Plan Advisor* helps you determine a plan based on the
*net profit* - interest income after plan fees.
It helps identify the point at which one plan becomes more profitable than another by
comparing the net profits of different plans.

## Privacy
Your data is not stored or sent to anyone.


