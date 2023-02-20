package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	billValue, billExists := bill[item]

	if exists {
		if !billExists {
			bill[item] = value
		} else {
			bill[item] = billValue + value
		}
		return true
	} else {
		return false
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExists := units[unit]
	billValue, billExists := bill[item]

	if !unitExists || !billExists || unitValue > billValue {
		return false
	} else if billValue == unitValue {
		delete(bill, item)
	} else {
		bill[item] -= unitValue
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (value int, exists bool) {
	value, exists = bill[item]
	return
}
