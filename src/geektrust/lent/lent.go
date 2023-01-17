package lent

type Lent struct {
	amount int
}

var lentInstances = make(map[string]*Lent)

func CreateLent(name string, amount int) {
	lentInstances[name] = &Lent{amount: amount}
}

func (self *Lent) Update(amount int) {
	self.amount = amount
}

func (self *Lent) GetAmount() int {
	return self.amount
}

func GetInstance(name string) *Lent {
	return lentInstances[name]
}
