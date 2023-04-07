package dto

import "time"

// base -------------- start
type UploadChainType struct {
	//上一次的上链hash，首次上链为空
	LastHash string `json:"lastHash"`
	//上链流程节点,每个节点不可跳跃，同一节点可以反复更新，进入下一个节点，不可返回上一节点
	UploadNode       string `json:"uploadNode"`
	UploaderName     string `json:"uploaderName"`
	UploaderRole     string `json:"uploaderRole"`
	UploaderRoleName string `json:"uploaderRoleName"`
	UploaderAccount  string `json:"uploaderAccount"`
}
type BankAccount struct {
	CardNo      string `json:"cardNo"`
	AccountType string `json:"accountType"`
	Bank        string `json:"bank"`
	Toibkn      string `json:"toibkn"`
	CardOwner   string `json:"cardOwner"`
	//银行卡医保局是否列为黑名单，此时不允许支付
	ValidStatus string `json:"validStatus"`
}

// base -------------- end

//配送企业上链信息记录 生效时上链 企业社会识别编码不可修改 其他内容更新可重复上链
// key: "company" + MibCode + CompanyCode
type UploadChainCompany struct {
	UploadChainType
	CompanyName string `json:"companyName"`
	//企业社会识别编码不可修改
	CompanyCode  string `json:"companyCode"`
	ManagerName  string `json:"managerName"`
	ManagerPhone string `json:"managerPhone"`
	//授权所属医保局code
	MibCode string `json:"mibCode"`
	//授信总额度
	CreditAmount int64 `json:"creditAmount"`
	//授信可用时间段  每个企业只会有一个授信设置，无效的时间段一律视为授信额度为0
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
	//配送企业专户
	BankAccount BankAccount
}

//医疗机构上链信息记录 生效时上链 企业社会识别编码不可修改 其他内容更新可重复上链
// key : "mi"+ MibCode + MiCode
type UploadChainMi struct {
	UploadChainType
	//授权所属医保局code
	MibCode string `json:"mibCode"`
	//社会识别编码 不可修改
	MiName string `json:"miName"`
	MiCode string `json:"miCode"`
	//机构类别
	MiType       string `json:"miType"`
	ManagerName  string `json:"managerName"`
	ManagerPhone string `json:"managerPhone"`
	//专户额度
	LimitBalance int64 `json:"limitBalance"`
	//专户
	SpecialAccount BankAccount
	//一般户
	NormalAccount BankAccount
}

//药品信息上链 药品通用code 不可修改
// key : "med" + MedicineCode
type UploadChainMedicine struct {
	UploadChainType
	MedicineName string `json:"medicineName"`
	//药品通用code 唯一键
	MedicineCode string `json:"medicineCode"`
	//省药品编码
	MedicinePCode string `json:"medicinePCode"`
	//药品规格（如20mg,2ml:200ug）
	MedicineSpecification string `json:"medicineSpecification"`
	//剂型（如片剂）
	MedicineForm string `json:"medicineForm"`
	//药品包装(如盒，支，瓶)
	MedicinePackage string `json:"medicinePackage"`
	//包装单位
	MedicineUnit  string `json:"medicineUnit"`
	CalculateRate int64  `json:"calculateRate"`
	Manufacturer  string `json:"manufacturer"`
}

//配置信息上链
type UploadChainConfig struct {
	UploadChainType
	//具体的配置key，有些涉及判断的key会单独提供key
	Key       string    `json:"key"`
	TypeOne   string    `json:"typeOne"`
	TypeTwo   string    `json:"typeTwo"`
	TypeThree string    `json:"typeThree"`
	Value     string    `json:"value"`
	StartAt   time.Time `json:"startAt"`
	EndAt     time.Time `json:"endAt"`
	Explain   string    `json:"explain"`
}

//目前合同上的金额不作为限制依据了， 主合同在生效后 即为敲定，不可修改
// key: "contract" + ContractNo
type UploadChainContract struct {
	UploadChainType
	ContractNo    string `json:"contractNo"`
	TenderCompany string `json:"tenderCompany"`
	//药品通用code 唯一键
	MedicineCode string `json:"medicineCode"`
	//省药品编码
	MedicinePCode string `json:"medicinePCode"`
	//配送企业社会识别编码不可修改
	CompanyCode string `json:"companyCode"`
	//所属医保局code
	MibCode       string `json:"mibCode"`
	MedicinePrice int64  `json:"medicinePrice"`
	//各个企业合计总采购量（包装）
	TotalAmount int64     `json:"totalAmount"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
}

//子合同信息 主合同在生效后可生效 生效即为敲定，不可修改，需要校验主合同已经存在
// key: "contractMi"+ ContractNo + MiCode
type UploadChainContractMi struct {
	UploadChainType
	ContractNo string `json:"contractNo"`
	//医院社会识别编码 不可修改
	MiCode string `json:"miCode"`
	//约定采购量(包装)
	MedicineAmount int64 `json:"medicineAmount"`
	//合同文件md5
	ContractMd5 string `json:"contractMd5"`
	//合同文件地址
	ContractUrl string `json:"contractUrl"`
}

//配送计划上链 // key: "shipPlan"+OrderNo
type UploadChainShipment struct {
	UploadChainType
	//招采单号
	TpOrderNo string `json:"tpOrderNo"`
	//平台单号
	OrderNo string `json:"orderNo"`
	//数据来源平台code
	TpCode string `json:"tpCode"`
	//授权所属医保局code
	MibCode        string `json:"mibCode"`
	ShipmentAmount int64  `json:"shipmentAmount"`
	OrderAmount    int64  `json:"orderAmount"`
	//所属合同与医院
	ContractNo string `json:"contractNo"`
	MiCode     string `json:"miCode"`
	//药品通用code 唯一键
	MedicineCode string `json:"medicineCode"`
	//省药品编码
	MedicinePCode string `json:"medicinePCode"`
	//配送企业社会识别编码不可修改
	CompanyCode string `json:"companyCode"`
	//招采同步时间
	SyncAt time.Time `json:"syncAt"`
}

//配送单相关 // key: "shipment"+ShipmentNo
type UploadChainShipmentOrder struct {
	UploadChainType
	Plans []UploadChainShipment `json:"plans"`
	//配送单号
	ShipmentNo       string    `json:"shipmentNo"`
	InvoiceMd5       string    `json:"invoiceMd5"`
	InvoiceUrl       string    `json:"invoiceUrl"`
	InvoiceType      string    `json:"invoiceType"`
	InvoiceTotal     int64     `json:"invoiceTotal"`
	InvoiceTax       int64     `json:"invoiceTax"`
	InvoiceCheckCode string    `json:"invoiceCheckCode"`
	InvoiceNo        string    `json:"invoiceNo"`
	InvoiceCode      string    `json:"invoiceCode"`
	InvoiceDate      time.Time `json:"invoiceDate"`
	//发货时间&生成配送码时间
	SendOutAt       time.Time `json:"sendOutAt"`
	ShipmentCode    string    `json:"shipmentCode"`
	ShipmentCompany string    `json:"shipmentCompany"`
	// 配送状态 WAITING SENDING  RECEIVED FAIL
	ShipmentStatus string `json:"shipmentStatus"`

	//扫码收货时间
	ReceiveAt time.Time `json:"receiveAt"`
	//入库时间 目前无
	ConfirmAt time.Time `json:"confirmAt"`
}

//支付单 上链 // key: "Pay" + ShipmentNo
type UploadChainPayOrder struct {
	UploadChainType
	//支付的配送单号
	ShipmentNo string `json:"shipmentNo"`
	//配送单总金额 分
	OrderAmount int64  `json:"orderAmount"`
	MiCode      string `json:"miCode"`
	CompanyCode string `json:"companyCode"`
	//支付结果节点
	//平台支付号
	OutTradeNo string    `json:"outTradeNo"`
	PayType    string    `json:"payType"`
	PayAt      time.Time `json:"payAt"`
	//支付状态 理论上只有支付失败时，平台会告知线下处置完毕
	PayStatus string `json:"payStatus"`
	// 支付单下层私有状态, 记录是否有保理信息(有保理信息则处于保理中, 没有则不管)
	FactoringStatus string `json:"factoringStatus"`
}

type PayOrder struct {
	ShipmentNo string `json:"shipmentNo"`
}

//保理单 上链 // key: "Factoring"+FactoringOrderNo
type UploadChainFactoringOrder struct {
	UploadChainType
	//申请节点（基本信息节点）
	//保理单使用什么还款单做基础，只有尚未回款的订单可以做担保
	PayOrders []PayOrder
	//谁保理(默认是医院货款给企业，企业去还款所以不涉及谁还款)
	CompanyCode    string `json:"companyCode"`
	FactoringMoney int64  `json:"factoringMoney"`
	//平台保理申请号
	FactoringOrderNo string    `json:"factoringOrderNo"`
	ApplyDate        time.Time `json:"applyDate"`
	ApplyAt          time.Time `json:"applyAt"`
	RefundDate       time.Time `json:"refundDate"`

	//保理申请结果节点
	FactoringBankNo string `json:"factoringBankNo"`
	//实际收到金额
	ActualReceiveMoney int64 `json:"actualReceiveMoney"`
	//保理状态 理论上此项为银行推送，平台自身不会上送此状态
	FactoringStatus string `json:"factoringStatus"`

	//保理还款结果节点
	RefundAt time.Time `json:"refundAt"`
	//仅线下处置时由平台推送，其他状态为银行推送
	RefundStatus string `json:"refundStatus"`
}
