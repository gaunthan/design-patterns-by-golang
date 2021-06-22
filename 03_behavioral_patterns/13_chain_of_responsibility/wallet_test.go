package chainofresponsibility

import "fmt"

func Example_one() {
	wallet := NewWallet("Wechat Pay", 100)
	wallet.Chain(
		NewWallet("Alipay", 80)).Chain(
		NewWallet("Bank Card", 200))

	outputFailedPayment(wallet.Pay(50), 50)
	outputFailedPayment(wallet.Pay(100), 100)
	outputFailedPayment(wallet.Pay(60), 60)
	outputFailedPayment(wallet.Pay(150), 150)

	// Output:
	// Wechat Pay: succeed to deduct 50$
	// Bank Card: succeed to deduct 100$
	// Alipay: succeed to deduct 60$
	// error: wallet balance less then 150$
}

func outputFailedPayment(ok bool, amount int) {
	if !ok {
		fmt.Printf("error: wallet balance less then %d$\n", amount)
	}
}
