// Code generated by MockGen. DO NOT EDIT.
// Source: x/storage/types/expected_keepers.go

// Package types is a generated GoMock package.
package types

import (
	context "context"
	big "math/big"
	reflect "reflect"

	math "cosmossdk.io/math"
	resource "github.com/bnb-chain/greenfield/types/resource"
	types "github.com/bnb-chain/greenfield/x/payment/types"
	types0 "github.com/bnb-chain/greenfield/x/permission/types"
	types1 "github.com/bnb-chain/greenfield/x/sp/types"
	types2 "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx context.Context, addr types2.AccAddress) types2.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types2.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(name string) types2.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", name)
	ret0, _ := ret[0].(types2.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), name)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// GetAllBalances mocks base method.
func (m *MockBankKeeper) GetAllBalances(ctx types2.Context, addr types2.AccAddress) types2.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBalances", ctx, addr)
	ret0, _ := ret[0].(types2.Coins)
	return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances.
func (mr *MockBankKeeperMockRecorder) GetAllBalances(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankKeeper)(nil).GetAllBalances), ctx, addr)
}

// GetBalance mocks base method.
func (m *MockBankKeeper) GetBalance(ctx types2.Context, addr types2.AccAddress, denom string) types2.Coin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, addr, denom)
	ret0, _ := ret[0].(types2.Coin)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockBankKeeperMockRecorder) GetBalance(ctx, addr, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockBankKeeper)(nil).GetBalance), ctx, addr, denom)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx types2.Context, senderModule string, recipientAddr types2.AccAddress, amt types2.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx types2.Context, addr types2.AccAddress) types2.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types2.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}

// MockSpKeeper is a mock of SpKeeper interface.
type MockSpKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockSpKeeperMockRecorder
}

// MockSpKeeperMockRecorder is the mock recorder for MockSpKeeper.
type MockSpKeeperMockRecorder struct {
	mock *MockSpKeeper
}

// NewMockSpKeeper creates a new mock instance.
func NewMockSpKeeper(ctrl *gomock.Controller) *MockSpKeeper {
	mock := &MockSpKeeper{ctrl: ctrl}
	mock.recorder = &MockSpKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpKeeper) EXPECT() *MockSpKeeperMockRecorder {
	return m.recorder
}

// GetSpStoragePriceByTime mocks base method.
func (m *MockSpKeeper) GetSpStoragePriceByTime(ctx types2.Context, spAddr types2.AccAddress, time int64) (types1.SpStoragePrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpStoragePriceByTime", ctx, spAddr, time)
	ret0, _ := ret[0].(types1.SpStoragePrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpStoragePriceByTime indicates an expected call of GetSpStoragePriceByTime.
func (mr *MockSpKeeperMockRecorder) GetSpStoragePriceByTime(ctx, spAddr, time interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpStoragePriceByTime", reflect.TypeOf((*MockSpKeeper)(nil).GetSpStoragePriceByTime), ctx, spAddr, time)
}

// GetStorageProvider mocks base method.
func (m *MockSpKeeper) GetStorageProvider(ctx types2.Context, addr types2.AccAddress) (*types1.StorageProvider, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageProvider", ctx, addr)
	ret0, _ := ret[0].(*types1.StorageProvider)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetStorageProvider indicates an expected call of GetStorageProvider.
func (mr *MockSpKeeperMockRecorder) GetStorageProvider(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageProvider", reflect.TypeOf((*MockSpKeeper)(nil).GetStorageProvider), ctx, addr)
}

// GetStorageProviderByGcAddr mocks base method.
func (m *MockSpKeeper) GetStorageProviderByGcAddr(ctx types2.Context, gcAddr types2.AccAddress) (*types1.StorageProvider, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageProviderByGcAddr", ctx, gcAddr)
	ret0, _ := ret[0].(*types1.StorageProvider)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetStorageProviderByGcAddr indicates an expected call of GetStorageProviderByGcAddr.
func (mr *MockSpKeeperMockRecorder) GetStorageProviderByGcAddr(ctx, gcAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageProviderByGcAddr", reflect.TypeOf((*MockSpKeeper)(nil).GetStorageProviderByGcAddr), ctx, gcAddr)
}

// GetStorageProviderBySealAddr mocks base method.
func (m *MockSpKeeper) GetStorageProviderBySealAddr(ctx types2.Context, sealAddr types2.AccAddress) (*types1.StorageProvider, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageProviderBySealAddr", ctx, sealAddr)
	ret0, _ := ret[0].(*types1.StorageProvider)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetStorageProviderBySealAddr indicates an expected call of GetStorageProviderBySealAddr.
func (mr *MockSpKeeperMockRecorder) GetStorageProviderBySealAddr(ctx, sealAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageProviderBySealAddr", reflect.TypeOf((*MockSpKeeper)(nil).GetStorageProviderBySealAddr), ctx, sealAddr)
}

// IsStorageProviderExistAndInService mocks base method.
func (m *MockSpKeeper) IsStorageProviderExistAndInService(ctx types2.Context, addr types2.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsStorageProviderExistAndInService", ctx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsStorageProviderExistAndInService indicates an expected call of IsStorageProviderExistAndInService.
func (mr *MockSpKeeperMockRecorder) IsStorageProviderExistAndInService(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStorageProviderExistAndInService", reflect.TypeOf((*MockSpKeeper)(nil).IsStorageProviderExistAndInService), ctx, addr)
}

// SetSecondarySpStorePrice mocks base method.
func (m *MockSpKeeper) SetSecondarySpStorePrice(ctx types2.Context, secondarySpStorePrice types1.SecondarySpStorePrice) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSecondarySpStorePrice", ctx, secondarySpStorePrice)
}

// SetSecondarySpStorePrice indicates an expected call of SetSecondarySpStorePrice.
func (mr *MockSpKeeperMockRecorder) SetSecondarySpStorePrice(ctx, secondarySpStorePrice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSecondarySpStorePrice", reflect.TypeOf((*MockSpKeeper)(nil).SetSecondarySpStorePrice), ctx, secondarySpStorePrice)
}

// SetSpStoragePrice mocks base method.
func (m *MockSpKeeper) SetSpStoragePrice(ctx types2.Context, SpStoragePrice types1.SpStoragePrice) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSpStoragePrice", ctx, SpStoragePrice)
}

// SetSpStoragePrice indicates an expected call of SetSpStoragePrice.
func (mr *MockSpKeeperMockRecorder) SetSpStoragePrice(ctx, SpStoragePrice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSpStoragePrice", reflect.TypeOf((*MockSpKeeper)(nil).SetSpStoragePrice), ctx, SpStoragePrice)
}

// MockPaymentKeeper is a mock of PaymentKeeper interface.
type MockPaymentKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentKeeperMockRecorder
}

// MockPaymentKeeperMockRecorder is the mock recorder for MockPaymentKeeper.
type MockPaymentKeeperMockRecorder struct {
	mock *MockPaymentKeeper
}

// NewMockPaymentKeeper creates a new mock instance.
func NewMockPaymentKeeper(ctrl *gomock.Controller) *MockPaymentKeeper {
	mock := &MockPaymentKeeper{ctrl: ctrl}
	mock.recorder = &MockPaymentKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentKeeper) EXPECT() *MockPaymentKeeperMockRecorder {
	return m.recorder
}

// ApplyUserFlowsList mocks base method.
func (m *MockPaymentKeeper) ApplyUserFlowsList(ctx types2.Context, userFlows []types.UserFlows) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyUserFlowsList", ctx, userFlows)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyUserFlowsList indicates an expected call of ApplyUserFlowsList.
func (mr *MockPaymentKeeperMockRecorder) ApplyUserFlowsList(ctx, userFlows interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyUserFlowsList", reflect.TypeOf((*MockPaymentKeeper)(nil).ApplyUserFlowsList), ctx, userFlows)
}

// GetParams mocks base method.
func (m *MockPaymentKeeper) GetParams(ctx types2.Context) types.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockPaymentKeeperMockRecorder) GetParams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockPaymentKeeper)(nil).GetParams), ctx)
}

// GetStoragePrice mocks base method.
func (m *MockPaymentKeeper) GetStoragePrice(ctx types2.Context, params types.StoragePriceParams) (types.StoragePrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoragePrice", ctx, params)
	ret0, _ := ret[0].(types.StoragePrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoragePrice indicates an expected call of GetStoragePrice.
func (mr *MockPaymentKeeperMockRecorder) GetStoragePrice(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePrice", reflect.TypeOf((*MockPaymentKeeper)(nil).GetStoragePrice), ctx, params)
}

// GetStreamRecord mocks base method.
func (m *MockPaymentKeeper) GetStreamRecord(ctx types2.Context, account types2.AccAddress) (*types.StreamRecord, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStreamRecord", ctx, account)
	ret0, _ := ret[0].(*types.StreamRecord)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetStreamRecord indicates an expected call of GetStreamRecord.
func (mr *MockPaymentKeeperMockRecorder) GetStreamRecord(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStreamRecord", reflect.TypeOf((*MockPaymentKeeper)(nil).GetStreamRecord), ctx, account)
}

// IsPaymentAccountOwner mocks base method.
func (m *MockPaymentKeeper) IsPaymentAccountOwner(ctx types2.Context, addr, owner types2.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPaymentAccountOwner", ctx, addr, owner)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPaymentAccountOwner indicates an expected call of IsPaymentAccountOwner.
func (mr *MockPaymentKeeperMockRecorder) IsPaymentAccountOwner(ctx, addr, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPaymentAccountOwner", reflect.TypeOf((*MockPaymentKeeper)(nil).IsPaymentAccountOwner), ctx, addr, owner)
}

// UpdateStreamRecordByAddr mocks base method.
func (m *MockPaymentKeeper) UpdateStreamRecordByAddr(ctx types2.Context, change *types.StreamRecordChange) (*types.StreamRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStreamRecordByAddr", ctx, change)
	ret0, _ := ret[0].(*types.StreamRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStreamRecordByAddr indicates an expected call of UpdateStreamRecordByAddr.
func (mr *MockPaymentKeeperMockRecorder) UpdateStreamRecordByAddr(ctx, change interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStreamRecordByAddr", reflect.TypeOf((*MockPaymentKeeper)(nil).UpdateStreamRecordByAddr), ctx, change)
}

// MockPermissionKeeper is a mock of PermissionKeeper interface.
type MockPermissionKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockPermissionKeeperMockRecorder
}

// MockPermissionKeeperMockRecorder is the mock recorder for MockPermissionKeeper.
type MockPermissionKeeperMockRecorder struct {
	mock *MockPermissionKeeper
}

// NewMockPermissionKeeper creates a new mock instance.
func NewMockPermissionKeeper(ctrl *gomock.Controller) *MockPermissionKeeper {
	mock := &MockPermissionKeeper{ctrl: ctrl}
	mock.recorder = &MockPermissionKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPermissionKeeper) EXPECT() *MockPermissionKeeperMockRecorder {
	return m.recorder
}

// AddGroupMember mocks base method.
func (m *MockPermissionKeeper) AddGroupMember(ctx types2.Context, groupID math.Uint, member types2.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddGroupMember", ctx, groupID, member)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddGroupMember indicates an expected call of AddGroupMember.
func (mr *MockPermissionKeeperMockRecorder) AddGroupMember(ctx, groupID, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGroupMember", reflect.TypeOf((*MockPermissionKeeper)(nil).AddGroupMember), ctx, groupID, member)
}

// DeletePolicy mocks base method.
func (m *MockPermissionKeeper) DeletePolicy(ctx types2.Context, principal *types0.Principal, resourceType resource.ResourceType, resourceID math.Uint) (math.Uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePolicy", ctx, principal, resourceType, resourceID)
	ret0, _ := ret[0].(math.Uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePolicy indicates an expected call of DeletePolicy.
func (mr *MockPermissionKeeperMockRecorder) DeletePolicy(ctx, principal, resourceType, resourceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePolicy", reflect.TypeOf((*MockPermissionKeeper)(nil).DeletePolicy), ctx, principal, resourceType, resourceID)
}

// GetGroupMember mocks base method.
func (m *MockPermissionKeeper) GetGroupMember(ctx types2.Context, groupID math.Uint, member types2.AccAddress) (*types0.GroupMember, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMember", ctx, groupID, member)
	ret0, _ := ret[0].(*types0.GroupMember)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetGroupMember indicates an expected call of GetGroupMember.
func (mr *MockPermissionKeeperMockRecorder) GetGroupMember(ctx, groupID, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMember", reflect.TypeOf((*MockPermissionKeeper)(nil).GetGroupMember), ctx, groupID, member)
}

// GetGroupMemberByID mocks base method.
func (m *MockPermissionKeeper) GetGroupMemberByID(ctx types2.Context, groupMemberID math.Uint) (*types0.GroupMember, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMemberByID", ctx, groupMemberID)
	ret0, _ := ret[0].(*types0.GroupMember)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetGroupMemberByID indicates an expected call of GetGroupMemberByID.
func (mr *MockPermissionKeeperMockRecorder) GetGroupMemberByID(ctx, groupMemberID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMemberByID", reflect.TypeOf((*MockPermissionKeeper)(nil).GetGroupMemberByID), ctx, groupMemberID)
}

// GetPolicyByID mocks base method.
func (m *MockPermissionKeeper) GetPolicyByID(ctx types2.Context, policyID math.Uint) (*types0.Policy, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyByID", ctx, policyID)
	ret0, _ := ret[0].(*types0.Policy)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPolicyByID indicates an expected call of GetPolicyByID.
func (mr *MockPermissionKeeperMockRecorder) GetPolicyByID(ctx, policyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyByID", reflect.TypeOf((*MockPermissionKeeper)(nil).GetPolicyByID), ctx, policyID)
}

// GetPolicyForAccount mocks base method.
func (m *MockPermissionKeeper) GetPolicyForAccount(ctx types2.Context, resourceID math.Uint, resourceType resource.ResourceType, addr types2.AccAddress) (*types0.Policy, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyForAccount", ctx, resourceID, resourceType, addr)
	ret0, _ := ret[0].(*types0.Policy)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPolicyForAccount indicates an expected call of GetPolicyForAccount.
func (mr *MockPermissionKeeperMockRecorder) GetPolicyForAccount(ctx, resourceID, resourceType, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyForAccount", reflect.TypeOf((*MockPermissionKeeper)(nil).GetPolicyForAccount), ctx, resourceID, resourceType, addr)
}

// GetPolicyForGroup mocks base method.
func (m *MockPermissionKeeper) GetPolicyForGroup(ctx types2.Context, resourceID math.Uint, resourceType resource.ResourceType, groupID math.Uint) (*types0.Policy, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyForGroup", ctx, resourceID, resourceType, groupID)
	ret0, _ := ret[0].(*types0.Policy)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPolicyForGroup indicates an expected call of GetPolicyForGroup.
func (mr *MockPermissionKeeperMockRecorder) GetPolicyForGroup(ctx, resourceID, resourceType, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyForGroup", reflect.TypeOf((*MockPermissionKeeper)(nil).GetPolicyForGroup), ctx, resourceID, resourceType, groupID)
}

// PutPolicy mocks base method.
func (m *MockPermissionKeeper) PutPolicy(ctx types2.Context, policy *types0.Policy) (math.Uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutPolicy", ctx, policy)
	ret0, _ := ret[0].(math.Uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutPolicy indicates an expected call of PutPolicy.
func (mr *MockPermissionKeeperMockRecorder) PutPolicy(ctx, policy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPolicy", reflect.TypeOf((*MockPermissionKeeper)(nil).PutPolicy), ctx, policy)
}

// RemoveGroupMember mocks base method.
func (m *MockPermissionKeeper) RemoveGroupMember(ctx types2.Context, groupID math.Uint, member types2.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveGroupMember", ctx, groupID, member)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveGroupMember indicates an expected call of RemoveGroupMember.
func (mr *MockPermissionKeeperMockRecorder) RemoveGroupMember(ctx, groupID, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveGroupMember", reflect.TypeOf((*MockPermissionKeeper)(nil).RemoveGroupMember), ctx, groupID, member)
}

// VerifyPolicy mocks base method.
func (m *MockPermissionKeeper) VerifyPolicy(ctx types2.Context, resourceID math.Uint, resourceType resource.ResourceType, operator types2.AccAddress, action types0.ActionType, opts *types0.VerifyOptions) types0.Effect {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPolicy", ctx, resourceID, resourceType, operator, action, opts)
	ret0, _ := ret[0].(types0.Effect)
	return ret0
}

// VerifyPolicy indicates an expected call of VerifyPolicy.
func (mr *MockPermissionKeeperMockRecorder) VerifyPolicy(ctx, resourceID, resourceType, operator, action, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPolicy", reflect.TypeOf((*MockPermissionKeeper)(nil).VerifyPolicy), ctx, resourceID, resourceType, operator, action, opts)
}

// MockCrossChainKeeper is a mock of CrossChainKeeper interface.
type MockCrossChainKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockCrossChainKeeperMockRecorder
}

// MockCrossChainKeeperMockRecorder is the mock recorder for MockCrossChainKeeper.
type MockCrossChainKeeperMockRecorder struct {
	mock *MockCrossChainKeeper
}

// NewMockCrossChainKeeper creates a new mock instance.
func NewMockCrossChainKeeper(ctrl *gomock.Controller) *MockCrossChainKeeper {
	mock := &MockCrossChainKeeper{ctrl: ctrl}
	mock.recorder = &MockCrossChainKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCrossChainKeeper) EXPECT() *MockCrossChainKeeperMockRecorder {
	return m.recorder
}

// CreateRawIBCPackageWithFee mocks base method.
func (m *MockCrossChainKeeper) CreateRawIBCPackageWithFee(ctx types2.Context, channelID types2.ChannelID, packageType types2.CrossChainPackageType, packageLoad []byte, relayerFee, ackRelayerFee *big.Int) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRawIBCPackageWithFee", ctx, channelID, packageType, packageLoad, relayerFee, ackRelayerFee)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRawIBCPackageWithFee indicates an expected call of CreateRawIBCPackageWithFee.
func (mr *MockCrossChainKeeperMockRecorder) CreateRawIBCPackageWithFee(ctx, channelID, packageType, packageLoad, relayerFee, ackRelayerFee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRawIBCPackageWithFee", reflect.TypeOf((*MockCrossChainKeeper)(nil).CreateRawIBCPackageWithFee), ctx, channelID, packageType, packageLoad, relayerFee, ackRelayerFee)
}

// RegisterChannel mocks base method.
func (m *MockCrossChainKeeper) RegisterChannel(name string, id types2.ChannelID, app types2.CrossChainApplication) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterChannel", name, id, app)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterChannel indicates an expected call of RegisterChannel.
func (mr *MockCrossChainKeeperMockRecorder) RegisterChannel(name, id, app interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterChannel", reflect.TypeOf((*MockCrossChainKeeper)(nil).RegisterChannel), name, id, app)
}

// ForceDeleteAccountPolicyForResource mocks base method.
func (m *MockPermissionKeeper) ForceDeleteAccountPolicyForResource(ctx types2.Context, maxDelete, deletedCount uint64, resourceType resource.ResourceType, resourceID math.Uint) (uint64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceDeleteAccountPolicyForResource", ctx, resourceType, resourceID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ForceDeleteAccountPolicyForResource indicates an expected call of ForceDeleteAccountPolicyForResource.
func (mr *MockPermissionKeeperMockRecorder) ForceDeleteAccountPolicyForResource(ctx, maxDelete, deletedCount, resourceType, resourceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceDeleteAccountPolicyForResource", reflect.TypeOf((*MockPermissionKeeper)(nil).ForceDeleteAccountPolicyForResource), ctx, maxDelete, deletedCount, resourceType, resourceID)
}

// ForceDeleteGroupPolicyForResource mocks base method.
func (m *MockPermissionKeeper) ForceDeleteGroupPolicyForResource(ctx types2.Context, maxDelete, deletedCount uint64, resourceType resource.ResourceType, resourceID math.Uint) (uint64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceDeleteGroupPolicyForResource", ctx, resourceType, resourceID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ForceDeleteGroupPolicyForResource indicates an expected call of ForceDeleteGroupPolicyForResource.
func (mr *MockPermissionKeeperMockRecorder) ForceDeleteGroupPolicyForResource(ctx, maxDelete, deletedCount, resourceType, resourceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceDeleteGroupPolicyForResource", reflect.TypeOf((*MockPermissionKeeper)(nil).ForceDeleteGroupPolicyForResource), ctx, maxDelete, deletedCount, resourceType, resourceID)
}

// ForceDeleteGroupMembers mocks base method.
func (m *MockPermissionKeeper) ForceDeleteGroupMembers(ctx types2.Context, groupId math.Uint) {
	m.ctrl.T.Helper()
	_ = m.ctrl.Call(m, "ForceDeleteGroupMembers", ctx, groupId)
}

// ForceDeleteGroupMembers indicates an expected call of ForceDeleteGroupMembers.
func (mr *MockPermissionKeeperMockRecorder) ForceDeleteGroupMembers(ctx, groupId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceDeleteGroupMembers", reflect.TypeOf((*MockPermissionKeeper)(nil).ForceDeleteGroupMembers), ctx, groupId)
}
