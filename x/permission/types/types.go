package types

import (
	"regexp"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	gnfd "github.com/bnb-chain/greenfield/types"
	"github.com/bnb-chain/greenfield/types/common"
	"github.com/bnb-chain/greenfield/types/resource"
)

type (
	Int  = math.Int
	Uint = math.Uint
)

type VerifyOptions struct {
	Resource   string
	WantedSize *uint64
}

var (
	BucketAllowedActions = map[ActionType]bool{
		ACTION_UPDATE_BUCKET_INFO: true,
		ACTION_DELETE_BUCKET:      true,

		ACTION_CREATE_OBJECT:  true,
		ACTION_DELETE_OBJECT:  true,
		ACTION_GET_OBJECT:     true,
		ACTION_COPY_OBJECT:    true,
		ACTION_EXECUTE_OBJECT: true,
		ACTION_LIST_OBJECT:    true,

		ACTION_TYPE_ALL: true,
	}
	ObjectAllowedActions = map[ActionType]bool{
		ACTION_CREATE_OBJECT:  true,
		ACTION_DELETE_OBJECT:  true,
		ACTION_GET_OBJECT:     true,
		ACTION_COPY_OBJECT:    true,
		ACTION_EXECUTE_OBJECT: true,
		ACTION_LIST_OBJECT:    true,

		ACTION_TYPE_ALL: true,
	}
	GroupAllowedActions = map[ActionType]bool{
		ACTION_UPDATE_GROUP_MEMBER: true,
		ACTION_DELETE_GROUP:        true,

		ACTION_TYPE_ALL: true,
	}
)

func NewDefaultPolicyForGroupMember(groupID math.Uint, member sdk.AccAddress) *Policy {
	return &Policy{
		Principal:       NewPrincipalWithAccount(member),
		ResourceType:    resource.RESOURCE_TYPE_GROUP,
		ResourceId:      groupID,
		MemberStatement: NewMemberStatement(),
	}
}

func (p *Policy) Eval(action ActionType, blockTime time.Time, opts *VerifyOptions) (Effect, *Policy) {
	// 1. the policy is expired, need delete
	if p.ExpirationTime != nil && p.ExpirationTime.Before(blockTime) {
		// Notice: We do not actively delete policies that expire for users.
		return EFFECT_PASS, nil
	}
	allowed := false
	updated := false
	// 2. check all the statements
	for i, s := range p.Statements {
		if s.ExpirationTime != nil && s.ExpirationTime.Before(blockTime) {
			continue
		}
		e, updatedStatement := s.Eval(action, opts)
		// statement need to be updated
		if updatedStatement != nil {
			updated = true
			p.Statements[i] = updatedStatement
		}
		if e == EFFECT_DENY {
			return EFFECT_DENY, nil
		} else if e == EFFECT_ALLOW {
			allowed = true
		}
	}
	if allowed {
		if updated {
			return EFFECT_ALLOW, p
		} else {
			return EFFECT_ALLOW, nil
		}
	}
	return EFFECT_PASS, nil
}

func (p *Policy) GetGroupMemberStatement() (*Statement, bool) {
	for _, s := range p.Statements {
		for _, act := range s.Actions {
			if act == ACTION_GROUP_MEMBER {
				return s, true
			}
		}
	}
	return nil, false
}

func NewMemberStatement() *Statement {
	return &Statement{
		Effect:    EFFECT_ALLOW,
		Resources: nil,
		Actions:   []ActionType{ACTION_GROUP_MEMBER},
	}

}
func (s *Statement) Eval(action ActionType, opts *VerifyOptions) (Effect, *Statement) {
	// If 'resource' is not nil, it implies that the user intends to access a sub-resource, which would
	// be specified in 's.Resources'. Therefore, if the sub-resource in the statement is nil, we will ignore this statement.
	if opts != nil && opts.Resource != "" && s.Resources == nil {
		return EFFECT_PASS, nil
	}
	// If 'resource' is not nil, and 's.Resource' is also not nil, it indicates that we should verify whether
	// the resource that the user intends to access matches any items in 's.Resource'
	if opts != nil && opts.Resource != "" && s.Resources != nil {
		isMatch := false
		for _, res := range s.Resources {
			reg := regexp.MustCompile(res)
			if reg == nil {
				continue
			}
			matchRes := reg.MatchString(opts.Resource)
			if matchRes {
				isMatch = matchRes
				break
			}
		}
		if !isMatch {
			return EFFECT_PASS, nil
		}
	}

	for _, act := range s.Actions {
		if act == action || act == ACTION_TYPE_ALL {
			// Action matched, if effect is deny, then return deny
			if s.Effect == EFFECT_DENY {
				return EFFECT_DENY, nil
			}
			// There is special handling for ACTION_CREATE_OBJECT.
			// userA grant CreateObject permission to userB, but only allows he to create a limit size of object.
			// If exceeded, rejected
			if action == ACTION_CREATE_OBJECT && s.LimitSize != nil && opts != nil && opts.WantedSize != nil {
				if s.LimitSize.GetValue() >= *opts.WantedSize {
					s.LimitSize = &common.UInt64Value{Value: s.LimitSize.GetValue() - *opts.WantedSize}
					return EFFECT_ALLOW, s
				} else {
					return EFFECT_DENY, nil
				}
			}
			return s.Effect, nil
		}
	}

	return EFFECT_PASS, nil
}

func (s *Statement) ValidateBasic(resType resource.ResourceType) error {
	if s.Effect == EFFECT_PASS {
		return ErrInvalidStatement.Wrap("Not allowed to set EFFECT_PASS.")
	}
	switch resType {
	case resource.RESOURCE_TYPE_BUCKET:
		containsCreateObject := false
		for _, a := range s.Actions {
			if !BucketAllowedActions[a] {
				return ErrInvalidStatement.Wrapf("%s not allowed to be used on bucket.", a.String())
			}
			if a == ACTION_CREATE_OBJECT {
				containsCreateObject = true
			}
		}
		for _, r := range s.Resources {
			var grn gnfd.GRN
			err := grn.ParseFromString(r, true)
			if err != nil {
				return err
			}
		}

		if !containsCreateObject && s.LimitSize != nil {
			return ErrInvalidStatement.Wrap("The LimitSize option can only be used with CreateObject actions at the bucket level. .")
		}
	case resource.RESOURCE_TYPE_OBJECT:
		for _, a := range s.Actions {
			if !ObjectAllowedActions[a] {
				return ErrInvalidStatement.Wrapf("%s not allowed to be used on object.", a.String())
			}
		}
		if s.LimitSize != nil {
			return ErrInvalidStatement.Wrap("The LimitSize option can only be used with CreateObject actions at the bucket level. ")
		}
	case resource.RESOURCE_TYPE_GROUP:
		for _, a := range s.Actions {
			if !GroupAllowedActions[a] {
				return ErrInvalidStatement.Wrapf("%s not allowed to be used on group.", a.String())
			}
		}
		if s.LimitSize != nil {
			return ErrInvalidStatement.Wrap("The LimitSize option can only be used with CreateObject actions at the bucket level. ")
		}
	}

	return nil
}
