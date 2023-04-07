package dto

type NodeStatus interface {
	// 判断节点是否 跳跃/回退
	IsNextOK(nextNode string) bool
	// 返回 node 对应的 int
	GetStatusInt(node string) int
}

// shipment 节点
var ChainShipmentNodeMap = map[string]int{
	"info_hash":             0,
	"qr_code_hash":          1,
	"shipment_start_hash":   2,
	"shipment_receive_hash": 3,
	"shipment_enter_hash":   4,
}

// shipment
func (u *UploadChainShipment) IsNextOK(nextNode string) bool {
	if u.UploadNode == "" {
		// 新建
		return true
	}
	return u.GetStatusInt(u.UploadNode) == u.GetStatusInt(nextNode) ||
		u.GetStatusInt(u.UploadNode)+1 == u.GetStatusInt(nextNode)
}

func (u *UploadChainShipment) GetStatusInt(node string) int {
	return ChainShipmentNodeMap[node]
}

// shipmentOrder node
var ChainShipmentOrderNodeMap = map[string]int{
	"qr_code_hash":          0,
	"shipment_start_hash":   1,
	"shipment_receive_hash": 2,
	"shipment_enter_hash":   3,
}

func (u *UploadChainShipmentOrder) IsNextOK(nextNode string) bool {
	if u.UploadNode == "" {
		// 新建
		return true
	}
	return u.GetStatusInt(u.UploadNode) == u.GetStatusInt(nextNode) ||
		u.GetStatusInt(u.UploadNode)+1 == u.GetStatusInt(nextNode)
}

func (u *UploadChainShipmentOrder) GetStatusInt(node string) int {
	return ChainShipmentOrderNodeMap[node]
}

// 付款单
// 付款单节点
var ChainPayOrderNodeMap = map[string]int{
	"info_hash":   0,
	"result_hash": 1,
}

func (u *UploadChainPayOrder) IsNextOK(nextNode string) bool {
	if u.UploadNode == "" {
		return true
	}
	return u.GetStatusInt(u.UploadNode) == u.GetStatusInt(nextNode) ||
		u.GetStatusInt(u.UploadNode)+1 == u.GetStatusInt(nextNode)
}

func (u *UploadChainPayOrder) GetStatusInt(node string) int {
	return ChainPayOrderNodeMap[node]
}

// 保理单
// 保理单 node

var ChainFactoringOrderNodeMap = map[string]int{
	"apply_hash":    0,
	"result_hash":   1,
	"pay_back_hash": 2,
}

func (u *UploadChainFactoringOrder) IsNextOK(nextNode string) bool {
	if u.UploadNode == "" {
		return true
	}
	return u.GetStatusInt(u.UploadNode) == u.GetStatusInt(nextNode) ||
		u.GetStatusInt(u.UploadNode)+1 == u.GetStatusInt(nextNode)
}

func (u *UploadChainFactoringOrder) GetStatusInt(node string) int {
	return ChainFactoringOrderNodeMap[node]
}
