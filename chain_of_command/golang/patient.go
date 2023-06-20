package main

type Patient struct {
	name              string
	registrationDone  bool // 注册状态
	doctorCheckUpDone bool // 医生是否检查完成
	medicineDone      bool // 是否取完了药品
	paymentDone       bool // 是否已经支付
}
