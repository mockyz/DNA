// SPDX-License-Identifier: LGPL-3.0-or-later
// Copyright 2019 DNA Dev team
//
/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */
package auth

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/DNAProject/DNA/common"
	"github.com/DNAProject/DNA/smartcontract/service/native/utils"
	"github.com/stretchr/testify/assert"
)

var (
	admin    []byte
	newAdmin []byte
	p1       []byte
	p2       []byte
)
var (
	funcs           = []string{"foo1", "foo2"}
	role            = "role"
	GasContractAddr = utils.GasContractAddress
)

func init() {
	admin = make([]byte, 32)
	newAdmin = make([]byte, 32)
	p1 = make([]byte, 20)
	p2 = make([]byte, 20)
	rand.Read(admin)
	rand.Read(newAdmin)
	rand.Read(p1)
	rand.Read(p2)
}
func TestSerialization_Init(t *testing.T) {
	param := &InitContractAdminParam{
		AdminOntID: admin,
	}
	sink := common.NewZeroCopySink(nil)
	param.Serialization(sink)
	rd := common.NewZeroCopySource(sink.Bytes())

	param2 := new(InitContractAdminParam)

	if err := param2.Deserialization(rd); err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(param.AdminOntID, param2.AdminOntID) != 0 {
		t.Fatalf("failed")
	}
}

func TestSerialization_Transfer(t *testing.T) {
	param := &TransferParam{
		ContractAddr:  GasContractAddr,
		NewAdminOntID: newAdmin,
	}
	sink := common.NewZeroCopySink(nil)
	param.Serialization(sink)
	rd := common.NewZeroCopySource(sink.Bytes())

	param2 := new(TransferParam)
	if err := param2.Deserialization(rd); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, param, param2)
}

func TestSerialization_AssignFuncs(t *testing.T) {
	param := &FuncsToRoleParam{
		ContractAddr: GasContractAddr,
		AdminOntID:   admin,
		Role:         []byte("role"),
		FuncNames:    funcs,
	}
	bf := common.NewZeroCopySink(nil)
	param.Serialization(bf)
	rd := common.NewZeroCopySource(bf.Bytes())

	param2 := new(FuncsToRoleParam)
	if err := param2.Deserialization(rd); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, param, param2)
}

func TestSerialization_AssignOntIDs(t *testing.T) {
	param := &DIDsToRoleParam{
		ContractAddr: GasContractAddr,
		AdminOntID:   admin,
		Role:         []byte(role),
		Persons:      [][]byte{[]byte{0x03, 0x04, 0x05, 0x06}, []byte{0x07, 0x08, 0x09, 0x0a}},
	}
	bf := common.NewZeroCopySink(nil)
	param.Serialization(bf)
	rd := common.NewZeroCopySource(bf.Bytes())
	param2 := new(DIDsToRoleParam)
	if err := param2.Deserialization(rd); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, param, param2)
}

func TestSerialization_Delegate(t *testing.T) {
	param := &DelegateParam{
		ContractAddr: GasContractAddr,
		From:         p1,
		To:           p2,
		Role:         []byte(role),
		Period:       60 * 60 * 24,
		Level:        3,
	}
	bf := common.NewZeroCopySink(nil)
	param.Serialization(bf)
	rd := common.NewZeroCopySource(bf.Bytes())
	param2 := new(DelegateParam)
	if err := param2.Deserialization(rd); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, param, param2)
}

func TestSerialization_Withdraw(t *testing.T) {
	param := &WithdrawParam{
		ContractAddr: GasContractAddr,
		Initiator:    p1,
		Delegate:     p2,
		Role:         []byte(role),
	}
	bf := common.NewZeroCopySink(nil)
	param.Serialization(bf)
	rd := common.NewZeroCopySource(bf.Bytes())
	param2 := new(WithdrawParam)
	if err := param2.Deserialization(rd); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, param, param2)
}

func TestSerialization_VerifyToken(t *testing.T) {
	param := &VerifyTokenParam{
		ContractAddr: GasContractAddr,
		Caller:       p1,
		Fn:           "foo1",
	}
	bf := common.NewZeroCopySink(nil)
	param.Serialization(bf)
	rd := common.NewZeroCopySource(bf.Bytes())
	param2 := new(VerifyTokenParam)
	if err := param2.Deserialization(rd); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, param, param2)
}
