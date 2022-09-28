// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: api/proto/src/transaction.proto

package walletApi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"type"
	Type TransactionType `protobuf:"varint,1,opt,name=Type,proto3,enum=wallet_api.TransactionType" json:"type"`
	// @inject_tag: json:"status"
	Status TransactionStatus `protobuf:"varint,2,opt,name=Status,proto3,enum=wallet_api.TransactionStatus" json:"status"`
	// @inject_tag: json:"amount"
	Amount int64 `protobuf:"varint,3,opt,name=Amount,proto3" json:"amount"`
	// @inject_tag: json:"wallet_id"
	WalletID string `protobuf:"bytes,4,opt,name=WalletID,proto3" json:"wallet_id"`
	// @inject_tag: json:"description"
	Description string `protobuf:"bytes,5,opt,name=Description,proto3" json:"description"`
	// @inject_tag: json:"ref_id"
	RefID string `protobuf:"bytes,6,opt,name=RefID,proto3" json:"ref_id"`
}

func (x *TransactionCreateReq) Reset() {
	*x = TransactionCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_src_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionCreateReq) ProtoMessage() {}

func (x *TransactionCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_src_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionCreateReq.ProtoReflect.Descriptor instead.
func (*TransactionCreateReq) Descriptor() ([]byte, []int) {
	return file_api_proto_src_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionCreateReq) GetType() TransactionType {
	if x != nil {
		return x.Type
	}
	return Transaction_Deposit
}

func (x *TransactionCreateReq) GetStatus() TransactionStatus {
	if x != nil {
		return x.Status
	}
	return Transaction_Success
}

func (x *TransactionCreateReq) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransactionCreateReq) GetWalletID() string {
	if x != nil {
		return x.WalletID
	}
	return ""
}

func (x *TransactionCreateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TransactionCreateReq) GetRefID() string {
	if x != nil {
		return x.RefID
	}
	return ""
}

type TransactionNewStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"new_status"
	NewStatus TransactionStatus `protobuf:"varint,2,opt,name=NewStatus,proto3,enum=wallet_api.TransactionStatus" json:"new_status"`
}

func (x *TransactionNewStatusReq) Reset() {
	*x = TransactionNewStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_src_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionNewStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionNewStatusReq) ProtoMessage() {}

func (x *TransactionNewStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_src_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionNewStatusReq.ProtoReflect.Descriptor instead.
func (*TransactionNewStatusReq) Descriptor() ([]byte, []int) {
	return file_api_proto_src_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionNewStatusReq) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *TransactionNewStatusReq) GetNewStatus() TransactionStatus {
	if x != nil {
		return x.NewStatus
	}
	return Transaction_Success
}

type TransactionListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"wallet_id"
	WalletID string `protobuf:"bytes,1,opt,name=WalletID,proto3" json:"wallet_id"`
}

func (x *TransactionListReq) Reset() {
	*x = TransactionListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_src_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionListReq) ProtoMessage() {}

func (x *TransactionListReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_src_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionListReq.ProtoReflect.Descriptor instead.
func (*TransactionListReq) Descriptor() ([]byte, []int) {
	return file_api_proto_src_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionListReq) GetWalletID() string {
	if x != nil {
		return x.WalletID
	}
	return ""
}

type Transactions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"transactions"
	Transactions []*Transaction `protobuf:"bytes,1,rep,name=Transactions,proto3" json:"transactions"`
}

func (x *Transactions) Reset() {
	*x = Transactions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_src_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transactions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transactions) ProtoMessage() {}

func (x *Transactions) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_src_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transactions.ProtoReflect.Descriptor instead.
func (*Transactions) Descriptor() ([]byte, []int) {
	return file_api_proto_src_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *Transactions) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

var File_api_proto_src_transaction_proto protoreflect.FileDescriptor

var file_api_proto_src_transaction_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1a, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x14, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x30, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1c, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x65, 0x66, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x52, 0x65, 0x66, 0x49, 0x44, 0x22, 0x67, 0x0a, 0x17, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x3c, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x30, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x49, 0x44, 0x22, 0x4b, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x3b, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x32, 0xe9, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x20, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0c,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x17, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1e, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x18, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x2d, 0x76, 0x61, 0x72,
	0x6d, 0x61, 0x7a, 0x79, 0x61, 0x72, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_src_transaction_proto_rawDescOnce sync.Once
	file_api_proto_src_transaction_proto_rawDescData = file_api_proto_src_transaction_proto_rawDesc
)

func file_api_proto_src_transaction_proto_rawDescGZIP() []byte {
	file_api_proto_src_transaction_proto_rawDescOnce.Do(func() {
		file_api_proto_src_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_src_transaction_proto_rawDescData)
	})
	return file_api_proto_src_transaction_proto_rawDescData
}

var file_api_proto_src_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_proto_src_transaction_proto_goTypes = []interface{}{
	(*TransactionCreateReq)(nil),    // 0: wallet_api.TransactionCreateReq
	(*TransactionNewStatusReq)(nil), // 1: wallet_api.TransactionNewStatusReq
	(*TransactionListReq)(nil),      // 2: wallet_api.TransactionListReq
	(*Transactions)(nil),            // 3: wallet_api.Transactions
	(TransactionType)(0),            // 4: wallet_api.Transaction.type
	(TransactionStatus)(0),          // 5: wallet_api.Transaction.status
	(*Transaction)(nil),             // 6: wallet_api.Transaction
}
var file_api_proto_src_transaction_proto_depIdxs = []int32{
	4, // 0: wallet_api.TransactionCreateReq.Type:type_name -> wallet_api.Transaction.type
	5, // 1: wallet_api.TransactionCreateReq.Status:type_name -> wallet_api.Transaction.status
	5, // 2: wallet_api.TransactionNewStatusReq.NewStatus:type_name -> wallet_api.Transaction.status
	6, // 3: wallet_api.Transactions.Transactions:type_name -> wallet_api.Transaction
	0, // 4: wallet_api.TransactionService.Create:input_type -> wallet_api.TransactionCreateReq
	1, // 5: wallet_api.TransactionService.ChangeStatus:input_type -> wallet_api.TransactionNewStatusReq
	2, // 6: wallet_api.TransactionService.List:input_type -> wallet_api.TransactionListReq
	6, // 7: wallet_api.TransactionService.Create:output_type -> wallet_api.Transaction
	6, // 8: wallet_api.TransactionService.ChangeStatus:output_type -> wallet_api.Transaction
	3, // 9: wallet_api.TransactionService.List:output_type -> wallet_api.Transactions
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_src_transaction_proto_init() }
func file_api_proto_src_transaction_proto_init() {
	if File_api_proto_src_transaction_proto != nil {
		return
	}
	file_api_proto_src_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_src_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionCreateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_src_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionNewStatusReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_src_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_src_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transactions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_src_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_src_transaction_proto_goTypes,
		DependencyIndexes: file_api_proto_src_transaction_proto_depIdxs,
		MessageInfos:      file_api_proto_src_transaction_proto_msgTypes,
	}.Build()
	File_api_proto_src_transaction_proto = out.File
	file_api_proto_src_transaction_proto_rawDesc = nil
	file_api_proto_src_transaction_proto_goTypes = nil
	file_api_proto_src_transaction_proto_depIdxs = nil
}