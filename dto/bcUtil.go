package dto

func GetKeyOfConfig(txID string) string {
	return txID
}

func GetKeyOfShipmentOrder(ShipmentNo string) string {
	return "shipment" + ShipmentNo
}

// key: "company" + MibCode + CompanyCode
func GetKeyOfCompany(MibCode, CompanyCode string) string {
	return "company" + MibCode + CompanyCode
}

// key : "mi" + MibCode + MiCode
func GetKeyOfMi(MibCode, MiCode string) string {
	return "mi" + MibCode + MiCode
}

// key: "contract"+ContractNo
func GetKeyOfContract(ContractNo string) string {
	return "contract" + ContractNo
}

// key:"contractMi" + ContractNo + MiCode
func GetKeyOfContractMi(ContractNo, MiCode string) string {
	return "contractMi" + ContractNo + MiCode
}
func GetKeyOfShipment(OrderNo string) string {
	return "shipPlan" + OrderNo
}
func GetKeyOfPay(ShipmentNo string) string {
	return "pay" + ShipmentNo
}
func GetKeyOfFactoring(FactoringOrderNo string) string {
	return "factoring" + FactoringOrderNo
}
