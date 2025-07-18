// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/scg/shared/v1/enums.proto

package sharedv1

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

type CurrencyCode int32

const (
	CurrencyCode_CURRENCY_CODE_UNSPECIFIED CurrencyCode = 0
	CurrencyCode_CURRENCY_CODE_EUR         CurrencyCode = 1
	CurrencyCode_CURRENCY_CODE_USD         CurrencyCode = 2
)

// Enum value maps for CurrencyCode.
var (
	CurrencyCode_name = map[int32]string{
		0: "CURRENCY_CODE_UNSPECIFIED",
		1: "CURRENCY_CODE_EUR",
		2: "CURRENCY_CODE_USD",
	}
	CurrencyCode_value = map[string]int32{
		"CURRENCY_CODE_UNSPECIFIED": 0,
		"CURRENCY_CODE_EUR":         1,
		"CURRENCY_CODE_USD":         2,
	}
)

func (x CurrencyCode) Enum() *CurrencyCode {
	p := new(CurrencyCode)
	*p = x
	return p
}

func (x CurrencyCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CurrencyCode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_shared_v1_enums_proto_enumTypes[0].Descriptor()
}

func (CurrencyCode) Type() protoreflect.EnumType {
	return &file_proto_scg_shared_v1_enums_proto_enumTypes[0]
}

func (x CurrencyCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CurrencyCode.Descriptor instead.
func (CurrencyCode) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_shared_v1_enums_proto_rawDescGZIP(), []int{0}
}

var File_proto_scg_shared_v1_enums_proto protoreflect.FileDescriptor

const file_proto_scg_shared_v1_enums_proto_rawDesc = "" +
	"\n" +
	"\x1fproto/scg/shared/v1/enums.proto\x12\x13proto.scg.shared.v1*[\n" +
	"\fCurrencyCode\x12\x1d\n" +
	"\x19CURRENCY_CODE_UNSPECIFIED\x10\x00\x12\x15\n" +
	"\x11CURRENCY_CODE_EUR\x10\x01\x12\x15\n" +
	"\x11CURRENCY_CODE_USD\x10\x02BHZFgithub.com/hbttundar/scg-contracts/gen/go/proto/scg/shared/v1;sharedv1b\x06proto3"

var (
	file_proto_scg_shared_v1_enums_proto_rawDescOnce sync.Once
	file_proto_scg_shared_v1_enums_proto_rawDescData []byte
)

func file_proto_scg_shared_v1_enums_proto_rawDescGZIP() []byte {
	file_proto_scg_shared_v1_enums_proto_rawDescOnce.Do(func() {
		file_proto_scg_shared_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_scg_shared_v1_enums_proto_rawDesc), len(file_proto_scg_shared_v1_enums_proto_rawDesc)))
	})
	return file_proto_scg_shared_v1_enums_proto_rawDescData
}

var file_proto_scg_shared_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_scg_shared_v1_enums_proto_goTypes = []any{
	(CurrencyCode)(0), // 0: proto.scg.shared.v1.CurrencyCode
}
var file_proto_scg_shared_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_scg_shared_v1_enums_proto_init() }
func file_proto_scg_shared_v1_enums_proto_init() {
	if File_proto_scg_shared_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_scg_shared_v1_enums_proto_rawDesc), len(file_proto_scg_shared_v1_enums_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_scg_shared_v1_enums_proto_goTypes,
		DependencyIndexes: file_proto_scg_shared_v1_enums_proto_depIdxs,
		EnumInfos:         file_proto_scg_shared_v1_enums_proto_enumTypes,
	}.Build()
	File_proto_scg_shared_v1_enums_proto = out.File
	file_proto_scg_shared_v1_enums_proto_goTypes = nil
	file_proto_scg_shared_v1_enums_proto_depIdxs = nil
}
