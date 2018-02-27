package vendingMachine

type VendingMachine struct {
	insertedMoney int
	coins         map[string]int
	items         map[string]int
}

func NewVendingMachine(coins, items map[string]int) *VendingMachine {
	return &VendingMachine{coins: coins, items: items}
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	m.insertedMoney += m.coins[coin]
}

func (m *VendingMachine) SelectSD() string {
	return m.selectItem("SD")
}
func (m *VendingMachine) SelectCC() string {
	return m.selectItem("CC")
}
func (m *VendingMachine) selectItem(item string) string {
	price, ok := m.find(item)
	if !ok {
		return ""
	}
	change := m.insertedMoney - price
	m.insertedMoney -= price + change
	return item + m.change(change)
}

func (m *VendingMachine) find(s string) (price int, ok bool) {
	price, ok = m.items[s]
	return
}
func (m VendingMachine) change(c int) (change string) {
	values := [...]int{10, 5, 2, 1}
	coins := [...]string{"T", "F", "TW", "O"}

	for i := 0; c > 0; i++ {
		if c >= values[i] {
			change += ", " + coins[i]
			c -= values[i]
			i--
		}
	}
	
	return
}
func (m *VendingMachine) CoinReturn() string {
	change := m.change(m.insertedMoney)
	return change[2:len(change)]
}
