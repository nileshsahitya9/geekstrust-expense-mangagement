package account

import (
	"fmt"
	"geektrust/common"
	"geektrust/lent"
	"geektrust/variables"
	"sort"
)

type Account struct {
	name   string
	amount int
}

type Members struct {
	name string
}

var accountInstances = make(map[string][]*Account)

var memberInstances = make(map[string]*Members)

// Creation

func Add(name string) {

	if CountMembers() >= variables.MAX_MEMBERS {
		common.FormattedOutput(variables.HOUSEFULL_TEXT)
		return
	}

	AccountCreation(name, variables.ZERO_VALUE)
	memberInstances[name] = &Members{name}
	common.FormattedOutput(variables.SUCCESS_TEXT)
}

func CountMembers() int {
	return len(memberInstances)
}

func AccountCreation(name string, amount int) {

	members := GetAllMembers()
	for _, element := range members {
		CreateDues(element, name, amount)
		CreateDues(name, element, amount)
	}

}

func CreateDues(name string, lender string, amount int) {
	accountInstances[name] = append(accountInstances[name], &Account{lender, amount})
}

func Remove(name string) {
	if !isEligibleToRemove(name) {
		common.FormattedOutput(variables.FAILURE_TEXT)
		return
	}

	delete(accountInstances, name)
	delete(memberInstances, name)
	common.FormattedOutput(variables.SUCCESS_TEXT)
}

func GetMember(name string) *Members {
	if _, ok := memberInstances[name]; ok {
		return memberInstances[name]
	}
	return nil
}

func IsMemberExists(name string) bool {
	member := GetMember(name)

	if member == nil {
		return false
	}

	return true
}

func GetAllMembers() []string {
	members := make([]string, 0)

	for key, _ := range memberInstances {
		members = append(members, key)
	}

	return members

}

func Dues(name string) {
	owerDues := GetDues(name)
	keys := make([]string, 0, len(owerDues))

	for k := range owerDues {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	sort.SliceStable(keys, func(i, j int) bool {
		return *owerDues[keys[i]] > *owerDues[keys[j]]
	})

	for _, k := range keys {
		fmt.Println(k, *owerDues[k])
	}

}

func ClearDues(ower, lender, amount string) {
	amountInt := common.ConvertToInteger(amount)
	if isTransactionValid(ower, lender, amountInt) {
		updateTransaction(ower, lender, amountInt)
	}

}

func GetDues(name string) map[string]*int {
	if _, ok := accountInstances[name]; ok {
		debt := make(map[string]*int)
		for i := 0; i < len(accountInstances[name]); i++ {
			value := accountInstances[name][i]
			debt[value.name] = &value.amount
		}
		return debt
	}

	return nil
}

func isTransactionValid(ower string, lender string, amount int) bool {

	owerDues := GetDues(ower)

	if owerDues == nil {
		common.FormattedOutput(variables.INCORRECT_PAYMENT_TEXT)
		return false
	}

	var duesAmount int
	if v, ok := owerDues[lender]; ok {
		duesAmount = *v
	} else {
		common.FormattedOutput(variables.INCORRECT_PAYMENT_TEXT)
		return false
	}

	if amount > duesAmount {
		common.FormattedOutput(variables.INCORRECT_PAYMENT_TEXT)
		return false
	}

	return true

}

func updateTransaction(ower string, lender string, amount int) {
	owerDues := GetDues(ower)
	var remainingAmount int
	if v, ok := owerDues[lender]; ok {
		*v -= amount
		remainingAmount = *v
	}

	UpdateLenter(lender, remainingAmount)

	common.FormattedOutput(remainingAmount)
}

func isEligibleToRemove(name string) bool {
	lentAmount := getLentAmount(name)

	if lentAmount != variables.ZERO_VALUE {
		return false
	}

	owerDues := GetDues(name)

	for _, value := range owerDues {
		if *value != variables.ZERO_VALUE {
			return false
		}
	}

	return true
}

func PreviousSpends(name string) map[string]*int {
	spend := GetDues(name)
	return spend
}

func UpdateDues(name string, lender string, amount int) {
	accounts := accountInstances[name]

	for _, element := range accounts {
		if element.name == lender {

			element.amount += amount
		}
	}

}

func UpdateLenter(lender string, amount int) {
	lenter := lent.GetInstance(lender)

	if lenter == nil {
		common.FormattedOutput(variables.MEMBER_NOT_FOUND)
		return
	}

	lenter.Update(amount)
}

func getLentAmount(name string) int {
	lenter := lent.GetInstance(name)

	if lenter == nil {
		return -1
	}

	return lenter.GetAmount()
}
