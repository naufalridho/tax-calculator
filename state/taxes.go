package state

const (
	TaxCodeFood          = 1
	TaxCodeTobacco       = 2
	TaxCodeEntertainment = 3
)

var TaxCodeName = map[int]string{
	TaxCodeFood:          "Food",
	TaxCodeTobacco:       "Tobacco",
	TaxCodeEntertainment: "Entertainment",
}
