// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/scg/onboarding/v1/enums.proto

package onboardingv1

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

// OnboardingStatus represents the status of an onboarding process
type OnboardingStatus int32

const (
	OnboardingStatus_ONBOARDING_STATUS_UNSPECIFIED    OnboardingStatus = 0
	OnboardingStatus_ONBOARDING_STATUS_INITIATED      OnboardingStatus = 1 // Onboarding process has been initiated
	OnboardingStatus_ONBOARDING_STATUS_IN_PROGRESS    OnboardingStatus = 2 // Onboarding is in progress
	OnboardingStatus_ONBOARDING_STATUS_PENDING_REVIEW OnboardingStatus = 3 // Onboarding is pending review
	OnboardingStatus_ONBOARDING_STATUS_APPROVED       OnboardingStatus = 4 // Onboarding has been approved
	OnboardingStatus_ONBOARDING_STATUS_REJECTED       OnboardingStatus = 5 // Onboarding has been rejected
	OnboardingStatus_ONBOARDING_STATUS_COMPLETED      OnboardingStatus = 6 // Onboarding has been completed
	OnboardingStatus_ONBOARDING_STATUS_CANCELLED      OnboardingStatus = 7 // Onboarding has been cancelled
)

// Enum value maps for OnboardingStatus.
var (
	OnboardingStatus_name = map[int32]string{
		0: "ONBOARDING_STATUS_UNSPECIFIED",
		1: "ONBOARDING_STATUS_INITIATED",
		2: "ONBOARDING_STATUS_IN_PROGRESS",
		3: "ONBOARDING_STATUS_PENDING_REVIEW",
		4: "ONBOARDING_STATUS_APPROVED",
		5: "ONBOARDING_STATUS_REJECTED",
		6: "ONBOARDING_STATUS_COMPLETED",
		7: "ONBOARDING_STATUS_CANCELLED",
	}
	OnboardingStatus_value = map[string]int32{
		"ONBOARDING_STATUS_UNSPECIFIED":    0,
		"ONBOARDING_STATUS_INITIATED":      1,
		"ONBOARDING_STATUS_IN_PROGRESS":    2,
		"ONBOARDING_STATUS_PENDING_REVIEW": 3,
		"ONBOARDING_STATUS_APPROVED":       4,
		"ONBOARDING_STATUS_REJECTED":       5,
		"ONBOARDING_STATUS_COMPLETED":      6,
		"ONBOARDING_STATUS_CANCELLED":      7,
	}
)

func (x OnboardingStatus) Enum() *OnboardingStatus {
	p := new(OnboardingStatus)
	*p = x
	return p
}

func (x OnboardingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OnboardingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_onboarding_v1_enums_proto_enumTypes[0].Descriptor()
}

func (OnboardingStatus) Type() protoreflect.EnumType {
	return &file_proto_scg_onboarding_v1_enums_proto_enumTypes[0]
}

func (x OnboardingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OnboardingStatus.Descriptor instead.
func (OnboardingStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_onboarding_v1_enums_proto_rawDescGZIP(), []int{0}
}

// VerificationType represents different types of verification
type VerificationType int32

const (
	VerificationType_VERIFICATION_TYPE_UNSPECIFIED VerificationType = 0
	VerificationType_VERIFICATION_TYPE_EMAIL       VerificationType = 1 // Email verification
	VerificationType_VERIFICATION_TYPE_PHONE       VerificationType = 2 // Phone verification
	VerificationType_VERIFICATION_TYPE_DOCUMENT    VerificationType = 3 // Document verification
	VerificationType_VERIFICATION_TYPE_IDENTITY    VerificationType = 4 // Identity verification
	VerificationType_VERIFICATION_TYPE_ADDRESS     VerificationType = 5 // Address verification
	VerificationType_VERIFICATION_TYPE_BUSINESS    VerificationType = 6 // Business verification
)

// Enum value maps for VerificationType.
var (
	VerificationType_name = map[int32]string{
		0: "VERIFICATION_TYPE_UNSPECIFIED",
		1: "VERIFICATION_TYPE_EMAIL",
		2: "VERIFICATION_TYPE_PHONE",
		3: "VERIFICATION_TYPE_DOCUMENT",
		4: "VERIFICATION_TYPE_IDENTITY",
		5: "VERIFICATION_TYPE_ADDRESS",
		6: "VERIFICATION_TYPE_BUSINESS",
	}
	VerificationType_value = map[string]int32{
		"VERIFICATION_TYPE_UNSPECIFIED": 0,
		"VERIFICATION_TYPE_EMAIL":       1,
		"VERIFICATION_TYPE_PHONE":       2,
		"VERIFICATION_TYPE_DOCUMENT":    3,
		"VERIFICATION_TYPE_IDENTITY":    4,
		"VERIFICATION_TYPE_ADDRESS":     5,
		"VERIFICATION_TYPE_BUSINESS":    6,
	}
)

func (x VerificationType) Enum() *VerificationType {
	p := new(VerificationType)
	*p = x
	return p
}

func (x VerificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VerificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_onboarding_v1_enums_proto_enumTypes[1].Descriptor()
}

func (VerificationType) Type() protoreflect.EnumType {
	return &file_proto_scg_onboarding_v1_enums_proto_enumTypes[1]
}

func (x VerificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VerificationType.Descriptor instead.
func (VerificationType) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_onboarding_v1_enums_proto_rawDescGZIP(), []int{1}
}

// DocumentType represents different types of documents
type DocumentType int32

const (
	DocumentType_DOCUMENT_TYPE_UNSPECIFIED      DocumentType = 0
	DocumentType_DOCUMENT_TYPE_ID_CARD          DocumentType = 1 // ID card
	DocumentType_DOCUMENT_TYPE_PASSPORT         DocumentType = 2 // Passport
	DocumentType_DOCUMENT_TYPE_DRIVERS_LICENSE  DocumentType = 3 // Driver's license
	DocumentType_DOCUMENT_TYPE_BUSINESS_LICENSE DocumentType = 4 // Business license
	DocumentType_DOCUMENT_TYPE_TAX_ID           DocumentType = 5 // Tax ID
	DocumentType_DOCUMENT_TYPE_UTILITY_BILL     DocumentType = 6 // Utility bill
	DocumentType_DOCUMENT_TYPE_BANK_STATEMENT   DocumentType = 7 // Bank statement
	DocumentType_DOCUMENT_TYPE_CERTIFICATE      DocumentType = 8 // Certificate
	DocumentType_DOCUMENT_TYPE_OTHER            DocumentType = 9 // Other document type
)

// Enum value maps for DocumentType.
var (
	DocumentType_name = map[int32]string{
		0: "DOCUMENT_TYPE_UNSPECIFIED",
		1: "DOCUMENT_TYPE_ID_CARD",
		2: "DOCUMENT_TYPE_PASSPORT",
		3: "DOCUMENT_TYPE_DRIVERS_LICENSE",
		4: "DOCUMENT_TYPE_BUSINESS_LICENSE",
		5: "DOCUMENT_TYPE_TAX_ID",
		6: "DOCUMENT_TYPE_UTILITY_BILL",
		7: "DOCUMENT_TYPE_BANK_STATEMENT",
		8: "DOCUMENT_TYPE_CERTIFICATE",
		9: "DOCUMENT_TYPE_OTHER",
	}
	DocumentType_value = map[string]int32{
		"DOCUMENT_TYPE_UNSPECIFIED":      0,
		"DOCUMENT_TYPE_ID_CARD":          1,
		"DOCUMENT_TYPE_PASSPORT":         2,
		"DOCUMENT_TYPE_DRIVERS_LICENSE":  3,
		"DOCUMENT_TYPE_BUSINESS_LICENSE": 4,
		"DOCUMENT_TYPE_TAX_ID":           5,
		"DOCUMENT_TYPE_UTILITY_BILL":     6,
		"DOCUMENT_TYPE_BANK_STATEMENT":   7,
		"DOCUMENT_TYPE_CERTIFICATE":      8,
		"DOCUMENT_TYPE_OTHER":            9,
	}
)

func (x DocumentType) Enum() *DocumentType {
	p := new(DocumentType)
	*p = x
	return p
}

func (x DocumentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DocumentType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_onboarding_v1_enums_proto_enumTypes[2].Descriptor()
}

func (DocumentType) Type() protoreflect.EnumType {
	return &file_proto_scg_onboarding_v1_enums_proto_enumTypes[2]
}

func (x DocumentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DocumentType.Descriptor instead.
func (DocumentType) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_onboarding_v1_enums_proto_rawDescGZIP(), []int{2}
}

// DocumentStatus represents the status of a document
type DocumentStatus int32

const (
	DocumentStatus_DOCUMENT_STATUS_UNSPECIFIED     DocumentStatus = 0
	DocumentStatus_DOCUMENT_STATUS_PENDING         DocumentStatus = 1 // Document is pending review
	DocumentStatus_DOCUMENT_STATUS_APPROVED        DocumentStatus = 2 // Document has been approved
	DocumentStatus_DOCUMENT_STATUS_REJECTED        DocumentStatus = 3 // Document has been rejected
	DocumentStatus_DOCUMENT_STATUS_EXPIRED         DocumentStatus = 4 // Document has expired
	DocumentStatus_DOCUMENT_STATUS_REQUIRES_UPDATE DocumentStatus = 5 // Document requires update
)

// Enum value maps for DocumentStatus.
var (
	DocumentStatus_name = map[int32]string{
		0: "DOCUMENT_STATUS_UNSPECIFIED",
		1: "DOCUMENT_STATUS_PENDING",
		2: "DOCUMENT_STATUS_APPROVED",
		3: "DOCUMENT_STATUS_REJECTED",
		4: "DOCUMENT_STATUS_EXPIRED",
		5: "DOCUMENT_STATUS_REQUIRES_UPDATE",
	}
	DocumentStatus_value = map[string]int32{
		"DOCUMENT_STATUS_UNSPECIFIED":     0,
		"DOCUMENT_STATUS_PENDING":         1,
		"DOCUMENT_STATUS_APPROVED":        2,
		"DOCUMENT_STATUS_REJECTED":        3,
		"DOCUMENT_STATUS_EXPIRED":         4,
		"DOCUMENT_STATUS_REQUIRES_UPDATE": 5,
	}
)

func (x DocumentStatus) Enum() *DocumentStatus {
	p := new(DocumentStatus)
	*p = x
	return p
}

func (x DocumentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DocumentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_onboarding_v1_enums_proto_enumTypes[3].Descriptor()
}

func (DocumentStatus) Type() protoreflect.EnumType {
	return &file_proto_scg_onboarding_v1_enums_proto_enumTypes[3]
}

func (x DocumentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DocumentStatus.Descriptor instead.
func (DocumentStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_onboarding_v1_enums_proto_rawDescGZIP(), []int{3}
}

// OnboardingStepType represents different types of onboarding steps
type OnboardingStepType int32

const (
	OnboardingStepType_ONBOARDING_STEP_TYPE_UNSPECIFIED   OnboardingStepType = 0
	OnboardingStepType_ONBOARDING_STEP_TYPE_REGISTRATION  OnboardingStepType = 1 // Registration step
	OnboardingStepType_ONBOARDING_STEP_TYPE_VERIFICATION  OnboardingStepType = 2 // Verification step
	OnboardingStepType_ONBOARDING_STEP_TYPE_DOCUMENTATION OnboardingStepType = 3 // Documentation step
	OnboardingStepType_ONBOARDING_STEP_TYPE_TRAINING      OnboardingStepType = 4 // Training step
	OnboardingStepType_ONBOARDING_STEP_TYPE_INTEGRATION   OnboardingStepType = 5 // Integration step
	OnboardingStepType_ONBOARDING_STEP_TYPE_CONFIGURATION OnboardingStepType = 6 // Configuration step
	OnboardingStepType_ONBOARDING_STEP_TYPE_APPROVAL      OnboardingStepType = 7 // Approval step
	OnboardingStepType_ONBOARDING_STEP_TYPE_ACTIVATION    OnboardingStepType = 8 // Activation step
)

// Enum value maps for OnboardingStepType.
var (
	OnboardingStepType_name = map[int32]string{
		0: "ONBOARDING_STEP_TYPE_UNSPECIFIED",
		1: "ONBOARDING_STEP_TYPE_REGISTRATION",
		2: "ONBOARDING_STEP_TYPE_VERIFICATION",
		3: "ONBOARDING_STEP_TYPE_DOCUMENTATION",
		4: "ONBOARDING_STEP_TYPE_TRAINING",
		5: "ONBOARDING_STEP_TYPE_INTEGRATION",
		6: "ONBOARDING_STEP_TYPE_CONFIGURATION",
		7: "ONBOARDING_STEP_TYPE_APPROVAL",
		8: "ONBOARDING_STEP_TYPE_ACTIVATION",
	}
	OnboardingStepType_value = map[string]int32{
		"ONBOARDING_STEP_TYPE_UNSPECIFIED":   0,
		"ONBOARDING_STEP_TYPE_REGISTRATION":  1,
		"ONBOARDING_STEP_TYPE_VERIFICATION":  2,
		"ONBOARDING_STEP_TYPE_DOCUMENTATION": 3,
		"ONBOARDING_STEP_TYPE_TRAINING":      4,
		"ONBOARDING_STEP_TYPE_INTEGRATION":   5,
		"ONBOARDING_STEP_TYPE_CONFIGURATION": 6,
		"ONBOARDING_STEP_TYPE_APPROVAL":      7,
		"ONBOARDING_STEP_TYPE_ACTIVATION":    8,
	}
)

func (x OnboardingStepType) Enum() *OnboardingStepType {
	p := new(OnboardingStepType)
	*p = x
	return p
}

func (x OnboardingStepType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OnboardingStepType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_onboarding_v1_enums_proto_enumTypes[4].Descriptor()
}

func (OnboardingStepType) Type() protoreflect.EnumType {
	return &file_proto_scg_onboarding_v1_enums_proto_enumTypes[4]
}

func (x OnboardingStepType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OnboardingStepType.Descriptor instead.
func (OnboardingStepType) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_onboarding_v1_enums_proto_rawDescGZIP(), []int{4}
}

// OnboardingStepStatus represents the status of an onboarding step
type OnboardingStepStatus int32

const (
	OnboardingStepStatus_ONBOARDING_STEP_STATUS_UNSPECIFIED OnboardingStepStatus = 0
	OnboardingStepStatus_ONBOARDING_STEP_STATUS_NOT_STARTED OnboardingStepStatus = 1 // Step has not been started
	OnboardingStepStatus_ONBOARDING_STEP_STATUS_IN_PROGRESS OnboardingStepStatus = 2 // Step is in progress
	OnboardingStepStatus_ONBOARDING_STEP_STATUS_COMPLETED   OnboardingStepStatus = 3 // Step has been completed
	OnboardingStepStatus_ONBOARDING_STEP_STATUS_SKIPPED     OnboardingStepStatus = 4 // Step has been skipped
	OnboardingStepStatus_ONBOARDING_STEP_STATUS_FAILED      OnboardingStepStatus = 5 // Step has failed
	OnboardingStepStatus_ONBOARDING_STEP_STATUS_PENDING     OnboardingStepStatus = 6 // Step is pending
)

// Enum value maps for OnboardingStepStatus.
var (
	OnboardingStepStatus_name = map[int32]string{
		0: "ONBOARDING_STEP_STATUS_UNSPECIFIED",
		1: "ONBOARDING_STEP_STATUS_NOT_STARTED",
		2: "ONBOARDING_STEP_STATUS_IN_PROGRESS",
		3: "ONBOARDING_STEP_STATUS_COMPLETED",
		4: "ONBOARDING_STEP_STATUS_SKIPPED",
		5: "ONBOARDING_STEP_STATUS_FAILED",
		6: "ONBOARDING_STEP_STATUS_PENDING",
	}
	OnboardingStepStatus_value = map[string]int32{
		"ONBOARDING_STEP_STATUS_UNSPECIFIED": 0,
		"ONBOARDING_STEP_STATUS_NOT_STARTED": 1,
		"ONBOARDING_STEP_STATUS_IN_PROGRESS": 2,
		"ONBOARDING_STEP_STATUS_COMPLETED":   3,
		"ONBOARDING_STEP_STATUS_SKIPPED":     4,
		"ONBOARDING_STEP_STATUS_FAILED":      5,
		"ONBOARDING_STEP_STATUS_PENDING":     6,
	}
)

func (x OnboardingStepStatus) Enum() *OnboardingStepStatus {
	p := new(OnboardingStepStatus)
	*p = x
	return p
}

func (x OnboardingStepStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OnboardingStepStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_onboarding_v1_enums_proto_enumTypes[5].Descriptor()
}

func (OnboardingStepStatus) Type() protoreflect.EnumType {
	return &file_proto_scg_onboarding_v1_enums_proto_enumTypes[5]
}

func (x OnboardingStepStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OnboardingStepStatus.Descriptor instead.
func (OnboardingStepStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_onboarding_v1_enums_proto_rawDescGZIP(), []int{5}
}

var File_proto_scg_onboarding_v1_enums_proto protoreflect.FileDescriptor

const file_proto_scg_onboarding_v1_enums_proto_rawDesc = "" +
	"\n" +
	"#proto/scg/onboarding/v1/enums.proto\x12\x17proto.scg.onboarding.v1*\xa1\x02\n" +
	"\x10OnboardingStatus\x12!\n" +
	"\x1dONBOARDING_STATUS_UNSPECIFIED\x10\x00\x12\x1f\n" +
	"\x1bONBOARDING_STATUS_INITIATED\x10\x01\x12!\n" +
	"\x1dONBOARDING_STATUS_IN_PROGRESS\x10\x02\x12$\n" +
	" ONBOARDING_STATUS_PENDING_REVIEW\x10\x03\x12\x1e\n" +
	"\x1aONBOARDING_STATUS_APPROVED\x10\x04\x12\x1e\n" +
	"\x1aONBOARDING_STATUS_REJECTED\x10\x05\x12\x1f\n" +
	"\x1bONBOARDING_STATUS_COMPLETED\x10\x06\x12\x1f\n" +
	"\x1bONBOARDING_STATUS_CANCELLED\x10\a*\xee\x01\n" +
	"\x10VerificationType\x12!\n" +
	"\x1dVERIFICATION_TYPE_UNSPECIFIED\x10\x00\x12\x1b\n" +
	"\x17VERIFICATION_TYPE_EMAIL\x10\x01\x12\x1b\n" +
	"\x17VERIFICATION_TYPE_PHONE\x10\x02\x12\x1e\n" +
	"\x1aVERIFICATION_TYPE_DOCUMENT\x10\x03\x12\x1e\n" +
	"\x1aVERIFICATION_TYPE_IDENTITY\x10\x04\x12\x1d\n" +
	"\x19VERIFICATION_TYPE_ADDRESS\x10\x05\x12\x1e\n" +
	"\x1aVERIFICATION_TYPE_BUSINESS\x10\x06*\xbf\x02\n" +
	"\fDocumentType\x12\x1d\n" +
	"\x19DOCUMENT_TYPE_UNSPECIFIED\x10\x00\x12\x19\n" +
	"\x15DOCUMENT_TYPE_ID_CARD\x10\x01\x12\x1a\n" +
	"\x16DOCUMENT_TYPE_PASSPORT\x10\x02\x12!\n" +
	"\x1dDOCUMENT_TYPE_DRIVERS_LICENSE\x10\x03\x12\"\n" +
	"\x1eDOCUMENT_TYPE_BUSINESS_LICENSE\x10\x04\x12\x18\n" +
	"\x14DOCUMENT_TYPE_TAX_ID\x10\x05\x12\x1e\n" +
	"\x1aDOCUMENT_TYPE_UTILITY_BILL\x10\x06\x12 \n" +
	"\x1cDOCUMENT_TYPE_BANK_STATEMENT\x10\a\x12\x1d\n" +
	"\x19DOCUMENT_TYPE_CERTIFICATE\x10\b\x12\x17\n" +
	"\x13DOCUMENT_TYPE_OTHER\x10\t*\xcc\x01\n" +
	"\x0eDocumentStatus\x12\x1f\n" +
	"\x1bDOCUMENT_STATUS_UNSPECIFIED\x10\x00\x12\x1b\n" +
	"\x17DOCUMENT_STATUS_PENDING\x10\x01\x12\x1c\n" +
	"\x18DOCUMENT_STATUS_APPROVED\x10\x02\x12\x1c\n" +
	"\x18DOCUMENT_STATUS_REJECTED\x10\x03\x12\x1b\n" +
	"\x17DOCUMENT_STATUS_EXPIRED\x10\x04\x12#\n" +
	"\x1fDOCUMENT_STATUS_REQUIRES_UPDATE\x10\x05*\xe9\x02\n" +
	"\x12OnboardingStepType\x12$\n" +
	" ONBOARDING_STEP_TYPE_UNSPECIFIED\x10\x00\x12%\n" +
	"!ONBOARDING_STEP_TYPE_REGISTRATION\x10\x01\x12%\n" +
	"!ONBOARDING_STEP_TYPE_VERIFICATION\x10\x02\x12&\n" +
	"\"ONBOARDING_STEP_TYPE_DOCUMENTATION\x10\x03\x12!\n" +
	"\x1dONBOARDING_STEP_TYPE_TRAINING\x10\x04\x12$\n" +
	" ONBOARDING_STEP_TYPE_INTEGRATION\x10\x05\x12&\n" +
	"\"ONBOARDING_STEP_TYPE_CONFIGURATION\x10\x06\x12!\n" +
	"\x1dONBOARDING_STEP_TYPE_APPROVAL\x10\a\x12#\n" +
	"\x1fONBOARDING_STEP_TYPE_ACTIVATION\x10\b*\x9f\x02\n" +
	"\x14OnboardingStepStatus\x12&\n" +
	"\"ONBOARDING_STEP_STATUS_UNSPECIFIED\x10\x00\x12&\n" +
	"\"ONBOARDING_STEP_STATUS_NOT_STARTED\x10\x01\x12&\n" +
	"\"ONBOARDING_STEP_STATUS_IN_PROGRESS\x10\x02\x12$\n" +
	" ONBOARDING_STEP_STATUS_COMPLETED\x10\x03\x12\"\n" +
	"\x1eONBOARDING_STEP_STATUS_SKIPPED\x10\x04\x12!\n" +
	"\x1dONBOARDING_STEP_STATUS_FAILED\x10\x05\x12\"\n" +
	"\x1eONBOARDING_STEP_STATUS_PENDING\x10\x06BPZNgithub.com/hbttundar/scg-contracts/gen/go/proto/scg/onboarding/v1;onboardingv1b\x06proto3"

var (
	file_proto_scg_onboarding_v1_enums_proto_rawDescOnce sync.Once
	file_proto_scg_onboarding_v1_enums_proto_rawDescData []byte
)

func file_proto_scg_onboarding_v1_enums_proto_rawDescGZIP() []byte {
	file_proto_scg_onboarding_v1_enums_proto_rawDescOnce.Do(func() {
		file_proto_scg_onboarding_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_scg_onboarding_v1_enums_proto_rawDesc), len(file_proto_scg_onboarding_v1_enums_proto_rawDesc)))
	})
	return file_proto_scg_onboarding_v1_enums_proto_rawDescData
}

var file_proto_scg_onboarding_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_proto_scg_onboarding_v1_enums_proto_goTypes = []any{
	(OnboardingStatus)(0),     // 0: proto.scg.onboarding.v1.OnboardingStatus
	(VerificationType)(0),     // 1: proto.scg.onboarding.v1.VerificationType
	(DocumentType)(0),         // 2: proto.scg.onboarding.v1.DocumentType
	(DocumentStatus)(0),       // 3: proto.scg.onboarding.v1.DocumentStatus
	(OnboardingStepType)(0),   // 4: proto.scg.onboarding.v1.OnboardingStepType
	(OnboardingStepStatus)(0), // 5: proto.scg.onboarding.v1.OnboardingStepStatus
}
var file_proto_scg_onboarding_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_scg_onboarding_v1_enums_proto_init() }
func file_proto_scg_onboarding_v1_enums_proto_init() {
	if File_proto_scg_onboarding_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_scg_onboarding_v1_enums_proto_rawDesc), len(file_proto_scg_onboarding_v1_enums_proto_rawDesc)),
			NumEnums:      6,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_scg_onboarding_v1_enums_proto_goTypes,
		DependencyIndexes: file_proto_scg_onboarding_v1_enums_proto_depIdxs,
		EnumInfos:         file_proto_scg_onboarding_v1_enums_proto_enumTypes,
	}.Build()
	File_proto_scg_onboarding_v1_enums_proto = out.File
	file_proto_scg_onboarding_v1_enums_proto_goTypes = nil
	file_proto_scg_onboarding_v1_enums_proto_depIdxs = nil
}
