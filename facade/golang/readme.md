# Go **外观**模式讲解和代码示例

**外观**是一种结构型设计模式， 能为复杂系统、 程序库或框架提供一个简单 （但有限） 的接口。

尽管外观模式降低了程序的整体复杂度， 但它同时也有助于将不需要的依赖移动到同一个位置。

## 概念示例

人们很容易低估使用信用卡订购披萨时幕后工作的复杂程度。 在整个过程中会有不少的子系统发挥作用。 下面是其中的一部分：

- 检查账户
- 检查安全码
- 借记/贷记余额
- 账簿录入
- 发送消息通知

在如此复杂的系统中， 可以说是一步错步步错， 很容易就会引发大的问题。 这就是为什么我们需要外观模式， 让客户端可以使用一个简单的接口来处理众多组件。 客户端只需要输入卡片详情、 安全码、 支付金额以及操作类型即可。 

外观模式会与多种组件进一步地进行沟通， 而又不会向客户端暴露其内部的复杂性。

####  **walletFacade.go:** 外观

```go
package main

import (
	"fmt"
)

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func newWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	var walletFacade = &WalletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &Notification{},
		ledger:       &Ledger{},
	}
	fmt.Println("Account created")
	return walletFacade
}

func (w *WalletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Start add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) deductMoneyFromWallet(accountID string, securityCode, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "debit", amount)
	return nil
}

```

####  **account.go:** 复杂子系统的组成部分

```go
package main

import "fmt"

type Account struct {
	name string
}

func newAccount(name string) *Account {
	return &Account{
		name: name,
	}
}

func (a *Account) checkAccount(name string) error {
	if a.name != name {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

```

####  **securityCode.go:** 复杂子系统的组成部分

```go
package main

import "fmt"

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}

func (s *SecurityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}

```

####  **wallet.go:** 复杂子系统的组成部分

```go
package main

import "fmt"

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	w.balance -= amount
	return nil
}

```

####  **ledger.go:** 复杂子系统的组成部分

```go
package main

import "fmt"

type Ledger struct{}

func (s *Ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}

```

####  **notification.go:** 复杂子系统的组成部分

```go
package main

import "fmt"

type Notification struct{}

func (n *Notification) sendWalletCreditNotification() {
	fmt.Println("sending wallet credit notification")
}

func (n *Notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

```

####  **main.go:** 客户端代码

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()
	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("error: %s \n", err.Error())
	}
	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("error: %s \n", err.Error())
	}
}

```

####  **output.txt:** 执行结果

```

```