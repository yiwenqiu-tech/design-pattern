package main

import "design-pattern/factory"

type Factory interface {
	CreateChart() factory.Chart
}
