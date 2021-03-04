//   ____                           _           _   _____         _
//  / ___| ___ _ __   ___ _ __ __ _| |_ ___  __| | |_   _|__  ___| |_ ___
// | |  _ / _ \ '_ \ / _ \ '__/ _` | __/ _ \/ _` |   | |/ _ \/ __| __/ __|
// | |_| |  __/ | | |  __/ | | (_| | ||  __/ (_| |   | |  __/\__ \ |_\__ \
//  \____|\___|_| |_|\___|_|  \__,_|\__\___|\__,_|   |_|\___||___/\__|___/
//                         Please do not modify

package koinos_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/koinos/koinos-types-golang"
	"testing"
)

// ----------------------------------------
//  Struct: ActiveBlockData
// ----------------------------------------

func TestActiveBlockData(t *testing.T) {
	o := koinos.NewActiveBlockData()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeActiveBlockData(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test previous_block
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeActiveBlockData(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test transaction_merkle_root
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeActiveBlockData(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test passive_data_merkle_root
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeActiveBlockData(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test height
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeActiveBlockData(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test timestamp
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeActiveBlockData(vb)
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

	jo := koinos.NewActiveBlockData()
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
	o := koinos.NewPassiveBlockData()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializePassiveBlockData(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewPassiveBlockData()
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
//  Struct: ActiveTransactionData
// ----------------------------------------

func TestActiveTransactionData(t *testing.T) {
	o := koinos.NewActiveTransactionData()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeActiveTransactionData(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test resource_limit
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeActiveTransactionData(vb)
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

	jo := koinos.NewActiveTransactionData()
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
	o := koinos.NewPassiveTransactionData()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializePassiveTransactionData(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewPassiveTransactionData()
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
//  Struct: UnusedExtensionsType
// ----------------------------------------

func TestUnusedExtensionsType(t *testing.T) {
	o := koinos.NewUnusedExtensionsType()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeUnusedExtensionsType(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewUnusedExtensionsType()
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
//  Struct: ReservedOperation
// ----------------------------------------

func TestReservedOperation(t *testing.T) {
	o := koinos.NewReservedOperation()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeReservedOperation(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewReservedOperation()
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
	o := koinos.NewNopOperation()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeNopOperation(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewNopOperation()
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
//  Typedef: ContractIDType
// ----------------------------------------

func TestContractIDType(t *testing.T) {
	o := koinos.NewContractIDType()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeContractIDType(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeContractIDType(vb)
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

	jo := koinos.NewContractIDType()
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
//  Struct: CreateSystemContractOperation
// ----------------------------------------

func TestCreateSystemContractOperation(t *testing.T) {
	o := koinos.NewCreateSystemContractOperation()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeCreateSystemContractOperation(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test contract_id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeCreateSystemContractOperation(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test bytecode
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeCreateSystemContractOperation(vb)
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

	jo := koinos.NewCreateSystemContractOperation()
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
	o := koinos.NewContractCallOperation()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeContractCallOperation(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test contract_id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeContractCallOperation(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test entry_point
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeContractCallOperation(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test args
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeContractCallOperation(vb)
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

	jo := koinos.NewContractCallOperation()
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
	o := koinos.NewSystemCallTargetReserved()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeSystemCallTargetReserved(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewSystemCallTargetReserved()
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
	vals := []koinos.ThunkID{
		koinos.ThunkIDPrints,
		koinos.ThunkIDVerifyBlockHeader,
		koinos.ThunkIDApplyBlock,
		koinos.ThunkIDApplyTransaction,
		koinos.ThunkIDApplyReservedOperation,
		koinos.ThunkIDApplyUploadContractOperation,
		koinos.ThunkIDApplyExecuteContractOperation,
		koinos.ThunkIDApplySetSystemCallOperation,
		koinos.ThunkIDDbPutObject,
		koinos.ThunkIDDbGetObject,
		koinos.ThunkIDDbGetNextObject,
		koinos.ThunkIDDbGetPrevObject,
		koinos.ThunkIDExecuteContract,
		koinos.ThunkIDGetContractArgsSize,
		koinos.ThunkIDGetContractArgs,
		koinos.ThunkIDSetContractReturn,
		koinos.ThunkIDExitContract,
		koinos.ThunkIDGetHeadInfo,
		koinos.ThunkIDHash,
		koinos.ThunkIDVerifyBlockSignature,
		koinos.ThunkIDVerifyMerkleRoot,
		koinos.ThunkIDGetTransactionPayer,
		koinos.ThunkIDGetMaxAccountResources,
		koinos.ThunkIDGetTransactionResourceLimit,
	}

	// Make sure all types properly serialize
	for _, x := range vals {
		vb := koinos.NewVariableBlob()
		vb = x.Serialize(vb)

		nvb := koinos.NewVariableBlob()
		ox := koinos.UInt32(x)
		nvb = ox.Serialize(nvb)

		if !bytes.Equal(*vb, *nvb) {
			t.Errorf("Serialized enum does match ideal serialization.")
		}

		_, y, err := koinos.DeserializeThunkID(vb)
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

		r := koinos.NewThunkID()
		if err = json.Unmarshal(jx, r); err != nil {
			t.Error(err)
		}
	}

	// Find a value that is NOT an enum value
	w := getInvalidThunkID()

	tw := koinos.UInt32(w)
	vb := koinos.NewVariableBlob()
	n, _, err := koinos.DeserializeThunkID(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	vb = tw.Serialize(vb)

	if _, _, err := koinos.DeserializeThunkID(vb); err == nil {
		t.Errorf("Deserializing an invalid value did not return an error.")
	}

	je := koinos.NewThunkID()
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

		vb := koinos.NewVariableBlob()
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

func getInvalidThunkID() koinos.ThunkID {
	w := koinos.ThunkIDGetTransactionResourceLimit
	for koinos.IsValidThunkID(w) {
		w++
	}

	return w
}
// ----------------------------------------
//  Struct: ContractCallBundle
// ----------------------------------------

func TestContractCallBundle(t *testing.T) {
	o := koinos.NewContractCallBundle()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeContractCallBundle(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test contract_id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeContractCallBundle(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test entry_point
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeContractCallBundle(vb)
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

	jo := koinos.NewContractCallBundle()
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
	o := koinos.NewSystemCallTarget()
	exerciseSystemCallTargetSerialization(o, t)
	{
		v := koinos.NewSystemCallTarget()
		v.Value = koinos.NewSystemCallTargetReserved()
		exerciseSystemCallTargetSerialization(v, t)

	}
	{
		v := koinos.NewSystemCallTarget()
		v.Value = koinos.NewThunkID()
		exerciseSystemCallTargetSerialization(v, t)

		vb := koinos.VariableBlob{1}
		n, _, err := koinos.DeserializeSystemCallTarget(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewSystemCallTarget()
		v.Value = koinos.NewContractCallBundle()
		exerciseSystemCallTargetSerialization(v, t)

		vb := koinos.VariableBlob{2}
		n, _, err := koinos.DeserializeSystemCallTarget(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}

	// Test bad variant tag
	vb := koinos.VariableBlob{0x80}
	n, _, err := koinos.DeserializeSystemCallTarget(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = koinos.VariableBlob{3}
	n, _, err = koinos.DeserializeSystemCallTarget(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = koinos.NewSystemCallTarget()
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

		variant := koinos.SystemCallTarget{Value: int64(0)}
		vb := koinos.NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := koinos.SystemCallTarget{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseSystemCallTargetSerialization(v *koinos.SystemCallTarget, t *testing.T) {
	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := koinos.DeserializeSystemCallTarget(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := koinos.NewSystemCallTarget()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Struct: SetSystemCallOperation
// ----------------------------------------

func TestSetSystemCallOperation(t *testing.T) {
	o := koinos.NewSetSystemCallOperation()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeSetSystemCallOperation(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test call_id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeSetSystemCallOperation(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test target
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeSetSystemCallOperation(vb)
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

	jo := koinos.NewSetSystemCallOperation()
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
	o := koinos.NewOperation()
	exerciseOperationSerialization(o, t)
	{
		v := koinos.NewOperation()
		v.Value = koinos.NewReservedOperation()
		exerciseOperationSerialization(v, t)

	}
	{
		v := koinos.NewOperation()
		v.Value = koinos.NewNopOperation()
		exerciseOperationSerialization(v, t)

	}
	{
		v := koinos.NewOperation()
		v.Value = koinos.NewCreateSystemContractOperation()
		exerciseOperationSerialization(v, t)

		vb := koinos.VariableBlob{2}
		n, _, err := koinos.DeserializeOperation(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewOperation()
		v.Value = koinos.NewContractCallOperation()
		exerciseOperationSerialization(v, t)

		vb := koinos.VariableBlob{3}
		n, _, err := koinos.DeserializeOperation(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewOperation()
		v.Value = koinos.NewSetSystemCallOperation()
		exerciseOperationSerialization(v, t)

		vb := koinos.VariableBlob{4}
		n, _, err := koinos.DeserializeOperation(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}

	// Test bad variant tag
	vb := koinos.VariableBlob{0x80}
	n, _, err := koinos.DeserializeOperation(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = koinos.VariableBlob{5}
	n, _, err = koinos.DeserializeOperation(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = koinos.NewOperation()
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

		variant := koinos.Operation{Value: int64(0)}
		vb := koinos.NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := koinos.Operation{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseOperationSerialization(v *koinos.Operation, t *testing.T) {
	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := koinos.DeserializeOperation(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := koinos.NewOperation()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Struct: Transaction
// ----------------------------------------

func TestTransaction(t *testing.T) {
	o := koinos.NewTransaction()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeTransaction(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test active_data
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeTransaction(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test passive_data
	vb = &koinos.VariableBlob{0x00}
	n, _, err = koinos.DeserializeTransaction(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test signature_data
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeTransaction(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test operations
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeTransaction(vb)
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

	jo := koinos.NewTransaction()
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
	o := koinos.NewBlock()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeBlock(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test active_data
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeBlock(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test passive_data
	vb = &koinos.VariableBlob{0x00}
	n, _, err = koinos.DeserializeBlock(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test signature_data
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeBlock(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test transactions
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeBlock(vb)
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

	jo := koinos.NewBlock()
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
//  Struct: BlockReceipt
// ----------------------------------------

func TestBlockReceipt(t *testing.T) {
	o := koinos.NewBlockReceipt()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeBlockReceipt(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewBlockReceipt()
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
//  Struct: BlockItem
// ----------------------------------------

func TestBlockItem(t *testing.T) {
	o := koinos.NewBlockItem()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeBlockItem(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test block_id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeBlockItem(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test block_height
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeBlockItem(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test block
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeBlockItem(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test block_receipt
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeBlockItem(vb)
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

	jo := koinos.NewBlockItem()
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
//  Struct: BlockRecord
// ----------------------------------------

func TestBlockRecord(t *testing.T) {
	o := koinos.NewBlockRecord()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeBlockRecord(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test block_id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeBlockRecord(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test block_height
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeBlockRecord(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test previous_block_ids
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeBlockRecord(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test block
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeBlockRecord(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test block_receipt
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeBlockRecord(vb)
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

	jo := koinos.NewBlockRecord()
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
//  Struct: TransactionItem
// ----------------------------------------

func TestTransactionItem(t *testing.T) {
	o := koinos.NewTransactionItem()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeTransactionItem(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test transaction
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeTransactionItem(vb)
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

	jo := koinos.NewTransactionItem()
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
//  Struct: TransactionRecord
// ----------------------------------------

func TestTransactionRecord(t *testing.T) {
	o := koinos.NewTransactionRecord()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeTransactionRecord(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test transaction
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeTransactionRecord(vb)
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

	jo := koinos.NewTransactionRecord()
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
//  Struct: BlockStoreReservedRequest
// ----------------------------------------

func TestBlockStoreReservedRequest(t *testing.T) {
	o := koinos.NewBlockStoreReservedRequest()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeBlockStoreReservedRequest(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewBlockStoreReservedRequest()
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
//  Struct: BlockStoreReservedResponse
// ----------------------------------------

func TestBlockStoreReservedResponse(t *testing.T) {
	o := koinos.NewBlockStoreReservedResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeBlockStoreReservedResponse(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewBlockStoreReservedResponse()
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
//  Struct: GetBlocksByIDRequest
// ----------------------------------------

func TestGetBlocksByIDRequest(t *testing.T) {
	o := koinos.NewGetBlocksByIDRequest()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetBlocksByIDRequest(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test block_id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeGetBlocksByIDRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test return_block_blob
	vb = &koinos.VariableBlob{0x00}
	n, _, err = koinos.DeserializeGetBlocksByIDRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test return_receipt_blob
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeGetBlocksByIDRequest(vb)
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

	jo := koinos.NewGetBlocksByIDRequest()
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
//  Struct: GetBlocksByIDResponse
// ----------------------------------------

func TestGetBlocksByIDResponse(t *testing.T) {
	o := koinos.NewGetBlocksByIDResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetBlocksByIDResponse(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test block_items
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeGetBlocksByIDResponse(vb)
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

	jo := koinos.NewGetBlocksByIDResponse()
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
//  Struct: GetBlocksByHeightRequest
// ----------------------------------------

func TestGetBlocksByHeightRequest(t *testing.T) {
	o := koinos.NewGetBlocksByHeightRequest()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetBlocksByHeightRequest(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test head_block_id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeGetBlocksByHeightRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test ancestor_start_height
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeGetBlocksByHeightRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test num_blocks
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeGetBlocksByHeightRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test return_block
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeGetBlocksByHeightRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test return_receipt
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeGetBlocksByHeightRequest(vb)
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

	jo := koinos.NewGetBlocksByHeightRequest()
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
//  Struct: GetBlocksByHeightResponse
// ----------------------------------------

func TestGetBlocksByHeightResponse(t *testing.T) {
	o := koinos.NewGetBlocksByHeightResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetBlocksByHeightResponse(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test block_items
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeGetBlocksByHeightResponse(vb)
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

	jo := koinos.NewGetBlocksByHeightResponse()
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
//  Struct: AddBlockRequest
// ----------------------------------------

func TestAddBlockRequest(t *testing.T) {
	o := koinos.NewAddBlockRequest()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeAddBlockRequest(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test block_to_add
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeAddBlockRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test previous_block_id
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeAddBlockRequest(vb)
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

	jo := koinos.NewAddBlockRequest()
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
//  Struct: AddBlockResponse
// ----------------------------------------

func TestAddBlockResponse(t *testing.T) {
	o := koinos.NewAddBlockResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeAddBlockResponse(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewAddBlockResponse()
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
//  Struct: AddTransactionRequest
// ----------------------------------------

func TestAddTransactionRequest(t *testing.T) {
	o := koinos.NewAddTransactionRequest()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeAddTransactionRequest(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test transaction_id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeAddTransactionRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test transaction
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeAddTransactionRequest(vb)
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

	jo := koinos.NewAddTransactionRequest()
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
//  Struct: AddTransactionResponse
// ----------------------------------------

func TestAddTransactionResponse(t *testing.T) {
	o := koinos.NewAddTransactionResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeAddTransactionResponse(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewAddTransactionResponse()
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
//  Struct: GetTransactionsByIDRequest
// ----------------------------------------

func TestGetTransactionsByIDRequest(t *testing.T) {
	o := koinos.NewGetTransactionsByIDRequest()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetTransactionsByIDRequest(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test transaction_ids
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeGetTransactionsByIDRequest(vb)
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

	jo := koinos.NewGetTransactionsByIDRequest()
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
//  Struct: GetTransactionsByIDResponse
// ----------------------------------------

func TestGetTransactionsByIDResponse(t *testing.T) {
	o := koinos.NewGetTransactionsByIDResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetTransactionsByIDResponse(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test transaction_items
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeGetTransactionsByIDResponse(vb)
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

	jo := koinos.NewGetTransactionsByIDResponse()
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
//  Struct: BlockStoreErrorResponse
// ----------------------------------------

func TestBlockStoreErrorResponse(t *testing.T) {
	o := koinos.NewBlockStoreErrorResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeBlockStoreErrorResponse(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test error_text
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeBlockStoreErrorResponse(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test error_data
	vb = &koinos.VariableBlob{0x00}
	n, _, err = koinos.DeserializeBlockStoreErrorResponse(vb)
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

	jo := koinos.NewBlockStoreErrorResponse()
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
//  Variant: BlockStoreRequest
// ----------------------------------------

func TestBlockStoreRequest(t *testing.T) {
	o := koinos.NewBlockStoreRequest()
	exerciseBlockStoreRequestSerialization(o, t)
	{
		v := koinos.NewBlockStoreRequest()
		v.Value = koinos.NewBlockStoreReservedRequest()
		exerciseBlockStoreRequestSerialization(v, t)

	}
	{
		v := koinos.NewBlockStoreRequest()
		v.Value = koinos.NewGetBlocksByIDRequest()
		exerciseBlockStoreRequestSerialization(v, t)

		vb := koinos.VariableBlob{1}
		n, _, err := koinos.DeserializeBlockStoreRequest(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewBlockStoreRequest()
		v.Value = koinos.NewGetBlocksByHeightRequest()
		exerciseBlockStoreRequestSerialization(v, t)

		vb := koinos.VariableBlob{2}
		n, _, err := koinos.DeserializeBlockStoreRequest(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewBlockStoreRequest()
		v.Value = koinos.NewAddBlockRequest()
		exerciseBlockStoreRequestSerialization(v, t)

		vb := koinos.VariableBlob{3}
		n, _, err := koinos.DeserializeBlockStoreRequest(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewBlockStoreRequest()
		v.Value = koinos.NewAddTransactionRequest()
		exerciseBlockStoreRequestSerialization(v, t)

		vb := koinos.VariableBlob{4}
		n, _, err := koinos.DeserializeBlockStoreRequest(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewBlockStoreRequest()
		v.Value = koinos.NewGetTransactionsByIDRequest()
		exerciseBlockStoreRequestSerialization(v, t)

		vb := koinos.VariableBlob{5}
		n, _, err := koinos.DeserializeBlockStoreRequest(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}

	// Test bad variant tag
	vb := koinos.VariableBlob{0x80}
	n, _, err := koinos.DeserializeBlockStoreRequest(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = koinos.VariableBlob{6}
	n, _, err = koinos.DeserializeBlockStoreRequest(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = koinos.NewBlockStoreRequest()
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

		variant := koinos.BlockStoreRequest{Value: int64(0)}
		vb := koinos.NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := koinos.BlockStoreRequest{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseBlockStoreRequestSerialization(v *koinos.BlockStoreRequest, t *testing.T) {
	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := koinos.DeserializeBlockStoreRequest(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := koinos.NewBlockStoreRequest()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Variant: BlockStoreResponse
// ----------------------------------------

func TestBlockStoreResponse(t *testing.T) {
	o := koinos.NewBlockStoreResponse()
	exerciseBlockStoreResponseSerialization(o, t)
	{
		v := koinos.NewBlockStoreResponse()
		v.Value = koinos.NewBlockStoreReservedResponse()
		exerciseBlockStoreResponseSerialization(v, t)

	}
	{
		v := koinos.NewBlockStoreResponse()
		v.Value = koinos.NewBlockStoreErrorResponse()
		exerciseBlockStoreResponseSerialization(v, t)

		vb := koinos.VariableBlob{1}
		n, _, err := koinos.DeserializeBlockStoreResponse(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewBlockStoreResponse()
		v.Value = koinos.NewGetBlocksByIDResponse()
		exerciseBlockStoreResponseSerialization(v, t)

		vb := koinos.VariableBlob{2}
		n, _, err := koinos.DeserializeBlockStoreResponse(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewBlockStoreResponse()
		v.Value = koinos.NewGetBlocksByHeightResponse()
		exerciseBlockStoreResponseSerialization(v, t)

		vb := koinos.VariableBlob{3}
		n, _, err := koinos.DeserializeBlockStoreResponse(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewBlockStoreResponse()
		v.Value = koinos.NewAddBlockResponse()
		exerciseBlockStoreResponseSerialization(v, t)

	}
	{
		v := koinos.NewBlockStoreResponse()
		v.Value = koinos.NewAddTransactionResponse()
		exerciseBlockStoreResponseSerialization(v, t)

	}
	{
		v := koinos.NewBlockStoreResponse()
		v.Value = koinos.NewGetTransactionsByIDResponse()
		exerciseBlockStoreResponseSerialization(v, t)

		vb := koinos.VariableBlob{6}
		n, _, err := koinos.DeserializeBlockStoreResponse(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}

	// Test bad variant tag
	vb := koinos.VariableBlob{0x80}
	n, _, err := koinos.DeserializeBlockStoreResponse(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = koinos.VariableBlob{7}
	n, _, err = koinos.DeserializeBlockStoreResponse(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = koinos.NewBlockStoreResponse()
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

		variant := koinos.BlockStoreResponse{Value: int64(0)}
		vb := koinos.NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := koinos.BlockStoreResponse{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseBlockStoreResponseSerialization(v *koinos.BlockStoreResponse, t *testing.T) {
	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := koinos.DeserializeBlockStoreResponse(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := koinos.NewBlockStoreResponse()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Struct: TransactionTopology
// ----------------------------------------

func TestTransactionTopology(t *testing.T) {
	o := koinos.NewTransactionTopology()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeTransactionTopology(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeTransactionTopology(vb)
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

	jo := koinos.NewTransactionTopology()
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
//  Struct: TransactionAccepted
// ----------------------------------------

func TestTransactionAccepted(t *testing.T) {
	o := koinos.NewTransactionAccepted()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeTransactionAccepted(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test topology
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeTransactionAccepted(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test transaction
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeTransactionAccepted(vb)
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

	jo := koinos.NewTransactionAccepted()
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
//  Struct: BlockTopology
// ----------------------------------------

func TestBlockTopology(t *testing.T) {
	o := koinos.NewBlockTopology()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeBlockTopology(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeBlockTopology(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test height
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeBlockTopology(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test previous
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeBlockTopology(vb)
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

	jo := koinos.NewBlockTopology()
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
//  Struct: BlockAccepted
// ----------------------------------------

func TestBlockAccepted(t *testing.T) {
	o := koinos.NewBlockAccepted()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeBlockAccepted(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test topology
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeBlockAccepted(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test block
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeBlockAccepted(vb)
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

	jo := koinos.NewBlockAccepted()
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
//  Enum: SystemCallID
// ----------------------------------------

func TestSystemCallID(t *testing.T) {
	vals := []koinos.SystemCallID{
		koinos.SystemCallIDPrints,
		koinos.SystemCallIDVerifyBlockHeader,
		koinos.SystemCallIDApplyBlock,
		koinos.SystemCallIDApplyTransaction,
		koinos.SystemCallIDApplyReservedOperation,
		koinos.SystemCallIDApplyUploadContractOperation,
		koinos.SystemCallIDApplyExecuteContractOperation,
		koinos.SystemCallIDApplySetSystemCallOperation,
		koinos.SystemCallIDDbPutObject,
		koinos.SystemCallIDDbGetObject,
		koinos.SystemCallIDDbGetNextObject,
		koinos.SystemCallIDDbGetPrevObject,
		koinos.SystemCallIDExecuteContract,
		koinos.SystemCallIDGetContractArgsSize,
		koinos.SystemCallIDGetContractArgs,
		koinos.SystemCallIDSetContractReturn,
		koinos.SystemCallIDExitContract,
		koinos.SystemCallIDGetHeadInfo,
		koinos.SystemCallIDHash,
		koinos.SystemCallIDVerifyBlockSignature,
		koinos.SystemCallIDVerifyMerkleRoot,
		koinos.SystemCallIDGetTransactionPayer,
		koinos.SystemCallIDGetMaxAccountResources,
		koinos.SystemCallIDGetTransactionResourceLimit,
	}

	// Make sure all types properly serialize
	for _, x := range vals {
		vb := koinos.NewVariableBlob()
		vb = x.Serialize(vb)

		nvb := koinos.NewVariableBlob()
		ox := koinos.UInt32(x)
		nvb = ox.Serialize(nvb)

		if !bytes.Equal(*vb, *nvb) {
			t.Errorf("Serialized enum does match ideal serialization.")
		}

		_, y, err := koinos.DeserializeSystemCallID(vb)
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

		r := koinos.NewSystemCallID()
		if err = json.Unmarshal(jx, r); err != nil {
			t.Error(err)
		}
	}

	// Find a value that is NOT an enum value
	w := getInvalidSystemCallID()

	tw := koinos.UInt32(w)
	vb := koinos.NewVariableBlob()
	n, _, err := koinos.DeserializeSystemCallID(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	vb = tw.Serialize(vb)

	if _, _, err := koinos.DeserializeSystemCallID(vb); err == nil {
		t.Errorf("Deserializing an invalid value did not return an error.")
	}

	je := koinos.NewSystemCallID()
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

		vb := koinos.NewVariableBlob()
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

func getInvalidSystemCallID() koinos.SystemCallID {
	w := koinos.SystemCallIDGetTransactionResourceLimit
	for koinos.IsValidSystemCallID(w) {
		w++
	}

	return w
}
// ----------------------------------------
//  Struct: HeadInfo
// ----------------------------------------

func TestHeadInfo(t *testing.T) {
	o := koinos.NewHeadInfo()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeHeadInfo(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeHeadInfo(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test previous_id
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeHeadInfo(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test height
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeHeadInfo(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test last_irreversible_height
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeHeadInfo(vb)
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

	jo := koinos.NewHeadInfo()
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
//  Typedef: AccountType
// ----------------------------------------

func TestAccountType(t *testing.T) {
	o := koinos.NewAccountType()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeAccountType(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeAccountType(vb)
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

	jo := koinos.NewAccountType()
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
//  Struct: VoidType
// ----------------------------------------

func TestVoidType(t *testing.T) {
	o := koinos.NewVoidType()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeVoidType(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewVoidType()
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
	o := koinos.NewPrintsArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializePrintsArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test message
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializePrintsArgs(vb)
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

	jo := koinos.NewPrintsArgs()
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
//  Typedef: PrintsReturn
// ----------------------------------------

func TestPrintsReturn(t *testing.T) {
	o := koinos.NewPrintsReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializePrintsReturn(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewPrintsReturn()
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
//  Struct: VerifyBlockSignatureArgs
// ----------------------------------------

func TestVerifyBlockSignatureArgs(t *testing.T) {
	o := koinos.NewVerifyBlockSignatureArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeVerifyBlockSignatureArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test signature_data
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeVerifyBlockSignatureArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test digest
	vb = &koinos.VariableBlob{0x00}
	n, _, err = koinos.DeserializeVerifyBlockSignatureArgs(vb)
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

	jo := koinos.NewVerifyBlockSignatureArgs()
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
//  Typedef: VerifyBlockSignatureReturn
// ----------------------------------------

func TestVerifyBlockSignatureReturn(t *testing.T) {
	o := koinos.NewVerifyBlockSignatureReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeVerifyBlockSignatureReturn(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeVerifyBlockSignatureReturn(vb)
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

	jo := koinos.NewVerifyBlockSignatureReturn()
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
	o := koinos.NewVerifyMerkleRootArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeVerifyMerkleRootArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test root
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeVerifyMerkleRootArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test hashes
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeVerifyMerkleRootArgs(vb)
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

	jo := koinos.NewVerifyMerkleRootArgs()
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
//  Typedef: VerifyMerkleRootReturn
// ----------------------------------------

func TestVerifyMerkleRootReturn(t *testing.T) {
	o := koinos.NewVerifyMerkleRootReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeVerifyMerkleRootReturn(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeVerifyMerkleRootReturn(vb)
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

	jo := koinos.NewVerifyMerkleRootReturn()
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
	o := koinos.NewApplyBlockArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeApplyBlockArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test block
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeApplyBlockArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test enable_check_passive_data
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeApplyBlockArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test enable_check_block_signature
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeApplyBlockArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test enable_check_transaction_signatures
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeApplyBlockArgs(vb)
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

	jo := koinos.NewApplyBlockArgs()
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
//  Typedef: ApplyBlockReturn
// ----------------------------------------

func TestApplyBlockReturn(t *testing.T) {
	o := koinos.NewApplyBlockReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeApplyBlockReturn(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewApplyBlockReturn()
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
	o := koinos.NewApplyTransactionArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeApplyTransactionArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test trx
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeApplyTransactionArgs(vb)
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

	jo := koinos.NewApplyTransactionArgs()
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
//  Typedef: ApplyTransactionReturn
// ----------------------------------------

func TestApplyTransactionReturn(t *testing.T) {
	o := koinos.NewApplyTransactionReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeApplyTransactionReturn(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewApplyTransactionReturn()
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
	o := koinos.NewApplyUploadContractOperationArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeApplyUploadContractOperationArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test op
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeApplyUploadContractOperationArgs(vb)
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

	jo := koinos.NewApplyUploadContractOperationArgs()
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
//  Typedef: ApplyUploadContractOperationReturn
// ----------------------------------------

func TestApplyUploadContractOperationReturn(t *testing.T) {
	o := koinos.NewApplyUploadContractOperationReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeApplyUploadContractOperationReturn(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewApplyUploadContractOperationReturn()
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
	o := koinos.NewApplyReservedOperationArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeApplyReservedOperationArgs(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewApplyReservedOperationArgs()
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
//  Typedef: ApplyReservedOperationReturn
// ----------------------------------------

func TestApplyReservedOperationReturn(t *testing.T) {
	o := koinos.NewApplyReservedOperationReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeApplyReservedOperationReturn(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewApplyReservedOperationReturn()
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
	o := koinos.NewApplyExecuteContractOperationArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeApplyExecuteContractOperationArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test op
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeApplyExecuteContractOperationArgs(vb)
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

	jo := koinos.NewApplyExecuteContractOperationArgs()
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
//  Typedef: ApplyExecuteContractOperationReturn
// ----------------------------------------

func TestApplyExecuteContractOperationReturn(t *testing.T) {
	o := koinos.NewApplyExecuteContractOperationReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeApplyExecuteContractOperationReturn(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewApplyExecuteContractOperationReturn()
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
	o := koinos.NewApplySetSystemCallOperationArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeApplySetSystemCallOperationArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test op
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeApplySetSystemCallOperationArgs(vb)
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

	jo := koinos.NewApplySetSystemCallOperationArgs()
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
//  Typedef: ApplySetSystemCallOperationReturn
// ----------------------------------------

func TestApplySetSystemCallOperationReturn(t *testing.T) {
	o := koinos.NewApplySetSystemCallOperationReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeApplySetSystemCallOperationReturn(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewApplySetSystemCallOperationReturn()
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
	o := koinos.NewDbPutObjectArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeDbPutObjectArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test space
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeDbPutObjectArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test key
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeDbPutObjectArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test obj
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeDbPutObjectArgs(vb)
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

	jo := koinos.NewDbPutObjectArgs()
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
//  Typedef: DbPutObjectReturn
// ----------------------------------------

func TestDbPutObjectReturn(t *testing.T) {
	o := koinos.NewDbPutObjectReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeDbPutObjectReturn(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeDbPutObjectReturn(vb)
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

	jo := koinos.NewDbPutObjectReturn()
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
	o := koinos.NewDbGetObjectArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeDbGetObjectArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test space
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeDbGetObjectArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test key
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeDbGetObjectArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test object_size_hint
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeDbGetObjectArgs(vb)
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

	jo := koinos.NewDbGetObjectArgs()
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
//  Typedef: DbGetObjectReturn
// ----------------------------------------

func TestDbGetObjectReturn(t *testing.T) {
	o := koinos.NewDbGetObjectReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeDbGetObjectReturn(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeDbGetObjectReturn(vb)
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

	jo := koinos.NewDbGetObjectReturn()
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
	o := koinos.NewDbGetNextObjectArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeDbGetNextObjectArgs(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeDbGetNextObjectArgs(vb)
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

	jo := koinos.NewDbGetNextObjectArgs()
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
//  Typedef: DbGetNextObjectReturn
// ----------------------------------------

func TestDbGetNextObjectReturn(t *testing.T) {
	o := koinos.NewDbGetNextObjectReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeDbGetNextObjectReturn(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeDbGetNextObjectReturn(vb)
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

	jo := koinos.NewDbGetNextObjectReturn()
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
	o := koinos.NewDbGetPrevObjectArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeDbGetPrevObjectArgs(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeDbGetPrevObjectArgs(vb)
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

	jo := koinos.NewDbGetPrevObjectArgs()
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
//  Typedef: DbGetPrevObjectReturn
// ----------------------------------------

func TestDbGetPrevObjectReturn(t *testing.T) {
	o := koinos.NewDbGetPrevObjectReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeDbGetPrevObjectReturn(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeDbGetPrevObjectReturn(vb)
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

	jo := koinos.NewDbGetPrevObjectReturn()
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
	o := koinos.NewExecuteContractArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeExecuteContractArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test contract_id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeExecuteContractArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test entry_point
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeExecuteContractArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test args
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeExecuteContractArgs(vb)
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

	jo := koinos.NewExecuteContractArgs()
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
//  Typedef: ExecuteContractReturn
// ----------------------------------------

func TestExecuteContractReturn(t *testing.T) {
	o := koinos.NewExecuteContractReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeExecuteContractReturn(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeExecuteContractReturn(vb)
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

	jo := koinos.NewExecuteContractReturn()
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
	o := koinos.NewGetContractArgsSizeArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetContractArgsSizeArgs(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewGetContractArgsSizeArgs()
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
//  Typedef: GetContractArgsSizeReturn
// ----------------------------------------

func TestGetContractArgsSizeReturn(t *testing.T) {
	o := koinos.NewGetContractArgsSizeReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetContractArgsSizeReturn(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeGetContractArgsSizeReturn(vb)
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

	jo := koinos.NewGetContractArgsSizeReturn()
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
	o := koinos.NewGetContractArgsArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetContractArgsArgs(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewGetContractArgsArgs()
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
//  Typedef: GetContractArgsReturn
// ----------------------------------------

func TestGetContractArgsReturn(t *testing.T) {
	o := koinos.NewGetContractArgsReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetContractArgsReturn(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeGetContractArgsReturn(vb)
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

	jo := koinos.NewGetContractArgsReturn()
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
	o := koinos.NewSetContractReturnArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeSetContractReturnArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test value
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeSetContractReturnArgs(vb)
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

	jo := koinos.NewSetContractReturnArgs()
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
//  Typedef: SetContractReturnReturn
// ----------------------------------------

func TestSetContractReturnReturn(t *testing.T) {
	o := koinos.NewSetContractReturnReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeSetContractReturnReturn(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewSetContractReturnReturn()
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
	o := koinos.NewExitContractArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeExitContractArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test exit_code
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeExitContractArgs(vb)
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

	jo := koinos.NewExitContractArgs()
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
//  Typedef: ExitContractReturn
// ----------------------------------------

func TestExitContractReturn(t *testing.T) {
	o := koinos.NewExitContractReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeExitContractReturn(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewExitContractReturn()
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
	o := koinos.NewGetHeadInfoArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetHeadInfoArgs(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewGetHeadInfoArgs()
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
//  Typedef: GetHeadInfoReturn
// ----------------------------------------

func TestGetHeadInfoReturn(t *testing.T) {
	o := koinos.NewGetHeadInfoReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetHeadInfoReturn(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeGetHeadInfoReturn(vb)
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

	jo := koinos.NewGetHeadInfoReturn()
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
	o := koinos.NewHashArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeHashArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test code
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeHashArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test obj
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeHashArgs(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test size
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeHashArgs(vb)
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

	jo := koinos.NewHashArgs()
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
//  Typedef: HashReturn
// ----------------------------------------

func TestHashReturn(t *testing.T) {
	o := koinos.NewHashReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeHashReturn(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeHashReturn(vb)
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

	jo := koinos.NewHashReturn()
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
//  Struct: GetTransactionPayerArgs
// ----------------------------------------

func TestGetTransactionPayerArgs(t *testing.T) {
	o := koinos.NewGetTransactionPayerArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetTransactionPayerArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test transaction
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeGetTransactionPayerArgs(vb)
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

	jo := koinos.NewGetTransactionPayerArgs()
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
//  Typedef: GetTransactionPayerReturn
// ----------------------------------------

func TestGetTransactionPayerReturn(t *testing.T) {
	o := koinos.NewGetTransactionPayerReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetTransactionPayerReturn(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeGetTransactionPayerReturn(vb)
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

	jo := koinos.NewGetTransactionPayerReturn()
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
//  Struct: GetMaxAccountResourcesArgs
// ----------------------------------------

func TestGetMaxAccountResourcesArgs(t *testing.T) {
	o := koinos.NewGetMaxAccountResourcesArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetMaxAccountResourcesArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test account
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeGetMaxAccountResourcesArgs(vb)
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

	jo := koinos.NewGetMaxAccountResourcesArgs()
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
//  Typedef: GetMaxAccountResourcesReturn
// ----------------------------------------

func TestGetMaxAccountResourcesReturn(t *testing.T) {
	o := koinos.NewGetMaxAccountResourcesReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetMaxAccountResourcesReturn(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeGetMaxAccountResourcesReturn(vb)
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

	jo := koinos.NewGetMaxAccountResourcesReturn()
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
//  Struct: GetTransactionResourceLimitArgs
// ----------------------------------------

func TestGetTransactionResourceLimitArgs(t *testing.T) {
	o := koinos.NewGetTransactionResourceLimitArgs()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetTransactionResourceLimitArgs(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test transaction
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeGetTransactionResourceLimitArgs(vb)
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

	jo := koinos.NewGetTransactionResourceLimitArgs()
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
//  Typedef: GetTransactionResourceLimitReturn
// ----------------------------------------

func TestGetTransactionResourceLimitReturn(t *testing.T) {
	o := koinos.NewGetTransactionResourceLimitReturn()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetTransactionResourceLimitReturn(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeGetTransactionResourceLimitReturn(vb)
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

	jo := koinos.NewGetTransactionResourceLimitReturn()
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
//  Struct: ChainReservedRequest
// ----------------------------------------

func TestChainReservedRequest(t *testing.T) {
	o := koinos.NewChainReservedRequest()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeChainReservedRequest(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewChainReservedRequest()
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
//  Struct: SubmitBlockRequest
// ----------------------------------------

func TestSubmitBlockRequest(t *testing.T) {
	o := koinos.NewSubmitBlockRequest()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeSubmitBlockRequest(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test topology
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeSubmitBlockRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test block
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeSubmitBlockRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test verify_passive_data
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeSubmitBlockRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test verify_block_signature
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeSubmitBlockRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test verify_transaction_signatures
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeSubmitBlockRequest(vb)
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

	jo := koinos.NewSubmitBlockRequest()
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
//  Struct: SubmitTransactionRequest
// ----------------------------------------

func TestSubmitTransactionRequest(t *testing.T) {
	o := koinos.NewSubmitTransactionRequest()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeSubmitTransactionRequest(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test topology
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeSubmitTransactionRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test transaction
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeSubmitTransactionRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test verify_passive_data
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeSubmitTransactionRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test verify_transaction_signatures
	vb = &koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, _, err = koinos.DeserializeSubmitTransactionRequest(vb)
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

	jo := koinos.NewSubmitTransactionRequest()
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
//  Struct: GetHeadInfoRequest
// ----------------------------------------

func TestGetHeadInfoRequest(t *testing.T) {
	o := koinos.NewGetHeadInfoRequest()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetHeadInfoRequest(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewGetHeadInfoRequest()
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
//  Struct: GetChainIDRequest
// ----------------------------------------

func TestGetChainIDRequest(t *testing.T) {
	o := koinos.NewGetChainIDRequest()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetChainIDRequest(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewGetChainIDRequest()
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
//  Struct: GetPendingTransactionsRequest
// ----------------------------------------

func TestGetPendingTransactionsRequest(t *testing.T) {
	o := koinos.NewGetPendingTransactionsRequest()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetPendingTransactionsRequest(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test start
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeGetPendingTransactionsRequest(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test limit
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeGetPendingTransactionsRequest(vb)
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

	jo := koinos.NewGetPendingTransactionsRequest()
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
//  Struct: GetForkHeadsRequest
// ----------------------------------------

func TestGetForkHeadsRequest(t *testing.T) {
	o := koinos.NewGetForkHeadsRequest()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetForkHeadsRequest(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewGetForkHeadsRequest()
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
//  Variant: ChainRPCRequest
// ----------------------------------------

func TestChainRPCRequest(t *testing.T) {
	o := koinos.NewChainRPCRequest()
	exerciseChainRPCRequestSerialization(o, t)
	{
		v := koinos.NewChainRPCRequest()
		v.Value = koinos.NewChainReservedRequest()
		exerciseChainRPCRequestSerialization(v, t)

	}
	{
		v := koinos.NewChainRPCRequest()
		v.Value = koinos.NewSubmitBlockRequest()
		exerciseChainRPCRequestSerialization(v, t)

		vb := koinos.VariableBlob{1}
		n, _, err := koinos.DeserializeChainRPCRequest(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewChainRPCRequest()
		v.Value = koinos.NewSubmitTransactionRequest()
		exerciseChainRPCRequestSerialization(v, t)

		vb := koinos.VariableBlob{2}
		n, _, err := koinos.DeserializeChainRPCRequest(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewChainRPCRequest()
		v.Value = koinos.NewGetHeadInfoRequest()
		exerciseChainRPCRequestSerialization(v, t)

	}
	{
		v := koinos.NewChainRPCRequest()
		v.Value = koinos.NewGetChainIDRequest()
		exerciseChainRPCRequestSerialization(v, t)

	}
	{
		v := koinos.NewChainRPCRequest()
		v.Value = koinos.NewGetPendingTransactionsRequest()
		exerciseChainRPCRequestSerialization(v, t)

		vb := koinos.VariableBlob{5}
		n, _, err := koinos.DeserializeChainRPCRequest(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewChainRPCRequest()
		v.Value = koinos.NewGetForkHeadsRequest()
		exerciseChainRPCRequestSerialization(v, t)

	}

	// Test bad variant tag
	vb := koinos.VariableBlob{0x80}
	n, _, err := koinos.DeserializeChainRPCRequest(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = koinos.VariableBlob{7}
	n, _, err = koinos.DeserializeChainRPCRequest(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = koinos.NewChainRPCRequest()
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

		variant := koinos.ChainRPCRequest{Value: int64(0)}
		vb := koinos.NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := koinos.ChainRPCRequest{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseChainRPCRequestSerialization(v *koinos.ChainRPCRequest, t *testing.T) {
	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := koinos.DeserializeChainRPCRequest(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := koinos.NewChainRPCRequest()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Struct: ChainReservedResponse
// ----------------------------------------

func TestChainReservedResponse(t *testing.T) {
	o := koinos.NewChainReservedResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeChainReservedResponse(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewChainReservedResponse()
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
//  Struct: ChainErrorResponse
// ----------------------------------------

func TestChainErrorResponse(t *testing.T) {
	o := koinos.NewChainErrorResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeChainErrorResponse(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test error_text
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeChainErrorResponse(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test error_data
	vb = &koinos.VariableBlob{0x00}
	n, _, err = koinos.DeserializeChainErrorResponse(vb)
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

	jo := koinos.NewChainErrorResponse()
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
//  Struct: SubmitBlockResponse
// ----------------------------------------

func TestSubmitBlockResponse(t *testing.T) {
	o := koinos.NewSubmitBlockResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeSubmitBlockResponse(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewSubmitBlockResponse()
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
//  Struct: SubmitTransactionResponse
// ----------------------------------------

func TestSubmitTransactionResponse(t *testing.T) {
	o := koinos.NewSubmitTransactionResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeSubmitTransactionResponse(vb)
	if err != nil {
		t.Error(err)
	}
	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewSubmitTransactionResponse()
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
//  Struct: GetHeadInfoResponse
// ----------------------------------------

func TestGetHeadInfoResponse(t *testing.T) {
	o := koinos.NewGetHeadInfoResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetHeadInfoResponse(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeGetHeadInfoResponse(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test height
	vb = &koinos.VariableBlob{0x00, 0x00}
	n, _, err = koinos.DeserializeGetHeadInfoResponse(vb)
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

	jo := koinos.NewGetHeadInfoResponse()
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
//  Struct: GetChainIDResponse
// ----------------------------------------

func TestGetChainIDResponse(t *testing.T) {
	o := koinos.NewGetChainIDResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetChainIDResponse(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test chain_id
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeGetChainIDResponse(vb)
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

	jo := koinos.NewGetChainIDResponse()
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
//  Struct: GetPendingTransactionsResponse
// ----------------------------------------

func TestGetPendingTransactionsResponse(t *testing.T) {
	o := koinos.NewGetPendingTransactionsResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetPendingTransactionsResponse(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test transactions
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeGetPendingTransactionsResponse(vb)
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

	jo := koinos.NewGetPendingTransactionsResponse()
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
//  Struct: GetForkHeadsResponse
// ----------------------------------------

func TestGetForkHeadsResponse(t *testing.T) {
	o := koinos.NewGetForkHeadsResponse()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetForkHeadsResponse(vb)
	if err != nil {
		t.Error(err)
	}

	var n uint64
	// Test fork_heads
	vb = &koinos.VariableBlob{}
	n, _, err = koinos.DeserializeGetForkHeadsResponse(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test last_irreversible_block
	vb = &koinos.VariableBlob{0x00}
	n, _, err = koinos.DeserializeGetForkHeadsResponse(vb)
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

	jo := koinos.NewGetForkHeadsResponse()
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
//  Variant: ChainRPCResponse
// ----------------------------------------

func TestChainRPCResponse(t *testing.T) {
	o := koinos.NewChainRPCResponse()
	exerciseChainRPCResponseSerialization(o, t)
	{
		v := koinos.NewChainRPCResponse()
		v.Value = koinos.NewChainReservedResponse()
		exerciseChainRPCResponseSerialization(v, t)

	}
	{
		v := koinos.NewChainRPCResponse()
		v.Value = koinos.NewChainErrorResponse()
		exerciseChainRPCResponseSerialization(v, t)

		vb := koinos.VariableBlob{1}
		n, _, err := koinos.DeserializeChainRPCResponse(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewChainRPCResponse()
		v.Value = koinos.NewSubmitBlockResponse()
		exerciseChainRPCResponseSerialization(v, t)

	}
	{
		v := koinos.NewChainRPCResponse()
		v.Value = koinos.NewSubmitTransactionResponse()
		exerciseChainRPCResponseSerialization(v, t)

	}
	{
		v := koinos.NewChainRPCResponse()
		v.Value = koinos.NewGetHeadInfoResponse()
		exerciseChainRPCResponseSerialization(v, t)

		vb := koinos.VariableBlob{4}
		n, _, err := koinos.DeserializeChainRPCResponse(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewChainRPCResponse()
		v.Value = koinos.NewGetChainIDResponse()
		exerciseChainRPCResponseSerialization(v, t)

		vb := koinos.VariableBlob{5}
		n, _, err := koinos.DeserializeChainRPCResponse(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewChainRPCResponse()
		v.Value = koinos.NewGetPendingTransactionsResponse()
		exerciseChainRPCResponseSerialization(v, t)

		vb := koinos.VariableBlob{6}
		n, _, err := koinos.DeserializeChainRPCResponse(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewChainRPCResponse()
		v.Value = koinos.NewGetForkHeadsResponse()
		exerciseChainRPCResponseSerialization(v, t)

		vb := koinos.VariableBlob{7}
		n, _, err := koinos.DeserializeChainRPCResponse(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}

	// Test bad variant tag
	vb := koinos.VariableBlob{0x80}
	n, _, err := koinos.DeserializeChainRPCResponse(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = koinos.VariableBlob{8}
	n, _, err = koinos.DeserializeChainRPCResponse(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = koinos.NewChainRPCResponse()
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

		variant := koinos.ChainRPCResponse{Value: int64(0)}
		vb := koinos.NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := koinos.ChainRPCResponse{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseChainRPCResponseSerialization(v *koinos.ChainRPCResponse, t *testing.T) {
	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := koinos.DeserializeChainRPCResponse(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := koinos.NewChainRPCResponse()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Typedef: SignatureType
// ----------------------------------------

func TestSignatureType(t *testing.T) {
	o := koinos.NewSignatureType()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeSignatureType(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeSignatureType(vb)
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

	jo := koinos.NewSignatureType()
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
//  Typedef: ReservedQueryParams
// ----------------------------------------

func TestReservedQueryParams(t *testing.T) {
	o := koinos.NewReservedQueryParams()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeReservedQueryParams(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewReservedQueryParams()
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
//  Typedef: GetHeadInfoParams
// ----------------------------------------

func TestGetHeadInfoParams(t *testing.T) {
	o := koinos.NewGetHeadInfoParams()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetHeadInfoParams(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewGetHeadInfoParams()
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
//  Typedef: GetChainIDParams
// ----------------------------------------

func TestGetChainIDParams(t *testing.T) {
	o := koinos.NewGetChainIDParams()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetChainIDParams(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewGetChainIDParams()
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
//  Typedef: GetPendingTransactionsParams
// ----------------------------------------

func TestGetPendingTransactionsParams(t *testing.T) {
	o := koinos.NewGetPendingTransactionsParams()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetPendingTransactionsParams(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeGetPendingTransactionsParams(vb)
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

	jo := koinos.NewGetPendingTransactionsParams()
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
//  Variant: QueryParamItem
// ----------------------------------------

func TestQueryParamItem(t *testing.T) {
	o := koinos.NewQueryParamItem()
	exerciseQueryParamItemSerialization(o, t)
	{
		v := koinos.NewQueryParamItem()
		v.Value = koinos.NewReservedQueryParams()
		exerciseQueryParamItemSerialization(v, t)

	}
	{
		v := koinos.NewQueryParamItem()
		v.Value = koinos.NewGetHeadInfoParams()
		exerciseQueryParamItemSerialization(v, t)

	}
	{
		v := koinos.NewQueryParamItem()
		v.Value = koinos.NewGetChainIDParams()
		exerciseQueryParamItemSerialization(v, t)

	}
	{
		v := koinos.NewQueryParamItem()
		v.Value = koinos.NewGetPendingTransactionsParams()
		exerciseQueryParamItemSerialization(v, t)

		vb := koinos.VariableBlob{3}
		n, _, err := koinos.DeserializeQueryParamItem(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}

	// Test bad variant tag
	vb := koinos.VariableBlob{0x80}
	n, _, err := koinos.DeserializeQueryParamItem(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = koinos.VariableBlob{4}
	n, _, err = koinos.DeserializeQueryParamItem(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = koinos.NewQueryParamItem()
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

		variant := koinos.QueryParamItem{Value: int64(0)}
		vb := koinos.NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := koinos.QueryParamItem{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseQueryParamItemSerialization(v *koinos.QueryParamItem, t *testing.T) {
	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := koinos.DeserializeQueryParamItem(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := koinos.NewQueryParamItem()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Typedef: QuerySubmission
// ----------------------------------------

func TestQuerySubmission(t *testing.T) {
	o := koinos.NewQuerySubmission()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeQuerySubmission(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeQuerySubmission(vb)
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

	jo := koinos.NewQuerySubmission()
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
//  Typedef: ReservedQueryResult
// ----------------------------------------

func TestReservedQueryResult(t *testing.T) {
	o := koinos.NewReservedQueryResult()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeReservedQueryResult(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewReservedQueryResult()
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
//  Typedef: QueryError
// ----------------------------------------

func TestQueryError(t *testing.T) {
	o := koinos.NewQueryError()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeQueryError(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeQueryError(vb)
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

	jo := koinos.NewQueryError()
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
//  Typedef: GetHeadInfoResult
// ----------------------------------------

func TestGetHeadInfoResult(t *testing.T) {
	o := koinos.NewGetHeadInfoResult()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetHeadInfoResult(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeGetHeadInfoResult(vb)
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

	jo := koinos.NewGetHeadInfoResult()
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
//  Typedef: GetChainIDResult
// ----------------------------------------

func TestGetChainIDResult(t *testing.T) {
	o := koinos.NewGetChainIDResult()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetChainIDResult(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeGetChainIDResult(vb)
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

	jo := koinos.NewGetChainIDResult()
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
//  Typedef: GetPendingTransactionsResult
// ----------------------------------------

func TestGetPendingTransactionsResult(t *testing.T) {
	o := koinos.NewGetPendingTransactionsResult()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetPendingTransactionsResult(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeGetPendingTransactionsResult(vb)
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

	jo := koinos.NewGetPendingTransactionsResult()
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
	o := koinos.NewQueryItemResult()
	exerciseQueryItemResultSerialization(o, t)
	{
		v := koinos.NewQueryItemResult()
		v.Value = koinos.NewReservedQueryResult()
		exerciseQueryItemResultSerialization(v, t)

	}
	{
		v := koinos.NewQueryItemResult()
		v.Value = koinos.NewQueryError()
		exerciseQueryItemResultSerialization(v, t)

		vb := koinos.VariableBlob{1}
		n, _, err := koinos.DeserializeQueryItemResult(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewQueryItemResult()
		v.Value = koinos.NewGetHeadInfoResult()
		exerciseQueryItemResultSerialization(v, t)

		vb := koinos.VariableBlob{2}
		n, _, err := koinos.DeserializeQueryItemResult(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewQueryItemResult()
		v.Value = koinos.NewGetChainIDResult()
		exerciseQueryItemResultSerialization(v, t)

		vb := koinos.VariableBlob{3}
		n, _, err := koinos.DeserializeQueryItemResult(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewQueryItemResult()
		v.Value = koinos.NewGetPendingTransactionsResult()
		exerciseQueryItemResultSerialization(v, t)

		vb := koinos.VariableBlob{4}
		n, _, err := koinos.DeserializeQueryItemResult(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}

	// Test bad variant tag
	vb := koinos.VariableBlob{0x80}
	n, _, err := koinos.DeserializeQueryItemResult(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = koinos.VariableBlob{5}
	n, _, err = koinos.DeserializeQueryItemResult(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = koinos.NewQueryItemResult()
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

		variant := koinos.QueryItemResult{Value: int64(0)}
		vb := koinos.NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := koinos.QueryItemResult{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseQueryItemResultSerialization(v *koinos.QueryItemResult, t *testing.T) {
	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := koinos.DeserializeQueryItemResult(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := koinos.NewQueryItemResult()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Typedef: QuerySubmissionResult
// ----------------------------------------

func TestQuerySubmissionResult(t *testing.T) {
	o := koinos.NewQuerySubmissionResult()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeQuerySubmissionResult(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeQuerySubmissionResult(vb)
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

	jo := koinos.NewQuerySubmissionResult()
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
//  Typedef: ReservedSubmission
// ----------------------------------------

func TestReservedSubmission(t *testing.T) {
	o := koinos.NewReservedSubmission()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeReservedSubmission(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewReservedSubmission()
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
//  Typedef: BlockSubmission
// ----------------------------------------

func TestBlockSubmission(t *testing.T) {
	o := koinos.NewBlockSubmission()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeBlockSubmission(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeBlockSubmission(vb)
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

	jo := koinos.NewBlockSubmission()
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
//  Typedef: TransactionSubmission
// ----------------------------------------

func TestTransactionSubmission(t *testing.T) {
	o := koinos.NewTransactionSubmission()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeTransactionSubmission(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeTransactionSubmission(vb)
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

	jo := koinos.NewTransactionSubmission()
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
//  Typedef: GetForkHeadsSubmission
// ----------------------------------------

func TestGetForkHeadsSubmission(t *testing.T) {
	o := koinos.NewGetForkHeadsSubmission()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetForkHeadsSubmission(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewGetForkHeadsSubmission()
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
//  Variant: SubmissionItem
// ----------------------------------------

func TestSubmissionItem(t *testing.T) {
	o := koinos.NewSubmissionItem()
	exerciseSubmissionItemSerialization(o, t)
	{
		v := koinos.NewSubmissionItem()
		v.Value = koinos.NewReservedSubmission()
		exerciseSubmissionItemSerialization(v, t)

	}
	{
		v := koinos.NewSubmissionItem()
		v.Value = koinos.NewBlockSubmission()
		exerciseSubmissionItemSerialization(v, t)

		vb := koinos.VariableBlob{1}
		n, _, err := koinos.DeserializeSubmissionItem(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewSubmissionItem()
		v.Value = koinos.NewTransactionSubmission()
		exerciseSubmissionItemSerialization(v, t)

		vb := koinos.VariableBlob{2}
		n, _, err := koinos.DeserializeSubmissionItem(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewSubmissionItem()
		v.Value = koinos.NewQuerySubmission()
		exerciseSubmissionItemSerialization(v, t)

		vb := koinos.VariableBlob{3}
		n, _, err := koinos.DeserializeSubmissionItem(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewSubmissionItem()
		v.Value = koinos.NewGetForkHeadsSubmission()
		exerciseSubmissionItemSerialization(v, t)

	}

	// Test bad variant tag
	vb := koinos.VariableBlob{0x80}
	n, _, err := koinos.DeserializeSubmissionItem(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = koinos.VariableBlob{5}
	n, _, err = koinos.DeserializeSubmissionItem(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = koinos.NewSubmissionItem()
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

		variant := koinos.SubmissionItem{Value: int64(0)}
		vb := koinos.NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := koinos.SubmissionItem{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseSubmissionItemSerialization(v *koinos.SubmissionItem, t *testing.T) {
	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := koinos.DeserializeSubmissionItem(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := koinos.NewSubmissionItem()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}

// ----------------------------------------
//  Typedef: ReservedSubmissionResult
// ----------------------------------------

func TestReservedSubmissionResult(t *testing.T) {
	o := koinos.NewReservedSubmissionResult()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeReservedSubmissionResult(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewReservedSubmissionResult()
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
//  Typedef: BlockSubmissionResult
// ----------------------------------------

func TestBlockSubmissionResult(t *testing.T) {
	o := koinos.NewBlockSubmissionResult()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeBlockSubmissionResult(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewBlockSubmissionResult()
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
//  Typedef: TransactionSubmissionResult
// ----------------------------------------

func TestTransactionSubmissionResult(t *testing.T) {
	o := koinos.NewTransactionSubmissionResult()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeTransactionSubmissionResult(vb)
	if err != nil {
		t.Error(err)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewTransactionSubmissionResult()
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
//  Typedef: GetForkHeadsSubmissionResult
// ----------------------------------------

func TestGetForkHeadsSubmissionResult(t *testing.T) {
	o := koinos.NewGetForkHeadsSubmissionResult()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeGetForkHeadsSubmissionResult(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeGetForkHeadsSubmissionResult(vb)
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

	jo := koinos.NewGetForkHeadsSubmissionResult()
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
//  Typedef: SubmissionErrorResult
// ----------------------------------------

func TestSubmissionErrorResult(t *testing.T) {
	o := koinos.NewSubmissionErrorResult()

	vb := koinos.NewVariableBlob()
	vb = o.Serialize(vb)

	_, _, err := koinos.DeserializeSubmissionErrorResult(vb)
	if err != nil {
		t.Error(err)
	}

	vb = koinos.NewVariableBlob()
	size, _, err := koinos.DeserializeSubmissionErrorResult(vb)
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

	jo := koinos.NewSubmissionErrorResult()
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
//  Variant: SubmissionResult
// ----------------------------------------

func TestSubmissionResult(t *testing.T) {
	o := koinos.NewSubmissionResult()
	exerciseSubmissionResultSerialization(o, t)
	{
		v := koinos.NewSubmissionResult()
		v.Value = koinos.NewReservedSubmissionResult()
		exerciseSubmissionResultSerialization(v, t)

	}
	{
		v := koinos.NewSubmissionResult()
		v.Value = koinos.NewBlockSubmissionResult()
		exerciseSubmissionResultSerialization(v, t)

	}
	{
		v := koinos.NewSubmissionResult()
		v.Value = koinos.NewTransactionSubmissionResult()
		exerciseSubmissionResultSerialization(v, t)

	}
	{
		v := koinos.NewSubmissionResult()
		v.Value = koinos.NewQuerySubmissionResult()
		exerciseSubmissionResultSerialization(v, t)

		vb := koinos.VariableBlob{3}
		n, _, err := koinos.DeserializeSubmissionResult(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewSubmissionResult()
		v.Value = koinos.NewGetForkHeadsSubmissionResult()
		exerciseSubmissionResultSerialization(v, t)

		vb := koinos.VariableBlob{4}
		n, _, err := koinos.DeserializeSubmissionResult(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}
	{
		v := koinos.NewSubmissionResult()
		v.Value = koinos.NewSubmissionErrorResult()
		exerciseSubmissionResultSerialization(v, t)

		vb := koinos.VariableBlob{5}
		n, _, err := koinos.DeserializeSubmissionResult(&vb)
		if err == nil {
			t.Errorf("err == nil")
		}
		if n != 0 {
			t.Errorf("Bytes were consumed on error")
		}
	}

	// Test bad variant tag
	vb := koinos.VariableBlob{0x80}
	n, _, err := koinos.DeserializeSubmissionResult(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test unknown tag
	vb = koinos.VariableBlob{6}
	n, _, err = koinos.DeserializeSubmissionResult(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}

	// Test nonsensical json
	o = koinos.NewSubmissionResult()
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

		variant := koinos.SubmissionResult{Value: int64(0)}
		vb := koinos.NewVariableBlob()
		_ = variant.Serialize(vb)
	}()

	// Test panic when jsonifying an unknown variant tag
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Serializing an incorrect variant tag did not panic")
			}
		}()

		variant := koinos.SubmissionResult{Value: int64(0)}
		_, _ = json.Marshal(&variant)
	}()
}

func exerciseSubmissionResultSerialization(v *koinos.SubmissionResult, t *testing.T) {
	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, _, err := koinos.DeserializeSubmissionResult(vb)
	if err != nil {
		t.Error(err)
	}

	jv, jerr := json.Marshal(v)
	if jerr != nil {
		t.Error(jerr)
	}

	nv := koinos.NewSubmissionResult()
	if jerr = json.Unmarshal(jv, nv); jerr != nil {
		t.Error(jerr)
	}
}


// ----------------------------------------
//  FixedBlob20
// ----------------------------------------

func TestFixedBlob20(t *testing.T) {
	fb := koinos.NewFixedBlob20()
	for i := 0; i < 20; i++ {
		fb[i] = byte((20 + i) % 256)
	}

	vb := koinos.NewVariableBlob()
	vb = fb.Serialize(vb)

	size, nfb, err := koinos.DeserializeFixedBlob20(vb)
	if err != nil {
		t.Error(err)
	}

	if size != 20 {
		t.Errorf("Fixedblob deserialization consumed %d bytes. Expected %d bytes.", size, 20)
	}

	if !bytes.Equal(fb[:], nfb[:]) {
		t.Errorf("Binary deserialized fixed blob does not match source.")
	}

	vb = koinos.NewVariableBlob()
	size, _, err = koinos.DeserializeFixedBlob20(vb)
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

	jfb := koinos.NewFixedBlob20()
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
	s := koinos.EncodeBytes(wfb)
	err = json.Unmarshal([]byte("\""+s+"\""), &jfb)
	if err == nil {
		t.Errorf("deserialized JSON from a string which was not the proper size.")
	}

	// Encode with invalid start character of '0' and ensure it returns an error
	wfb = make([]byte, 20)
	for i := 0; i < 20; i++ {
		wfb[i] = byte((20 + i) % 256)
	}
	s = koinos.EncodeBytes(wfb)
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
	fb := koinos.NewFixedBlob65()
	for i := 0; i < 65; i++ {
		fb[i] = byte((65 + i) % 256)
	}

	vb := koinos.NewVariableBlob()
	vb = fb.Serialize(vb)

	size, nfb, err := koinos.DeserializeFixedBlob65(vb)
	if err != nil {
		t.Error(err)
	}

	if size != 65 {
		t.Errorf("Fixedblob deserialization consumed %d bytes. Expected %d bytes.", size, 65)
	}

	if !bytes.Equal(fb[:], nfb[:]) {
		t.Errorf("Binary deserialized fixed blob does not match source.")
	}

	vb = koinos.NewVariableBlob()
	size, _, err = koinos.DeserializeFixedBlob65(vb)
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

	jfb := koinos.NewFixedBlob65()
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
	s := koinos.EncodeBytes(wfb)
	err = json.Unmarshal([]byte("\""+s+"\""), &jfb)
	if err == nil {
		t.Errorf("deserialized JSON from a string which was not the proper size.")
	}

	// Encode with invalid start character of '0' and ensure it returns an error
	wfb = make([]byte, 65)
	for i := 0; i < 65; i++ {
		wfb[i] = byte((65 + i) % 256)
	}
	s = koinos.EncodeBytes(wfb)
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
	o := koinos.NewOpaqueActiveBlockData()

	o.Box()
	if !o.IsBoxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if !o.IsBoxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if o.IsBoxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if o.IsBoxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.Box() // Call Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.Unbox()
	vb := o.GetBlob() // Implicit Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("GetBlob did not cause boxing.")
	}

	o.Box()

	// Test serialization

	vb = koinos.NewVariableBlob()
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

	_, _, err2 := koinos.DeserializeOpaqueActiveBlockData(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewOpaqueActiveBlockData()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"koinos::protocol::active_block_data\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jo.Unbox()
	if !jo.IsBoxed() {
		t.Errorf("Unboxed incompatible serialization")
	}

	expected := []byte("{\"opaque\":{\"type\":\"koinos::protocol::active_block_data\",\"value\":\"zt1Zv2yaZ\"}}")
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

	jerr = json.Unmarshal([]byte("{\"previous_block\":{\"bougsagdg\":\"adgagf\"},\"transaction_merkle_root\":{\"bougsagdg\":\"adgagf\"},\"passive_data_merkle_root\":{\"bougsagdg\":\"adgagf\"},\"height\":{\"bougsagdg\":\"adgagf\"},\"timestamp\":{\"bougsagdg\":\"adgagf\"}}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	// Test alternative constructors
	vb = koinos.NewVariableBlob()
	o = koinos.NewOpaqueActiveBlockDataFromBlob(vb)

	if !o.IsBoxed() || !bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Create opaque from blob failed.")
	}

	slice := append([]byte(*vb), 0)
	vb = (*koinos.VariableBlob)(&slice)
	if bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Opaque blob pointer leaked")
	}

	n := koinos.NewActiveBlockData()
	o = koinos.NewOpaqueActiveBlockDataFromNative(*n)
	nativePtr, _ := o.GetNative()

	if o.IsBoxed() || nativePtr == n {
		t.Errorf("Create opaque from native failed.")
	}
}

// ----------------------------------------
//  OpaqueActiveTransactionData
// ----------------------------------------

func TestOpaqueActiveTransactionData(t *testing.T) {
	o := koinos.NewOpaqueActiveTransactionData()

	o.Box()
	if !o.IsBoxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if !o.IsBoxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if o.IsBoxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if o.IsBoxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.Box() // Call Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.Unbox()
	vb := o.GetBlob() // Implicit Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("GetBlob did not cause boxing.")
	}

	o.Box()

	// Test serialization

	vb = koinos.NewVariableBlob()
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

	_, _, err2 := koinos.DeserializeOpaqueActiveTransactionData(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewOpaqueActiveTransactionData()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"koinos::protocol::active_transaction_data\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jo.Unbox()
	if !jo.IsBoxed() {
		t.Errorf("Unboxed incompatible serialization")
	}

	expected := []byte("{\"opaque\":{\"type\":\"koinos::protocol::active_transaction_data\",\"value\":\"zt1Zv2yaZ\"}}")
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

	jerr = json.Unmarshal([]byte("{\"resource_limit\":{\"bougsagdg\":\"adgagf\"}}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	// Test alternative constructors
	vb = koinos.NewVariableBlob()
	o = koinos.NewOpaqueActiveTransactionDataFromBlob(vb)

	if !o.IsBoxed() || !bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Create opaque from blob failed.")
	}

	slice := append([]byte(*vb), 0)
	vb = (*koinos.VariableBlob)(&slice)
	if bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Opaque blob pointer leaked")
	}

	n := koinos.NewActiveTransactionData()
	o = koinos.NewOpaqueActiveTransactionDataFromNative(*n)
	nativePtr, _ := o.GetNative()

	if o.IsBoxed() || nativePtr == n {
		t.Errorf("Create opaque from native failed.")
	}
}

// ----------------------------------------
//  OpaqueBlock
// ----------------------------------------

func TestOpaqueBlock(t *testing.T) {
	o := koinos.NewOpaqueBlock()

	o.Box()
	if !o.IsBoxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if !o.IsBoxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if o.IsBoxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if o.IsBoxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.Box() // Call Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.Unbox()
	vb := o.GetBlob() // Implicit Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("GetBlob did not cause boxing.")
	}

	o.Box()

	// Test serialization

	vb = koinos.NewVariableBlob()
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

	_, _, err2 := koinos.DeserializeOpaqueBlock(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewOpaqueBlock()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"koinos::protocol::block\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jo.Unbox()
	if !jo.IsBoxed() {
		t.Errorf("Unboxed incompatible serialization")
	}

	expected := []byte("{\"opaque\":{\"type\":\"koinos::protocol::block\",\"value\":\"zt1Zv2yaZ\"}}")
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

	jerr = json.Unmarshal([]byte("{\"active_data\":{\"bougsagdg\":\"adgagf\"},\"passive_data\":{\"bougsagdg\":\"adgagf\"},\"signature_data\":{\"bougsagdg\":\"adgagf\"},\"transactions\":{\"bougsagdg\":\"adgagf\"}}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	// Test alternative constructors
	vb = koinos.NewVariableBlob()
	o = koinos.NewOpaqueBlockFromBlob(vb)

	if !o.IsBoxed() || !bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Create opaque from blob failed.")
	}

	slice := append([]byte(*vb), 0)
	vb = (*koinos.VariableBlob)(&slice)
	if bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Opaque blob pointer leaked")
	}

	n := koinos.NewBlock()
	o = koinos.NewOpaqueBlockFromNative(*n)
	nativePtr, _ := o.GetNative()

	if o.IsBoxed() || nativePtr == n {
		t.Errorf("Create opaque from native failed.")
	}
}

// ----------------------------------------
//  OpaqueBlockReceipt
// ----------------------------------------

func TestOpaqueBlockReceipt(t *testing.T) {
	o := koinos.NewOpaqueBlockReceipt()

	o.Box()
	if !o.IsBoxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if !o.IsBoxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if o.IsBoxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if o.IsBoxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.Box() // Call Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.Unbox()
	vb := o.GetBlob() // Implicit Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("GetBlob did not cause boxing.")
	}

	o.Box()

	// Test serialization

	vb = koinos.NewVariableBlob()
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

	_, _, err2 := koinos.DeserializeOpaqueBlockReceipt(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewOpaqueBlockReceipt()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"koinos::protocol::block_receipt\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jo.Unbox()
	if !jo.IsBoxed() {
		t.Errorf("Unboxed incompatible serialization")
	}

	expected := []byte("{\"opaque\":{\"type\":\"koinos::protocol::block_receipt\",\"value\":\"zt1Zv2yaZ\"}}")
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


	// Test alternative constructors
	vb = koinos.NewVariableBlob()
	o = koinos.NewOpaqueBlockReceiptFromBlob(vb)

	if !o.IsBoxed() || !bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Create opaque from blob failed.")
	}

	slice := append([]byte(*vb), 0)
	vb = (*koinos.VariableBlob)(&slice)
	if bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Opaque blob pointer leaked")
	}

	n := koinos.NewBlockReceipt()
	o = koinos.NewOpaqueBlockReceiptFromNative(*n)
	nativePtr, _ := o.GetNative()

	if o.IsBoxed() || nativePtr == n {
		t.Errorf("Create opaque from native failed.")
	}
}

// ----------------------------------------
//  OpaquePassiveBlockData
// ----------------------------------------

func TestOpaquePassiveBlockData(t *testing.T) {
	o := koinos.NewOpaquePassiveBlockData()

	o.Box()
	if !o.IsBoxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if !o.IsBoxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if o.IsBoxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if o.IsBoxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.Box() // Call Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.Unbox()
	vb := o.GetBlob() // Implicit Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("GetBlob did not cause boxing.")
	}

	o.Box()

	// Test serialization

	vb = koinos.NewVariableBlob()
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

	_, _, err2 := koinos.DeserializeOpaquePassiveBlockData(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewOpaquePassiveBlockData()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"koinos::protocol::passive_block_data\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jo.Unbox()
	if !jo.IsBoxed() {
		t.Errorf("Unboxed incompatible serialization")
	}

	expected := []byte("{\"opaque\":{\"type\":\"koinos::protocol::passive_block_data\",\"value\":\"zt1Zv2yaZ\"}}")
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


	// Test alternative constructors
	vb = koinos.NewVariableBlob()
	o = koinos.NewOpaquePassiveBlockDataFromBlob(vb)

	if !o.IsBoxed() || !bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Create opaque from blob failed.")
	}

	slice := append([]byte(*vb), 0)
	vb = (*koinos.VariableBlob)(&slice)
	if bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Opaque blob pointer leaked")
	}

	n := koinos.NewPassiveBlockData()
	o = koinos.NewOpaquePassiveBlockDataFromNative(*n)
	nativePtr, _ := o.GetNative()

	if o.IsBoxed() || nativePtr == n {
		t.Errorf("Create opaque from native failed.")
	}
}

// ----------------------------------------
//  OpaquePassiveTransactionData
// ----------------------------------------

func TestOpaquePassiveTransactionData(t *testing.T) {
	o := koinos.NewOpaquePassiveTransactionData()

	o.Box()
	if !o.IsBoxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if !o.IsBoxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if o.IsBoxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if o.IsBoxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.Box() // Call Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.Unbox()
	vb := o.GetBlob() // Implicit Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("GetBlob did not cause boxing.")
	}

	o.Box()

	// Test serialization

	vb = koinos.NewVariableBlob()
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

	_, _, err2 := koinos.DeserializeOpaquePassiveTransactionData(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewOpaquePassiveTransactionData()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"koinos::protocol::passive_transaction_data\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jo.Unbox()
	if !jo.IsBoxed() {
		t.Errorf("Unboxed incompatible serialization")
	}

	expected := []byte("{\"opaque\":{\"type\":\"koinos::protocol::passive_transaction_data\",\"value\":\"zt1Zv2yaZ\"}}")
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


	// Test alternative constructors
	vb = koinos.NewVariableBlob()
	o = koinos.NewOpaquePassiveTransactionDataFromBlob(vb)

	if !o.IsBoxed() || !bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Create opaque from blob failed.")
	}

	slice := append([]byte(*vb), 0)
	vb = (*koinos.VariableBlob)(&slice)
	if bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Opaque blob pointer leaked")
	}

	n := koinos.NewPassiveTransactionData()
	o = koinos.NewOpaquePassiveTransactionDataFromNative(*n)
	nativePtr, _ := o.GetNative()

	if o.IsBoxed() || nativePtr == n {
		t.Errorf("Create opaque from native failed.")
	}
}

// ----------------------------------------
//  OpaqueQueryItemResult
// ----------------------------------------

func TestOpaqueQueryItemResult(t *testing.T) {
	o := koinos.NewOpaqueQueryItemResult()

	o.Box()
	if !o.IsBoxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if !o.IsBoxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if o.IsBoxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if o.IsBoxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.Box() // Call Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.Unbox()
	vb := o.GetBlob() // Implicit Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("GetBlob did not cause boxing.")
	}

	o.Box()

	// Test serialization

	vb = koinos.NewVariableBlob()
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

	_, _, err2 := koinos.DeserializeOpaqueQueryItemResult(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewOpaqueQueryItemResult()
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
	if !jo.IsBoxed() {
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

	// Test alternative constructors
	vb = koinos.NewVariableBlob()
	o = koinos.NewOpaqueQueryItemResultFromBlob(vb)

	if !o.IsBoxed() || !bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Create opaque from blob failed.")
	}

	slice := append([]byte(*vb), 0)
	vb = (*koinos.VariableBlob)(&slice)
	if bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Opaque blob pointer leaked")
	}

	n := koinos.NewQueryItemResult()
	o = koinos.NewOpaqueQueryItemResultFromNative(*n)
	nativePtr, _ := o.GetNative()

	if o.IsBoxed() || nativePtr == n {
		t.Errorf("Create opaque from native failed.")
	}
}

// ----------------------------------------
//  OpaqueQueryParamItem
// ----------------------------------------

func TestOpaqueQueryParamItem(t *testing.T) {
	o := koinos.NewOpaqueQueryParamItem()

	o.Box()
	if !o.IsBoxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if !o.IsBoxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if o.IsBoxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if o.IsBoxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.Box() // Call Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.Unbox()
	vb := o.GetBlob() // Implicit Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("GetBlob did not cause boxing.")
	}

	o.Box()

	// Test serialization

	vb = koinos.NewVariableBlob()
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

	_, _, err2 := koinos.DeserializeOpaqueQueryParamItem(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewOpaqueQueryParamItem()
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
	if !jo.IsBoxed() {
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

	// Test alternative constructors
	vb = koinos.NewVariableBlob()
	o = koinos.NewOpaqueQueryParamItemFromBlob(vb)

	if !o.IsBoxed() || !bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Create opaque from blob failed.")
	}

	slice := append([]byte(*vb), 0)
	vb = (*koinos.VariableBlob)(&slice)
	if bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Opaque blob pointer leaked")
	}

	n := koinos.NewQueryParamItem()
	o = koinos.NewOpaqueQueryParamItemFromNative(*n)
	nativePtr, _ := o.GetNative()

	if o.IsBoxed() || nativePtr == n {
		t.Errorf("Create opaque from native failed.")
	}
}

// ----------------------------------------
//  OpaqueTransaction
// ----------------------------------------

func TestOpaqueTransaction(t *testing.T) {
	o := koinos.NewOpaqueTransaction()

	o.Box()
	if !o.IsBoxed() {
		t.Errorf("Opaque is unboxed but should not be.")
	}

	// Test getting native on Boxed
	_, err := o.GetNative()
	if err == nil {
		t.Errorf("Getting native on boxed should fail.")
	}

	o.Box() // Call Box() on Boxed
	if !o.IsBoxed() {
		t.Errorf("Boxed -> Boxed failed.")
	}

	o.Unbox() // Call Unbox() on Boxed
	if o.IsBoxed() {
		t.Errorf("Boxed -> Uboxed failed.")
	}

	// Test getting native on Unboxed
	_, err = o.GetNative()
	if err != nil {
		t.Errorf("Getting native on boxed should not fail.")
	}

	o.Unbox() // Call Unbox() on Unboxed
	if o.IsBoxed() {
		t.Errorf("Unboxed -> Unboxed failed.")
	}

	o.Box() // Call Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("Unboxed -> Boxed failed.")
	}

	o.Unbox()
	vb := o.GetBlob() // Implicit Box() on Unboxed
	if !o.IsBoxed() {
		t.Errorf("GetBlob did not cause boxing.")
	}

	o.Box()

	// Test serialization

	vb = koinos.NewVariableBlob()
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

	_, _, err2 := koinos.DeserializeOpaqueTransaction(vb)
	if err2 != nil {
		t.Error(err2)
	}

	v, jerr := json.Marshal(o)
	if jerr != nil {
		t.Error(jerr)
	}

	jo := koinos.NewOpaqueTransaction()
	jerr = json.Unmarshal(v, jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jerr = json.Unmarshal([]byte("\"!@#$%^&*\""), jo)
	if jerr == nil {
		t.Errorf("Unmarshaling nonsense JSON did not give error.")
	}

	jerr = json.Unmarshal([]byte("{\"opaque\":{\"type\":\"koinos::protocol::transaction\",\"value\":\"zt1Zv2yaZ\"}}"), jo)
	if jerr != nil {
		t.Error(jerr)
	}

	jo.Unbox()
	if !jo.IsBoxed() {
		t.Errorf("Unboxed incompatible serialization")
	}

	expected := []byte("{\"opaque\":{\"type\":\"koinos::protocol::transaction\",\"value\":\"zt1Zv2yaZ\"}}")
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

	jerr = json.Unmarshal([]byte("{\"active_data\":{\"bougsagdg\":\"adgagf\"},\"passive_data\":{\"bougsagdg\":\"adgagf\"},\"signature_data\":{\"bougsagdg\":\"adgagf\"},\"operations\":{\"bougsagdg\":\"adgagf\"}}"), jo)
	if jerr == nil {
		t.Errorf("jerr == nil")
	}

	// Test alternative constructors
	vb = koinos.NewVariableBlob()
	o = koinos.NewOpaqueTransactionFromBlob(vb)

	if !o.IsBoxed() || !bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Create opaque from blob failed.")
	}

	slice := append([]byte(*vb), 0)
	vb = (*koinos.VariableBlob)(&slice)
	if bytes.Equal([]byte(*vb), []byte(*o.GetBlob())) {
		t.Errorf("Opaque blob pointer leaked")
	}

	n := koinos.NewTransaction()
	o = koinos.NewOpaqueTransactionFromNative(*n)
	nativePtr, _ := o.GetNative()

	if o.IsBoxed() || nativePtr == n {
		t.Errorf("Create opaque from native failed.")
	}
}


// ----------------------------------------
//  VectorBlockItem
// ----------------------------------------

func TestVectorBlockItem(t *testing.T) {
	v := koinos.NewVectorBlockItem()
	for i := 0; i < 16; i++ {
		no := koinos.NewBlockItem()
		*v = append(*v, *no)
	}

	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, nv, err := koinos.DeserializeVectorBlockItem(vb)
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

	jv := koinos.NewVectorBlockItem()
	err = json.Unmarshal(j, &jv)
	if err != nil {
		t.Error(err)
	}

	// Test no data in the vector
	vb = &koinos.VariableBlob{0x01}
	n, _, err := koinos.DeserializeVectorBlockItem(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}
}

// ----------------------------------------
//  VectorBlockTopology
// ----------------------------------------

func TestVectorBlockTopology(t *testing.T) {
	v := koinos.NewVectorBlockTopology()
	for i := 0; i < 16; i++ {
		no := koinos.NewBlockTopology()
		*v = append(*v, *no)
	}

	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, nv, err := koinos.DeserializeVectorBlockTopology(vb)
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

	jv := koinos.NewVectorBlockTopology()
	err = json.Unmarshal(j, &jv)
	if err != nil {
		t.Error(err)
	}

	// Test no data in the vector
	vb = &koinos.VariableBlob{0x01}
	n, _, err := koinos.DeserializeVectorBlockTopology(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}
}

// ----------------------------------------
//  VectorMultihash
// ----------------------------------------

func TestVectorMultihash(t *testing.T) {
	v := koinos.NewVectorMultihash()
	for i := 0; i < 16; i++ {
		no := koinos.NewMultihash()
		*v = append(*v, *no)
	}

	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, nv, err := koinos.DeserializeVectorMultihash(vb)
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

	jv := koinos.NewVectorMultihash()
	err = json.Unmarshal(j, &jv)
	if err != nil {
		t.Error(err)
	}

	// Test no data in the vector
	vb = &koinos.VariableBlob{0x01}
	n, _, err := koinos.DeserializeVectorMultihash(vb)
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
	v := koinos.NewVectorOpaqueTransaction()
	for i := 0; i < 16; i++ {
		no := koinos.NewOpaqueTransaction()
		*v = append(*v, *no)
	}

	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, nv, err := koinos.DeserializeVectorOpaqueTransaction(vb)
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

	jv := koinos.NewVectorOpaqueTransaction()
	err = json.Unmarshal(j, &jv)
	if err != nil {
		t.Error(err)
	}

	// Test no data in the vector
	vb = &koinos.VariableBlob{0x01}
	n, _, err := koinos.DeserializeVectorOpaqueTransaction(vb)
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
	v := koinos.NewVectorOperation()
	for i := 0; i < 16; i++ {
		no := koinos.NewOperation()
		*v = append(*v, *no)
	}

	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, nv, err := koinos.DeserializeVectorOperation(vb)
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

	jv := koinos.NewVectorOperation()
	err = json.Unmarshal(j, &jv)
	if err != nil {
		t.Error(err)
	}

	// Test no data in the vector
	vb = &koinos.VariableBlob{0x01}
	n, _, err := koinos.DeserializeVectorOperation(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}
}

// ----------------------------------------
//  VectorTransaction
// ----------------------------------------

func TestVectorTransaction(t *testing.T) {
	v := koinos.NewVectorTransaction()
	for i := 0; i < 16; i++ {
		no := koinos.NewTransaction()
		*v = append(*v, *no)
	}

	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, nv, err := koinos.DeserializeVectorTransaction(vb)
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

	jv := koinos.NewVectorTransaction()
	err = json.Unmarshal(j, &jv)
	if err != nil {
		t.Error(err)
	}

	// Test no data in the vector
	vb = &koinos.VariableBlob{0x01}
	n, _, err := koinos.DeserializeVectorTransaction(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}
}

// ----------------------------------------
//  VectorTransactionItem
// ----------------------------------------

func TestVectorTransactionItem(t *testing.T) {
	v := koinos.NewVectorTransactionItem()
	for i := 0; i < 16; i++ {
		no := koinos.NewTransactionItem()
		*v = append(*v, *no)
	}

	vb := koinos.NewVariableBlob()
	vb = v.Serialize(vb)

	_, nv, err := koinos.DeserializeVectorTransactionItem(vb)
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

	jv := koinos.NewVectorTransactionItem()
	err = json.Unmarshal(j, &jv)
	if err != nil {
		t.Error(err)
	}

	// Test no data in the vector
	vb = &koinos.VariableBlob{0x01}
	n, _, err := koinos.DeserializeVectorTransactionItem(vb)
	if err == nil {
		t.Errorf("err == nil")
	}
	if n != 0 {
		t.Errorf("Bytes were consumed on error")
	}
}


