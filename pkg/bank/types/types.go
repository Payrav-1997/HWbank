package types

//Money представляет собой денежую сумму в минимальных кдиницах
type Money int64

//PAN Для карты
type PAN string





//Card пердставляет информацию о платёжной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Active     bool
  }

  //PaymentSource для выбора карты 
  type PaymentSource struct {
	Type   string
	Number string
	Balance Money
}