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
		resultStr += eqStr + " Earthquakes + " + ltStr + " Lightning"
		if o.FireballUsed {
			resultStr += " + Fireball"
		}
		if o.GiantarrowUsed {
			resultStr += " + Giantarrow"
		}
		resultStr += "\n"
	}
	return resultStr
}
