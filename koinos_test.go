//   ____                           _           _   _____         _
//  / ___| ___ _ __   ___ _ __ __ _| |_ ___  __| | |_   _|__  ___| |_ ___
// | |  _ / _ \ '_ \ / _ \ '__/ _` | __/ _ \/ _` |   | |/ _ \/ __| __/ __|
// | |_| |  __/ | | |  __/ | | (_| | ||  __/ (_| |   | |  __/\__ \ |_\__ \
//  \____|\___|_| |_|\___|_|  \__,_|\__\___|\__,_|   |_|\___||___/\__|___/
//                         Please do not modify

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "koinos"
	"testing"
)

// ----------------------------------------
//  Struct: UnusedExtensionsType
// ----------------------------------------

func TestUnusedExtensionsType(t *testing.T) {
	o := NewUnusedExtensionsType()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeUnusedExtensionsType(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewUnusedExtensionsType()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: SignatureType
// ----------------------------------------

func TestSignatureType(t *testing.T) {
	o := NewSignatureType()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeSignatureType(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeSignatureType(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewSignatureType()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: ContractIDType
// ----------------------------------------

func TestContractIDType(t *testing.T) {
	o := NewContractIDType()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeContractIDType(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeContractIDType(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewContractIDType()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: ReservedOperation
// ----------------------------------------

func TestReservedOperation(t *testing.T) {
	o := NewReservedOperation()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeReservedOperation(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewReservedOperation()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: NopOperation
// ----------------------------------------

func TestNopOperation(t *testing.T) {
	o := NewNopOperation()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeNopOperation(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewNopOperation()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: CreateSystemContractOperation
// ----------------------------------------

func TestCreateSystemContractOperation(t *testing.T) {
	o := NewCreateSystemContractOperation()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeCreateSystemContractOperation(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test contract_id
	vb = &VariableBlob{}
	n, _, err = DeserializeCreateSystemContractOperation(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test bytecode
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeCreateSystemContractOperation(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewCreateSystemContractOperation()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: SystemCallTargetReserved
// ----------------------------------------

func TestSystemCallTargetReserved(t *testing.T) {
	o := NewSystemCallTargetReserved()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeSystemCallTargetReserved(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewSystemCallTargetReserved()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Enum: ThunkID
// ----------------------------------------

func TestThunkID(t *testing.T) {
	vals := []ThunkID{
		ThunkIDPrints,
		ThunkIDVerifyBlockHeader,
		ThunkIDApplyBlock,
		ThunkIDApplyTransaction,
		ThunkIDApplyReservedOperation,
		ThunkIDApplyUploadContractOperation,
		ThunkIDApplyExecuteContractOperation,
		ThunkIDApplySetSystemCallOperation,
		ThunkIDDbPutObject,
		ThunkIDDbGetObject,
		ThunkIDDbGetNextObject,
		ThunkIDDbGetPrevObject,
		ThunkIDExecuteContract,
		ThunkIDGetContractArgsSize,
		ThunkIDGetContractArgs,
		ThunkIDSetContractReturn,
		ThunkIDExitContract,
		ThunkIDGetHeadInfo,
		ThunkIDHash,
		ThunkIDVerifyBlockSig,
		ThunkIDVerifyMerkleRoot,
	}

	// Make sure all types properly serialize
	for _, x := range vals {
		vb := NewVariableBlob()
		vb = x.Serialize(vb)

		nvb := NewVariableBlob()
		ox := UInt32(x)
		nvb = ox.Serialize(nvb)

		if !bytes.Equal(*vb, *nvb) {
			t.Errorf("Serialized enum does match ideal serialization.")
		}

		_, y, err := DeserializeThunkID(vb)
		if err != nil {
			t.Error(err)
		}

		if x != *y {
			t.Errorf("Deserialized enum does not match original.")
		}

		// Test JSON
		jx, err := json.Marshal(x)
		if err != nil {
			t.Error(err)
		}

		jox, err := json.Marshal(ox)
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(jx, jox) {
			t.Error(err)
		}

		r := NewThunkID()
		if err = json.Unmarshal(jx, r); err != nil {
			t.Error(err)
		}
	}

	// Find a value that is NOT an enum value
	w := getInvalidThunkID()

	tw := UInt32(w)
	vb := NewVariableBlob()
	n, _, err := DeserializeThunkID(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	vb = tw.Serialize(vb)

	if _, _, err := DeserializeThunkID(vb); err == nil {
		t.Errorf("Deserializing an invalid value did not return an error.")
	}

	je := NewThunkID()
	if err := json.Unmarshal([]byte(fmt.Sprint(tw)), &je); err == nil {
		t.Errorf("Deserializing an invalid JSON value did not return an error.")
	}

	if err = json.Unmarshal([]byte("\"foobar\""), &je); err == nil {
		t.Errorf("err == nil")
	}
}

func TestThunkIDPanic(t *testing.T) {
	// Find a value that is NOT an enum value
	w := getInvalidThunkID()

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an invalid enum value did not panic.")
			}
		}()

		vb := NewVariableBlob()
		vb = w.Serialize(vb)
	}()

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an invalid enum value did not panic.")
			}
		}()

		_, _ = json.Marshal(w)
	}()
}

func getInvalidThunkID() ThunkID {
	w := ThunkIDVerifyMerkleRoot
	for IsValidThunkID(w) {
		w++
	}

	return w
}
// ----------------------------------------
//  Struct: ContractCallBundle
// ----------------------------------------

func TestContractCallBundle(t *testing.T) {
	o := NewContractCallBundle()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeContractCallBundle(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test contract_id
	vb = &VariableBlob{}
	n, _, err = DeserializeContractCallBundle(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test entry_point
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeContractCallBundle(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewContractCallBundle()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Variant: SystemCallTarget
// ----------------------------------------

func TestSystemCallTarget(t *testing.T) {
	o := NewSystemCallTarget()
	exerciseSystemCallTargetSerialization(o, t)
	{
		v := NewSystemCallTarget()
		v.Value = NewSystemCallTargetReserved()
		exerciseSystemCallTargetSerialization(v, t)

	}
	{
		v := NewSystemCallTarget()
		v.Value = NewThunkID()
		exerciseSystemCallTargetSerialization(v, t)

		vb := VariableBlob{1}
		n, _, err := DeserializeSystemCallTarget(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := NewSystemCallTarget()
		v.Value = NewContractCallBundle()
		exerciseSystemCallTargetSerialization(v, t)

		vb := VariableBlob{2}
		n, _, err := DeserializeSystemCallTarget(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}

	// Test bad variant tag
	vb := VariableBlob{0x80}
	n, _, err := DeserializeSystemCallTarget(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = VariableBlob{3}
	n, _, err = DeserializeSystemCallTarget(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = NewSystemCallTarget()
	if jerr := json.Unmarshal([]byte("\"!@#$%^&*\""), o); jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	// Test bad tag
	if jerr := json.Unmarshal([]byte("{\"type\":\"foobar\", \"value\":0}"), o); jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	// Test panic when serializing an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := SystemCallTarget{Value: int64(0)}
		vb := NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := SystemCallTarget{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseSystemCallTargetSerialization(v *SystemCallTarget, t *testing.T) {
	vb := NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := DeserializeSystemCallTarget(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := NewSystemCallTarget()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Struct: SetSystemCallOperation
// ----------------------------------------

func TestSetSystemCallOperation(t *testing.T) {
	o := NewSetSystemCallOperation()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeSetSystemCallOperation(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test call_id
	vb = &VariableBlob{}
	n, _, err = DeserializeSetSystemCallOperation(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test target
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeSetSystemCallOperation(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewSetSystemCallOperation()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: ContractCallOperation
// ----------------------------------------

func TestContractCallOperation(t *testing.T) {
	o := NewContractCallOperation()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeContractCallOperation(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test contract_id
	vb = &VariableBlob{}
	n, _, err = DeserializeContractCallOperation(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test entry_point
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeContractCallOperation(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test args
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeContractCallOperation(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewContractCallOperation()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Variant: Operation
// ----------------------------------------

func TestOperation(t *testing.T) {
	o := NewOperation()
	exerciseOperationSerialization(o, t)
	{
		v := NewOperation()
		v.Value = NewReservedOperation()
		exerciseOperationSerialization(v, t)

	}
	{
		v := NewOperation()
		v.Value = NewNopOperation()
		exerciseOperationSerialization(v, t)

	}
	{
		v := NewOperation()
		v.Value = NewCreateSystemContractOperation()
		exerciseOperationSerialization(v, t)

		vb := VariableBlob{2}
		n, _, err := DeserializeOperation(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := NewOperation()
		v.Value = NewContractCallOperation()
		exerciseOperationSerialization(v, t)

		vb := VariableBlob{3}
		n, _, err := DeserializeOperation(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := NewOperation()
		v.Value = NewSetSystemCallOperation()
		exerciseOperationSerialization(v, t)

		vb := VariableBlob{4}
		n, _, err := DeserializeOperation(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}

	// Test bad variant tag
	vb := VariableBlob{0x80}
	n, _, err := DeserializeOperation(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = VariableBlob{5}
	n, _, err = DeserializeOperation(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = NewOperation()
	if jerr := json.Unmarshal([]byte("\"!@#$%^&*\""), o); jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	// Test bad tag
	if jerr := json.Unmarshal([]byte("{\"type\":\"foobar\", \"value\":0}"), o); jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	// Test panic when serializing an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := Operation{Value: int64(0)}
		vb := NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := Operation{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseOperationSerialization(v *Operation, t *testing.T) {
	vb := NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := DeserializeOperation(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := NewOperation()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Struct: ActiveTransactionData
// ----------------------------------------

func TestActiveTransactionData(t *testing.T) {
	o := NewActiveTransactionData()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeActiveTransactionData(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewActiveTransactionData()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: PassiveTransactionData
// ----------------------------------------

func TestPassiveTransactionData(t *testing.T) {
	o := NewPassiveTransactionData()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializePassiveTransactionData(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewPassiveTransactionData()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: Transaction
// ----------------------------------------

func TestTransaction(t *testing.T) {
	o := NewTransaction()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeTransaction(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test active_data
	vb = &VariableBlob{}
	n, _, err = DeserializeTransaction(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test passive_data
	vb = &VariableBlob{0x00}
	n, _, err = DeserializeTransaction(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test signature_data
	vb = &VariableBlob{0x00, 0x00}
	n, _, err = DeserializeTransaction(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test operations
	vb = &VariableBlob{0x00, 0x00, 0x00}
	n, _, err = DeserializeTransaction(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewTransaction()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Enum: HeaderHashIndex
// ----------------------------------------

func TestHeaderHashIndex(t *testing.T) {
	vals := []HeaderHashIndex{
		HeaderHashIndexPreviousBlockHashIndex,
		HeaderHashIndexTransactionMerkleRootHashIndex,
		HeaderHashIndexPassiveDataMerkleRootHashIndex,
		HeaderHashIndexNumHeaderHashes,
	}

	// Make sure all types properly serialize
	for _, x := range vals {
		vb := NewVariableBlob()
		vb = x.Serialize(vb)

		nvb := NewVariableBlob()
		ox := UInt32(x)
		nvb = ox.Serialize(nvb)

		if !bytes.Equal(*vb, *nvb) {
			t.Errorf("Serialized enum does match ideal serialization.")
		}

		_, y, err := DeserializeHeaderHashIndex(vb)
		if err != nil {
			t.Error(err)
		}

		if x != *y {
			t.Errorf("Deserialized enum does not match original.")
		}

		// Test JSON
		jx, err := json.Marshal(x)
		if err != nil {
			t.Error(err)
		}

		jox, err := json.Marshal(ox)
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(jx, jox) {
			t.Error(err)
		}

		r := NewHeaderHashIndex()
		if err = json.Unmarshal(jx, r); err != nil {
			t.Error(err)
		}
	}

	// Find a value that is NOT an enum value
	w := getInvalidHeaderHashIndex()

	tw := UInt32(w)
	vb := NewVariableBlob()
	n, _, err := DeserializeHeaderHashIndex(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	vb = tw.Serialize(vb)

	if _, _, err := DeserializeHeaderHashIndex(vb); err == nil {
		t.Errorf("Deserializing an invalid value did not return an error.")
	}

	je := NewHeaderHashIndex()
	if err := json.Unmarshal([]byte(fmt.Sprint(tw)), &je); err == nil {
		t.Errorf("Deserializing an invalid JSON value did not return an error.")
	}

	if err = json.Unmarshal([]byte("\"foobar\""), &je); err == nil {
		t.Errorf("err == nil")
	}
}

func TestHeaderHashIndexPanic(t *testing.T) {
	// Find a value that is NOT an enum value
	w := getInvalidHeaderHashIndex()

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an invalid enum value did not panic.")
			}
		}()

		vb := NewVariableBlob()
		vb = w.Serialize(vb)
	}()

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an invalid enum value did not panic.")
			}
		}()

		_, _ = json.Marshal(w)
	}()
}

func getInvalidHeaderHashIndex() HeaderHashIndex {
	w := HeaderHashIndexNumHeaderHashes
	for IsValidHeaderHashIndex(w) {
		w++
	}

	return w
}
// ----------------------------------------
//  Struct: ActiveBlockData
// ----------------------------------------

func TestActiveBlockData(t *testing.T) {
	o := NewActiveBlockData()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeActiveBlockData(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test header_hashes
	vb = &VariableBlob{}
	n, _, err = DeserializeActiveBlockData(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test height
	vb = &VariableBlob{0x00, 0x00, 0x00}
	n, _, err = DeserializeActiveBlockData(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test timestamp
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeActiveBlockData(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewActiveBlockData()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: PassiveBlockData
// ----------------------------------------

func TestPassiveBlockData(t *testing.T) {
	o := NewPassiveBlockData()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializePassiveBlockData(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewPassiveBlockData()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}



// ----------------------------------------
//  Struct: Block
// ----------------------------------------

func TestBlock(t *testing.T) {
	o := NewBlock()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeBlock(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test active_data
	vb = &VariableBlob{}
	n, _, err = DeserializeBlock(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test passive_data
	vb = &VariableBlob{0x00}
	n, _, err = DeserializeBlock(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test signature_data
	vb = &VariableBlob{0x00, 0x00}
	n, _, err = DeserializeBlock(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test transactions
	vb = &VariableBlob{0x00, 0x00, 0x00}
	n, _, err = DeserializeBlock(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewBlock()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: ReservedQueryParams
// ----------------------------------------

func TestReservedQueryParams(t *testing.T) {
	o := NewReservedQueryParams()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeReservedQueryParams(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewReservedQueryParams()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: GetHeadInfoParams
// ----------------------------------------

func TestGetHeadInfoParams(t *testing.T) {
	o := NewGetHeadInfoParams()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeGetHeadInfoParams(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewGetHeadInfoParams()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Variant: QueryParamItem
// ----------------------------------------

func TestQueryParamItem(t *testing.T) {
	o := NewQueryParamItem()
	exerciseQueryParamItemSerialization(o, t)
	{
		v := NewQueryParamItem()
		v.Value = NewReservedQueryParams()
		exerciseQueryParamItemSerialization(v, t)

	}
	{
		v := NewQueryParamItem()
		v.Value = NewGetHeadInfoParams()
		exerciseQueryParamItemSerialization(v, t)

	}

	// Test bad variant tag
	vb := VariableBlob{0x80}
	n, _, err := DeserializeQueryParamItem(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = VariableBlob{2}
	n, _, err = DeserializeQueryParamItem(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = NewQueryParamItem()
	if jerr := json.Unmarshal([]byte("\"!@#$%^&*\""), o); jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	// Test bad tag
	if jerr := json.Unmarshal([]byte("{\"type\":\"foobar\", \"value\":0}"), o); jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	// Test panic when serializing an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := QueryParamItem{Value: int64(0)}
		vb := NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := QueryParamItem{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseQueryParamItemSerialization(v *QueryParamItem, t *testing.T) {
	vb := NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := DeserializeQueryParamItem(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := NewQueryParamItem()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Typedef: QuerySubmission
// ----------------------------------------

func TestQuerySubmission(t *testing.T) {
	o := NewQuerySubmission()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeQuerySubmission(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeQuerySubmission(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewQuerySubmission()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: ReservedQueryResult
// ----------------------------------------

func TestReservedQueryResult(t *testing.T) {
	o := NewReservedQueryResult()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeReservedQueryResult(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewReservedQueryResult()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: QueryError
// ----------------------------------------

func TestQueryError(t *testing.T) {
	o := NewQueryError()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeQueryError(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test error_text
	vb = &VariableBlob{}
	n, _, err = DeserializeQueryError(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewQueryError()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: HeadInfo
// ----------------------------------------

func TestHeadInfo(t *testing.T) {
	o := NewHeadInfo()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeHeadInfo(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test id
	vb = &VariableBlob{}
	n, _, err = DeserializeHeadInfo(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test height
	vb = &VariableBlob{0x00, 0x00}
	n, _, err = DeserializeHeadInfo(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewHeadInfo()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: GetHeadInfoResult
// ----------------------------------------

func TestGetHeadInfoResult(t *testing.T) {
	o := NewGetHeadInfoResult()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeGetHeadInfoResult(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeGetHeadInfoResult(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewGetHeadInfoResult()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Variant: QueryItemResult
// ----------------------------------------

func TestQueryItemResult(t *testing.T) {
	o := NewQueryItemResult()
	exerciseQueryItemResultSerialization(o, t)
	{
		v := NewQueryItemResult()
		v.Value = NewReservedQueryResult()
		exerciseQueryItemResultSerialization(v, t)

	}
	{
		v := NewQueryItemResult()
		v.Value = NewQueryError()
		exerciseQueryItemResultSerialization(v, t)

		vb := VariableBlob{1}
		n, _, err := DeserializeQueryItemResult(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := NewQueryItemResult()
		v.Value = NewGetHeadInfoResult()
		exerciseQueryItemResultSerialization(v, t)

		vb := VariableBlob{2}
		n, _, err := DeserializeQueryItemResult(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}

	// Test bad variant tag
	vb := VariableBlob{0x80}
	n, _, err := DeserializeQueryItemResult(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = VariableBlob{3}
	n, _, err = DeserializeQueryItemResult(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = NewQueryItemResult()
	if jerr := json.Unmarshal([]byte("\"!@#$%^&*\""), o); jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	// Test bad tag
	if jerr := json.Unmarshal([]byte("{\"type\":\"foobar\", \"value\":0}"), o); jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	// Test panic when serializing an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := QueryItemResult{Value: int64(0)}
		vb := NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := QueryItemResult{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseQueryItemResultSerialization(v *QueryItemResult, t *testing.T) {
	vb := NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := DeserializeQueryItemResult(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := NewQueryItemResult()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Typedef: QuerySubmissionResult
// ----------------------------------------

func TestQuerySubmissionResult(t *testing.T) {
	o := NewQuerySubmissionResult()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeQuerySubmissionResult(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeQuerySubmissionResult(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewQuerySubmissionResult()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: BlockTopology
// ----------------------------------------

func TestBlockTopology(t *testing.T) {
	o := NewBlockTopology()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeBlockTopology(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test id
	vb = &VariableBlob{}
	n, _, err = DeserializeBlockTopology(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test height
	vb = &VariableBlob{0x00, 0x00}
	n, _, err = DeserializeBlockTopology(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test previous
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeBlockTopology(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewBlockTopology()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: ReservedSubmission
// ----------------------------------------

func TestReservedSubmission(t *testing.T) {
	o := NewReservedSubmission()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeReservedSubmission(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewReservedSubmission()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: BlockSubmission
// ----------------------------------------

func TestBlockSubmission(t *testing.T) {
	o := NewBlockSubmission()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeBlockSubmission(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test topology
	vb = &VariableBlob{}
	n, _, err = DeserializeBlockSubmission(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test block
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeBlockSubmission(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test verify_passive_data
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeBlockSubmission(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test verify_block_signature
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeBlockSubmission(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test verify_transaction_signatures
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeBlockSubmission(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewBlockSubmission()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: TransactionSubmission
// ----------------------------------------

func TestTransactionSubmission(t *testing.T) {
	o := NewTransactionSubmission()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeTransactionSubmission(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test active_bytes
	vb = &VariableBlob{}
	n, _, err = DeserializeTransactionSubmission(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test passive_bytes
	vb = &VariableBlob{0x00}
	n, _, err = DeserializeTransactionSubmission(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewTransactionSubmission()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Variant: SubmissionItem
// ----------------------------------------

func TestSubmissionItem(t *testing.T) {
	o := NewSubmissionItem()
	exerciseSubmissionItemSerialization(o, t)
	{
		v := NewSubmissionItem()
		v.Value = NewReservedSubmission()
		exerciseSubmissionItemSerialization(v, t)

	}
	{
		v := NewSubmissionItem()
		v.Value = NewBlockSubmission()
		exerciseSubmissionItemSerialization(v, t)

		vb := VariableBlob{1}
		n, _, err := DeserializeSubmissionItem(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := NewSubmissionItem()
		v.Value = NewTransactionSubmission()
		exerciseSubmissionItemSerialization(v, t)

		vb := VariableBlob{2}
		n, _, err := DeserializeSubmissionItem(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := NewSubmissionItem()
		v.Value = NewQuerySubmission()
		exerciseSubmissionItemSerialization(v, t)

		vb := VariableBlob{3}
		n, _, err := DeserializeSubmissionItem(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}

	// Test bad variant tag
	vb := VariableBlob{0x80}
	n, _, err := DeserializeSubmissionItem(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = VariableBlob{4}
	n, _, err = DeserializeSubmissionItem(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = NewSubmissionItem()
	if jerr := json.Unmarshal([]byte("\"!@#$%^&*\""), o); jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	// Test bad tag
	if jerr := json.Unmarshal([]byte("{\"type\":\"foobar\", \"value\":0}"), o); jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	// Test panic when serializing an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := SubmissionItem{Value: int64(0)}
		vb := NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := SubmissionItem{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseSubmissionItemSerialization(v *SubmissionItem, t *testing.T) {
	vb := NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := DeserializeSubmissionItem(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := NewSubmissionItem()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Struct: ReservedSubmissionResult
// ----------------------------------------

func TestReservedSubmissionResult(t *testing.T) {
	o := NewReservedSubmissionResult()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeReservedSubmissionResult(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewReservedSubmissionResult()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: BlockSubmissionResult
// ----------------------------------------

func TestBlockSubmissionResult(t *testing.T) {
	o := NewBlockSubmissionResult()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeBlockSubmissionResult(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewBlockSubmissionResult()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: TransactionSubmissionResult
// ----------------------------------------

func TestTransactionSubmissionResult(t *testing.T) {
	o := NewTransactionSubmissionResult()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeTransactionSubmissionResult(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewTransactionSubmissionResult()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: SubmissionErrorResult
// ----------------------------------------

func TestSubmissionErrorResult(t *testing.T) {
	o := NewSubmissionErrorResult()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeSubmissionErrorResult(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test error_text
	vb = &VariableBlob{}
	n, _, err = DeserializeSubmissionErrorResult(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewSubmissionErrorResult()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Variant: SubmissionResult
// ----------------------------------------

func TestSubmissionResult(t *testing.T) {
	o := NewSubmissionResult()
	exerciseSubmissionResultSerialization(o, t)
	{
		v := NewSubmissionResult()
		v.Value = NewReservedSubmissionResult()
		exerciseSubmissionResultSerialization(v, t)

	}
	{
		v := NewSubmissionResult()
		v.Value = NewBlockSubmissionResult()
		exerciseSubmissionResultSerialization(v, t)

	}
	{
		v := NewSubmissionResult()
		v.Value = NewTransactionSubmissionResult()
		exerciseSubmissionResultSerialization(v, t)

	}
	{
		v := NewSubmissionResult()
		v.Value = NewQuerySubmissionResult()
		exerciseSubmissionResultSerialization(v, t)

		vb := VariableBlob{3}
		n, _, err := DeserializeSubmissionResult(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := NewSubmissionResult()
		v.Value = NewSubmissionErrorResult()
		exerciseSubmissionResultSerialization(v, t)

		vb := VariableBlob{4}
		n, _, err := DeserializeSubmissionResult(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}

	// Test bad variant tag
	vb := VariableBlob{0x80}
	n, _, err := DeserializeSubmissionResult(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = VariableBlob{5}
	n, _, err = DeserializeSubmissionResult(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = NewSubmissionResult()
	if jerr := json.Unmarshal([]byte("\"!@#$%^&*\""), o); jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	// Test bad tag
	if jerr := json.Unmarshal([]byte("{\"type\":\"foobar\", \"value\":0}"), o); jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	// Test panic when serializing an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := SubmissionResult{Value: int64(0)}
		vb := NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := SubmissionResult{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseSubmissionResultSerialization(v *SubmissionResult, t *testing.T) {
	vb := NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := DeserializeSubmissionResult(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := NewSubmissionResult()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Struct: VoidType
// ----------------------------------------

func TestVoidType(t *testing.T) {
	o := NewVoidType()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeVoidType(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewVoidType()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: PrintsArgs
// ----------------------------------------

func TestPrintsArgs(t *testing.T) {
	o := NewPrintsArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializePrintsArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test message
	vb = &VariableBlob{}
	n, _, err = DeserializePrintsArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewPrintsArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: PrintsRet
// ----------------------------------------

func TestPrintsRet(t *testing.T) {
	o := NewPrintsRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializePrintsRet(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewPrintsRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: VerifyBlockSigArgs
// ----------------------------------------

func TestVerifyBlockSigArgs(t *testing.T) {
	o := NewVerifyBlockSigArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeVerifyBlockSigArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test sig_data
	vb = &VariableBlob{}
	n, _, err = DeserializeVerifyBlockSigArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test digest
	vb = &VariableBlob{0x00}
	n, _, err = DeserializeVerifyBlockSigArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewVerifyBlockSigArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: VerifyBlockSigRet
// ----------------------------------------

func TestVerifyBlockSigRet(t *testing.T) {
	o := NewVerifyBlockSigRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeVerifyBlockSigRet(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeVerifyBlockSigRet(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewVerifyBlockSigRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: VerifyMerkleRootArgs
// ----------------------------------------

func TestVerifyMerkleRootArgs(t *testing.T) {
	o := NewVerifyMerkleRootArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeVerifyMerkleRootArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test root
	vb = &VariableBlob{}
	n, _, err = DeserializeVerifyMerkleRootArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test hashes
	vb = &VariableBlob{0x00, 0x00}
	n, _, err = DeserializeVerifyMerkleRootArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewVerifyMerkleRootArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: VerifyMerkleRootRet
// ----------------------------------------

func TestVerifyMerkleRootRet(t *testing.T) {
	o := NewVerifyMerkleRootRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeVerifyMerkleRootRet(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeVerifyMerkleRootRet(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewVerifyMerkleRootRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: ApplyBlockArgs
// ----------------------------------------

func TestApplyBlockArgs(t *testing.T) {
	o := NewApplyBlockArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeApplyBlockArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test block
	vb = &VariableBlob{}
	n, _, err = DeserializeApplyBlockArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test enable_check_passive_data
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeApplyBlockArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test enable_check_block_signature
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeApplyBlockArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test enable_check_transaction_signatures
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeApplyBlockArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewApplyBlockArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: ApplyBlockRet
// ----------------------------------------

func TestApplyBlockRet(t *testing.T) {
	o := NewApplyBlockRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeApplyBlockRet(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewApplyBlockRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: ApplyTransactionArgs
// ----------------------------------------

func TestApplyTransactionArgs(t *testing.T) {
	o := NewApplyTransactionArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeApplyTransactionArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test trx
	vb = &VariableBlob{}
	n, _, err = DeserializeApplyTransactionArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewApplyTransactionArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: ApplyTransactionRet
// ----------------------------------------

func TestApplyTransactionRet(t *testing.T) {
	o := NewApplyTransactionRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeApplyTransactionRet(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewApplyTransactionRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: ApplyUploadContractOperationArgs
// ----------------------------------------

func TestApplyUploadContractOperationArgs(t *testing.T) {
	o := NewApplyUploadContractOperationArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeApplyUploadContractOperationArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test op
	vb = &VariableBlob{}
	n, _, err = DeserializeApplyUploadContractOperationArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewApplyUploadContractOperationArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: ApplyUploadContractOperationRet
// ----------------------------------------

func TestApplyUploadContractOperationRet(t *testing.T) {
	o := NewApplyUploadContractOperationRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeApplyUploadContractOperationRet(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewApplyUploadContractOperationRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: ApplyReservedOperationArgs
// ----------------------------------------

func TestApplyReservedOperationArgs(t *testing.T) {
	o := NewApplyReservedOperationArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeApplyReservedOperationArgs(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewApplyReservedOperationArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: ApplyReservedOperationRet
// ----------------------------------------

func TestApplyReservedOperationRet(t *testing.T) {
	o := NewApplyReservedOperationRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeApplyReservedOperationRet(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewApplyReservedOperationRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: ApplyExecuteContractOperationArgs
// ----------------------------------------

func TestApplyExecuteContractOperationArgs(t *testing.T) {
	o := NewApplyExecuteContractOperationArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeApplyExecuteContractOperationArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test op
	vb = &VariableBlob{}
	n, _, err = DeserializeApplyExecuteContractOperationArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewApplyExecuteContractOperationArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: ApplyExecuteContractOperationRet
// ----------------------------------------

func TestApplyExecuteContractOperationRet(t *testing.T) {
	o := NewApplyExecuteContractOperationRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeApplyExecuteContractOperationRet(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewApplyExecuteContractOperationRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: ApplySetSystemCallOperationArgs
// ----------------------------------------

func TestApplySetSystemCallOperationArgs(t *testing.T) {
	o := NewApplySetSystemCallOperationArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeApplySetSystemCallOperationArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test op
	vb = &VariableBlob{}
	n, _, err = DeserializeApplySetSystemCallOperationArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewApplySetSystemCallOperationArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: ApplySetSystemCallOperationRet
// ----------------------------------------

func TestApplySetSystemCallOperationRet(t *testing.T) {
	o := NewApplySetSystemCallOperationRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeApplySetSystemCallOperationRet(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewApplySetSystemCallOperationRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: DbPutObjectArgs
// ----------------------------------------

func TestDbPutObjectArgs(t *testing.T) {
	o := NewDbPutObjectArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeDbPutObjectArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test space
	vb = &VariableBlob{}
	n, _, err = DeserializeDbPutObjectArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test key
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeDbPutObjectArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test obj
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeDbPutObjectArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewDbPutObjectArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: DbPutObjectRet
// ----------------------------------------

func TestDbPutObjectRet(t *testing.T) {
	o := NewDbPutObjectRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeDbPutObjectRet(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeDbPutObjectRet(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewDbPutObjectRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: DbGetObjectArgs
// ----------------------------------------

func TestDbGetObjectArgs(t *testing.T) {
	o := NewDbGetObjectArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeDbGetObjectArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test space
	vb = &VariableBlob{}
	n, _, err = DeserializeDbGetObjectArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test key
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeDbGetObjectArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test object_size_hint
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeDbGetObjectArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewDbGetObjectArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: DbGetObjectRet
// ----------------------------------------

func TestDbGetObjectRet(t *testing.T) {
	o := NewDbGetObjectRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeDbGetObjectRet(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeDbGetObjectRet(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewDbGetObjectRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: DbGetNextObjectArgs
// ----------------------------------------

func TestDbGetNextObjectArgs(t *testing.T) {
	o := NewDbGetNextObjectArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeDbGetNextObjectArgs(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeDbGetNextObjectArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewDbGetNextObjectArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: DbGetNextObjectRet
// ----------------------------------------

func TestDbGetNextObjectRet(t *testing.T) {
	o := NewDbGetNextObjectRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeDbGetNextObjectRet(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeDbGetNextObjectRet(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewDbGetNextObjectRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: DbGetPrevObjectArgs
// ----------------------------------------

func TestDbGetPrevObjectArgs(t *testing.T) {
	o := NewDbGetPrevObjectArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeDbGetPrevObjectArgs(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeDbGetPrevObjectArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewDbGetPrevObjectArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: DbGetPrevObjectRet
// ----------------------------------------

func TestDbGetPrevObjectRet(t *testing.T) {
	o := NewDbGetPrevObjectRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeDbGetPrevObjectRet(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeDbGetPrevObjectRet(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewDbGetPrevObjectRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: ExecuteContractArgs
// ----------------------------------------

func TestExecuteContractArgs(t *testing.T) {
	o := NewExecuteContractArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeExecuteContractArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test contract_id
	vb = &VariableBlob{}
	n, _, err = DeserializeExecuteContractArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test entry_point
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeExecuteContractArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test args
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeExecuteContractArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewExecuteContractArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: ExecuteContractRet
// ----------------------------------------

func TestExecuteContractRet(t *testing.T) {
	o := NewExecuteContractRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeExecuteContractRet(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeExecuteContractRet(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewExecuteContractRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: GetContractArgsSizeArgs
// ----------------------------------------

func TestGetContractArgsSizeArgs(t *testing.T) {
	o := NewGetContractArgsSizeArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeGetContractArgsSizeArgs(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewGetContractArgsSizeArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: GetContractArgsSizeRet
// ----------------------------------------

func TestGetContractArgsSizeRet(t *testing.T) {
	o := NewGetContractArgsSizeRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeGetContractArgsSizeRet(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeGetContractArgsSizeRet(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewGetContractArgsSizeRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: GetContractArgsArgs
// ----------------------------------------

func TestGetContractArgsArgs(t *testing.T) {
	o := NewGetContractArgsArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeGetContractArgsArgs(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewGetContractArgsArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: GetContractArgsRet
// ----------------------------------------

func TestGetContractArgsRet(t *testing.T) {
	o := NewGetContractArgsRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeGetContractArgsRet(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeGetContractArgsRet(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewGetContractArgsRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: SetContractReturnArgs
// ----------------------------------------

func TestSetContractReturnArgs(t *testing.T) {
	o := NewSetContractReturnArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeSetContractReturnArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test ret
	vb = &VariableBlob{}
	n, _, err = DeserializeSetContractReturnArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewSetContractReturnArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: SetContractReturnRet
// ----------------------------------------

func TestSetContractReturnRet(t *testing.T) {
	o := NewSetContractReturnRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeSetContractReturnRet(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewSetContractReturnRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: ExitContractArgs
// ----------------------------------------

func TestExitContractArgs(t *testing.T) {
	o := NewExitContractArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeExitContractArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test exit_code
	vb = &VariableBlob{}
	n, _, err = DeserializeExitContractArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewExitContractArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: ExitContractRet
// ----------------------------------------

func TestExitContractRet(t *testing.T) {
	o := NewExitContractRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeExitContractRet(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewExitContractRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: GetHeadInfoArgs
// ----------------------------------------

func TestGetHeadInfoArgs(t *testing.T) {
	o := NewGetHeadInfoArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeGetHeadInfoArgs(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewGetHeadInfoArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: GetHeadInfoRet
// ----------------------------------------

func TestGetHeadInfoRet(t *testing.T) {
	o := NewGetHeadInfoRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeGetHeadInfoRet(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeGetHeadInfoRet(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewGetHeadInfoRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Struct: HashArgs
// ----------------------------------------

func TestHashArgs(t *testing.T) {
	o := NewHashArgs()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeHashArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test code
	vb = &VariableBlob{}
	n, _, err = DeserializeHashArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test obj
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeHashArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test size
	vb = &VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = DeserializeHashArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewHashArgs()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Typedef: HashRet
// ----------------------------------------

func TestHashRet(t *testing.T) {
	o := NewHashRet()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeHashRet(vb)
	if err != nil {
		t.Error(err)
	}

	vb = NewVariableBlob()
	size, _, err := DeserializeHashRet(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewHashRet()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}

// ----------------------------------------
//  Enum: SystemCallID
// ----------------------------------------

func TestSystemCallID(t *testing.T) {
	vals := []SystemCallID{
		SystemCallIDPrints,
		SystemCallIDVerifyBlockHeader,
		SystemCallIDApplyBlock,
		SystemCallIDApplyTransaction,
		SystemCallIDApplyReservedOperation,
		SystemCallIDApplyUploadContractOperation,
		SystemCallIDApplyExecuteContractOperation,
		SystemCallIDApplySetSystemCallOperation,
		SystemCallIDDbPutObject,
		SystemCallIDDbGetObject,
		SystemCallIDDbGetNextObject,
		SystemCallIDDbGetPrevObject,
		SystemCallIDExecuteContract,
		SystemCallIDGetContractArgsSize,
		SystemCallIDGetContractArgs,
		SystemCallIDSetContractReturn,
		SystemCallIDExitContract,
		SystemCallIDGetHeadInfo,
		SystemCallIDHash,
		SystemCallIDVerifyBlockSig,
		SystemCallIDVerifyMerkleRoot,
	}

	// Make sure all types properly serialize
	for _, x := range vals {
		vb := NewVariableBlob()
		vb = x.Serialize(vb)

		nvb := NewVariableBlob()
		ox := UInt32(x)
		nvb = ox.Serialize(nvb)

		if !bytes.Equal(*vb, *nvb) {
			t.Errorf("Serialized enum does match ideal serialization.")
		}

		_, y, err := DeserializeSystemCallID(vb)
		if err != nil {
			t.Error(err)
		}

		if x != *y {
			t.Errorf("Deserialized enum does not match original.")
		}

		// Test JSON
		jx, err := json.Marshal(x)
		if err != nil {
			t.Error(err)
		}

		jox, err := json.Marshal(ox)
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(jx, jox) {
			t.Error(err)
		}

		r := NewSystemCallID()
		if err = json.Unmarshal(jx, r); err != nil {
			t.Error(err)
		}
	}

	// Find a value that is NOT an enum value
	w := getInvalidSystemCallID()

	tw := UInt32(w)
	vb := NewVariableBlob()
	n, _, err := DeserializeSystemCallID(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	vb = tw.Serialize(vb)

	if _, _, err := DeserializeSystemCallID(vb); err == nil {
		t.Errorf("Deserializing an invalid value did not return an error.")
	}

	je := NewSystemCallID()
	if err := json.Unmarshal([]byte(fmt.Sprint(tw)), &je); err == nil {
		t.Errorf("Deserializing an invalid JSON value did not return an error.")
	}

	if err = json.Unmarshal([]byte("\"foobar\""), &je); err == nil {
		t.Errorf("err == nil")
	}
}

func TestSystemCallIDPanic(t *testing.T) {
	// Find a value that is NOT an enum value
	w := getInvalidSystemCallID()

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an invalid enum value did not panic.")
			}
		}()

		vb := NewVariableBlob()
		vb = w.Serialize(vb)
	}()

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an invalid enum value did not panic.")
			}
		}()

		_, _ = json.Marshal(w)
	}()
}

func getInvalidSystemCallID() SystemCallID {
	w := SystemCallIDVerifyMerkleRoot
	for IsValidSystemCallID(w) {
		w++
	}

	return w
}
// ----------------------------------------
//  Struct: BlockPart
// ----------------------------------------

func TestBlockPart(t *testing.T) {
	o := NewBlockPart()

	vb := NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := DeserializeBlockPart(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test active_data
	vb = &VariableBlob{}
	n, _, err = DeserializeBlockPart(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test passive_data
	vb = &VariableBlob{0x00}
	n, _, err = DeserializeBlockPart(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test sig_data
	vb = &VariableBlob{0x00, 0x00}
	n, _, err = DeserializeBlockPart(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewBlockPart()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("[1,2,3,4,5]"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{1:2, 3:4}"), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}
}


// ----------------------------------------
//  FixedBlob20
// ----------------------------------------

func TestFixedBlob20(t *testing.T) {
	fb := NewFixedBlob20()
	for i := 0; i < 20; i++ {
		fb[i] = byte((20 + i) % 256)
	}

	vb := NewVariableBlob()
	vb = fb.Serialize(vb)

	size, nfb, err := DeserializeFixedBlob20(vb)
	if err != nil {
		t.Error(err)
	}

	if size != 20 {
		t.Errorf("Fixedblob deserialization consumed %d bytes. Expected %d bytes.", size, 20)
	}

	if !bytes.Equal(fb[:], nfb[:]) {
		t.Errorf("Binary deserialized fixed blob does not match source.")
	}

	vb = NewVariableBlob()
	size, _, err = DeserializeFixedBlob20(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	j, err := json.Marshal(fb)
	if err != nil {
		t.Error(err)
	}

	jfb := NewFixedBlob20()
	err = json.Unmarshal(j, &jfb)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(fb[:], jfb[:]) {
		t.Errorf("JSON deserialized fixed blob does not match source.")
	}

	err = json.Unmarshal([]byte("[1,2,3,4,5]"), &jfb)
	if err == nil {
		t.Errorf("Fixed blob unmarshaled from invalid JSON object.")
	}

	err = json.Unmarshal([]byte("\"gdtdgh5td\""), &jfb)
	if err == nil {
		t.Errorf("Fixed blob unmarshaled from invalid JSON object.")
	}

	// Encode a value that is incorrectly sized and ensure it returns an error
	wfb := make([]byte, 20+1)
	for i := 0; i < 20+1; i++ {
		wfb[i] = byte((20 + i) % 256)
	}
	s := EncodeBytes(wfb)
	err = json.Unmarshal([]byte("\""+s+"\""), &jfb)
	if err == nil {
		t.Errorf("deserialized JSON from a string which was not the proper size.")
	}

	// Encode with invalid start character of '0' and ensure it returns an error
	wfb = make([]byte, 20)
	for i := 0; i < 20; i++ {
		wfb[i] = byte((20 + i) % 256)
	}
	s = EncodeBytes(wfb)
	s = "0" + s[1:]
	err = json.Unmarshal([]byte("\""+s+"\""), &jfb)
	if err == nil {
		t.Errorf("Deserialized JSON from a string which contained an invalid encoding character.")
	}
}

// ----------------------------------------
//  FixedBlob65
// ----------------------------------------

func TestFixedBlob65(t *testing.T) {
	fb := NewFixedBlob65()
	for i := 0; i < 65; i++ {
		fb[i] = byte((65 + i) % 256)
	}

	vb := NewVariableBlob()
	vb = fb.Serialize(vb)

	size, nfb, err := DeserializeFixedBlob65(vb)
	if err != nil {
		t.Error(err)
	}

	if size != 65 {
		t.Errorf("Fixedblob deserialization consumed %d bytes. Expected %d bytes.", size, 65)
	}

	if !bytes.Equal(fb[:], nfb[:]) {
		t.Errorf("Binary deserialized fixed blob does not match source.")
	}

	vb = NewVariableBlob()
	size, _, err = DeserializeFixedBlob65(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if size != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	j, err := json.Marshal(fb)
	if err != nil {
		t.Error(err)
	}

	jfb := NewFixedBlob65()
	err = json.Unmarshal(j, &jfb)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(fb[:], jfb[:]) {
		t.Errorf("JSON deserialized fixed blob does not match source.")
	}

	err = json.Unmarshal([]byte("[1,2,3,4,5]"), &jfb)
	if err == nil {
		t.Errorf("Fixed blob unmarshaled from invalid JSON object.")
	}

	err = json.Unmarshal([]byte("\"gdtdgh5td\""), &jfb)
	if err == nil {
		t.Errorf("Fixed blob unmarshaled from invalid JSON object.")
	}

	// Encode a value that is incorrectly sized and ensure it returns an error
	wfb := make([]byte, 65+1)
	for i := 0; i < 65+1; i++ {
		wfb[i] = byte((65 + i) % 256)
	}
	s := EncodeBytes(wfb)
	err = json.Unmarshal([]byte("\""+s+"\""), &jfb)
	if err == nil {
		t.Errorf("deserialized JSON from a string which was not the proper size.")
	}

	// Encode with invalid start character of '0' and ensure it returns an error
	wfb = make([]byte, 65)
	for i := 0; i < 65; i++ {
		wfb[i] = byte((65 + i) % 256)
	}
	s = EncodeBytes(wfb)
	s = "0" + s[1:]
	err = json.Unmarshal([]byte("\""+s+"\""), &jfb)
	if err == nil {
		t.Errorf("Deserialized JSON from a string which contained an invalid encoding character.")
	}
}



// ----------------------------------------
//  OpaqueActiveBlockData
// ----------------------------------------

func TestOpaqueActiveBlockData(t *testing.T) {
	o := NewOpaqueActiveBlockData()

	o.Box()
	if o.IsUnboxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if o.IsUnboxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.MakeImmutable() // Call MakeImmutable() on Boxed
	if o.IsUnboxed() {
		t.Errorf("Unboxed -> MakeImmutable failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if !o.IsUnboxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err == nil {
		t.Errorf("Getting native on Unboxed should fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if !o.IsUnboxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.MakeImmutable() // Call MakeImmutable() on Unboxed
	if !o.IsUnboxed() {
		t.Errorf("Unboxed -> MakeImmutable failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Unboxed
	if !o.IsMutable() {
		t.Errorf("Unboxed -> Mutable failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Mutable
	if !o.IsMutable() {
		t.Errorf("Mutable -> Mutable failed.")
	}

	// Test getting native on Mutable
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Mutable
	if !o.IsMutable() {
		t.Errorf("Mutable -> Unboxed failed.")
	}

	o.Box() // Call Box() on Mutable
	if o.IsUnboxed() {
		t.Errorf("Mutable -> Boxed failed.")
	}

	o.MakeMutable()
	o.MakeImmutable() // Call MakeImmutable() on Mutable
	if !o.IsUnboxed() {
		t.Errorf("Mutable -> Immutable failed.")
	}

	o.Box() // Call Box() on Unboxed
	if o.IsUnboxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Boxed
	if !o.IsMutable() {
		t.Errorf("Boxed -> Mutable failed.")
	}

	// Test serialization

	vb := NewVariableBlob()
	vb = o.Serialize(vb)
	b := o.GetBlob()

	o.Box()
	if !bytes.Equal(*b, (*vb)[1:]) {
		t.Errorf("GetBlob and Serialization do not match")
	}

	o.Unbox()
	b = o.GetBlob()
	if !bytes.Equal(*b, (*vb)[1:]) {
		t.Errorf("GetBlob and Serialization do not match")
	}

	_, _, err2 := DeserializeOpaqueActiveBlockData(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewOpaqueActiveBlockData()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"koinos::types::protocol::active_block_data\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jo.Unbox()
	if jo.IsUnboxed() {
		t.Errorf("Unboxed incompatible serialization")
	}

	expected := []byte("{\"opaque\":{\"type\":\"koinos::types::protocol::active_block_data\",\"value\":\"zt1Zv2yaZ\"}}")
	v, jerr = json.Marshal(jo)
	if jerr != nil {
		t.Error(jerr)
	}
	if !bytes.Equal(v, expected) {
		t.Errorf("Marshal unknown data to json failed. Expected: %s, Was: %s", expected, v)
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"foobar\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":10}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	jerr = json.Unmarshal([]byte("{\"header_hashes\":0,\"height\":0,\"timestamp\":0}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}
}

// ----------------------------------------
//  OpaqueActiveTransactionData
// ----------------------------------------

func TestOpaqueActiveTransactionData(t *testing.T) {
	o := NewOpaqueActiveTransactionData()

	o.Box()
	if o.IsUnboxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if o.IsUnboxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.MakeImmutable() // Call MakeImmutable() on Boxed
	if o.IsUnboxed() {
		t.Errorf("Unboxed -> MakeImmutable failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if !o.IsUnboxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err == nil {
		t.Errorf("Getting native on Unboxed should fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if !o.IsUnboxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.MakeImmutable() // Call MakeImmutable() on Unboxed
	if !o.IsUnboxed() {
		t.Errorf("Unboxed -> MakeImmutable failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Unboxed
	if !o.IsMutable() {
		t.Errorf("Unboxed -> Mutable failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Mutable
	if !o.IsMutable() {
		t.Errorf("Mutable -> Mutable failed.")
	}

	// Test getting native on Mutable
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Mutable
	if !o.IsMutable() {
		t.Errorf("Mutable -> Unboxed failed.")
	}

	o.Box() // Call Box() on Mutable
	if o.IsUnboxed() {
		t.Errorf("Mutable -> Boxed failed.")
	}

	o.MakeMutable()
	o.MakeImmutable() // Call MakeImmutable() on Mutable
	if !o.IsUnboxed() {
		t.Errorf("Mutable -> Immutable failed.")
	}

	o.Box() // Call Box() on Unboxed
	if o.IsUnboxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Boxed
	if !o.IsMutable() {
		t.Errorf("Boxed -> Mutable failed.")
	}

	// Test serialization

	vb := NewVariableBlob()
	vb = o.Serialize(vb)
	b := o.GetBlob()

	o.Box()
	if !bytes.Equal(*b, (*vb)[1:]) {
		t.Errorf("GetBlob and Serialization do not match")
	}

	o.Unbox()
	b = o.GetBlob()
	if !bytes.Equal(*b, (*vb)[1:]) {
		t.Errorf("GetBlob and Serialization do not match")
	}

	_, _, err2 := DeserializeOpaqueActiveTransactionData(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewOpaqueActiveTransactionData()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"koinos::types::protocol::active_transaction_data\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jo.Unbox()
	if jo.IsUnboxed() {
		t.Errorf("Unboxed incompatible serialization")
	}

	expected := []byte("{\"opaque\":{\"type\":\"koinos::types::protocol::active_transaction_data\",\"value\":\"zt1Zv2yaZ\"}}")
	v, jerr = json.Marshal(jo)
	if jerr != nil {
		t.Error(jerr)
	}
	if !bytes.Equal(v, expected) {
		t.Errorf("Marshal unknown data to json failed. Expected: %s, Was: %s", expected, v)
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"foobar\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":10}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

}

// ----------------------------------------
//  OpaquePassiveBlockData
// ----------------------------------------

func TestOpaquePassiveBlockData(t *testing.T) {
	o := NewOpaquePassiveBlockData()

	o.Box()
	if o.IsUnboxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if o.IsUnboxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.MakeImmutable() // Call MakeImmutable() on Boxed
	if o.IsUnboxed() {
		t.Errorf("Unboxed -> MakeImmutable failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if !o.IsUnboxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err == nil {
		t.Errorf("Getting native on Unboxed should fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if !o.IsUnboxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.MakeImmutable() // Call MakeImmutable() on Unboxed
	if !o.IsUnboxed() {
		t.Errorf("Unboxed -> MakeImmutable failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Unboxed
	if !o.IsMutable() {
		t.Errorf("Unboxed -> Mutable failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Mutable
	if !o.IsMutable() {
		t.Errorf("Mutable -> Mutable failed.")
	}

	// Test getting native on Mutable
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Mutable
	if !o.IsMutable() {
		t.Errorf("Mutable -> Unboxed failed.")
	}

	o.Box() // Call Box() on Mutable
	if o.IsUnboxed() {
		t.Errorf("Mutable -> Boxed failed.")
	}

	o.MakeMutable()
	o.MakeImmutable() // Call MakeImmutable() on Mutable
	if !o.IsUnboxed() {
		t.Errorf("Mutable -> Immutable failed.")
	}

	o.Box() // Call Box() on Unboxed
	if o.IsUnboxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Boxed
	if !o.IsMutable() {
		t.Errorf("Boxed -> Mutable failed.")
	}

	// Test serialization

	vb := NewVariableBlob()
	vb = o.Serialize(vb)
	b := o.GetBlob()

	o.Box()
	if !bytes.Equal(*b, (*vb)[1:]) {
		t.Errorf("GetBlob and Serialization do not match")
	}

	o.Unbox()
	b = o.GetBlob()
	if !bytes.Equal(*b, (*vb)[1:]) {
		t.Errorf("GetBlob and Serialization do not match")
	}

	_, _, err2 := DeserializeOpaquePassiveBlockData(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewOpaquePassiveBlockData()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"koinos::types::protocol::passive_block_data\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jo.Unbox()
	if jo.IsUnboxed() {
		t.Errorf("Unboxed incompatible serialization")
	}

	expected := []byte("{\"opaque\":{\"type\":\"koinos::types::protocol::passive_block_data\",\"value\":\"zt1Zv2yaZ\"}}")
	v, jerr = json.Marshal(jo)
	if jerr != nil {
		t.Error(jerr)
	}
	if !bytes.Equal(v, expected) {
		t.Errorf("Marshal unknown data to json failed. Expected: %s, Was: %s", expected, v)
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"foobar\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":10}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

}

// ----------------------------------------
//  OpaquePassiveTransactionData
// ----------------------------------------

func TestOpaquePassiveTransactionData(t *testing.T) {
	o := NewOpaquePassiveTransactionData()

	o.Box()
	if o.IsUnboxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if o.IsUnboxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.MakeImmutable() // Call MakeImmutable() on Boxed
	if o.IsUnboxed() {
		t.Errorf("Unboxed -> MakeImmutable failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if !o.IsUnboxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err == nil {
		t.Errorf("Getting native on Unboxed should fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if !o.IsUnboxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.MakeImmutable() // Call MakeImmutable() on Unboxed
	if !o.IsUnboxed() {
		t.Errorf("Unboxed -> MakeImmutable failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Unboxed
	if !o.IsMutable() {
		t.Errorf("Unboxed -> Mutable failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Mutable
	if !o.IsMutable() {
		t.Errorf("Mutable -> Mutable failed.")
	}

	// Test getting native on Mutable
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Mutable
	if !o.IsMutable() {
		t.Errorf("Mutable -> Unboxed failed.")
	}

	o.Box() // Call Box() on Mutable
	if o.IsUnboxed() {
		t.Errorf("Mutable -> Boxed failed.")
	}

	o.MakeMutable()
	o.MakeImmutable() // Call MakeImmutable() on Mutable
	if !o.IsUnboxed() {
		t.Errorf("Mutable -> Immutable failed.")
	}

	o.Box() // Call Box() on Unboxed
	if o.IsUnboxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Boxed
	if !o.IsMutable() {
		t.Errorf("Boxed -> Mutable failed.")
	}

	// Test serialization

	vb := NewVariableBlob()
	vb = o.Serialize(vb)
	b := o.GetBlob()

	o.Box()
	if !bytes.Equal(*b, (*vb)[1:]) {
		t.Errorf("GetBlob and Serialization do not match")
	}

	o.Unbox()
	b = o.GetBlob()
	if !bytes.Equal(*b, (*vb)[1:]) {
		t.Errorf("GetBlob and Serialization do not match")
	}

	_, _, err2 := DeserializeOpaquePassiveTransactionData(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewOpaquePassiveTransactionData()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"koinos::types::protocol::passive_transaction_data\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jo.Unbox()
	if jo.IsUnboxed() {
		t.Errorf("Unboxed incompatible serialization")
	}

	expected := []byte("{\"opaque\":{\"type\":\"koinos::types::protocol::passive_transaction_data\",\"value\":\"zt1Zv2yaZ\"}}")
	v, jerr = json.Marshal(jo)
	if jerr != nil {
		t.Error(jerr)
	}
	if !bytes.Equal(v, expected) {
		t.Errorf("Marshal unknown data to json failed. Expected: %s, Was: %s", expected, v)
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"foobar\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":10}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

}

// ----------------------------------------
//  OpaqueQueryItemResult
// ----------------------------------------

func TestOpaqueQueryItemResult(t *testing.T) {
	o := NewOpaqueQueryItemResult()

	o.Box()
	if o.IsUnboxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if o.IsUnboxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.MakeImmutable() // Call MakeImmutable() on Boxed
	if o.IsUnboxed() {
		t.Errorf("Unboxed -> MakeImmutable failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if !o.IsUnboxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err == nil {
		t.Errorf("Getting native on Unboxed should fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if !o.IsUnboxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.MakeImmutable() // Call MakeImmutable() on Unboxed
	if !o.IsUnboxed() {
		t.Errorf("Unboxed -> MakeImmutable failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Unboxed
	if !o.IsMutable() {
		t.Errorf("Unboxed -> Mutable failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Mutable
	if !o.IsMutable() {
		t.Errorf("Mutable -> Mutable failed.")
	}

	// Test getting native on Mutable
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Mutable
	if !o.IsMutable() {
		t.Errorf("Mutable -> Unboxed failed.")
	}

	o.Box() // Call Box() on Mutable
	if o.IsUnboxed() {
		t.Errorf("Mutable -> Boxed failed.")
	}

	o.MakeMutable()
	o.MakeImmutable() // Call MakeImmutable() on Mutable
	if !o.IsUnboxed() {
		t.Errorf("Mutable -> Immutable failed.")
	}

	o.Box() // Call Box() on Unboxed
	if o.IsUnboxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Boxed
	if !o.IsMutable() {
		t.Errorf("Boxed -> Mutable failed.")
	}

	// Test serialization

	vb := NewVariableBlob()
	vb = o.Serialize(vb)
	b := o.GetBlob()

	o.Box()
	if !bytes.Equal(*b, (*vb)[1:]) {
		t.Errorf("GetBlob and Serialization do not match")
	}

	o.Unbox()
	b = o.GetBlob()
	if !bytes.Equal(*b, (*vb)[1:]) {
		t.Errorf("GetBlob and Serialization do not match")
	}

	_, _, err2 := DeserializeOpaqueQueryItemResult(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewOpaqueQueryItemResult()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"koinos::types::rpc::query_item_result\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jo.Unbox()
	if jo.IsUnboxed() {
		t.Errorf("Unboxed incompatible serialization")
	}

	expected := []byte("{\"opaque\":{\"type\":\"koinos::types::rpc::query_item_result\",\"value\":\"zt1Zv2yaZ\"}}")
	v, jerr = json.Marshal(jo)
	if jerr != nil {
		t.Error(jerr)
	}
	if !bytes.Equal(v, expected) {
		t.Errorf("Marshal unknown data to json failed. Expected: %s, Was: %s", expected, v)
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"foobar\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":10}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	jerr = json.Unmarshal([]byte("{}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}
}

// ----------------------------------------
//  OpaqueQueryParamItem
// ----------------------------------------

func TestOpaqueQueryParamItem(t *testing.T) {
	o := NewOpaqueQueryParamItem()

	o.Box()
	if o.IsUnboxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if o.IsUnboxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.MakeImmutable() // Call MakeImmutable() on Boxed
	if o.IsUnboxed() {
		t.Errorf("Unboxed -> MakeImmutable failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if !o.IsUnboxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err == nil {
		t.Errorf("Getting native on Unboxed should fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if !o.IsUnboxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.MakeImmutable() // Call MakeImmutable() on Unboxed
	if !o.IsUnboxed() {
		t.Errorf("Unboxed -> MakeImmutable failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Unboxed
	if !o.IsMutable() {
		t.Errorf("Unboxed -> Mutable failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Mutable
	if !o.IsMutable() {
		t.Errorf("Mutable -> Mutable failed.")
	}

	// Test getting native on Mutable
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Mutable
	if !o.IsMutable() {
		t.Errorf("Mutable -> Unboxed failed.")
	}

	o.Box() // Call Box() on Mutable
	if o.IsUnboxed() {
		t.Errorf("Mutable -> Boxed failed.")
	}

	o.MakeMutable()
	o.MakeImmutable() // Call MakeImmutable() on Mutable
	if !o.IsUnboxed() {
		t.Errorf("Mutable -> Immutable failed.")
	}

	o.Box() // Call Box() on Unboxed
	if o.IsUnboxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Boxed
	if !o.IsMutable() {
		t.Errorf("Boxed -> Mutable failed.")
	}

	// Test serialization

	vb := NewVariableBlob()
	vb = o.Serialize(vb)
	b := o.GetBlob()

	o.Box()
	if !bytes.Equal(*b, (*vb)[1:]) {
		t.Errorf("GetBlob and Serialization do not match")
	}

	o.Unbox()
	b = o.GetBlob()
	if !bytes.Equal(*b, (*vb)[1:]) {
		t.Errorf("GetBlob and Serialization do not match")
	}

	_, _, err2 := DeserializeOpaqueQueryParamItem(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewOpaqueQueryParamItem()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"koinos::types::rpc::query_param_item\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jo.Unbox()
	if jo.IsUnboxed() {
		t.Errorf("Unboxed incompatible serialization")
	}

	expected := []byte("{\"opaque\":{\"type\":\"koinos::types::rpc::query_param_item\",\"value\":\"zt1Zv2yaZ\"}}")
	v, jerr = json.Marshal(jo)
	if jerr != nil {
		t.Error(jerr)
	}
	if !bytes.Equal(v, expected) {
		t.Errorf("Marshal unknown data to json failed. Expected: %s, Was: %s", expected, v)
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"foobar\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":10}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	jerr = json.Unmarshal([]byte("{}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}
}

// ----------------------------------------
//  OpaqueTransaction
// ----------------------------------------

func TestOpaqueTransaction(t *testing.T) {
	o := NewOpaqueTransaction()

	o.Box()
	if o.IsUnboxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if o.IsUnboxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.MakeImmutable() // Call MakeImmutable() on Boxed
	if o.IsUnboxed() {
		t.Errorf("Unboxed -> MakeImmutable failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if !o.IsUnboxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err == nil {
		t.Errorf("Getting native on Unboxed should fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if !o.IsUnboxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.MakeImmutable() // Call MakeImmutable() on Unboxed
	if !o.IsUnboxed() {
		t.Errorf("Unboxed -> MakeImmutable failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Unboxed
	if !o.IsMutable() {
		t.Errorf("Unboxed -> Mutable failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Mutable
	if !o.IsMutable() {
		t.Errorf("Mutable -> Mutable failed.")
	}

	// Test getting native on Mutable
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Mutable
	if !o.IsMutable() {
		t.Errorf("Mutable -> Unboxed failed.")
	}

	o.Box() // Call Box() on Mutable
	if o.IsUnboxed() {
		t.Errorf("Mutable -> Boxed failed.")
	}

	o.MakeMutable()
	o.MakeImmutable() // Call MakeImmutable() on Mutable
	if !o.IsUnboxed() {
		t.Errorf("Mutable -> Immutable failed.")
	}

	o.Box() // Call Box() on Unboxed
	if o.IsUnboxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.MakeMutable() // Call MakeMutable() on Boxed
	if !o.IsMutable() {
		t.Errorf("Boxed -> Mutable failed.")
	}

	// Test serialization

	vb := NewVariableBlob()
	vb = o.Serialize(vb)
	b := o.GetBlob()

	o.Box()
	if !bytes.Equal(*b, (*vb)[1:]) {
		t.Errorf("GetBlob and Serialization do not match")
	}

	o.Unbox()
	b = o.GetBlob()
	if !bytes.Equal(*b, (*vb)[1:]) {
		t.Errorf("GetBlob and Serialization do not match")
	}

	_, _, err2 := DeserializeOpaqueTransaction(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := NewOpaqueTransaction()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"koinos::types::protocol::transaction\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jo.Unbox()
	if jo.IsUnboxed() {
		t.Errorf("Unboxed incompatible serialization")
	}

	expected := []byte("{\"opaque\":{\"type\":\"koinos::types::protocol::transaction\",\"value\":\"zt1Zv2yaZ\"}}")
	v, jerr = json.Marshal(jo)
	if jerr != nil {
		t.Error(jerr)
	}
	if !bytes.Equal(v, expected) {
		t.Errorf("Marshal unknown data to json failed. Expected: %s, Was: %s", expected, v)
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"foobar\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":10}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	jerr = json.Unmarshal([]byte("{\"active_data\":0,\"passive_data\":0,\"signature_data\":0,\"operations\":0}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}
}


// ----------------------------------------
//  VectorMultihash
// ----------------------------------------

func TestVectorMultihash(t *testing.T) {
	v := NewVectorMultihash()
	for i := 0; i < 16; i++ {
		no := NewMultihash()
		*v = append(*v, *no)
	}

	vb := NewVariableBlob()
	vb = v.Serialize(vb)

	_, nv, err := DeserializeVectorMultihash(vb)
	if err != nil {
		t.Error(err)
	}

	if len(*nv) != len(*v) {
		t.Errorf("Deserialized vector length does not match original.")
	}

	j, err := json.Marshal(v)
	if err != nil {
		t.Error(err)
	}

	jv := NewVectorMultihash()
	err = json.Unmarshal(j, &jv)
	if err != nil {
		t.Error(err)
	}

	// Test no data in the vector
	vb = &VariableBlob{0x01}
	n, _, err := DeserializeVectorMultihash(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}
}

// ----------------------------------------
//  VectorOpaqueTransaction
// ----------------------------------------

func TestVectorOpaqueTransaction(t *testing.T) {
	v := NewVectorOpaqueTransaction()
	for i := 0; i < 16; i++ {
		no := NewOpaqueTransaction()
		*v = append(*v, *no)
	}

	vb := NewVariableBlob()
	vb = v.Serialize(vb)

	_, nv, err := DeserializeVectorOpaqueTransaction(vb)
	if err != nil {
		t.Error(err)
	}

	if len(*nv) != len(*v) {
		t.Errorf("Deserialized vector length does not match original.")
	}

	j, err := json.Marshal(v)
	if err != nil {
		t.Error(err)
	}

	jv := NewVectorOpaqueTransaction()
	err = json.Unmarshal(j, &jv)
	if err != nil {
		t.Error(err)
	}

	// Test no data in the vector
	vb = &VariableBlob{0x01}
	n, _, err := DeserializeVectorOpaqueTransaction(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}
}

// ----------------------------------------
//  VectorOperation
// ----------------------------------------

func TestVectorOperation(t *testing.T) {
	v := NewVectorOperation()
	for i := 0; i < 16; i++ {
		no := NewOperation()
		*v = append(*v, *no)
	}

	vb := NewVariableBlob()
	vb = v.Serialize(vb)

	_, nv, err := DeserializeVectorOperation(vb)
	if err != nil {
		t.Error(err)
	}

	if len(*nv) != len(*v) {
		t.Errorf("Deserialized vector length does not match original.")
	}

	j, err := json.Marshal(v)
	if err != nil {
		t.Error(err)
	}

	jv := NewVectorOperation()
	err = json.Unmarshal(j, &jv)
	if err != nil {
		t.Error(err)
	}

	// Test no data in the vector
	vb = &VariableBlob{0x01}
	n, _, err := DeserializeVectorOperation(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}
}


