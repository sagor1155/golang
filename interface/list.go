package main

// an abstract type, a protocol, a contract
// no implementation
type printer interface {
	print()
}

// list is a slice of printer
type list []printer

func (l list) print() {
	for _, item := range l {
		item.print()
	}
}

func (l list) discount(ratio float64) {
	type discounter interface {
		discount(float64)
	}

	for _, item := range l {
		// g, isGame := item.(*game) // extract dynamic value
		// if isGame {
		// 	g.discount(ratio)
		// }
		// t, isToy := item.(*toy) // extract dynamic value
		// if isToy {
		// 	t.discount(ratio)
		// }

		// or,

		// check if the item has discount method
		// g, ok := item.(interface{ discount(float64) })
		// if !ok {
		// 	continue
		// }
		// g.discount(ratio)

		// or,

		//type assertion: used to check whether an interface value provides the method that you want
		// g, ok := item.(discounter)
		// if !ok {
		// 	continue
		// }
		// g.discount(ratio)

		// or,

		if item, ok := item.(discounter); ok {
			item.discount(ratio)
		}

	}
}
