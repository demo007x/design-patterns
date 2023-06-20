# Go **责任链**模式讲解和代码示例

**责任链**是一种行为设计模式， 允许你将请求沿着处理者链进行发送， 直至其中一个处理者对其进行处理。

该模式允许多个对象来对请求进行处理， 而无需让发送者类与具体接收者类相耦合。 链可在运行时由遵循标准处理者接口的任意处理者动态生成。

## 概念示例

让我们来看看一个医院应用的责任链模式例子。 医院中会有多个部门， 如：

- 前台
- 医生
- 药房
- 收银

病人来访时， 他们首先都会去前台， 然后是看医生、 取药， 最后结账。 也就是说， 病人需要通过一条部门链， 每个部门都在完成其职能后将病人进一步沿着链条输送。

此模式适用于有多个候选选项处理相同请求的情形， 适用于不希望客户端选择接收者 （因为多个对象都可处理请求） 的情形， 还适用于想将客户端同接收者解耦时。 客户端只需要链中的首个元素即可。

正如示例中的医院， 患者在到达后首先去的就是前台。 然后根据患者的当前状态， 前台会将其指向链上的下一个处理者。

####  **department.go:** 处理者接口

```go
package main

type Department interface {
	execute(*Patient)
	setNext(Department)
}

```

####  **reception.go:** 具体处理者

```go
package main

import "fmt"

// 前台
type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

```

####  **doctor.go:** 具体处理者

```go
package main

import "fmt"

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}

	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

```

####  **medical.go:** 具体处理者

```go
package main

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

```

####  **cashier.go:** 具体处理者

```go
package main

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

```

####  **patient.go**

```go
package main

type Patient struct {
	name              string
	registrationDone  bool // 注册状态
	doctorCheckUpDone bool // 医生是否检查完成
	medicineDone      bool // 是否取完了药品
	paymentDone       bool // 是否已经支付
}

```

####  **main.go:** 客户端代码

```go
package main

func main() {
	cashier := &Cashier{}

	// set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}

	//Patient visiting
	reception.execute(patient)
}

```

####  **output.txt:** 执行结果

```html
Reception registering patient
Doctor checking patient
Medical giving medicine to patient
Cashier getting money from patient patient
```