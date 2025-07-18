// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/scg/inventory/v1/enums.proto

package inventoryv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InventoryItemStatus int32

const (
	InventoryItemStatus_INVENTORY_ITEM_STATUS_UNSPECIFIED InventoryItemStatus = 0
	InventoryItemStatus_INVENTORY_ITEM_STATUS_AVAILABLE   InventoryItemStatus = 1
	InventoryItemStatus_INVENTORY_ITEM_STATUS_RESERVED    InventoryItemStatus = 2
	InventoryItemStatus_INVENTORY_ITEM_STATUS_ALLOCATED   InventoryItemStatus = 3
	InventoryItemStatus_INVENTORY_ITEM_STATUS_BACKORDERED InventoryItemStatus = 4
	InventoryItemStatus_INVENTORY_ITEM_STATUS_DAMAGED     InventoryItemStatus = 5
	InventoryItemStatus_INVENTORY_ITEM_STATUS_EXPIRED     InventoryItemStatus = 6
	InventoryItemStatus_INVENTORY_ITEM_STATUS_QUARANTINED InventoryItemStatus = 7
	InventoryItemStatus_INVENTORY_ITEM_STATUS_RECALLED    InventoryItemStatus = 8
	InventoryItemStatus_INVENTORY_ITEM_STATUS_DISPOSED    InventoryItemStatus = 9
)

// Enum value maps for InventoryItemStatus.
var (
	InventoryItemStatus_name = map[int32]string{
		0: "INVENTORY_ITEM_STATUS_UNSPECIFIED",
		1: "INVENTORY_ITEM_STATUS_AVAILABLE",
		2: "INVENTORY_ITEM_STATUS_RESERVED",
		3: "INVENTORY_ITEM_STATUS_ALLOCATED",
		4: "INVENTORY_ITEM_STATUS_BACKORDERED",
		5: "INVENTORY_ITEM_STATUS_DAMAGED",
		6: "INVENTORY_ITEM_STATUS_EXPIRED",
		7: "INVENTORY_ITEM_STATUS_QUARANTINED",
		8: "INVENTORY_ITEM_STATUS_RECALLED",
		9: "INVENTORY_ITEM_STATUS_DISPOSED",
	}
	InventoryItemStatus_value = map[string]int32{
		"INVENTORY_ITEM_STATUS_UNSPECIFIED": 0,
		"INVENTORY_ITEM_STATUS_AVAILABLE":   1,
		"INVENTORY_ITEM_STATUS_RESERVED":    2,
		"INVENTORY_ITEM_STATUS_ALLOCATED":   3,
		"INVENTORY_ITEM_STATUS_BACKORDERED": 4,
		"INVENTORY_ITEM_STATUS_DAMAGED":     5,
		"INVENTORY_ITEM_STATUS_EXPIRED":     6,
		"INVENTORY_ITEM_STATUS_QUARANTINED": 7,
		"INVENTORY_ITEM_STATUS_RECALLED":    8,
		"INVENTORY_ITEM_STATUS_DISPOSED":    9,
	}
)

func (x InventoryItemStatus) Enum() *InventoryItemStatus {
	p := new(InventoryItemStatus)
	*p = x
	return p
}

func (x InventoryItemStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InventoryItemStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_inventory_v1_enums_proto_enumTypes[0].Descriptor()
}

func (InventoryItemStatus) Type() protoreflect.EnumType {
	return &file_proto_scg_inventory_v1_enums_proto_enumTypes[0]
}

func (x InventoryItemStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InventoryItemStatus.Descriptor instead.
func (InventoryItemStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_inventory_v1_enums_proto_rawDescGZIP(), []int{0}
}

type InventoryTransactionType int32

const (
	InventoryTransactionType_INVENTORY_TRANSACTION_TYPE_UNSPECIFIED InventoryTransactionType = 0
	InventoryTransactionType_INVENTORY_TRANSACTION_TYPE_RECEIPT     InventoryTransactionType = 1
	InventoryTransactionType_INVENTORY_TRANSACTION_TYPE_SHIPMENT    InventoryTransactionType = 2
	InventoryTransactionType_INVENTORY_TRANSACTION_TYPE_TRANSFER    InventoryTransactionType = 3
	InventoryTransactionType_INVENTORY_TRANSACTION_TYPE_ADJUSTMENT  InventoryTransactionType = 4
	InventoryTransactionType_INVENTORY_TRANSACTION_TYPE_RETURN      InventoryTransactionType = 5
	InventoryTransactionType_INVENTORY_TRANSACTION_TYPE_DISPOSAL    InventoryTransactionType = 6
	InventoryTransactionType_INVENTORY_TRANSACTION_TYPE_CYCLE_COUNT InventoryTransactionType = 7
)

// Enum value maps for InventoryTransactionType.
var (
	InventoryTransactionType_name = map[int32]string{
		0: "INVENTORY_TRANSACTION_TYPE_UNSPECIFIED",
		1: "INVENTORY_TRANSACTION_TYPE_RECEIPT",
		2: "INVENTORY_TRANSACTION_TYPE_SHIPMENT",
		3: "INVENTORY_TRANSACTION_TYPE_TRANSFER",
		4: "INVENTORY_TRANSACTION_TYPE_ADJUSTMENT",
		5: "INVENTORY_TRANSACTION_TYPE_RETURN",
		6: "INVENTORY_TRANSACTION_TYPE_DISPOSAL",
		7: "INVENTORY_TRANSACTION_TYPE_CYCLE_COUNT",
	}
	InventoryTransactionType_value = map[string]int32{
		"INVENTORY_TRANSACTION_TYPE_UNSPECIFIED": 0,
		"INVENTORY_TRANSACTION_TYPE_RECEIPT":     1,
		"INVENTORY_TRANSACTION_TYPE_SHIPMENT":    2,
		"INVENTORY_TRANSACTION_TYPE_TRANSFER":    3,
		"INVENTORY_TRANSACTION_TYPE_ADJUSTMENT":  4,
		"INVENTORY_TRANSACTION_TYPE_RETURN":      5,
		"INVENTORY_TRANSACTION_TYPE_DISPOSAL":    6,
		"INVENTORY_TRANSACTION_TYPE_CYCLE_COUNT": 7,
	}
)

func (x InventoryTransactionType) Enum() *InventoryTransactionType {
	p := new(InventoryTransactionType)
	*p = x
	return p
}

func (x InventoryTransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InventoryTransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_inventory_v1_enums_proto_enumTypes[1].Descriptor()
}

func (InventoryTransactionType) Type() protoreflect.EnumType {
	return &file_proto_scg_inventory_v1_enums_proto_enumTypes[1]
}

func (x InventoryTransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InventoryTransactionType.Descriptor instead.
func (InventoryTransactionType) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_inventory_v1_enums_proto_rawDescGZIP(), []int{1}
}

type InventoryAdjustmentReason int32

const (
	InventoryAdjustmentReason_INVENTORY_ADJUSTMENT_REASON_UNSPECIFIED      InventoryAdjustmentReason = 0
	InventoryAdjustmentReason_INVENTORY_ADJUSTMENT_REASON_DAMAGE           InventoryAdjustmentReason = 1
	InventoryAdjustmentReason_INVENTORY_ADJUSTMENT_REASON_THEFT            InventoryAdjustmentReason = 2
	InventoryAdjustmentReason_INVENTORY_ADJUSTMENT_REASON_EXPIRATION       InventoryAdjustmentReason = 3
	InventoryAdjustmentReason_INVENTORY_ADJUSTMENT_REASON_QUALITY_ISSUE    InventoryAdjustmentReason = 4
	InventoryAdjustmentReason_INVENTORY_ADJUSTMENT_REASON_COUNT_CORRECTION InventoryAdjustmentReason = 5
	InventoryAdjustmentReason_INVENTORY_ADJUSTMENT_REASON_SYSTEM_ERROR     InventoryAdjustmentReason = 6
	InventoryAdjustmentReason_INVENTORY_ADJUSTMENT_REASON_OTHER            InventoryAdjustmentReason = 7
)

// Enum value maps for InventoryAdjustmentReason.
var (
	InventoryAdjustmentReason_name = map[int32]string{
		0: "INVENTORY_ADJUSTMENT_REASON_UNSPECIFIED",
		1: "INVENTORY_ADJUSTMENT_REASON_DAMAGE",
		2: "INVENTORY_ADJUSTMENT_REASON_THEFT",
		3: "INVENTORY_ADJUSTMENT_REASON_EXPIRATION",
		4: "INVENTORY_ADJUSTMENT_REASON_QUALITY_ISSUE",
		5: "INVENTORY_ADJUSTMENT_REASON_COUNT_CORRECTION",
		6: "INVENTORY_ADJUSTMENT_REASON_SYSTEM_ERROR",
		7: "INVENTORY_ADJUSTMENT_REASON_OTHER",
	}
	InventoryAdjustmentReason_value = map[string]int32{
		"INVENTORY_ADJUSTMENT_REASON_UNSPECIFIED":      0,
		"INVENTORY_ADJUSTMENT_REASON_DAMAGE":           1,
		"INVENTORY_ADJUSTMENT_REASON_THEFT":            2,
		"INVENTORY_ADJUSTMENT_REASON_EXPIRATION":       3,
		"INVENTORY_ADJUSTMENT_REASON_QUALITY_ISSUE":    4,
		"INVENTORY_ADJUSTMENT_REASON_COUNT_CORRECTION": 5,
		"INVENTORY_ADJUSTMENT_REASON_SYSTEM_ERROR":     6,
		"INVENTORY_ADJUSTMENT_REASON_OTHER":            7,
	}
)

func (x InventoryAdjustmentReason) Enum() *InventoryAdjustmentReason {
	p := new(InventoryAdjustmentReason)
	*p = x
	return p
}

func (x InventoryAdjustmentReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InventoryAdjustmentReason) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_inventory_v1_enums_proto_enumTypes[2].Descriptor()
}

func (InventoryAdjustmentReason) Type() protoreflect.EnumType {
	return &file_proto_scg_inventory_v1_enums_proto_enumTypes[2]
}

func (x InventoryAdjustmentReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InventoryAdjustmentReason.Descriptor instead.
func (InventoryAdjustmentReason) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_inventory_v1_enums_proto_rawDescGZIP(), []int{2}
}

type InventoryReservationStatus int32

const (
	InventoryReservationStatus_INVENTORY_RESERVATION_STATUS_UNSPECIFIED InventoryReservationStatus = 0
	InventoryReservationStatus_INVENTORY_RESERVATION_STATUS_PENDING     InventoryReservationStatus = 1
	InventoryReservationStatus_INVENTORY_RESERVATION_STATUS_CONFIRMED   InventoryReservationStatus = 2
	InventoryReservationStatus_INVENTORY_RESERVATION_STATUS_ALLOCATED   InventoryReservationStatus = 3
	InventoryReservationStatus_INVENTORY_RESERVATION_STATUS_FULFILLED   InventoryReservationStatus = 4
	InventoryReservationStatus_INVENTORY_RESERVATION_STATUS_CANCELLED   InventoryReservationStatus = 5
	InventoryReservationStatus_INVENTORY_RESERVATION_STATUS_EXPIRED     InventoryReservationStatus = 6
)

// Enum value maps for InventoryReservationStatus.
var (
	InventoryReservationStatus_name = map[int32]string{
		0: "INVENTORY_RESERVATION_STATUS_UNSPECIFIED",
		1: "INVENTORY_RESERVATION_STATUS_PENDING",
		2: "INVENTORY_RESERVATION_STATUS_CONFIRMED",
		3: "INVENTORY_RESERVATION_STATUS_ALLOCATED",
		4: "INVENTORY_RESERVATION_STATUS_FULFILLED",
		5: "INVENTORY_RESERVATION_STATUS_CANCELLED",
		6: "INVENTORY_RESERVATION_STATUS_EXPIRED",
	}
	InventoryReservationStatus_value = map[string]int32{
		"INVENTORY_RESERVATION_STATUS_UNSPECIFIED": 0,
		"INVENTORY_RESERVATION_STATUS_PENDING":     1,
		"INVENTORY_RESERVATION_STATUS_CONFIRMED":   2,
		"INVENTORY_RESERVATION_STATUS_ALLOCATED":   3,
		"INVENTORY_RESERVATION_STATUS_FULFILLED":   4,
		"INVENTORY_RESERVATION_STATUS_CANCELLED":   5,
		"INVENTORY_RESERVATION_STATUS_EXPIRED":     6,
	}
)

func (x InventoryReservationStatus) Enum() *InventoryReservationStatus {
	p := new(InventoryReservationStatus)
	*p = x
	return p
}

func (x InventoryReservationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InventoryReservationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_inventory_v1_enums_proto_enumTypes[3].Descriptor()
}

func (InventoryReservationStatus) Type() protoreflect.EnumType {
	return &file_proto_scg_inventory_v1_enums_proto_enumTypes[3]
}

func (x InventoryReservationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InventoryReservationStatus.Descriptor instead.
func (InventoryReservationStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_inventory_v1_enums_proto_rawDescGZIP(), []int{3}
}

var File_proto_scg_inventory_v1_enums_proto protoreflect.FileDescriptor

const file_proto_scg_inventory_v1_enums_proto_rawDesc = "" +
	"\n" +
	"\"proto/scg/inventory/v1/enums.proto\x12\x16proto.scg.inventory.v1*\x86\x03\n" +
	"\x13InventoryItemStatus\x12%\n" +
	"!INVENTORY_ITEM_STATUS_UNSPECIFIED\x10\x00\x12#\n" +
	"\x1fINVENTORY_ITEM_STATUS_AVAILABLE\x10\x01\x12\"\n" +
	"\x1eINVENTORY_ITEM_STATUS_RESERVED\x10\x02\x12#\n" +
	"\x1fINVENTORY_ITEM_STATUS_ALLOCATED\x10\x03\x12%\n" +
	"!INVENTORY_ITEM_STATUS_BACKORDERED\x10\x04\x12!\n" +
	"\x1dINVENTORY_ITEM_STATUS_DAMAGED\x10\x05\x12!\n" +
	"\x1dINVENTORY_ITEM_STATUS_EXPIRED\x10\x06\x12%\n" +
	"!INVENTORY_ITEM_STATUS_QUARANTINED\x10\a\x12\"\n" +
	"\x1eINVENTORY_ITEM_STATUS_RECALLED\x10\b\x12\"\n" +
	"\x1eINVENTORY_ITEM_STATUS_DISPOSED\x10\t*\xe7\x02\n" +
	"\x18InventoryTransactionType\x12*\n" +
	"&INVENTORY_TRANSACTION_TYPE_UNSPECIFIED\x10\x00\x12&\n" +
	"\"INVENTORY_TRANSACTION_TYPE_RECEIPT\x10\x01\x12'\n" +
	"#INVENTORY_TRANSACTION_TYPE_SHIPMENT\x10\x02\x12'\n" +
	"#INVENTORY_TRANSACTION_TYPE_TRANSFER\x10\x03\x12)\n" +
	"%INVENTORY_TRANSACTION_TYPE_ADJUSTMENT\x10\x04\x12%\n" +
	"!INVENTORY_TRANSACTION_TYPE_RETURN\x10\x05\x12'\n" +
	"#INVENTORY_TRANSACTION_TYPE_DISPOSAL\x10\x06\x12*\n" +
	"&INVENTORY_TRANSACTION_TYPE_CYCLE_COUNT\x10\a*\xf9\x02\n" +
	"\x19InventoryAdjustmentReason\x12+\n" +
	"'INVENTORY_ADJUSTMENT_REASON_UNSPECIFIED\x10\x00\x12&\n" +
	"\"INVENTORY_ADJUSTMENT_REASON_DAMAGE\x10\x01\x12%\n" +
	"!INVENTORY_ADJUSTMENT_REASON_THEFT\x10\x02\x12*\n" +
	"&INVENTORY_ADJUSTMENT_REASON_EXPIRATION\x10\x03\x12-\n" +
	")INVENTORY_ADJUSTMENT_REASON_QUALITY_ISSUE\x10\x04\x120\n" +
	",INVENTORY_ADJUSTMENT_REASON_COUNT_CORRECTION\x10\x05\x12,\n" +
	"(INVENTORY_ADJUSTMENT_REASON_SYSTEM_ERROR\x10\x06\x12%\n" +
	"!INVENTORY_ADJUSTMENT_REASON_OTHER\x10\a*\xce\x02\n" +
	"\x1aInventoryReservationStatus\x12,\n" +
	"(INVENTORY_RESERVATION_STATUS_UNSPECIFIED\x10\x00\x12(\n" +
	"$INVENTORY_RESERVATION_STATUS_PENDING\x10\x01\x12*\n" +
	"&INVENTORY_RESERVATION_STATUS_CONFIRMED\x10\x02\x12*\n" +
	"&INVENTORY_RESERVATION_STATUS_ALLOCATED\x10\x03\x12*\n" +
	"&INVENTORY_RESERVATION_STATUS_FULFILLED\x10\x04\x12*\n" +
	"&INVENTORY_RESERVATION_STATUS_CANCELLED\x10\x05\x12(\n" +
	"$INVENTORY_RESERVATION_STATUS_EXPIRED\x10\x06BNZLgithub.com/hbttundar/scg-contracts/gen/go/proto/scg/inventory/v1;inventoryv1b\x06proto3"

var (
	file_proto_scg_inventory_v1_enums_proto_rawDescOnce sync.Once
	file_proto_scg_inventory_v1_enums_proto_rawDescData []byte
)

func file_proto_scg_inventory_v1_enums_proto_rawDescGZIP() []byte {
	file_proto_scg_inventory_v1_enums_proto_rawDescOnce.Do(func() {
		file_proto_scg_inventory_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_scg_inventory_v1_enums_proto_rawDesc), len(file_proto_scg_inventory_v1_enums_proto_rawDesc)))
	})
	return file_proto_scg_inventory_v1_enums_proto_rawDescData
}

var file_proto_scg_inventory_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_proto_scg_inventory_v1_enums_proto_goTypes = []any{
	(InventoryItemStatus)(0),        // 0: proto.scg.inventory.v1.InventoryItemStatus
	(InventoryTransactionType)(0),   // 1: proto.scg.inventory.v1.InventoryTransactionType
	(InventoryAdjustmentReason)(0),  // 2: proto.scg.inventory.v1.InventoryAdjustmentReason
	(InventoryReservationStatus)(0), // 3: proto.scg.inventory.v1.InventoryReservationStatus
}
var file_proto_scg_inventory_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_scg_inventory_v1_enums_proto_init() }
func file_proto_scg_inventory_v1_enums_proto_init() {
	if File_proto_scg_inventory_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_scg_inventory_v1_enums_proto_rawDesc), len(file_proto_scg_inventory_v1_enums_proto_rawDesc)),
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_scg_inventory_v1_enums_proto_goTypes,
		DependencyIndexes: file_proto_scg_inventory_v1_enums_proto_depIdxs,
		EnumInfos:         file_proto_scg_inventory_v1_enums_proto_enumTypes,
	}.Build()
	File_proto_scg_inventory_v1_enums_proto = out.File
	file_proto_scg_inventory_v1_enums_proto_goTypes = nil
	file_proto_scg_inventory_v1_enums_proto_depIdxs = nil
}
