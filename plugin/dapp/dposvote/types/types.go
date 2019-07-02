// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"reflect"

	"github.com/33cn/chain33/types"
)

func init() {
	// init executor type
	types.RegistorExecutor(DPosX, NewType())
	types.AllowUserExec = append(types.AllowUserExec, ExecerDposVote)
	types.RegisterDappFork(DPosX, "Enable", 0)
}

// GuessType struct
type DPosType struct {
	types.ExecTypeBase
}

// NewType method
func NewType() *DPosType {
	c := &DPosType{}
	c.SetChild(c)
	return c
}

// GetPayload method
func (t *DPosType) GetPayload() types.Message {
	return &DposVoteAction{}
}

// GetTypeMap method
func (t *DPosType) GetTypeMap() map[string]int32 {
	return map[string]int32{
		"Regist":         DposVoteActionRegist,
		"CancelRegist":   DposVoteActionCancelRegist,
		"ReRegist":       DposVoteActionReRegist,
		"Vote":           DposVoteActionVote,
		"CancelVote":     DposVoteActionCancelVote,
		"RegistVrfM":     DposVoteActionRegistVrfM,
		"RegistVrfRP":    DposVoteActionRegistVrfRP,
	}
}

// GetLogMap method
func (t *DPosType) GetLogMap() map[int64]*types.LogInfo {
	return map[int64]*types.LogInfo{
		TyLogCandicatorRegist:       {Ty: reflect.TypeOf(ReceiptCandicator{}), Name: "TyLogCandicatorRegist"},
		TyLogCandicatorVoted:        {Ty: reflect.TypeOf(ReceiptCandicator{}), Name: "TyLogCandicatorVoted"},
		TyLogCandicatorCancelRegist: {Ty: reflect.TypeOf(ReceiptCandicator{}), Name: "TyLogCandicatorCancelRegist"},
		TyLogVrfMRegist:             {Ty: reflect.TypeOf(ReceiptVrf{}),        Name: "TyLogVrfMRegist"},
		TyLogVrfRPRegist:            {Ty: reflect.TypeOf(ReceiptVrf{}),        Name: "TyLogVrfRPRegist"},
	}
}
