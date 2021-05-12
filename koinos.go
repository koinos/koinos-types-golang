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
//  Struct: BlockHeader
// ----------------------------------------

// BlockHeader type
type BlockHeader struct {
    Previous Multihash `json:"previous"`
    Height BlockHeightType `json:"height"`
    Timestamp TimestampType `json:"timestamp"`
}

// NewBlockHeader factory
func NewBlockHeader() *BlockHeader {
	o := BlockHeader{}
	o.Previous = *NewMultihash()
	o.Height = *NewBlockHeightType()
	o.Timestamp = *NewTimestampType()
	return &o
}

// Serialize BlockHeader
func (n BlockHeader) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Previous.Serialize(vb)
	vb = n.Height.Serialize(vb)
	vb = n.Timestamp.Serialize(vb)
	return vb
}

// DeserializeBlockHeader function
func DeserializeBlockHeader(vb *VariableBlob) (uint64,*BlockHeader,error) {
	var i,j uint64 = 0,0
	s := BlockHeader{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tPrevious,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &BlockHeader{}, err
	}
	s.Previous = *tPrevious
	ovb = (*vb)[i:]
	j,tHeight,err := DeserializeBlockHeightType(&ovb); i+=j
	if err != nil {
		return 0, &BlockHeader{}, err
	}
	s.Height = *tHeight
	ovb = (*vb)[i:]
	j,tTimestamp,err := DeserializeTimestampType(&ovb); i+=j
	if err != nil {
		return 0, &BlockHeader{}, err
	}
	s.Timestamp = *tTimestamp
	return i, &s, nil
}

// ----------------------------------------
//  Struct: ActiveBlockData
// ----------------------------------------

// ActiveBlockData type
type ActiveBlockData struct {
    TransactionMerkleRoot Multihash `json:"transaction_merkle_root"`
    PassiveDataMerkleRoot Multihash `json:"passive_data_merkle_root"`
    SignerAddress Multihash `json:"signer_address"`
}

// NewActiveBlockData factory
func NewActiveBlockData() *ActiveBlockData {
	o := ActiveBlockData{}
	o.TransactionMerkleRoot = *NewMultihash()
	o.PassiveDataMerkleRoot = *NewMultihash()
	o.SignerAddress = *NewMultihash()
	return &o
}

// Serialize ActiveBlockData
func (n ActiveBlockData) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.TransactionMerkleRoot.Serialize(vb)
	vb = n.PassiveDataMerkleRoot.Serialize(vb)
	vb = n.SignerAddress.Serialize(vb)
	return vb
}

// DeserializeActiveBlockData function
func DeserializeActiveBlockData(vb *VariableBlob) (uint64,*ActiveBlockData,error) {
	var i,j uint64 = 0,0
	s := ActiveBlockData{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransactionMerkleRoot,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &ActiveBlockData{}, err
	}
	s.TransactionMerkleRoot = *tTransactionMerkleRoot
	ovb = (*vb)[i:]
	j,tPassiveDataMerkleRoot,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &ActiveBlockData{}, err
	}
	s.PassiveDataMerkleRoot = *tPassiveDataMerkleRoot
	ovb = (*vb)[i:]
	j,tSignerAddress,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &ActiveBlockData{}, err
	}
	s.SignerAddress = *tSignerAddress
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
//  Struct: CallContractOperation
// ----------------------------------------

// CallContractOperation type
type CallContractOperation struct {
    ContractID ContractIDType `json:"contract_id"`
    EntryPoint UInt32 `json:"entry_point"`
    Args VariableBlob `json:"args"`
    Extensions UnusedExtensionsType `json:"extensions"`
}

// NewCallContractOperation factory
func NewCallContractOperation() *CallContractOperation {
	o := CallContractOperation{}
	o.ContractID = *NewContractIDType()
	o.EntryPoint = *NewUInt32()
	o.Args = *NewVariableBlob()
	o.Extensions = *NewUnusedExtensionsType()
	return &o
}

// Serialize CallContractOperation
func (n CallContractOperation) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ContractID.Serialize(vb)
	vb = n.EntryPoint.Serialize(vb)
	vb = n.Args.Serialize(vb)
	vb = n.Extensions.Serialize(vb)
	return vb
}

// DeserializeCallContractOperation function
func DeserializeCallContractOperation(vb *VariableBlob) (uint64,*CallContractOperation,error) {
	var i,j uint64 = 0,0
	s := CallContractOperation{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tContractID,err := DeserializeContractIDType(&ovb); i+=j
	if err != nil {
		return 0, &CallContractOperation{}, err
	}
	s.ContractID = *tContractID
	ovb = (*vb)[i:]
	j,tEntryPoint,err := DeserializeUInt32(&ovb); i+=j
	if err != nil {
		return 0, &CallContractOperation{}, err
	}
	s.EntryPoint = *tEntryPoint
	ovb = (*vb)[i:]
	j,tArgs,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &CallContractOperation{}, err
	}
	s.Args = *tArgs
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

// {'name': 'thunk_id', 'entries': [{'name': 'prints', 'value': 2406348109, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'verify_block_header', 'value': 2369936044, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_block', 'value': 2372743592, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_transaction', 'value': 2306978015, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_reserved_operation', 'value': 2335970550, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_upload_contract_operation', 'value': 2290263390, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_execute_contract_operation', 'value': 2246607595, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_set_system_call_operation', 'value': 2264476812, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_put_object', 'value': 2181271013, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_get_object', 'value': 2288165080, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_get_next_object', 'value': 2263109703, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_get_prev_object', 'value': 2371348733, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'execute_contract', 'value': 2319711875, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_entry_point', 'value': 2155788133, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_contract_args_size', 'value': 2201456262, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_contract_args', 'value': 2383977862, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'set_contract_return', 'value': 2260230773, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'exit_contract', 'value': 2180390815, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_head_info', 'value': 2313106628, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'hash', 'value': 2326459719, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'verify_block_signature', 'value': 2635873417, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'verify_merkle_root', 'value': 2396642763, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_transaction_payer', 'value': 2645774470, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_max_account_resources', 'value': 2431733645, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_transaction_resource_limit', 'value': 2346680737, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_last_irreversible_block', 'value': 2160310419, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_caller', 'value': 2184258817, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'require_authority', 'value': 2315678077, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_transaction_signature', 'value': 2202278691, 'doc': '', 'info': {'type': 'EnumEntry'}}], 'tref': {'name': ['koinos', 'uint32'], 'targs': None, 'info': {'type': 'Typeref'}}, 'doc': '', 'info': {'type': 'EnumClass'}}

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
	ThunkIDGetEntryPoint ThunkID = 2155788133
	ThunkIDGetContractArgsSize ThunkID = 2201456262
	ThunkIDGetContractArgs ThunkID = 2383977862
	ThunkIDSetContractReturn ThunkID = 2260230773
	ThunkIDExitContract ThunkID = 2180390815
	ThunkIDGetHeadInfo ThunkID = 2313106628
	ThunkIDHash ThunkID = 2326459719
	ThunkIDVerifyBlockSignature ThunkID = 2635873417
	ThunkIDVerifyMerkleRoot ThunkID = 2396642763
	ThunkIDGetTransactionPayer ThunkID = 2645774470
	ThunkIDGetMaxAccountResources ThunkID = 2431733645
	ThunkIDGetTransactionResourceLimit ThunkID = 2346680737
	ThunkIDGetLastIrreversibleBlock ThunkID = 2160310419
	ThunkIDGetCaller ThunkID = 2184258817
	ThunkIDRequireAuthority ThunkID = 2315678077
	ThunkIDGetTransactionSignature ThunkID = 2202278691
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
		case ThunkIDGetEntryPoint:
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
		case ThunkIDVerifyBlockSignature:
			return true
		case ThunkIDVerifyMerkleRoot:
			return true
		case ThunkIDGetTransactionPayer:
			return true
		case ThunkIDGetMaxAccountResources:
			return true
		case ThunkIDGetTransactionResourceLimit:
			return true
		case ThunkIDGetLastIrreversibleBlock:
			return true
		case ThunkIDGetCaller:
			return true
		case ThunkIDRequireAuthority:
			return true
		case ThunkIDGetTransactionSignature:
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
			return "koinos::chain::system_call_target_reserved"
		case *ThunkID:
			return "koinos::chain::thunk_id"
		case *ContractCallBundle:
			return "koinos::chain::contract_call_bundle"
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
		case "koinos::chain::system_call_target_reserved":
			v := NewSystemCallTargetReserved()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::chain::thunk_id":
			v := NewThunkID()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::chain::contract_call_bundle":
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
		case *CallContractOperation:
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
			return "koinos::protocol::reserved_operation"
		case *NopOperation:
			return "koinos::protocol::nop_operation"
		case *CreateSystemContractOperation:
			return "koinos::protocol::create_system_contract_operation"
		case *CallContractOperation:
			return "koinos::protocol::call_contract_operation"
		case *SetSystemCallOperation:
			return "koinos::protocol::set_system_call_operation"
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
			k,x,err := DeserializeCallContractOperation(&ovb)
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
		case "koinos::protocol::reserved_operation":
			v := NewReservedOperation()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::protocol::nop_operation":
			v := NewNopOperation()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::protocol::create_system_contract_operation":
			v := NewCreateSystemContractOperation()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::protocol::call_contract_operation":
			v := NewCallContractOperation()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::protocol::set_system_call_operation":
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
    ResourceLimit UInt128 `json:"resource_limit"`
    Nonce UInt64 `json:"nonce"`
    Operations VectorOperation `json:"operations"`
}

// NewActiveTransactionData factory
func NewActiveTransactionData() *ActiveTransactionData {
	o := ActiveTransactionData{}
	o.ResourceLimit = *NewUInt128()
	o.Nonce = *NewUInt64()
	o.Operations = *NewVectorOperation()
	return &o
}

// Serialize ActiveTransactionData
func (n ActiveTransactionData) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ResourceLimit.Serialize(vb)
	vb = n.Nonce.Serialize(vb)
	vb = n.Operations.Serialize(vb)
	return vb
}

// DeserializeActiveTransactionData function
func DeserializeActiveTransactionData(vb *VariableBlob) (uint64,*ActiveTransactionData,error) {
	var i,j uint64 = 0,0
	s := ActiveTransactionData{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tResourceLimit,err := DeserializeUInt128(&ovb); i+=j
	if err != nil {
		return 0, &ActiveTransactionData{}, err
	}
	s.ResourceLimit = *tResourceLimit
	ovb = (*vb)[i:]
	j,tNonce,err := DeserializeUInt64(&ovb); i+=j
	if err != nil {
		return 0, &ActiveTransactionData{}, err
	}
	s.Nonce = *tNonce
	ovb = (*vb)[i:]
	j,tOperations,err := DeserializeVectorOperation(&ovb); i+=j
	if err != nil {
		return 0, &ActiveTransactionData{}, err
	}
	s.Operations = *tOperations
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
    ID Multihash `json:"id"`
    ActiveData OpaqueActiveTransactionData `json:"active_data"`
    PassiveData OpaquePassiveTransactionData `json:"passive_data"`
    SignatureData VariableBlob `json:"signature_data"`
}

// NewTransaction factory
func NewTransaction() *Transaction {
	o := Transaction{}
	o.ID = *NewMultihash()
	o.ActiveData = *NewOpaqueActiveTransactionData()
	o.PassiveData = *NewOpaquePassiveTransactionData()
	o.SignatureData = *NewVariableBlob()
	return &o
}

// Serialize Transaction
func (n Transaction) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ID.Serialize(vb)
	vb = n.ActiveData.Serialize(vb)
	vb = n.PassiveData.Serialize(vb)
	vb = n.SignatureData.Serialize(vb)
	return vb
}

// DeserializeTransaction function
func DeserializeTransaction(vb *VariableBlob) (uint64,*Transaction,error) {
	var i,j uint64 = 0,0
	s := Transaction{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tID,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &Transaction{}, err
	}
	s.ID = *tID
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
	return i, &s, nil
}

// ----------------------------------------
//  Struct: Block
// ----------------------------------------

// Block type
type Block struct {
    ID Multihash `json:"id"`
    Header BlockHeader `json:"header"`
    ActiveData OpaqueActiveBlockData `json:"active_data"`
    PassiveData OpaquePassiveBlockData `json:"passive_data"`
    SignatureData VariableBlob `json:"signature_data"`
    Transactions VectorTransaction `json:"transactions"`
}

// NewBlock factory
func NewBlock() *Block {
	o := Block{}
	o.ID = *NewMultihash()
	o.Header = *NewBlockHeader()
	o.ActiveData = *NewOpaqueActiveBlockData()
	o.PassiveData = *NewOpaquePassiveBlockData()
	o.SignatureData = *NewVariableBlob()
	o.Transactions = *NewVectorTransaction()
	return &o
}

// Serialize Block
func (n Block) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ID.Serialize(vb)
	vb = n.Header.Serialize(vb)
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
	j,tID,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &Block{}, err
	}
	s.ID = *tID
	ovb = (*vb)[i:]
	j,tHeader,err := DeserializeBlockHeader(&ovb); i+=j
	if err != nil {
		return 0, &Block{}, err
	}
	s.Header = *tHeader
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
	j,tTransactions,err := DeserializeVectorTransaction(&ovb); i+=j
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
//  Struct: BlockRecord
// ----------------------------------------

// BlockRecord type
type BlockRecord struct {
    BlockID Multihash `json:"block_id"`
    BlockHeight BlockHeightType `json:"block_height"`
    PreviousBlockIds VectorMultihash `json:"previous_block_ids"`
    Block Block `json:"block"`
    BlockReceipt OpaqueBlockReceipt `json:"block_receipt"`
}

// NewBlockRecord factory
func NewBlockRecord() *BlockRecord {
	o := BlockRecord{}
	o.BlockID = *NewMultihash()
	o.BlockHeight = *NewBlockHeightType()
	o.PreviousBlockIds = *NewVectorMultihash()
	o.Block = *NewBlock()
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
	j,tBlock,err := DeserializeBlock(&ovb); i+=j
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
//  Struct: TransactionItem
// ----------------------------------------

// TransactionItem type
type TransactionItem struct {
    Transaction Transaction `json:"transaction"`
}

// NewTransactionItem factory
func NewTransactionItem() *TransactionItem {
	o := TransactionItem{}
	o.Transaction = *NewTransaction()
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
	j,tTransaction,err := DeserializeTransaction(&ovb); i+=j
	if err != nil {
		return 0, &TransactionItem{}, err
	}
	s.Transaction = *tTransaction
	return i, &s, nil
}

// ----------------------------------------
//  Struct: TransactionRecord
// ----------------------------------------

// TransactionRecord type
type TransactionRecord struct {
    Transaction Transaction `json:"transaction"`
}

// NewTransactionRecord factory
func NewTransactionRecord() *TransactionRecord {
	o := TransactionRecord{}
	o.Transaction = *NewTransaction()
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
	j,tTransaction,err := DeserializeTransaction(&ovb); i+=j
	if err != nil {
		return 0, &TransactionRecord{}, err
	}
	s.Transaction = *tTransaction
	return i, &s, nil
}

// ----------------------------------------
//  Struct: BlockStoreReservedRequest
// ----------------------------------------

// BlockStoreReservedRequest type
type BlockStoreReservedRequest struct {
}

// NewBlockStoreReservedRequest factory
func NewBlockStoreReservedRequest() *BlockStoreReservedRequest {
	o := BlockStoreReservedRequest{}
	return &o
}

// Serialize BlockStoreReservedRequest
func (n BlockStoreReservedRequest) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeBlockStoreReservedRequest function
func DeserializeBlockStoreReservedRequest(vb *VariableBlob) (uint64,*BlockStoreReservedRequest,error) {
	var i uint64 = 0
	s := BlockStoreReservedRequest{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: BlockStoreReservedResponse
// ----------------------------------------

// BlockStoreReservedResponse type
type BlockStoreReservedResponse struct {
}

// NewBlockStoreReservedResponse factory
func NewBlockStoreReservedResponse() *BlockStoreReservedResponse {
	o := BlockStoreReservedResponse{}
	return &o
}

// Serialize BlockStoreReservedResponse
func (n BlockStoreReservedResponse) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeBlockStoreReservedResponse function
func DeserializeBlockStoreReservedResponse(vb *VariableBlob) (uint64,*BlockStoreReservedResponse,error) {
	var i uint64 = 0
	s := BlockStoreReservedResponse{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetBlocksByIDRequest
// ----------------------------------------

// GetBlocksByIDRequest type
type GetBlocksByIDRequest struct {
    BlockID VectorMultihash `json:"block_id"`
    ReturnBlockBlob Boolean `json:"return_block_blob"`
    ReturnReceiptBlob Boolean `json:"return_receipt_blob"`
}

// NewGetBlocksByIDRequest factory
func NewGetBlocksByIDRequest() *GetBlocksByIDRequest {
	o := GetBlocksByIDRequest{}
	o.BlockID = *NewVectorMultihash()
	o.ReturnBlockBlob = *NewBoolean()
	o.ReturnReceiptBlob = *NewBoolean()
	return &o
}

// Serialize GetBlocksByIDRequest
func (n GetBlocksByIDRequest) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.BlockID.Serialize(vb)
	vb = n.ReturnBlockBlob.Serialize(vb)
	vb = n.ReturnReceiptBlob.Serialize(vb)
	return vb
}

// DeserializeGetBlocksByIDRequest function
func DeserializeGetBlocksByIDRequest(vb *VariableBlob) (uint64,*GetBlocksByIDRequest,error) {
	var i,j uint64 = 0,0
	s := GetBlocksByIDRequest{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tBlockID,err := DeserializeVectorMultihash(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByIDRequest{}, err
	}
	s.BlockID = *tBlockID
	ovb = (*vb)[i:]
	j,tReturnBlockBlob,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByIDRequest{}, err
	}
	s.ReturnBlockBlob = *tReturnBlockBlob
	ovb = (*vb)[i:]
	j,tReturnReceiptBlob,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByIDRequest{}, err
	}
	s.ReturnReceiptBlob = *tReturnReceiptBlob
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetBlocksByIDResponse
// ----------------------------------------

// GetBlocksByIDResponse type
type GetBlocksByIDResponse struct {
    BlockItems VectorBlockItem `json:"block_items"`
}

// NewGetBlocksByIDResponse factory
func NewGetBlocksByIDResponse() *GetBlocksByIDResponse {
	o := GetBlocksByIDResponse{}
	o.BlockItems = *NewVectorBlockItem()
	return &o
}

// Serialize GetBlocksByIDResponse
func (n GetBlocksByIDResponse) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.BlockItems.Serialize(vb)
	return vb
}

// DeserializeGetBlocksByIDResponse function
func DeserializeGetBlocksByIDResponse(vb *VariableBlob) (uint64,*GetBlocksByIDResponse,error) {
	var i,j uint64 = 0,0
	s := GetBlocksByIDResponse{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tBlockItems,err := DeserializeVectorBlockItem(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByIDResponse{}, err
	}
	s.BlockItems = *tBlockItems
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetBlocksByHeightRequest
// ----------------------------------------

// GetBlocksByHeightRequest type
type GetBlocksByHeightRequest struct {
    HeadBlockID Multihash `json:"head_block_id"`
    AncestorStartHeight BlockHeightType `json:"ancestor_start_height"`
    NumBlocks UInt32 `json:"num_blocks"`
    ReturnBlock Boolean `json:"return_block"`
    ReturnReceipt Boolean `json:"return_receipt"`
}

// NewGetBlocksByHeightRequest factory
func NewGetBlocksByHeightRequest() *GetBlocksByHeightRequest {
	o := GetBlocksByHeightRequest{}
	o.HeadBlockID = *NewMultihash()
	o.AncestorStartHeight = *NewBlockHeightType()
	o.NumBlocks = *NewUInt32()
	o.ReturnBlock = *NewBoolean()
	o.ReturnReceipt = *NewBoolean()
	return &o
}

// Serialize GetBlocksByHeightRequest
func (n GetBlocksByHeightRequest) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.HeadBlockID.Serialize(vb)
	vb = n.AncestorStartHeight.Serialize(vb)
	vb = n.NumBlocks.Serialize(vb)
	vb = n.ReturnBlock.Serialize(vb)
	vb = n.ReturnReceipt.Serialize(vb)
	return vb
}

// DeserializeGetBlocksByHeightRequest function
func DeserializeGetBlocksByHeightRequest(vb *VariableBlob) (uint64,*GetBlocksByHeightRequest,error) {
	var i,j uint64 = 0,0
	s := GetBlocksByHeightRequest{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tHeadBlockID,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByHeightRequest{}, err
	}
	s.HeadBlockID = *tHeadBlockID
	ovb = (*vb)[i:]
	j,tAncestorStartHeight,err := DeserializeBlockHeightType(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByHeightRequest{}, err
	}
	s.AncestorStartHeight = *tAncestorStartHeight
	ovb = (*vb)[i:]
	j,tNumBlocks,err := DeserializeUInt32(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByHeightRequest{}, err
	}
	s.NumBlocks = *tNumBlocks
	ovb = (*vb)[i:]
	j,tReturnBlock,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByHeightRequest{}, err
	}
	s.ReturnBlock = *tReturnBlock
	ovb = (*vb)[i:]
	j,tReturnReceipt,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByHeightRequest{}, err
	}
	s.ReturnReceipt = *tReturnReceipt
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetBlocksByHeightResponse
// ----------------------------------------

// GetBlocksByHeightResponse type
type GetBlocksByHeightResponse struct {
    BlockItems VectorBlockItem `json:"block_items"`
}

// NewGetBlocksByHeightResponse factory
func NewGetBlocksByHeightResponse() *GetBlocksByHeightResponse {
	o := GetBlocksByHeightResponse{}
	o.BlockItems = *NewVectorBlockItem()
	return &o
}

// Serialize GetBlocksByHeightResponse
func (n GetBlocksByHeightResponse) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.BlockItems.Serialize(vb)
	return vb
}

// DeserializeGetBlocksByHeightResponse function
func DeserializeGetBlocksByHeightResponse(vb *VariableBlob) (uint64,*GetBlocksByHeightResponse,error) {
	var i,j uint64 = 0,0
	s := GetBlocksByHeightResponse{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tBlockItems,err := DeserializeVectorBlockItem(&ovb); i+=j
	if err != nil {
		return 0, &GetBlocksByHeightResponse{}, err
	}
	s.BlockItems = *tBlockItems
	return i, &s, nil
}

// ----------------------------------------
//  Struct: AddBlockRequest
// ----------------------------------------

// AddBlockRequest type
type AddBlockRequest struct {
    BlockToAdd BlockItem `json:"block_to_add"`
}

// NewAddBlockRequest factory
func NewAddBlockRequest() *AddBlockRequest {
	o := AddBlockRequest{}
	o.BlockToAdd = *NewBlockItem()
	return &o
}

// Serialize AddBlockRequest
func (n AddBlockRequest) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.BlockToAdd.Serialize(vb)
	return vb
}

// DeserializeAddBlockRequest function
func DeserializeAddBlockRequest(vb *VariableBlob) (uint64,*AddBlockRequest,error) {
	var i,j uint64 = 0,0
	s := AddBlockRequest{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tBlockToAdd,err := DeserializeBlockItem(&ovb); i+=j
	if err != nil {
		return 0, &AddBlockRequest{}, err
	}
	s.BlockToAdd = *tBlockToAdd
	return i, &s, nil
}

// ----------------------------------------
//  Struct: AddBlockResponse
// ----------------------------------------

// AddBlockResponse type
type AddBlockResponse struct {
}

// NewAddBlockResponse factory
func NewAddBlockResponse() *AddBlockResponse {
	o := AddBlockResponse{}
	return &o
}

// Serialize AddBlockResponse
func (n AddBlockResponse) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeAddBlockResponse function
func DeserializeAddBlockResponse(vb *VariableBlob) (uint64,*AddBlockResponse,error) {
	var i uint64 = 0
	s := AddBlockResponse{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: AddTransactionRequest
// ----------------------------------------

// AddTransactionRequest type
type AddTransactionRequest struct {
    Transaction Transaction `json:"transaction"`
}

// NewAddTransactionRequest factory
func NewAddTransactionRequest() *AddTransactionRequest {
	o := AddTransactionRequest{}
	o.Transaction = *NewTransaction()
	return &o
}

// Serialize AddTransactionRequest
func (n AddTransactionRequest) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Transaction.Serialize(vb)
	return vb
}

// DeserializeAddTransactionRequest function
func DeserializeAddTransactionRequest(vb *VariableBlob) (uint64,*AddTransactionRequest,error) {
	var i,j uint64 = 0,0
	s := AddTransactionRequest{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransaction,err := DeserializeTransaction(&ovb); i+=j
	if err != nil {
		return 0, &AddTransactionRequest{}, err
	}
	s.Transaction = *tTransaction
	return i, &s, nil
}

// ----------------------------------------
//  Struct: AddTransactionResponse
// ----------------------------------------

// AddTransactionResponse type
type AddTransactionResponse struct {
}

// NewAddTransactionResponse factory
func NewAddTransactionResponse() *AddTransactionResponse {
	o := AddTransactionResponse{}
	return &o
}

// Serialize AddTransactionResponse
func (n AddTransactionResponse) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeAddTransactionResponse function
func DeserializeAddTransactionResponse(vb *VariableBlob) (uint64,*AddTransactionResponse,error) {
	var i uint64 = 0
	s := AddTransactionResponse{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetTransactionsByIDRequest
// ----------------------------------------

// GetTransactionsByIDRequest type
type GetTransactionsByIDRequest struct {
    TransactionIds VectorMultihash `json:"transaction_ids"`
}

// NewGetTransactionsByIDRequest factory
func NewGetTransactionsByIDRequest() *GetTransactionsByIDRequest {
	o := GetTransactionsByIDRequest{}
	o.TransactionIds = *NewVectorMultihash()
	return &o
}

// Serialize GetTransactionsByIDRequest
func (n GetTransactionsByIDRequest) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.TransactionIds.Serialize(vb)
	return vb
}

// DeserializeGetTransactionsByIDRequest function
func DeserializeGetTransactionsByIDRequest(vb *VariableBlob) (uint64,*GetTransactionsByIDRequest,error) {
	var i,j uint64 = 0,0
	s := GetTransactionsByIDRequest{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransactionIds,err := DeserializeVectorMultihash(&ovb); i+=j
	if err != nil {
		return 0, &GetTransactionsByIDRequest{}, err
	}
	s.TransactionIds = *tTransactionIds
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetTransactionsByIDResponse
// ----------------------------------------

// GetTransactionsByIDResponse type
type GetTransactionsByIDResponse struct {
    TransactionItems VectorTransactionItem `json:"transaction_items"`
}

// NewGetTransactionsByIDResponse factory
func NewGetTransactionsByIDResponse() *GetTransactionsByIDResponse {
	o := GetTransactionsByIDResponse{}
	o.TransactionItems = *NewVectorTransactionItem()
	return &o
}

// Serialize GetTransactionsByIDResponse
func (n GetTransactionsByIDResponse) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.TransactionItems.Serialize(vb)
	return vb
}

// DeserializeGetTransactionsByIDResponse function
func DeserializeGetTransactionsByIDResponse(vb *VariableBlob) (uint64,*GetTransactionsByIDResponse,error) {
	var i,j uint64 = 0,0
	s := GetTransactionsByIDResponse{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransactionItems,err := DeserializeVectorTransactionItem(&ovb); i+=j
	if err != nil {
		return 0, &GetTransactionsByIDResponse{}, err
	}
	s.TransactionItems = *tTransactionItems
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetHighestBlockRequest
// ----------------------------------------

// GetHighestBlockRequest type
type GetHighestBlockRequest struct {
}

// NewGetHighestBlockRequest factory
func NewGetHighestBlockRequest() *GetHighestBlockRequest {
	o := GetHighestBlockRequest{}
	return &o
}

// Serialize GetHighestBlockRequest
func (n GetHighestBlockRequest) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeGetHighestBlockRequest function
func DeserializeGetHighestBlockRequest(vb *VariableBlob) (uint64,*GetHighestBlockRequest,error) {
	var i uint64 = 0
	s := GetHighestBlockRequest{}
	
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
//  Struct: GetHighestBlockResponse
// ----------------------------------------

// GetHighestBlockResponse type
type GetHighestBlockResponse struct {
    Topology BlockTopology `json:"topology"`
}

// NewGetHighestBlockResponse factory
func NewGetHighestBlockResponse() *GetHighestBlockResponse {
	o := GetHighestBlockResponse{}
	o.Topology = *NewBlockTopology()
	return &o
}

// Serialize GetHighestBlockResponse
func (n GetHighestBlockResponse) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Topology.Serialize(vb)
	return vb
}

// DeserializeGetHighestBlockResponse function
func DeserializeGetHighestBlockResponse(vb *VariableBlob) (uint64,*GetHighestBlockResponse,error) {
	var i,j uint64 = 0,0
	s := GetHighestBlockResponse{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTopology,err := DeserializeBlockTopology(&ovb); i+=j
	if err != nil {
		return 0, &GetHighestBlockResponse{}, err
	}
	s.Topology = *tTopology
	return i, &s, nil
}

// ----------------------------------------
//  Struct: BlockStoreErrorResponse
// ----------------------------------------

// BlockStoreErrorResponse type
type BlockStoreErrorResponse struct {
    ErrorText String `json:"error_text"`
    ErrorData String `json:"error_data"`
}

// NewBlockStoreErrorResponse factory
func NewBlockStoreErrorResponse() *BlockStoreErrorResponse {
	o := BlockStoreErrorResponse{}
	o.ErrorText = *NewString()
	o.ErrorData = *NewString()
	return &o
}

// Serialize BlockStoreErrorResponse
func (n BlockStoreErrorResponse) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ErrorText.Serialize(vb)
	vb = n.ErrorData.Serialize(vb)
	return vb
}

// DeserializeBlockStoreErrorResponse function
func DeserializeBlockStoreErrorResponse(vb *VariableBlob) (uint64,*BlockStoreErrorResponse,error) {
	var i,j uint64 = 0,0
	s := BlockStoreErrorResponse{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tErrorText,err := DeserializeString(&ovb); i+=j
	if err != nil {
		return 0, &BlockStoreErrorResponse{}, err
	}
	s.ErrorText = *tErrorText
	ovb = (*vb)[i:]
	j,tErrorData,err := DeserializeString(&ovb); i+=j
	if err != nil {
		return 0, &BlockStoreErrorResponse{}, err
	}
	s.ErrorData = *tErrorData
	return i, &s, nil
}

// ----------------------------------------
//  Variant: BlockStoreRequest
// ----------------------------------------

// BlockStoreRequest type
type BlockStoreRequest struct {
	Value interface{}
}

// NewBlockStoreRequest factory
func NewBlockStoreRequest() *BlockStoreRequest {
	v := BlockStoreRequest{}
	v.Value = NewBlockStoreReservedRequest()
	return &v
}

// Serialize BlockStoreRequest
func (n BlockStoreRequest) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *BlockStoreReservedRequest:
			i = 0
		case *GetBlocksByIDRequest:
			i = 1
		case *GetBlocksByHeightRequest:
			i = 2
		case *AddBlockRequest:
			i = 3
		case *AddTransactionRequest:
			i = 4
		case *GetTransactionsByIDRequest:
			i = 5
		case *GetHighestBlockRequest:
			i = 6
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName BlockStoreRequest
func (n BlockStoreRequest) TypeToName() (string) {
	switch n.Value.(type) {
		case *BlockStoreReservedRequest:
			return "koinos::rpc::block_store::block_store_reserved_request"
		case *GetBlocksByIDRequest:
			return "koinos::rpc::block_store::get_blocks_by_id_request"
		case *GetBlocksByHeightRequest:
			return "koinos::rpc::block_store::get_blocks_by_height_request"
		case *AddBlockRequest:
			return "koinos::rpc::block_store::add_block_request"
		case *AddTransactionRequest:
			return "koinos::rpc::block_store::add_transaction_request"
		case *GetTransactionsByIDRequest:
			return "koinos::rpc::block_store::get_transactions_by_id_request"
		case *GetHighestBlockRequest:
			return "koinos::rpc::block_store::get_highest_block_request"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON BlockStoreRequest
func (n BlockStoreRequest) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeBlockStoreRequest function
func DeserializeBlockStoreRequest(vb *VariableBlob) (uint64,*BlockStoreRequest,error) {
	var v BlockStoreRequest
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewBlockStoreReservedRequest()
		case 1:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetBlocksByIDRequest(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 2:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetBlocksByHeightRequest(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 3:
			ovb := (*vb)[i:]
			k,x,err := DeserializeAddBlockRequest(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 4:
			ovb := (*vb)[i:]
			k,x,err := DeserializeAddTransactionRequest(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 5:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetTransactionsByIDRequest(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 6:
			v.Value = NewGetHighestBlockRequest()
		default:
			return 0, &v, errors.New("unknown variant tag")
	}
	return uint64(i)+j,&v,nil
}

// UnmarshalJSON *BlockStoreRequest
func (n *BlockStoreRequest) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::rpc::block_store::block_store_reserved_request":
			v := NewBlockStoreReservedRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::block_store::get_blocks_by_id_request":
			v := NewGetBlocksByIDRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::block_store::get_blocks_by_height_request":
			v := NewGetBlocksByHeightRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::block_store::add_block_request":
			v := NewAddBlockRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::block_store::add_transaction_request":
			v := NewAddTransactionRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::block_store::get_transactions_by_id_request":
			v := NewGetTransactionsByIDRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::block_store::get_highest_block_request":
			v := NewGetHighestBlockRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

	return nil
}


// ----------------------------------------
//  Variant: BlockStoreResponse
// ----------------------------------------

// BlockStoreResponse type
type BlockStoreResponse struct {
	Value interface{}
}

// NewBlockStoreResponse factory
func NewBlockStoreResponse() *BlockStoreResponse {
	v := BlockStoreResponse{}
	v.Value = NewBlockStoreReservedResponse()
	return &v
}

// Serialize BlockStoreResponse
func (n BlockStoreResponse) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *BlockStoreReservedResponse:
			i = 0
		case *BlockStoreErrorResponse:
			i = 1
		case *GetBlocksByIDResponse:
			i = 2
		case *GetBlocksByHeightResponse:
			i = 3
		case *AddBlockResponse:
			i = 4
		case *AddTransactionResponse:
			i = 5
		case *GetTransactionsByIDResponse:
			i = 6
		case *GetHighestBlockResponse:
			i = 7
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName BlockStoreResponse
func (n BlockStoreResponse) TypeToName() (string) {
	switch n.Value.(type) {
		case *BlockStoreReservedResponse:
			return "koinos::rpc::block_store::block_store_reserved_response"
		case *BlockStoreErrorResponse:
			return "koinos::rpc::block_store::block_store_error_response"
		case *GetBlocksByIDResponse:
			return "koinos::rpc::block_store::get_blocks_by_id_response"
		case *GetBlocksByHeightResponse:
			return "koinos::rpc::block_store::get_blocks_by_height_response"
		case *AddBlockResponse:
			return "koinos::rpc::block_store::add_block_response"
		case *AddTransactionResponse:
			return "koinos::rpc::block_store::add_transaction_response"
		case *GetTransactionsByIDResponse:
			return "koinos::rpc::block_store::get_transactions_by_id_response"
		case *GetHighestBlockResponse:
			return "koinos::rpc::block_store::get_highest_block_response"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON BlockStoreResponse
func (n BlockStoreResponse) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeBlockStoreResponse function
func DeserializeBlockStoreResponse(vb *VariableBlob) (uint64,*BlockStoreResponse,error) {
	var v BlockStoreResponse
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewBlockStoreReservedResponse()
		case 1:
			ovb := (*vb)[i:]
			k,x,err := DeserializeBlockStoreErrorResponse(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 2:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetBlocksByIDResponse(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 3:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetBlocksByHeightResponse(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 4:
			v.Value = NewAddBlockResponse()
		case 5:
			v.Value = NewAddTransactionResponse()
		case 6:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetTransactionsByIDResponse(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 7:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetHighestBlockResponse(&ovb)
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

// UnmarshalJSON *BlockStoreResponse
func (n *BlockStoreResponse) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::rpc::block_store::block_store_reserved_response":
			v := NewBlockStoreReservedResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::block_store::block_store_error_response":
			v := NewBlockStoreErrorResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::block_store::get_blocks_by_id_response":
			v := NewGetBlocksByIDResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::block_store::get_blocks_by_height_response":
			v := NewGetBlocksByHeightResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::block_store::add_block_response":
			v := NewAddBlockResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::block_store::add_transaction_response":
			v := NewAddTransactionResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::block_store::get_transactions_by_id_response":
			v := NewGetTransactionsByIDResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::block_store::get_highest_block_response":
			v := NewGetHighestBlockResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

	return nil
}


// ----------------------------------------
//  Typedef: AccountType
// ----------------------------------------

// AccountType type
type AccountType VariableBlob

// NewAccountType factory
func NewAccountType() *AccountType {
	o := AccountType(*NewVariableBlob())
	return &o
}

// Serialize AccountType
func (n AccountType) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VariableBlob(n)
	return ox.Serialize(vb)
}

// DeserializeAccountType function
func DeserializeAccountType(vb *VariableBlob) (uint64,*AccountType,error) {
	var ot AccountType
	i,n,err := DeserializeVariableBlob(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = AccountType(*n)
	return i,&ot,nil}

// MarshalJSON AccountType
func (n AccountType) MarshalJSON() ([]byte, error) {
	v := VariableBlob(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *AccountType
func (n *AccountType) UnmarshalJSON(data []byte) error {
	v := VariableBlob(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = AccountType(v)
	return nil
}


// ----------------------------------------
//  Struct: TransactionAccepted
// ----------------------------------------

// TransactionAccepted type
type TransactionAccepted struct {
    Transaction Transaction `json:"transaction"`
    Payer AccountType `json:"payer"`
    MaxPayerResources UInt128 `json:"max_payer_resources"`
    TrxResourceLimit UInt128 `json:"trx_resource_limit"`
    Height BlockHeightType `json:"height"`
}

// NewTransactionAccepted factory
func NewTransactionAccepted() *TransactionAccepted {
	o := TransactionAccepted{}
	o.Transaction = *NewTransaction()
	o.Payer = *NewAccountType()
	o.MaxPayerResources = *NewUInt128()
	o.TrxResourceLimit = *NewUInt128()
	o.Height = *NewBlockHeightType()
	return &o
}

// Serialize TransactionAccepted
func (n TransactionAccepted) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Transaction.Serialize(vb)
	vb = n.Payer.Serialize(vb)
	vb = n.MaxPayerResources.Serialize(vb)
	vb = n.TrxResourceLimit.Serialize(vb)
	vb = n.Height.Serialize(vb)
	return vb
}

// DeserializeTransactionAccepted function
func DeserializeTransactionAccepted(vb *VariableBlob) (uint64,*TransactionAccepted,error) {
	var i,j uint64 = 0,0
	s := TransactionAccepted{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransaction,err := DeserializeTransaction(&ovb); i+=j
	if err != nil {
		return 0, &TransactionAccepted{}, err
	}
	s.Transaction = *tTransaction
	ovb = (*vb)[i:]
	j,tPayer,err := DeserializeAccountType(&ovb); i+=j
	if err != nil {
		return 0, &TransactionAccepted{}, err
	}
	s.Payer = *tPayer
	ovb = (*vb)[i:]
	j,tMaxPayerResources,err := DeserializeUInt128(&ovb); i+=j
	if err != nil {
		return 0, &TransactionAccepted{}, err
	}
	s.MaxPayerResources = *tMaxPayerResources
	ovb = (*vb)[i:]
	j,tTrxResourceLimit,err := DeserializeUInt128(&ovb); i+=j
	if err != nil {
		return 0, &TransactionAccepted{}, err
	}
	s.TrxResourceLimit = *tTrxResourceLimit
	ovb = (*vb)[i:]
	j,tHeight,err := DeserializeBlockHeightType(&ovb); i+=j
	if err != nil {
		return 0, &TransactionAccepted{}, err
	}
	s.Height = *tHeight
	return i, &s, nil
}

// ----------------------------------------
//  Struct: BlockAccepted
// ----------------------------------------

// BlockAccepted type
type BlockAccepted struct {
    Block Block `json:"block"`
}

// NewBlockAccepted factory
func NewBlockAccepted() *BlockAccepted {
	o := BlockAccepted{}
	o.Block = *NewBlock()
	return &o
}

// Serialize BlockAccepted
func (n BlockAccepted) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Block.Serialize(vb)
	return vb
}

// DeserializeBlockAccepted function
func DeserializeBlockAccepted(vb *VariableBlob) (uint64,*BlockAccepted,error) {
	var i,j uint64 = 0,0
	s := BlockAccepted{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tBlock,err := DeserializeBlock(&ovb); i+=j
	if err != nil {
		return 0, &BlockAccepted{}, err
	}
	s.Block = *tBlock
	return i, &s, nil
}

// ----------------------------------------
//  Struct: BlockIrreversible
// ----------------------------------------

// BlockIrreversible type
type BlockIrreversible struct {
    Topology BlockTopology `json:"topology"`
}

// NewBlockIrreversible factory
func NewBlockIrreversible() *BlockIrreversible {
	o := BlockIrreversible{}
	o.Topology = *NewBlockTopology()
	return &o
}

// Serialize BlockIrreversible
func (n BlockIrreversible) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Topology.Serialize(vb)
	return vb
}

// DeserializeBlockIrreversible function
func DeserializeBlockIrreversible(vb *VariableBlob) (uint64,*BlockIrreversible,error) {
	var i,j uint64 = 0,0
	s := BlockIrreversible{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTopology,err := DeserializeBlockTopology(&ovb); i+=j
	if err != nil {
		return 0, &BlockIrreversible{}, err
	}
	s.Topology = *tTopology
	return i, &s, nil
}

// ----------------------------------------
//  Struct: ForkHeads
// ----------------------------------------

// ForkHeads type
type ForkHeads struct {
    ForkHeads VectorBlockTopology `json:"fork_heads"`
    LastIrreversibleBlock BlockTopology `json:"last_irreversible_block"`
}

// NewForkHeads factory
func NewForkHeads() *ForkHeads {
	o := ForkHeads{}
	o.ForkHeads = *NewVectorBlockTopology()
	o.LastIrreversibleBlock = *NewBlockTopology()
	return &o
}

// Serialize ForkHeads
func (n ForkHeads) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ForkHeads.Serialize(vb)
	vb = n.LastIrreversibleBlock.Serialize(vb)
	return vb
}

// DeserializeForkHeads function
func DeserializeForkHeads(vb *VariableBlob) (uint64,*ForkHeads,error) {
	var i,j uint64 = 0,0
	s := ForkHeads{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tForkHeads,err := DeserializeVectorBlockTopology(&ovb); i+=j
	if err != nil {
		return 0, &ForkHeads{}, err
	}
	s.ForkHeads = *tForkHeads
	ovb = (*vb)[i:]
	j,tLastIrreversibleBlock,err := DeserializeBlockTopology(&ovb); i+=j
	if err != nil {
		return 0, &ForkHeads{}, err
	}
	s.LastIrreversibleBlock = *tLastIrreversibleBlock
	return i, &s, nil
}

// ----------------------------------------
//  Enum: SystemCallID
// ----------------------------------------

// {'name': 'system_call_id', 'entries': [{'name': 'prints', 'value': 2602735937, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_block', 'value': 2494255093, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_transaction', 'value': 2643154394, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_reserved_operation', 'value': 2594724132, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_upload_contract_operation', 'value': 2658052407, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_execute_contract_operation', 'value': 2451064454, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'apply_set_system_call_operation', 'value': 2507777116, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_put_object', 'value': 2535376802, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_get_object', 'value': 2540087547, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_get_next_object', 'value': 2577635560, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'db_get_prev_object', 'value': 2614326908, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'execute_contract', 'value': 2562796798, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_entry_point', 'value': 2683439069, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_contract_args_size', 'value': 2601357273, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_contract_args', 'value': 2679873944, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'set_contract_return', 'value': 2672414186, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'exit_contract', 'value': 2564781488, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_head_info', 'value': 2507125293, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'hash', 'value': 2574716420, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'verify_block_signature', 'value': 2411308443, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'verify_merkle_root', 'value': 2574132409, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_transaction_payer', 'value': 2259188725, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_max_account_resources', 'value': 2217503873, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_transaction_resource_limit', 'value': 2571171461, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_last_irreversible_block', 'value': 2503814711, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_caller', 'value': 2484563039, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'require_authority', 'value': 2492589352, 'doc': '', 'info': {'type': 'EnumEntry'}}, {'name': 'get_transaction_signature', 'value': 2646862900, 'doc': '', 'info': {'type': 'EnumEntry'}}], 'tref': {'name': ['koinos', 'uint32'], 'targs': None, 'info': {'type': 'Typeref'}}, 'doc': '', 'info': {'type': 'EnumClass'}}

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
	SystemCallIDGetEntryPoint SystemCallID = 2683439069
	SystemCallIDGetContractArgsSize SystemCallID = 2601357273
	SystemCallIDGetContractArgs SystemCallID = 2679873944
	SystemCallIDSetContractReturn SystemCallID = 2672414186
	SystemCallIDExitContract SystemCallID = 2564781488
	SystemCallIDGetHeadInfo SystemCallID = 2507125293
	SystemCallIDHash SystemCallID = 2574716420
	SystemCallIDVerifyBlockSignature SystemCallID = 2411308443
	SystemCallIDVerifyMerkleRoot SystemCallID = 2574132409
	SystemCallIDGetTransactionPayer SystemCallID = 2259188725
	SystemCallIDGetMaxAccountResources SystemCallID = 2217503873
	SystemCallIDGetTransactionResourceLimit SystemCallID = 2571171461
	SystemCallIDGetLastIrreversibleBlock SystemCallID = 2503814711
	SystemCallIDGetCaller SystemCallID = 2484563039
	SystemCallIDRequireAuthority SystemCallID = 2492589352
	SystemCallIDGetTransactionSignature SystemCallID = 2646862900
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
		case SystemCallIDGetEntryPoint:
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
		case SystemCallIDVerifyBlockSignature:
			return true
		case SystemCallIDVerifyMerkleRoot:
			return true
		case SystemCallIDGetTransactionPayer:
			return true
		case SystemCallIDGetMaxAccountResources:
			return true
		case SystemCallIDGetTransactionResourceLimit:
			return true
		case SystemCallIDGetLastIrreversibleBlock:
			return true
		case SystemCallIDGetCaller:
			return true
		case SystemCallIDRequireAuthority:
			return true
		case SystemCallIDGetTransactionSignature:
			return true
		default:
			return false
	}
}


// ----------------------------------------
//  Struct: HeadInfo
// ----------------------------------------

// HeadInfo type
type HeadInfo struct {
    HeadTopology BlockTopology `json:"head_topology"`
    LastIrreversibleHeight BlockHeightType `json:"last_irreversible_height"`
}

// NewHeadInfo factory
func NewHeadInfo() *HeadInfo {
	o := HeadInfo{}
	o.HeadTopology = *NewBlockTopology()
	o.LastIrreversibleHeight = *NewBlockHeightType()
	return &o
}

// Serialize HeadInfo
func (n HeadInfo) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.HeadTopology.Serialize(vb)
	vb = n.LastIrreversibleHeight.Serialize(vb)
	return vb
}

// DeserializeHeadInfo function
func DeserializeHeadInfo(vb *VariableBlob) (uint64,*HeadInfo,error) {
	var i,j uint64 = 0,0
	s := HeadInfo{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tHeadTopology,err := DeserializeBlockTopology(&ovb); i+=j
	if err != nil {
		return 0, &HeadInfo{}, err
	}
	s.HeadTopology = *tHeadTopology
	ovb = (*vb)[i:]
	j,tLastIrreversibleHeight,err := DeserializeBlockHeightType(&ovb); i+=j
	if err != nil {
		return 0, &HeadInfo{}, err
	}
	s.LastIrreversibleHeight = *tLastIrreversibleHeight
	return i, &s, nil
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
//  Typedef: PrintsReturn
// ----------------------------------------

// PrintsReturn type
type PrintsReturn VoidType

// NewPrintsReturn factory
func NewPrintsReturn() *PrintsReturn {
	o := PrintsReturn(*NewVoidType())
	return &o
}

// Serialize PrintsReturn
func (n PrintsReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializePrintsReturn function
func DeserializePrintsReturn(vb *VariableBlob) (uint64,*PrintsReturn,error) {
	var ot PrintsReturn
	return 0,&ot,nil}

// MarshalJSON PrintsReturn
func (n PrintsReturn) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *PrintsReturn
func (n *PrintsReturn) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = PrintsReturn(v)
	return nil
}


// ----------------------------------------
//  Struct: VerifyBlockSignatureArgs
// ----------------------------------------

// VerifyBlockSignatureArgs type
type VerifyBlockSignatureArgs struct {
    SignatureData VariableBlob `json:"signature_data"`
    Digest Multihash `json:"digest"`
}

// NewVerifyBlockSignatureArgs factory
func NewVerifyBlockSignatureArgs() *VerifyBlockSignatureArgs {
	o := VerifyBlockSignatureArgs{}
	o.SignatureData = *NewVariableBlob()
	o.Digest = *NewMultihash()
	return &o
}

// Serialize VerifyBlockSignatureArgs
func (n VerifyBlockSignatureArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.SignatureData.Serialize(vb)
	vb = n.Digest.Serialize(vb)
	return vb
}

// DeserializeVerifyBlockSignatureArgs function
func DeserializeVerifyBlockSignatureArgs(vb *VariableBlob) (uint64,*VerifyBlockSignatureArgs,error) {
	var i,j uint64 = 0,0
	s := VerifyBlockSignatureArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tSignatureData,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &VerifyBlockSignatureArgs{}, err
	}
	s.SignatureData = *tSignatureData
	ovb = (*vb)[i:]
	j,tDigest,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &VerifyBlockSignatureArgs{}, err
	}
	s.Digest = *tDigest
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: VerifyBlockSignatureReturn
// ----------------------------------------

// VerifyBlockSignatureReturn type
type VerifyBlockSignatureReturn Boolean

// NewVerifyBlockSignatureReturn factory
func NewVerifyBlockSignatureReturn() *VerifyBlockSignatureReturn {
	o := VerifyBlockSignatureReturn(*NewBoolean())
	return &o
}

// Serialize VerifyBlockSignatureReturn
func (n VerifyBlockSignatureReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := Boolean(n)
	return ox.Serialize(vb)
}

// DeserializeVerifyBlockSignatureReturn function
func DeserializeVerifyBlockSignatureReturn(vb *VariableBlob) (uint64,*VerifyBlockSignatureReturn,error) {
	var ot VerifyBlockSignatureReturn
	i,n,err := DeserializeBoolean(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = VerifyBlockSignatureReturn(*n)
	return i,&ot,nil}

// MarshalJSON VerifyBlockSignatureReturn
func (n VerifyBlockSignatureReturn) MarshalJSON() ([]byte, error) {
	v := Boolean(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *VerifyBlockSignatureReturn
func (n *VerifyBlockSignatureReturn) UnmarshalJSON(data []byte) error {
	v := Boolean(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = VerifyBlockSignatureReturn(v)
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
//  Typedef: VerifyMerkleRootReturn
// ----------------------------------------

// VerifyMerkleRootReturn type
type VerifyMerkleRootReturn Boolean

// NewVerifyMerkleRootReturn factory
func NewVerifyMerkleRootReturn() *VerifyMerkleRootReturn {
	o := VerifyMerkleRootReturn(*NewBoolean())
	return &o
}

// Serialize VerifyMerkleRootReturn
func (n VerifyMerkleRootReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := Boolean(n)
	return ox.Serialize(vb)
}

// DeserializeVerifyMerkleRootReturn function
func DeserializeVerifyMerkleRootReturn(vb *VariableBlob) (uint64,*VerifyMerkleRootReturn,error) {
	var ot VerifyMerkleRootReturn
	i,n,err := DeserializeBoolean(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = VerifyMerkleRootReturn(*n)
	return i,&ot,nil}

// MarshalJSON VerifyMerkleRootReturn
func (n VerifyMerkleRootReturn) MarshalJSON() ([]byte, error) {
	v := Boolean(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *VerifyMerkleRootReturn
func (n *VerifyMerkleRootReturn) UnmarshalJSON(data []byte) error {
	v := Boolean(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = VerifyMerkleRootReturn(v)
	return nil
}


// ----------------------------------------
//  Struct: ApplyBlockArgs
// ----------------------------------------

// ApplyBlockArgs type
type ApplyBlockArgs struct {
    Block Block `json:"block"`
    CheckPassiveData Boolean `json:"check_passive_data"`
    CheckBlockSignature Boolean `json:"check_block_signature"`
    CheckTransactionSignatures Boolean `json:"check_transaction_signatures"`
}

// NewApplyBlockArgs factory
func NewApplyBlockArgs() *ApplyBlockArgs {
	o := ApplyBlockArgs{}
	o.Block = *NewBlock()
	o.CheckPassiveData = *NewBoolean()
	o.CheckBlockSignature = *NewBoolean()
	o.CheckTransactionSignatures = *NewBoolean()
	return &o
}

// Serialize ApplyBlockArgs
func (n ApplyBlockArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Block.Serialize(vb)
	vb = n.CheckPassiveData.Serialize(vb)
	vb = n.CheckBlockSignature.Serialize(vb)
	vb = n.CheckTransactionSignatures.Serialize(vb)
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
	j,tCheckPassiveData,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &ApplyBlockArgs{}, err
	}
	s.CheckPassiveData = *tCheckPassiveData
	ovb = (*vb)[i:]
	j,tCheckBlockSignature,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &ApplyBlockArgs{}, err
	}
	s.CheckBlockSignature = *tCheckBlockSignature
	ovb = (*vb)[i:]
	j,tCheckTransactionSignatures,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &ApplyBlockArgs{}, err
	}
	s.CheckTransactionSignatures = *tCheckTransactionSignatures
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: ApplyBlockReturn
// ----------------------------------------

// ApplyBlockReturn type
type ApplyBlockReturn VoidType

// NewApplyBlockReturn factory
func NewApplyBlockReturn() *ApplyBlockReturn {
	o := ApplyBlockReturn(*NewVoidType())
	return &o
}

// Serialize ApplyBlockReturn
func (n ApplyBlockReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeApplyBlockReturn function
func DeserializeApplyBlockReturn(vb *VariableBlob) (uint64,*ApplyBlockReturn,error) {
	var ot ApplyBlockReturn
	return 0,&ot,nil}

// MarshalJSON ApplyBlockReturn
func (n ApplyBlockReturn) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ApplyBlockReturn
func (n *ApplyBlockReturn) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ApplyBlockReturn(v)
	return nil
}


// ----------------------------------------
//  Struct: ApplyTransactionArgs
// ----------------------------------------

// ApplyTransactionArgs type
type ApplyTransactionArgs struct {
    Transaction Transaction `json:"transaction"`
}

// NewApplyTransactionArgs factory
func NewApplyTransactionArgs() *ApplyTransactionArgs {
	o := ApplyTransactionArgs{}
	o.Transaction = *NewTransaction()
	return &o
}

// Serialize ApplyTransactionArgs
func (n ApplyTransactionArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Transaction.Serialize(vb)
	return vb
}

// DeserializeApplyTransactionArgs function
func DeserializeApplyTransactionArgs(vb *VariableBlob) (uint64,*ApplyTransactionArgs,error) {
	var i,j uint64 = 0,0
	s := ApplyTransactionArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransaction,err := DeserializeTransaction(&ovb); i+=j
	if err != nil {
		return 0, &ApplyTransactionArgs{}, err
	}
	s.Transaction = *tTransaction
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: ApplyTransactionReturn
// ----------------------------------------

// ApplyTransactionReturn type
type ApplyTransactionReturn VoidType

// NewApplyTransactionReturn factory
func NewApplyTransactionReturn() *ApplyTransactionReturn {
	o := ApplyTransactionReturn(*NewVoidType())
	return &o
}

// Serialize ApplyTransactionReturn
func (n ApplyTransactionReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeApplyTransactionReturn function
func DeserializeApplyTransactionReturn(vb *VariableBlob) (uint64,*ApplyTransactionReturn,error) {
	var ot ApplyTransactionReturn
	return 0,&ot,nil}

// MarshalJSON ApplyTransactionReturn
func (n ApplyTransactionReturn) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ApplyTransactionReturn
func (n *ApplyTransactionReturn) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ApplyTransactionReturn(v)
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
//  Typedef: ApplyUploadContractOperationReturn
// ----------------------------------------

// ApplyUploadContractOperationReturn type
type ApplyUploadContractOperationReturn VoidType

// NewApplyUploadContractOperationReturn factory
func NewApplyUploadContractOperationReturn() *ApplyUploadContractOperationReturn {
	o := ApplyUploadContractOperationReturn(*NewVoidType())
	return &o
}

// Serialize ApplyUploadContractOperationReturn
func (n ApplyUploadContractOperationReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeApplyUploadContractOperationReturn function
func DeserializeApplyUploadContractOperationReturn(vb *VariableBlob) (uint64,*ApplyUploadContractOperationReturn,error) {
	var ot ApplyUploadContractOperationReturn
	return 0,&ot,nil}

// MarshalJSON ApplyUploadContractOperationReturn
func (n ApplyUploadContractOperationReturn) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ApplyUploadContractOperationReturn
func (n *ApplyUploadContractOperationReturn) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ApplyUploadContractOperationReturn(v)
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
//  Typedef: ApplyReservedOperationReturn
// ----------------------------------------

// ApplyReservedOperationReturn type
type ApplyReservedOperationReturn VoidType

// NewApplyReservedOperationReturn factory
func NewApplyReservedOperationReturn() *ApplyReservedOperationReturn {
	o := ApplyReservedOperationReturn(*NewVoidType())
	return &o
}

// Serialize ApplyReservedOperationReturn
func (n ApplyReservedOperationReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeApplyReservedOperationReturn function
func DeserializeApplyReservedOperationReturn(vb *VariableBlob) (uint64,*ApplyReservedOperationReturn,error) {
	var ot ApplyReservedOperationReturn
	return 0,&ot,nil}

// MarshalJSON ApplyReservedOperationReturn
func (n ApplyReservedOperationReturn) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ApplyReservedOperationReturn
func (n *ApplyReservedOperationReturn) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ApplyReservedOperationReturn(v)
	return nil
}


// ----------------------------------------
//  Struct: ApplyExecuteContractOperationArgs
// ----------------------------------------

// ApplyExecuteContractOperationArgs type
type ApplyExecuteContractOperationArgs struct {
    Op CallContractOperation `json:"op"`
}

// NewApplyExecuteContractOperationArgs factory
func NewApplyExecuteContractOperationArgs() *ApplyExecuteContractOperationArgs {
	o := ApplyExecuteContractOperationArgs{}
	o.Op = *NewCallContractOperation()
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
	j,tOp,err := DeserializeCallContractOperation(&ovb); i+=j
	if err != nil {
		return 0, &ApplyExecuteContractOperationArgs{}, err
	}
	s.Op = *tOp
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: ApplyExecuteContractOperationReturn
// ----------------------------------------

// ApplyExecuteContractOperationReturn type
type ApplyExecuteContractOperationReturn VoidType

// NewApplyExecuteContractOperationReturn factory
func NewApplyExecuteContractOperationReturn() *ApplyExecuteContractOperationReturn {
	o := ApplyExecuteContractOperationReturn(*NewVoidType())
	return &o
}

// Serialize ApplyExecuteContractOperationReturn
func (n ApplyExecuteContractOperationReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeApplyExecuteContractOperationReturn function
func DeserializeApplyExecuteContractOperationReturn(vb *VariableBlob) (uint64,*ApplyExecuteContractOperationReturn,error) {
	var ot ApplyExecuteContractOperationReturn
	return 0,&ot,nil}

// MarshalJSON ApplyExecuteContractOperationReturn
func (n ApplyExecuteContractOperationReturn) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ApplyExecuteContractOperationReturn
func (n *ApplyExecuteContractOperationReturn) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ApplyExecuteContractOperationReturn(v)
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
//  Typedef: ApplySetSystemCallOperationReturn
// ----------------------------------------

// ApplySetSystemCallOperationReturn type
type ApplySetSystemCallOperationReturn VoidType

// NewApplySetSystemCallOperationReturn factory
func NewApplySetSystemCallOperationReturn() *ApplySetSystemCallOperationReturn {
	o := ApplySetSystemCallOperationReturn(*NewVoidType())
	return &o
}

// Serialize ApplySetSystemCallOperationReturn
func (n ApplySetSystemCallOperationReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeApplySetSystemCallOperationReturn function
func DeserializeApplySetSystemCallOperationReturn(vb *VariableBlob) (uint64,*ApplySetSystemCallOperationReturn,error) {
	var ot ApplySetSystemCallOperationReturn
	return 0,&ot,nil}

// MarshalJSON ApplySetSystemCallOperationReturn
func (n ApplySetSystemCallOperationReturn) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ApplySetSystemCallOperationReturn
func (n *ApplySetSystemCallOperationReturn) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ApplySetSystemCallOperationReturn(v)
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
//  Typedef: DbPutObjectReturn
// ----------------------------------------

// DbPutObjectReturn type
type DbPutObjectReturn Boolean

// NewDbPutObjectReturn factory
func NewDbPutObjectReturn() *DbPutObjectReturn {
	o := DbPutObjectReturn(*NewBoolean())
	return &o
}

// Serialize DbPutObjectReturn
func (n DbPutObjectReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := Boolean(n)
	return ox.Serialize(vb)
}

// DeserializeDbPutObjectReturn function
func DeserializeDbPutObjectReturn(vb *VariableBlob) (uint64,*DbPutObjectReturn,error) {
	var ot DbPutObjectReturn
	i,n,err := DeserializeBoolean(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = DbPutObjectReturn(*n)
	return i,&ot,nil}

// MarshalJSON DbPutObjectReturn
func (n DbPutObjectReturn) MarshalJSON() ([]byte, error) {
	v := Boolean(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *DbPutObjectReturn
func (n *DbPutObjectReturn) UnmarshalJSON(data []byte) error {
	v := Boolean(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = DbPutObjectReturn(v)
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
//  Typedef: DbGetObjectReturn
// ----------------------------------------

// DbGetObjectReturn type
type DbGetObjectReturn VariableBlob

// NewDbGetObjectReturn factory
func NewDbGetObjectReturn() *DbGetObjectReturn {
	o := DbGetObjectReturn(*NewVariableBlob())
	return &o
}

// Serialize DbGetObjectReturn
func (n DbGetObjectReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VariableBlob(n)
	return ox.Serialize(vb)
}

// DeserializeDbGetObjectReturn function
func DeserializeDbGetObjectReturn(vb *VariableBlob) (uint64,*DbGetObjectReturn,error) {
	var ot DbGetObjectReturn
	i,n,err := DeserializeVariableBlob(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = DbGetObjectReturn(*n)
	return i,&ot,nil}

// MarshalJSON DbGetObjectReturn
func (n DbGetObjectReturn) MarshalJSON() ([]byte, error) {
	v := VariableBlob(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *DbGetObjectReturn
func (n *DbGetObjectReturn) UnmarshalJSON(data []byte) error {
	v := VariableBlob(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = DbGetObjectReturn(v)
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
//  Typedef: DbGetNextObjectReturn
// ----------------------------------------

// DbGetNextObjectReturn type
type DbGetNextObjectReturn DbGetObjectReturn

// NewDbGetNextObjectReturn factory
func NewDbGetNextObjectReturn() *DbGetNextObjectReturn {
	o := DbGetNextObjectReturn(*NewDbGetObjectReturn())
	return &o
}

// Serialize DbGetNextObjectReturn
func (n DbGetNextObjectReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := DbGetObjectReturn(n)
	return ox.Serialize(vb)
}

// DeserializeDbGetNextObjectReturn function
func DeserializeDbGetNextObjectReturn(vb *VariableBlob) (uint64,*DbGetNextObjectReturn,error) {
	var ot DbGetNextObjectReturn
	i,n,err := DeserializeDbGetObjectReturn(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = DbGetNextObjectReturn(*n)
	return i,&ot,nil}

// MarshalJSON DbGetNextObjectReturn
func (n DbGetNextObjectReturn) MarshalJSON() ([]byte, error) {
	v := DbGetObjectReturn(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *DbGetNextObjectReturn
func (n *DbGetNextObjectReturn) UnmarshalJSON(data []byte) error {
	v := DbGetObjectReturn(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = DbGetNextObjectReturn(v)
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
//  Typedef: DbGetPrevObjectReturn
// ----------------------------------------

// DbGetPrevObjectReturn type
type DbGetPrevObjectReturn DbGetObjectReturn

// NewDbGetPrevObjectReturn factory
func NewDbGetPrevObjectReturn() *DbGetPrevObjectReturn {
	o := DbGetPrevObjectReturn(*NewDbGetObjectReturn())
	return &o
}

// Serialize DbGetPrevObjectReturn
func (n DbGetPrevObjectReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := DbGetObjectReturn(n)
	return ox.Serialize(vb)
}

// DeserializeDbGetPrevObjectReturn function
func DeserializeDbGetPrevObjectReturn(vb *VariableBlob) (uint64,*DbGetPrevObjectReturn,error) {
	var ot DbGetPrevObjectReturn
	i,n,err := DeserializeDbGetObjectReturn(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = DbGetPrevObjectReturn(*n)
	return i,&ot,nil}

// MarshalJSON DbGetPrevObjectReturn
func (n DbGetPrevObjectReturn) MarshalJSON() ([]byte, error) {
	v := DbGetObjectReturn(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *DbGetPrevObjectReturn
func (n *DbGetPrevObjectReturn) UnmarshalJSON(data []byte) error {
	v := DbGetObjectReturn(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = DbGetPrevObjectReturn(v)
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
//  Typedef: ExecuteContractReturn
// ----------------------------------------

// ExecuteContractReturn type
type ExecuteContractReturn VariableBlob

// NewExecuteContractReturn factory
func NewExecuteContractReturn() *ExecuteContractReturn {
	o := ExecuteContractReturn(*NewVariableBlob())
	return &o
}

// Serialize ExecuteContractReturn
func (n ExecuteContractReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VariableBlob(n)
	return ox.Serialize(vb)
}

// DeserializeExecuteContractReturn function
func DeserializeExecuteContractReturn(vb *VariableBlob) (uint64,*ExecuteContractReturn,error) {
	var ot ExecuteContractReturn
	i,n,err := DeserializeVariableBlob(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = ExecuteContractReturn(*n)
	return i,&ot,nil}

// MarshalJSON ExecuteContractReturn
func (n ExecuteContractReturn) MarshalJSON() ([]byte, error) {
	v := VariableBlob(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ExecuteContractReturn
func (n *ExecuteContractReturn) UnmarshalJSON(data []byte) error {
	v := VariableBlob(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ExecuteContractReturn(v)
	return nil
}


// ----------------------------------------
//  Typedef: GetEntryPointArgs
// ----------------------------------------

// GetEntryPointArgs type
type GetEntryPointArgs VoidType

// NewGetEntryPointArgs factory
func NewGetEntryPointArgs() *GetEntryPointArgs {
	o := GetEntryPointArgs(*NewVoidType())
	return &o
}

// Serialize GetEntryPointArgs
func (n GetEntryPointArgs) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeGetEntryPointArgs function
func DeserializeGetEntryPointArgs(vb *VariableBlob) (uint64,*GetEntryPointArgs,error) {
	var ot GetEntryPointArgs
	return 0,&ot,nil}

// MarshalJSON GetEntryPointArgs
func (n GetEntryPointArgs) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetEntryPointArgs
func (n *GetEntryPointArgs) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetEntryPointArgs(v)
	return nil
}


// ----------------------------------------
//  Typedef: GetEntryPointReturn
// ----------------------------------------

// GetEntryPointReturn type
type GetEntryPointReturn UInt32

// NewGetEntryPointReturn factory
func NewGetEntryPointReturn() *GetEntryPointReturn {
	o := GetEntryPointReturn(*NewUInt32())
	return &o
}

// Serialize GetEntryPointReturn
func (n GetEntryPointReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := UInt32(n)
	return ox.Serialize(vb)
}

// DeserializeGetEntryPointReturn function
func DeserializeGetEntryPointReturn(vb *VariableBlob) (uint64,*GetEntryPointReturn,error) {
	var ot GetEntryPointReturn
	i,n,err := DeserializeUInt32(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = GetEntryPointReturn(*n)
	return i,&ot,nil}

// MarshalJSON GetEntryPointReturn
func (n GetEntryPointReturn) MarshalJSON() ([]byte, error) {
	v := UInt32(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetEntryPointReturn
func (n *GetEntryPointReturn) UnmarshalJSON(data []byte) error {
	v := UInt32(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetEntryPointReturn(v)
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
//  Typedef: GetContractArgsSizeReturn
// ----------------------------------------

// GetContractArgsSizeReturn type
type GetContractArgsSizeReturn UInt32

// NewGetContractArgsSizeReturn factory
func NewGetContractArgsSizeReturn() *GetContractArgsSizeReturn {
	o := GetContractArgsSizeReturn(*NewUInt32())
	return &o
}

// Serialize GetContractArgsSizeReturn
func (n GetContractArgsSizeReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := UInt32(n)
	return ox.Serialize(vb)
}

// DeserializeGetContractArgsSizeReturn function
func DeserializeGetContractArgsSizeReturn(vb *VariableBlob) (uint64,*GetContractArgsSizeReturn,error) {
	var ot GetContractArgsSizeReturn
	i,n,err := DeserializeUInt32(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = GetContractArgsSizeReturn(*n)
	return i,&ot,nil}

// MarshalJSON GetContractArgsSizeReturn
func (n GetContractArgsSizeReturn) MarshalJSON() ([]byte, error) {
	v := UInt32(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetContractArgsSizeReturn
func (n *GetContractArgsSizeReturn) UnmarshalJSON(data []byte) error {
	v := UInt32(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetContractArgsSizeReturn(v)
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
//  Typedef: GetContractArgsReturn
// ----------------------------------------

// GetContractArgsReturn type
type GetContractArgsReturn VariableBlob

// NewGetContractArgsReturn factory
func NewGetContractArgsReturn() *GetContractArgsReturn {
	o := GetContractArgsReturn(*NewVariableBlob())
	return &o
}

// Serialize GetContractArgsReturn
func (n GetContractArgsReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VariableBlob(n)
	return ox.Serialize(vb)
}

// DeserializeGetContractArgsReturn function
func DeserializeGetContractArgsReturn(vb *VariableBlob) (uint64,*GetContractArgsReturn,error) {
	var ot GetContractArgsReturn
	i,n,err := DeserializeVariableBlob(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = GetContractArgsReturn(*n)
	return i,&ot,nil}

// MarshalJSON GetContractArgsReturn
func (n GetContractArgsReturn) MarshalJSON() ([]byte, error) {
	v := VariableBlob(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetContractArgsReturn
func (n *GetContractArgsReturn) UnmarshalJSON(data []byte) error {
	v := VariableBlob(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetContractArgsReturn(v)
	return nil
}


// ----------------------------------------
//  Struct: SetContractReturnArgs
// ----------------------------------------

// SetContractReturnArgs type
type SetContractReturnArgs struct {
    Value VariableBlob `json:"value"`
}

// NewSetContractReturnArgs factory
func NewSetContractReturnArgs() *SetContractReturnArgs {
	o := SetContractReturnArgs{}
	o.Value = *NewVariableBlob()
	return &o
}

// Serialize SetContractReturnArgs
func (n SetContractReturnArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Value.Serialize(vb)
	return vb
}

// DeserializeSetContractReturnArgs function
func DeserializeSetContractReturnArgs(vb *VariableBlob) (uint64,*SetContractReturnArgs,error) {
	var i,j uint64 = 0,0
	s := SetContractReturnArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tValue,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &SetContractReturnArgs{}, err
	}
	s.Value = *tValue
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: SetContractReturnReturn
// ----------------------------------------

// SetContractReturnReturn type
type SetContractReturnReturn VoidType

// NewSetContractReturnReturn factory
func NewSetContractReturnReturn() *SetContractReturnReturn {
	o := SetContractReturnReturn(*NewVoidType())
	return &o
}

// Serialize SetContractReturnReturn
func (n SetContractReturnReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeSetContractReturnReturn function
func DeserializeSetContractReturnReturn(vb *VariableBlob) (uint64,*SetContractReturnReturn,error) {
	var ot SetContractReturnReturn
	return 0,&ot,nil}

// MarshalJSON SetContractReturnReturn
func (n SetContractReturnReturn) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *SetContractReturnReturn
func (n *SetContractReturnReturn) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = SetContractReturnReturn(v)
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
//  Typedef: ExitContractReturn
// ----------------------------------------

// ExitContractReturn type
type ExitContractReturn VoidType

// NewExitContractReturn factory
func NewExitContractReturn() *ExitContractReturn {
	o := ExitContractReturn(*NewVoidType())
	return &o
}

// Serialize ExitContractReturn
func (n ExitContractReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeExitContractReturn function
func DeserializeExitContractReturn(vb *VariableBlob) (uint64,*ExitContractReturn,error) {
	var ot ExitContractReturn
	return 0,&ot,nil}

// MarshalJSON ExitContractReturn
func (n ExitContractReturn) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *ExitContractReturn
func (n *ExitContractReturn) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = ExitContractReturn(v)
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
//  Typedef: GetHeadInfoReturn
// ----------------------------------------

// GetHeadInfoReturn type
type GetHeadInfoReturn HeadInfo

// NewGetHeadInfoReturn factory
func NewGetHeadInfoReturn() *GetHeadInfoReturn {
	o := GetHeadInfoReturn(*NewHeadInfo())
	return &o
}

// Serialize GetHeadInfoReturn
func (n GetHeadInfoReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := HeadInfo(n)
	return ox.Serialize(vb)
}

// DeserializeGetHeadInfoReturn function
func DeserializeGetHeadInfoReturn(vb *VariableBlob) (uint64,*GetHeadInfoReturn,error) {
	var ot GetHeadInfoReturn
	i,n,err := DeserializeHeadInfo(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = GetHeadInfoReturn(*n)
	return i,&ot,nil}

// MarshalJSON GetHeadInfoReturn
func (n GetHeadInfoReturn) MarshalJSON() ([]byte, error) {
	v := HeadInfo(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetHeadInfoReturn
func (n *GetHeadInfoReturn) UnmarshalJSON(data []byte) error {
	v := HeadInfo(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetHeadInfoReturn(v)
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
//  Typedef: HashReturn
// ----------------------------------------

// HashReturn type
type HashReturn Multihash

// NewHashReturn factory
func NewHashReturn() *HashReturn {
	o := HashReturn(*NewMultihash())
	return &o
}

// Serialize HashReturn
func (n HashReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := Multihash(n)
	return ox.Serialize(vb)
}

// DeserializeHashReturn function
func DeserializeHashReturn(vb *VariableBlob) (uint64,*HashReturn,error) {
	var ot HashReturn
	i,n,err := DeserializeMultihash(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = HashReturn(*n)
	return i,&ot,nil}

// MarshalJSON HashReturn
func (n HashReturn) MarshalJSON() ([]byte, error) {
	v := Multihash(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *HashReturn
func (n *HashReturn) UnmarshalJSON(data []byte) error {
	v := Multihash(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = HashReturn(v)
	return nil
}


// ----------------------------------------
//  Struct: GetTransactionPayerArgs
// ----------------------------------------

// GetTransactionPayerArgs type
type GetTransactionPayerArgs struct {
    Transaction Transaction `json:"transaction"`
}

// NewGetTransactionPayerArgs factory
func NewGetTransactionPayerArgs() *GetTransactionPayerArgs {
	o := GetTransactionPayerArgs{}
	o.Transaction = *NewTransaction()
	return &o
}

// Serialize GetTransactionPayerArgs
func (n GetTransactionPayerArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Transaction.Serialize(vb)
	return vb
}

// DeserializeGetTransactionPayerArgs function
func DeserializeGetTransactionPayerArgs(vb *VariableBlob) (uint64,*GetTransactionPayerArgs,error) {
	var i,j uint64 = 0,0
	s := GetTransactionPayerArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransaction,err := DeserializeTransaction(&ovb); i+=j
	if err != nil {
		return 0, &GetTransactionPayerArgs{}, err
	}
	s.Transaction = *tTransaction
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: GetTransactionPayerReturn
// ----------------------------------------

// GetTransactionPayerReturn type
type GetTransactionPayerReturn AccountType

// NewGetTransactionPayerReturn factory
func NewGetTransactionPayerReturn() *GetTransactionPayerReturn {
	o := GetTransactionPayerReturn(*NewAccountType())
	return &o
}

// Serialize GetTransactionPayerReturn
func (n GetTransactionPayerReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := AccountType(n)
	return ox.Serialize(vb)
}

// DeserializeGetTransactionPayerReturn function
func DeserializeGetTransactionPayerReturn(vb *VariableBlob) (uint64,*GetTransactionPayerReturn,error) {
	var ot GetTransactionPayerReturn
	i,n,err := DeserializeAccountType(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = GetTransactionPayerReturn(*n)
	return i,&ot,nil}

// MarshalJSON GetTransactionPayerReturn
func (n GetTransactionPayerReturn) MarshalJSON() ([]byte, error) {
	v := AccountType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetTransactionPayerReturn
func (n *GetTransactionPayerReturn) UnmarshalJSON(data []byte) error {
	v := AccountType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetTransactionPayerReturn(v)
	return nil
}


// ----------------------------------------
//  Struct: GetMaxAccountResourcesArgs
// ----------------------------------------

// GetMaxAccountResourcesArgs type
type GetMaxAccountResourcesArgs struct {
    Account AccountType `json:"account"`
}

// NewGetMaxAccountResourcesArgs factory
func NewGetMaxAccountResourcesArgs() *GetMaxAccountResourcesArgs {
	o := GetMaxAccountResourcesArgs{}
	o.Account = *NewAccountType()
	return &o
}

// Serialize GetMaxAccountResourcesArgs
func (n GetMaxAccountResourcesArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Account.Serialize(vb)
	return vb
}

// DeserializeGetMaxAccountResourcesArgs function
func DeserializeGetMaxAccountResourcesArgs(vb *VariableBlob) (uint64,*GetMaxAccountResourcesArgs,error) {
	var i,j uint64 = 0,0
	s := GetMaxAccountResourcesArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tAccount,err := DeserializeAccountType(&ovb); i+=j
	if err != nil {
		return 0, &GetMaxAccountResourcesArgs{}, err
	}
	s.Account = *tAccount
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: GetMaxAccountResourcesReturn
// ----------------------------------------

// GetMaxAccountResourcesReturn type
type GetMaxAccountResourcesReturn UInt128

// NewGetMaxAccountResourcesReturn factory
func NewGetMaxAccountResourcesReturn() *GetMaxAccountResourcesReturn {
	o := GetMaxAccountResourcesReturn(*NewUInt128())
	return &o
}

// Serialize GetMaxAccountResourcesReturn
func (n GetMaxAccountResourcesReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := UInt128(n)
	return ox.Serialize(vb)
}

// DeserializeGetMaxAccountResourcesReturn function
func DeserializeGetMaxAccountResourcesReturn(vb *VariableBlob) (uint64,*GetMaxAccountResourcesReturn,error) {
	var ot GetMaxAccountResourcesReturn
	i,n,err := DeserializeUInt128(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = GetMaxAccountResourcesReturn(*n)
	return i,&ot,nil}

// MarshalJSON GetMaxAccountResourcesReturn
func (n GetMaxAccountResourcesReturn) MarshalJSON() ([]byte, error) {
	v := UInt128(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetMaxAccountResourcesReturn
func (n *GetMaxAccountResourcesReturn) UnmarshalJSON(data []byte) error {
	v := UInt128(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetMaxAccountResourcesReturn(v)
	return nil
}


// ----------------------------------------
//  Struct: GetTransactionResourceLimitArgs
// ----------------------------------------

// GetTransactionResourceLimitArgs type
type GetTransactionResourceLimitArgs struct {
    Transaction Transaction `json:"transaction"`
}

// NewGetTransactionResourceLimitArgs factory
func NewGetTransactionResourceLimitArgs() *GetTransactionResourceLimitArgs {
	o := GetTransactionResourceLimitArgs{}
	o.Transaction = *NewTransaction()
	return &o
}

// Serialize GetTransactionResourceLimitArgs
func (n GetTransactionResourceLimitArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Transaction.Serialize(vb)
	return vb
}

// DeserializeGetTransactionResourceLimitArgs function
func DeserializeGetTransactionResourceLimitArgs(vb *VariableBlob) (uint64,*GetTransactionResourceLimitArgs,error) {
	var i,j uint64 = 0,0
	s := GetTransactionResourceLimitArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransaction,err := DeserializeTransaction(&ovb); i+=j
	if err != nil {
		return 0, &GetTransactionResourceLimitArgs{}, err
	}
	s.Transaction = *tTransaction
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: GetTransactionResourceLimitReturn
// ----------------------------------------

// GetTransactionResourceLimitReturn type
type GetTransactionResourceLimitReturn UInt128

// NewGetTransactionResourceLimitReturn factory
func NewGetTransactionResourceLimitReturn() *GetTransactionResourceLimitReturn {
	o := GetTransactionResourceLimitReturn(*NewUInt128())
	return &o
}

// Serialize GetTransactionResourceLimitReturn
func (n GetTransactionResourceLimitReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := UInt128(n)
	return ox.Serialize(vb)
}

// DeserializeGetTransactionResourceLimitReturn function
func DeserializeGetTransactionResourceLimitReturn(vb *VariableBlob) (uint64,*GetTransactionResourceLimitReturn,error) {
	var ot GetTransactionResourceLimitReturn
	i,n,err := DeserializeUInt128(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = GetTransactionResourceLimitReturn(*n)
	return i,&ot,nil}

// MarshalJSON GetTransactionResourceLimitReturn
func (n GetTransactionResourceLimitReturn) MarshalJSON() ([]byte, error) {
	v := UInt128(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetTransactionResourceLimitReturn
func (n *GetTransactionResourceLimitReturn) UnmarshalJSON(data []byte) error {
	v := UInt128(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetTransactionResourceLimitReturn(v)
	return nil
}


// ----------------------------------------
//  Struct: GetLastIrreversibleBlockArgs
// ----------------------------------------

// GetLastIrreversibleBlockArgs type
type GetLastIrreversibleBlockArgs struct {
}

// NewGetLastIrreversibleBlockArgs factory
func NewGetLastIrreversibleBlockArgs() *GetLastIrreversibleBlockArgs {
	o := GetLastIrreversibleBlockArgs{}
	return &o
}

// Serialize GetLastIrreversibleBlockArgs
func (n GetLastIrreversibleBlockArgs) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeGetLastIrreversibleBlockArgs function
func DeserializeGetLastIrreversibleBlockArgs(vb *VariableBlob) (uint64,*GetLastIrreversibleBlockArgs,error) {
	var i uint64 = 0
	s := GetLastIrreversibleBlockArgs{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: GetLastIrreversibleBlockReturn
// ----------------------------------------

// GetLastIrreversibleBlockReturn type
type GetLastIrreversibleBlockReturn BlockHeightType

// NewGetLastIrreversibleBlockReturn factory
func NewGetLastIrreversibleBlockReturn() *GetLastIrreversibleBlockReturn {
	o := GetLastIrreversibleBlockReturn(*NewBlockHeightType())
	return &o
}

// Serialize GetLastIrreversibleBlockReturn
func (n GetLastIrreversibleBlockReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := BlockHeightType(n)
	return ox.Serialize(vb)
}

// DeserializeGetLastIrreversibleBlockReturn function
func DeserializeGetLastIrreversibleBlockReturn(vb *VariableBlob) (uint64,*GetLastIrreversibleBlockReturn,error) {
	var ot GetLastIrreversibleBlockReturn
	i,n,err := DeserializeBlockHeightType(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = GetLastIrreversibleBlockReturn(*n)
	return i,&ot,nil}

// MarshalJSON GetLastIrreversibleBlockReturn
func (n GetLastIrreversibleBlockReturn) MarshalJSON() ([]byte, error) {
	v := BlockHeightType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetLastIrreversibleBlockReturn
func (n *GetLastIrreversibleBlockReturn) UnmarshalJSON(data []byte) error {
	v := BlockHeightType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetLastIrreversibleBlockReturn(v)
	return nil
}


// ----------------------------------------
//  Struct: GetCallerArgs
// ----------------------------------------

// GetCallerArgs type
type GetCallerArgs struct {
}

// NewGetCallerArgs factory
func NewGetCallerArgs() *GetCallerArgs {
	o := GetCallerArgs{}
	return &o
}

// Serialize GetCallerArgs
func (n GetCallerArgs) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeGetCallerArgs function
func DeserializeGetCallerArgs(vb *VariableBlob) (uint64,*GetCallerArgs,error) {
	var i uint64 = 0
	s := GetCallerArgs{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: GetCallerReturn
// ----------------------------------------

// GetCallerReturn type
type GetCallerReturn AccountType

// NewGetCallerReturn factory
func NewGetCallerReturn() *GetCallerReturn {
	o := GetCallerReturn(*NewAccountType())
	return &o
}

// Serialize GetCallerReturn
func (n GetCallerReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := AccountType(n)
	return ox.Serialize(vb)
}

// DeserializeGetCallerReturn function
func DeserializeGetCallerReturn(vb *VariableBlob) (uint64,*GetCallerReturn,error) {
	var ot GetCallerReturn
	i,n,err := DeserializeAccountType(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = GetCallerReturn(*n)
	return i,&ot,nil}

// MarshalJSON GetCallerReturn
func (n GetCallerReturn) MarshalJSON() ([]byte, error) {
	v := AccountType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetCallerReturn
func (n *GetCallerReturn) UnmarshalJSON(data []byte) error {
	v := AccountType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetCallerReturn(v)
	return nil
}


// ----------------------------------------
//  Struct: RequireAuthorityArgs
// ----------------------------------------

// RequireAuthorityArgs type
type RequireAuthorityArgs struct {
    Account AccountType `json:"account"`
}

// NewRequireAuthorityArgs factory
func NewRequireAuthorityArgs() *RequireAuthorityArgs {
	o := RequireAuthorityArgs{}
	o.Account = *NewAccountType()
	return &o
}

// Serialize RequireAuthorityArgs
func (n RequireAuthorityArgs) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Account.Serialize(vb)
	return vb
}

// DeserializeRequireAuthorityArgs function
func DeserializeRequireAuthorityArgs(vb *VariableBlob) (uint64,*RequireAuthorityArgs,error) {
	var i,j uint64 = 0,0
	s := RequireAuthorityArgs{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tAccount,err := DeserializeAccountType(&ovb); i+=j
	if err != nil {
		return 0, &RequireAuthorityArgs{}, err
	}
	s.Account = *tAccount
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: RequireAuthorityReturn
// ----------------------------------------

// RequireAuthorityReturn type
type RequireAuthorityReturn VoidType

// NewRequireAuthorityReturn factory
func NewRequireAuthorityReturn() *RequireAuthorityReturn {
	o := RequireAuthorityReturn(*NewVoidType())
	return &o
}

// Serialize RequireAuthorityReturn
func (n RequireAuthorityReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VoidType(n)
	return ox.Serialize(vb)
}

// DeserializeRequireAuthorityReturn function
func DeserializeRequireAuthorityReturn(vb *VariableBlob) (uint64,*RequireAuthorityReturn,error) {
	var ot RequireAuthorityReturn
	return 0,&ot,nil}

// MarshalJSON RequireAuthorityReturn
func (n RequireAuthorityReturn) MarshalJSON() ([]byte, error) {
	v := VoidType(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *RequireAuthorityReturn
func (n *RequireAuthorityReturn) UnmarshalJSON(data []byte) error {
	v := VoidType(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = RequireAuthorityReturn(v)
	return nil
}


// ----------------------------------------
//  Struct: GetTransactionSignatureArgs
// ----------------------------------------

// GetTransactionSignatureArgs type
type GetTransactionSignatureArgs struct {
}

// NewGetTransactionSignatureArgs factory
func NewGetTransactionSignatureArgs() *GetTransactionSignatureArgs {
	o := GetTransactionSignatureArgs{}
	return &o
}

// Serialize GetTransactionSignatureArgs
func (n GetTransactionSignatureArgs) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeGetTransactionSignatureArgs function
func DeserializeGetTransactionSignatureArgs(vb *VariableBlob) (uint64,*GetTransactionSignatureArgs,error) {
	var i uint64 = 0
	s := GetTransactionSignatureArgs{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Typedef: GetTransactionSignatureReturn
// ----------------------------------------

// GetTransactionSignatureReturn type
type GetTransactionSignatureReturn VariableBlob

// NewGetTransactionSignatureReturn factory
func NewGetTransactionSignatureReturn() *GetTransactionSignatureReturn {
	o := GetTransactionSignatureReturn(*NewVariableBlob())
	return &o
}

// Serialize GetTransactionSignatureReturn
func (n GetTransactionSignatureReturn) Serialize(vb *VariableBlob) *VariableBlob {
	ox := VariableBlob(n)
	return ox.Serialize(vb)
}

// DeserializeGetTransactionSignatureReturn function
func DeserializeGetTransactionSignatureReturn(vb *VariableBlob) (uint64,*GetTransactionSignatureReturn,error) {
	var ot GetTransactionSignatureReturn
	i,n,err := DeserializeVariableBlob(vb)
	if err != nil {
		return 0,&ot,err
	}
	ot = GetTransactionSignatureReturn(*n)
	return i,&ot,nil}

// MarshalJSON GetTransactionSignatureReturn
func (n GetTransactionSignatureReturn) MarshalJSON() ([]byte, error) {
	v := VariableBlob(n)
	return json.Marshal(&v)
}

// UnmarshalJSON *GetTransactionSignatureReturn
func (n *GetTransactionSignatureReturn) UnmarshalJSON(data []byte) error {
	v := VariableBlob(*n);
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*n = GetTransactionSignatureReturn(v)
	return nil
}


// ----------------------------------------
//  Struct: ChainReservedRequest
// ----------------------------------------

// ChainReservedRequest type
type ChainReservedRequest struct {
}

// NewChainReservedRequest factory
func NewChainReservedRequest() *ChainReservedRequest {
	o := ChainReservedRequest{}
	return &o
}

// Serialize ChainReservedRequest
func (n ChainReservedRequest) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeChainReservedRequest function
func DeserializeChainReservedRequest(vb *VariableBlob) (uint64,*ChainReservedRequest,error) {
	var i uint64 = 0
	s := ChainReservedRequest{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: SubmitBlockRequest
// ----------------------------------------

// SubmitBlockRequest type
type SubmitBlockRequest struct {
    Block Block `json:"block"`
    VerifyPassiveData Boolean `json:"verify_passive_data"`
    VerifyBlockSignature Boolean `json:"verify_block_signature"`
    VerifyTransactionSignatures Boolean `json:"verify_transaction_signatures"`
}

// NewSubmitBlockRequest factory
func NewSubmitBlockRequest() *SubmitBlockRequest {
	o := SubmitBlockRequest{}
	o.Block = *NewBlock()
	o.VerifyPassiveData = *NewBoolean()
	o.VerifyBlockSignature = *NewBoolean()
	o.VerifyTransactionSignatures = *NewBoolean()
	return &o
}

// Serialize SubmitBlockRequest
func (n SubmitBlockRequest) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Block.Serialize(vb)
	vb = n.VerifyPassiveData.Serialize(vb)
	vb = n.VerifyBlockSignature.Serialize(vb)
	vb = n.VerifyTransactionSignatures.Serialize(vb)
	return vb
}

// DeserializeSubmitBlockRequest function
func DeserializeSubmitBlockRequest(vb *VariableBlob) (uint64,*SubmitBlockRequest,error) {
	var i,j uint64 = 0,0
	s := SubmitBlockRequest{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tBlock,err := DeserializeBlock(&ovb); i+=j
	if err != nil {
		return 0, &SubmitBlockRequest{}, err
	}
	s.Block = *tBlock
	ovb = (*vb)[i:]
	j,tVerifyPassiveData,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &SubmitBlockRequest{}, err
	}
	s.VerifyPassiveData = *tVerifyPassiveData
	ovb = (*vb)[i:]
	j,tVerifyBlockSignature,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &SubmitBlockRequest{}, err
	}
	s.VerifyBlockSignature = *tVerifyBlockSignature
	ovb = (*vb)[i:]
	j,tVerifyTransactionSignatures,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &SubmitBlockRequest{}, err
	}
	s.VerifyTransactionSignatures = *tVerifyTransactionSignatures
	return i, &s, nil
}

// ----------------------------------------
//  Struct: SubmitTransactionRequest
// ----------------------------------------

// SubmitTransactionRequest type
type SubmitTransactionRequest struct {
    Transaction Transaction `json:"transaction"`
    VerifyPassiveData Boolean `json:"verify_passive_data"`
    VerifyTransactionSignatures Boolean `json:"verify_transaction_signatures"`
}

// NewSubmitTransactionRequest factory
func NewSubmitTransactionRequest() *SubmitTransactionRequest {
	o := SubmitTransactionRequest{}
	o.Transaction = *NewTransaction()
	o.VerifyPassiveData = *NewBoolean()
	o.VerifyTransactionSignatures = *NewBoolean()
	return &o
}

// Serialize SubmitTransactionRequest
func (n SubmitTransactionRequest) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Transaction.Serialize(vb)
	vb = n.VerifyPassiveData.Serialize(vb)
	vb = n.VerifyTransactionSignatures.Serialize(vb)
	return vb
}

// DeserializeSubmitTransactionRequest function
func DeserializeSubmitTransactionRequest(vb *VariableBlob) (uint64,*SubmitTransactionRequest,error) {
	var i,j uint64 = 0,0
	s := SubmitTransactionRequest{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransaction,err := DeserializeTransaction(&ovb); i+=j
	if err != nil {
		return 0, &SubmitTransactionRequest{}, err
	}
	s.Transaction = *tTransaction
	ovb = (*vb)[i:]
	j,tVerifyPassiveData,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &SubmitTransactionRequest{}, err
	}
	s.VerifyPassiveData = *tVerifyPassiveData
	ovb = (*vb)[i:]
	j,tVerifyTransactionSignatures,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &SubmitTransactionRequest{}, err
	}
	s.VerifyTransactionSignatures = *tVerifyTransactionSignatures
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetHeadInfoRequest
// ----------------------------------------

// GetHeadInfoRequest type
type GetHeadInfoRequest struct {
}

// NewGetHeadInfoRequest factory
func NewGetHeadInfoRequest() *GetHeadInfoRequest {
	o := GetHeadInfoRequest{}
	return &o
}

// Serialize GetHeadInfoRequest
func (n GetHeadInfoRequest) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeGetHeadInfoRequest function
func DeserializeGetHeadInfoRequest(vb *VariableBlob) (uint64,*GetHeadInfoRequest,error) {
	var i uint64 = 0
	s := GetHeadInfoRequest{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetChainIDRequest
// ----------------------------------------

// GetChainIDRequest type
type GetChainIDRequest struct {
}

// NewGetChainIDRequest factory
func NewGetChainIDRequest() *GetChainIDRequest {
	o := GetChainIDRequest{}
	return &o
}

// Serialize GetChainIDRequest
func (n GetChainIDRequest) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeGetChainIDRequest function
func DeserializeGetChainIDRequest(vb *VariableBlob) (uint64,*GetChainIDRequest,error) {
	var i uint64 = 0
	s := GetChainIDRequest{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetForkHeadsRequest
// ----------------------------------------

// GetForkHeadsRequest type
type GetForkHeadsRequest struct {
}

// NewGetForkHeadsRequest factory
func NewGetForkHeadsRequest() *GetForkHeadsRequest {
	o := GetForkHeadsRequest{}
	return &o
}

// Serialize GetForkHeadsRequest
func (n GetForkHeadsRequest) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeGetForkHeadsRequest function
func DeserializeGetForkHeadsRequest(vb *VariableBlob) (uint64,*GetForkHeadsRequest,error) {
	var i uint64 = 0
	s := GetForkHeadsRequest{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: ReadContractRequest
// ----------------------------------------

// ReadContractRequest type
type ReadContractRequest struct {
    ContractID ContractIDType `json:"contract_id"`
    EntryPoint UInt32 `json:"entry_point"`
    Args VariableBlob `json:"args"`
}

// NewReadContractRequest factory
func NewReadContractRequest() *ReadContractRequest {
	o := ReadContractRequest{}
	o.ContractID = *NewContractIDType()
	o.EntryPoint = *NewUInt32()
	o.Args = *NewVariableBlob()
	return &o
}

// Serialize ReadContractRequest
func (n ReadContractRequest) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ContractID.Serialize(vb)
	vb = n.EntryPoint.Serialize(vb)
	vb = n.Args.Serialize(vb)
	return vb
}

// DeserializeReadContractRequest function
func DeserializeReadContractRequest(vb *VariableBlob) (uint64,*ReadContractRequest,error) {
	var i,j uint64 = 0,0
	s := ReadContractRequest{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tContractID,err := DeserializeContractIDType(&ovb); i+=j
	if err != nil {
		return 0, &ReadContractRequest{}, err
	}
	s.ContractID = *tContractID
	ovb = (*vb)[i:]
	j,tEntryPoint,err := DeserializeUInt32(&ovb); i+=j
	if err != nil {
		return 0, &ReadContractRequest{}, err
	}
	s.EntryPoint = *tEntryPoint
	ovb = (*vb)[i:]
	j,tArgs,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &ReadContractRequest{}, err
	}
	s.Args = *tArgs
	return i, &s, nil
}

// ----------------------------------------
//  Variant: ChainRPCRequest
// ----------------------------------------

// ChainRPCRequest type
type ChainRPCRequest struct {
	Value interface{}
}

// NewChainRPCRequest factory
func NewChainRPCRequest() *ChainRPCRequest {
	v := ChainRPCRequest{}
	v.Value = NewChainReservedRequest()
	return &v
}

// Serialize ChainRPCRequest
func (n ChainRPCRequest) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *ChainReservedRequest:
			i = 0
		case *SubmitBlockRequest:
			i = 1
		case *SubmitTransactionRequest:
			i = 2
		case *GetHeadInfoRequest:
			i = 3
		case *GetChainIDRequest:
			i = 4
		case *GetForkHeadsRequest:
			i = 5
		case *ReadContractRequest:
			i = 6
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName ChainRPCRequest
func (n ChainRPCRequest) TypeToName() (string) {
	switch n.Value.(type) {
		case *ChainReservedRequest:
			return "koinos::rpc::chain::chain_reserved_request"
		case *SubmitBlockRequest:
			return "koinos::rpc::chain::submit_block_request"
		case *SubmitTransactionRequest:
			return "koinos::rpc::chain::submit_transaction_request"
		case *GetHeadInfoRequest:
			return "koinos::rpc::chain::get_head_info_request"
		case *GetChainIDRequest:
			return "koinos::rpc::chain::get_chain_id_request"
		case *GetForkHeadsRequest:
			return "koinos::rpc::chain::get_fork_heads_request"
		case *ReadContractRequest:
			return "koinos::rpc::chain::read_contract_request"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON ChainRPCRequest
func (n ChainRPCRequest) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeChainRPCRequest function
func DeserializeChainRPCRequest(vb *VariableBlob) (uint64,*ChainRPCRequest,error) {
	var v ChainRPCRequest
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewChainReservedRequest()
		case 1:
			ovb := (*vb)[i:]
			k,x,err := DeserializeSubmitBlockRequest(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 2:
			ovb := (*vb)[i:]
			k,x,err := DeserializeSubmitTransactionRequest(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 3:
			v.Value = NewGetHeadInfoRequest()
		case 4:
			v.Value = NewGetChainIDRequest()
		case 5:
			v.Value = NewGetForkHeadsRequest()
		case 6:
			ovb := (*vb)[i:]
			k,x,err := DeserializeReadContractRequest(&ovb)
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

// UnmarshalJSON *ChainRPCRequest
func (n *ChainRPCRequest) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::rpc::chain::chain_reserved_request":
			v := NewChainReservedRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::submit_block_request":
			v := NewSubmitBlockRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::submit_transaction_request":
			v := NewSubmitTransactionRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::get_head_info_request":
			v := NewGetHeadInfoRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::get_chain_id_request":
			v := NewGetChainIDRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::get_fork_heads_request":
			v := NewGetForkHeadsRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::read_contract_request":
			v := NewReadContractRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

	return nil
}


// ----------------------------------------
//  Struct: ChainReservedResponse
// ----------------------------------------

// ChainReservedResponse type
type ChainReservedResponse struct {
}

// NewChainReservedResponse factory
func NewChainReservedResponse() *ChainReservedResponse {
	o := ChainReservedResponse{}
	return &o
}

// Serialize ChainReservedResponse
func (n ChainReservedResponse) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeChainReservedResponse function
func DeserializeChainReservedResponse(vb *VariableBlob) (uint64,*ChainReservedResponse,error) {
	var i uint64 = 0
	s := ChainReservedResponse{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: ChainErrorResponse
// ----------------------------------------

// ChainErrorResponse type
type ChainErrorResponse struct {
    ErrorText String `json:"error_text"`
    ErrorData String `json:"error_data"`
}

// NewChainErrorResponse factory
func NewChainErrorResponse() *ChainErrorResponse {
	o := ChainErrorResponse{}
	o.ErrorText = *NewString()
	o.ErrorData = *NewString()
	return &o
}

// Serialize ChainErrorResponse
func (n ChainErrorResponse) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ErrorText.Serialize(vb)
	vb = n.ErrorData.Serialize(vb)
	return vb
}

// DeserializeChainErrorResponse function
func DeserializeChainErrorResponse(vb *VariableBlob) (uint64,*ChainErrorResponse,error) {
	var i,j uint64 = 0,0
	s := ChainErrorResponse{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tErrorText,err := DeserializeString(&ovb); i+=j
	if err != nil {
		return 0, &ChainErrorResponse{}, err
	}
	s.ErrorText = *tErrorText
	ovb = (*vb)[i:]
	j,tErrorData,err := DeserializeString(&ovb); i+=j
	if err != nil {
		return 0, &ChainErrorResponse{}, err
	}
	s.ErrorData = *tErrorData
	return i, &s, nil
}

// ----------------------------------------
//  Struct: SubmitBlockResponse
// ----------------------------------------

// SubmitBlockResponse type
type SubmitBlockResponse struct {
}

// NewSubmitBlockResponse factory
func NewSubmitBlockResponse() *SubmitBlockResponse {
	o := SubmitBlockResponse{}
	return &o
}

// Serialize SubmitBlockResponse
func (n SubmitBlockResponse) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeSubmitBlockResponse function
func DeserializeSubmitBlockResponse(vb *VariableBlob) (uint64,*SubmitBlockResponse,error) {
	var i uint64 = 0
	s := SubmitBlockResponse{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: SubmitTransactionResponse
// ----------------------------------------

// SubmitTransactionResponse type
type SubmitTransactionResponse struct {
}

// NewSubmitTransactionResponse factory
func NewSubmitTransactionResponse() *SubmitTransactionResponse {
	o := SubmitTransactionResponse{}
	return &o
}

// Serialize SubmitTransactionResponse
func (n SubmitTransactionResponse) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeSubmitTransactionResponse function
func DeserializeSubmitTransactionResponse(vb *VariableBlob) (uint64,*SubmitTransactionResponse,error) {
	var i uint64 = 0
	s := SubmitTransactionResponse{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetHeadInfoResponse
// ----------------------------------------

// GetHeadInfoResponse type
type GetHeadInfoResponse struct {
    HeadTopology BlockTopology `json:"head_topology"`
    LastIrreversibleHeight BlockHeightType `json:"last_irreversible_height"`
}

// NewGetHeadInfoResponse factory
func NewGetHeadInfoResponse() *GetHeadInfoResponse {
	o := GetHeadInfoResponse{}
	o.HeadTopology = *NewBlockTopology()
	o.LastIrreversibleHeight = *NewBlockHeightType()
	return &o
}

// Serialize GetHeadInfoResponse
func (n GetHeadInfoResponse) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.HeadTopology.Serialize(vb)
	vb = n.LastIrreversibleHeight.Serialize(vb)
	return vb
}

// DeserializeGetHeadInfoResponse function
func DeserializeGetHeadInfoResponse(vb *VariableBlob) (uint64,*GetHeadInfoResponse,error) {
	var i,j uint64 = 0,0
	s := GetHeadInfoResponse{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tHeadTopology,err := DeserializeBlockTopology(&ovb); i+=j
	if err != nil {
		return 0, &GetHeadInfoResponse{}, err
	}
	s.HeadTopology = *tHeadTopology
	ovb = (*vb)[i:]
	j,tLastIrreversibleHeight,err := DeserializeBlockHeightType(&ovb); i+=j
	if err != nil {
		return 0, &GetHeadInfoResponse{}, err
	}
	s.LastIrreversibleHeight = *tLastIrreversibleHeight
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetChainIDResponse
// ----------------------------------------

// GetChainIDResponse type
type GetChainIDResponse struct {
    ChainID Multihash `json:"chain_id"`
}

// NewGetChainIDResponse factory
func NewGetChainIDResponse() *GetChainIDResponse {
	o := GetChainIDResponse{}
	o.ChainID = *NewMultihash()
	return &o
}

// Serialize GetChainIDResponse
func (n GetChainIDResponse) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ChainID.Serialize(vb)
	return vb
}

// DeserializeGetChainIDResponse function
func DeserializeGetChainIDResponse(vb *VariableBlob) (uint64,*GetChainIDResponse,error) {
	var i,j uint64 = 0,0
	s := GetChainIDResponse{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tChainID,err := DeserializeMultihash(&ovb); i+=j
	if err != nil {
		return 0, &GetChainIDResponse{}, err
	}
	s.ChainID = *tChainID
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetForkHeadsResponse
// ----------------------------------------

// GetForkHeadsResponse type
type GetForkHeadsResponse struct {
    ForkHeads VectorBlockTopology `json:"fork_heads"`
    LastIrreversibleBlock BlockTopology `json:"last_irreversible_block"`
}

// NewGetForkHeadsResponse factory
func NewGetForkHeadsResponse() *GetForkHeadsResponse {
	o := GetForkHeadsResponse{}
	o.ForkHeads = *NewVectorBlockTopology()
	o.LastIrreversibleBlock = *NewBlockTopology()
	return &o
}

// Serialize GetForkHeadsResponse
func (n GetForkHeadsResponse) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ForkHeads.Serialize(vb)
	vb = n.LastIrreversibleBlock.Serialize(vb)
	return vb
}

// DeserializeGetForkHeadsResponse function
func DeserializeGetForkHeadsResponse(vb *VariableBlob) (uint64,*GetForkHeadsResponse,error) {
	var i,j uint64 = 0,0
	s := GetForkHeadsResponse{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tForkHeads,err := DeserializeVectorBlockTopology(&ovb); i+=j
	if err != nil {
		return 0, &GetForkHeadsResponse{}, err
	}
	s.ForkHeads = *tForkHeads
	ovb = (*vb)[i:]
	j,tLastIrreversibleBlock,err := DeserializeBlockTopology(&ovb); i+=j
	if err != nil {
		return 0, &GetForkHeadsResponse{}, err
	}
	s.LastIrreversibleBlock = *tLastIrreversibleBlock
	return i, &s, nil
}

// ----------------------------------------
//  Struct: ReadContractResponse
// ----------------------------------------

// ReadContractResponse type
type ReadContractResponse struct {
    Result VariableBlob `json:"result"`
    Logs String `json:"logs"`
}

// NewReadContractResponse factory
func NewReadContractResponse() *ReadContractResponse {
	o := ReadContractResponse{}
	o.Result = *NewVariableBlob()
	o.Logs = *NewString()
	return &o
}

// Serialize ReadContractResponse
func (n ReadContractResponse) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Result.Serialize(vb)
	vb = n.Logs.Serialize(vb)
	return vb
}

// DeserializeReadContractResponse function
func DeserializeReadContractResponse(vb *VariableBlob) (uint64,*ReadContractResponse,error) {
	var i,j uint64 = 0,0
	s := ReadContractResponse{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tResult,err := DeserializeVariableBlob(&ovb); i+=j
	if err != nil {
		return 0, &ReadContractResponse{}, err
	}
	s.Result = *tResult
	ovb = (*vb)[i:]
	j,tLogs,err := DeserializeString(&ovb); i+=j
	if err != nil {
		return 0, &ReadContractResponse{}, err
	}
	s.Logs = *tLogs
	return i, &s, nil
}

// ----------------------------------------
//  Variant: ChainRPCResponse
// ----------------------------------------

// ChainRPCResponse type
type ChainRPCResponse struct {
	Value interface{}
}

// NewChainRPCResponse factory
func NewChainRPCResponse() *ChainRPCResponse {
	v := ChainRPCResponse{}
	v.Value = NewChainReservedResponse()
	return &v
}

// Serialize ChainRPCResponse
func (n ChainRPCResponse) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *ChainReservedResponse:
			i = 0
		case *ChainErrorResponse:
			i = 1
		case *SubmitBlockResponse:
			i = 2
		case *SubmitTransactionResponse:
			i = 3
		case *GetHeadInfoResponse:
			i = 4
		case *GetChainIDResponse:
			i = 5
		case *GetForkHeadsResponse:
			i = 6
		case *ReadContractResponse:
			i = 7
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName ChainRPCResponse
func (n ChainRPCResponse) TypeToName() (string) {
	switch n.Value.(type) {
		case *ChainReservedResponse:
			return "koinos::rpc::chain::chain_reserved_response"
		case *ChainErrorResponse:
			return "koinos::rpc::chain::chain_error_response"
		case *SubmitBlockResponse:
			return "koinos::rpc::chain::submit_block_response"
		case *SubmitTransactionResponse:
			return "koinos::rpc::chain::submit_transaction_response"
		case *GetHeadInfoResponse:
			return "koinos::rpc::chain::get_head_info_response"
		case *GetChainIDResponse:
			return "koinos::rpc::chain::get_chain_id_response"
		case *GetForkHeadsResponse:
			return "koinos::rpc::chain::get_fork_heads_response"
		case *ReadContractResponse:
			return "koinos::rpc::chain::read_contract_response"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON ChainRPCResponse
func (n ChainRPCResponse) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeChainRPCResponse function
func DeserializeChainRPCResponse(vb *VariableBlob) (uint64,*ChainRPCResponse,error) {
	var v ChainRPCResponse
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewChainReservedResponse()
		case 1:
			ovb := (*vb)[i:]
			k,x,err := DeserializeChainErrorResponse(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 2:
			v.Value = NewSubmitBlockResponse()
		case 3:
			v.Value = NewSubmitTransactionResponse()
		case 4:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetHeadInfoResponse(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 5:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetChainIDResponse(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 6:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetForkHeadsResponse(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 7:
			ovb := (*vb)[i:]
			k,x,err := DeserializeReadContractResponse(&ovb)
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

// UnmarshalJSON *ChainRPCResponse
func (n *ChainRPCResponse) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::rpc::chain::chain_reserved_response":
			v := NewChainReservedResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::chain_error_response":
			v := NewChainErrorResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::submit_block_response":
			v := NewSubmitBlockResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::submit_transaction_response":
			v := NewSubmitTransactionResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::get_head_info_response":
			v := NewGetHeadInfoResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::get_chain_id_response":
			v := NewGetChainIDResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::get_fork_heads_response":
			v := NewGetForkHeadsResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::chain::read_contract_response":
			v := NewReadContractResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

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
//  Struct: MempoolReservedRequest
// ----------------------------------------

// MempoolReservedRequest type
type MempoolReservedRequest struct {
}

// NewMempoolReservedRequest factory
func NewMempoolReservedRequest() *MempoolReservedRequest {
	o := MempoolReservedRequest{}
	return &o
}

// Serialize MempoolReservedRequest
func (n MempoolReservedRequest) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeMempoolReservedRequest function
func DeserializeMempoolReservedRequest(vb *VariableBlob) (uint64,*MempoolReservedRequest,error) {
	var i uint64 = 0
	s := MempoolReservedRequest{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: CheckPendingAccountResourcesRequest
// ----------------------------------------

// CheckPendingAccountResourcesRequest type
type CheckPendingAccountResourcesRequest struct {
    Payer AccountType `json:"payer"`
    MaxPayerResources UInt128 `json:"max_payer_resources"`
    TrxResourceLimit UInt128 `json:"trx_resource_limit"`
}

// NewCheckPendingAccountResourcesRequest factory
func NewCheckPendingAccountResourcesRequest() *CheckPendingAccountResourcesRequest {
	o := CheckPendingAccountResourcesRequest{}
	o.Payer = *NewAccountType()
	o.MaxPayerResources = *NewUInt128()
	o.TrxResourceLimit = *NewUInt128()
	return &o
}

// Serialize CheckPendingAccountResourcesRequest
func (n CheckPendingAccountResourcesRequest) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Payer.Serialize(vb)
	vb = n.MaxPayerResources.Serialize(vb)
	vb = n.TrxResourceLimit.Serialize(vb)
	return vb
}

// DeserializeCheckPendingAccountResourcesRequest function
func DeserializeCheckPendingAccountResourcesRequest(vb *VariableBlob) (uint64,*CheckPendingAccountResourcesRequest,error) {
	var i,j uint64 = 0,0
	s := CheckPendingAccountResourcesRequest{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tPayer,err := DeserializeAccountType(&ovb); i+=j
	if err != nil {
		return 0, &CheckPendingAccountResourcesRequest{}, err
	}
	s.Payer = *tPayer
	ovb = (*vb)[i:]
	j,tMaxPayerResources,err := DeserializeUInt128(&ovb); i+=j
	if err != nil {
		return 0, &CheckPendingAccountResourcesRequest{}, err
	}
	s.MaxPayerResources = *tMaxPayerResources
	ovb = (*vb)[i:]
	j,tTrxResourceLimit,err := DeserializeUInt128(&ovb); i+=j
	if err != nil {
		return 0, &CheckPendingAccountResourcesRequest{}, err
	}
	s.TrxResourceLimit = *tTrxResourceLimit
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetPendingTransactionsRequest
// ----------------------------------------

// GetPendingTransactionsRequest type
type GetPendingTransactionsRequest struct {
    Limit UInt64 `json:"limit"`
}

// NewGetPendingTransactionsRequest factory
func NewGetPendingTransactionsRequest() *GetPendingTransactionsRequest {
	o := GetPendingTransactionsRequest{}
	o.Limit = *NewUInt64()
	return &o
}

// Serialize GetPendingTransactionsRequest
func (n GetPendingTransactionsRequest) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Limit.Serialize(vb)
	return vb
}

// DeserializeGetPendingTransactionsRequest function
func DeserializeGetPendingTransactionsRequest(vb *VariableBlob) (uint64,*GetPendingTransactionsRequest,error) {
	var i,j uint64 = 0,0
	s := GetPendingTransactionsRequest{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tLimit,err := DeserializeUInt64(&ovb); i+=j
	if err != nil {
		return 0, &GetPendingTransactionsRequest{}, err
	}
	s.Limit = *tLimit
	return i, &s, nil
}

// ----------------------------------------
//  Variant: MempoolRPCRequest
// ----------------------------------------

// MempoolRPCRequest type
type MempoolRPCRequest struct {
	Value interface{}
}

// NewMempoolRPCRequest factory
func NewMempoolRPCRequest() *MempoolRPCRequest {
	v := MempoolRPCRequest{}
	v.Value = NewMempoolReservedRequest()
	return &v
}

// Serialize MempoolRPCRequest
func (n MempoolRPCRequest) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *MempoolReservedRequest:
			i = 0
		case *CheckPendingAccountResourcesRequest:
			i = 1
		case *GetPendingTransactionsRequest:
			i = 2
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName MempoolRPCRequest
func (n MempoolRPCRequest) TypeToName() (string) {
	switch n.Value.(type) {
		case *MempoolReservedRequest:
			return "koinos::rpc::mempool::mempool_reserved_request"
		case *CheckPendingAccountResourcesRequest:
			return "koinos::rpc::mempool::check_pending_account_resources_request"
		case *GetPendingTransactionsRequest:
			return "koinos::rpc::mempool::get_pending_transactions_request"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON MempoolRPCRequest
func (n MempoolRPCRequest) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeMempoolRPCRequest function
func DeserializeMempoolRPCRequest(vb *VariableBlob) (uint64,*MempoolRPCRequest,error) {
	var v MempoolRPCRequest
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewMempoolReservedRequest()
		case 1:
			ovb := (*vb)[i:]
			k,x,err := DeserializeCheckPendingAccountResourcesRequest(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 2:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetPendingTransactionsRequest(&ovb)
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

// UnmarshalJSON *MempoolRPCRequest
func (n *MempoolRPCRequest) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::rpc::mempool::mempool_reserved_request":
			v := NewMempoolReservedRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::mempool::check_pending_account_resources_request":
			v := NewCheckPendingAccountResourcesRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::mempool::get_pending_transactions_request":
			v := NewGetPendingTransactionsRequest()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

	return nil
}


// ----------------------------------------
//  Struct: MempoolReservedResponse
// ----------------------------------------

// MempoolReservedResponse type
type MempoolReservedResponse struct {
}

// NewMempoolReservedResponse factory
func NewMempoolReservedResponse() *MempoolReservedResponse {
	o := MempoolReservedResponse{}
	return &o
}

// Serialize MempoolReservedResponse
func (n MempoolReservedResponse) Serialize(vb *VariableBlob) *VariableBlob {
	return vb
}

// DeserializeMempoolReservedResponse function
func DeserializeMempoolReservedResponse(vb *VariableBlob) (uint64,*MempoolReservedResponse,error) {
	var i uint64 = 0
	s := MempoolReservedResponse{}
	
	return i, &s, nil
}

// ----------------------------------------
//  Struct: MempoolErrorResponse
// ----------------------------------------

// MempoolErrorResponse type
type MempoolErrorResponse struct {
    ErrorText String `json:"error_text"`
    ErrorData String `json:"error_data"`
}

// NewMempoolErrorResponse factory
func NewMempoolErrorResponse() *MempoolErrorResponse {
	o := MempoolErrorResponse{}
	o.ErrorText = *NewString()
	o.ErrorData = *NewString()
	return &o
}

// Serialize MempoolErrorResponse
func (n MempoolErrorResponse) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.ErrorText.Serialize(vb)
	vb = n.ErrorData.Serialize(vb)
	return vb
}

// DeserializeMempoolErrorResponse function
func DeserializeMempoolErrorResponse(vb *VariableBlob) (uint64,*MempoolErrorResponse,error) {
	var i,j uint64 = 0,0
	s := MempoolErrorResponse{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tErrorText,err := DeserializeString(&ovb); i+=j
	if err != nil {
		return 0, &MempoolErrorResponse{}, err
	}
	s.ErrorText = *tErrorText
	ovb = (*vb)[i:]
	j,tErrorData,err := DeserializeString(&ovb); i+=j
	if err != nil {
		return 0, &MempoolErrorResponse{}, err
	}
	s.ErrorData = *tErrorData
	return i, &s, nil
}

// ----------------------------------------
//  Struct: CheckPendingAccountResourcesResponse
// ----------------------------------------

// CheckPendingAccountResourcesResponse type
type CheckPendingAccountResourcesResponse struct {
    Success Boolean `json:"success"`
}

// NewCheckPendingAccountResourcesResponse factory
func NewCheckPendingAccountResourcesResponse() *CheckPendingAccountResourcesResponse {
	o := CheckPendingAccountResourcesResponse{}
	o.Success = *NewBoolean()
	return &o
}

// Serialize CheckPendingAccountResourcesResponse
func (n CheckPendingAccountResourcesResponse) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Success.Serialize(vb)
	return vb
}

// DeserializeCheckPendingAccountResourcesResponse function
func DeserializeCheckPendingAccountResourcesResponse(vb *VariableBlob) (uint64,*CheckPendingAccountResourcesResponse,error) {
	var i,j uint64 = 0,0
	s := CheckPendingAccountResourcesResponse{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tSuccess,err := DeserializeBoolean(&ovb); i+=j
	if err != nil {
		return 0, &CheckPendingAccountResourcesResponse{}, err
	}
	s.Success = *tSuccess
	return i, &s, nil
}

// ----------------------------------------
//  Struct: GetPendingTransactionsResponse
// ----------------------------------------

// GetPendingTransactionsResponse type
type GetPendingTransactionsResponse struct {
    Transactions VectorTransaction `json:"transactions"`
}

// NewGetPendingTransactionsResponse factory
func NewGetPendingTransactionsResponse() *GetPendingTransactionsResponse {
	o := GetPendingTransactionsResponse{}
	o.Transactions = *NewVectorTransaction()
	return &o
}

// Serialize GetPendingTransactionsResponse
func (n GetPendingTransactionsResponse) Serialize(vb *VariableBlob) *VariableBlob {
	vb = n.Transactions.Serialize(vb)
	return vb
}

// DeserializeGetPendingTransactionsResponse function
func DeserializeGetPendingTransactionsResponse(vb *VariableBlob) (uint64,*GetPendingTransactionsResponse,error) {
	var i,j uint64 = 0,0
	s := GetPendingTransactionsResponse{}
	var ovb VariableBlob
	ovb = (*vb)[i:]
	j,tTransactions,err := DeserializeVectorTransaction(&ovb); i+=j
	if err != nil {
		return 0, &GetPendingTransactionsResponse{}, err
	}
	s.Transactions = *tTransactions
	return i, &s, nil
}

// ----------------------------------------
//  Variant: MempoolRPCResponse
// ----------------------------------------

// MempoolRPCResponse type
type MempoolRPCResponse struct {
	Value interface{}
}

// NewMempoolRPCResponse factory
func NewMempoolRPCResponse() *MempoolRPCResponse {
	v := MempoolRPCResponse{}
	v.Value = NewMempoolReservedResponse()
	return &v
}

// Serialize MempoolRPCResponse
func (n MempoolRPCResponse) Serialize(vb *VariableBlob) *VariableBlob {
	var i uint64
	switch n.Value.(type) {
		case *MempoolReservedResponse:
			i = 0
		case *MempoolErrorResponse:
			i = 1
		case *CheckPendingAccountResourcesResponse:
			i = 2
		case *GetPendingTransactionsResponse:
			i = 3
		default:
			panic("Unknown variant type")
	}

	vb = EncodeVarint(vb, i)
	ser,_ := n.Value.(Serializeable)
	return ser.Serialize(vb)
}

// TypeToName MempoolRPCResponse
func (n MempoolRPCResponse) TypeToName() (string) {
	switch n.Value.(type) {
		case *MempoolReservedResponse:
			return "koinos::rpc::mempool::mempool_reserved_response"
		case *MempoolErrorResponse:
			return "koinos::rpc::mempool::mempool_error_response"
		case *CheckPendingAccountResourcesResponse:
			return "koinos::rpc::mempool::check_pending_account_resources_response"
		case *GetPendingTransactionsResponse:
			return "koinos::rpc::mempool::get_pending_transactions_response"
		default:
			panic("Variant type is not serializeable.")
	}
}

// MarshalJSON MempoolRPCResponse
func (n MempoolRPCResponse) MarshalJSON() ([]byte, error) {
	variant := struct {
		Type string `json:"type"`
		Value *interface{} `json:"value"`
	}{
		n.TypeToName(),
		&n.Value,
	}

	return json.Marshal(&variant)
}

// DeserializeMempoolRPCResponse function
func DeserializeMempoolRPCResponse(vb *VariableBlob) (uint64,*MempoolRPCResponse,error) {
	var v MempoolRPCResponse
	typeID,i := binary.Uvarint(*vb)
	if i <= 0 {
		return 0, &v, errors.New("could not deserialize variant tag")
	}
	var j uint64

	switch( typeID ) {
		case 0:
			v.Value = NewMempoolReservedResponse()
		case 1:
			ovb := (*vb)[i:]
			k,x,err := DeserializeMempoolErrorResponse(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 2:
			ovb := (*vb)[i:]
			k,x,err := DeserializeCheckPendingAccountResourcesResponse(&ovb)
			if err != nil {
				return 0, &v, err
			}
			j = k
			v.Value = x
		case 3:
			ovb := (*vb)[i:]
			k,x,err := DeserializeGetPendingTransactionsResponse(&ovb)
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

// UnmarshalJSON *MempoolRPCResponse
func (n *MempoolRPCResponse) UnmarshalJSON(data []byte) error {
	variant := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}

	err := json.Unmarshal(data, &variant)
	if err != nil {
		return err
	}

	switch variant.Type {
		case "koinos::rpc::mempool::mempool_reserved_response":
			v := NewMempoolReservedResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::mempool::mempool_error_response":
			v := NewMempoolErrorResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::mempool::check_pending_account_resources_response":
			v := NewCheckPendingAccountResourcesResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		case "koinos::rpc::mempool::get_pending_transactions_response":
			v := NewGetPendingTransactionsResponse()
			json.Unmarshal(variant.Value, &v)
			n.Value = v
		default:
			return errors.New("unknown variant type: " + variant.Type)
	}

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
	v.Opaque.Type = "koinos::protocol::active_block_data"
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
		if strings.Compare(obj.Opaque.Type, "koinos::protocol::active_block_data") != 0 {
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
	v.Opaque.Type = "koinos::protocol::active_transaction_data"
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
		if strings.Compare(obj.Opaque.Type, "koinos::protocol::active_transaction_data") != 0 {
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
	v.Opaque.Type = "koinos::protocol::block"
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
		if strings.Compare(obj.Opaque.Type, "koinos::protocol::block") != 0 {
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
	v.Opaque.Type = "koinos::protocol::block_receipt"
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
		if strings.Compare(obj.Opaque.Type, "koinos::protocol::block_receipt") != 0 {
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
	v.Opaque.Type = "koinos::protocol::passive_block_data"
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
		if strings.Compare(obj.Opaque.Type, "koinos::protocol::passive_block_data") != 0 {
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
	v.Opaque.Type = "koinos::protocol::passive_transaction_data"
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
		if strings.Compare(obj.Opaque.Type, "koinos::protocol::passive_transaction_data") != 0 {
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
//  VectorBlockTopology
// ----------------------------------------

// VectorBlockTopology type
type VectorBlockTopology []BlockTopology

// NewVectorBlockTopology factory
func NewVectorBlockTopology() *VectorBlockTopology {
	o := VectorBlockTopology(make([]BlockTopology, 0))
	return &o
}

// Serialize VectorBlockTopology
func (n VectorBlockTopology) Serialize(vb *VariableBlob) *VariableBlob {
	header := make([]byte, binary.MaxVarintLen64)
	bytes := binary.PutUvarint(header, uint64(len(n)))
	ovb := append(*vb, header[:bytes]...)
	vb = &ovb
	for _, item := range n {
		vb = item.Serialize(vb)
	}

	return vb
}
// DeserializeVectorBlockTopology function
func DeserializeVectorBlockTopology(vb *VariableBlob) (uint64,*VectorBlockTopology,error) {
	var result VectorBlockTopology
	size,bytes := binary.Uvarint(*vb)
	if bytes <= 0 {
		return 0, &result, errors.New("could not deserialize multihash id")
	}
	result = VectorBlockTopology(make([]BlockTopology, 0, size))
	i := uint64(bytes)
	var j uint64
	var item *BlockTopology
	var err error
	for num := uint64(0); num < size; num++ {
		ovb := (*vb)[i:]
		j,item,err = DeserializeBlockTopology(&ovb)
		if nil != err {
			var v VectorBlockTopology
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
//  VectorTransaction
// ----------------------------------------

// VectorTransaction type
type VectorTransaction []Transaction

// NewVectorTransaction factory
func NewVectorTransaction() *VectorTransaction {
	o := VectorTransaction(make([]Transaction, 0))
	return &o
}

// Serialize VectorTransaction
func (n VectorTransaction) Serialize(vb *VariableBlob) *VariableBlob {
	header := make([]byte, binary.MaxVarintLen64)
	bytes := binary.PutUvarint(header, uint64(len(n)))
	ovb := append(*vb, header[:bytes]...)
	vb = &ovb
	for _, item := range n {
		vb = item.Serialize(vb)
	}

	return vb
}
// DeserializeVectorTransaction function
func DeserializeVectorTransaction(vb *VariableBlob) (uint64,*VectorTransaction,error) {
	var result VectorTransaction
	size,bytes := binary.Uvarint(*vb)
	if bytes <= 0 {
		return 0, &result, errors.New("could not deserialize multihash id")
	}
	result = VectorTransaction(make([]Transaction, 0, size))
	i := uint64(bytes)
	var j uint64
	var item *Transaction
	var err error
	for num := uint64(0); num < size; num++ {
		ovb := (*vb)[i:]
		j,item,err = DeserializeTransaction(&ovb)
		if nil != err {
			var v VectorTransaction
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


