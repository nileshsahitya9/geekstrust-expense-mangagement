package calculation

import (
	"geektrust/account"
	"geektrust/lent"
	"geektrust/spend"
	"geektrust/variables"
)

func Calculation(userInput [][]string) {

	for _, input := range userInput {
		switch input[0] {
		case variables.MOVE_IN:
			account.Add(input[1])
			lent.CreateLent(input[1], 0)
		case variables.SPEND:
			spend.Spend(input[1], input[2], input[3:])
		case variables.DUES:
			account.Dues(input[1])
		case variables.CLEAR_DUE:
			account.ClearDues(input[1], input[2], input[3])
		case variables.MOVE_OUT:
			account.Remove(input[1])
		}
	}

}
