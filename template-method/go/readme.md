# Go **模板方法**模式讲解和代码示例

**模版方法**是一种行为设计模式， 它在基类中定义了一个算法的框架， 允许子类在不修改结构的情况下重写算法的特定步骤。

## 概念示例

让我们来考虑一个一次性密码功能 （OTP） 的例子。 将 OTP 传递给用户的方式多种多样 （短信、 邮件等）。 但无论是短信还是邮件， 整个 OTP 流程都是相同的：

1. 生成随机的 n 位数字。
2. 在缓存中保存这组数字以便进行后续验证。
3. 准备内容。
4. 发送通知。

后续引入的任何新 OTP 类型都很有可能需要进行相同的上述步骤。

因此， 我们会有这样的一个场景， 其中某个特定操作的步骤是相同的， 但实现方式却可能有所不同。 这正是适合考虑使用模板方法模式的情况。

首先， 我们定义一个由固定数量的方法组成的基础模板算法。 这就是我们的模板方法。 然后我们将实现每一个步骤方法， 但不会改变模板方法。

####  **otp.go:** 模板方法

```go
package main

type IOtp interface {
	genRandomOPT(int) string
	saveOPTCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

func (o *Otp) genAndSendOPT(optLength int) error {
	opt := o.iOtp.genRandomOPT(optLength)
	o.iOtp.saveOPTCache(opt)
	message := o.iOtp.getMessage(opt)
	if err := o.iOtp.sendNotification(message); err != nil {
		return err
	}
	return nil
}

```

####  **sms.go:** 具体实施

```go
package main

import (
	"fmt"
)

type Sms struct {
	Otp
}

func (s *Sms) genRandomOPT(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s \n", randomOTP)
	return randomOTP
}

func (s *Sms) saveOPTCache(otp string) {
	fmt.Printf("SMS: saving otp %s", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

```

####  **email.go:** 具体实施

```go
package main

import (
	"fmt"
)

type Email struct {
	Otp
}

func (s *Email) genRandomOPT(len int) string {
	randomOTP := "2345"
	fmt.Printf("Email: generating random otp %s \n", randomOTP)
	return randomOTP
}

func (s *Email) saveOPTCache(otp string) {
	fmt.Printf("Email: saving otp %s to cache", otp)
}

func (s *Email) getMessage(otp string) string {
	return "Email otp for login is " + otp
}

func (s *Email) sendNotification(message string) error {
	fmt.Printf("Email: sending email %s \n", message)
	return nil
}

```

####  **main.go:** 客户端代码

```go
package main

import "fmt"

func main() {
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOPT(4)

	fmt.Println("")
	emailOtp := &Email{}
	o = Otp{
		iOtp: emailOtp,
	}
	o.genAndSendOPT(4)
}

```

####  **output.txt:** 执行结果

```tex
SMS: generating random otp 1234 
SMS: saving otp 1234SMS: sending sms: SMS OTP for login is 1234

Email: generating random otp 2345 
Email: saving otp 2345 to cacheEmail: sending email Email otp for login is 2345 
```