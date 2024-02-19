//
// Protos for function calls to the Staking adaptor.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: staking.proto

package steward_proto

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

// Represents call data for the Staking adaptor V1
type StakingAdaptorV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Function:
	//	*StakingAdaptorV1_RevokeApproval
	//	*StakingAdaptorV1_Mint_
	//	*StakingAdaptorV1_RequestBurn_
	//	*StakingAdaptorV1_CompleteBurn_
	//	*StakingAdaptorV1_CancelBurn_
	//	*StakingAdaptorV1_Wrap_
	//	*StakingAdaptorV1_Unwrap_
	//	*StakingAdaptorV1_MintErc20
	//	*StakingAdaptorV1_RemoveClaimedRequest_
	Function isStakingAdaptorV1_Function `protobuf_oneof:"function"`
}

func (x *StakingAdaptorV1) Reset() {
	*x = StakingAdaptorV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakingAdaptorV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakingAdaptorV1) ProtoMessage() {}

func (x *StakingAdaptorV1) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakingAdaptorV1.ProtoReflect.Descriptor instead.
func (*StakingAdaptorV1) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{0}
}

func (m *StakingAdaptorV1) GetFunction() isStakingAdaptorV1_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (x *StakingAdaptorV1) GetRevokeApproval() *RevokeApproval {
	if x, ok := x.GetFunction().(*StakingAdaptorV1_RevokeApproval); ok {
		return x.RevokeApproval
	}
	return nil
}

func (x *StakingAdaptorV1) GetMint() *StakingAdaptorV1_Mint {
	if x, ok := x.GetFunction().(*StakingAdaptorV1_Mint_); ok {
		return x.Mint
	}
	return nil
}

func (x *StakingAdaptorV1) GetRequestBurn() *StakingAdaptorV1_RequestBurn {
	if x, ok := x.GetFunction().(*StakingAdaptorV1_RequestBurn_); ok {
		return x.RequestBurn
	}
	return nil
}

func (x *StakingAdaptorV1) GetCompleteBurn() *StakingAdaptorV1_CompleteBurn {
	if x, ok := x.GetFunction().(*StakingAdaptorV1_CompleteBurn_); ok {
		return x.CompleteBurn
	}
	return nil
}

func (x *StakingAdaptorV1) GetCancelBurn() *StakingAdaptorV1_CancelBurn {
	if x, ok := x.GetFunction().(*StakingAdaptorV1_CancelBurn_); ok {
		return x.CancelBurn
	}
	return nil
}

func (x *StakingAdaptorV1) GetWrap() *StakingAdaptorV1_Wrap {
	if x, ok := x.GetFunction().(*StakingAdaptorV1_Wrap_); ok {
		return x.Wrap
	}
	return nil
}

func (x *StakingAdaptorV1) GetUnwrap() *StakingAdaptorV1_Unwrap {
	if x, ok := x.GetFunction().(*StakingAdaptorV1_Unwrap_); ok {
		return x.Unwrap
	}
	return nil
}

func (x *StakingAdaptorV1) GetMintErc20() *StakingAdaptorV1_MintERC20 {
	if x, ok := x.GetFunction().(*StakingAdaptorV1_MintErc20); ok {
		return x.MintErc20
	}
	return nil
}

func (x *StakingAdaptorV1) GetRemoveClaimedRequest() *StakingAdaptorV1_RemoveClaimedRequest {
	if x, ok := x.GetFunction().(*StakingAdaptorV1_RemoveClaimedRequest_); ok {
		return x.RemoveClaimedRequest
	}
	return nil
}

type isStakingAdaptorV1_Function interface {
	isStakingAdaptorV1_Function()
}

type StakingAdaptorV1_RevokeApproval struct {
	// Represents function `revokeApproval(ERC20 asset, address spender)`
	RevokeApproval *RevokeApproval `protobuf:"bytes,1,opt,name=revoke_approval,json=revokeApproval,proto3,oneof"`
}

type StakingAdaptorV1_Mint_ struct {
	// Represents function `mint(uint256 amount, uint256 minAmountOut, bytes calldata wildcard)`
	Mint *StakingAdaptorV1_Mint `protobuf:"bytes,2,opt,name=mint,proto3,oneof"`
}

type StakingAdaptorV1_RequestBurn_ struct {
	// Represents function `requestBurn(uint256 amount, bytes calldata wildcard)`
	RequestBurn *StakingAdaptorV1_RequestBurn `protobuf:"bytes,3,opt,name=request_burn,json=requestBurn,proto3,oneof"`
}

type StakingAdaptorV1_CompleteBurn_ struct {
	// Represents function `completeBurn(uint256 id, uint256 minAmountOut, bytes calldata wildcard)`
	CompleteBurn *StakingAdaptorV1_CompleteBurn `protobuf:"bytes,4,opt,name=complete_burn,json=completeBurn,proto3,oneof"`
}

type StakingAdaptorV1_CancelBurn_ struct {
	// Represents function `cancelBurn(uint256 id, bytes calldata wildcard)`
	CancelBurn *StakingAdaptorV1_CancelBurn `protobuf:"bytes,5,opt,name=cancel_burn,json=cancelBurn,proto3,oneof"`
}

type StakingAdaptorV1_Wrap_ struct {
	// Represents function `wrap(uint256 amount, uint256 minAmountOut, bytes calldata wildcard)`
	Wrap *StakingAdaptorV1_Wrap `protobuf:"bytes,6,opt,name=wrap,proto3,oneof"`
}

type StakingAdaptorV1_Unwrap_ struct {
	// Represents function `unwrap(uint256 amount, uint256 minAmountOut, bytes calldata wildcard)`
	Unwrap *StakingAdaptorV1_Unwrap `protobuf:"bytes,7,opt,name=unwrap,proto3,oneof"`
}

type StakingAdaptorV1_MintErc20 struct {
	// Represents function `mintERC20(ERC20 depositAsset, uint256 amount, uint256 minAmountOut, bytes calldata wildcard)`
	MintErc20 *StakingAdaptorV1_MintERC20 `protobuf:"bytes,8,opt,name=mint_erc20,json=mintErc20,proto3,oneof"`
}

type StakingAdaptorV1_RemoveClaimedRequest_ struct {
	// Represents function `removeClaimedRequest(uint256, bytes calldata)`
	RemoveClaimedRequest *StakingAdaptorV1_RemoveClaimedRequest `protobuf:"bytes,9,opt,name=remove_claimed_request,json=removeClaimedRequest,proto3,oneof"`
}

func (*StakingAdaptorV1_RevokeApproval) isStakingAdaptorV1_Function() {}

func (*StakingAdaptorV1_Mint_) isStakingAdaptorV1_Function() {}

func (*StakingAdaptorV1_RequestBurn_) isStakingAdaptorV1_Function() {}

func (*StakingAdaptorV1_CompleteBurn_) isStakingAdaptorV1_Function() {}

func (*StakingAdaptorV1_CancelBurn_) isStakingAdaptorV1_Function() {}

func (*StakingAdaptorV1_Wrap_) isStakingAdaptorV1_Function() {}

func (*StakingAdaptorV1_Unwrap_) isStakingAdaptorV1_Function() {}

func (*StakingAdaptorV1_MintErc20) isStakingAdaptorV1_Function() {}

func (*StakingAdaptorV1_RemoveClaimedRequest_) isStakingAdaptorV1_Function() {}

type StakingAdaptorV1Calls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calls []*StakingAdaptorV1 `protobuf:"bytes,1,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *StakingAdaptorV1Calls) Reset() {
	*x = StakingAdaptorV1Calls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakingAdaptorV1Calls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakingAdaptorV1Calls) ProtoMessage() {}

func (x *StakingAdaptorV1Calls) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakingAdaptorV1Calls.ProtoReflect.Descriptor instead.
func (*StakingAdaptorV1Calls) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{1}
}

func (x *StakingAdaptorV1Calls) GetCalls() []*StakingAdaptorV1 {
	if x != nil {
		return x.Calls
	}
	return nil
}

//
// Allows a strategist to `mint` a derivative asset using the chains native asset.
//
// Represents the function `mint(uint256 amount, uint256 minAmountOut, bytes calldata wildcard)`
type StakingAdaptorV1_Mint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The amount of the asset to mint
	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// The minimum amount of the asset to receive
	MinAmountOut string `protobuf:"bytes,2,opt,name=min_amount_out,json=minAmountOut,proto3" json:"min_amount_out,omitempty"`
	// Arbitrary ABI encoded data that can be used by inheriting adaptors
	Wildcard string `protobuf:"bytes,3,opt,name=wildcard,proto3" json:"wildcard,omitempty"`
}

func (x *StakingAdaptorV1_Mint) Reset() {
	*x = StakingAdaptorV1_Mint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakingAdaptorV1_Mint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakingAdaptorV1_Mint) ProtoMessage() {}

func (x *StakingAdaptorV1_Mint) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakingAdaptorV1_Mint.ProtoReflect.Descriptor instead.
func (*StakingAdaptorV1_Mint) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StakingAdaptorV1_Mint) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *StakingAdaptorV1_Mint) GetMinAmountOut() string {
	if x != nil {
		return x.MinAmountOut
	}
	return ""
}

func (x *StakingAdaptorV1_Mint) GetWildcard() string {
	if x != nil {
		return x.Wildcard
	}
	return ""
}

//
// Allows a strategist to request to burn/withdraw a derivative for a chains native asset.
//
// Represents the function `requestBurn(uint256 amount, bytes calldata wildcard)`
type StakingAdaptorV1_RequestBurn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The amount of the asset to burn
	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// Arbitrary ABI encoded data that can be used by inheriting adaptors
	Wildcard string `protobuf:"bytes,2,opt,name=wildcard,proto3" json:"wildcard,omitempty"`
}

func (x *StakingAdaptorV1_RequestBurn) Reset() {
	*x = StakingAdaptorV1_RequestBurn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakingAdaptorV1_RequestBurn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakingAdaptorV1_RequestBurn) ProtoMessage() {}

func (x *StakingAdaptorV1_RequestBurn) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakingAdaptorV1_RequestBurn.ProtoReflect.Descriptor instead.
func (*StakingAdaptorV1_RequestBurn) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{0, 1}
}

func (x *StakingAdaptorV1_RequestBurn) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *StakingAdaptorV1_RequestBurn) GetWildcard() string {
	if x != nil {
		return x.Wildcard
	}
	return ""
}

//
// Allows a strategist to complete a burn/withdraw of a derivative asset for a native asset.
//
// Represents the function `completeBurn(uint256 id, uint256 minAmountOut, bytes calldata wildcard)`
type StakingAdaptorV1_CompleteBurn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the burn request
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The minimum amount of the asset to receive
	MinAmountOut string `protobuf:"bytes,2,opt,name=min_amount_out,json=minAmountOut,proto3" json:"min_amount_out,omitempty"`
	// Arbitrary ABI encoded data that can be used by inheriting adaptors
	Wildcard string `protobuf:"bytes,3,opt,name=wildcard,proto3" json:"wildcard,omitempty"`
}

func (x *StakingAdaptorV1_CompleteBurn) Reset() {
	*x = StakingAdaptorV1_CompleteBurn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakingAdaptorV1_CompleteBurn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakingAdaptorV1_CompleteBurn) ProtoMessage() {}

func (x *StakingAdaptorV1_CompleteBurn) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakingAdaptorV1_CompleteBurn.ProtoReflect.Descriptor instead.
func (*StakingAdaptorV1_CompleteBurn) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{0, 2}
}

func (x *StakingAdaptorV1_CompleteBurn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StakingAdaptorV1_CompleteBurn) GetMinAmountOut() string {
	if x != nil {
		return x.MinAmountOut
	}
	return ""
}

func (x *StakingAdaptorV1_CompleteBurn) GetWildcard() string {
	if x != nil {
		return x.Wildcard
	}
	return ""
}

//
// Allows a strategist to cancel an active burn/withdraw request.
//
// Represents the function `cancelBurn(uint256 id, bytes calldata wildcard)`
type StakingAdaptorV1_CancelBurn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the burn request
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Arbitrary ABI encoded data that can be used by inheriting adaptors
	Wildcard string `protobuf:"bytes,2,opt,name=wildcard,proto3" json:"wildcard,omitempty"`
}

func (x *StakingAdaptorV1_CancelBurn) Reset() {
	*x = StakingAdaptorV1_CancelBurn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakingAdaptorV1_CancelBurn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakingAdaptorV1_CancelBurn) ProtoMessage() {}

func (x *StakingAdaptorV1_CancelBurn) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakingAdaptorV1_CancelBurn.ProtoReflect.Descriptor instead.
func (*StakingAdaptorV1_CancelBurn) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{0, 3}
}

func (x *StakingAdaptorV1_CancelBurn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StakingAdaptorV1_CancelBurn) GetWildcard() string {
	if x != nil {
		return x.Wildcard
	}
	return ""
}

//
// Allows a strategist to wrap a derivative asset.
//
// Represents the function `wrap(uint256 amount, uint256 minAmountOut, bytes calldata wildcard)`
type StakingAdaptorV1_Wrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The amount of the asset to wrap
	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// The minimum amount of the asset to receive
	MinAmountOut string `protobuf:"bytes,2,opt,name=min_amount_out,json=minAmountOut,proto3" json:"min_amount_out,omitempty"`
	// Arbitrary ABI encoded data that can be used by inheriting adaptors
	Wildcard string `protobuf:"bytes,3,opt,name=wildcard,proto3" json:"wildcard,omitempty"`
}

func (x *StakingAdaptorV1_Wrap) Reset() {
	*x = StakingAdaptorV1_Wrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakingAdaptorV1_Wrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakingAdaptorV1_Wrap) ProtoMessage() {}

func (x *StakingAdaptorV1_Wrap) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakingAdaptorV1_Wrap.ProtoReflect.Descriptor instead.
func (*StakingAdaptorV1_Wrap) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{0, 4}
}

func (x *StakingAdaptorV1_Wrap) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *StakingAdaptorV1_Wrap) GetMinAmountOut() string {
	if x != nil {
		return x.MinAmountOut
	}
	return ""
}

func (x *StakingAdaptorV1_Wrap) GetWildcard() string {
	if x != nil {
		return x.Wildcard
	}
	return ""
}

//
// Allows a strategist to unwrap a wrapped derivative asset.
//
// Represents the function `unwrap(uint256 amount, uint256 minAmountOut, bytes calldata wildcard)`
type StakingAdaptorV1_Unwrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The amount of the asset to unwrap
	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// The minimum amount of the asset to receive
	MinAmountOut string `protobuf:"bytes,2,opt,name=min_amount_out,json=minAmountOut,proto3" json:"min_amount_out,omitempty"`
	// Arbitrary ABI encoded data that can be used by inheriting adaptors
	Wildcard string `protobuf:"bytes,3,opt,name=wildcard,proto3" json:"wildcard,omitempty"`
}

func (x *StakingAdaptorV1_Unwrap) Reset() {
	*x = StakingAdaptorV1_Unwrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakingAdaptorV1_Unwrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakingAdaptorV1_Unwrap) ProtoMessage() {}

func (x *StakingAdaptorV1_Unwrap) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakingAdaptorV1_Unwrap.ProtoReflect.Descriptor instead.
func (*StakingAdaptorV1_Unwrap) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{0, 5}
}

func (x *StakingAdaptorV1_Unwrap) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *StakingAdaptorV1_Unwrap) GetMinAmountOut() string {
	if x != nil {
		return x.MinAmountOut
	}
	return ""
}

func (x *StakingAdaptorV1_Unwrap) GetWildcard() string {
	if x != nil {
		return x.Wildcard
	}
	return ""
}

//
// Allows a strategist to mint a derivative asset using an ERC20.
//
// Represents the function `mintERC20(ERC20 depositAsset, uint256 amount, uint256 minAmountOut, bytes calldata wildcard)`
type StakingAdaptorV1_MintERC20 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the ERC20 asset to deposit
	DepositAsset string `protobuf:"bytes,1,opt,name=deposit_asset,json=depositAsset,proto3" json:"deposit_asset,omitempty"`
	// The amount of the asset to mint
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// The minimum amount of the asset to receive
	MinAmountOut string `protobuf:"bytes,3,opt,name=min_amount_out,json=minAmountOut,proto3" json:"min_amount_out,omitempty"`
	// Arbitrary ABI encoded data that can be used by inheriting adaptors
	Wildcard string `protobuf:"bytes,4,opt,name=wildcard,proto3" json:"wildcard,omitempty"`
}

func (x *StakingAdaptorV1_MintERC20) Reset() {
	*x = StakingAdaptorV1_MintERC20{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakingAdaptorV1_MintERC20) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakingAdaptorV1_MintERC20) ProtoMessage() {}

func (x *StakingAdaptorV1_MintERC20) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakingAdaptorV1_MintERC20.ProtoReflect.Descriptor instead.
func (*StakingAdaptorV1_MintERC20) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{0, 6}
}

func (x *StakingAdaptorV1_MintERC20) GetDepositAsset() string {
	if x != nil {
		return x.DepositAsset
	}
	return ""
}

func (x *StakingAdaptorV1_MintERC20) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *StakingAdaptorV1_MintERC20) GetMinAmountOut() string {
	if x != nil {
		return x.MinAmountOut
	}
	return ""
}

func (x *StakingAdaptorV1_MintERC20) GetWildcard() string {
	if x != nil {
		return x.Wildcard
	}
	return ""
}

//
// Allows strategist to remove a request from `requestIds` if it has already been claimed.
//
// Represents the function `removeClaimedRequest(uint256, bytes calldata)`
type StakingAdaptorV1_RemoveClaimedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the request to remove
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Arbitrary ABI encoded data that can be used by inheriting adaptors
	Wildcard string `protobuf:"bytes,2,opt,name=wildcard,proto3" json:"wildcard,omitempty"`
}

func (x *StakingAdaptorV1_RemoveClaimedRequest) Reset() {
	*x = StakingAdaptorV1_RemoveClaimedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staking_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakingAdaptorV1_RemoveClaimedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakingAdaptorV1_RemoveClaimedRequest) ProtoMessage() {}

func (x *StakingAdaptorV1_RemoveClaimedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staking_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakingAdaptorV1_RemoveClaimedRequest.ProtoReflect.Descriptor instead.
func (*StakingAdaptorV1_RemoveClaimedRequest) Descriptor() ([]byte, []int) {
	return file_staking_proto_rawDescGZIP(), []int{0, 7}
}

func (x *StakingAdaptorV1_RemoveClaimedRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StakingAdaptorV1_RemoveClaimedRequest) GetWildcard() string {
	if x != nil {
		return x.Wildcard
	}
	return ""
}

var File_staking_proto protoreflect.FileDescriptor

var file_staking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x1a, 0x0a, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x0b, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x12, 0x45, 0x0a, 0x0f,
	0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x34, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x53,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x2e,
	0x4d, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x69, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x0c,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x75, 0x72, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e,
	0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x75, 0x72, 0x6e, 0x48, 0x00, 0x52, 0x0b,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x75, 0x72, 0x6e, 0x12, 0x50, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x72, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e,
	0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x72, 0x6e, 0x48, 0x00, 0x52,
	0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x72, 0x6e, 0x12, 0x4a, 0x0a,
	0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x62, 0x75, 0x72, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e,
	0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31,
	0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x75, 0x72, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x63,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x75, 0x72, 0x6e, 0x12, 0x37, 0x0a, 0x04, 0x77, 0x72, 0x61,
	0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x34, 0x2e, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x6f, 0x72, 0x56, 0x31, 0x2e, 0x57, 0x72, 0x61, 0x70, 0x48, 0x00, 0x52, 0x04, 0x77, 0x72,
	0x61, 0x70, 0x12, 0x3d, 0x0a, 0x06, 0x75, 0x6e, 0x77, 0x72, 0x61, 0x70, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e,
	0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31,
	0x2e, 0x55, 0x6e, 0x77, 0x72, 0x61, 0x70, 0x48, 0x00, 0x52, 0x06, 0x75, 0x6e, 0x77, 0x72, 0x61,
	0x70, 0x12, 0x47, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x74, 0x5f, 0x65, 0x72, 0x63, 0x32, 0x30, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x34, 0x2e, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f,
	0x72, 0x56, 0x31, 0x2e, 0x4d, 0x69, 0x6e, 0x74, 0x45, 0x52, 0x43, 0x32, 0x30, 0x48, 0x00, 0x52,
	0x09, 0x6d, 0x69, 0x6e, 0x74, 0x45, 0x72, 0x63, 0x32, 0x30, 0x12, 0x69, 0x0a, 0x16, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x74, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43,
	0x6c, 0x61, 0x69, 0x6d, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x14, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x60, 0x0a, 0x04, 0x4d, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d,
	0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x77,
	0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77,
	0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x1a, 0x41, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x75, 0x72, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x1a, 0x60, 0x0a, 0x0c, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x72, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x69,
	0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x1a, 0x38, 0x0a, 0x0a,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x75, 0x72, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x69,
	0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x69,
	0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x1a, 0x60, 0x0a, 0x04, 0x57, 0x72, 0x61, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6d, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x1a, 0x62, 0x0a, 0x06, 0x55, 0x6e, 0x77, 0x72,
	0x61, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x69,
	0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x1a, 0x8a, 0x01, 0x0a,
	0x09, 0x4d, 0x69, 0x6e, 0x74, 0x45, 0x52, 0x43, 0x32, 0x30, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x1a, 0x42, 0x0a, 0x14, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x42, 0x0a, 0x0a,
	0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x15, 0x53, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x43, 0x61, 0x6c,
	0x6c, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x53,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x52,
	0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x73, 0x74, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staking_proto_rawDescOnce sync.Once
	file_staking_proto_rawDescData = file_staking_proto_rawDesc
)

func file_staking_proto_rawDescGZIP() []byte {
	file_staking_proto_rawDescOnce.Do(func() {
		file_staking_proto_rawDescData = protoimpl.X.CompressGZIP(file_staking_proto_rawDescData)
	})
	return file_staking_proto_rawDescData
}

var file_staking_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_staking_proto_goTypes = []interface{}{
	(*StakingAdaptorV1)(nil),                      // 0: steward.v4.StakingAdaptorV1
	(*StakingAdaptorV1Calls)(nil),                 // 1: steward.v4.StakingAdaptorV1Calls
	(*StakingAdaptorV1_Mint)(nil),                 // 2: steward.v4.StakingAdaptorV1.Mint
	(*StakingAdaptorV1_RequestBurn)(nil),          // 3: steward.v4.StakingAdaptorV1.RequestBurn
	(*StakingAdaptorV1_CompleteBurn)(nil),         // 4: steward.v4.StakingAdaptorV1.CompleteBurn
	(*StakingAdaptorV1_CancelBurn)(nil),           // 5: steward.v4.StakingAdaptorV1.CancelBurn
	(*StakingAdaptorV1_Wrap)(nil),                 // 6: steward.v4.StakingAdaptorV1.Wrap
	(*StakingAdaptorV1_Unwrap)(nil),               // 7: steward.v4.StakingAdaptorV1.Unwrap
	(*StakingAdaptorV1_MintERC20)(nil),            // 8: steward.v4.StakingAdaptorV1.MintERC20
	(*StakingAdaptorV1_RemoveClaimedRequest)(nil), // 9: steward.v4.StakingAdaptorV1.RemoveClaimedRequest
	(*RevokeApproval)(nil),                        // 10: steward.v4.RevokeApproval
}
var file_staking_proto_depIdxs = []int32{
	10, // 0: steward.v4.StakingAdaptorV1.revoke_approval:type_name -> steward.v4.RevokeApproval
	2,  // 1: steward.v4.StakingAdaptorV1.mint:type_name -> steward.v4.StakingAdaptorV1.Mint
	3,  // 2: steward.v4.StakingAdaptorV1.request_burn:type_name -> steward.v4.StakingAdaptorV1.RequestBurn
	4,  // 3: steward.v4.StakingAdaptorV1.complete_burn:type_name -> steward.v4.StakingAdaptorV1.CompleteBurn
	5,  // 4: steward.v4.StakingAdaptorV1.cancel_burn:type_name -> steward.v4.StakingAdaptorV1.CancelBurn
	6,  // 5: steward.v4.StakingAdaptorV1.wrap:type_name -> steward.v4.StakingAdaptorV1.Wrap
	7,  // 6: steward.v4.StakingAdaptorV1.unwrap:type_name -> steward.v4.StakingAdaptorV1.Unwrap
	8,  // 7: steward.v4.StakingAdaptorV1.mint_erc20:type_name -> steward.v4.StakingAdaptorV1.MintERC20
	9,  // 8: steward.v4.StakingAdaptorV1.remove_claimed_request:type_name -> steward.v4.StakingAdaptorV1.RemoveClaimedRequest
	0,  // 9: steward.v4.StakingAdaptorV1Calls.calls:type_name -> steward.v4.StakingAdaptorV1
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_staking_proto_init() }
func file_staking_proto_init() {
	if File_staking_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_staking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakingAdaptorV1); i {
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
		file_staking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakingAdaptorV1Calls); i {
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
		file_staking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakingAdaptorV1_Mint); i {
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
		file_staking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakingAdaptorV1_RequestBurn); i {
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
		file_staking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakingAdaptorV1_CompleteBurn); i {
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
		file_staking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakingAdaptorV1_CancelBurn); i {
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
		file_staking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakingAdaptorV1_Wrap); i {
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
		file_staking_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakingAdaptorV1_Unwrap); i {
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
		file_staking_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakingAdaptorV1_MintERC20); i {
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
		file_staking_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakingAdaptorV1_RemoveClaimedRequest); i {
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
	file_staking_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StakingAdaptorV1_RevokeApproval)(nil),
		(*StakingAdaptorV1_Mint_)(nil),
		(*StakingAdaptorV1_RequestBurn_)(nil),
		(*StakingAdaptorV1_CompleteBurn_)(nil),
		(*StakingAdaptorV1_CancelBurn_)(nil),
		(*StakingAdaptorV1_Wrap_)(nil),
		(*StakingAdaptorV1_Unwrap_)(nil),
		(*StakingAdaptorV1_MintErc20)(nil),
		(*StakingAdaptorV1_RemoveClaimedRequest_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_staking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_staking_proto_goTypes,
		DependencyIndexes: file_staking_proto_depIdxs,
		MessageInfos:      file_staking_proto_msgTypes,
	}.Build()
	File_staking_proto = out.File
	file_staking_proto_rawDesc = nil
	file_staking_proto_goTypes = nil
	file_staking_proto_depIdxs = nil
}
