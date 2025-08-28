package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
		newUnitMap := map[string]int{}
    	newUnitMap["quarter_of_a_dozen"] = 3
		newUnitMap["half_of_a_dozen"] =	6
		newUnitMap["dozen"] = 12
		newUnitMap["small_gross"] =	120
		newUnitMap["gross"] = 144
		newUnitMap["great_gross"] =	1728
    return newUnitMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
    if exists == true{
        bill[item] += value
    }
    return exists
    
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, exists := bill[item]
    if exists{
    	removedAmount ,exists := units[unit]
        if exists{
            newAmount := bill[item] - removedAmount
            if newAmount < 0{
                return false
            }else if newAmount == 0{
                delete(bill, item)
            }else{
                bill[item] = newAmount
            }
        }
        return exists
    }
	return exists
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
    return qty, ok
}
