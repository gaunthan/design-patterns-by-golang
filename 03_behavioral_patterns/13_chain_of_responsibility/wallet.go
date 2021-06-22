package chainofresponsibility

import "fmt"

type Wallet interface {
	Balance() int
	Pay(amount int) bool
	Chain(wallet Wallet) Wallet
}

type walletImpl struct {
	balance int
	next    Wallet
}

func newWalletImpl(balance int, next Wallet) walletImpl {
	return walletImpl{
		balance: balance,
		next:    next,
	}
}

func (p *walletImpl) Balance() int {
	return p.balance
}

func (p *walletImpl) Pay(amount int) bool {
	return false
}

func (p *walletImpl) Chain(wallet Wallet) Wallet {
	p.next = wallet
	return p.next
}

type WechatPay struct {
	walletImpl
}

func (p *WechatPay) Pay(amount int) bool {
	if p.balance < amount {
		return p.next != nil && p.next.Pay(amount)
	}
	p.balance -= amount
	fmt.Printf("Wechat Pay: succeed to deduct %d$\n", amount)
	return true
}

type AliPay struct {
	walletImpl
}

func (p *AliPay) Pay(amount int) bool {
	if p.balance < amount {
		return p.next != nil && p.next.Pay(amount)
	}
	p.balance -= amount
	fmt.Printf("Alipay: succeed to deduct %d$\n", amount)
	return true
}

type BankCardPay struct {
	walletImpl
}

func (p *BankCardPay) Pay(amount int) bool {
	if p.balance < amount {
		return p.next != nil && p.next.Pay(amount)
	}
	p.balance -= amount
	fmt.Printf("Bank Card: succeed to deduct %d$\n", amount)
	return true
}

func NewWallet(wallet string, balance int) Wallet {
	switch wallet {
	case "Wechat Pay":
		return &WechatPay{walletImpl: newWalletImpl(balance, nil)}
	case "Alipay":
		return &AliPay{walletImpl: newWalletImpl(balance, nil)}
	case "Bank Card":
		return &BankCardPay{walletImpl: newWalletImpl(balance, nil)}
	default:
		panic(fmt.Sprintf("unknown wallet: %s", wallet))
	}
}
