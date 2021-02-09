//   ____                           _           _    ____          _
//  / ___| ___ _ __   ___ _ __ __ _| |_ ___  __| |  / ___|___   __| | ___
// | |  _ / _ \ '_ \ / _ \ '__/ _` | __/ _ \/ _` | | |   / _ \ / _` |/ _ \
// | |_| |  __/ | | |  __/ | | (_| | ||  __/ (_| | | |__| (_) | (_| |  __/
//  \____|\___|_| |_|\___|_|  \__,_|\__\___|\__,_|  \____\___/ \__,_|\___|
//                         Please do not modify

package koinos

import (
	"fmt"
	"errors"
	"encoding/binary"
	"encoding/json"
	"strings"
)

// ----------------------------------------
//  Struct: UnusedExtensionsType
// ----------------------------------------

// UnusedExtensionsType type
type UnusedExtensionsType struct {
}

// NewUnusedExtensionsType factory
func NewUnusedExtensionsType() *UnusedExtensionsType {
	o := UnusedExtensionsType{}
	return &o
}

// Serialize UnusedExtensionsType
func (n UnusedExtensionsType) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeUnusedExtensionsType function
func DeserializeUnusedExtensionsType(vb *VariableBlob) (uint64,*UnusedExtensionsType,error) {
	var i uint64 = 0
	s := UnusedExtensionsType{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: ReservedOperation
// ----------------------------------------

// ReservedOperation type
type ReservedOperation struct {
    Extensions UnusedExtensionsType `json:"extensions"`
}

// NewReservedOperation factory
func NewReservedOperation() *ReservedOperation {
	o := ReservedOperation{}
	o.Extensions = *NewUnusedExtensionsType()
	return &o
}

// Serialize ReservedOperation
func (n ReservedOperation) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Extensions.Serialize(vb)
	return vb
}

// DeserializeReservedOperation function
func DeserializeReservedOperation(vb *VariableBlob) (uint64,*ReservedOperation,error) {
	var i uint64 = 0
	s := ReservedOperation{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: NopOperation
// ----------------------------------------

// NopOperation type
type NopOperation struct {
    Extensions UnusedExtensionsType `json:"extensions"`
}

// NewNopOperation factory
func NewNopOperation() *NopOperation {
	o := NopOperation{}
	o.Extensions = *NewUnusedExtensionsType()
	return &o
}

// Serialize NopOperation
func (n NopOperation) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Extensions.Serialize(vb)
	return vb
}

// DeserializeNopOperation function
func DeserializeNopOperation(vb *VariableBlob) (uint64,*NopOperation,error) {
	var i uint64 = 0
	s := NopOperation{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: ContractIDType
// ----------------------------------------

// ContractIDType type
type ContractIDType FixedBlob20

// NewContractIDType factory
func NewContractIDType() *ContractIDType {
	o := ContractIDType(*NewFixedBlob20())
	return &o
}

// Serialize ContractIDType
func (n ContractIDType) Serialize(vb *VariableBlob) *VariableBlob {
	ox := FixedBlob20(n)
	return ox.Serialize(vb)
}

// DeserializeContractIDType function
func DeserializeContractIDType(vb *VariableBlob) (uint64,*ContractIDType,error) {
	var ot ContractIDType
	i,n,err := DeserializeFixedBlob20(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = ContractIDType(*n)
	return i,&ot,nil}

// MarshalJSON ContractIDType
func (n ContractIDType) MarshalJSON() ([]byte, error) {
	v := FixedBlob20(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ContractIDType
func (n *ContractIDType) UnmarshalJSON(data []byte) error {
	v := FixedBlob20(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ContractIDType(v)
	return nil
}


// ----------------------------------------
//  Struct: CreateSystemContractOperation
// ----------------------------------------

// CreateSystemContractOperation type
type CreateSystemContractOperation struct {
    ContractID ContractIDType `json:"contract_id"`
    Bytecode VariableBlob `json:"bytecode"`
    Extensions UnusedExtensionsType `json:"extensions"`
}

// NewCreateSystemContractOperation factory
func NewCreateSystemContractOperation() *CreateSystemContractOperation {
	o := CreateSystemContractOperation{}
	o.ContractID = *NewContractIDType()
	o.Bytecode = *NewVariableBlob()
	o.Extensions = *NewUnusedExtensionsType()
	return &o
}

// Serialize CreateSystemContractOperation
func (n CreateSystemContractOperation) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ContractID.Serialize(vb)
	vb = n.Bytecode.Serialize(vb)
	vb = n.Extensions.Serialize(vb)
	return vb
}

// DeserializeCreateSystemContractOperation function
func DeserializeCreateSystemContractOperation(vb *VariableBlob) (uint64,*CreateSystemContractOperation,error) {
	var i,j uint64 = 0,0
	s := CreateSystemContractOperation{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tContractID,err := DeserializeContractIDType(&ovb); i+=j
	if err != nil {
		return 0, &CreateSystemContractOperation{}, err
	}
	s.ContractID = *tContractID
	ovb = (*vb)[i:]
	j,tBytecode,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &CreateSystemContractOperation{}, err
	}
	s.Bytecode = *tBytecode
	return i, &s, nil
}

// ----------------------------------------
//  Struct: SystemCallTargetReserved
// ----------------------------------------

// SystemCallTargetReserved type
type SystemCallTargetReserved struct {
}

// NewSystemCallTargetReserved factory
func NewSystemCallTargetReserved() *SystemCallTargetReserved {
	o := SystemCallTargetReserved{}
	return &o
}

// Serialize SystemCallTargetReserved
func (n SystemCallTargetReserved) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeSystemCallTargetReserved function
func DeserializeSystemCallTargetReserved(vb *VariableBlob) (uint64,*SystemCallTargetReserved,error) {
	var i uint64 = 0
	s := SystemCallTargetReserved{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Enum: ThunkID
// ----------------------------------------

// {'name': 'thunk_id', 'entries': [{'name': 'prints', 'value': 2406348109, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'verify_block_header', 'value': 2369936044, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_block', 'value': 2372743592, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_transaction', 'value': 2306978015, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_reserved_operation', 'value': 2335970550, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_upload_contract_operation', 'value': 2290263390, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_execute_contract_operation', 'value': 2246607595, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_set_system_call_operation', 'value': 2264476812, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_put_object', 'value': 2181271013, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_get_object', 'value': 2288165080, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_get_next_object', 'value': 2263109703, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_get_prev_object', 'value': 2371348733, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'execute_contract', 'value': 2319711875, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_contract_args_size', 'value': 2201456262, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_contract_args', 'value': 2383977862, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'set_contract_return', 'value': 2260230773, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'exit_contract', 'value': 2180390815, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_head_info', 'value': 2313106628, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'hash', 'value': 2326459719, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'verify_block_sig', 'value': 2300919863, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'verify_merkle_root', 'value': 2396642763, 'doc': '', 'info': {'type': 'EnumEntry'}}], 'tref': {'name': ['koinos', 'types', 'uint32'], 'targs': None, 'info': {'type': 'Typeref'}}, 'doc': '', 'info': {'type': 'EnumClass'}}

// ThunkID type
type ThunkID UInt32

// NewThunkID factory
func NewThunkID() *ThunkID {
	o := ThunkID(2406348109)
	return &o
}

// ThunkID values
const (
	ThunkIDPrints ThunkID = 2406348109
	ThunkIDVerifyBlockHeader ThunkID = 2369936044
	ThunkIDApplyBlock ThunkID = 2372743592
	ThunkIDApplyTransaction ThunkID = 2306978015
	ThunkIDApplyReservedOperation ThunkID = 2335970550
	ThunkIDApplyUploadContractOperation ThunkID = 2290263390
	ThunkIDApplyExecuteContractOperation ThunkID = 2246607595
	ThunkIDApplySetSystemCallOperation ThunkID = 2264476812
	ThunkIDDbPutObject ThunkID = 2181271013
	ThunkIDDbGetObject ThunkID = 2288165080
	ThunkIDDbGetNextObject ThunkID = 2263109703
	ThunkIDDbGetPrevObject ThunkID = 2371348733
	ThunkIDExecuteContract ThunkID = 2319711875
	ThunkIDGetContractArgsSize ThunkID = 2201456262
	ThunkIDGetContractArgs ThunkID = 2383977862
	ThunkIDSetContractReturn ThunkID = 2260230773
	ThunkIDExitContract ThunkID = 2180390815
	ThunkIDGetHeadInfo ThunkID = 2313106628
	ThunkIDHash ThunkID = 2326459719
	ThunkIDVerifyBlockSig ThunkID = 2300919863
	ThunkIDVerifyMerkleRoot ThunkID = 2396642763
)

// Serialize ThunkID
func (n ThunkID) Serialize(vb *VariableBlob) *VariableBlob {
	if !IsValidThunkID(n) {
		panic("Attempting to serialize an invalid value")
	}
	x := UInt32(n)
	return x.Serialize(vb)
}

// DeserializeThunkID function
func DeserializeThunkID(vb *VariableBlob) (uint64,*ThunkID,error) {
	i,item,err := DeserializeUInt32(vb)
	var x ThunkID
	if err != nil {
		return 0,&x,err
	}

	x = ThunkID(*item)
	if !IsValidThunkID(x) {
		return i,&x,fmt.Errorf("invalid ThunkID: %d", x)
	}
	return i,&x,nil
}

// MarshalJSON ThunkID
func (n ThunkID) MarshalJSON() ([]byte, error) {
	if !IsValidThunkID(n) {
		panic("Attempting to serialize an invalid value")
	}

	return json.Marshal(UInt32(n))
}

// UnmarshalJSON *ThunkID
func (n *ThunkID) UnmarshalJSON(b []byte) error {
	var o UInt32
	if err := json.Unmarshal(b, &o); err != nil {
		return err
	}

	ov := ThunkID(o)

	if !IsValidThunkID(ov) {
		return fmt.Errorf("invalid ThunkID: %d", o)
	}

	*n = ov
	return nil
}

// IsValidThunkID validator
func IsValidThunkID(v ThunkID) bool {
	switch v {
		case ThunkIDPrints:
			return true
		case ThunkIDVerifyBlockHeader:
			return true
		case ThunkIDApplyBlock:
			return true
		case ThunkIDApplyTransaction:
			return true
		case ThunkIDApplyReservedOperation:
			return true
		case ThunkIDApplyUploadContractOperation:
			return true
		case ThunkIDApplyExecuteContractOperation:
			return true
		case ThunkIDApplySetSystemCallOperation:
			return true
		case ThunkIDDbPutObject:
			return true
		case ThunkIDDbGetObject:
			return true
		case ThunkIDDbGetNextObject:
			return true
		case ThunkIDDbGetPrevObject:
			return true
		case ThunkIDExecuteContract:
			return true
		case ThunkIDGetContractArgsSize:
			return true
		case ThunkIDGetContractArgs:
			return true
		case ThunkIDSetContractReturn:
			return true
		case ThunkIDExitContract:
			return true
		case ThunkIDGetHeadInfo:
			return true
		case ThunkIDHash:
			return true
		case ThunkIDVerifyBlockSig:
			return true
		case ThunkIDVerifyMerkleRoot:
			return true
		default:
			return false
	}
}


// ----------------------------------------
//  Struct: ContractCallBundle
// ----------------------------------------

// ContractCallBundle type
type ContractCallBundle struct {
    ContractID ContractIDType `json:"contract_id"`
    EntryPoint UInt32 `json:"entry_point"`
}

// NewContractCallBundle factory
func NewContractCallBundle() *ContractCallBundle {
	o := ContractCallBundle{}
	o.ContractID = *NewContractIDType()
	o.EntryPoint = *NewUInt32()
	return &o
}

// Serialize ContractCallBundle
func (n ContractCallBundle) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ContractID.Serialize(vb)
	vb = n.EntryPoint.Serialize(vb)
	return vb
}

// DeserializeContractCallBundle function
func DeserializeContractCallBundle(vb *VariableBlob) (uint64,*ContractCallBundle,error) {
	var i,j uint64 = 0,0
	s := ContractCallBundle{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tContractID,err := DeserializeContractIDType(&ovb); i+=j
	if err != nil {
		return 0, &ContractCallBundle{}, err
	}
	s.ContractID = *tContractID
	ovb = (*vb)[i:]
	j,tEntryPoint,err := DeserializeUInt32(&ovb); i+=j
	if err != nil {
		return 0, &ContractCallBundle{}, err
	}
	s.EntryPoint = *tEntryPoint
	return i, &s, nil
}

// ----------------------------------------
//  Variant: SystemCallTarget
// ----------------------------------------

// SystemCallTarget type
type SystemCallTarget struct {
	Value interface{}
}

// NewSystemCallTarget factory
func NewSystemCallTarget() *SystemCallTarget {
	v := SystemCallTarget{}
	v.Value = NewSystemCallTargetReserved()
	return &v
}

// Serialize SystemCallTarget
func (n SystemCallTarget) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *SystemCallTargetReserved:
			i = 0
		case *ThunkID:
			i = 1
		case *ContractCallBundle:
			i = 2
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName SystemCallTarget
func (n SystemCallTarget) TypeToName() (string) {
	switch n.Value.(type) {
		case *SystemCallTargetReserved:
			return "koinos::types::system::system_call_target_reserved"
		case *ThunkID:
			return "koinos::types::thunks::thunk_id"
		case *ContractCallBundle:
			return "koinos::types::system::contract_call_bundle"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON SystemCallTarget
func (n SystemCallTarget) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeSystemCallTarget function
func DeserializeSystemCallTarget(vb *VariableBlob) (uint64,*SystemCallTarget,error) {
	var v SystemCallTarget
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewSystemCallTargetReserved()
		case 1:
			ovb := (*vb)[i:]
			k,x,err := DeserializeThunkID(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 2:
			ovb := (*vb)[i:]
			k,x,err := DeserializeContractCallBundle(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		default:
			return 0, &v, errors.New("unknown variant tag")
	}
	return uint64(i)+j,&v,nil
}

// UnmarshalJSON *SystemCallTarget
func (n *SystemCallTarget) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::types::system::system_call_target_reserved":
			v := NewSystemCallTargetReserved()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::thunks::thunk_id":
			v := NewThunkID()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::system::contract_call_bundle":
			v := NewContractCallBundle()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

	return nil
}


// ----------------------------------------
//  Struct: SetSystemCallOperation
// ----------------------------------------

// SetSystemCallOperation type
type SetSystemCallOperation struct {
    CallID UInt32 `json:"call_id"`
    Target SystemCallTarget `json:"target"`
    Extensions UnusedExtensionsType `json:"extensions"`
}

// NewSetSystemCallOperation factory
func NewSetSystemCallOperation() *SetSystemCallOperation {
	o := SetSystemCallOperation{}
	o.CallID = *NewUInt32()
	o.Target = *NewSystemCallTarget()
	o.Extensions = *NewUnusedExtensionsType()
	return &o
}

// Serialize SetSystemCallOperation
func (n SetSystemCallOperation) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.CallID.Serialize(vb)
	vb = n.Target.Serialize(vb)
	vb = n.Extensions.Serialize(vb)
	return vb
}

// DeserializeSetSystemCallOperation function
func DeserializeSetSystemCallOperation(vb *VariableBlob) (uint64,*SetSystemCallOperation,error) {
	var i,j uint64 = 0,0
	s := SetSystemCallOperation{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tCallID,err := DeserializeUInt32(&ovb); i+=j
	if err != nil {
		return 0, &SetSystemCallOperation{}, err
	}
	s.CallID = *tCallID
	ovb = (*vb)[i:]
	j,tTarget,err := DeserializeSystemCallTarget(&ovb); i+=j
	if err != nil {
		return 0, &SetSystemCallOperation{}, err
	}
	s.Target = *tTarget
	return i, &s, nil
}

// ----------------------------------------
//  Struct: ContractCallOperation
// ----------------------------------------

// ContractCallOperation type
type ContractCallOperation struct {
    ContractID ContractIDType `json:"contract_id"`
    EntryPoint UInt32 `json:"entry_point"`
    Args VariableBlob `json:"args"`
    Extensions UnusedExtensionsType `json:"extensions"`
}

// NewContractCallOperation factory
func NewContractCallOperation() *ContractCallOperation {
	o := ContractCallOperation{}
	o.ContractID = *NewContractIDType()
	o.EntryPoint = *NewUInt32()
	o.Args = *NewVariableBlob()
	o.Extensions = *NewUnusedExtensionsType()
	return &o
}

// Serialize ContractCallOperation
func (n ContractCallOperation) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ContractID.Serialize(vb)
	vb = n.EntryPoint.Serialize(vb)
	vb = n.Args.Serialize(vb)
	vb = n.Extensions.Serialize(vb)
	return vb
}

// DeserializeContractCallOperation function
func DeserializeContractCallOperation(vb *VariableBlob) (uint64,*ContractCallOperation,error) {
	var i,j uint64 = 0,0
	s := ContractCallOperation{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tContractID,err := DeserializeContractIDType(&ovb); i+=j
	if err != nil {
		return 0, &ContractCallOperation{}, err
	}
	s.ContractID = *tContractID
	ovb = (*vb)[i:]
	j,tEntryPoint,err := DeserializeUInt32(&ovb); i+=j
	if err != nil {
		return 0, &ContractCallOperation{}, err
	}
	s.EntryPoint = *tEntryPoint
	ovb = (*vb)[i:]
	j,tArgs,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &ContractCallOperation{}, err
	}
	s.Args = *tArgs
	return i, &s, nil
}

// ----------------------------------------
//  Variant: Operation
// ----------------------------------------

// Operation type
type Operation struct {
	Value interface{}
}

// NewOperation factory
func NewOperation() *Operation {
	v := Operation{}
	v.Value = NewReservedOperation()
	return &v
}

// Serialize Operation
func (n Operation) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *ReservedOperation:
			i = 0
		case *NopOperation:
			i = 1
		case *CreateSystemContractOperation:
			i = 2
		case *ContractCallOperation:
			i = 3
		case *SetSystemCallOperation:
			i = 4
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName Operation
func (n Operation) TypeToName() (string) {
	switch n.Value.(type) {
		case *ReservedOperation:
			return "koinos::types::protocol::reserved_operation"
		case *NopOperation:
			return "koinos::types::protocol::nop_operation"
		case *CreateSystemContractOperation:
			return "koinos::types::protocol::create_system_contract_operation"
		case *ContractCallOperation:
			return "koinos::types::protocol::contract_call_operation"
		case *SetSystemCallOperation:
			return "koinos::types::protocol::set_system_call_operation"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON Operation
func (n Operation) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeOperation function
func DeserializeOperation(vb *VariableBlob) (uint64,*Operation,error) {
	var v Operation
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewReservedOperation()
		case 1:
			v.Value = NewNopOperation()
		case 2:
			ovb := (*vb)[i:]
			k,x,err := DeserializeCreateSystemContractOperation(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 3:
			ovb := (*vb)[i:]
			k,x,err := DeserializeContractCallOperation(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 4:
			ovb := (*vb)[i:]
			k,x,err := DeserializeSetSystemCallOperation(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		default:
			return 0, &v, errors.New("unknown variant tag")
	}
	return uint64(i)+j,&v,nil
}

// UnmarshalJSON *Operation
func (n *Operation) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::types::protocol::reserved_operation":
			v := NewReservedOperation()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::protocol::nop_operation":
			v := NewNopOperation()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::protocol::create_system_contract_operation":
			v := NewCreateSystemContractOperation()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::protocol::contract_call_operation":
			v := NewContractCallOperation()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::protocol::set_system_call_operation":
			v := NewSetSystemCallOperation()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

	return nil
}


// ----------------------------------------
//  Struct: ActiveTransactionData
// ----------------------------------------

// ActiveTransactionData type
type ActiveTransactionData struct {
}

// NewActiveTransactionData factory
func NewActiveTransactionData() *ActiveTransactionData {
	o := ActiveTransactionData{}
	return &o
}

// Serialize ActiveTransactionData
func (n ActiveTransactionData) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeActiveTransactionData function
func DeserializeActiveTransactionData(vb *VariableBlob) (uint64,*ActiveTransactionData,error) {
	var i uint64 = 0
	s := ActiveTransactionData{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: PassiveTransactionData
// ----------------------------------------

// PassiveTransactionData type
type PassiveTransactionData struct {
}

// NewPassiveTransactionData factory
func NewPassiveTransactionData() *PassiveTransactionData {
	o := PassiveTransactionData{}
	return &o
}

// Serialize PassiveTransactionData
func (n PassiveTransactionData) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializePassiveTransactionData function
func DeserializePassiveTransactionData(vb *VariableBlob) (uint64,*PassiveTransactionData,error) {
	var i uint64 = 0
	s := PassiveTransactionData{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: Transaction
// ----------------------------------------

// Transaction type
type Transaction struct {
    ActiveData OpaqueActiveTransactionData `json:"active_data"`
    PassiveData OpaquePassiveTransactionData `json:"passive_data"`
    SignatureData VariableBlob `json:"signature_data"`
    Operations VectorOperation `json:"operations"`
}

// NewTransaction factory
func NewTransaction() *Transaction {
	o := Transaction{}
	o.ActiveData = *NewOpaqueActiveTransactionData()
	o.PassiveData = *NewOpaquePassiveTransactionData()
	o.SignatureData = *NewVariableBlob()
	o.Operations = *NewVectorOperation()
	return &o
}

// Serialize Transaction
func (n Transaction) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ActiveData.Serialize(vb)
	vb = n.PassiveData.Serialize(vb)
	vb = n.SignatureData.Serialize(vb)
	vb = n.Operations.Serialize(vb)
	return vb
}

// DeserializeTransaction function
func DeserializeTransaction(vb *VariableBlob) (uint64,*Transaction,error) {
	var i,j uint64 = 0,0
	s := Transaction{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tActiveData,err := DeserializeOpaqueActiveTransactionData(&ovb); i+=j
	if err != nil {
		return 0, &Transaction{}, err
	}
	s.ActiveData = *tActiveData
	ovb = (*vb)[i:]
	j,tPassiveData,err := DeserializeOpaquePassiveTransactionData(&ovb); i+=j
	if err != nil {
		return 0, &Transaction{}, err
	}
	s.PassiveData = *tPassiveData
	ovb = (*vb)[i:]
	j,tSignatureData,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &Transaction{}, err
	}
	s.SignatureData = *tSignatureData
	ovb = (*vb)[i:]
	j,tOperations,err := DeserializeVectorOperation(&ovb); i+=j
	if err != nil {
		return 0, &Transaction{}, err
	}
	s.Operations = *tOperations
	return i, &s, nil
}

// ----------------------------------------
//  Enum: HeaderHashIndex
// ----------------------------------------

// {'name': 'header_hash_index', 'entries': [{'name': 'previous_block_hash_index', 'value': 0, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'transaction_merkle_root_hash_index', 'value': 1, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'passive_data_merkle_root_hash_index', 'value': 2, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'NUM_HEADER_HASHES', 'value': 3, 'doc': '', 'info': {'type': 'EnumEntry'}}], 'tref': {'name': ['koinos', 'types', 'uint32'], 'targs': None, 'info': {'type': 'Typeref'}}, 'doc': '/**\n    * Number of header hashes.\n    */', 'info': {'type': 'EnumClass'}}

// HeaderHashIndex type
type HeaderHashIndex UInt32

// NewHeaderHashIndex factory
func NewHeaderHashIndex() *HeaderHashIndex {
	o := HeaderHashIndex(0)
	return &o
}

// HeaderHashIndex values
const (
	HeaderHashIndexPreviousBlockHashIndex HeaderHashIndex = 0
	HeaderHashIndexTransactionMerkleRootHashIndex HeaderHashIndex = 1
	HeaderHashIndexPassiveDataMerkleRootHashIndex HeaderHashIndex = 2
	HeaderHashIndexNumHeaderHashes HeaderHashIndex = 3
)

// Serialize HeaderHashIndex
func (n HeaderHashIndex) Serialize(vb *VariableBlob) *VariableBlob {
	if !IsValidHeaderHashIndex(n) {
		panic("Attempting to serialize an invalid value")
	}
	x := UInt32(n)
	return x.Serialize(vb)
}

// DeserializeHeaderHashIndex function
func DeserializeHeaderHashIndex(vb *VariableBlob) (uint64,*HeaderHashIndex,error) {
	i,item,err := DeserializeUInt32(vb)
	var x HeaderHashIndex
	if err != nil {
		return 0,&x,err
	}

	x = HeaderHashIndex(*item)
	if !IsValidHeaderHashIndex(x) {
		return i,&x,fmt.Errorf("invalid HeaderHashIndex: %d", x)
	}
	return i,&x,nil
}

// MarshalJSON HeaderHashIndex
func (n HeaderHashIndex) MarshalJSON() ([]byte, error) {
	if !IsValidHeaderHashIndex(n) {
		panic("Attempting to serialize an invalid value")
	}

	return json.Marshal(UInt32(n))
}

// UnmarshalJSON *HeaderHashIndex
func (n *HeaderHashIndex) UnmarshalJSON(b []byte) error {
	var o UInt32
	if err := json.Unmarshal(b, &o); err != nil {
		return err
	}

	ov := HeaderHashIndex(o)

	if !IsValidHeaderHashIndex(ov) {
		return fmt.Errorf("invalid HeaderHashIndex: %d", o)
	}

	*n = ov
	return nil
}

// IsValidHeaderHashIndex validator
func IsValidHeaderHashIndex(v HeaderHashIndex) bool {
	switch v {
		case HeaderHashIndexPreviousBlockHashIndex:
			return true
		case HeaderHashIndexTransactionMerkleRootHashIndex:
			return true
		case HeaderHashIndexPassiveDataMerkleRootHashIndex:
			return true
		case HeaderHashIndexNumHeaderHashes:
			return true
		default:
			return false
	}
}


// ----------------------------------------
//  Struct: ActiveBlockData
// ----------------------------------------

// ActiveBlockData type
type ActiveBlockData struct {
    HeaderHashes MultihashVector `json:"header_hashes"`
    Height BlockHeightType `json:"height"`
    Timestamp TimestampType `json:"timestamp"`
}

// NewActiveBlockData factory
func NewActiveBlockData() *ActiveBlockData {
	o := ActiveBlockData{}
	o.HeaderHashes = *NewMultihashVector()
	o.Height = *NewBlockHeightType()
	o.Timestamp = *NewTimestampType()
	return &o
}

// Serialize ActiveBlockData
func (n ActiveBlockData) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.HeaderHashes.Serialize(vb)
	vb = n.Height.Serialize(vb)
	vb = n.Timestamp.Serialize(vb)
	return vb
}

// DeserializeActiveBlockData function
func DeserializeActiveBlockData(vb *VariableBlob) (uint64,*ActiveBlockData,error) {
	var i,j uint64 = 0,0
	s := ActiveBlockData{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tHeaderHashes,err := DeserializeMultihashVector(&ovb); i+=j
	if err != nil {
		return 0, &ActiveBlockData{}, err
	}
	s.HeaderHashes = *tHeaderHashes
	ovb = (*vb)[i:]
	j,tHeight,err := DeserializeBlockHeightType(&ovb); i+=j
	if err != nil {
		return 0, &ActiveBlockData{}, err
	}
	s.Height = *tHeight
	ovb = (*vb)[i:]
	j,tTimestamp,err := DeserializeTimestampType(&ovb); i+=j
	if err != nil {
		return 0, &ActiveBlockData{}, err
	}
	s.Timestamp = *tTimestamp
	return i, &s, nil
}

// ----------------------------------------
//  Struct: PassiveBlockData
// ----------------------------------------

// PassiveBlockData type
type PassiveBlockData struct {
}

// NewPassiveBlockData factory
func NewPassiveBlockData() *PassiveBlockData {
	o := PassiveBlockData{}
	return &o
}

// Serialize PassiveBlockData
func (n PassiveBlockData) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializePassiveBlockData function
func DeserializePassiveBlockData(vb *VariableBlob) (uint64,*PassiveBlockData,error) {
	var i uint64 = 0
	s := PassiveBlockData{}
	
	return i, &s, nil
}




// ----------------------------------------
//  Struct: Block
// ----------------------------------------

// Block type
type Block struct {
    ActiveData OpaqueActiveBlockData `json:"active_data"`
    PassiveData OpaquePassiveBlockData `json:"passive_data"`
    SignatureData VariableBlob `json:"signature_data"`
    Transactions VectorOpaqueTransaction `json:"transactions"`
}

// NewBlock factory
func NewBlock() *Block {
	o := Block{}
	o.ActiveData = *NewOpaqueActiveBlockData()
	o.PassiveData = *NewOpaquePassiveBlockData()
	o.SignatureData = *NewVariableBlob()
	o.Transactions = *NewVectorOpaqueTransaction()
	return &o
}

// Serialize Block
func (n Block) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ActiveData.Serialize(vb)
	vb = n.PassiveData.Serialize(vb)
	vb = n.SignatureData.Serialize(vb)
	vb = n.Transactions.Serialize(vb)
	return vb
}

// DeserializeBlock function
func DeserializeBlock(vb *VariableBlob) (uint64,*Block,error) {
	var i,j uint64 = 0,0
	s := Block{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tActiveData,err := DeserializeOpaqueActiveBlockData(&ovb); i+=j
	if err != nil {
		return 0, &Block{}, err
	}
	s.ActiveData = *tActiveData
	ovb = (*vb)[i:]
	j,tPassiveData,err := DeserializeOpaquePassiveBlockData(&ovb); i+=j
	if err != nil {
		return 0, &Block{}, err
	}
	s.PassiveData = *tPassiveData
	ovb = (*vb)[i:]
	j,tSignatureData,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &Block{}, err
	}
	s.SignatureData = *tSignatureData
	ovb = (*vb)[i:]
	j,tTransactions,err := DeserializeVectorOpaqueTransaction(&ovb); i+=j
	if err != nil {
		return 0, &Block{}, err
	}
	s.Transactions = *tTransactions
	return i, &s, nil
}

// ----------------------------------------
//  Struct: BlockReceipt
// ----------------------------------------

// BlockReceipt type
type BlockReceipt struct {
}

// NewBlockReceipt factory
func NewBlockReceipt() *BlockReceipt {
	o := BlockReceipt{}
	return &o
}

// Serialize BlockReceipt
func (n BlockReceipt) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeBlockReceipt function
func DeserializeBlockReceipt(vb *VariableBlob) (uint64,*BlockReceipt,error) {
	var i uint64 = 0
	s := BlockReceipt{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: ReservedReq
// ----------------------------------------

// ReservedReq type
type ReservedReq struct {
}

// NewReservedReq factory
func NewReservedReq() *ReservedReq {
	o := ReservedReq{}
	return &o
}

// Serialize ReservedReq
func (n ReservedReq) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeReservedReq function
func DeserializeReservedReq(vb *VariableBlob) (uint64,*ReservedReq,error) {
	var i uint64 = 0
	s := ReservedReq{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: ReservedResp
// ----------------------------------------

// ReservedResp type
type ReservedResp struct {
}

// NewReservedResp factory
func NewReservedResp() *ReservedResp {
	o := ReservedResp{}
	return &o
}

// Serialize ReservedResp
func (n ReservedResp) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeReservedResp function
func DeserializeReservedResp(vb *VariableBlob) (uint64,*ReservedResp,error) {
	var i uint64 = 0
	s := ReservedResp{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetBlocksByIDReq
// ----------------------------------------

// GetBlocksByIDReq type
type GetBlocksByIDReq struct {
    BlockID VectorMultihash `json:"block_id"`
    ReturnBlockBlob Boolean `json:"return_block_blob"`
    ReturnReceiptBlob Boolean `json:"return_receipt_blob"`
}

// NewGetBlocksByIDReq factory
func NewGetBlocksByIDReq() *GetBlocksByIDReq {
	o := GetBlocksByIDReq{}
	o.BlockID = *NewVectorMultihash()
	o.ReturnBlockBlob = *NewBoolean()
	o.ReturnReceiptBlob = *NewBoolean()
	return &o
}

// Serialize GetBlocksByIDReq
func (n GetBlocksByIDReq) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.BlockID.Serialize(vb)
	vb = n.ReturnBlockBlob.Serialize(vb)
	vb = n.ReturnReceiptBlob.Serialize(vb)
	return vb
}

// DeserializeGetBlocksByIDReq function
func DeserializeGetBlocksByIDReq(vb *VariableBlob) (uint64,*GetBlocksByIDReq,error) {
	var i,j uint64 = 0,0
	s := GetBlocksByIDReq{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tBlockID,err := DeserializeVectorMultihash(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByIDReq{}, err
	}
	s.BlockID = *tBlockID
	ovb = (*vb)[i:]
	j,tReturnBlockBlob,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByIDReq{}, err
	}
	s.ReturnBlockBlob = *tReturnBlockBlob
	ovb = (*vb)[i:]
	j,tReturnReceiptBlob,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByIDReq{}, err
	}
	s.ReturnReceiptBlob = *tReturnReceiptBlob
	return i, &s, nil
}

// ----------------------------------------
//  Struct: BlockItem
// ----------------------------------------

// BlockItem type
type BlockItem struct {
    BlockID Multihash `json:"block_id"`
    BlockHeight BlockHeightType `json:"block_height"`
    Block OpaqueBlock `json:"block"`
    BlockReceipt OpaqueBlockReceipt `json:"block_receipt"`
}

// NewBlockItem factory
func NewBlockItem() *BlockItem {
	o := BlockItem{}
	o.BlockID = *NewMultihash()
	o.BlockHeight = *NewBlockHeightType()
	o.Block = *NewOpaqueBlock()
	o.BlockReceipt = *NewOpaqueBlockReceipt()
	return &o
}

// Serialize BlockItem
func (n BlockItem) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.BlockID.Serialize(vb)
	vb = n.BlockHeight.Serialize(vb)
	vb = n.Block.Serialize(vb)
	vb = n.BlockReceipt.Serialize(vb)
	return vb
}

// DeserializeBlockItem function
func DeserializeBlockItem(vb *VariableBlob) (uint64,*BlockItem,error) {
	var i,j uint64 = 0,0
	s := BlockItem{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tBlockID,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &BlockItem{}, err
	}
	s.BlockID = *tBlockID
	ovb = (*vb)[i:]
	j,tBlockHeight,err := DeserializeBlockHeightType(&ovb); i+=j
	if err != nil {
		return 0, &BlockItem{}, err
	}
	s.BlockHeight = *tBlockHeight
	ovb = (*vb)[i:]
	j,tBlock,err := DeserializeOpaqueBlock(&ovb); i+=j
	if err != nil {
		return 0, &BlockItem{}, err
	}
	s.Block = *tBlock
	ovb = (*vb)[i:]
	j,tBlockReceipt,err := DeserializeOpaqueBlockReceipt(&ovb); i+=j
	if err != nil {
		return 0, &BlockItem{}, err
	}
	s.BlockReceipt = *tBlockReceipt
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetBlocksByIDResp
// ----------------------------------------

// GetBlocksByIDResp type
type GetBlocksByIDResp struct {
    BlockItems VectorBlockItem `json:"block_items"`
}

// NewGetBlocksByIDResp factory
func NewGetBlocksByIDResp() *GetBlocksByIDResp {
	o := GetBlocksByIDResp{}
	o.BlockItems = *NewVectorBlockItem()
	return &o
}

// Serialize GetBlocksByIDResp
func (n GetBlocksByIDResp) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.BlockItems.Serialize(vb)
	return vb
}

// DeserializeGetBlocksByIDResp function
func DeserializeGetBlocksByIDResp(vb *VariableBlob) (uint64,*GetBlocksByIDResp,error) {
	var i,j uint64 = 0,0
	s := GetBlocksByIDResp{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tBlockItems,err := DeserializeVectorBlockItem(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByIDResp{}, err
	}
	s.BlockItems = *tBlockItems
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetBlocksByHeightReq
// ----------------------------------------

// GetBlocksByHeightReq type
type GetBlocksByHeightReq struct {
    HeadBlockID Multihash `json:"head_block_id"`
    AncestorStartHeight BlockHeightType `json:"ancestor_start_height"`
    NumBlocks UInt32 `json:"num_blocks"`
    ReturnBlock Boolean `json:"return_block"`
    ReturnReceipt Boolean `json:"return_receipt"`
}

// NewGetBlocksByHeightReq factory
func NewGetBlocksByHeightReq() *GetBlocksByHeightReq {
	o := GetBlocksByHeightReq{}
	o.HeadBlockID = *NewMultihash()
	o.AncestorStartHeight = *NewBlockHeightType()
	o.NumBlocks = *NewUInt32()
	o.ReturnBlock = *NewBoolean()
	o.ReturnReceipt = *NewBoolean()
	return &o
}

// Serialize GetBlocksByHeightReq
func (n GetBlocksByHeightReq) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.HeadBlockID.Serialize(vb)
	vb = n.AncestorStartHeight.Serialize(vb)
	vb = n.NumBlocks.Serialize(vb)
	vb = n.ReturnBlock.Serialize(vb)
	vb = n.ReturnReceipt.Serialize(vb)
	return vb
}

// DeserializeGetBlocksByHeightReq function
func DeserializeGetBlocksByHeightReq(vb *VariableBlob) (uint64,*GetBlocksByHeightReq,error) {
	var i,j uint64 = 0,0
	s := GetBlocksByHeightReq{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tHeadBlockID,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByHeightReq{}, err
	}
	s.HeadBlockID = *tHeadBlockID
	ovb = (*vb)[i:]
	j,tAncestorStartHeight,err := DeserializeBlockHeightType(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByHeightReq{}, err
	}
	s.AncestorStartHeight = *tAncestorStartHeight
	ovb = (*vb)[i:]
	j,tNumBlocks,err := DeserializeUInt32(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByHeightReq{}, err
	}
	s.NumBlocks = *tNumBlocks
	ovb = (*vb)[i:]
	j,tReturnBlock,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByHeightReq{}, err
	}
	s.ReturnBlock = *tReturnBlock
	ovb = (*vb)[i:]
	j,tReturnReceipt,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByHeightReq{}, err
	}
	s.ReturnReceipt = *tReturnReceipt
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetBlocksByHeightResp
// ----------------------------------------

// GetBlocksByHeightResp type
type GetBlocksByHeightResp struct {
    BlockItems VectorBlockItem `json:"block_items"`
}

// NewGetBlocksByHeightResp factory
func NewGetBlocksByHeightResp() *GetBlocksByHeightResp {
	o := GetBlocksByHeightResp{}
	o.BlockItems = *NewVectorBlockItem()
	return &o
}

// Serialize GetBlocksByHeightResp
func (n GetBlocksByHeightResp) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.BlockItems.Serialize(vb)
	return vb
}

// DeserializeGetBlocksByHeightResp function
func DeserializeGetBlocksByHeightResp(vb *VariableBlob) (uint64,*GetBlocksByHeightResp,error) {
	var i,j uint64 = 0,0
	s := GetBlocksByHeightResp{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tBlockItems,err := DeserializeVectorBlockItem(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByHeightResp{}, err
	}
	s.BlockItems = *tBlockItems
	return i, &s, nil
}

// ----------------------------------------
//  Struct: AddBlockReq
// ----------------------------------------

// AddBlockReq type
type AddBlockReq struct {
    BlockToAdd BlockItem `json:"block_to_add"`
    PreviousBlockID Multihash `json:"previous_block_id"`
}

// NewAddBlockReq factory
func NewAddBlockReq() *AddBlockReq {
	o := AddBlockReq{}
	o.BlockToAdd = *NewBlockItem()
	o.PreviousBlockID = *NewMultihash()
	return &o
}

// Serialize AddBlockReq
func (n AddBlockReq) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.BlockToAdd.Serialize(vb)
	vb = n.PreviousBlockID.Serialize(vb)
	return vb
}

// DeserializeAddBlockReq function
func DeserializeAddBlockReq(vb *VariableBlob) (uint64,*AddBlockReq,error) {
	var i,j uint64 = 0,0
	s := AddBlockReq{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tBlockToAdd,err := DeserializeBlockItem(&ovb); i+=j
	if err != nil {
		return 0, &AddBlockReq{}, err
	}
	s.BlockToAdd = *tBlockToAdd
	ovb = (*vb)[i:]
	j,tPreviousBlockID,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &AddBlockReq{}, err
	}
	s.PreviousBlockID = *tPreviousBlockID
	return i, &s, nil
}

// ----------------------------------------
//  Struct: AddBlockResp
// ----------------------------------------

// AddBlockResp type
type AddBlockResp struct {
}

// NewAddBlockResp factory
func NewAddBlockResp() *AddBlockResp {
	o := AddBlockResp{}
	return &o
}

// Serialize AddBlockResp
func (n AddBlockResp) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeAddBlockResp function
func DeserializeAddBlockResp(vb *VariableBlob) (uint64,*AddBlockResp,error) {
	var i uint64 = 0
	s := AddBlockResp{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: BlockRecord
// ----------------------------------------

// BlockRecord type
type BlockRecord struct {
    BlockID Multihash `json:"block_id"`
    BlockHeight BlockHeightType `json:"block_height"`
    PreviousBlockIds VectorMultihash `json:"previous_block_ids"`
    Block OpaqueBlock `json:"block"`
    BlockReceipt OpaqueBlockReceipt `json:"block_receipt"`
}

// NewBlockRecord factory
func NewBlockRecord() *BlockRecord {
	o := BlockRecord{}
	o.BlockID = *NewMultihash()
	o.BlockHeight = *NewBlockHeightType()
	o.PreviousBlockIds = *NewVectorMultihash()
	o.Block = *NewOpaqueBlock()
	o.BlockReceipt = *NewOpaqueBlockReceipt()
	return &o
}

// Serialize BlockRecord
func (n BlockRecord) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.BlockID.Serialize(vb)
	vb = n.BlockHeight.Serialize(vb)
	vb = n.PreviousBlockIds.Serialize(vb)
	vb = n.Block.Serialize(vb)
	vb = n.BlockReceipt.Serialize(vb)
	return vb
}

// DeserializeBlockRecord function
func DeserializeBlockRecord(vb *VariableBlob) (uint64,*BlockRecord,error) {
	var i,j uint64 = 0,0
	s := BlockRecord{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tBlockID,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &BlockRecord{}, err
	}
	s.BlockID = *tBlockID
	ovb = (*vb)[i:]
	j,tBlockHeight,err := DeserializeBlockHeightType(&ovb); i+=j
	if err != nil {
		return 0, &BlockRecord{}, err
	}
	s.BlockHeight = *tBlockHeight
	ovb = (*vb)[i:]
	j,tPreviousBlockIds,err := DeserializeVectorMultihash(&ovb); i+=j
	if err != nil {
		return 0, &BlockRecord{}, err
	}
	s.PreviousBlockIds = *tPreviousBlockIds
	ovb = (*vb)[i:]
	j,tBlock,err := DeserializeOpaqueBlock(&ovb); i+=j
	if err != nil {
		return 0, &BlockRecord{}, err
	}
	s.Block = *tBlock
	ovb = (*vb)[i:]
	j,tBlockReceipt,err := DeserializeOpaqueBlockReceipt(&ovb); i+=j
	if err != nil {
		return 0, &BlockRecord{}, err
	}
	s.BlockReceipt = *tBlockReceipt
	return i, &s, nil
}

// ----------------------------------------
//  Struct: AddTransactionReq
// ----------------------------------------

// AddTransactionReq type
type AddTransactionReq struct {
    TransactionID Multihash `json:"transaction_id"`
    Transaction OpaqueTransaction `json:"transaction"`
}

// NewAddTransactionReq factory
func NewAddTransactionReq() *AddTransactionReq {
	o := AddTransactionReq{}
	o.TransactionID = *NewMultihash()
	o.Transaction = *NewOpaqueTransaction()
	return &o
}

// Serialize AddTransactionReq
func (n AddTransactionReq) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.TransactionID.Serialize(vb)
	vb = n.Transaction.Serialize(vb)
	return vb
}

// DeserializeAddTransactionReq function
func DeserializeAddTransactionReq(vb *VariableBlob) (uint64,*AddTransactionReq,error) {
	var i,j uint64 = 0,0
	s := AddTransactionReq{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransactionID,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &AddTransactionReq{}, err
	}
	s.TransactionID = *tTransactionID
	ovb = (*vb)[i:]
	j,tTransaction,err := DeserializeOpaqueTransaction(&ovb); i+=j
	if err != nil {
		return 0, &AddTransactionReq{}, err
	}
	s.Transaction = *tTransaction
	return i, &s, nil
}

// ----------------------------------------
//  Struct: AddTransactionResp
// ----------------------------------------

// AddTransactionResp type
type AddTransactionResp struct {
}

// NewAddTransactionResp factory
func NewAddTransactionResp() *AddTransactionResp {
	o := AddTransactionResp{}
	return &o
}

// Serialize AddTransactionResp
func (n AddTransactionResp) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeAddTransactionResp function
func DeserializeAddTransactionResp(vb *VariableBlob) (uint64,*AddTransactionResp,error) {
	var i uint64 = 0
	s := AddTransactionResp{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: TransactionRecord
// ----------------------------------------

// TransactionRecord type
type TransactionRecord struct {
    Transaction OpaqueTransaction `json:"transaction"`
}

// NewTransactionRecord factory
func NewTransactionRecord() *TransactionRecord {
	o := TransactionRecord{}
	o.Transaction = *NewOpaqueTransaction()
	return &o
}

// Serialize TransactionRecord
func (n TransactionRecord) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Transaction.Serialize(vb)
	return vb
}

// DeserializeTransactionRecord function
func DeserializeTransactionRecord(vb *VariableBlob) (uint64,*TransactionRecord,error) {
	var i,j uint64 = 0,0
	s := TransactionRecord{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransaction,err := DeserializeOpaqueTransaction(&ovb); i+=j
	if err != nil {
		return 0, &TransactionRecord{}, err
	}
	s.Transaction = *tTransaction
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetTransactionsByIDReq
// ----------------------------------------

// GetTransactionsByIDReq type
type GetTransactionsByIDReq struct {
    TransactionIds VectorMultihash `json:"transaction_ids"`
}

// NewGetTransactionsByIDReq factory
func NewGetTransactionsByIDReq() *GetTransactionsByIDReq {
	o := GetTransactionsByIDReq{}
	o.TransactionIds = *NewVectorMultihash()
	return &o
}

// Serialize GetTransactionsByIDReq
func (n GetTransactionsByIDReq) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.TransactionIds.Serialize(vb)
	return vb
}

// DeserializeGetTransactionsByIDReq function
func DeserializeGetTransactionsByIDReq(vb *VariableBlob) (uint64,*GetTransactionsByIDReq,error) {
	var i,j uint64 = 0,0
	s := GetTransactionsByIDReq{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransactionIds,err := DeserializeVectorMultihash(&ovb); i+=j
	if err != nil {
		return 0, &GetTransactionsByIDReq{}, err
	}
	s.TransactionIds = *tTransactionIds
	return i, &s, nil
}

// ----------------------------------------
//  Struct: TransactionItem
// ----------------------------------------

// TransactionItem type
type TransactionItem struct {
    Transaction OpaqueTransaction `json:"transaction"`
}

// NewTransactionItem factory
func NewTransactionItem() *TransactionItem {
	o := TransactionItem{}
	o.Transaction = *NewOpaqueTransaction()
	return &o
}

// Serialize TransactionItem
func (n TransactionItem) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Transaction.Serialize(vb)
	return vb
}

// DeserializeTransactionItem function
func DeserializeTransactionItem(vb *VariableBlob) (uint64,*TransactionItem,error) {
	var i,j uint64 = 0,0
	s := TransactionItem{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransaction,err := DeserializeOpaqueTransaction(&ovb); i+=j
	if err != nil {
		return 0, &TransactionItem{}, err
	}
	s.Transaction = *tTransaction
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetTransactionsByIDResp
// ----------------------------------------

// GetTransactionsByIDResp type
type GetTransactionsByIDResp struct {
    TransactionItems VectorTransactionItem `json:"transaction_items"`
}

// NewGetTransactionsByIDResp factory
func NewGetTransactionsByIDResp() *GetTransactionsByIDResp {
	o := GetTransactionsByIDResp{}
	o.TransactionItems = *NewVectorTransactionItem()
	return &o
}

// Serialize GetTransactionsByIDResp
func (n GetTransactionsByIDResp) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.TransactionItems.Serialize(vb)
	return vb
}

// DeserializeGetTransactionsByIDResp function
func DeserializeGetTransactionsByIDResp(vb *VariableBlob) (uint64,*GetTransactionsByIDResp,error) {
	var i,j uint64 = 0,0
	s := GetTransactionsByIDResp{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransactionItems,err := DeserializeVectorTransactionItem(&ovb); i+=j
	if err != nil {
		return 0, &GetTransactionsByIDResp{}, err
	}
	s.TransactionItems = *tTransactionItems
	return i, &s, nil
}

// ----------------------------------------
//  Struct: BlockStoreError
// ----------------------------------------

// BlockStoreError type
type BlockStoreError struct {
    ErrorText String `json:"error_text"`
}

// NewBlockStoreError factory
func NewBlockStoreError() *BlockStoreError {
	o := BlockStoreError{}
	o.ErrorText = *NewString()
	return &o
}

// Serialize BlockStoreError
func (n BlockStoreError) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ErrorText.Serialize(vb)
	return vb
}

// DeserializeBlockStoreError function
func DeserializeBlockStoreError(vb *VariableBlob) (uint64,*BlockStoreError,error) {
	var i,j uint64 = 0,0
	s := BlockStoreError{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tErrorText,err := DeserializeString(&ovb); i+=j
	if err != nil {
		return 0, &BlockStoreError{}, err
	}
	s.ErrorText = *tErrorText
	return i, &s, nil
}

// ----------------------------------------
//  Variant: BlockStoreReq
// ----------------------------------------

// BlockStoreReq type
type BlockStoreReq struct {
	Value interface{}
}

// NewBlockStoreReq factory
func NewBlockStoreReq() *BlockStoreReq {
	v := BlockStoreReq{}
	v.Value = NewReservedReq()
	return &v
}

// Serialize BlockStoreReq
func (n BlockStoreReq) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *ReservedReq:
			i = 0
		case *GetBlocksByIDReq:
			i = 1
		case *GetBlocksByHeightReq:
			i = 2
		case *AddBlockReq:
			i = 3
		case *AddTransactionReq:
			i = 4
		case *GetTransactionsByIDReq:
			i = 5
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName BlockStoreReq
func (n BlockStoreReq) TypeToName() (string) {
	switch n.Value.(type) {
		case *ReservedReq:
			return "koinos::types::block_store::reserved_req"
		case *GetBlocksByIDReq:
			return "koinos::types::block_store::get_blocks_by_id_req"
		case *GetBlocksByHeightReq:
			return "koinos::types::block_store::get_blocks_by_height_req"
		case *AddBlockReq:
			return "koinos::types::block_store::add_block_req"
		case *AddTransactionReq:
			return "koinos::types::block_store::add_transaction_req"
		case *GetTransactionsByIDReq:
			return "koinos::types::block_store::get_transactions_by_id_req"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON BlockStoreReq
func (n BlockStoreReq) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeBlockStoreReq function
func DeserializeBlockStoreReq(vb *VariableBlob) (uint64,*BlockStoreReq,error) {
	var v BlockStoreReq
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewReservedReq()
		case 1:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetBlocksByIDReq(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 2:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetBlocksByHeightReq(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 3:
			ovb := (*vb)[i:]
			k,x,err := DeserializeAddBlockReq(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 4:
			ovb := (*vb)[i:]
			k,x,err := DeserializeAddTransactionReq(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 5:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetTransactionsByIDReq(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		default:
			return 0, &v, errors.New("unknown variant tag")
	}
	return uint64(i)+j,&v,nil
}

// UnmarshalJSON *BlockStoreReq
func (n *BlockStoreReq) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::types::block_store::reserved_req":
			v := NewReservedReq()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::block_store::get_blocks_by_id_req":
			v := NewGetBlocksByIDReq()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::block_store::get_blocks_by_height_req":
			v := NewGetBlocksByHeightReq()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::block_store::add_block_req":
			v := NewAddBlockReq()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::block_store::add_transaction_req":
			v := NewAddTransactionReq()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::block_store::get_transactions_by_id_req":
			v := NewGetTransactionsByIDReq()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

	return nil
}


// ----------------------------------------
//  Variant: BlockStoreResp
// ----------------------------------------

// BlockStoreResp type
type BlockStoreResp struct {
	Value interface{}
}

// NewBlockStoreResp factory
func NewBlockStoreResp() *BlockStoreResp {
	v := BlockStoreResp{}
	v.Value = NewReservedResp()
	return &v
}

// Serialize BlockStoreResp
func (n BlockStoreResp) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *ReservedResp:
			i = 0
		case *BlockStoreError:
			i = 1
		case *GetBlocksByIDResp:
			i = 2
		case *GetBlocksByHeightResp:
			i = 3
		case *AddBlockResp:
			i = 4
		case *AddTransactionResp:
			i = 5
		case *GetTransactionsByIDResp:
			i = 6
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName BlockStoreResp
func (n BlockStoreResp) TypeToName() (string) {
	switch n.Value.(type) {
		case *ReservedResp:
			return "koinos::types::block_store::reserved_resp"
		case *BlockStoreError:
			return "koinos::types::block_store::block_store_error"
		case *GetBlocksByIDResp:
			return "koinos::types::block_store::get_blocks_by_id_resp"
		case *GetBlocksByHeightResp:
			return "koinos::types::block_store::get_blocks_by_height_resp"
		case *AddBlockResp:
			return "koinos::types::block_store::add_block_resp"
		case *AddTransactionResp:
			return "koinos::types::block_store::add_transaction_resp"
		case *GetTransactionsByIDResp:
			return "koinos::types::block_store::get_transactions_by_id_resp"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON BlockStoreResp
func (n BlockStoreResp) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeBlockStoreResp function
func DeserializeBlockStoreResp(vb *VariableBlob) (uint64,*BlockStoreResp,error) {
	var v BlockStoreResp
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewReservedResp()
		case 1:
			ovb := (*vb)[i:]
			k,x,err := DeserializeBlockStoreError(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 2:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetBlocksByIDResp(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 3:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetBlocksByHeightResp(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 4:
			v.Value = NewAddBlockResp()
		case 5:
			v.Value = NewAddTransactionResp()
		case 6:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetTransactionsByIDResp(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		default:
			return 0, &v, errors.New("unknown variant tag")
	}
	return uint64(i)+j,&v,nil
}

// UnmarshalJSON *BlockStoreResp
func (n *BlockStoreResp) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::types::block_store::reserved_resp":
			v := NewReservedResp()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::block_store::block_store_error":
			v := NewBlockStoreError()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::block_store::get_blocks_by_id_resp":
			v := NewGetBlocksByIDResp()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::block_store::get_blocks_by_height_resp":
			v := NewGetBlocksByHeightResp()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::block_store::add_block_resp":
			v := NewAddBlockResp()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::block_store::add_transaction_resp":
			v := NewAddTransactionResp()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::block_store::get_transactions_by_id_resp":
			v := NewGetTransactionsByIDResp()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

	return nil
}


// ----------------------------------------
//  Struct: HeadInfo
// ----------------------------------------

// HeadInfo type
type HeadInfo struct {
    ID Multihash `json:"id"`
    Height BlockHeightType `json:"height"`
}

// NewHeadInfo factory
func NewHeadInfo() *HeadInfo {
	o := HeadInfo{}
	o.ID = *NewMultihash()
	o.Height = *NewBlockHeightType()
	return &o
}

// Serialize HeadInfo
func (n HeadInfo) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ID.Serialize(vb)
	vb = n.Height.Serialize(vb)
	return vb
}

// DeserializeHeadInfo function
func DeserializeHeadInfo(vb *VariableBlob) (uint64,*HeadInfo,error) {
	var i,j uint64 = 0,0
	s := HeadInfo{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tID,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &HeadInfo{}, err
	}
	s.ID = *tID
	ovb = (*vb)[i:]
	j,tHeight,err := DeserializeBlockHeightType(&ovb); i+=j
	if err != nil {
		return 0, &HeadInfo{}, err
	}
	s.Height = *tHeight
	return i, &s, nil
}

// ----------------------------------------
//  Struct: BlockPart
// ----------------------------------------

// BlockPart type
type BlockPart struct {
    ActiveData VariableBlob `json:"active_data"`
    PassiveData VariableBlob `json:"passive_data"`
    SigData VariableBlob `json:"sig_data"`
}

// NewBlockPart factory
func NewBlockPart() *BlockPart {
	o := BlockPart{}
	o.ActiveData = *NewVariableBlob()
	o.PassiveData = *NewVariableBlob()
	o.SigData = *NewVariableBlob()
	return &o
}

// Serialize BlockPart
func (n BlockPart) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ActiveData.Serialize(vb)
	vb = n.PassiveData.Serialize(vb)
	vb = n.SigData.Serialize(vb)
	return vb
}

// DeserializeBlockPart function
func DeserializeBlockPart(vb *VariableBlob) (uint64,*BlockPart,error) {
	var i,j uint64 = 0,0
	s := BlockPart{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tActiveData,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &BlockPart{}, err
	}
	s.ActiveData = *tActiveData
	ovb = (*vb)[i:]
	j,tPassiveData,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &BlockPart{}, err
	}
	s.PassiveData = *tPassiveData
	ovb = (*vb)[i:]
	j,tSigData,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &BlockPart{}, err
	}
	s.SigData = *tSigData
	return i, &s, nil
}

// ----------------------------------------
//  Struct: ReservedRPCParams
// ----------------------------------------

// ReservedRPCParams type
type ReservedRPCParams struct {
}

// NewReservedRPCParams factory
func NewReservedRPCParams() *ReservedRPCParams {
	o := ReservedRPCParams{}
	return &o
}

// Serialize ReservedRPCParams
func (n ReservedRPCParams) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeReservedRPCParams function
func DeserializeReservedRPCParams(vb *VariableBlob) (uint64,*ReservedRPCParams,error) {
	var i uint64 = 0
	s := ReservedRPCParams{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: BlockTopology
// ----------------------------------------

// BlockTopology type
type BlockTopology struct {
    ID Multihash `json:"id"`
    Height BlockHeightType `json:"height"`
    Previous Multihash `json:"previous"`
}

// NewBlockTopology factory
func NewBlockTopology() *BlockTopology {
	o := BlockTopology{}
	o.ID = *NewMultihash()
	o.Height = *NewBlockHeightType()
	o.Previous = *NewMultihash()
	return &o
}

// Serialize BlockTopology
func (n BlockTopology) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ID.Serialize(vb)
	vb = n.Height.Serialize(vb)
	vb = n.Previous.Serialize(vb)
	return vb
}

// DeserializeBlockTopology function
func DeserializeBlockTopology(vb *VariableBlob) (uint64,*BlockTopology,error) {
	var i,j uint64 = 0,0
	s := BlockTopology{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tID,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &BlockTopology{}, err
	}
	s.ID = *tID
	ovb = (*vb)[i:]
	j,tHeight,err := DeserializeBlockHeightType(&ovb); i+=j
	if err != nil {
		return 0, &BlockTopology{}, err
	}
	s.Height = *tHeight
	ovb = (*vb)[i:]
	j,tPrevious,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &BlockTopology{}, err
	}
	s.Previous = *tPrevious
	return i, &s, nil
}

// ----------------------------------------
//  Struct: SubmitBlockParams
// ----------------------------------------

// SubmitBlockParams type
type SubmitBlockParams struct {
    Topology BlockTopology `json:"topology"`
    Block Block `json:"block"`
    VerifyPassiveData Boolean `json:"verify_passive_data"`
    VerifyBlockSignature Boolean `json:"verify_block_signature"`
    VerifyTransactionSignatures Boolean `json:"verify_transaction_signatures"`
}

// NewSubmitBlockParams factory
func NewSubmitBlockParams() *SubmitBlockParams {
	o := SubmitBlockParams{}
	o.Topology = *NewBlockTopology()
	o.Block = *NewBlock()
	o.VerifyPassiveData = *NewBoolean()
	o.VerifyBlockSignature = *NewBoolean()
	o.VerifyTransactionSignatures = *NewBoolean()
	return &o
}

// Serialize SubmitBlockParams
func (n SubmitBlockParams) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Topology.Serialize(vb)
	vb = n.Block.Serialize(vb)
	vb = n.VerifyPassiveData.Serialize(vb)
	vb = n.VerifyBlockSignature.Serialize(vb)
	vb = n.VerifyTransactionSignatures.Serialize(vb)
	return vb
}

// DeserializeSubmitBlockParams function
func DeserializeSubmitBlockParams(vb *VariableBlob) (uint64,*SubmitBlockParams,error) {
	var i,j uint64 = 0,0
	s := SubmitBlockParams{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTopology,err := DeserializeBlockTopology(&ovb); i+=j
	if err != nil {
		return 0, &SubmitBlockParams{}, err
	}
	s.Topology = *tTopology
	ovb = (*vb)[i:]
	j,tBlock,err := DeserializeBlock(&ovb); i+=j
	if err != nil {
		return 0, &SubmitBlockParams{}, err
	}
	s.Block = *tBlock
	ovb = (*vb)[i:]
	j,tVerifyPassiveData,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &SubmitBlockParams{}, err
	}
	s.VerifyPassiveData = *tVerifyPassiveData
	ovb = (*vb)[i:]
	j,tVerifyBlockSignature,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &SubmitBlockParams{}, err
	}
	s.VerifyBlockSignature = *tVerifyBlockSignature
	ovb = (*vb)[i:]
	j,tVerifyTransactionSignatures,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &SubmitBlockParams{}, err
	}
	s.VerifyTransactionSignatures = *tVerifyTransactionSignatures
	return i, &s, nil
}

// ----------------------------------------
//  Struct: SubmitTransactionParams
// ----------------------------------------

// SubmitTransactionParams type
type SubmitTransactionParams struct {
    ActiveBytes VariableBlob `json:"active_bytes"`
    PassiveBytes VariableBlob `json:"passive_bytes"`
}

// NewSubmitTransactionParams factory
func NewSubmitTransactionParams() *SubmitTransactionParams {
	o := SubmitTransactionParams{}
	o.ActiveBytes = *NewVariableBlob()
	o.PassiveBytes = *NewVariableBlob()
	return &o
}

// Serialize SubmitTransactionParams
func (n SubmitTransactionParams) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ActiveBytes.Serialize(vb)
	vb = n.PassiveBytes.Serialize(vb)
	return vb
}

// DeserializeSubmitTransactionParams function
func DeserializeSubmitTransactionParams(vb *VariableBlob) (uint64,*SubmitTransactionParams,error) {
	var i,j uint64 = 0,0
	s := SubmitTransactionParams{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tActiveBytes,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &SubmitTransactionParams{}, err
	}
	s.ActiveBytes = *tActiveBytes
	ovb = (*vb)[i:]
	j,tPassiveBytes,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &SubmitTransactionParams{}, err
	}
	s.PassiveBytes = *tPassiveBytes
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetHeadInfoParams
// ----------------------------------------

// GetHeadInfoParams type
type GetHeadInfoParams struct {
}

// NewGetHeadInfoParams factory
func NewGetHeadInfoParams() *GetHeadInfoParams {
	o := GetHeadInfoParams{}
	return &o
}

// Serialize GetHeadInfoParams
func (n GetHeadInfoParams) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeGetHeadInfoParams function
func DeserializeGetHeadInfoParams(vb *VariableBlob) (uint64,*GetHeadInfoParams,error) {
	var i uint64 = 0
	s := GetHeadInfoParams{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetChainIDParams
// ----------------------------------------

// GetChainIDParams type
type GetChainIDParams struct {
}

// NewGetChainIDParams factory
func NewGetChainIDParams() *GetChainIDParams {
	o := GetChainIDParams{}
	return &o
}

// Serialize GetChainIDParams
func (n GetChainIDParams) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeGetChainIDParams function
func DeserializeGetChainIDParams(vb *VariableBlob) (uint64,*GetChainIDParams,error) {
	var i uint64 = 0
	s := GetChainIDParams{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Variant: ChainRPCParams
// ----------------------------------------

// ChainRPCParams type
type ChainRPCParams struct {
	Value interface{}
}

// NewChainRPCParams factory
func NewChainRPCParams() *ChainRPCParams {
	v := ChainRPCParams{}
	v.Value = NewReservedRPCParams()
	return &v
}

// Serialize ChainRPCParams
func (n ChainRPCParams) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *ReservedRPCParams:
			i = 0
		case *SubmitBlockParams:
			i = 1
		case *SubmitTransactionParams:
			i = 2
		case *GetHeadInfoParams:
			i = 3
		case *GetChainIDParams:
			i = 4
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName ChainRPCParams
func (n ChainRPCParams) TypeToName() (string) {
	switch n.Value.(type) {
		case *ReservedRPCParams:
			return "koinos::rpc::chain::reserved_rpc_params"
		case *SubmitBlockParams:
			return "koinos::rpc::chain::submit_block_params"
		case *SubmitTransactionParams:
			return "koinos::rpc::chain::submit_transaction_params"
		case *GetHeadInfoParams:
			return "koinos::rpc::chain::get_head_info_params"
		case *GetChainIDParams:
			return "koinos::rpc::chain::get_chain_id_params"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON ChainRPCParams
func (n ChainRPCParams) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeChainRPCParams function
func DeserializeChainRPCParams(vb *VariableBlob) (uint64,*ChainRPCParams,error) {
	var v ChainRPCParams
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewReservedRPCParams()
		case 1:
			ovb := (*vb)[i:]
			k,x,err := DeserializeSubmitBlockParams(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 2:
			ovb := (*vb)[i:]
			k,x,err := DeserializeSubmitTransactionParams(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 3:
			v.Value = NewGetHeadInfoParams()
		case 4:
			v.Value = NewGetChainIDParams()
		default:
			return 0, &v, errors.New("unknown variant tag")
	}
	return uint64(i)+j,&v,nil
}

// UnmarshalJSON *ChainRPCParams
func (n *ChainRPCParams) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::rpc::chain::reserved_rpc_params":
			v := NewReservedRPCParams()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::submit_block_params":
			v := NewSubmitBlockParams()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::submit_transaction_params":
			v := NewSubmitTransactionParams()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::get_head_info_params":
			v := NewGetHeadInfoParams()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::get_chain_id_params":
			v := NewGetChainIDParams()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

	return nil
}


// ----------------------------------------
//  Struct: ReservedRPCResult
// ----------------------------------------

// ReservedRPCResult type
type ReservedRPCResult struct {
}

// NewReservedRPCResult factory
func NewReservedRPCResult() *ReservedRPCResult {
	o := ReservedRPCResult{}
	return &o
}

// Serialize ReservedRPCResult
func (n ReservedRPCResult) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeReservedRPCResult function
func DeserializeReservedRPCResult(vb *VariableBlob) (uint64,*ReservedRPCResult,error) {
	var i uint64 = 0
	s := ReservedRPCResult{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: RPCError
// ----------------------------------------

// RPCError type
type RPCError struct {
    ErrorText String `json:"error_text"`
}

// NewRPCError factory
func NewRPCError() *RPCError {
	o := RPCError{}
	o.ErrorText = *NewString()
	return &o
}

// Serialize RPCError
func (n RPCError) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ErrorText.Serialize(vb)
	return vb
}

// DeserializeRPCError function
func DeserializeRPCError(vb *VariableBlob) (uint64,*RPCError,error) {
	var i,j uint64 = 0,0
	s := RPCError{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tErrorText,err := DeserializeString(&ovb); i+=j
	if err != nil {
		return 0, &RPCError{}, err
	}
	s.ErrorText = *tErrorText
	return i, &s, nil
}

// ----------------------------------------
//  Struct: SubmitBlockResult
// ----------------------------------------

// SubmitBlockResult type
type SubmitBlockResult struct {
}

// NewSubmitBlockResult factory
func NewSubmitBlockResult() *SubmitBlockResult {
	o := SubmitBlockResult{}
	return &o
}

// Serialize SubmitBlockResult
func (n SubmitBlockResult) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeSubmitBlockResult function
func DeserializeSubmitBlockResult(vb *VariableBlob) (uint64,*SubmitBlockResult,error) {
	var i uint64 = 0
	s := SubmitBlockResult{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: SubmitTransactionResult
// ----------------------------------------

// SubmitTransactionResult type
type SubmitTransactionResult struct {
}

// NewSubmitTransactionResult factory
func NewSubmitTransactionResult() *SubmitTransactionResult {
	o := SubmitTransactionResult{}
	return &o
}

// Serialize SubmitTransactionResult
func (n SubmitTransactionResult) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeSubmitTransactionResult function
func DeserializeSubmitTransactionResult(vb *VariableBlob) (uint64,*SubmitTransactionResult,error) {
	var i uint64 = 0
	s := SubmitTransactionResult{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetHeadInfoResult
// ----------------------------------------

// GetHeadInfoResult type
type GetHeadInfoResult struct {
    ID Multihash `json:"id"`
    Height BlockHeightType `json:"height"`
}

// NewGetHeadInfoResult factory
func NewGetHeadInfoResult() *GetHeadInfoResult {
	o := GetHeadInfoResult{}
	o.ID = *NewMultihash()
	o.Height = *NewBlockHeightType()
	return &o
}

// Serialize GetHeadInfoResult
func (n GetHeadInfoResult) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ID.Serialize(vb)
	vb = n.Height.Serialize(vb)
	return vb
}

// DeserializeGetHeadInfoResult function
func DeserializeGetHeadInfoResult(vb *VariableBlob) (uint64,*GetHeadInfoResult,error) {
	var i,j uint64 = 0,0
	s := GetHeadInfoResult{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tID,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &GetHeadInfoResult{}, err
	}
	s.ID = *tID
	ovb = (*vb)[i:]
	j,tHeight,err := DeserializeBlockHeightType(&ovb); i+=j
	if err != nil {
		return 0, &GetHeadInfoResult{}, err
	}
	s.Height = *tHeight
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetChainIDResult
// ----------------------------------------

// GetChainIDResult type
type GetChainIDResult struct {
    ChainID Multihash `json:"chain_id"`
}

// NewGetChainIDResult factory
func NewGetChainIDResult() *GetChainIDResult {
	o := GetChainIDResult{}
	o.ChainID = *NewMultihash()
	return &o
}

// Serialize GetChainIDResult
func (n GetChainIDResult) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ChainID.Serialize(vb)
	return vb
}

// DeserializeGetChainIDResult function
func DeserializeGetChainIDResult(vb *VariableBlob) (uint64,*GetChainIDResult,error) {
	var i,j uint64 = 0,0
	s := GetChainIDResult{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tChainID,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &GetChainIDResult{}, err
	}
	s.ChainID = *tChainID
	return i, &s, nil
}

// ----------------------------------------
//  Variant: ChainRPCResult
// ----------------------------------------

// ChainRPCResult type
type ChainRPCResult struct {
	Value interface{}
}

// NewChainRPCResult factory
func NewChainRPCResult() *ChainRPCResult {
	v := ChainRPCResult{}
	v.Value = NewReservedRPCResult()
	return &v
}

// Serialize ChainRPCResult
func (n ChainRPCResult) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *ReservedRPCResult:
			i = 0
		case *RPCError:
			i = 1
		case *SubmitBlockResult:
			i = 2
		case *SubmitTransactionResult:
			i = 3
		case *GetHeadInfoResult:
			i = 4
		case *GetChainIDResult:
			i = 5
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName ChainRPCResult
func (n ChainRPCResult) TypeToName() (string) {
	switch n.Value.(type) {
		case *ReservedRPCResult:
			return "koinos::rpc::chain::reserved_rpc_result"
		case *RPCError:
			return "koinos::rpc::chain::rpc_error"
		case *SubmitBlockResult:
			return "koinos::rpc::chain::submit_block_result"
		case *SubmitTransactionResult:
			return "koinos::rpc::chain::submit_transaction_result"
		case *GetHeadInfoResult:
			return "koinos::rpc::chain::get_head_info_result"
		case *GetChainIDResult:
			return "koinos::rpc::chain::get_chain_id_result"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON ChainRPCResult
func (n ChainRPCResult) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeChainRPCResult function
func DeserializeChainRPCResult(vb *VariableBlob) (uint64,*ChainRPCResult,error) {
	var v ChainRPCResult
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewReservedRPCResult()
		case 1:
			ovb := (*vb)[i:]
			k,x,err := DeserializeRPCError(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 2:
			v.Value = NewSubmitBlockResult()
		case 3:
			v.Value = NewSubmitTransactionResult()
		case 4:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetHeadInfoResult(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 5:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetChainIDResult(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		default:
			return 0, &v, errors.New("unknown variant tag")
	}
	return uint64(i)+j,&v,nil
}

// UnmarshalJSON *ChainRPCResult
func (n *ChainRPCResult) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::rpc::chain::reserved_rpc_result":
			v := NewReservedRPCResult()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::rpc_error":
			v := NewRPCError()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::submit_block_result":
			v := NewSubmitBlockResult()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::submit_transaction_result":
			v := NewSubmitTransactionResult()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::get_head_info_result":
			v := NewGetHeadInfoResult()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::get_chain_id_result":
			v := NewGetChainIDResult()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

	return nil
}


// ----------------------------------------
//  Typedef: ReservedQueryParams
// ----------------------------------------

// ReservedQueryParams type
type ReservedQueryParams ReservedRPCParams

// NewReservedQueryParams factory
func NewReservedQueryParams() *ReservedQueryParams {
	o := ReservedQueryParams(*NewReservedRPCParams())
	return &o
}

// Serialize ReservedQueryParams
func (n ReservedQueryParams) Serialize(vb *VariableBlob) *VariableBlob {
	ox := ReservedRPCParams(n)
	return ox.Serialize(vb)
}

// DeserializeReservedQueryParams function
func DeserializeReservedQueryParams(vb *VariableBlob) (uint64,*ReservedQueryParams,error) {
	var ot ReservedQueryParams
	return 0,&ot,nil}

// MarshalJSON ReservedQueryParams
func (n ReservedQueryParams) MarshalJSON() ([]byte, error) {
	v := ReservedRPCParams(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ReservedQueryParams
func (n *ReservedQueryParams) UnmarshalJSON(data []byte) error {
	v := ReservedRPCParams(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ReservedQueryParams(v)
	return nil
}








// ----------------------------------------
//  Variant: QueryParamItem
// ----------------------------------------

// QueryParamItem type
type QueryParamItem struct {
	Value interface{}
}

// NewQueryParamItem factory
func NewQueryParamItem() *QueryParamItem {
	v := QueryParamItem{}
	v.Value = NewReservedQueryParams()
	return &v
}

// Serialize QueryParamItem
func (n QueryParamItem) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *ReservedQueryParams:
			i = 0
		case *GetHeadInfoParams:
			i = 1
		case *GetChainIDParams:
			i = 2
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName QueryParamItem
func (n QueryParamItem) TypeToName() (string) {
	switch n.Value.(type) {
		case *ReservedQueryParams:
			return "koinos::types::rpc::reserved_query_params"
		case *GetHeadInfoParams:
			return "koinos::types::rpc::get_head_info_params"
		case *GetChainIDParams:
			return "koinos::types::rpc::get_chain_id_params"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON QueryParamItem
func (n QueryParamItem) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeQueryParamItem function
func DeserializeQueryParamItem(vb *VariableBlob) (uint64,*QueryParamItem,error) {
	var v QueryParamItem
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewReservedQueryParams()
		case 1:
			v.Value = NewGetHeadInfoParams()
		case 2:
			v.Value = NewGetChainIDParams()
		default:
			return 0, &v, errors.New("unknown variant tag")
	}
	return uint64(i)+j,&v,nil
}

// UnmarshalJSON *QueryParamItem
func (n *QueryParamItem) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::types::rpc::reserved_query_params":
			v := NewReservedQueryParams()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::rpc::get_head_info_params":
			v := NewGetHeadInfoParams()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::rpc::get_chain_id_params":
			v := NewGetChainIDParams()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

	return nil
}


// ----------------------------------------
//  Typedef: QuerySubmission
// ----------------------------------------

// QuerySubmission type
type QuerySubmission OpaqueQueryParamItem

// NewQuerySubmission factory
func NewQuerySubmission() *QuerySubmission {
	o := QuerySubmission(*NewOpaqueQueryParamItem())
	return &o
}

// Serialize QuerySubmission
func (n QuerySubmission) Serialize(vb *VariableBlob) *VariableBlob {
	ox := OpaqueQueryParamItem(n)
	return ox.Serialize(vb)
}

// DeserializeQuerySubmission function
func DeserializeQuerySubmission(vb *VariableBlob) (uint64,*QuerySubmission,error) {
	var ot QuerySubmission
	i,n,err := DeserializeOpaqueQueryParamItem(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = QuerySubmission(*n)
	return i,&ot,nil}

// MarshalJSON QuerySubmission
func (n QuerySubmission) MarshalJSON() ([]byte, error) {
	v := OpaqueQueryParamItem(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *QuerySubmission
func (n *QuerySubmission) UnmarshalJSON(data []byte) error {
	v := OpaqueQueryParamItem(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = QuerySubmission(v)
	return nil
}


// ----------------------------------------
//  Typedef: ReservedQueryResult
// ----------------------------------------

// ReservedQueryResult type
type ReservedQueryResult ReservedRPCResult

// NewReservedQueryResult factory
func NewReservedQueryResult() *ReservedQueryResult {
	o := ReservedQueryResult(*NewReservedRPCResult())
	return &o
}

// Serialize ReservedQueryResult
func (n ReservedQueryResult) Serialize(vb *VariableBlob) *VariableBlob {
	ox := ReservedRPCResult(n)
	return ox.Serialize(vb)
}

// DeserializeReservedQueryResult function
func DeserializeReservedQueryResult(vb *VariableBlob) (uint64,*ReservedQueryResult,error) {
	var ot ReservedQueryResult
	return 0,&ot,nil}

// MarshalJSON ReservedQueryResult
func (n ReservedQueryResult) MarshalJSON() ([]byte, error) {
	v := ReservedRPCResult(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ReservedQueryResult
func (n *ReservedQueryResult) UnmarshalJSON(data []byte) error {
	v := ReservedRPCResult(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ReservedQueryResult(v)
	return nil
}


// ----------------------------------------
//  Typedef: QueryError
// ----------------------------------------

// QueryError type
type QueryError RPCError

// NewQueryError factory
func NewQueryError() *QueryError {
	o := QueryError(*NewRPCError())
	return &o
}

// Serialize QueryError
func (n QueryError) Serialize(vb *VariableBlob) *VariableBlob {
	ox := RPCError(n)
	return ox.Serialize(vb)
}

// DeserializeQueryError function
func DeserializeQueryError(vb *VariableBlob) (uint64,*QueryError,error) {
	var ot QueryError
	i,n,err := DeserializeRPCError(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = QueryError(*n)
	return i,&ot,nil}

// MarshalJSON QueryError
func (n QueryError) MarshalJSON() ([]byte, error) {
	v := RPCError(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *QueryError
func (n *QueryError) UnmarshalJSON(data []byte) error {
	v := RPCError(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = QueryError(v)
	return nil
}








// ----------------------------------------
//  Variant: QueryItemResult
// ----------------------------------------

// QueryItemResult type
type QueryItemResult struct {
	Value interface{}
}

// NewQueryItemResult factory
func NewQueryItemResult() *QueryItemResult {
	v := QueryItemResult{}
	v.Value = NewReservedQueryResult()
	return &v
}

// Serialize QueryItemResult
func (n QueryItemResult) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *ReservedQueryResult:
			i = 0
		case *QueryError:
			i = 1
		case *GetHeadInfoResult:
			i = 2
		case *GetChainIDResult:
			i = 3
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName QueryItemResult
func (n QueryItemResult) TypeToName() (string) {
	switch n.Value.(type) {
		case *ReservedQueryResult:
			return "koinos::types::rpc::reserved_query_result"
		case *QueryError:
			return "koinos::types::rpc::query_error"
		case *GetHeadInfoResult:
			return "koinos::types::rpc::get_head_info_result"
		case *GetChainIDResult:
			return "koinos::types::rpc::get_chain_id_result"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON QueryItemResult
func (n QueryItemResult) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeQueryItemResult function
func DeserializeQueryItemResult(vb *VariableBlob) (uint64,*QueryItemResult,error) {
	var v QueryItemResult
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewReservedQueryResult()
		case 1:
			ovb := (*vb)[i:]
			k,x,err := DeserializeQueryError(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 2:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetHeadInfoResult(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 3:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetChainIDResult(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		default:
			return 0, &v, errors.New("unknown variant tag")
	}
	return uint64(i)+j,&v,nil
}

// UnmarshalJSON *QueryItemResult
func (n *QueryItemResult) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::types::rpc::reserved_query_result":
			v := NewReservedQueryResult()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::rpc::query_error":
			v := NewQueryError()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::rpc::get_head_info_result":
			v := NewGetHeadInfoResult()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::rpc::get_chain_id_result":
			v := NewGetChainIDResult()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

	return nil
}


// ----------------------------------------
//  Typedef: QuerySubmissionResult
// ----------------------------------------

// QuerySubmissionResult type
type QuerySubmissionResult OpaqueQueryItemResult

// NewQuerySubmissionResult factory
func NewQuerySubmissionResult() *QuerySubmissionResult {
	o := QuerySubmissionResult(*NewOpaqueQueryItemResult())
	return &o
}

// Serialize QuerySubmissionResult
func (n QuerySubmissionResult) Serialize(vb *VariableBlob) *VariableBlob {
	ox := OpaqueQueryItemResult(n)
	return ox.Serialize(vb)
}

// DeserializeQuerySubmissionResult function
func DeserializeQuerySubmissionResult(vb *VariableBlob) (uint64,*QuerySubmissionResult,error) {
	var ot QuerySubmissionResult
	i,n,err := DeserializeOpaqueQueryItemResult(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = QuerySubmissionResult(*n)
	return i,&ot,nil}

// MarshalJSON QuerySubmissionResult
func (n QuerySubmissionResult) MarshalJSON() ([]byte, error) {
	v := OpaqueQueryItemResult(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *QuerySubmissionResult
func (n *QuerySubmissionResult) UnmarshalJSON(data []byte) error {
	v := OpaqueQueryItemResult(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = QuerySubmissionResult(v)
	return nil
}


// ----------------------------------------
//  Typedef: ReservedSubmission
// ----------------------------------------

// ReservedSubmission type
type ReservedSubmission ReservedRPCParams

// NewReservedSubmission factory
func NewReservedSubmission() *ReservedSubmission {
	o := ReservedSubmission(*NewReservedRPCParams())
	return &o
}

// Serialize ReservedSubmission
func (n ReservedSubmission) Serialize(vb *VariableBlob) *VariableBlob {
	ox := ReservedRPCParams(n)
	return ox.Serialize(vb)
}

// DeserializeReservedSubmission function
func DeserializeReservedSubmission(vb *VariableBlob) (uint64,*ReservedSubmission,error) {
	var ot ReservedSubmission
	return 0,&ot,nil}

// MarshalJSON ReservedSubmission
func (n ReservedSubmission) MarshalJSON() ([]byte, error) {
	v := ReservedRPCParams(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ReservedSubmission
func (n *ReservedSubmission) UnmarshalJSON(data []byte) error {
	v := ReservedRPCParams(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ReservedSubmission(v)
	return nil
}


// ----------------------------------------
//  Typedef: BlockSubmission
// ----------------------------------------

// BlockSubmission type
type BlockSubmission SubmitBlockParams

// NewBlockSubmission factory
func NewBlockSubmission() *BlockSubmission {
	o := BlockSubmission(*NewSubmitBlockParams())
	return &o
}

// Serialize BlockSubmission
func (n BlockSubmission) Serialize(vb *VariableBlob) *VariableBlob {
	ox := SubmitBlockParams(n)
	return ox.Serialize(vb)
}

// DeserializeBlockSubmission function
func DeserializeBlockSubmission(vb *VariableBlob) (uint64,*BlockSubmission,error) {
	var ot BlockSubmission
	i,n,err := DeserializeSubmitBlockParams(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = BlockSubmission(*n)
	return i,&ot,nil}

// MarshalJSON BlockSubmission
func (n BlockSubmission) MarshalJSON() ([]byte, error) {
	v := SubmitBlockParams(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *BlockSubmission
func (n *BlockSubmission) UnmarshalJSON(data []byte) error {
	v := SubmitBlockParams(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = BlockSubmission(v)
	return nil
}


// ----------------------------------------
//  Typedef: TransactionSubmission
// ----------------------------------------

// TransactionSubmission type
type TransactionSubmission SubmitTransactionParams

// NewTransactionSubmission factory
func NewTransactionSubmission() *TransactionSubmission {
	o := TransactionSubmission(*NewSubmitTransactionParams())
	return &o
}

// Serialize TransactionSubmission
func (n TransactionSubmission) Serialize(vb *VariableBlob) *VariableBlob {
	ox := SubmitTransactionParams(n)
	return ox.Serialize(vb)
}

// DeserializeTransactionSubmission function
func DeserializeTransactionSubmission(vb *VariableBlob) (uint64,*TransactionSubmission,error) {
	var ot TransactionSubmission
	i,n,err := DeserializeSubmitTransactionParams(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = TransactionSubmission(*n)
	return i,&ot,nil}

// MarshalJSON TransactionSubmission
func (n TransactionSubmission) MarshalJSON() ([]byte, error) {
	v := SubmitTransactionParams(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *TransactionSubmission
func (n *TransactionSubmission) UnmarshalJSON(data []byte) error {
	v := SubmitTransactionParams(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = TransactionSubmission(v)
	return nil
}


// ----------------------------------------
//  Variant: SubmissionItem
// ----------------------------------------

// SubmissionItem type
type SubmissionItem struct {
	Value interface{}
}

// NewSubmissionItem factory
func NewSubmissionItem() *SubmissionItem {
	v := SubmissionItem{}
	v.Value = NewReservedSubmission()
	return &v
}

// Serialize SubmissionItem
func (n SubmissionItem) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *ReservedSubmission:
			i = 0
		case *BlockSubmission:
			i = 1
		case *TransactionSubmission:
			i = 2
		case *QuerySubmission:
			i = 3
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName SubmissionItem
func (n SubmissionItem) TypeToName() (string) {
	switch n.Value.(type) {
		case *ReservedSubmission:
			return "koinos::types::rpc::reserved_submission"
		case *BlockSubmission:
			return "koinos::types::rpc::block_submission"
		case *TransactionSubmission:
			return "koinos::types::rpc::transaction_submission"
		case *QuerySubmission:
			return "koinos::types::rpc::query_submission"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON SubmissionItem
func (n SubmissionItem) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeSubmissionItem function
func DeserializeSubmissionItem(vb *VariableBlob) (uint64,*SubmissionItem,error) {
	var v SubmissionItem
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewReservedSubmission()
		case 1:
			ovb := (*vb)[i:]
			k,x,err := DeserializeBlockSubmission(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 2:
			ovb := (*vb)[i:]
			k,x,err := DeserializeTransactionSubmission(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 3:
			ovb := (*vb)[i:]
			k,x,err := DeserializeQuerySubmission(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		default:
			return 0, &v, errors.New("unknown variant tag")
	}
	return uint64(i)+j,&v,nil
}

// UnmarshalJSON *SubmissionItem
func (n *SubmissionItem) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::types::rpc::reserved_submission":
			v := NewReservedSubmission()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::rpc::block_submission":
			v := NewBlockSubmission()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::rpc::transaction_submission":
			v := NewTransactionSubmission()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::rpc::query_submission":
			v := NewQuerySubmission()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

	return nil
}


// ----------------------------------------
//  Typedef: ReservedSubmissionResult
// ----------------------------------------

// ReservedSubmissionResult type
type ReservedSubmissionResult ReservedRPCResult

// NewReservedSubmissionResult factory
func NewReservedSubmissionResult() *ReservedSubmissionResult {
	o := ReservedSubmissionResult(*NewReservedRPCResult())
	return &o
}

// Serialize ReservedSubmissionResult
func (n ReservedSubmissionResult) Serialize(vb *VariableBlob) *VariableBlob {
	ox := ReservedRPCResult(n)
	return ox.Serialize(vb)
}

// DeserializeReservedSubmissionResult function
func DeserializeReservedSubmissionResult(vb *VariableBlob) (uint64,*ReservedSubmissionResult,error) {
	var ot ReservedSubmissionResult
	return 0,&ot,nil}

// MarshalJSON ReservedSubmissionResult
func (n ReservedSubmissionResult) MarshalJSON() ([]byte, error) {
	v := ReservedRPCResult(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ReservedSubmissionResult
func (n *ReservedSubmissionResult) UnmarshalJSON(data []byte) error {
	v := ReservedRPCResult(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ReservedSubmissionResult(v)
	return nil
}


// ----------------------------------------
//  Typedef: BlockSubmissionResult
// ----------------------------------------

// BlockSubmissionResult type
type BlockSubmissionResult SubmitBlockResult

// NewBlockSubmissionResult factory
func NewBlockSubmissionResult() *BlockSubmissionResult {
	o := BlockSubmissionResult(*NewSubmitBlockResult())
	return &o
}

// Serialize BlockSubmissionResult
func (n BlockSubmissionResult) Serialize(vb *VariableBlob) *VariableBlob {
	ox := SubmitBlockResult(n)
	return ox.Serialize(vb)
}

// DeserializeBlockSubmissionResult function
func DeserializeBlockSubmissionResult(vb *VariableBlob) (uint64,*BlockSubmissionResult,error) {
	var ot BlockSubmissionResult
	return 0,&ot,nil}

// MarshalJSON BlockSubmissionResult
func (n BlockSubmissionResult) MarshalJSON() ([]byte, error) {
	v := SubmitBlockResult(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *BlockSubmissionResult
func (n *BlockSubmissionResult) UnmarshalJSON(data []byte) error {
	v := SubmitBlockResult(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = BlockSubmissionResult(v)
	return nil
}


// ----------------------------------------
//  Typedef: TransactionSubmissionResult
// ----------------------------------------

// TransactionSubmissionResult type
type TransactionSubmissionResult SubmitTransactionResult

// NewTransactionSubmissionResult factory
func NewTransactionSubmissionResult() *TransactionSubmissionResult {
	o := TransactionSubmissionResult(*NewSubmitTransactionResult())
	return &o
}

// Serialize TransactionSubmissionResult
func (n TransactionSubmissionResult) Serialize(vb *VariableBlob) *VariableBlob {
	ox := SubmitTransactionResult(n)
	return ox.Serialize(vb)
}

// DeserializeTransactionSubmissionResult function
func DeserializeTransactionSubmissionResult(vb *VariableBlob) (uint64,*TransactionSubmissionResult,error) {
	var ot TransactionSubmissionResult
	return 0,&ot,nil}

// MarshalJSON TransactionSubmissionResult
func (n TransactionSubmissionResult) MarshalJSON() ([]byte, error) {
	v := SubmitTransactionResult(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *TransactionSubmissionResult
func (n *TransactionSubmissionResult) UnmarshalJSON(data []byte) error {
	v := SubmitTransactionResult(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = TransactionSubmissionResult(v)
	return nil
}


// ----------------------------------------
//  Typedef: SubmissionErrorResult
// ----------------------------------------

// SubmissionErrorResult type
type SubmissionErrorResult RPCError

// NewSubmissionErrorResult factory
func NewSubmissionErrorResult() *SubmissionErrorResult {
	o := SubmissionErrorResult(*NewRPCError())
	return &o
}

// Serialize SubmissionErrorResult
func (n SubmissionErrorResult) Serialize(vb *VariableBlob) *VariableBlob {
	ox := RPCError(n)
	return ox.Serialize(vb)
}

// DeserializeSubmissionErrorResult function
func DeserializeSubmissionErrorResult(vb *VariableBlob) (uint64,*SubmissionErrorResult,error) {
	var ot SubmissionErrorResult
	i,n,err := DeserializeRPCError(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = SubmissionErrorResult(*n)
	return i,&ot,nil}

// MarshalJSON SubmissionErrorResult
func (n SubmissionErrorResult) MarshalJSON() ([]byte, error) {
	v := RPCError(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *SubmissionErrorResult
func (n *SubmissionErrorResult) UnmarshalJSON(data []byte) error {
	v := RPCError(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = SubmissionErrorResult(v)
	return nil
}


// ----------------------------------------
//  Variant: SubmissionResult
// ----------------------------------------

// SubmissionResult type
type SubmissionResult struct {
	Value interface{}
}

// NewSubmissionResult factory
func NewSubmissionResult() *SubmissionResult {
	v := SubmissionResult{}
	v.Value = NewReservedSubmissionResult()
	return &v
}

// Serialize SubmissionResult
func (n SubmissionResult) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *ReservedSubmissionResult:
			i = 0
		case *BlockSubmissionResult:
			i = 1
		case *TransactionSubmissionResult:
			i = 2
		case *QuerySubmissionResult:
			i = 3
		case *SubmissionErrorResult:
			i = 4
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName SubmissionResult
func (n SubmissionResult) TypeToName() (string) {
	switch n.Value.(type) {
		case *ReservedSubmissionResult:
			return "koinos::types::rpc::reserved_submission_result"
		case *BlockSubmissionResult:
			return "koinos::types::rpc::block_submission_result"
		case *TransactionSubmissionResult:
			return "koinos::types::rpc::transaction_submission_result"
		case *QuerySubmissionResult:
			return "koinos::types::rpc::query_submission_result"
		case *SubmissionErrorResult:
			return "koinos::types::rpc::submission_error_result"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON SubmissionResult
func (n SubmissionResult) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeSubmissionResult function
func DeserializeSubmissionResult(vb *VariableBlob) (uint64,*SubmissionResult,error) {
	var v SubmissionResult
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewReservedSubmissionResult()
		case 1:
			v.Value = NewBlockSubmissionResult()
		case 2:
			v.Value = NewTransactionSubmissionResult()
		case 3:
			ovb := (*vb)[i:]
			k,x,err := DeserializeQuerySubmissionResult(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 4:
			ovb := (*vb)[i:]
			k,x,err := DeserializeSubmissionErrorResult(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		default:
			return 0, &v, errors.New("unknown variant tag")
	}
	return uint64(i)+j,&v,nil
}

// UnmarshalJSON *SubmissionResult
func (n *SubmissionResult) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::types::rpc::reserved_submission_result":
			v := NewReservedSubmissionResult()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::rpc::block_submission_result":
			v := NewBlockSubmissionResult()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::rpc::transaction_submission_result":
			v := NewTransactionSubmissionResult()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::rpc::query_submission_result":
			v := NewQuerySubmissionResult()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::types::rpc::submission_error_result":
			v := NewSubmissionErrorResult()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

	return nil
}


// ----------------------------------------
//  Enum: SystemCallID
// ----------------------------------------

// {'name': 'system_call_id', 'entries': [{'name': 'prints', 'value': 2602735937, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'verify_block_header', 'value': 2519027790, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_block', 'value': 2494255093, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_transaction', 'value': 2643154394, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_reserved_operation', 'value': 2594724132, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_upload_contract_operation', 'value': 2658052407, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_execute_contract_operation', 'value': 2451064454, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_set_system_call_operation', 'value': 2507777116, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_put_object', 'value': 2535376802, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_get_object', 'value': 2540087547, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_get_next_object', 'value': 2577635560, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_get_prev_object', 'value': 2614326908, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'execute_contract', 'value': 2562796798, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_contract_args_size', 'value': 2601357273, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_contract_args', 'value': 2679873944, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'set_contract_return', 'value': 2672414186, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'exit_contract', 'value': 2564781488, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_head_info', 'value': 2507125293, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'hash', 'value': 2574716420, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'verify_block_sig', 'value': 2508417296, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'verify_merkle_root', 'value': 2574132409, 'doc': '', 'info': {'type': 'EnumEntry'}}], 'tref': {'name': ['koinos', 'types', 'uint32'], 'targs': None, 'info': {'type': 'Typeref'}}, 'doc': '', 'info': {'type': 'EnumClass'}}

// SystemCallID type
type SystemCallID UInt32

// NewSystemCallID factory
func NewSystemCallID() *SystemCallID {
	o := SystemCallID(2602735937)
	return &o
}

// SystemCallID values
const (
	SystemCallIDPrints SystemCallID = 2602735937
	SystemCallIDVerifyBlockHeader SystemCallID = 2519027790
	SystemCallIDApplyBlock SystemCallID = 2494255093
	SystemCallIDApplyTransaction SystemCallID = 2643154394
	SystemCallIDApplyReservedOperation SystemCallID = 2594724132
	SystemCallIDApplyUploadContractOperation SystemCallID = 2658052407
	SystemCallIDApplyExecuteContractOperation SystemCallID = 2451064454
	SystemCallIDApplySetSystemCallOperation SystemCallID = 2507777116
	SystemCallIDDbPutObject SystemCallID = 2535376802
	SystemCallIDDbGetObject SystemCallID = 2540087547
	SystemCallIDDbGetNextObject SystemCallID = 2577635560
	SystemCallIDDbGetPrevObject SystemCallID = 2614326908
	SystemCallIDExecuteContract SystemCallID = 2562796798
	SystemCallIDGetContractArgsSize SystemCallID = 2601357273
	SystemCallIDGetContractArgs SystemCallID = 2679873944
	SystemCallIDSetContractReturn SystemCallID = 2672414186
	SystemCallIDExitContract SystemCallID = 2564781488
	SystemCallIDGetHeadInfo SystemCallID = 2507125293
	SystemCallIDHash SystemCallID = 2574716420
	SystemCallIDVerifyBlockSig SystemCallID = 2508417296
	SystemCallIDVerifyMerkleRoot SystemCallID = 2574132409
)

// Serialize SystemCallID
func (n SystemCallID) Serialize(vb *VariableBlob) *VariableBlob {
	if !IsValidSystemCallID(n) {
		panic("Attempting to serialize an invalid value")
	}
	x := UInt32(n)
	return x.Serialize(vb)
}

// DeserializeSystemCallID function
func DeserializeSystemCallID(vb *VariableBlob) (uint64,*SystemCallID,error) {
	i,item,err := DeserializeUInt32(vb)
	var x SystemCallID
	if err != nil {
		return 0,&x,err
	}

	x = SystemCallID(*item)
	if !IsValidSystemCallID(x) {
		return i,&x,fmt.Errorf("invalid SystemCallID: %d", x)
	}
	return i,&x,nil
}

// MarshalJSON SystemCallID
func (n SystemCallID) MarshalJSON() ([]byte, error) {
	if !IsValidSystemCallID(n) {
		panic("Attempting to serialize an invalid value")
	}

	return json.Marshal(UInt32(n))
}

// UnmarshalJSON *SystemCallID
func (n *SystemCallID) UnmarshalJSON(b []byte) error {
	var o UInt32
	if err := json.Unmarshal(b, &o); err != nil {
		return err
	}

	ov := SystemCallID(o)

	if !IsValidSystemCallID(ov) {
		return fmt.Errorf("invalid SystemCallID: %d", o)
	}

	*n = ov
	return nil
}

// IsValidSystemCallID validator
func IsValidSystemCallID(v SystemCallID) bool {
	switch v {
		case SystemCallIDPrints:
			return true
		case SystemCallIDVerifyBlockHeader:
			return true
		case SystemCallIDApplyBlock:
			return true
		case SystemCallIDApplyTransaction:
			return true
		case SystemCallIDApplyReservedOperation:
			return true
		case SystemCallIDApplyUploadContractOperation:
			return true
		case SystemCallIDApplyExecuteContractOperation:
			return true
		case SystemCallIDApplySetSystemCallOperation:
			return true
		case SystemCallIDDbPutObject:
			return true
		case SystemCallIDDbGetObject:
			return true
		case SystemCallIDDbGetNextObject:
			return true
		case SystemCallIDDbGetPrevObject:
			return true
		case SystemCallIDExecuteContract:
			return true
		case SystemCallIDGetContractArgsSize:
			return true
		case SystemCallIDGetContractArgs:
			return true
		case SystemCallIDSetContractReturn:
			return true
		case SystemCallIDExitContract:
			return true
		case SystemCallIDGetHeadInfo:
			return true
		case SystemCallIDHash:
			return true
		case SystemCallIDVerifyBlockSig:
			return true
		case SystemCallIDVerifyMerkleRoot:
			return true
		default:
			return false
	}
}


// ----------------------------------------
//  Struct: VoidType
// ----------------------------------------

// VoidType type
type VoidType struct {
}

// NewVoidType factory
func NewVoidType() *VoidType {
	o := VoidType{}
	return &o
}

// Serialize VoidType
func (n VoidType) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeVoidType function
func DeserializeVoidType(vb *VariableBlob) (uint64,*VoidType,error) {
	var i uint64 = 0
	s := VoidType{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: PrintsArgs
// ----------------------------------------

// PrintsArgs type
type PrintsArgs struct {
    Message String `json:"message"`
}

// NewPrintsArgs factory
func NewPrintsArgs() *PrintsArgs {
	o := PrintsArgs{}
	o.Message = *NewString()
	return &o
}

// Serialize PrintsArgs
func (n PrintsArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Message.Serialize(vb)
	return vb
}

// DeserializePrintsArgs function
func DeserializePrintsArgs(vb *VariableBlob) (uint64,*PrintsArgs,error) {
	var i,j uint64 = 0,0
	s := PrintsArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tMessage,err := DeserializeString(&ovb); i+=j
	if err != nil {
		return 0, &PrintsArgs{}, err
	}
	s.Message = *tMessage
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: PrintsRet
// ----------------------------------------

// PrintsRet type
type PrintsRet VoidType

// NewPrintsRet factory
func NewPrintsRet() *PrintsRet {
	o := PrintsRet(*NewVoidType())
	return &o
}

// Serialize PrintsRet
func (n PrintsRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializePrintsRet function
func DeserializePrintsRet(vb *VariableBlob) (uint64,*PrintsRet,error) {
	var ot PrintsRet
	return 0,&ot,nil}

// MarshalJSON PrintsRet
func (n PrintsRet) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *PrintsRet
func (n *PrintsRet) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = PrintsRet(v)
	return nil
}


// ----------------------------------------
//  Struct: VerifyBlockSigArgs
// ----------------------------------------

// VerifyBlockSigArgs type
type VerifyBlockSigArgs struct {
    SigData VariableBlob `json:"sig_data"`
    Digest Multihash `json:"digest"`
}

// NewVerifyBlockSigArgs factory
func NewVerifyBlockSigArgs() *VerifyBlockSigArgs {
	o := VerifyBlockSigArgs{}
	o.SigData = *NewVariableBlob()
	o.Digest = *NewMultihash()
	return &o
}

// Serialize VerifyBlockSigArgs
func (n VerifyBlockSigArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.SigData.Serialize(vb)
	vb = n.Digest.Serialize(vb)
	return vb
}

// DeserializeVerifyBlockSigArgs function
func DeserializeVerifyBlockSigArgs(vb *VariableBlob) (uint64,*VerifyBlockSigArgs,error) {
	var i,j uint64 = 0,0
	s := VerifyBlockSigArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tSigData,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &VerifyBlockSigArgs{}, err
	}
	s.SigData = *tSigData
	ovb = (*vb)[i:]
	j,tDigest,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &VerifyBlockSigArgs{}, err
	}
	s.Digest = *tDigest
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: VerifyBlockSigRet
// ----------------------------------------

// VerifyBlockSigRet type
type VerifyBlockSigRet Boolean

// NewVerifyBlockSigRet factory
func NewVerifyBlockSigRet() *VerifyBlockSigRet {
	o := VerifyBlockSigRet(*NewBoolean())
	return &o
}

// Serialize VerifyBlockSigRet
func (n VerifyBlockSigRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := Boolean(n)
	return ox.Serialize(vb)
}

// DeserializeVerifyBlockSigRet function
func DeserializeVerifyBlockSigRet(vb *VariableBlob) (uint64,*VerifyBlockSigRet,error) {
	var ot VerifyBlockSigRet
	i,n,err := DeserializeBoolean(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = VerifyBlockSigRet(*n)
	return i,&ot,nil}

// MarshalJSON VerifyBlockSigRet
func (n VerifyBlockSigRet) MarshalJSON() ([]byte, error) {
	v := Boolean(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *VerifyBlockSigRet
func (n *VerifyBlockSigRet) UnmarshalJSON(data []byte) error {
	v := Boolean(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = VerifyBlockSigRet(v)
	return nil
}


// ----------------------------------------
//  Struct: VerifyMerkleRootArgs
// ----------------------------------------

// VerifyMerkleRootArgs type
type VerifyMerkleRootArgs struct {
    Root Multihash `json:"root"`
    Hashes VectorMultihash `json:"hashes"`
}

// NewVerifyMerkleRootArgs factory
func NewVerifyMerkleRootArgs() *VerifyMerkleRootArgs {
	o := VerifyMerkleRootArgs{}
	o.Root = *NewMultihash()
	o.Hashes = *NewVectorMultihash()
	return &o
}

// Serialize VerifyMerkleRootArgs
func (n VerifyMerkleRootArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Root.Serialize(vb)
	vb = n.Hashes.Serialize(vb)
	return vb
}

// DeserializeVerifyMerkleRootArgs function
func DeserializeVerifyMerkleRootArgs(vb *VariableBlob) (uint64,*VerifyMerkleRootArgs,error) {
	var i,j uint64 = 0,0
	s := VerifyMerkleRootArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tRoot,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &VerifyMerkleRootArgs{}, err
	}
	s.Root = *tRoot
	ovb = (*vb)[i:]
	j,tHashes,err := DeserializeVectorMultihash(&ovb); i+=j
	if err != nil {
		return 0, &VerifyMerkleRootArgs{}, err
	}
	s.Hashes = *tHashes
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: VerifyMerkleRootRet
// ----------------------------------------

// VerifyMerkleRootRet type
type VerifyMerkleRootRet Boolean

// NewVerifyMerkleRootRet factory
func NewVerifyMerkleRootRet() *VerifyMerkleRootRet {
	o := VerifyMerkleRootRet(*NewBoolean())
	return &o
}

// Serialize VerifyMerkleRootRet
func (n VerifyMerkleRootRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := Boolean(n)
	return ox.Serialize(vb)
}

// DeserializeVerifyMerkleRootRet function
func DeserializeVerifyMerkleRootRet(vb *VariableBlob) (uint64,*VerifyMerkleRootRet,error) {
	var ot VerifyMerkleRootRet
	i,n,err := DeserializeBoolean(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = VerifyMerkleRootRet(*n)
	return i,&ot,nil}

// MarshalJSON VerifyMerkleRootRet
func (n VerifyMerkleRootRet) MarshalJSON() ([]byte, error) {
	v := Boolean(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *VerifyMerkleRootRet
func (n *VerifyMerkleRootRet) UnmarshalJSON(data []byte) error {
	v := Boolean(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = VerifyMerkleRootRet(v)
	return nil
}


// ----------------------------------------
//  Struct: ApplyBlockArgs
// ----------------------------------------

// ApplyBlockArgs type
type ApplyBlockArgs struct {
    Block Block `json:"block"`
    EnableCheckPassiveData Boolean `json:"enable_check_passive_data"`
    EnableCheckBlockSignature Boolean `json:"enable_check_block_signature"`
    EnableCheckTransactionSignatures Boolean `json:"enable_check_transaction_signatures"`
}

// NewApplyBlockArgs factory
func NewApplyBlockArgs() *ApplyBlockArgs {
	o := ApplyBlockArgs{}
	o.Block = *NewBlock()
	o.EnableCheckPassiveData = *NewBoolean()
	o.EnableCheckBlockSignature = *NewBoolean()
	o.EnableCheckTransactionSignatures = *NewBoolean()
	return &o
}

// Serialize ApplyBlockArgs
func (n ApplyBlockArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Block.Serialize(vb)
	vb = n.EnableCheckPassiveData.Serialize(vb)
	vb = n.EnableCheckBlockSignature.Serialize(vb)
	vb = n.EnableCheckTransactionSignatures.Serialize(vb)
	return vb
}

// DeserializeApplyBlockArgs function
func DeserializeApplyBlockArgs(vb *VariableBlob) (uint64,*ApplyBlockArgs,error) {
	var i,j uint64 = 0,0
	s := ApplyBlockArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tBlock,err := DeserializeBlock(&ovb); i+=j
	if err != nil {
		return 0, &ApplyBlockArgs{}, err
	}
	s.Block = *tBlock
	ovb = (*vb)[i:]
	j,tEnableCheckPassiveData,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &ApplyBlockArgs{}, err
	}
	s.EnableCheckPassiveData = *tEnableCheckPassiveData
	ovb = (*vb)[i:]
	j,tEnableCheckBlockSignature,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &ApplyBlockArgs{}, err
	}
	s.EnableCheckBlockSignature = *tEnableCheckBlockSignature
	ovb = (*vb)[i:]
	j,tEnableCheckTransactionSignatures,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &ApplyBlockArgs{}, err
	}
	s.EnableCheckTransactionSignatures = *tEnableCheckTransactionSignatures
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: ApplyBlockRet
// ----------------------------------------

// ApplyBlockRet type
type ApplyBlockRet VoidType

// NewApplyBlockRet factory
func NewApplyBlockRet() *ApplyBlockRet {
	o := ApplyBlockRet(*NewVoidType())
	return &o
}

// Serialize ApplyBlockRet
func (n ApplyBlockRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeApplyBlockRet function
func DeserializeApplyBlockRet(vb *VariableBlob) (uint64,*ApplyBlockRet,error) {
	var ot ApplyBlockRet
	return 0,&ot,nil}

// MarshalJSON ApplyBlockRet
func (n ApplyBlockRet) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ApplyBlockRet
func (n *ApplyBlockRet) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ApplyBlockRet(v)
	return nil
}


// ----------------------------------------
//  Struct: ApplyTransactionArgs
// ----------------------------------------

// ApplyTransactionArgs type
type ApplyTransactionArgs struct {
    Trx OpaqueTransaction `json:"trx"`
}

// NewApplyTransactionArgs factory
func NewApplyTransactionArgs() *ApplyTransactionArgs {
	o := ApplyTransactionArgs{}
	o.Trx = *NewOpaqueTransaction()
	return &o
}

// Serialize ApplyTransactionArgs
func (n ApplyTransactionArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Trx.Serialize(vb)
	return vb
}

// DeserializeApplyTransactionArgs function
func DeserializeApplyTransactionArgs(vb *VariableBlob) (uint64,*ApplyTransactionArgs,error) {
	var i,j uint64 = 0,0
	s := ApplyTransactionArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTrx,err := DeserializeOpaqueTransaction(&ovb); i+=j
	if err != nil {
		return 0, &ApplyTransactionArgs{}, err
	}
	s.Trx = *tTrx
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: ApplyTransactionRet
// ----------------------------------------

// ApplyTransactionRet type
type ApplyTransactionRet VoidType

// NewApplyTransactionRet factory
func NewApplyTransactionRet() *ApplyTransactionRet {
	o := ApplyTransactionRet(*NewVoidType())
	return &o
}

// Serialize ApplyTransactionRet
func (n ApplyTransactionRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeApplyTransactionRet function
func DeserializeApplyTransactionRet(vb *VariableBlob) (uint64,*ApplyTransactionRet,error) {
	var ot ApplyTransactionRet
	return 0,&ot,nil}

// MarshalJSON ApplyTransactionRet
func (n ApplyTransactionRet) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ApplyTransactionRet
func (n *ApplyTransactionRet) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ApplyTransactionRet(v)
	return nil
}


// ----------------------------------------
//  Struct: ApplyUploadContractOperationArgs
// ----------------------------------------

// ApplyUploadContractOperationArgs type
type ApplyUploadContractOperationArgs struct {
    Op CreateSystemContractOperation `json:"op"`
}

// NewApplyUploadContractOperationArgs factory
func NewApplyUploadContractOperationArgs() *ApplyUploadContractOperationArgs {
	o := ApplyUploadContractOperationArgs{}
	o.Op = *NewCreateSystemContractOperation()
	return &o
}

// Serialize ApplyUploadContractOperationArgs
func (n ApplyUploadContractOperationArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Op.Serialize(vb)
	return vb
}

// DeserializeApplyUploadContractOperationArgs function
func DeserializeApplyUploadContractOperationArgs(vb *VariableBlob) (uint64,*ApplyUploadContractOperationArgs,error) {
	var i,j uint64 = 0,0
	s := ApplyUploadContractOperationArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tOp,err := DeserializeCreateSystemContractOperation(&ovb); i+=j
	if err != nil {
		return 0, &ApplyUploadContractOperationArgs{}, err
	}
	s.Op = *tOp
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: ApplyUploadContractOperationRet
// ----------------------------------------

// ApplyUploadContractOperationRet type
type ApplyUploadContractOperationRet VoidType

// NewApplyUploadContractOperationRet factory
func NewApplyUploadContractOperationRet() *ApplyUploadContractOperationRet {
	o := ApplyUploadContractOperationRet(*NewVoidType())
	return &o
}

// Serialize ApplyUploadContractOperationRet
func (n ApplyUploadContractOperationRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeApplyUploadContractOperationRet function
func DeserializeApplyUploadContractOperationRet(vb *VariableBlob) (uint64,*ApplyUploadContractOperationRet,error) {
	var ot ApplyUploadContractOperationRet
	return 0,&ot,nil}

// MarshalJSON ApplyUploadContractOperationRet
func (n ApplyUploadContractOperationRet) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ApplyUploadContractOperationRet
func (n *ApplyUploadContractOperationRet) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ApplyUploadContractOperationRet(v)
	return nil
}


// ----------------------------------------
//  Struct: ApplyReservedOperationArgs
// ----------------------------------------

// ApplyReservedOperationArgs type
type ApplyReservedOperationArgs struct {
    Op ReservedOperation `json:"op"`
}

// NewApplyReservedOperationArgs factory
func NewApplyReservedOperationArgs() *ApplyReservedOperationArgs {
	o := ApplyReservedOperationArgs{}
	o.Op = *NewReservedOperation()
	return &o
}

// Serialize ApplyReservedOperationArgs
func (n ApplyReservedOperationArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Op.Serialize(vb)
	return vb
}

// DeserializeApplyReservedOperationArgs function
func DeserializeApplyReservedOperationArgs(vb *VariableBlob) (uint64,*ApplyReservedOperationArgs,error) {
	var i uint64 = 0
	s := ApplyReservedOperationArgs{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: ApplyReservedOperationRet
// ----------------------------------------

// ApplyReservedOperationRet type
type ApplyReservedOperationRet VoidType

// NewApplyReservedOperationRet factory
func NewApplyReservedOperationRet() *ApplyReservedOperationRet {
	o := ApplyReservedOperationRet(*NewVoidType())
	return &o
}

// Serialize ApplyReservedOperationRet
func (n ApplyReservedOperationRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeApplyReservedOperationRet function
func DeserializeApplyReservedOperationRet(vb *VariableBlob) (uint64,*ApplyReservedOperationRet,error) {
	var ot ApplyReservedOperationRet
	return 0,&ot,nil}

// MarshalJSON ApplyReservedOperationRet
func (n ApplyReservedOperationRet) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ApplyReservedOperationRet
func (n *ApplyReservedOperationRet) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ApplyReservedOperationRet(v)
	return nil
}


// ----------------------------------------
//  Struct: ApplyExecuteContractOperationArgs
// ----------------------------------------

// ApplyExecuteContractOperationArgs type
type ApplyExecuteContractOperationArgs struct {
    Op ContractCallOperation `json:"op"`
}

// NewApplyExecuteContractOperationArgs factory
func NewApplyExecuteContractOperationArgs() *ApplyExecuteContractOperationArgs {
	o := ApplyExecuteContractOperationArgs{}
	o.Op = *NewContractCallOperation()
	return &o
}

// Serialize ApplyExecuteContractOperationArgs
func (n ApplyExecuteContractOperationArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Op.Serialize(vb)
	return vb
}

// DeserializeApplyExecuteContractOperationArgs function
func DeserializeApplyExecuteContractOperationArgs(vb *VariableBlob) (uint64,*ApplyExecuteContractOperationArgs,error) {
	var i,j uint64 = 0,0
	s := ApplyExecuteContractOperationArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tOp,err := DeserializeContractCallOperation(&ovb); i+=j
	if err != nil {
		return 0, &ApplyExecuteContractOperationArgs{}, err
	}
	s.Op = *tOp
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: ApplyExecuteContractOperationRet
// ----------------------------------------

// ApplyExecuteContractOperationRet type
type ApplyExecuteContractOperationRet VoidType

// NewApplyExecuteContractOperationRet factory
func NewApplyExecuteContractOperationRet() *ApplyExecuteContractOperationRet {
	o := ApplyExecuteContractOperationRet(*NewVoidType())
	return &o
}

// Serialize ApplyExecuteContractOperationRet
func (n ApplyExecuteContractOperationRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeApplyExecuteContractOperationRet function
func DeserializeApplyExecuteContractOperationRet(vb *VariableBlob) (uint64,*ApplyExecuteContractOperationRet,error) {
	var ot ApplyExecuteContractOperationRet
	return 0,&ot,nil}

// MarshalJSON ApplyExecuteContractOperationRet
func (n ApplyExecuteContractOperationRet) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ApplyExecuteContractOperationRet
func (n *ApplyExecuteContractOperationRet) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ApplyExecuteContractOperationRet(v)
	return nil
}


// ----------------------------------------
//  Struct: ApplySetSystemCallOperationArgs
// ----------------------------------------

// ApplySetSystemCallOperationArgs type
type ApplySetSystemCallOperationArgs struct {
    Op SetSystemCallOperation `json:"op"`
}

// NewApplySetSystemCallOperationArgs factory
func NewApplySetSystemCallOperationArgs() *ApplySetSystemCallOperationArgs {
	o := ApplySetSystemCallOperationArgs{}
	o.Op = *NewSetSystemCallOperation()
	return &o
}

// Serialize ApplySetSystemCallOperationArgs
func (n ApplySetSystemCallOperationArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Op.Serialize(vb)
	return vb
}

// DeserializeApplySetSystemCallOperationArgs function
func DeserializeApplySetSystemCallOperationArgs(vb *VariableBlob) (uint64,*ApplySetSystemCallOperationArgs,error) {
	var i,j uint64 = 0,0
	s := ApplySetSystemCallOperationArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tOp,err := DeserializeSetSystemCallOperation(&ovb); i+=j
	if err != nil {
		return 0, &ApplySetSystemCallOperationArgs{}, err
	}
	s.Op = *tOp
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: ApplySetSystemCallOperationRet
// ----------------------------------------

// ApplySetSystemCallOperationRet type
type ApplySetSystemCallOperationRet VoidType

// NewApplySetSystemCallOperationRet factory
func NewApplySetSystemCallOperationRet() *ApplySetSystemCallOperationRet {
	o := ApplySetSystemCallOperationRet(*NewVoidType())
	return &o
}

// Serialize ApplySetSystemCallOperationRet
func (n ApplySetSystemCallOperationRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeApplySetSystemCallOperationRet function
func DeserializeApplySetSystemCallOperationRet(vb *VariableBlob) (uint64,*ApplySetSystemCallOperationRet,error) {
	var ot ApplySetSystemCallOperationRet
	return 0,&ot,nil}

// MarshalJSON ApplySetSystemCallOperationRet
func (n ApplySetSystemCallOperationRet) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ApplySetSystemCallOperationRet
func (n *ApplySetSystemCallOperationRet) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ApplySetSystemCallOperationRet(v)
	return nil
}


// ----------------------------------------
//  Struct: DbPutObjectArgs
// ----------------------------------------

// DbPutObjectArgs type
type DbPutObjectArgs struct {
    Space UInt256 `json:"space"`
    Key UInt256 `json:"key"`
    Obj VariableBlob `json:"obj"`
}

// NewDbPutObjectArgs factory
func NewDbPutObjectArgs() *DbPutObjectArgs {
	o := DbPutObjectArgs{}
	o.Space = *NewUInt256()
	o.Key = *NewUInt256()
	o.Obj = *NewVariableBlob()
	return &o
}

// Serialize DbPutObjectArgs
func (n DbPutObjectArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Space.Serialize(vb)
	vb = n.Key.Serialize(vb)
	vb = n.Obj.Serialize(vb)
	return vb
}

// DeserializeDbPutObjectArgs function
func DeserializeDbPutObjectArgs(vb *VariableBlob) (uint64,*DbPutObjectArgs,error) {
	var i,j uint64 = 0,0
	s := DbPutObjectArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tSpace,err := DeserializeUInt256(&ovb); i+=j
	if err != nil {
		return 0, &DbPutObjectArgs{}, err
	}
	s.Space = *tSpace
	ovb = (*vb)[i:]
	j,tKey,err := DeserializeUInt256(&ovb); i+=j
	if err != nil {
		return 0, &DbPutObjectArgs{}, err
	}
	s.Key = *tKey
	ovb = (*vb)[i:]
	j,tObj,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &DbPutObjectArgs{}, err
	}
	s.Obj = *tObj
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: DbPutObjectRet
// ----------------------------------------

// DbPutObjectRet type
type DbPutObjectRet Boolean

// NewDbPutObjectRet factory
func NewDbPutObjectRet() *DbPutObjectRet {
	o := DbPutObjectRet(*NewBoolean())
	return &o
}

// Serialize DbPutObjectRet
func (n DbPutObjectRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := Boolean(n)
	return ox.Serialize(vb)
}

// DeserializeDbPutObjectRet function
func DeserializeDbPutObjectRet(vb *VariableBlob) (uint64,*DbPutObjectRet,error) {
	var ot DbPutObjectRet
	i,n,err := DeserializeBoolean(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = DbPutObjectRet(*n)
	return i,&ot,nil}

// MarshalJSON DbPutObjectRet
func (n DbPutObjectRet) MarshalJSON() ([]byte, error) {
	v := Boolean(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *DbPutObjectRet
func (n *DbPutObjectRet) UnmarshalJSON(data []byte) error {
	v := Boolean(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = DbPutObjectRet(v)
	return nil
}


// ----------------------------------------
//  Struct: DbGetObjectArgs
// ----------------------------------------

// DbGetObjectArgs type
type DbGetObjectArgs struct {
    Space UInt256 `json:"space"`
    Key UInt256 `json:"key"`
    ObjectSizeHint Int32 `json:"object_size_hint"`
}

// NewDbGetObjectArgs factory
func NewDbGetObjectArgs() *DbGetObjectArgs {
	o := DbGetObjectArgs{}
	o.Space = *NewUInt256()
	o.Key = *NewUInt256()
	o.ObjectSizeHint = *NewInt32()
	return &o
}

// Serialize DbGetObjectArgs
func (n DbGetObjectArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Space.Serialize(vb)
	vb = n.Key.Serialize(vb)
	vb = n.ObjectSizeHint.Serialize(vb)
	return vb
}

// DeserializeDbGetObjectArgs function
func DeserializeDbGetObjectArgs(vb *VariableBlob) (uint64,*DbGetObjectArgs,error) {
	var i,j uint64 = 0,0
	s := DbGetObjectArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tSpace,err := DeserializeUInt256(&ovb); i+=j
	if err != nil {
		return 0, &DbGetObjectArgs{}, err
	}
	s.Space = *tSpace
	ovb = (*vb)[i:]
	j,tKey,err := DeserializeUInt256(&ovb); i+=j
	if err != nil {
		return 0, &DbGetObjectArgs{}, err
	}
	s.Key = *tKey
	ovb = (*vb)[i:]
	j,tObjectSizeHint,err := DeserializeInt32(&ovb); i+=j
	if err != nil {
		return 0, &DbGetObjectArgs{}, err
	}
	s.ObjectSizeHint = *tObjectSizeHint
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: DbGetObjectRet
// ----------------------------------------

// DbGetObjectRet type
type DbGetObjectRet VariableBlob

// NewDbGetObjectRet factory
func NewDbGetObjectRet() *DbGetObjectRet {
	o := DbGetObjectRet(*NewVariableBlob())
	return &o
}

// Serialize DbGetObjectRet
func (n DbGetObjectRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VariableBlob(n)
	return ox.Serialize(vb)
}

// DeserializeDbGetObjectRet function
func DeserializeDbGetObjectRet(vb *VariableBlob) (uint64,*DbGetObjectRet,error) {
	var ot DbGetObjectRet
	i,n,err := DeserializeVariableBlob(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = DbGetObjectRet(*n)
	return i,&ot,nil}

// MarshalJSON DbGetObjectRet
func (n DbGetObjectRet) MarshalJSON() ([]byte, error) {
	v := VariableBlob(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *DbGetObjectRet
func (n *DbGetObjectRet) UnmarshalJSON(data []byte) error {
	v := VariableBlob(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = DbGetObjectRet(v)
	return nil
}


// ----------------------------------------
//  Typedef: DbGetNextObjectArgs
// ----------------------------------------

// DbGetNextObjectArgs type
type DbGetNextObjectArgs DbGetObjectArgs

// NewDbGetNextObjectArgs factory
func NewDbGetNextObjectArgs() *DbGetNextObjectArgs {
	o := DbGetNextObjectArgs(*NewDbGetObjectArgs())
	return &o
}

// Serialize DbGetNextObjectArgs
func (n DbGetNextObjectArgs) Serialize(vb *VariableBlob) *VariableBlob {
	ox := DbGetObjectArgs(n)
	return ox.Serialize(vb)
}

// DeserializeDbGetNextObjectArgs function
func DeserializeDbGetNextObjectArgs(vb *VariableBlob) (uint64,*DbGetNextObjectArgs,error) {
	var ot DbGetNextObjectArgs
	i,n,err := DeserializeDbGetObjectArgs(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = DbGetNextObjectArgs(*n)
	return i,&ot,nil}

// MarshalJSON DbGetNextObjectArgs
func (n DbGetNextObjectArgs) MarshalJSON() ([]byte, error) {
	v := DbGetObjectArgs(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *DbGetNextObjectArgs
func (n *DbGetNextObjectArgs) UnmarshalJSON(data []byte) error {
	v := DbGetObjectArgs(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = DbGetNextObjectArgs(v)
	return nil
}


// ----------------------------------------
//  Typedef: DbGetNextObjectRet
// ----------------------------------------

// DbGetNextObjectRet type
type DbGetNextObjectRet DbGetObjectRet

// NewDbGetNextObjectRet factory
func NewDbGetNextObjectRet() *DbGetNextObjectRet {
	o := DbGetNextObjectRet(*NewDbGetObjectRet())
	return &o
}

// Serialize DbGetNextObjectRet
func (n DbGetNextObjectRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := DbGetObjectRet(n)
	return ox.Serialize(vb)
}

// DeserializeDbGetNextObjectRet function
func DeserializeDbGetNextObjectRet(vb *VariableBlob) (uint64,*DbGetNextObjectRet,error) {
	var ot DbGetNextObjectRet
	i,n,err := DeserializeDbGetObjectRet(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = DbGetNextObjectRet(*n)
	return i,&ot,nil}

// MarshalJSON DbGetNextObjectRet
func (n DbGetNextObjectRet) MarshalJSON() ([]byte, error) {
	v := DbGetObjectRet(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *DbGetNextObjectRet
func (n *DbGetNextObjectRet) UnmarshalJSON(data []byte) error {
	v := DbGetObjectRet(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = DbGetNextObjectRet(v)
	return nil
}


// ----------------------------------------
//  Typedef: DbGetPrevObjectArgs
// ----------------------------------------

// DbGetPrevObjectArgs type
type DbGetPrevObjectArgs DbGetObjectArgs

// NewDbGetPrevObjectArgs factory
func NewDbGetPrevObjectArgs() *DbGetPrevObjectArgs {
	o := DbGetPrevObjectArgs(*NewDbGetObjectArgs())
	return &o
}

// Serialize DbGetPrevObjectArgs
func (n DbGetPrevObjectArgs) Serialize(vb *VariableBlob) *VariableBlob {
	ox := DbGetObjectArgs(n)
	return ox.Serialize(vb)
}

// DeserializeDbGetPrevObjectArgs function
func DeserializeDbGetPrevObjectArgs(vb *VariableBlob) (uint64,*DbGetPrevObjectArgs,error) {
	var ot DbGetPrevObjectArgs
	i,n,err := DeserializeDbGetObjectArgs(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = DbGetPrevObjectArgs(*n)
	return i,&ot,nil}

// MarshalJSON DbGetPrevObjectArgs
func (n DbGetPrevObjectArgs) MarshalJSON() ([]byte, error) {
	v := DbGetObjectArgs(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *DbGetPrevObjectArgs
func (n *DbGetPrevObjectArgs) UnmarshalJSON(data []byte) error {
	v := DbGetObjectArgs(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = DbGetPrevObjectArgs(v)
	return nil
}


// ----------------------------------------
//  Typedef: DbGetPrevObjectRet
// ----------------------------------------

// DbGetPrevObjectRet type
type DbGetPrevObjectRet DbGetObjectRet

// NewDbGetPrevObjectRet factory
func NewDbGetPrevObjectRet() *DbGetPrevObjectRet {
	o := DbGetPrevObjectRet(*NewDbGetObjectRet())
	return &o
}

// Serialize DbGetPrevObjectRet
func (n DbGetPrevObjectRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := DbGetObjectRet(n)
	return ox.Serialize(vb)
}

// DeserializeDbGetPrevObjectRet function
func DeserializeDbGetPrevObjectRet(vb *VariableBlob) (uint64,*DbGetPrevObjectRet,error) {
	var ot DbGetPrevObjectRet
	i,n,err := DeserializeDbGetObjectRet(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = DbGetPrevObjectRet(*n)
	return i,&ot,nil}

// MarshalJSON DbGetPrevObjectRet
func (n DbGetPrevObjectRet) MarshalJSON() ([]byte, error) {
	v := DbGetObjectRet(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *DbGetPrevObjectRet
func (n *DbGetPrevObjectRet) UnmarshalJSON(data []byte) error {
	v := DbGetObjectRet(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = DbGetPrevObjectRet(v)
	return nil
}


// ----------------------------------------
//  Struct: ExecuteContractArgs
// ----------------------------------------

// ExecuteContractArgs type
type ExecuteContractArgs struct {
    ContractID ContractIDType `json:"contract_id"`
    EntryPoint UInt32 `json:"entry_point"`
    Args VariableBlob `json:"args"`
}

// NewExecuteContractArgs factory
func NewExecuteContractArgs() *ExecuteContractArgs {
	o := ExecuteContractArgs{}
	o.ContractID = *NewContractIDType()
	o.EntryPoint = *NewUInt32()
	o.Args = *NewVariableBlob()
	return &o
}

// Serialize ExecuteContractArgs
func (n ExecuteContractArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ContractID.Serialize(vb)
	vb = n.EntryPoint.Serialize(vb)
	vb = n.Args.Serialize(vb)
	return vb
}

// DeserializeExecuteContractArgs function
func DeserializeExecuteContractArgs(vb *VariableBlob) (uint64,*ExecuteContractArgs,error) {
	var i,j uint64 = 0,0
	s := ExecuteContractArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tContractID,err := DeserializeContractIDType(&ovb); i+=j
	if err != nil {
		return 0, &ExecuteContractArgs{}, err
	}
	s.ContractID = *tContractID
	ovb = (*vb)[i:]
	j,tEntryPoint,err := DeserializeUInt32(&ovb); i+=j
	if err != nil {
		return 0, &ExecuteContractArgs{}, err
	}
	s.EntryPoint = *tEntryPoint
	ovb = (*vb)[i:]
	j,tArgs,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &ExecuteContractArgs{}, err
	}
	s.Args = *tArgs
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: ExecuteContractRet
// ----------------------------------------

// ExecuteContractRet type
type ExecuteContractRet VariableBlob

// NewExecuteContractRet factory
func NewExecuteContractRet() *ExecuteContractRet {
	o := ExecuteContractRet(*NewVariableBlob())
	return &o
}

// Serialize ExecuteContractRet
func (n ExecuteContractRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VariableBlob(n)
	return ox.Serialize(vb)
}

// DeserializeExecuteContractRet function
func DeserializeExecuteContractRet(vb *VariableBlob) (uint64,*ExecuteContractRet,error) {
	var ot ExecuteContractRet
	i,n,err := DeserializeVariableBlob(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = ExecuteContractRet(*n)
	return i,&ot,nil}

// MarshalJSON ExecuteContractRet
func (n ExecuteContractRet) MarshalJSON() ([]byte, error) {
	v := VariableBlob(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ExecuteContractRet
func (n *ExecuteContractRet) UnmarshalJSON(data []byte) error {
	v := VariableBlob(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ExecuteContractRet(v)
	return nil
}


// ----------------------------------------
//  Typedef: GetContractArgsSizeArgs
// ----------------------------------------

// GetContractArgsSizeArgs type
type GetContractArgsSizeArgs VoidType

// NewGetContractArgsSizeArgs factory
func NewGetContractArgsSizeArgs() *GetContractArgsSizeArgs {
	o := GetContractArgsSizeArgs(*NewVoidType())
	return &o
}

// Serialize GetContractArgsSizeArgs
func (n GetContractArgsSizeArgs) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeGetContractArgsSizeArgs function
func DeserializeGetContractArgsSizeArgs(vb *VariableBlob) (uint64,*GetContractArgsSizeArgs,error) {
	var ot GetContractArgsSizeArgs
	return 0,&ot,nil}

// MarshalJSON GetContractArgsSizeArgs
func (n GetContractArgsSizeArgs) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetContractArgsSizeArgs
func (n *GetContractArgsSizeArgs) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetContractArgsSizeArgs(v)
	return nil
}


// ----------------------------------------
//  Typedef: GetContractArgsSizeRet
// ----------------------------------------

// GetContractArgsSizeRet type
type GetContractArgsSizeRet UInt32

// NewGetContractArgsSizeRet factory
func NewGetContractArgsSizeRet() *GetContractArgsSizeRet {
	o := GetContractArgsSizeRet(*NewUInt32())
	return &o
}

// Serialize GetContractArgsSizeRet
func (n GetContractArgsSizeRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := UInt32(n)
	return ox.Serialize(vb)
}

// DeserializeGetContractArgsSizeRet function
func DeserializeGetContractArgsSizeRet(vb *VariableBlob) (uint64,*GetContractArgsSizeRet,error) {
	var ot GetContractArgsSizeRet
	i,n,err := DeserializeUInt32(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = GetContractArgsSizeRet(*n)
	return i,&ot,nil}

// MarshalJSON GetContractArgsSizeRet
func (n GetContractArgsSizeRet) MarshalJSON() ([]byte, error) {
	v := UInt32(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetContractArgsSizeRet
func (n *GetContractArgsSizeRet) UnmarshalJSON(data []byte) error {
	v := UInt32(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetContractArgsSizeRet(v)
	return nil
}


// ----------------------------------------
//  Typedef: GetContractArgsArgs
// ----------------------------------------

// GetContractArgsArgs type
type GetContractArgsArgs VoidType

// NewGetContractArgsArgs factory
func NewGetContractArgsArgs() *GetContractArgsArgs {
	o := GetContractArgsArgs(*NewVoidType())
	return &o
}

// Serialize GetContractArgsArgs
func (n GetContractArgsArgs) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeGetContractArgsArgs function
func DeserializeGetContractArgsArgs(vb *VariableBlob) (uint64,*GetContractArgsArgs,error) {
	var ot GetContractArgsArgs
	return 0,&ot,nil}

// MarshalJSON GetContractArgsArgs
func (n GetContractArgsArgs) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetContractArgsArgs
func (n *GetContractArgsArgs) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetContractArgsArgs(v)
	return nil
}


// ----------------------------------------
//  Typedef: GetContractArgsRet
// ----------------------------------------

// GetContractArgsRet type
type GetContractArgsRet VariableBlob

// NewGetContractArgsRet factory
func NewGetContractArgsRet() *GetContractArgsRet {
	o := GetContractArgsRet(*NewVariableBlob())
	return &o
}

// Serialize GetContractArgsRet
func (n GetContractArgsRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VariableBlob(n)
	return ox.Serialize(vb)
}

// DeserializeGetContractArgsRet function
func DeserializeGetContractArgsRet(vb *VariableBlob) (uint64,*GetContractArgsRet,error) {
	var ot GetContractArgsRet
	i,n,err := DeserializeVariableBlob(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = GetContractArgsRet(*n)
	return i,&ot,nil}

// MarshalJSON GetContractArgsRet
func (n GetContractArgsRet) MarshalJSON() ([]byte, error) {
	v := VariableBlob(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetContractArgsRet
func (n *GetContractArgsRet) UnmarshalJSON(data []byte) error {
	v := VariableBlob(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetContractArgsRet(v)
	return nil
}


// ----------------------------------------
//  Struct: SetContractReturnArgs
// ----------------------------------------

// SetContractReturnArgs type
type SetContractReturnArgs struct {
    Ret VariableBlob `json:"ret"`
}

// NewSetContractReturnArgs factory
func NewSetContractReturnArgs() *SetContractReturnArgs {
	o := SetContractReturnArgs{}
	o.Ret = *NewVariableBlob()
	return &o
}

// Serialize SetContractReturnArgs
func (n SetContractReturnArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Ret.Serialize(vb)
	return vb
}

// DeserializeSetContractReturnArgs function
func DeserializeSetContractReturnArgs(vb *VariableBlob) (uint64,*SetContractReturnArgs,error) {
	var i,j uint64 = 0,0
	s := SetContractReturnArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tRet,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &SetContractReturnArgs{}, err
	}
	s.Ret = *tRet
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: SetContractReturnRet
// ----------------------------------------

// SetContractReturnRet type
type SetContractReturnRet VoidType

// NewSetContractReturnRet factory
func NewSetContractReturnRet() *SetContractReturnRet {
	o := SetContractReturnRet(*NewVoidType())
	return &o
}

// Serialize SetContractReturnRet
func (n SetContractReturnRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeSetContractReturnRet function
func DeserializeSetContractReturnRet(vb *VariableBlob) (uint64,*SetContractReturnRet,error) {
	var ot SetContractReturnRet
	return 0,&ot,nil}

// MarshalJSON SetContractReturnRet
func (n SetContractReturnRet) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *SetContractReturnRet
func (n *SetContractReturnRet) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = SetContractReturnRet(v)
	return nil
}


// ----------------------------------------
//  Struct: ExitContractArgs
// ----------------------------------------

// ExitContractArgs type
type ExitContractArgs struct {
    ExitCode UInt8 `json:"exit_code"`
}

// NewExitContractArgs factory
func NewExitContractArgs() *ExitContractArgs {
	o := ExitContractArgs{}
	o.ExitCode = *NewUInt8()
	return &o
}

// Serialize ExitContractArgs
func (n ExitContractArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ExitCode.Serialize(vb)
	return vb
}

// DeserializeExitContractArgs function
func DeserializeExitContractArgs(vb *VariableBlob) (uint64,*ExitContractArgs,error) {
	var i,j uint64 = 0,0
	s := ExitContractArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tExitCode,err := DeserializeUInt8(&ovb); i+=j
	if err != nil {
		return 0, &ExitContractArgs{}, err
	}
	s.ExitCode = *tExitCode
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: ExitContractRet
// ----------------------------------------

// ExitContractRet type
type ExitContractRet VoidType

// NewExitContractRet factory
func NewExitContractRet() *ExitContractRet {
	o := ExitContractRet(*NewVoidType())
	return &o
}

// Serialize ExitContractRet
func (n ExitContractRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeExitContractRet function
func DeserializeExitContractRet(vb *VariableBlob) (uint64,*ExitContractRet,error) {
	var ot ExitContractRet
	return 0,&ot,nil}

// MarshalJSON ExitContractRet
func (n ExitContractRet) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ExitContractRet
func (n *ExitContractRet) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ExitContractRet(v)
	return nil
}


// ----------------------------------------
//  Typedef: GetHeadInfoArgs
// ----------------------------------------

// GetHeadInfoArgs type
type GetHeadInfoArgs VoidType

// NewGetHeadInfoArgs factory
func NewGetHeadInfoArgs() *GetHeadInfoArgs {
	o := GetHeadInfoArgs(*NewVoidType())
	return &o
}

// Serialize GetHeadInfoArgs
func (n GetHeadInfoArgs) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeGetHeadInfoArgs function
func DeserializeGetHeadInfoArgs(vb *VariableBlob) (uint64,*GetHeadInfoArgs,error) {
	var ot GetHeadInfoArgs
	return 0,&ot,nil}

// MarshalJSON GetHeadInfoArgs
func (n GetHeadInfoArgs) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetHeadInfoArgs
func (n *GetHeadInfoArgs) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetHeadInfoArgs(v)
	return nil
}


// ----------------------------------------
//  Typedef: GetHeadInfoRet
// ----------------------------------------

// GetHeadInfoRet type
type GetHeadInfoRet HeadInfo

// NewGetHeadInfoRet factory
func NewGetHeadInfoRet() *GetHeadInfoRet {
	o := GetHeadInfoRet(*NewHeadInfo())
	return &o
}

// Serialize GetHeadInfoRet
func (n GetHeadInfoRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := HeadInfo(n)
	return ox.Serialize(vb)
}

// DeserializeGetHeadInfoRet function
func DeserializeGetHeadInfoRet(vb *VariableBlob) (uint64,*GetHeadInfoRet,error) {
	var ot GetHeadInfoRet
	i,n,err := DeserializeHeadInfo(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = GetHeadInfoRet(*n)
	return i,&ot,nil}

// MarshalJSON GetHeadInfoRet
func (n GetHeadInfoRet) MarshalJSON() ([]byte, error) {
	v := HeadInfo(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetHeadInfoRet
func (n *GetHeadInfoRet) UnmarshalJSON(data []byte) error {
	v := HeadInfo(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetHeadInfoRet(v)
	return nil
}


// ----------------------------------------
//  Struct: HashArgs
// ----------------------------------------

// HashArgs type
type HashArgs struct {
    Code UInt64 `json:"code"`
    Obj VariableBlob `json:"obj"`
    Size UInt64 `json:"size"`
}

// NewHashArgs factory
func NewHashArgs() *HashArgs {
	o := HashArgs{}
	o.Code = *NewUInt64()
	o.Obj = *NewVariableBlob()
	o.Size = *NewUInt64()
	return &o
}

// Serialize HashArgs
func (n HashArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Code.Serialize(vb)
	vb = n.Obj.Serialize(vb)
	vb = n.Size.Serialize(vb)
	return vb
}

// DeserializeHashArgs function
func DeserializeHashArgs(vb *VariableBlob) (uint64,*HashArgs,error) {
	var i,j uint64 = 0,0
	s := HashArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tCode,err := DeserializeUInt64(&ovb); i+=j
	if err != nil {
		return 0, &HashArgs{}, err
	}
	s.Code = *tCode
	ovb = (*vb)[i:]
	j,tObj,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &HashArgs{}, err
	}
	s.Obj = *tObj
	ovb = (*vb)[i:]
	j,tSize,err := DeserializeUInt64(&ovb); i+=j
	if err != nil {
		return 0, &HashArgs{}, err
	}
	s.Size = *tSize
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: HashRet
// ----------------------------------------

// HashRet type
type HashRet Multihash

// NewHashRet factory
func NewHashRet() *HashRet {
	o := HashRet(*NewMultihash())
	return &o
}

// Serialize HashRet
func (n HashRet) Serialize(vb *VariableBlob) *VariableBlob {
	ox := Multihash(n)
	return ox.Serialize(vb)
}

// DeserializeHashRet function
func DeserializeHashRet(vb *VariableBlob) (uint64,*HashRet,error) {
	var ot HashRet
	i,n,err := DeserializeMultihash(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = HashRet(*n)
	return i,&ot,nil}

// MarshalJSON HashRet
func (n HashRet) MarshalJSON() ([]byte, error) {
	v := Multihash(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *HashRet
func (n *HashRet) UnmarshalJSON(data []byte) error {
	v := Multihash(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = HashRet(v)
	return nil
}


// ----------------------------------------
//  Typedef: SignatureType
// ----------------------------------------

// SignatureType type
type SignatureType FixedBlob65

// NewSignatureType factory
func NewSignatureType() *SignatureType {
	o := SignatureType(*NewFixedBlob65())
	return &o
}

// Serialize SignatureType
func (n SignatureType) Serialize(vb *VariableBlob) *VariableBlob {
	ox := FixedBlob65(n)
	return ox.Serialize(vb)
}

// DeserializeSignatureType function
func DeserializeSignatureType(vb *VariableBlob) (uint64,*SignatureType,error) {
	var ot SignatureType
	i,n,err := DeserializeFixedBlob65(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = SignatureType(*n)
	return i,&ot,nil}

// MarshalJSON SignatureType
func (n SignatureType) MarshalJSON() ([]byte, error) {
	v := FixedBlob65(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *SignatureType
func (n *SignatureType) UnmarshalJSON(data []byte) error {
	v := FixedBlob65(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = SignatureType(v)
	return nil
}



// ----------------------------------------
//  FixedBlob20
// ----------------------------------------

// FixedBlob20 type
type FixedBlob20 [20]byte

// NewFixedBlob20 factory
func NewFixedBlob20() *FixedBlob20 {
	var fb FixedBlob20
	return &fb
}

// Serialize FixedBlob20
func (n FixedBlob20) Serialize(vb *VariableBlob) *VariableBlob {
	ovb := append(*vb, n[:]...)
	return &ovb
}

// DeserializeFixedBlob20 function
func DeserializeFixedBlob20(vb *VariableBlob) (uint64, *FixedBlob20,error) {
	var result FixedBlob20
	if len(*vb) < 20 {
		return 0,&result,errors.New("unexpected eof")
	}
	copy(result[:], *vb)
	return 20,&result,nil
}

// MarshalJSON FixedBlob20
func (n FixedBlob20) MarshalJSON() ([]byte, error) {
	nfb := NewVariableBlob()
	nfb = n.Serialize(nfb)
	s := EncodeBytes(*nfb)
	return json.Marshal(s)
}

// UnmarshalJSON FixedBlob20
func (n *FixedBlob20) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	db,err := DecodeBytes(s)
	if err != nil {
		return err
	}

	if len(db) != 20 {
		return fmt.Errorf("deserialized blob length %d does not match size %d", len(db), 20)
	}

	copy((*n)[:], db)

	return nil
}

// ----------------------------------------
//  FixedBlob65
// ----------------------------------------

// FixedBlob65 type
type FixedBlob65 [65]byte

// NewFixedBlob65 factory
func NewFixedBlob65() *FixedBlob65 {
	var fb FixedBlob65
	return &fb
}

// Serialize FixedBlob65
func (n FixedBlob65) Serialize(vb *VariableBlob) *VariableBlob {
	ovb := append(*vb, n[:]...)
	return &ovb
}

// DeserializeFixedBlob65 function
func DeserializeFixedBlob65(vb *VariableBlob) (uint64, *FixedBlob65,error) {
	var result FixedBlob65
	if len(*vb) < 65 {
		return 0,&result,errors.New("unexpected eof")
	}
	copy(result[:], *vb)
	return 65,&result,nil
}

// MarshalJSON FixedBlob65
func (n FixedBlob65) MarshalJSON() ([]byte, error) {
	nfb := NewVariableBlob()
	nfb = n.Serialize(nfb)
	s := EncodeBytes(*nfb)
	return json.Marshal(s)
}

// UnmarshalJSON FixedBlob65
func (n *FixedBlob65) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	db,err := DecodeBytes(s)
	if err != nil {
		return err
	}

	if len(db) != 65 {
		return fmt.Errorf("deserialized blob length %d does not match size %d", len(db), 65)
	}

	copy((*n)[:], db)

	return nil
}



// ----------------------------------------
//  Opaque Helpers
// ----------------------------------------

type opaqueJSON struct {
	Opaque struct {
		Type string `json:"type"`
		Value VariableBlob `json:"value"`
	}`json:"opaque"`
}

// ----------------------------------------
//  OpaqueActiveBlockData
// ----------------------------------------

// OpaqueActiveBlockData type
type OpaqueActiveBlockData struct {
	blob *VariableBlob
	native *ActiveBlockData
}

// NewOpaqueActiveBlockData factory
func NewOpaqueActiveBlockData() *OpaqueActiveBlockData {
	o := OpaqueActiveBlockData{}
	o.native = NewActiveBlockData()
	return &o
}

// NewOpaqueActiveBlockDataFromBlob factory
func NewOpaqueActiveBlockDataFromBlob(vb *VariableBlob) *OpaqueActiveBlockData {
	o := OpaqueActiveBlockData{}
	o.blob = vb
	return &o
}

// NewOpaqueActiveBlockDataFromNative factory
func NewOpaqueActiveBlockDataFromNative(n ActiveBlockData) *OpaqueActiveBlockData {
	o := OpaqueActiveBlockData{}
	o.native = &n
	return &o
}

// GetBlob *OpaqueActiveBlockData
func (n *OpaqueActiveBlockData) GetBlob() *VariableBlob {
	if !n.IsBoxed() {
		n.Box()
	}

	return n.blob
}

// GetNative *OpaqueActiveBlockData
func (n *OpaqueActiveBlockData) GetNative() (*ActiveBlockData,error) {
	if( n.native == nil ) {
		return nil,errors.New("opaque type not unboxed")
	}

	return n.native,nil;
}

// Box *OpaqueActiveBlockData
func (n *OpaqueActiveBlockData) Box() {
	if (n.native != nil) {
		n.serializeNative()
		n.native = nil
	}
}

// Unbox *OpaqueActiveBlockData
func (n *OpaqueActiveBlockData) Unbox() {
	if (n.native == nil && n.blob != nil) {
		var err error
		var b uint64
		b,n.native,err = DeserializeActiveBlockData(n.blob)

		if err != nil || b != uint64(len(*n.blob)) {
			n.native = nil
		}
	}
}

// IsBoxed *OpaqueActiveBlockData
func (n *OpaqueActiveBlockData) IsBoxed() bool {
	return n.native == nil;
}

func (n *OpaqueActiveBlockData) serializeNative() {
	vb := NewVariableBlob()
	n.blob = n.native.Serialize(vb)
}

// Serialize OpaqueActiveBlockData
func (n OpaqueActiveBlockData) Serialize(vb *VariableBlob) *VariableBlob {
	n.Box()
	return n.blob.Serialize(vb)
}

// DeserializeOpaqueActiveBlockData function
func DeserializeOpaqueActiveBlockData(vb *VariableBlob) (uint64,*OpaqueActiveBlockData,error) {
	size,nv,err := DeserializeVariableBlob(vb)
	var o OpaqueActiveBlockData
	if err != nil {
		return 0, &o, err
	}
	o = OpaqueActiveBlockData{blob:nv, native:nil}
	return size,&o,nil
}

// MarshalJSON OpaqueActiveBlockData
func (n OpaqueActiveBlockData) MarshalJSON() ([]byte, error) {
	n.Unbox()

	if !n.IsBoxed() {
		return json.Marshal(&n.native)
	}

	v := opaqueJSON{}
	v.Opaque.Type = "koinos::types::protocol::active_block_data"
	v.Opaque.Value = *n.blob

	return json.Marshal(&v)
}

// UnmarshalJSON *OpaqueActiveBlockData
func (n *OpaqueActiveBlockData) UnmarshalJSON(data []byte) (error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	_, isOpaque := j["opaque"]

	if isOpaque {
		var obj opaqueJSON

		if err := json.Unmarshal(data, &obj); err != nil {
			return err
		}
		if strings.Compare(obj.Opaque.Type, "koinos::types::protocol::active_block_data") != 0 {
			return errors.New("unexpected opaque type name")
		}
		n.blob = &obj.Opaque.Value
		n.native = nil
	} else {
		if err := json.Unmarshal(data, &n.native); err != nil {
			return err
		}
	}

	return nil
}

// ----------------------------------------
//  OpaqueActiveTransactionData
// ----------------------------------------

// OpaqueActiveTransactionData type
type OpaqueActiveTransactionData struct {
	blob *VariableBlob
	native *ActiveTransactionData
}

// NewOpaqueActiveTransactionData factory
func NewOpaqueActiveTransactionData() *OpaqueActiveTransactionData {
	o := OpaqueActiveTransactionData{}
	o.native = NewActiveTransactionData()
	return &o
}

// NewOpaqueActiveTransactionDataFromBlob factory
func NewOpaqueActiveTransactionDataFromBlob(vb *VariableBlob) *OpaqueActiveTransactionData {
	o := OpaqueActiveTransactionData{}
	o.blob = vb
	return &o
}

// NewOpaqueActiveTransactionDataFromNative factory
func NewOpaqueActiveTransactionDataFromNative(n ActiveTransactionData) *OpaqueActiveTransactionData {
	o := OpaqueActiveTransactionData{}
	o.native = &n
	return &o
}

// GetBlob *OpaqueActiveTransactionData
func (n *OpaqueActiveTransactionData) GetBlob() *VariableBlob {
	if !n.IsBoxed() {
		n.Box()
	}

	return n.blob
}

// GetNative *OpaqueActiveTransactionData
func (n *OpaqueActiveTransactionData) GetNative() (*ActiveTransactionData,error) {
	if( n.native == nil ) {
		return nil,errors.New("opaque type not unboxed")
	}

	return n.native,nil;
}

// Box *OpaqueActiveTransactionData
func (n *OpaqueActiveTransactionData) Box() {
	if (n.native != nil) {
		n.serializeNative()
		n.native = nil
	}
}

// Unbox *OpaqueActiveTransactionData
func (n *OpaqueActiveTransactionData) Unbox() {
	if (n.native == nil && n.blob != nil) {
		var err error
		var b uint64
		b,n.native,err = DeserializeActiveTransactionData(n.blob)

		if err != nil || b != uint64(len(*n.blob)) {
			n.native = nil
		}
	}
}

// IsBoxed *OpaqueActiveTransactionData
func (n *OpaqueActiveTransactionData) IsBoxed() bool {
	return n.native == nil;
}

func (n *OpaqueActiveTransactionData) serializeNative() {
	vb := NewVariableBlob()
	n.blob = n.native.Serialize(vb)
}

// Serialize OpaqueActiveTransactionData
func (n OpaqueActiveTransactionData) Serialize(vb *VariableBlob) *VariableBlob {
	n.Box()
	return n.blob.Serialize(vb)
}

// DeserializeOpaqueActiveTransactionData function
func DeserializeOpaqueActiveTransactionData(vb *VariableBlob) (uint64,*OpaqueActiveTransactionData,error) {
	size,nv,err := DeserializeVariableBlob(vb)
	var o OpaqueActiveTransactionData
	if err != nil {
		return 0, &o, err
	}
	o = OpaqueActiveTransactionData{blob:nv, native:nil}
	return size,&o,nil
}

// MarshalJSON OpaqueActiveTransactionData
func (n OpaqueActiveTransactionData) MarshalJSON() ([]byte, error) {
	n.Unbox()

	if !n.IsBoxed() {
		return json.Marshal(&n.native)
	}

	v := opaqueJSON{}
	v.Opaque.Type = "koinos::types::protocol::active_transaction_data"
	v.Opaque.Value = *n.blob

	return json.Marshal(&v)
}

// UnmarshalJSON *OpaqueActiveTransactionData
func (n *OpaqueActiveTransactionData) UnmarshalJSON(data []byte) (error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	_, isOpaque := j["opaque"]

	if isOpaque {
		var obj opaqueJSON

		if err := json.Unmarshal(data, &obj); err != nil {
			return err
		}
		if strings.Compare(obj.Opaque.Type, "koinos::types::protocol::active_transaction_data") != 0 {
			return errors.New("unexpected opaque type name")
		}
		n.blob = &obj.Opaque.Value
		n.native = nil
	} else {
		if err := json.Unmarshal(data, &n.native); err != nil {
			return err
		}
	}

	return nil
}

// ----------------------------------------
//  OpaqueBlock
// ----------------------------------------

// OpaqueBlock type
type OpaqueBlock struct {
	blob *VariableBlob
	native *Block
}

// NewOpaqueBlock factory
func NewOpaqueBlock() *OpaqueBlock {
	o := OpaqueBlock{}
	o.native = NewBlock()
	return &o
}

// NewOpaqueBlockFromBlob factory
func NewOpaqueBlockFromBlob(vb *VariableBlob) *OpaqueBlock {
	o := OpaqueBlock{}
	o.blob = vb
	return &o
}

// NewOpaqueBlockFromNative factory
func NewOpaqueBlockFromNative(n Block) *OpaqueBlock {
	o := OpaqueBlock{}
	o.native = &n
	return &o
}

// GetBlob *OpaqueBlock
func (n *OpaqueBlock) GetBlob() *VariableBlob {
	if !n.IsBoxed() {
		n.Box()
	}

	return n.blob
}

// GetNative *OpaqueBlock
func (n *OpaqueBlock) GetNative() (*Block,error) {
	if( n.native == nil ) {
		return nil,errors.New("opaque type not unboxed")
	}

	return n.native,nil;
}

// Box *OpaqueBlock
func (n *OpaqueBlock) Box() {
	if (n.native != nil) {
		n.serializeNative()
		n.native = nil
	}
}

// Unbox *OpaqueBlock
func (n *OpaqueBlock) Unbox() {
	if (n.native == nil && n.blob != nil) {
		var err error
		var b uint64
		b,n.native,err = DeserializeBlock(n.blob)

		if err != nil || b != uint64(len(*n.blob)) {
			n.native = nil
		}
	}
}

// IsBoxed *OpaqueBlock
func (n *OpaqueBlock) IsBoxed() bool {
	return n.native == nil;
}

func (n *OpaqueBlock) serializeNative() {
	vb := NewVariableBlob()
	n.blob = n.native.Serialize(vb)
}

// Serialize OpaqueBlock
func (n OpaqueBlock) Serialize(vb *VariableBlob) *VariableBlob {
	n.Box()
	return n.blob.Serialize(vb)
}

// DeserializeOpaqueBlock function
func DeserializeOpaqueBlock(vb *VariableBlob) (uint64,*OpaqueBlock,error) {
	size,nv,err := DeserializeVariableBlob(vb)
	var o OpaqueBlock
	if err != nil {
		return 0, &o, err
	}
	o = OpaqueBlock{blob:nv, native:nil}
	return size,&o,nil
}

// MarshalJSON OpaqueBlock
func (n OpaqueBlock) MarshalJSON() ([]byte, error) {
	n.Unbox()

	if !n.IsBoxed() {
		return json.Marshal(&n.native)
	}

	v := opaqueJSON{}
	v.Opaque.Type = "koinos::types::protocol::block"
	v.Opaque.Value = *n.blob

	return json.Marshal(&v)
}

// UnmarshalJSON *OpaqueBlock
func (n *OpaqueBlock) UnmarshalJSON(data []byte) (error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	_, isOpaque := j["opaque"]

	if isOpaque {
		var obj opaqueJSON

		if err := json.Unmarshal(data, &obj); err != nil {
			return err
		}
		if strings.Compare(obj.Opaque.Type, "koinos::types::protocol::block") != 0 {
			return errors.New("unexpected opaque type name")
		}
		n.blob = &obj.Opaque.Value
		n.native = nil
	} else {
		if err := json.Unmarshal(data, &n.native); err != nil {
			return err
		}
	}

	return nil
}

// ----------------------------------------
//  OpaqueBlockReceipt
// ----------------------------------------

// OpaqueBlockReceipt type
type OpaqueBlockReceipt struct {
	blob *VariableBlob
	native *BlockReceipt
}

// NewOpaqueBlockReceipt factory
func NewOpaqueBlockReceipt() *OpaqueBlockReceipt {
	o := OpaqueBlockReceipt{}
	o.native = NewBlockReceipt()
	return &o
}

// NewOpaqueBlockReceiptFromBlob factory
func NewOpaqueBlockReceiptFromBlob(vb *VariableBlob) *OpaqueBlockReceipt {
	o := OpaqueBlockReceipt{}
	o.blob = vb
	return &o
}

// NewOpaqueBlockReceiptFromNative factory
func NewOpaqueBlockReceiptFromNative(n BlockReceipt) *OpaqueBlockReceipt {
	o := OpaqueBlockReceipt{}
	o.native = &n
	return &o
}

// GetBlob *OpaqueBlockReceipt
func (n *OpaqueBlockReceipt) GetBlob() *VariableBlob {
	if !n.IsBoxed() {
		n.Box()
	}

	return n.blob
}

// GetNative *OpaqueBlockReceipt
func (n *OpaqueBlockReceipt) GetNative() (*BlockReceipt,error) {
	if( n.native == nil ) {
		return nil,errors.New("opaque type not unboxed")
	}

	return n.native,nil;
}

// Box *OpaqueBlockReceipt
func (n *OpaqueBlockReceipt) Box() {
	if (n.native != nil) {
		n.serializeNative()
		n.native = nil
	}
}

// Unbox *OpaqueBlockReceipt
func (n *OpaqueBlockReceipt) Unbox() {
	if (n.native == nil && n.blob != nil) {
		var err error
		var b uint64
		b,n.native,err = DeserializeBlockReceipt(n.blob)

		if err != nil || b != uint64(len(*n.blob)) {
			n.native = nil
		}
	}
}

// IsBoxed *OpaqueBlockReceipt
func (n *OpaqueBlockReceipt) IsBoxed() bool {
	return n.native == nil;
}

func (n *OpaqueBlockReceipt) serializeNative() {
	vb := NewVariableBlob()
	n.blob = n.native.Serialize(vb)
}

// Serialize OpaqueBlockReceipt
func (n OpaqueBlockReceipt) Serialize(vb *VariableBlob) *VariableBlob {
	n.Box()
	return n.blob.Serialize(vb)
}

// DeserializeOpaqueBlockReceipt function
func DeserializeOpaqueBlockReceipt(vb *VariableBlob) (uint64,*OpaqueBlockReceipt,error) {
	size,nv,err := DeserializeVariableBlob(vb)
	var o OpaqueBlockReceipt
	if err != nil {
		return 0, &o, err
	}
	o = OpaqueBlockReceipt{blob:nv, native:nil}
	return size,&o,nil
}

// MarshalJSON OpaqueBlockReceipt
func (n OpaqueBlockReceipt) MarshalJSON() ([]byte, error) {
	n.Unbox()

	if !n.IsBoxed() {
		return json.Marshal(&n.native)
	}

	v := opaqueJSON{}
	v.Opaque.Type = "koinos::types::protocol::block_receipt"
	v.Opaque.Value = *n.blob

	return json.Marshal(&v)
}

// UnmarshalJSON *OpaqueBlockReceipt
func (n *OpaqueBlockReceipt) UnmarshalJSON(data []byte) (error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	_, isOpaque := j["opaque"]

	if isOpaque {
		var obj opaqueJSON

		if err := json.Unmarshal(data, &obj); err != nil {
			return err
		}
		if strings.Compare(obj.Opaque.Type, "koinos::types::protocol::block_receipt") != 0 {
			return errors.New("unexpected opaque type name")
		}
		n.blob = &obj.Opaque.Value
		n.native = nil
	} else {
		if err := json.Unmarshal(data, &n.native); err != nil {
			return err
		}
	}

	return nil
}

// ----------------------------------------
//  OpaquePassiveBlockData
// ----------------------------------------

// OpaquePassiveBlockData type
type OpaquePassiveBlockData struct {
	blob *VariableBlob
	native *PassiveBlockData
}

// NewOpaquePassiveBlockData factory
func NewOpaquePassiveBlockData() *OpaquePassiveBlockData {
	o := OpaquePassiveBlockData{}
	o.native = NewPassiveBlockData()
	return &o
}

// NewOpaquePassiveBlockDataFromBlob factory
func NewOpaquePassiveBlockDataFromBlob(vb *VariableBlob) *OpaquePassiveBlockData {
	o := OpaquePassiveBlockData{}
	o.blob = vb
	return &o
}

// NewOpaquePassiveBlockDataFromNative factory
func NewOpaquePassiveBlockDataFromNative(n PassiveBlockData) *OpaquePassiveBlockData {
	o := OpaquePassiveBlockData{}
	o.native = &n
	return &o
}

// GetBlob *OpaquePassiveBlockData
func (n *OpaquePassiveBlockData) GetBlob() *VariableBlob {
	if !n.IsBoxed() {
		n.Box()
	}

	return n.blob
}

// GetNative *OpaquePassiveBlockData
func (n *OpaquePassiveBlockData) GetNative() (*PassiveBlockData,error) {
	if( n.native == nil ) {
		return nil,errors.New("opaque type not unboxed")
	}

	return n.native,nil;
}

// Box *OpaquePassiveBlockData
func (n *OpaquePassiveBlockData) Box() {
	if (n.native != nil) {
		n.serializeNative()
		n.native = nil
	}
}

// Unbox *OpaquePassiveBlockData
func (n *OpaquePassiveBlockData) Unbox() {
	if (n.native == nil && n.blob != nil) {
		var err error
		var b uint64
		b,n.native,err = DeserializePassiveBlockData(n.blob)

		if err != nil || b != uint64(len(*n.blob)) {
			n.native = nil
		}
	}
}

// IsBoxed *OpaquePassiveBlockData
func (n *OpaquePassiveBlockData) IsBoxed() bool {
	return n.native == nil;
}

func (n *OpaquePassiveBlockData) serializeNative() {
	vb := NewVariableBlob()
	n.blob = n.native.Serialize(vb)
}

// Serialize OpaquePassiveBlockData
func (n OpaquePassiveBlockData) Serialize(vb *VariableBlob) *VariableBlob {
	n.Box()
	return n.blob.Serialize(vb)
}

// DeserializeOpaquePassiveBlockData function
func DeserializeOpaquePassiveBlockData(vb *VariableBlob) (uint64,*OpaquePassiveBlockData,error) {
	size,nv,err := DeserializeVariableBlob(vb)
	var o OpaquePassiveBlockData
	if err != nil {
		return 0, &o, err
	}
	o = OpaquePassiveBlockData{blob:nv, native:nil}
	return size,&o,nil
}

// MarshalJSON OpaquePassiveBlockData
func (n OpaquePassiveBlockData) MarshalJSON() ([]byte, error) {
	n.Unbox()

	if !n.IsBoxed() {
		return json.Marshal(&n.native)
	}

	v := opaqueJSON{}
	v.Opaque.Type = "koinos::types::protocol::passive_block_data"
	v.Opaque.Value = *n.blob

	return json.Marshal(&v)
}

// UnmarshalJSON *OpaquePassiveBlockData
func (n *OpaquePassiveBlockData) UnmarshalJSON(data []byte) (error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	_, isOpaque := j["opaque"]

	if isOpaque {
		var obj opaqueJSON

		if err := json.Unmarshal(data, &obj); err != nil {
			return err
		}
		if strings.Compare(obj.Opaque.Type, "koinos::types::protocol::passive_block_data") != 0 {
			return errors.New("unexpected opaque type name")
		}
		n.blob = &obj.Opaque.Value
		n.native = nil
	} else {
		if err := json.Unmarshal(data, &n.native); err != nil {
			return err
		}
	}

	return nil
}

// ----------------------------------------
//  OpaquePassiveTransactionData
// ----------------------------------------

// OpaquePassiveTransactionData type
type OpaquePassiveTransactionData struct {
	blob *VariableBlob
	native *PassiveTransactionData
}

// NewOpaquePassiveTransactionData factory
func NewOpaquePassiveTransactionData() *OpaquePassiveTransactionData {
	o := OpaquePassiveTransactionData{}
	o.native = NewPassiveTransactionData()
	return &o
}

// NewOpaquePassiveTransactionDataFromBlob factory
func NewOpaquePassiveTransactionDataFromBlob(vb *VariableBlob) *OpaquePassiveTransactionData {
	o := OpaquePassiveTransactionData{}
	o.blob = vb
	return &o
}

// NewOpaquePassiveTransactionDataFromNative factory
func NewOpaquePassiveTransactionDataFromNative(n PassiveTransactionData) *OpaquePassiveTransactionData {
	o := OpaquePassiveTransactionData{}
	o.native = &n
	return &o
}

// GetBlob *OpaquePassiveTransactionData
func (n *OpaquePassiveTransactionData) GetBlob() *VariableBlob {
	if !n.IsBoxed() {
		n.Box()
	}

	return n.blob
}

// GetNative *OpaquePassiveTransactionData
func (n *OpaquePassiveTransactionData) GetNative() (*PassiveTransactionData,error) {
	if( n.native == nil ) {
		return nil,errors.New("opaque type not unboxed")
	}

	return n.native,nil;
}

// Box *OpaquePassiveTransactionData
func (n *OpaquePassiveTransactionData) Box() {
	if (n.native != nil) {
		n.serializeNative()
		n.native = nil
	}
}

// Unbox *OpaquePassiveTransactionData
func (n *OpaquePassiveTransactionData) Unbox() {
	if (n.native == nil && n.blob != nil) {
		var err error
		var b uint64
		b,n.native,err = DeserializePassiveTransactionData(n.blob)

		if err != nil || b != uint64(len(*n.blob)) {
			n.native = nil
		}
	}
}

// IsBoxed *OpaquePassiveTransactionData
func (n *OpaquePassiveTransactionData) IsBoxed() bool {
	return n.native == nil;
}

func (n *OpaquePassiveTransactionData) serializeNative() {
	vb := NewVariableBlob()
	n.blob = n.native.Serialize(vb)
}

// Serialize OpaquePassiveTransactionData
func (n OpaquePassiveTransactionData) Serialize(vb *VariableBlob) *VariableBlob {
	n.Box()
	return n.blob.Serialize(vb)
}

// DeserializeOpaquePassiveTransactionData function
func DeserializeOpaquePassiveTransactionData(vb *VariableBlob) (uint64,*OpaquePassiveTransactionData,error) {
	size,nv,err := DeserializeVariableBlob(vb)
	var o OpaquePassiveTransactionData
	if err != nil {
		return 0, &o, err
	}
	o = OpaquePassiveTransactionData{blob:nv, native:nil}
	return size,&o,nil
}

// MarshalJSON OpaquePassiveTransactionData
func (n OpaquePassiveTransactionData) MarshalJSON() ([]byte, error) {
	n.Unbox()

	if !n.IsBoxed() {
		return json.Marshal(&n.native)
	}

	v := opaqueJSON{}
	v.Opaque.Type = "koinos::types::protocol::passive_transaction_data"
	v.Opaque.Value = *n.blob

	return json.Marshal(&v)
}

// UnmarshalJSON *OpaquePassiveTransactionData
func (n *OpaquePassiveTransactionData) UnmarshalJSON(data []byte) (error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	_, isOpaque := j["opaque"]

	if isOpaque {
		var obj opaqueJSON

		if err := json.Unmarshal(data, &obj); err != nil {
			return err
		}
		if strings.Compare(obj.Opaque.Type, "koinos::types::protocol::passive_transaction_data") != 0 {
			return errors.New("unexpected opaque type name")
		}
		n.blob = &obj.Opaque.Value
		n.native = nil
	} else {
		if err := json.Unmarshal(data, &n.native); err != nil {
			return err
		}
	}

	return nil
}

// ----------------------------------------
//  OpaqueQueryItemResult
// ----------------------------------------

// OpaqueQueryItemResult type
type OpaqueQueryItemResult struct {
	blob *VariableBlob
	native *QueryItemResult
}

// NewOpaqueQueryItemResult factory
func NewOpaqueQueryItemResult() *OpaqueQueryItemResult {
	o := OpaqueQueryItemResult{}
	o.native = NewQueryItemResult()
	return &o
}

// NewOpaqueQueryItemResultFromBlob factory
func NewOpaqueQueryItemResultFromBlob(vb *VariableBlob) *OpaqueQueryItemResult {
	o := OpaqueQueryItemResult{}
	o.blob = vb
	return &o
}

// NewOpaqueQueryItemResultFromNative factory
func NewOpaqueQueryItemResultFromNative(n QueryItemResult) *OpaqueQueryItemResult {
	o := OpaqueQueryItemResult{}
	o.native = &n
	return &o
}

// GetBlob *OpaqueQueryItemResult
func (n *OpaqueQueryItemResult) GetBlob() *VariableBlob {
	if !n.IsBoxed() {
		n.Box()
	}

	return n.blob
}

// GetNative *OpaqueQueryItemResult
func (n *OpaqueQueryItemResult) GetNative() (*QueryItemResult,error) {
	if( n.native == nil ) {
		return nil,errors.New("opaque type not unboxed")
	}

	return n.native,nil;
}

// Box *OpaqueQueryItemResult
func (n *OpaqueQueryItemResult) Box() {
	if (n.native != nil) {
		n.serializeNative()
		n.native = nil
	}
}

// Unbox *OpaqueQueryItemResult
func (n *OpaqueQueryItemResult) Unbox() {
	if (n.native == nil && n.blob != nil) {
		var err error
		var b uint64
		b,n.native,err = DeserializeQueryItemResult(n.blob)

		if err != nil || b != uint64(len(*n.blob)) {
			n.native = nil
		}
	}
}

// IsBoxed *OpaqueQueryItemResult
func (n *OpaqueQueryItemResult) IsBoxed() bool {
	return n.native == nil;
}

func (n *OpaqueQueryItemResult) serializeNative() {
	vb := NewVariableBlob()
	n.blob = n.native.Serialize(vb)
}

// Serialize OpaqueQueryItemResult
func (n OpaqueQueryItemResult) Serialize(vb *VariableBlob) *VariableBlob {
	n.Box()
	return n.blob.Serialize(vb)
}

// DeserializeOpaqueQueryItemResult function
func DeserializeOpaqueQueryItemResult(vb *VariableBlob) (uint64,*OpaqueQueryItemResult,error) {
	size,nv,err := DeserializeVariableBlob(vb)
	var o OpaqueQueryItemResult
	if err != nil {
		return 0, &o, err
	}
	o = OpaqueQueryItemResult{blob:nv, native:nil}
	return size,&o,nil
}

// MarshalJSON OpaqueQueryItemResult
func (n OpaqueQueryItemResult) MarshalJSON() ([]byte, error) {
	n.Unbox()

	if !n.IsBoxed() {
		return json.Marshal(&n.native)
	}

	v := opaqueJSON{}
	v.Opaque.Type = "koinos::types::rpc::query_item_result"
	v.Opaque.Value = *n.blob

	return json.Marshal(&v)
}

// UnmarshalJSON *OpaqueQueryItemResult
func (n *OpaqueQueryItemResult) UnmarshalJSON(data []byte) (error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	_, isOpaque := j["opaque"]

	if isOpaque {
		var obj opaqueJSON

		if err := json.Unmarshal(data, &obj); err != nil {
			return err
		}
		if strings.Compare(obj.Opaque.Type, "koinos::types::rpc::query_item_result") != 0 {
			return errors.New("unexpected opaque type name")
		}
		n.blob = &obj.Opaque.Value
		n.native = nil
	} else {
		if err := json.Unmarshal(data, &n.native); err != nil {
			return err
		}
	}

	return nil
}

// ----------------------------------------
//  OpaqueQueryParamItem
// ----------------------------------------

// OpaqueQueryParamItem type
type OpaqueQueryParamItem struct {
	blob *VariableBlob
	native *QueryParamItem
}

// NewOpaqueQueryParamItem factory
func NewOpaqueQueryParamItem() *OpaqueQueryParamItem {
	o := OpaqueQueryParamItem{}
	o.native = NewQueryParamItem()
	return &o
}

// NewOpaqueQueryParamItemFromBlob factory
func NewOpaqueQueryParamItemFromBlob(vb *VariableBlob) *OpaqueQueryParamItem {
	o := OpaqueQueryParamItem{}
	o.blob = vb
	return &o
}

// NewOpaqueQueryParamItemFromNative factory
func NewOpaqueQueryParamItemFromNative(n QueryParamItem) *OpaqueQueryParamItem {
	o := OpaqueQueryParamItem{}
	o.native = &n
	return &o
}

// GetBlob *OpaqueQueryParamItem
func (n *OpaqueQueryParamItem) GetBlob() *VariableBlob {
	if !n.IsBoxed() {
		n.Box()
	}

	return n.blob
}

// GetNative *OpaqueQueryParamItem
func (n *OpaqueQueryParamItem) GetNative() (*QueryParamItem,error) {
	if( n.native == nil ) {
		return nil,errors.New("opaque type not unboxed")
	}

	return n.native,nil;
}

// Box *OpaqueQueryParamItem
func (n *OpaqueQueryParamItem) Box() {
	if (n.native != nil) {
		n.serializeNative()
		n.native = nil
	}
}

// Unbox *OpaqueQueryParamItem
func (n *OpaqueQueryParamItem) Unbox() {
	if (n.native == nil && n.blob != nil) {
		var err error
		var b uint64
		b,n.native,err = DeserializeQueryParamItem(n.blob)

		if err != nil || b != uint64(len(*n.blob)) {
			n.native = nil
		}
	}
}

// IsBoxed *OpaqueQueryParamItem
func (n *OpaqueQueryParamItem) IsBoxed() bool {
	return n.native == nil;
}

func (n *OpaqueQueryParamItem) serializeNative() {
	vb := NewVariableBlob()
	n.blob = n.native.Serialize(vb)
}

// Serialize OpaqueQueryParamItem
func (n OpaqueQueryParamItem) Serialize(vb *VariableBlob) *VariableBlob {
	n.Box()
	return n.blob.Serialize(vb)
}

// DeserializeOpaqueQueryParamItem function
func DeserializeOpaqueQueryParamItem(vb *VariableBlob) (uint64,*OpaqueQueryParamItem,error) {
	size,nv,err := DeserializeVariableBlob(vb)
	var o OpaqueQueryParamItem
	if err != nil {
		return 0, &o, err
	}
	o = OpaqueQueryParamItem{blob:nv, native:nil}
	return size,&o,nil
}

// MarshalJSON OpaqueQueryParamItem
func (n OpaqueQueryParamItem) MarshalJSON() ([]byte, error) {
	n.Unbox()

	if !n.IsBoxed() {
		return json.Marshal(&n.native)
	}

	v := opaqueJSON{}
	v.Opaque.Type = "koinos::types::rpc::query_param_item"
	v.Opaque.Value = *n.blob

	return json.Marshal(&v)
}

// UnmarshalJSON *OpaqueQueryParamItem
func (n *OpaqueQueryParamItem) UnmarshalJSON(data []byte) (error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	_, isOpaque := j["opaque"]

	if isOpaque {
		var obj opaqueJSON

		if err := json.Unmarshal(data, &obj); err != nil {
			return err
		}
		if strings.Compare(obj.Opaque.Type, "koinos::types::rpc::query_param_item") != 0 {
			return errors.New("unexpected opaque type name")
		}
		n.blob = &obj.Opaque.Value
		n.native = nil
	} else {
		if err := json.Unmarshal(data, &n.native); err != nil {
			return err
		}
	}

	return nil
}

// ----------------------------------------
//  OpaqueTransaction
// ----------------------------------------

// OpaqueTransaction type
type OpaqueTransaction struct {
	blob *VariableBlob
	native *Transaction
}

// NewOpaqueTransaction factory
func NewOpaqueTransaction() *OpaqueTransaction {
	o := OpaqueTransaction{}
	o.native = NewTransaction()
	return &o
}

// NewOpaqueTransactionFromBlob factory
func NewOpaqueTransactionFromBlob(vb *VariableBlob) *OpaqueTransaction {
	o := OpaqueTransaction{}
	o.blob = vb
	return &o
}

// NewOpaqueTransactionFromNative factory
func NewOpaqueTransactionFromNative(n Transaction) *OpaqueTransaction {
	o := OpaqueTransaction{}
	o.native = &n
	return &o
}

// GetBlob *OpaqueTransaction
func (n *OpaqueTransaction) GetBlob() *VariableBlob {
	if !n.IsBoxed() {
		n.Box()
	}

	return n.blob
}

// GetNative *OpaqueTransaction
func (n *OpaqueTransaction) GetNative() (*Transaction,error) {
	if( n.native == nil ) {
		return nil,errors.New("opaque type not unboxed")
	}

	return n.native,nil;
}

// Box *OpaqueTransaction
func (n *OpaqueTransaction) Box() {
	if (n.native != nil) {
		n.serializeNative()
		n.native = nil
	}
}

// Unbox *OpaqueTransaction
func (n *OpaqueTransaction) Unbox() {
	if (n.native == nil && n.blob != nil) {
		var err error
		var b uint64
		b,n.native,err = DeserializeTransaction(n.blob)

		if err != nil || b != uint64(len(*n.blob)) {
			n.native = nil
		}
	}
}

// IsBoxed *OpaqueTransaction
func (n *OpaqueTransaction) IsBoxed() bool {
	return n.native == nil;
}

func (n *OpaqueTransaction) serializeNative() {
	vb := NewVariableBlob()
	n.blob = n.native.Serialize(vb)
}

// Serialize OpaqueTransaction
func (n OpaqueTransaction) Serialize(vb *VariableBlob) *VariableBlob {
	n.Box()
	return n.blob.Serialize(vb)
}

// DeserializeOpaqueTransaction function
func DeserializeOpaqueTransaction(vb *VariableBlob) (uint64,*OpaqueTransaction,error) {
	size,nv,err := DeserializeVariableBlob(vb)
	var o OpaqueTransaction
	if err != nil {
		return 0, &o, err
	}
	o = OpaqueTransaction{blob:nv, native:nil}
	return size,&o,nil
}

// MarshalJSON OpaqueTransaction
func (n OpaqueTransaction) MarshalJSON() ([]byte, error) {
	n.Unbox()

	if !n.IsBoxed() {
		return json.Marshal(&n.native)
	}

	v := opaqueJSON{}
	v.Opaque.Type = "koinos::types::protocol::transaction"
	v.Opaque.Value = *n.blob

	return json.Marshal(&v)
}

// UnmarshalJSON *OpaqueTransaction
func (n *OpaqueTransaction) UnmarshalJSON(data []byte) (error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	_, isOpaque := j["opaque"]

	if isOpaque {
		var obj opaqueJSON

		if err := json.Unmarshal(data, &obj); err != nil {
			return err
		}
		if strings.Compare(obj.Opaque.Type, "koinos::types::protocol::transaction") != 0 {
			return errors.New("unexpected opaque type name")
		}
		n.blob = &obj.Opaque.Value
		n.native = nil
	} else {
		if err := json.Unmarshal(data, &n.native); err != nil {
			return err
		}
	}

	return nil
}


// ----------------------------------------
//  VectorBlockItem
// ----------------------------------------

// VectorBlockItem type
type VectorBlockItem []BlockItem

// NewVectorBlockItem factory
func NewVectorBlockItem() *VectorBlockItem {
	o := VectorBlockItem(make([]BlockItem, 0))
	return &o
}

// Serialize VectorBlockItem
func (n VectorBlockItem) Serialize(vb *VariableBlob) *VariableBlob {
	header := make([]byte, binary.MaxVarintLen64)
	bytes := binary.PutUvarint(header, uint64(len(n)))
	ovb := append(*vb, header[:bytes]...)
	vb = &ovb
	for _, item := range n {
		vb = item.Serialize(vb)
	}

	return vb
}
// DeserializeVectorBlockItem function
func DeserializeVectorBlockItem(vb *VariableBlob) (uint64,*VectorBlockItem,error) {
	var result VectorBlockItem
	size,bytes := binary.Uvarint(*vb)
	if bytes <= 0 {
		return 0, &result, errors.New("could not deserialize multihash id")
	}
	result = VectorBlockItem(make([]BlockItem, 0, size))
	i := uint64(bytes)
	var j uint64
	var item *BlockItem
	var err error
	for num := uint64(0); num < size; num++ {
		ovb := (*vb)[i:]
		j,item,err = DeserializeBlockItem(&ovb)
		if nil != err {
			var v VectorBlockItem
			return 0,&v,err
		}
		i += j
		result = append(result, *item)
	}

	return i, &result, nil
}

// ----------------------------------------
//  VectorMultihash
// ----------------------------------------

// VectorMultihash type
type VectorMultihash []Multihash

// NewVectorMultihash factory
func NewVectorMultihash() *VectorMultihash {
	o := VectorMultihash(make([]Multihash, 0))
	return &o
}

// Serialize VectorMultihash
func (n VectorMultihash) Serialize(vb *VariableBlob) *VariableBlob {
	header := make([]byte, binary.MaxVarintLen64)
	bytes := binary.PutUvarint(header, uint64(len(n)))
	ovb := append(*vb, header[:bytes]...)
	vb = &ovb
	for _, item := range n {
		vb = item.Serialize(vb)
	}

	return vb
}
// DeserializeVectorMultihash function
func DeserializeVectorMultihash(vb *VariableBlob) (uint64,*VectorMultihash,error) {
	var result VectorMultihash
	size,bytes := binary.Uvarint(*vb)
	if bytes <= 0 {
		return 0, &result, errors.New("could not deserialize multihash id")
	}
	result = VectorMultihash(make([]Multihash, 0, size))
	i := uint64(bytes)
	var j uint64
	var item *Multihash
	var err error
	for num := uint64(0); num < size; num++ {
		ovb := (*vb)[i:]
		j,item,err = DeserializeMultihash(&ovb)
		if nil != err {
			var v VectorMultihash
			return 0,&v,err
		}
		i += j
		result = append(result, *item)
	}

	return i, &result, nil
}

// ----------------------------------------
//  VectorOpaqueTransaction
// ----------------------------------------

// VectorOpaqueTransaction type
type VectorOpaqueTransaction []OpaqueTransaction

// NewVectorOpaqueTransaction factory
func NewVectorOpaqueTransaction() *VectorOpaqueTransaction {
	o := VectorOpaqueTransaction(make([]OpaqueTransaction, 0))
	return &o
}

// Serialize VectorOpaqueTransaction
func (n VectorOpaqueTransaction) Serialize(vb *VariableBlob) *VariableBlob {
	header := make([]byte, binary.MaxVarintLen64)
	bytes := binary.PutUvarint(header, uint64(len(n)))
	ovb := append(*vb, header[:bytes]...)
	vb = &ovb
	for _, item := range n {
		vb = item.Serialize(vb)
	}

	return vb
}
// DeserializeVectorOpaqueTransaction function
func DeserializeVectorOpaqueTransaction(vb *VariableBlob) (uint64,*VectorOpaqueTransaction,error) {
	var result VectorOpaqueTransaction
	size,bytes := binary.Uvarint(*vb)
	if bytes <= 0 {
		return 0, &result, errors.New("could not deserialize multihash id")
	}
	result = VectorOpaqueTransaction(make([]OpaqueTransaction, 0, size))
	i := uint64(bytes)
	var j uint64
	var item *OpaqueTransaction
	var err error
	for num := uint64(0); num < size; num++ {
		ovb := (*vb)[i:]
		j,item,err = DeserializeOpaqueTransaction(&ovb)
		if nil != err {
			var v VectorOpaqueTransaction
			return 0,&v,err
		}
		i += j
		result = append(result, *item)
	}

	return i, &result, nil
}

// ----------------------------------------
//  VectorOperation
// ----------------------------------------

// VectorOperation type
type VectorOperation []Operation

// NewVectorOperation factory
func NewVectorOperation() *VectorOperation {
	o := VectorOperation(make([]Operation, 0))
	return &o
}

// Serialize VectorOperation
func (n VectorOperation) Serialize(vb *VariableBlob) *VariableBlob {
	header := make([]byte, binary.MaxVarintLen64)
	bytes := binary.PutUvarint(header, uint64(len(n)))
	ovb := append(*vb, header[:bytes]...)
	vb = &ovb
	for _, item := range n {
		vb = item.Serialize(vb)
	}

	return vb
}
// DeserializeVectorOperation function
func DeserializeVectorOperation(vb *VariableBlob) (uint64,*VectorOperation,error) {
	var result VectorOperation
	size,bytes := binary.Uvarint(*vb)
	if bytes <= 0 {
		return 0, &result, errors.New("could not deserialize multihash id")
	}
	result = VectorOperation(make([]Operation, 0, size))
	i := uint64(bytes)
	var j uint64
	var item *Operation
	var err error
	for num := uint64(0); num < size; num++ {
		ovb := (*vb)[i:]
		j,item,err = DeserializeOperation(&ovb)
		if nil != err {
			var v VectorOperation
			return 0,&v,err
		}
		i += j
		result = append(result, *item)
	}

	return i, &result, nil
}

// ----------------------------------------
//  VectorTransactionItem
// ----------------------------------------

// VectorTransactionItem type
type VectorTransactionItem []TransactionItem

// NewVectorTransactionItem factory
func NewVectorTransactionItem() *VectorTransactionItem {
	o := VectorTransactionItem(make([]TransactionItem, 0))
	return &o
}

// Serialize VectorTransactionItem
func (n VectorTransactionItem) Serialize(vb *VariableBlob) *VariableBlob {
	header := make([]byte, binary.MaxVarintLen64)
	bytes := binary.PutUvarint(header, uint64(len(n)))
	ovb := append(*vb, header[:bytes]...)
	vb = &ovb
	for _, item := range n {
		vb = item.Serialize(vb)
	}

	return vb
}
// DeserializeVectorTransactionItem function
func DeserializeVectorTransactionItem(vb *VariableBlob) (uint64,*VectorTransactionItem,error) {
	var result VectorTransactionItem
	size,bytes := binary.Uvarint(*vb)
	if bytes <= 0 {
		return 0, &result, errors.New("could not deserialize multihash id")
	}
	result = VectorTransactionItem(make([]TransactionItem, 0, size))
	i := uint64(bytes)
	var j uint64
	var item *TransactionItem
	var err error
	for num := uint64(0); num < size; num++ {
		ovb := (*vb)[i:]
		j,item,err = DeserializeTransactionItem(&ovb)
		if nil != err {
			var v VectorTransactionItem
			return 0,&v,err
		}
		i += j
		result = append(result, *item)
	}

	return i, &result, nil
}


