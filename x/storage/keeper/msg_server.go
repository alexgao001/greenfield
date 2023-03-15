package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	types2 "github.com/bnb-chain/greenfield/types"
	gnfderrors "github.com/bnb-chain/greenfield/types/errors"
	permtypes "github.com/bnb-chain/greenfield/x/permission/types"
	"github.com/bnb-chain/greenfield/x/storage/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateBucket(goCtx context.Context, msg *types.MsgCreateBucket) (*types.MsgCreateBucketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ownerAcc, err := sdk.AccAddressFromHexUnsafe(msg.Creator)
	if err != nil {
		return nil, err
	}

	primarySPAcc, err := sdk.AccAddressFromHexUnsafe(msg.PrimarySpAddress)
	if err != nil {
		return nil, err
	}

	id, err := k.Keeper.CreateBucket(ctx, ownerAcc, msg.BucketName, primarySPAcc, CreateBucketOptions{
		PaymentAddress:    msg.PaymentAddress,
		IsPublic:          msg.IsPublic,
		ReadQuota:         msg.ReadQuota,
		SourceType:        types.SOURCE_TYPE_ORIGIN,
		PrimarySpApproval: msg.PrimarySpApproval,
		ApprovalMsgBytes:  msg.GetApprovalBytes(),
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateBucketResponse{
		BucketId: id,
	}, nil
}

func (k msgServer) DeleteBucket(goCtx context.Context, msg *types.MsgDeleteBucket) (*types.MsgDeleteBucketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	operatorAcc, err := sdk.AccAddressFromHexUnsafe(msg.Operator)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.DeleteBucket(ctx, operatorAcc, msg.BucketName, DeleteBucketOptions{
		SourceType: types.SOURCE_TYPE_ORIGIN,
	})
	if err != nil {
		return nil, err
	}
	return &types.MsgDeleteBucketResponse{}, nil
}

func (k msgServer) UpdateBucketInfo(goCtx context.Context, msg *types.MsgUpdateBucketInfo) (*types.MsgUpdateBucketInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	operatorAcc, err := sdk.AccAddressFromHexUnsafe(msg.Operator)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.UpdateBucketInfo(ctx, operatorAcc, msg.BucketName, UpdateBucketOptions{
		SourceType:     types.SOURCE_TYPE_ORIGIN,
		ReadQuota:      msg.ReadQuota,
		PaymentAddress: msg.PaymentAddress,
	})
	if err != nil {
		return nil, err
	}
	return &types.MsgUpdateBucketInfoResponse{}, nil
}

func (k msgServer) CreateObject(goCtx context.Context, msg *types.MsgCreateObject) (*types.MsgCreateObjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ownerAcc, err := sdk.AccAddressFromHexUnsafe(msg.Creator)
	if err != nil {
		return nil, err
	}

	if len(msg.ExpectChecksums) != int(1+k.GetExpectSecondarySPNumForECObject(ctx)) {
		return nil, gnfderrors.ErrInvalidChecksum.Wrapf("ExpectChecksums missing, expect: %d, actual: %d",
			1+k.Keeper.RedundantParityChunkNum(ctx)+k.Keeper.RedundantParityChunkNum(ctx),
			len(msg.ExpectChecksums))
	}

	id, err := k.Keeper.CreateObject(ctx, ownerAcc, msg.BucketName, msg.ObjectName, msg.PayloadSize, CreateObjectOptions{
		SourceType:           types.SOURCE_TYPE_ORIGIN,
		IsPublic:             msg.IsPublic,
		ContentType:          msg.ContentType,
		RedundancyType:       msg.RedundancyType,
		Checksums:            msg.ExpectChecksums,
		PrimarySpApproval:    msg.PrimarySpApproval,
		ApprovalMsgBytes:     msg.GetApprovalBytes(),
		SecondarySpAddresses: msg.ExpectSecondarySpAddresses,
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateObjectResponse{
		ObjectId: id,
	}, nil
}

func (k msgServer) SealObject(goCtx context.Context, msg *types.MsgSealObject) (*types.MsgSealObjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	spSealAcc, err := sdk.AccAddressFromHexUnsafe(msg.Operator)
	if err != nil {
		return nil, err
	}

	expectSecondarySPNum := k.GetExpectSecondarySPNumForECObject(ctx)
	if len(msg.SecondarySpAddresses) != (int)(expectSecondarySPNum) {
		return nil, errors.Wrapf(gnfderrors.ErrInvalidSPAddress, "Missing SP expect (%d), but (%d)", expectSecondarySPNum,
			len(msg.SecondarySpAddresses))
	}

	if len(msg.SecondarySpSignatures) != (int)(expectSecondarySPNum) {
		return nil, errors.Wrapf(gnfderrors.ErrInvalidSPSignature, "Missing SP signatures, expect (%d), but (%d)",
			expectSecondarySPNum, len(msg.SecondarySpSignatures))
	}

	err = k.Keeper.SealObject(ctx, spSealAcc, msg.BucketName, msg.ObjectName, SealObjectOptions{
		SecondarySpAddresses:  msg.SecondarySpAddresses,
		SecondarySpSignatures: msg.SecondarySpSignatures,
	})

	if err != nil {
		return nil, err
	}

	return &types.MsgSealObjectResponse{}, nil
}

func (k msgServer) CancelCreateObject(goCtx context.Context, msg *types.MsgCancelCreateObject) (*types.MsgCancelCreateObjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	operatorAcc, err := sdk.AccAddressFromHexUnsafe(msg.Operator)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.CancelCreateObject(ctx, operatorAcc, msg.BucketName, msg.ObjectName, CancelCreateObjectOptions{SourceType: types.SOURCE_TYPE_ORIGIN})
	if err != nil {
		return nil, err
	}

	return &types.MsgCancelCreateObjectResponse{}, nil
}

func (k msgServer) CopyObject(goCtx context.Context, msg *types.MsgCopyObject) (*types.MsgCopyObjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ownerAcc, err := sdk.AccAddressFromHexUnsafe(msg.Operator)
	if err != nil {
		return nil, err
	}

	id, err := k.Keeper.CopyObject(ctx, ownerAcc, msg.SrcBucketName, msg.SrcObjectName, msg.DstBucketName, msg.DstObjectName, CopyObjectOptions{
		SourceType:        types.SOURCE_TYPE_ORIGIN,
		IsPublic:          false, // TODO: Need Impl
		PrimarySpApproval: msg.DstPrimarySpApproval,
		ApprovalMsgBytes:  msg.GetApprovalBytes(),
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgCopyObjectResponse{
		ObjectId: id,
	}, nil
}

func (k msgServer) DeleteObject(goCtx context.Context, msg *types.MsgDeleteObject) (*types.MsgDeleteObjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	operatorAcc, err := sdk.AccAddressFromHexUnsafe(msg.Operator)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.DeleteObject(ctx, operatorAcc, msg.BucketName, msg.ObjectName, DeleteObjectOptions{
		SourceType: types.SOURCE_TYPE_ORIGIN,
	})

	if err != nil {
		return nil, err
	}
	return &types.MsgDeleteObjectResponse{}, nil
}

func (k msgServer) RejectSealObject(goCtx context.Context, msg *types.MsgRejectSealObject) (*types.MsgRejectSealObjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	spAcc, err := sdk.AccAddressFromHexUnsafe(msg.Operator)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.RejectSealObject(ctx, spAcc, msg.BucketName, msg.ObjectName)
	if err != nil {
		return nil, err
	}
	return &types.MsgRejectSealObjectResponse{}, nil
}

func (k msgServer) CreateGroup(goCtx context.Context, msg *types.MsgCreateGroup) (*types.MsgCreateGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ownerAcc, err := sdk.AccAddressFromHexUnsafe(msg.Creator)
	if err != nil {
		return nil, err
	}

	id, err := k.Keeper.CreateGroup(ctx, ownerAcc, msg.GroupName, CreateGroupOptions{Members: msg.Members})
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateGroupResponse{
		GroupId: id,
	}, nil
}

func (k msgServer) DeleteGroup(goCtx context.Context, msg *types.MsgDeleteGroup) (*types.MsgDeleteGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	operatorAcc, err := sdk.AccAddressFromHexUnsafe(msg.Operator)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.DeleteGroup(ctx, operatorAcc, msg.GroupName, DeleteGroupOptions{SourceType: types.SOURCE_TYPE_ORIGIN})
	if err != nil {
		return nil, err
	}

	return &types.MsgDeleteGroupResponse{}, nil
}

func (k msgServer) LeaveGroup(goCtx context.Context, msg *types.MsgLeaveGroup) (*types.MsgLeaveGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	memberAcc, err := sdk.AccAddressFromHexUnsafe(msg.Member)
	if err != nil {
		return nil, err
	}

	ownerAcc, err := sdk.AccAddressFromHexUnsafe(msg.GroupOwner)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.LeaveGroup(ctx, memberAcc, ownerAcc, msg.GroupName, LeaveGroupOptions{SourceType: types.SOURCE_TYPE_ORIGIN})
	if err != nil {
		return nil, err
	}

	return &types.MsgLeaveGroupResponse{}, nil
}

func (k msgServer) UpdateGroupMember(goCtx context.Context, msg *types.MsgUpdateGroupMember) (*types.MsgUpdateGroupMemberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	operator, err := sdk.AccAddressFromHexUnsafe(msg.Operator)
	if err != nil {
		return nil, err
	}
	// Now only allowed group owner to update member
	err = k.Keeper.UpdateGroupMember(ctx, operator, msg.GroupName, UpdateGroupMemberOptions{
		SourceType:      types.SOURCE_TYPE_ORIGIN,
		MembersToAdd:    msg.MembersToAdd,
		MembersToDelete: msg.MembersToDelete,
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateGroupMemberResponse{}, nil
}

func (k msgServer) PutPolicy(goCtx context.Context, msg *types.MsgPutPolicy) (*types.MsgPutPolicyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	operatorAddr, err := sdk.AccAddressFromHexUnsafe(msg.Operator)
	if err != nil {
		return nil, err
	}

	var grn types2.GRN
	err = grn.ParseFromString(msg.Resource, false)
	if err != nil {
		return nil, err
	}

	policy := &permtypes.Policy{
		ResourceType: grn.ResourceType(),
		Principal:    msg.Principal,
		Statements:   msg.Statements,
	}

	policyID, err := k.Keeper.PutPolicy(ctx, operatorAddr, grn, policy)
	if err != nil {
		return nil, err
	}
	return &types.MsgPutPolicyResponse{Id: policyID}, nil

}

func (k msgServer) DeletePolicy(goCtx context.Context, msg *types.MsgDeletePolicy) (*types.MsgDeletePolicyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	operator, err := sdk.AccAddressFromHexUnsafe(msg.Operator)
	if err != nil {
		return nil, err
	}

	var grn types2.GRN
	err = grn.ParseFromString(msg.Resource, false)
	if err != nil {
		return nil, err
	}

	policyID, err := k.Keeper.DeletePolicy(ctx, operator, msg.Principal, grn)
	if err != nil {
		return nil, err
	}

	return &types.MsgDeletePolicyResponse{Id: policyID}, nil
}
