package golang_dependency_one_b

import gdp "github.com/mishamolnar/golang-dependency-two"

func Produce() string {
	return gdp.CapitalizeString("one", "two", " ")
}
