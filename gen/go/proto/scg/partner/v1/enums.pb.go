// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/scg/partner/v1/enums.proto

package partnerv1

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

type OrgType int32

const (
	OrgType_ORG_TYPE_UNSPECIFIED         OrgType = 0
	OrgType_ORG_TYPE_SUPPLIER            OrgType = 1
	OrgType_ORG_TYPE_BUYER               OrgType = 2
	OrgType_ORG_TYPE_CARRIER             OrgType = 3
	OrgType_ORG_TYPE_AUDITOR             OrgType = 4
	OrgType_ORG_TYPE_WAREHOUSE           OrgType = 5
	OrgType_ORG_TYPE_CUSTOMS_BROKER      OrgType = 6
	OrgType_ORG_TYPE_PORT_OPERATOR       OrgType = 7
	OrgType_ORG_TYPE_INSPECTION_AGENCY   OrgType = 8
	OrgType_ORG_TYPE_TECHNOLOGY_PROVIDER OrgType = 9
	OrgType_ORG_TYPE_FINANCIER           OrgType = 10
	OrgType_ORG_TYPE_RETAILER            OrgType = 11
	OrgType_ORG_TYPE_MANUFACTURER        OrgType = 12
	OrgType_ORG_TYPE_REGULATOR           OrgType = 13
)

// Enum value maps for OrgType.
var (
	OrgType_name = map[int32]string{
		0:  "ORG_TYPE_UNSPECIFIED",
		1:  "ORG_TYPE_SUPPLIER",
		2:  "ORG_TYPE_BUYER",
		3:  "ORG_TYPE_CARRIER",
		4:  "ORG_TYPE_AUDITOR",
		5:  "ORG_TYPE_WAREHOUSE",
		6:  "ORG_TYPE_CUSTOMS_BROKER",
		7:  "ORG_TYPE_PORT_OPERATOR",
		8:  "ORG_TYPE_INSPECTION_AGENCY",
		9:  "ORG_TYPE_TECHNOLOGY_PROVIDER",
		10: "ORG_TYPE_FINANCIER",
		11: "ORG_TYPE_RETAILER",
		12: "ORG_TYPE_MANUFACTURER",
		13: "ORG_TYPE_REGULATOR",
	}
	OrgType_value = map[string]int32{
		"ORG_TYPE_UNSPECIFIED":         0,
		"ORG_TYPE_SUPPLIER":            1,
		"ORG_TYPE_BUYER":               2,
		"ORG_TYPE_CARRIER":             3,
		"ORG_TYPE_AUDITOR":             4,
		"ORG_TYPE_WAREHOUSE":           5,
		"ORG_TYPE_CUSTOMS_BROKER":      6,
		"ORG_TYPE_PORT_OPERATOR":       7,
		"ORG_TYPE_INSPECTION_AGENCY":   8,
		"ORG_TYPE_TECHNOLOGY_PROVIDER": 9,
		"ORG_TYPE_FINANCIER":           10,
		"ORG_TYPE_RETAILER":            11,
		"ORG_TYPE_MANUFACTURER":        12,
		"ORG_TYPE_REGULATOR":           13,
	}
)

func (x OrgType) Enum() *OrgType {
	p := new(OrgType)
	*p = x
	return p
}

func (x OrgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrgType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_partner_v1_enums_proto_enumTypes[0].Descriptor()
}

func (OrgType) Type() protoreflect.EnumType {
	return &file_proto_scg_partner_v1_enums_proto_enumTypes[0]
}

func (x OrgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrgType.Descriptor instead.
func (OrgType) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_partner_v1_enums_proto_rawDescGZIP(), []int{0}
}

type PartnerStatus int32

const (
	PartnerStatus_PARTNER_STATUS_UNSPECIFIED PartnerStatus = 0
	PartnerStatus_PARTNER_STATUS_ACTIVE      PartnerStatus = 1
	PartnerStatus_PARTNER_STATUS_SUSPENDED   PartnerStatus = 2
	PartnerStatus_PARTNER_STATUS_BLACKLISTED PartnerStatus = 3
	PartnerStatus_PARTNER_STATUS_OBSOLETE    PartnerStatus = 4
)

// Enum value maps for PartnerStatus.
var (
	PartnerStatus_name = map[int32]string{
		0: "PARTNER_STATUS_UNSPECIFIED",
		1: "PARTNER_STATUS_ACTIVE",
		2: "PARTNER_STATUS_SUSPENDED",
		3: "PARTNER_STATUS_BLACKLISTED",
		4: "PARTNER_STATUS_OBSOLETE",
	}
	PartnerStatus_value = map[string]int32{
		"PARTNER_STATUS_UNSPECIFIED": 0,
		"PARTNER_STATUS_ACTIVE":      1,
		"PARTNER_STATUS_SUSPENDED":   2,
		"PARTNER_STATUS_BLACKLISTED": 3,
		"PARTNER_STATUS_OBSOLETE":    4,
	}
)

func (x PartnerStatus) Enum() *PartnerStatus {
	p := new(PartnerStatus)
	*p = x
	return p
}

func (x PartnerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PartnerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_partner_v1_enums_proto_enumTypes[1].Descriptor()
}

func (PartnerStatus) Type() protoreflect.EnumType {
	return &file_proto_scg_partner_v1_enums_proto_enumTypes[1]
}

func (x PartnerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PartnerStatus.Descriptor instead.
func (PartnerStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_partner_v1_enums_proto_rawDescGZIP(), []int{1}
}

var File_proto_scg_partner_v1_enums_proto protoreflect.FileDescriptor

const file_proto_scg_partner_v1_enums_proto_rawDesc = "" +
	"\n" +
	" proto/scg/partner/v1/enums.proto\x12\x14proto.scg.partner.v1*\xef\x02\n" +
	"\aOrgType\x12\x18\n" +
	"\x14ORG_TYPE_UNSPECIFIED\x10\x00\x12\x15\n" +
	"\x11ORG_TYPE_SUPPLIER\x10\x01\x12\x12\n" +
	"\x0eORG_TYPE_BUYER\x10\x02\x12\x14\n" +
	"\x10ORG_TYPE_CARRIER\x10\x03\x12\x14\n" +
	"\x10ORG_TYPE_AUDITOR\x10\x04\x12\x16\n" +
	"\x12ORG_TYPE_WAREHOUSE\x10\x05\x12\x1b\n" +
	"\x17ORG_TYPE_CUSTOMS_BROKER\x10\x06\x12\x1a\n" +
	"\x16ORG_TYPE_PORT_OPERATOR\x10\a\x12\x1e\n" +
	"\x1aORG_TYPE_INSPECTION_AGENCY\x10\b\x12 \n" +
	"\x1cORG_TYPE_TECHNOLOGY_PROVIDER\x10\t\x12\x16\n" +
	"\x12ORG_TYPE_FINANCIER\x10\n" +
	"\x12\x15\n" +
	"\x11ORG_TYPE_RETAILER\x10\v\x12\x19\n" +
	"\x15ORG_TYPE_MANUFACTURER\x10\f\x12\x16\n" +
	"\x12ORG_TYPE_REGULATOR\x10\r*\xa5\x01\n" +
	"\rPartnerStatus\x12\x1e\n" +
	"\x1aPARTNER_STATUS_UNSPECIFIED\x10\x00\x12\x19\n" +
	"\x15PARTNER_STATUS_ACTIVE\x10\x01\x12\x1c\n" +
	"\x18PARTNER_STATUS_SUSPENDED\x10\x02\x12\x1e\n" +
	"\x1aPARTNER_STATUS_BLACKLISTED\x10\x03\x12\x1b\n" +
	"\x17PARTNER_STATUS_OBSOLETE\x10\x04BJZHgithub.com/hbttundar/scg-contracts/gen/go/proto/scg/partner/v1;partnerv1b\x06proto3"

var (
	file_proto_scg_partner_v1_enums_proto_rawDescOnce sync.Once
	file_proto_scg_partner_v1_enums_proto_rawDescData []byte
)

func file_proto_scg_partner_v1_enums_proto_rawDescGZIP() []byte {
	file_proto_scg_partner_v1_enums_proto_rawDescOnce.Do(func() {
		file_proto_scg_partner_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_scg_partner_v1_enums_proto_rawDesc), len(file_proto_scg_partner_v1_enums_proto_rawDesc)))
	})
	return file_proto_scg_partner_v1_enums_proto_rawDescData
}

var file_proto_scg_partner_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_scg_partner_v1_enums_proto_goTypes = []any{
	(OrgType)(0),       // 0: proto.scg.partner.v1.OrgType
	(PartnerStatus)(0), // 1: proto.scg.partner.v1.PartnerStatus
}
var file_proto_scg_partner_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_scg_partner_v1_enums_proto_init() }
func file_proto_scg_partner_v1_enums_proto_init() {
	if File_proto_scg_partner_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_scg_partner_v1_enums_proto_rawDesc), len(file_proto_scg_partner_v1_enums_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_scg_partner_v1_enums_proto_goTypes,
		DependencyIndexes: file_proto_scg_partner_v1_enums_proto_depIdxs,
		EnumInfos:         file_proto_scg_partner_v1_enums_proto_enumTypes,
	}.Build()
	File_proto_scg_partner_v1_enums_proto = out.File
	file_proto_scg_partner_v1_enums_proto_goTypes = nil
	file_proto_scg_partner_v1_enums_proto_depIdxs = nil
}
