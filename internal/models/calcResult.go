package models

import "strconv"

type CalcResult struct {
	FireballUsed   bool
	GiantarrowUsed bool
	ZapQuakeCombos map[int]int
}

func (o CalcResult) ToString() string {
	resultStr := ""
	for eq, lt := range o.ZapQuakeCombos {
		eqStr := strconv.Itoa(eq)
		ltStr := strconv.Itoa(lt)
		if lt > 0 {
			resultStr += "Lightning:" + ltStr
		}
		if eq > 0 {
			if lt > 0 {
				resultStr += " + "
			}
			resultStr += "Earthquakes:" + eqStr
		}
		if o.FireballUsed {
			if lt > 0 || eq > 0 {
				resultStr += " + "
			}
			resultStr += "Fireball"
		}
		if o.GiantarrowUsed {
			if lt > 0 || eq > 0 || o.FireballUsed {
				resultStr += " + "
			}
			resultStr += "Giantarrow"
		}
		resultStr += "\n"
	}
	return resultStr
}
