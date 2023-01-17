package spend

import (
	"geektrust/account"
	"geektrust/common"
	"geektrust/variables"
)

func Spend(amount string, lender string, owes []string) {
	if len(owes) < 1 {
		return
	}

	splitedAmount := splitAmount(amount, len(owes)+1)
	adjustSpends(lender, owes, splitedAmount)
}

func splitAmount(amount string, n int) int {
	amountInt := common.ConvertToInteger(amount)
	splitAmount := amountInt / n
	return splitAmount
}

func adjustSpends(lender string, owes []string, splitedAmount int) {

	for i := 0; i < len(owes); i++ {
		if !account.IsMemberExists(owes[i]) {
			common.FormattedOutput(variables.MEMBER_NOT_FOUND)
			return
		}

		account.UpdateDues(owes[i], lender, splitedAmount)

		lenderSpend := account.PreviousSpends(lender)
		owerSpend := account.PreviousSpends(owes[i])
		updateSpend(lenderSpend, owerSpend, lender)
		account.UpdateLenter(lender, splitedAmount*len(owes))
	}

	common.FormattedOutput(variables.SUCCESS_TEXT)
}

func updateSpend(m1 map[string]*int, m2 map[string]*int, lender string) {
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; ok {
			if *v2 != variables.ZERO_VALUE {
				remainingAmount(v1, v2, m2[lender])
			}
		}
	}
}

func remainingAmount(v1, v2, v3 *int) {

	if *v2 == variables.ZERO_VALUE || *v3 == variables.ZERO_VALUE {
		return
	}

	if *v1 <= *v3 {
		*v2 += *v1
		*v3 -= *v1
		*v1 = variables.ZERO_VALUE
		return
	}

	remainingAmt := *v1 - *v3

	*v1 = remainingAmt
	*v2 += *v3
	*v3 = variables.ZERO_VALUE
}
