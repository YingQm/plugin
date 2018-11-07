// Code generated by protoc-gen-go. DO NOT EDIT.
// source: unfreeze.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	unfreeze.proto

It has these top-level messages:
	Unfreeze
	FixAmount
	LeftProportion
	UnfreezeAction
	UnfreezeCreate
	UnfreezeWithdraw
	UnfreezeTerminate
	ReceiptUnfreeze
	QueryUnfreezeWithdraw
	ReplyQueryUnfreezeWithdraw
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Unfreeze struct {
	// 解冻交易ID（唯一识别码）
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
	// 开始时间
	StartTime int64 `protobuf:"varint,2,opt,name=startTime" json:"startTime,omitempty"`
	// 币种
	AssetExec   string `protobuf:"bytes,3,opt,name=assetExec" json:"assetExec,omitempty"`
	AssetSymbol string `protobuf:"bytes,4,opt,name=assetSymbol" json:"assetSymbol,omitempty"`
	// 冻结总额
	TotalCount int64 `protobuf:"varint,5,opt,name=totalCount" json:"totalCount,omitempty"`
	// 发币人地址
	Initiator string `protobuf:"bytes,6,opt,name=initiator" json:"initiator,omitempty"`
	// 收币人地址
	Beneficiary string `protobuf:"bytes,7,opt,name=beneficiary" json:"beneficiary,omitempty"`
	// 解冻剩余币数
	Remaining int64 `protobuf:"varint,8,opt,name=remaining" json:"remaining,omitempty"`
	// 解冻方式（百分比；固额）
	Means string `protobuf:"bytes,9,opt,name=means" json:"means,omitempty"`
	// Types that are valid to be assigned to MeansOpt:
	//	*Unfreeze_FixAmount
	//	*Unfreeze_LeftProportion
	MeansOpt isUnfreeze_MeansOpt `protobuf_oneof:"meansOpt"`
}

func (m *Unfreeze) Reset()                    { *m = Unfreeze{} }
func (m *Unfreeze) String() string            { return proto.CompactTextString(m) }
func (*Unfreeze) ProtoMessage()               {}
func (*Unfreeze) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isUnfreeze_MeansOpt interface {
	isUnfreeze_MeansOpt()
}

type Unfreeze_FixAmount struct {
	FixAmount *FixAmount `protobuf:"bytes,10,opt,name=fixAmount,oneof"`
}
type Unfreeze_LeftProportion struct {
	LeftProportion *LeftProportion `protobuf:"bytes,11,opt,name=leftProportion,oneof"`
}

func (*Unfreeze_FixAmount) isUnfreeze_MeansOpt()      {}
func (*Unfreeze_LeftProportion) isUnfreeze_MeansOpt() {}

func (m *Unfreeze) GetMeansOpt() isUnfreeze_MeansOpt {
	if m != nil {
		return m.MeansOpt
	}
	return nil
}

func (m *Unfreeze) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

func (m *Unfreeze) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Unfreeze) GetAssetExec() string {
	if m != nil {
		return m.AssetExec
	}
	return ""
}

func (m *Unfreeze) GetAssetSymbol() string {
	if m != nil {
		return m.AssetSymbol
	}
	return ""
}

func (m *Unfreeze) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *Unfreeze) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *Unfreeze) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

func (m *Unfreeze) GetRemaining() int64 {
	if m != nil {
		return m.Remaining
	}
	return 0
}

func (m *Unfreeze) GetMeans() string {
	if m != nil {
		return m.Means
	}
	return ""
}

func (m *Unfreeze) GetFixAmount() *FixAmount {
	if x, ok := m.GetMeansOpt().(*Unfreeze_FixAmount); ok {
		return x.FixAmount
	}
	return nil
}

func (m *Unfreeze) GetLeftProportion() *LeftProportion {
	if x, ok := m.GetMeansOpt().(*Unfreeze_LeftProportion); ok {
		return x.LeftProportion
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Unfreeze) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Unfreeze_OneofMarshaler, _Unfreeze_OneofUnmarshaler, _Unfreeze_OneofSizer, []interface{}{
		(*Unfreeze_FixAmount)(nil),
		(*Unfreeze_LeftProportion)(nil),
	}
}

func _Unfreeze_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Unfreeze)
	// meansOpt
	switch x := m.MeansOpt.(type) {
	case *Unfreeze_FixAmount:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FixAmount); err != nil {
			return err
		}
	case *Unfreeze_LeftProportion:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LeftProportion); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Unfreeze.MeansOpt has unexpected type %T", x)
	}
	return nil
}

func _Unfreeze_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Unfreeze)
	switch tag {
	case 10: // meansOpt.fixAmount
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FixAmount)
		err := b.DecodeMessage(msg)
		m.MeansOpt = &Unfreeze_FixAmount{msg}
		return true, err
	case 11: // meansOpt.leftProportion
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LeftProportion)
		err := b.DecodeMessage(msg)
		m.MeansOpt = &Unfreeze_LeftProportion{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Unfreeze_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Unfreeze)
	// meansOpt
	switch x := m.MeansOpt.(type) {
	case *Unfreeze_FixAmount:
		s := proto.Size(x.FixAmount)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Unfreeze_LeftProportion:
		s := proto.Size(x.LeftProportion)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// 按时间固定额度解冻
type FixAmount struct {
	Period int64 `protobuf:"varint,1,opt,name=period" json:"period,omitempty"`
	Amount int64 `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *FixAmount) Reset()                    { *m = FixAmount{} }
func (m *FixAmount) String() string            { return proto.CompactTextString(m) }
func (*FixAmount) ProtoMessage()               {}
func (*FixAmount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FixAmount) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *FixAmount) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// 固定时间间隔按余量百分比解冻
type LeftProportion struct {
	Period        int64 `protobuf:"varint,1,opt,name=period" json:"period,omitempty"`
	TenThousandth int64 `protobuf:"varint,2,opt,name=tenThousandth" json:"tenThousandth,omitempty"`
}

func (m *LeftProportion) Reset()                    { *m = LeftProportion{} }
func (m *LeftProportion) String() string            { return proto.CompactTextString(m) }
func (*LeftProportion) ProtoMessage()               {}
func (*LeftProportion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LeftProportion) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *LeftProportion) GetTenThousandth() int64 {
	if m != nil {
		return m.TenThousandth
	}
	return 0
}

// message for execs.unfreeze
type UnfreezeAction struct {
	// Types that are valid to be assigned to Value:
	//	*UnfreezeAction_Create
	//	*UnfreezeAction_Withdraw
	//	*UnfreezeAction_Terminate
	Value isUnfreezeAction_Value `protobuf_oneof:"value"`
	Ty    int32                  `protobuf:"varint,4,opt,name=ty" json:"ty,omitempty"`
}

func (m *UnfreezeAction) Reset()                    { *m = UnfreezeAction{} }
func (m *UnfreezeAction) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeAction) ProtoMessage()               {}
func (*UnfreezeAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isUnfreezeAction_Value interface {
	isUnfreezeAction_Value()
}

type UnfreezeAction_Create struct {
	Create *UnfreezeCreate `protobuf:"bytes,1,opt,name=create,oneof"`
}
type UnfreezeAction_Withdraw struct {
	Withdraw *UnfreezeWithdraw `protobuf:"bytes,2,opt,name=withdraw,oneof"`
}
type UnfreezeAction_Terminate struct {
	Terminate *UnfreezeTerminate `protobuf:"bytes,3,opt,name=terminate,oneof"`
}

func (*UnfreezeAction_Create) isUnfreezeAction_Value()    {}
func (*UnfreezeAction_Withdraw) isUnfreezeAction_Value()  {}
func (*UnfreezeAction_Terminate) isUnfreezeAction_Value() {}

func (m *UnfreezeAction) GetValue() isUnfreezeAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *UnfreezeAction) GetCreate() *UnfreezeCreate {
	if x, ok := m.GetValue().(*UnfreezeAction_Create); ok {
		return x.Create
	}
	return nil
}

func (m *UnfreezeAction) GetWithdraw() *UnfreezeWithdraw {
	if x, ok := m.GetValue().(*UnfreezeAction_Withdraw); ok {
		return x.Withdraw
	}
	return nil
}

func (m *UnfreezeAction) GetTerminate() *UnfreezeTerminate {
	if x, ok := m.GetValue().(*UnfreezeAction_Terminate); ok {
		return x.Terminate
	}
	return nil
}

func (m *UnfreezeAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UnfreezeAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UnfreezeAction_OneofMarshaler, _UnfreezeAction_OneofUnmarshaler, _UnfreezeAction_OneofSizer, []interface{}{
		(*UnfreezeAction_Create)(nil),
		(*UnfreezeAction_Withdraw)(nil),
		(*UnfreezeAction_Terminate)(nil),
	}
}

func _UnfreezeAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UnfreezeAction)
	// value
	switch x := m.Value.(type) {
	case *UnfreezeAction_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *UnfreezeAction_Withdraw:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Withdraw); err != nil {
			return err
		}
	case *UnfreezeAction_Terminate:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Terminate); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UnfreezeAction.Value has unexpected type %T", x)
	}
	return nil
}

func _UnfreezeAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UnfreezeAction)
	switch tag {
	case 1: // value.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UnfreezeCreate)
		err := b.DecodeMessage(msg)
		m.Value = &UnfreezeAction_Create{msg}
		return true, err
	case 2: // value.withdraw
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UnfreezeWithdraw)
		err := b.DecodeMessage(msg)
		m.Value = &UnfreezeAction_Withdraw{msg}
		return true, err
	case 3: // value.terminate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UnfreezeTerminate)
		err := b.DecodeMessage(msg)
		m.Value = &UnfreezeAction_Terminate{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UnfreezeAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UnfreezeAction)
	// value
	switch x := m.Value.(type) {
	case *UnfreezeAction_Create:
		s := proto.Size(x.Create)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UnfreezeAction_Withdraw:
		s := proto.Size(x.Withdraw)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UnfreezeAction_Terminate:
		s := proto.Size(x.Terminate)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// action
type UnfreezeCreate struct {
	StartTime   int64  `protobuf:"varint,1,opt,name=startTime" json:"startTime,omitempty"`
	AssetExec   string `protobuf:"bytes,2,opt,name=assetExec" json:"assetExec,omitempty"`
	AssetSymbol string `protobuf:"bytes,3,opt,name=assetSymbol" json:"assetSymbol,omitempty"`
	TotalCount  int64  `protobuf:"varint,4,opt,name=totalCount" json:"totalCount,omitempty"`
	Beneficiary string `protobuf:"bytes,5,opt,name=beneficiary" json:"beneficiary,omitempty"`
	Means       string `protobuf:"bytes,6,opt,name=means" json:"means,omitempty"`
	// Types that are valid to be assigned to MeansOpt:
	//	*UnfreezeCreate_FixAmount
	//	*UnfreezeCreate_LeftProportion
	MeansOpt isUnfreezeCreate_MeansOpt `protobuf_oneof:"meansOpt"`
}

func (m *UnfreezeCreate) Reset()                    { *m = UnfreezeCreate{} }
func (m *UnfreezeCreate) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeCreate) ProtoMessage()               {}
func (*UnfreezeCreate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isUnfreezeCreate_MeansOpt interface {
	isUnfreezeCreate_MeansOpt()
}

type UnfreezeCreate_FixAmount struct {
	FixAmount *FixAmount `protobuf:"bytes,7,opt,name=fixAmount,oneof"`
}
type UnfreezeCreate_LeftProportion struct {
	LeftProportion *LeftProportion `protobuf:"bytes,8,opt,name=leftProportion,oneof"`
}

func (*UnfreezeCreate_FixAmount) isUnfreezeCreate_MeansOpt()      {}
func (*UnfreezeCreate_LeftProportion) isUnfreezeCreate_MeansOpt() {}

func (m *UnfreezeCreate) GetMeansOpt() isUnfreezeCreate_MeansOpt {
	if m != nil {
		return m.MeansOpt
	}
	return nil
}

func (m *UnfreezeCreate) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *UnfreezeCreate) GetAssetExec() string {
	if m != nil {
		return m.AssetExec
	}
	return ""
}

func (m *UnfreezeCreate) GetAssetSymbol() string {
	if m != nil {
		return m.AssetSymbol
	}
	return ""
}

func (m *UnfreezeCreate) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *UnfreezeCreate) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

func (m *UnfreezeCreate) GetMeans() string {
	if m != nil {
		return m.Means
	}
	return ""
}

func (m *UnfreezeCreate) GetFixAmount() *FixAmount {
	if x, ok := m.GetMeansOpt().(*UnfreezeCreate_FixAmount); ok {
		return x.FixAmount
	}
	return nil
}

func (m *UnfreezeCreate) GetLeftProportion() *LeftProportion {
	if x, ok := m.GetMeansOpt().(*UnfreezeCreate_LeftProportion); ok {
		return x.LeftProportion
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UnfreezeCreate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UnfreezeCreate_OneofMarshaler, _UnfreezeCreate_OneofUnmarshaler, _UnfreezeCreate_OneofSizer, []interface{}{
		(*UnfreezeCreate_FixAmount)(nil),
		(*UnfreezeCreate_LeftProportion)(nil),
	}
}

func _UnfreezeCreate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UnfreezeCreate)
	// meansOpt
	switch x := m.MeansOpt.(type) {
	case *UnfreezeCreate_FixAmount:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FixAmount); err != nil {
			return err
		}
	case *UnfreezeCreate_LeftProportion:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LeftProportion); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UnfreezeCreate.MeansOpt has unexpected type %T", x)
	}
	return nil
}

func _UnfreezeCreate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UnfreezeCreate)
	switch tag {
	case 7: // meansOpt.fixAmount
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FixAmount)
		err := b.DecodeMessage(msg)
		m.MeansOpt = &UnfreezeCreate_FixAmount{msg}
		return true, err
	case 8: // meansOpt.leftProportion
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LeftProportion)
		err := b.DecodeMessage(msg)
		m.MeansOpt = &UnfreezeCreate_LeftProportion{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UnfreezeCreate_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UnfreezeCreate)
	// meansOpt
	switch x := m.MeansOpt.(type) {
	case *UnfreezeCreate_FixAmount:
		s := proto.Size(x.FixAmount)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UnfreezeCreate_LeftProportion:
		s := proto.Size(x.LeftProportion)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type UnfreezeWithdraw struct {
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
}

func (m *UnfreezeWithdraw) Reset()                    { *m = UnfreezeWithdraw{} }
func (m *UnfreezeWithdraw) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeWithdraw) ProtoMessage()               {}
func (*UnfreezeWithdraw) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UnfreezeWithdraw) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

type UnfreezeTerminate struct {
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
}

func (m *UnfreezeTerminate) Reset()                    { *m = UnfreezeTerminate{} }
func (m *UnfreezeTerminate) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeTerminate) ProtoMessage()               {}
func (*UnfreezeTerminate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UnfreezeTerminate) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

// receipt
type ReceiptUnfreeze struct {
	Prev    *Unfreeze `protobuf:"bytes,1,opt,name=prev" json:"prev,omitempty"`
	Current *Unfreeze `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptUnfreeze) Reset()                    { *m = ReceiptUnfreeze{} }
func (m *ReceiptUnfreeze) String() string            { return proto.CompactTextString(m) }
func (*ReceiptUnfreeze) ProtoMessage()               {}
func (*ReceiptUnfreeze) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReceiptUnfreeze) GetPrev() *Unfreeze {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptUnfreeze) GetCurrent() *Unfreeze {
	if m != nil {
		return m.Current
	}
	return nil
}

// query
type QueryUnfreezeWithdraw struct {
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
}

func (m *QueryUnfreezeWithdraw) Reset()                    { *m = QueryUnfreezeWithdraw{} }
func (m *QueryUnfreezeWithdraw) String() string            { return proto.CompactTextString(m) }
func (*QueryUnfreezeWithdraw) ProtoMessage()               {}
func (*QueryUnfreezeWithdraw) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *QueryUnfreezeWithdraw) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

type ReplyQueryUnfreezeWithdraw struct {
	UnfreezeID      string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
	AvailableAmount int64  `protobuf:"varint,2,opt,name=availableAmount" json:"availableAmount,omitempty"`
}

func (m *ReplyQueryUnfreezeWithdraw) Reset()                    { *m = ReplyQueryUnfreezeWithdraw{} }
func (m *ReplyQueryUnfreezeWithdraw) String() string            { return proto.CompactTextString(m) }
func (*ReplyQueryUnfreezeWithdraw) ProtoMessage()               {}
func (*ReplyQueryUnfreezeWithdraw) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReplyQueryUnfreezeWithdraw) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

func (m *ReplyQueryUnfreezeWithdraw) GetAvailableAmount() int64 {
	if m != nil {
		return m.AvailableAmount
	}
	return 0
}

func init() {
	proto.RegisterType((*Unfreeze)(nil), "types.Unfreeze")
	proto.RegisterType((*FixAmount)(nil), "types.FixAmount")
	proto.RegisterType((*LeftProportion)(nil), "types.LeftProportion")
	proto.RegisterType((*UnfreezeAction)(nil), "types.UnfreezeAction")
	proto.RegisterType((*UnfreezeCreate)(nil), "types.UnfreezeCreate")
	proto.RegisterType((*UnfreezeWithdraw)(nil), "types.UnfreezeWithdraw")
	proto.RegisterType((*UnfreezeTerminate)(nil), "types.UnfreezeTerminate")
	proto.RegisterType((*ReceiptUnfreeze)(nil), "types.ReceiptUnfreeze")
	proto.RegisterType((*QueryUnfreezeWithdraw)(nil), "types.QueryUnfreezeWithdraw")
	proto.RegisterType((*ReplyQueryUnfreezeWithdraw)(nil), "types.ReplyQueryUnfreezeWithdraw")
}

func init() { proto.RegisterFile("unfreeze.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5f, 0x8f, 0xd2, 0x4e,
	0x14, 0xa5, 0x74, 0x0b, 0xf4, 0x92, 0x1f, 0xec, 0x6f, 0xe2, 0x6a, 0x63, 0x8c, 0x21, 0xd5, 0x07,
	0x7c, 0x41, 0xc3, 0xc6, 0x68, 0xe2, 0x83, 0x61, 0x57, 0x0d, 0x26, 0xc6, 0x3f, 0x15, 0xe3, 0xf3,
	0x50, 0x2e, 0x32, 0x49, 0x3b, 0xd3, 0x4c, 0xa7, 0xec, 0xd6, 0xaf, 0xe2, 0xe7, 0xf1, 0xc9, 0x2f,
	0x65, 0x3a, 0xfd, 0x03, 0x14, 0x57, 0x22, 0x8f, 0x73, 0xce, 0x3d, 0xe7, 0xde, 0xe9, 0x3d, 0x1d,
	0xe8, 0x25, 0x7c, 0x29, 0x11, 0xbf, 0xe3, 0x28, 0x92, 0x42, 0x09, 0x62, 0xa9, 0x34, 0xc2, 0xd8,
	0xfd, 0x61, 0x42, 0xe7, 0x4b, 0xc1, 0x90, 0xfb, 0x00, 0x65, 0xd5, 0xdb, 0x57, 0x8e, 0x31, 0x30,
	0x86, 0xb6, 0xb7, 0x85, 0x90, 0x7b, 0x60, 0xc7, 0x8a, 0x4a, 0x35, 0x63, 0x21, 0x3a, 0xcd, 0x81,
	0x31, 0x34, 0xbd, 0x0d, 0x90, 0xb1, 0x34, 0x8e, 0x51, 0xbd, 0xbe, 0x46, 0xdf, 0x31, 0xb5, 0x78,
	0x03, 0x90, 0x01, 0x74, 0xf5, 0xe1, 0x73, 0x1a, 0xce, 0x45, 0xe0, 0x9c, 0x68, 0x7e, 0x1b, 0xca,
	0xba, 0x2b, 0xa1, 0x68, 0x70, 0x29, 0x12, 0xae, 0x1c, 0x4b, 0xdb, 0x6f, 0x21, 0x99, 0x3f, 0xe3,
	0x4c, 0x31, 0xaa, 0x84, 0x74, 0x5a, 0xb9, 0x7f, 0x05, 0x64, 0xfe, 0x73, 0xe4, 0xb8, 0x64, 0x3e,
	0xa3, 0x32, 0x75, 0xda, 0xb9, 0xff, 0x16, 0x94, 0xe9, 0x25, 0x86, 0x94, 0x71, 0xc6, 0xbf, 0x39,
	0x9d, 0x7c, 0xfa, 0x0a, 0x20, 0xb7, 0xc0, 0x0a, 0x91, 0xf2, 0xd8, 0xb1, 0xb5, 0x32, 0x3f, 0x90,
	0x27, 0x60, 0x2f, 0xd9, 0xf5, 0x24, 0xd4, 0x23, 0xc1, 0xc0, 0x18, 0x76, 0xc7, 0xa7, 0x23, 0xfd,
	0xe5, 0x46, 0x6f, 0x4a, 0x7c, 0xda, 0xf0, 0x36, 0x45, 0xe4, 0x25, 0xf4, 0x02, 0x5c, 0xaa, 0x8f,
	0x52, 0x44, 0x42, 0x2a, 0x26, 0xb8, 0xd3, 0xd5, 0xb2, 0xb3, 0x42, 0xf6, 0x6e, 0x87, 0x9c, 0x36,
	0xbc, 0x5a, 0xf9, 0x05, 0x40, 0x47, 0xf7, 0xfe, 0x10, 0x29, 0xf7, 0x05, 0xd8, 0x55, 0x1b, 0x72,
	0x1b, 0x5a, 0x11, 0x4a, 0x26, 0x16, 0x7a, 0x33, 0xa6, 0x57, 0x9c, 0x32, 0x9c, 0xe6, 0x03, 0xe6,
	0x2b, 0x29, 0x4e, 0xee, 0x7b, 0xe8, 0xed, 0x36, 0xbb, 0xd1, 0xe1, 0x21, 0xfc, 0xa7, 0x90, 0xcf,
	0x56, 0x22, 0x89, 0x29, 0x5f, 0xa8, 0x55, 0x61, 0xb4, 0x0b, 0xba, 0xbf, 0x0c, 0xe8, 0x95, 0x51,
	0x99, 0xf8, 0xda, 0xf0, 0x31, 0xb4, 0x7c, 0x89, 0x54, 0xa1, 0x36, 0xdc, 0x5c, 0xb2, 0x2c, 0xbb,
	0xd4, 0xe4, 0xb4, 0xe1, 0x15, 0x65, 0xe4, 0x29, 0x74, 0xae, 0x98, 0x5a, 0x2d, 0x24, 0xbd, 0xd2,
	0x4d, 0xba, 0xe3, 0x3b, 0x35, 0xc9, 0xd7, 0x82, 0x9e, 0x36, 0xbc, 0xaa, 0x94, 0x3c, 0x07, 0x5b,
	0xa1, 0x0c, 0x19, 0xcf, 0x5a, 0x99, 0x5a, 0xe7, 0xd4, 0x74, 0xb3, 0x92, 0xcf, 0xd6, 0x51, 0x15,
	0x93, 0x1e, 0x34, 0x55, 0xaa, 0xd3, 0x66, 0x79, 0x4d, 0x95, 0x5e, 0xb4, 0xc1, 0x5a, 0xd3, 0x20,
	0x41, 0xf7, 0x67, 0x73, 0x73, 0x9b, 0x7c, 0xcc, 0xdd, 0x78, 0x1b, 0x7f, 0x8d, 0x77, 0xf3, 0x40,
	0xbc, 0xcd, 0x43, 0xf1, 0x3e, 0xd9, 0x8b, 0x77, 0x2d, 0xc0, 0xd6, 0x7e, 0x80, 0xab, 0x88, 0xb6,
	0x6e, 0x8c, 0x68, 0xfb, 0xb8, 0x88, 0x76, 0x8e, 0x8f, 0xe8, 0x18, 0x4e, 0xeb, 0xab, 0x3b, 0xf4,
	0x8e, 0xb8, 0xe7, 0xf0, 0xff, 0xde, 0xda, 0x0e, 0x8a, 0x28, 0xf4, 0x3d, 0xf4, 0x91, 0x45, 0xaa,
	0x7a, 0xaf, 0x1e, 0xc0, 0x49, 0x24, 0x71, 0x5d, 0x84, 0xaf, 0x5f, 0x4b, 0x84, 0xa7, 0x49, 0xf2,
	0x08, 0xda, 0x7e, 0x22, 0x25, 0x16, 0xff, 0xc7, 0x1f, 0xea, 0x4a, 0xde, 0x7d, 0x06, 0x67, 0x9f,
	0x12, 0x94, 0xe9, 0x3f, 0x5f, 0x68, 0x09, 0x77, 0x3d, 0x8c, 0x82, 0xf4, 0x28, 0x35, 0x19, 0x42,
	0x9f, 0xae, 0x29, 0x0b, 0xe8, 0x3c, 0xc0, 0xc9, 0xf6, 0x9f, 0x5c, 0x87, 0xe7, 0x2d, 0xfd, 0x76,
	0x9f, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x7b, 0x83, 0xa7, 0xcd, 0x05, 0x00, 0x00,
}
