// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package layerzero

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MessagingFee is an auto generated low-level Go binding around an user-defined struct.
type MessagingFee struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}

// MessagingParams is an auto generated low-level Go binding around an user-defined struct.
type MessagingParams struct {
	DstEid       uint32
	Receiver     [32]byte
	Message      []byte
	Options      []byte
	PayInLzToken bool
}

// MessagingReceipt is an auto generated low-level Go binding around an user-defined struct.
type MessagingReceipt struct {
	Guid  [32]byte
	Nonce uint64
	Fee   MessagingFee
}

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
}

// SetConfigParam is an auto generated low-level Go binding around an user-defined struct.
type SetConfigParam struct {
	Eid        uint32
	ConfigType uint32
	Config     []byte
}

// LayerzeroMetaData contains all meta data concerning the Layerzero contract.
var LayerzeroMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"LZ_AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_ComposeExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"LZ_ComposeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_DefaultReceiveLibUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_DefaultSendLibUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredNative\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"suppliedNative\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredLzToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"suppliedLzToken\",\"type\":\"uint256\"}],\"name\":\"LZ_InsufficientFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_InvalidExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"LZ_InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_InvalidPayloadHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_InvalidReceiveLibrary\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlyNonDefaultLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlyReceiveLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlyRegisteredLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlyRegisteredOrDefaultLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlySendLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_PathNotInitializable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_PathNotVerifiable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"LZ_PayloadHashNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_SameValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_SendReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_UnsupportedEid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_UnsupportedInterface\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_ZeroLzTokenFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer_NativeFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Transfer_ToAddressIsZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"}],\"name\":\"ComposeDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"ComposeSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"DefaultReceiveLibrarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldLib\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"DefaultReceiveLibraryTimeoutSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"DefaultSendLibrarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"DelegateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"InboundNonceSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"LibraryRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"LzComposeAlert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"LzReceiveAlert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"LzTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"name\":\"PacketBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"PacketDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"name\":\"PacketNilified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sendLibrary\",\"type\":\"address\"}],\"name\":\"PacketSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"name\":\"PacketVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"ReceiveLibrarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldLib\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"ReceiveLibraryTimeoutSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"SendLibrarySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EMPTY_PAYLOAD_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NIL_PAYLOAD_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockedLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"clear\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"}],\"name\":\"composeQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"}],\"name\":\"defaultReceiveLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"}],\"name\":\"defaultReceiveLibraryTimeout\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"}],\"name\":\"defaultSendLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oapp\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eid\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_configType\",\"type\":\"uint32\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"}],\"name\":\"getReceiveLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredLibraries\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSendContext\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"}],\"name\":\"getSendLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"}],\"name\":\"inboundNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"inboundNonce\",\"type\":\"uint64\"}],\"name\":\"inboundPayloadHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"initializable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"}],\"name\":\"isDefaultSendLibrary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"}],\"name\":\"isRegisteredLibrary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSendingMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"}],\"name\":\"isSupportedEid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_actualReceiveLib\",\"type\":\"address\"}],\"name\":\"isValidReceiveLibrary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"lazyInboundNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"_index\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzCompose\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"_index\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_reason\",\"type\":\"bytes\"}],\"name\":\"lzComposeAlert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_reason\",\"type\":\"bytes\"}],\"name\":\"lzReceiveAlert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lzToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_receiver\",\"type\":\"bytes32\"}],\"name\":\"nextGuid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"nilify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"}],\"name\":\"outboundNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"payInLzToken\",\"type\":\"bool\"}],\"internalType\":\"structMessagingParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"quote\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"}],\"name\":\"receiveLibraryTimeout\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"}],\"name\":\"registerLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"payInLzToken\",\"type\":\"bool\"}],\"internalType\":\"structMessagingParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"fee\",\"type\":\"tuple\"}],\"internalType\":\"structMessagingReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"_index\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendCompose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"configType\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"internalType\":\"structSetConfigParam[]\",\"name\":\"_params\",\"type\":\"tuple[]\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_newLib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gracePeriod\",\"type\":\"uint256\"}],\"name\":\"setDefaultReceiveLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expiry\",\"type\":\"uint256\"}],\"name\":\"setDefaultReceiveLibraryTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_newLib\",\"type\":\"address\"}],\"name\":\"setDefaultSendLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lzToken\",\"type\":\"address\"}],\"name\":\"setLzToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_newLib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gracePeriod\",\"type\":\"uint256\"}],\"name\":\"setReceiveLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expiry\",\"type\":\"uint256\"}],\"name\":\"setReceiveLibraryTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_newLib\",\"type\":\"address\"}],\"name\":\"setSendLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"}],\"name\":\"skip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"verifiable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LayerzeroABI is the input ABI used to generate the binding from.
// Deprecated: Use LayerzeroMetaData.ABI instead.
var LayerzeroABI = LayerzeroMetaData.ABI

// Layerzero is an auto generated Go binding around an Ethereum contract.
type Layerzero struct {
	LayerzeroCaller     // Read-only binding to the contract
	LayerzeroTransactor // Write-only binding to the contract
	LayerzeroFilterer   // Log filterer for contract events
}

// LayerzeroCaller is an auto generated read-only Go binding around an Ethereum contract.
type LayerzeroCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LayerzeroTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LayerzeroTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LayerzeroFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LayerzeroFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LayerzeroSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LayerzeroSession struct {
	Contract     *Layerzero        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LayerzeroCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LayerzeroCallerSession struct {
	Contract *LayerzeroCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// LayerzeroTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LayerzeroTransactorSession struct {
	Contract     *LayerzeroTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LayerzeroRaw is an auto generated low-level Go binding around an Ethereum contract.
type LayerzeroRaw struct {
	Contract *Layerzero // Generic contract binding to access the raw methods on
}

// LayerzeroCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LayerzeroCallerRaw struct {
	Contract *LayerzeroCaller // Generic read-only contract binding to access the raw methods on
}

// LayerzeroTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LayerzeroTransactorRaw struct {
	Contract *LayerzeroTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLayerzero creates a new instance of Layerzero, bound to a specific deployed contract.
func NewLayerzero(address common.Address, backend bind.ContractBackend) (*Layerzero, error) {
	contract, err := bindLayerzero(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Layerzero{LayerzeroCaller: LayerzeroCaller{contract: contract}, LayerzeroTransactor: LayerzeroTransactor{contract: contract}, LayerzeroFilterer: LayerzeroFilterer{contract: contract}}, nil
}

// NewLayerzeroCaller creates a new read-only instance of Layerzero, bound to a specific deployed contract.
func NewLayerzeroCaller(address common.Address, caller bind.ContractCaller) (*LayerzeroCaller, error) {
	contract, err := bindLayerzero(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LayerzeroCaller{contract: contract}, nil
}

// NewLayerzeroTransactor creates a new write-only instance of Layerzero, bound to a specific deployed contract.
func NewLayerzeroTransactor(address common.Address, transactor bind.ContractTransactor) (*LayerzeroTransactor, error) {
	contract, err := bindLayerzero(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LayerzeroTransactor{contract: contract}, nil
}

// NewLayerzeroFilterer creates a new log filterer instance of Layerzero, bound to a specific deployed contract.
func NewLayerzeroFilterer(address common.Address, filterer bind.ContractFilterer) (*LayerzeroFilterer, error) {
	contract, err := bindLayerzero(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LayerzeroFilterer{contract: contract}, nil
}

// bindLayerzero binds a generic wrapper to an already deployed contract.
func bindLayerzero(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LayerzeroMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Layerzero *LayerzeroRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Layerzero.Contract.LayerzeroCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Layerzero *LayerzeroRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Layerzero.Contract.LayerzeroTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Layerzero *LayerzeroRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Layerzero.Contract.LayerzeroTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Layerzero *LayerzeroCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Layerzero.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Layerzero *LayerzeroTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Layerzero.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Layerzero *LayerzeroTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Layerzero.Contract.contract.Transact(opts, method, params...)
}

// EMPTYPAYLOADHASH is a free data retrieval call binding the contract method 0xcb5026b9.
//
// Solidity: function EMPTY_PAYLOAD_HASH() view returns(bytes32)
func (_Layerzero *LayerzeroCaller) EMPTYPAYLOADHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "EMPTY_PAYLOAD_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EMPTYPAYLOADHASH is a free data retrieval call binding the contract method 0xcb5026b9.
//
// Solidity: function EMPTY_PAYLOAD_HASH() view returns(bytes32)
func (_Layerzero *LayerzeroSession) EMPTYPAYLOADHASH() ([32]byte, error) {
	return _Layerzero.Contract.EMPTYPAYLOADHASH(&_Layerzero.CallOpts)
}

// EMPTYPAYLOADHASH is a free data retrieval call binding the contract method 0xcb5026b9.
//
// Solidity: function EMPTY_PAYLOAD_HASH() view returns(bytes32)
func (_Layerzero *LayerzeroCallerSession) EMPTYPAYLOADHASH() ([32]byte, error) {
	return _Layerzero.Contract.EMPTYPAYLOADHASH(&_Layerzero.CallOpts)
}

// NILPAYLOADHASH is a free data retrieval call binding the contract method 0x2baf0be7.
//
// Solidity: function NIL_PAYLOAD_HASH() view returns(bytes32)
func (_Layerzero *LayerzeroCaller) NILPAYLOADHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "NIL_PAYLOAD_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NILPAYLOADHASH is a free data retrieval call binding the contract method 0x2baf0be7.
//
// Solidity: function NIL_PAYLOAD_HASH() view returns(bytes32)
func (_Layerzero *LayerzeroSession) NILPAYLOADHASH() ([32]byte, error) {
	return _Layerzero.Contract.NILPAYLOADHASH(&_Layerzero.CallOpts)
}

// NILPAYLOADHASH is a free data retrieval call binding the contract method 0x2baf0be7.
//
// Solidity: function NIL_PAYLOAD_HASH() view returns(bytes32)
func (_Layerzero *LayerzeroCallerSession) NILPAYLOADHASH() ([32]byte, error) {
	return _Layerzero.Contract.NILPAYLOADHASH(&_Layerzero.CallOpts)
}

// BlockedLibrary is a free data retrieval call binding the contract method 0x73318091.
//
// Solidity: function blockedLibrary() view returns(address)
func (_Layerzero *LayerzeroCaller) BlockedLibrary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "blockedLibrary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlockedLibrary is a free data retrieval call binding the contract method 0x73318091.
//
// Solidity: function blockedLibrary() view returns(address)
func (_Layerzero *LayerzeroSession) BlockedLibrary() (common.Address, error) {
	return _Layerzero.Contract.BlockedLibrary(&_Layerzero.CallOpts)
}

// BlockedLibrary is a free data retrieval call binding the contract method 0x73318091.
//
// Solidity: function blockedLibrary() view returns(address)
func (_Layerzero *LayerzeroCallerSession) BlockedLibrary() (common.Address, error) {
	return _Layerzero.Contract.BlockedLibrary(&_Layerzero.CallOpts)
}

// ComposeQueue is a free data retrieval call binding the contract method 0x35d330b0.
//
// Solidity: function composeQueue(address from, address to, bytes32 guid, uint16 index) view returns(bytes32 messageHash)
func (_Layerzero *LayerzeroCaller) ComposeQueue(opts *bind.CallOpts, from common.Address, to common.Address, guid [32]byte, index uint16) ([32]byte, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "composeQueue", from, to, guid, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComposeQueue is a free data retrieval call binding the contract method 0x35d330b0.
//
// Solidity: function composeQueue(address from, address to, bytes32 guid, uint16 index) view returns(bytes32 messageHash)
func (_Layerzero *LayerzeroSession) ComposeQueue(from common.Address, to common.Address, guid [32]byte, index uint16) ([32]byte, error) {
	return _Layerzero.Contract.ComposeQueue(&_Layerzero.CallOpts, from, to, guid, index)
}

// ComposeQueue is a free data retrieval call binding the contract method 0x35d330b0.
//
// Solidity: function composeQueue(address from, address to, bytes32 guid, uint16 index) view returns(bytes32 messageHash)
func (_Layerzero *LayerzeroCallerSession) ComposeQueue(from common.Address, to common.Address, guid [32]byte, index uint16) ([32]byte, error) {
	return _Layerzero.Contract.ComposeQueue(&_Layerzero.CallOpts, from, to, guid, index)
}

// DefaultReceiveLibrary is a free data retrieval call binding the contract method 0x6f50a803.
//
// Solidity: function defaultReceiveLibrary(uint32 srcEid) view returns(address lib)
func (_Layerzero *LayerzeroCaller) DefaultReceiveLibrary(opts *bind.CallOpts, srcEid uint32) (common.Address, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "defaultReceiveLibrary", srcEid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultReceiveLibrary is a free data retrieval call binding the contract method 0x6f50a803.
//
// Solidity: function defaultReceiveLibrary(uint32 srcEid) view returns(address lib)
func (_Layerzero *LayerzeroSession) DefaultReceiveLibrary(srcEid uint32) (common.Address, error) {
	return _Layerzero.Contract.DefaultReceiveLibrary(&_Layerzero.CallOpts, srcEid)
}

// DefaultReceiveLibrary is a free data retrieval call binding the contract method 0x6f50a803.
//
// Solidity: function defaultReceiveLibrary(uint32 srcEid) view returns(address lib)
func (_Layerzero *LayerzeroCallerSession) DefaultReceiveLibrary(srcEid uint32) (common.Address, error) {
	return _Layerzero.Contract.DefaultReceiveLibrary(&_Layerzero.CallOpts, srcEid)
}

// DefaultReceiveLibraryTimeout is a free data retrieval call binding the contract method 0x6e83f5bb.
//
// Solidity: function defaultReceiveLibraryTimeout(uint32 srcEid) view returns(address lib, uint256 expiry)
func (_Layerzero *LayerzeroCaller) DefaultReceiveLibraryTimeout(opts *bind.CallOpts, srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "defaultReceiveLibraryTimeout", srcEid)

	outstruct := new(struct {
		Lib    common.Address
		Expiry *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lib = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Expiry = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DefaultReceiveLibraryTimeout is a free data retrieval call binding the contract method 0x6e83f5bb.
//
// Solidity: function defaultReceiveLibraryTimeout(uint32 srcEid) view returns(address lib, uint256 expiry)
func (_Layerzero *LayerzeroSession) DefaultReceiveLibraryTimeout(srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	return _Layerzero.Contract.DefaultReceiveLibraryTimeout(&_Layerzero.CallOpts, srcEid)
}

// DefaultReceiveLibraryTimeout is a free data retrieval call binding the contract method 0x6e83f5bb.
//
// Solidity: function defaultReceiveLibraryTimeout(uint32 srcEid) view returns(address lib, uint256 expiry)
func (_Layerzero *LayerzeroCallerSession) DefaultReceiveLibraryTimeout(srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	return _Layerzero.Contract.DefaultReceiveLibraryTimeout(&_Layerzero.CallOpts, srcEid)
}

// DefaultSendLibrary is a free data retrieval call binding the contract method 0xf64be4c7.
//
// Solidity: function defaultSendLibrary(uint32 dstEid) view returns(address lib)
func (_Layerzero *LayerzeroCaller) DefaultSendLibrary(opts *bind.CallOpts, dstEid uint32) (common.Address, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "defaultSendLibrary", dstEid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultSendLibrary is a free data retrieval call binding the contract method 0xf64be4c7.
//
// Solidity: function defaultSendLibrary(uint32 dstEid) view returns(address lib)
func (_Layerzero *LayerzeroSession) DefaultSendLibrary(dstEid uint32) (common.Address, error) {
	return _Layerzero.Contract.DefaultSendLibrary(&_Layerzero.CallOpts, dstEid)
}

// DefaultSendLibrary is a free data retrieval call binding the contract method 0xf64be4c7.
//
// Solidity: function defaultSendLibrary(uint32 dstEid) view returns(address lib)
func (_Layerzero *LayerzeroCallerSession) DefaultSendLibrary(dstEid uint32) (common.Address, error) {
	return _Layerzero.Contract.DefaultSendLibrary(&_Layerzero.CallOpts, dstEid)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address oapp) view returns(address delegate)
func (_Layerzero *LayerzeroCaller) Delegates(opts *bind.CallOpts, oapp common.Address) (common.Address, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "delegates", oapp)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address oapp) view returns(address delegate)
func (_Layerzero *LayerzeroSession) Delegates(oapp common.Address) (common.Address, error) {
	return _Layerzero.Contract.Delegates(&_Layerzero.CallOpts, oapp)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address oapp) view returns(address delegate)
func (_Layerzero *LayerzeroCallerSession) Delegates(oapp common.Address) (common.Address, error) {
	return _Layerzero.Contract.Delegates(&_Layerzero.CallOpts, oapp)
}

// Eid is a free data retrieval call binding the contract method 0x416ecebf.
//
// Solidity: function eid() view returns(uint32)
func (_Layerzero *LayerzeroCaller) Eid(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "eid")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Eid is a free data retrieval call binding the contract method 0x416ecebf.
//
// Solidity: function eid() view returns(uint32)
func (_Layerzero *LayerzeroSession) Eid() (uint32, error) {
	return _Layerzero.Contract.Eid(&_Layerzero.CallOpts)
}

// Eid is a free data retrieval call binding the contract method 0x416ecebf.
//
// Solidity: function eid() view returns(uint32)
func (_Layerzero *LayerzeroCallerSession) Eid() (uint32, error) {
	return _Layerzero.Contract.Eid(&_Layerzero.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0x2b3197b9.
//
// Solidity: function getConfig(address _oapp, address _lib, uint32 _eid, uint32 _configType) view returns(bytes config)
func (_Layerzero *LayerzeroCaller) GetConfig(opts *bind.CallOpts, _oapp common.Address, _lib common.Address, _eid uint32, _configType uint32) ([]byte, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "getConfig", _oapp, _lib, _eid, _configType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0x2b3197b9.
//
// Solidity: function getConfig(address _oapp, address _lib, uint32 _eid, uint32 _configType) view returns(bytes config)
func (_Layerzero *LayerzeroSession) GetConfig(_oapp common.Address, _lib common.Address, _eid uint32, _configType uint32) ([]byte, error) {
	return _Layerzero.Contract.GetConfig(&_Layerzero.CallOpts, _oapp, _lib, _eid, _configType)
}

// GetConfig is a free data retrieval call binding the contract method 0x2b3197b9.
//
// Solidity: function getConfig(address _oapp, address _lib, uint32 _eid, uint32 _configType) view returns(bytes config)
func (_Layerzero *LayerzeroCallerSession) GetConfig(_oapp common.Address, _lib common.Address, _eid uint32, _configType uint32) ([]byte, error) {
	return _Layerzero.Contract.GetConfig(&_Layerzero.CallOpts, _oapp, _lib, _eid, _configType)
}

// GetReceiveLibrary is a free data retrieval call binding the contract method 0x402f8468.
//
// Solidity: function getReceiveLibrary(address _receiver, uint32 _srcEid) view returns(address lib, bool isDefault)
func (_Layerzero *LayerzeroCaller) GetReceiveLibrary(opts *bind.CallOpts, _receiver common.Address, _srcEid uint32) (struct {
	Lib       common.Address
	IsDefault bool
}, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "getReceiveLibrary", _receiver, _srcEid)

	outstruct := new(struct {
		Lib       common.Address
		IsDefault bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lib = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsDefault = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetReceiveLibrary is a free data retrieval call binding the contract method 0x402f8468.
//
// Solidity: function getReceiveLibrary(address _receiver, uint32 _srcEid) view returns(address lib, bool isDefault)
func (_Layerzero *LayerzeroSession) GetReceiveLibrary(_receiver common.Address, _srcEid uint32) (struct {
	Lib       common.Address
	IsDefault bool
}, error) {
	return _Layerzero.Contract.GetReceiveLibrary(&_Layerzero.CallOpts, _receiver, _srcEid)
}

// GetReceiveLibrary is a free data retrieval call binding the contract method 0x402f8468.
//
// Solidity: function getReceiveLibrary(address _receiver, uint32 _srcEid) view returns(address lib, bool isDefault)
func (_Layerzero *LayerzeroCallerSession) GetReceiveLibrary(_receiver common.Address, _srcEid uint32) (struct {
	Lib       common.Address
	IsDefault bool
}, error) {
	return _Layerzero.Contract.GetReceiveLibrary(&_Layerzero.CallOpts, _receiver, _srcEid)
}

// GetRegisteredLibraries is a free data retrieval call binding the contract method 0x9132e5c3.
//
// Solidity: function getRegisteredLibraries() view returns(address[])
func (_Layerzero *LayerzeroCaller) GetRegisteredLibraries(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "getRegisteredLibraries")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredLibraries is a free data retrieval call binding the contract method 0x9132e5c3.
//
// Solidity: function getRegisteredLibraries() view returns(address[])
func (_Layerzero *LayerzeroSession) GetRegisteredLibraries() ([]common.Address, error) {
	return _Layerzero.Contract.GetRegisteredLibraries(&_Layerzero.CallOpts)
}

// GetRegisteredLibraries is a free data retrieval call binding the contract method 0x9132e5c3.
//
// Solidity: function getRegisteredLibraries() view returns(address[])
func (_Layerzero *LayerzeroCallerSession) GetRegisteredLibraries() ([]common.Address, error) {
	return _Layerzero.Contract.GetRegisteredLibraries(&_Layerzero.CallOpts)
}

// GetSendContext is a free data retrieval call binding the contract method 0x14f651a9.
//
// Solidity: function getSendContext() view returns(uint32, address)
func (_Layerzero *LayerzeroCaller) GetSendContext(opts *bind.CallOpts) (uint32, common.Address, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "getSendContext")

	if err != nil {
		return *new(uint32), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetSendContext is a free data retrieval call binding the contract method 0x14f651a9.
//
// Solidity: function getSendContext() view returns(uint32, address)
func (_Layerzero *LayerzeroSession) GetSendContext() (uint32, common.Address, error) {
	return _Layerzero.Contract.GetSendContext(&_Layerzero.CallOpts)
}

// GetSendContext is a free data retrieval call binding the contract method 0x14f651a9.
//
// Solidity: function getSendContext() view returns(uint32, address)
func (_Layerzero *LayerzeroCallerSession) GetSendContext() (uint32, common.Address, error) {
	return _Layerzero.Contract.GetSendContext(&_Layerzero.CallOpts)
}

// GetSendLibrary is a free data retrieval call binding the contract method 0xb96a277f.
//
// Solidity: function getSendLibrary(address _sender, uint32 _dstEid) view returns(address lib)
func (_Layerzero *LayerzeroCaller) GetSendLibrary(opts *bind.CallOpts, _sender common.Address, _dstEid uint32) (common.Address, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "getSendLibrary", _sender, _dstEid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSendLibrary is a free data retrieval call binding the contract method 0xb96a277f.
//
// Solidity: function getSendLibrary(address _sender, uint32 _dstEid) view returns(address lib)
func (_Layerzero *LayerzeroSession) GetSendLibrary(_sender common.Address, _dstEid uint32) (common.Address, error) {
	return _Layerzero.Contract.GetSendLibrary(&_Layerzero.CallOpts, _sender, _dstEid)
}

// GetSendLibrary is a free data retrieval call binding the contract method 0xb96a277f.
//
// Solidity: function getSendLibrary(address _sender, uint32 _dstEid) view returns(address lib)
func (_Layerzero *LayerzeroCallerSession) GetSendLibrary(_sender common.Address, _dstEid uint32) (common.Address, error) {
	return _Layerzero.Contract.GetSendLibrary(&_Layerzero.CallOpts, _sender, _dstEid)
}

// InboundNonce is a free data retrieval call binding the contract method 0xa0dd43fc.
//
// Solidity: function inboundNonce(address _receiver, uint32 _srcEid, bytes32 _sender) view returns(uint64)
func (_Layerzero *LayerzeroCaller) InboundNonce(opts *bind.CallOpts, _receiver common.Address, _srcEid uint32, _sender [32]byte) (uint64, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "inboundNonce", _receiver, _srcEid, _sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InboundNonce is a free data retrieval call binding the contract method 0xa0dd43fc.
//
// Solidity: function inboundNonce(address _receiver, uint32 _srcEid, bytes32 _sender) view returns(uint64)
func (_Layerzero *LayerzeroSession) InboundNonce(_receiver common.Address, _srcEid uint32, _sender [32]byte) (uint64, error) {
	return _Layerzero.Contract.InboundNonce(&_Layerzero.CallOpts, _receiver, _srcEid, _sender)
}

// InboundNonce is a free data retrieval call binding the contract method 0xa0dd43fc.
//
// Solidity: function inboundNonce(address _receiver, uint32 _srcEid, bytes32 _sender) view returns(uint64)
func (_Layerzero *LayerzeroCallerSession) InboundNonce(_receiver common.Address, _srcEid uint32, _sender [32]byte) (uint64, error) {
	return _Layerzero.Contract.InboundNonce(&_Layerzero.CallOpts, _receiver, _srcEid, _sender)
}

// InboundPayloadHash is a free data retrieval call binding the contract method 0xc9fc7bcd.
//
// Solidity: function inboundPayloadHash(address receiver, uint32 srcEid, bytes32 sender, uint64 inboundNonce) view returns(bytes32 payloadHash)
func (_Layerzero *LayerzeroCaller) InboundPayloadHash(opts *bind.CallOpts, receiver common.Address, srcEid uint32, sender [32]byte, inboundNonce uint64) ([32]byte, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "inboundPayloadHash", receiver, srcEid, sender, inboundNonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboundPayloadHash is a free data retrieval call binding the contract method 0xc9fc7bcd.
//
// Solidity: function inboundPayloadHash(address receiver, uint32 srcEid, bytes32 sender, uint64 inboundNonce) view returns(bytes32 payloadHash)
func (_Layerzero *LayerzeroSession) InboundPayloadHash(receiver common.Address, srcEid uint32, sender [32]byte, inboundNonce uint64) ([32]byte, error) {
	return _Layerzero.Contract.InboundPayloadHash(&_Layerzero.CallOpts, receiver, srcEid, sender, inboundNonce)
}

// InboundPayloadHash is a free data retrieval call binding the contract method 0xc9fc7bcd.
//
// Solidity: function inboundPayloadHash(address receiver, uint32 srcEid, bytes32 sender, uint64 inboundNonce) view returns(bytes32 payloadHash)
func (_Layerzero *LayerzeroCallerSession) InboundPayloadHash(receiver common.Address, srcEid uint32, sender [32]byte, inboundNonce uint64) ([32]byte, error) {
	return _Layerzero.Contract.InboundPayloadHash(&_Layerzero.CallOpts, receiver, srcEid, sender, inboundNonce)
}

// Initializable is a free data retrieval call binding the contract method 0x861e1ca5.
//
// Solidity: function initializable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_Layerzero *LayerzeroCaller) Initializable(opts *bind.CallOpts, _origin Origin, _receiver common.Address) (bool, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "initializable", _origin, _receiver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initializable is a free data retrieval call binding the contract method 0x861e1ca5.
//
// Solidity: function initializable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_Layerzero *LayerzeroSession) Initializable(_origin Origin, _receiver common.Address) (bool, error) {
	return _Layerzero.Contract.Initializable(&_Layerzero.CallOpts, _origin, _receiver)
}

// Initializable is a free data retrieval call binding the contract method 0x861e1ca5.
//
// Solidity: function initializable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_Layerzero *LayerzeroCallerSession) Initializable(_origin Origin, _receiver common.Address) (bool, error) {
	return _Layerzero.Contract.Initializable(&_Layerzero.CallOpts, _origin, _receiver)
}

// IsDefaultSendLibrary is a free data retrieval call binding the contract method 0xdc93c8a2.
//
// Solidity: function isDefaultSendLibrary(address _sender, uint32 _dstEid) view returns(bool)
func (_Layerzero *LayerzeroCaller) IsDefaultSendLibrary(opts *bind.CallOpts, _sender common.Address, _dstEid uint32) (bool, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "isDefaultSendLibrary", _sender, _dstEid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDefaultSendLibrary is a free data retrieval call binding the contract method 0xdc93c8a2.
//
// Solidity: function isDefaultSendLibrary(address _sender, uint32 _dstEid) view returns(bool)
func (_Layerzero *LayerzeroSession) IsDefaultSendLibrary(_sender common.Address, _dstEid uint32) (bool, error) {
	return _Layerzero.Contract.IsDefaultSendLibrary(&_Layerzero.CallOpts, _sender, _dstEid)
}

// IsDefaultSendLibrary is a free data retrieval call binding the contract method 0xdc93c8a2.
//
// Solidity: function isDefaultSendLibrary(address _sender, uint32 _dstEid) view returns(bool)
func (_Layerzero *LayerzeroCallerSession) IsDefaultSendLibrary(_sender common.Address, _dstEid uint32) (bool, error) {
	return _Layerzero.Contract.IsDefaultSendLibrary(&_Layerzero.CallOpts, _sender, _dstEid)
}

// IsRegisteredLibrary is a free data retrieval call binding the contract method 0xdc706a62.
//
// Solidity: function isRegisteredLibrary(address lib) view returns(bool)
func (_Layerzero *LayerzeroCaller) IsRegisteredLibrary(opts *bind.CallOpts, lib common.Address) (bool, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "isRegisteredLibrary", lib)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredLibrary is a free data retrieval call binding the contract method 0xdc706a62.
//
// Solidity: function isRegisteredLibrary(address lib) view returns(bool)
func (_Layerzero *LayerzeroSession) IsRegisteredLibrary(lib common.Address) (bool, error) {
	return _Layerzero.Contract.IsRegisteredLibrary(&_Layerzero.CallOpts, lib)
}

// IsRegisteredLibrary is a free data retrieval call binding the contract method 0xdc706a62.
//
// Solidity: function isRegisteredLibrary(address lib) view returns(bool)
func (_Layerzero *LayerzeroCallerSession) IsRegisteredLibrary(lib common.Address) (bool, error) {
	return _Layerzero.Contract.IsRegisteredLibrary(&_Layerzero.CallOpts, lib)
}

// IsSendingMessage is a free data retrieval call binding the contract method 0x79624ca9.
//
// Solidity: function isSendingMessage() view returns(bool)
func (_Layerzero *LayerzeroCaller) IsSendingMessage(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "isSendingMessage")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSendingMessage is a free data retrieval call binding the contract method 0x79624ca9.
//
// Solidity: function isSendingMessage() view returns(bool)
func (_Layerzero *LayerzeroSession) IsSendingMessage() (bool, error) {
	return _Layerzero.Contract.IsSendingMessage(&_Layerzero.CallOpts)
}

// IsSendingMessage is a free data retrieval call binding the contract method 0x79624ca9.
//
// Solidity: function isSendingMessage() view returns(bool)
func (_Layerzero *LayerzeroCallerSession) IsSendingMessage() (bool, error) {
	return _Layerzero.Contract.IsSendingMessage(&_Layerzero.CallOpts)
}

// IsSupportedEid is a free data retrieval call binding the contract method 0x6750cd4c.
//
// Solidity: function isSupportedEid(uint32 _eid) view returns(bool)
func (_Layerzero *LayerzeroCaller) IsSupportedEid(opts *bind.CallOpts, _eid uint32) (bool, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "isSupportedEid", _eid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedEid is a free data retrieval call binding the contract method 0x6750cd4c.
//
// Solidity: function isSupportedEid(uint32 _eid) view returns(bool)
func (_Layerzero *LayerzeroSession) IsSupportedEid(_eid uint32) (bool, error) {
	return _Layerzero.Contract.IsSupportedEid(&_Layerzero.CallOpts, _eid)
}

// IsSupportedEid is a free data retrieval call binding the contract method 0x6750cd4c.
//
// Solidity: function isSupportedEid(uint32 _eid) view returns(bool)
func (_Layerzero *LayerzeroCallerSession) IsSupportedEid(_eid uint32) (bool, error) {
	return _Layerzero.Contract.IsSupportedEid(&_Layerzero.CallOpts, _eid)
}

// IsValidReceiveLibrary is a free data retrieval call binding the contract method 0x9d7f9775.
//
// Solidity: function isValidReceiveLibrary(address _receiver, uint32 _srcEid, address _actualReceiveLib) view returns(bool)
func (_Layerzero *LayerzeroCaller) IsValidReceiveLibrary(opts *bind.CallOpts, _receiver common.Address, _srcEid uint32, _actualReceiveLib common.Address) (bool, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "isValidReceiveLibrary", _receiver, _srcEid, _actualReceiveLib)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidReceiveLibrary is a free data retrieval call binding the contract method 0x9d7f9775.
//
// Solidity: function isValidReceiveLibrary(address _receiver, uint32 _srcEid, address _actualReceiveLib) view returns(bool)
func (_Layerzero *LayerzeroSession) IsValidReceiveLibrary(_receiver common.Address, _srcEid uint32, _actualReceiveLib common.Address) (bool, error) {
	return _Layerzero.Contract.IsValidReceiveLibrary(&_Layerzero.CallOpts, _receiver, _srcEid, _actualReceiveLib)
}

// IsValidReceiveLibrary is a free data retrieval call binding the contract method 0x9d7f9775.
//
// Solidity: function isValidReceiveLibrary(address _receiver, uint32 _srcEid, address _actualReceiveLib) view returns(bool)
func (_Layerzero *LayerzeroCallerSession) IsValidReceiveLibrary(_receiver common.Address, _srcEid uint32, _actualReceiveLib common.Address) (bool, error) {
	return _Layerzero.Contract.IsValidReceiveLibrary(&_Layerzero.CallOpts, _receiver, _srcEid, _actualReceiveLib)
}

// LazyInboundNonce is a free data retrieval call binding the contract method 0x5b17bb70.
//
// Solidity: function lazyInboundNonce(address receiver, uint32 srcEid, bytes32 sender) view returns(uint64 nonce)
func (_Layerzero *LayerzeroCaller) LazyInboundNonce(opts *bind.CallOpts, receiver common.Address, srcEid uint32, sender [32]byte) (uint64, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "lazyInboundNonce", receiver, srcEid, sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LazyInboundNonce is a free data retrieval call binding the contract method 0x5b17bb70.
//
// Solidity: function lazyInboundNonce(address receiver, uint32 srcEid, bytes32 sender) view returns(uint64 nonce)
func (_Layerzero *LayerzeroSession) LazyInboundNonce(receiver common.Address, srcEid uint32, sender [32]byte) (uint64, error) {
	return _Layerzero.Contract.LazyInboundNonce(&_Layerzero.CallOpts, receiver, srcEid, sender)
}

// LazyInboundNonce is a free data retrieval call binding the contract method 0x5b17bb70.
//
// Solidity: function lazyInboundNonce(address receiver, uint32 srcEid, bytes32 sender) view returns(uint64 nonce)
func (_Layerzero *LayerzeroCallerSession) LazyInboundNonce(receiver common.Address, srcEid uint32, sender [32]byte) (uint64, error) {
	return _Layerzero.Contract.LazyInboundNonce(&_Layerzero.CallOpts, receiver, srcEid, sender)
}

// LzToken is a free data retrieval call binding the contract method 0xe4fe1d94.
//
// Solidity: function lzToken() view returns(address)
func (_Layerzero *LayerzeroCaller) LzToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "lzToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LzToken is a free data retrieval call binding the contract method 0xe4fe1d94.
//
// Solidity: function lzToken() view returns(address)
func (_Layerzero *LayerzeroSession) LzToken() (common.Address, error) {
	return _Layerzero.Contract.LzToken(&_Layerzero.CallOpts)
}

// LzToken is a free data retrieval call binding the contract method 0xe4fe1d94.
//
// Solidity: function lzToken() view returns(address)
func (_Layerzero *LayerzeroCallerSession) LzToken() (common.Address, error) {
	return _Layerzero.Contract.LzToken(&_Layerzero.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_Layerzero *LayerzeroCaller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "nativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_Layerzero *LayerzeroSession) NativeToken() (common.Address, error) {
	return _Layerzero.Contract.NativeToken(&_Layerzero.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_Layerzero *LayerzeroCallerSession) NativeToken() (common.Address, error) {
	return _Layerzero.Contract.NativeToken(&_Layerzero.CallOpts)
}

// NextGuid is a free data retrieval call binding the contract method 0xaafe5e07.
//
// Solidity: function nextGuid(address _sender, uint32 _dstEid, bytes32 _receiver) view returns(bytes32)
func (_Layerzero *LayerzeroCaller) NextGuid(opts *bind.CallOpts, _sender common.Address, _dstEid uint32, _receiver [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "nextGuid", _sender, _dstEid, _receiver)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextGuid is a free data retrieval call binding the contract method 0xaafe5e07.
//
// Solidity: function nextGuid(address _sender, uint32 _dstEid, bytes32 _receiver) view returns(bytes32)
func (_Layerzero *LayerzeroSession) NextGuid(_sender common.Address, _dstEid uint32, _receiver [32]byte) ([32]byte, error) {
	return _Layerzero.Contract.NextGuid(&_Layerzero.CallOpts, _sender, _dstEid, _receiver)
}

// NextGuid is a free data retrieval call binding the contract method 0xaafe5e07.
//
// Solidity: function nextGuid(address _sender, uint32 _dstEid, bytes32 _receiver) view returns(bytes32)
func (_Layerzero *LayerzeroCallerSession) NextGuid(_sender common.Address, _dstEid uint32, _receiver [32]byte) ([32]byte, error) {
	return _Layerzero.Contract.NextGuid(&_Layerzero.CallOpts, _sender, _dstEid, _receiver)
}

// OutboundNonce is a free data retrieval call binding the contract method 0x9c6d7340.
//
// Solidity: function outboundNonce(address sender, uint32 dstEid, bytes32 receiver) view returns(uint64 nonce)
func (_Layerzero *LayerzeroCaller) OutboundNonce(opts *bind.CallOpts, sender common.Address, dstEid uint32, receiver [32]byte) (uint64, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "outboundNonce", sender, dstEid, receiver)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OutboundNonce is a free data retrieval call binding the contract method 0x9c6d7340.
//
// Solidity: function outboundNonce(address sender, uint32 dstEid, bytes32 receiver) view returns(uint64 nonce)
func (_Layerzero *LayerzeroSession) OutboundNonce(sender common.Address, dstEid uint32, receiver [32]byte) (uint64, error) {
	return _Layerzero.Contract.OutboundNonce(&_Layerzero.CallOpts, sender, dstEid, receiver)
}

// OutboundNonce is a free data retrieval call binding the contract method 0x9c6d7340.
//
// Solidity: function outboundNonce(address sender, uint32 dstEid, bytes32 receiver) view returns(uint64 nonce)
func (_Layerzero *LayerzeroCallerSession) OutboundNonce(sender common.Address, dstEid uint32, receiver [32]byte) (uint64, error) {
	return _Layerzero.Contract.OutboundNonce(&_Layerzero.CallOpts, sender, dstEid, receiver)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Layerzero *LayerzeroCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Layerzero *LayerzeroSession) Owner() (common.Address, error) {
	return _Layerzero.Contract.Owner(&_Layerzero.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Layerzero *LayerzeroCallerSession) Owner() (common.Address, error) {
	return _Layerzero.Contract.Owner(&_Layerzero.CallOpts)
}

// Quote is a free data retrieval call binding the contract method 0xddc28c58.
//
// Solidity: function quote((uint32,bytes32,bytes,bytes,bool) _params, address _sender) view returns((uint256,uint256))
func (_Layerzero *LayerzeroCaller) Quote(opts *bind.CallOpts, _params MessagingParams, _sender common.Address) (MessagingFee, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "quote", _params, _sender)

	if err != nil {
		return *new(MessagingFee), err
	}

	out0 := *abi.ConvertType(out[0], new(MessagingFee)).(*MessagingFee)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xddc28c58.
//
// Solidity: function quote((uint32,bytes32,bytes,bytes,bool) _params, address _sender) view returns((uint256,uint256))
func (_Layerzero *LayerzeroSession) Quote(_params MessagingParams, _sender common.Address) (MessagingFee, error) {
	return _Layerzero.Contract.Quote(&_Layerzero.CallOpts, _params, _sender)
}

// Quote is a free data retrieval call binding the contract method 0xddc28c58.
//
// Solidity: function quote((uint32,bytes32,bytes,bytes,bool) _params, address _sender) view returns((uint256,uint256))
func (_Layerzero *LayerzeroCallerSession) Quote(_params MessagingParams, _sender common.Address) (MessagingFee, error) {
	return _Layerzero.Contract.Quote(&_Layerzero.CallOpts, _params, _sender)
}

// ReceiveLibraryTimeout is a free data retrieval call binding the contract method 0xef667aa1.
//
// Solidity: function receiveLibraryTimeout(address receiver, uint32 srcEid) view returns(address lib, uint256 expiry)
func (_Layerzero *LayerzeroCaller) ReceiveLibraryTimeout(opts *bind.CallOpts, receiver common.Address, srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "receiveLibraryTimeout", receiver, srcEid)

	outstruct := new(struct {
		Lib    common.Address
		Expiry *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lib = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Expiry = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReceiveLibraryTimeout is a free data retrieval call binding the contract method 0xef667aa1.
//
// Solidity: function receiveLibraryTimeout(address receiver, uint32 srcEid) view returns(address lib, uint256 expiry)
func (_Layerzero *LayerzeroSession) ReceiveLibraryTimeout(receiver common.Address, srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	return _Layerzero.Contract.ReceiveLibraryTimeout(&_Layerzero.CallOpts, receiver, srcEid)
}

// ReceiveLibraryTimeout is a free data retrieval call binding the contract method 0xef667aa1.
//
// Solidity: function receiveLibraryTimeout(address receiver, uint32 srcEid) view returns(address lib, uint256 expiry)
func (_Layerzero *LayerzeroCallerSession) ReceiveLibraryTimeout(receiver common.Address, srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	return _Layerzero.Contract.ReceiveLibraryTimeout(&_Layerzero.CallOpts, receiver, srcEid)
}

// Verifiable is a free data retrieval call binding the contract method 0xc9a54a99.
//
// Solidity: function verifiable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_Layerzero *LayerzeroCaller) Verifiable(opts *bind.CallOpts, _origin Origin, _receiver common.Address) (bool, error) {
	var out []interface{}
	err := _Layerzero.contract.Call(opts, &out, "verifiable", _origin, _receiver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verifiable is a free data retrieval call binding the contract method 0xc9a54a99.
//
// Solidity: function verifiable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_Layerzero *LayerzeroSession) Verifiable(_origin Origin, _receiver common.Address) (bool, error) {
	return _Layerzero.Contract.Verifiable(&_Layerzero.CallOpts, _origin, _receiver)
}

// Verifiable is a free data retrieval call binding the contract method 0xc9a54a99.
//
// Solidity: function verifiable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_Layerzero *LayerzeroCallerSession) Verifiable(_origin Origin, _receiver common.Address) (bool, error) {
	return _Layerzero.Contract.Verifiable(&_Layerzero.CallOpts, _origin, _receiver)
}

// Burn is a paid mutator transaction binding the contract method 0x40f80683.
//
// Solidity: function burn(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_Layerzero *LayerzeroTransactor) Burn(opts *bind.TransactOpts, _oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "burn", _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Burn is a paid mutator transaction binding the contract method 0x40f80683.
//
// Solidity: function burn(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_Layerzero *LayerzeroSession) Burn(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Layerzero.Contract.Burn(&_Layerzero.TransactOpts, _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Burn is a paid mutator transaction binding the contract method 0x40f80683.
//
// Solidity: function burn(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_Layerzero *LayerzeroTransactorSession) Burn(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Layerzero.Contract.Burn(&_Layerzero.TransactOpts, _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Clear is a paid mutator transaction binding the contract method 0x2a56c1b0.
//
// Solidity: function clear(address _oapp, (uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message) returns()
func (_Layerzero *LayerzeroTransactor) Clear(opts *bind.TransactOpts, _oapp common.Address, _origin Origin, _guid [32]byte, _message []byte) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "clear", _oapp, _origin, _guid, _message)
}

// Clear is a paid mutator transaction binding the contract method 0x2a56c1b0.
//
// Solidity: function clear(address _oapp, (uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message) returns()
func (_Layerzero *LayerzeroSession) Clear(_oapp common.Address, _origin Origin, _guid [32]byte, _message []byte) (*types.Transaction, error) {
	return _Layerzero.Contract.Clear(&_Layerzero.TransactOpts, _oapp, _origin, _guid, _message)
}

// Clear is a paid mutator transaction binding the contract method 0x2a56c1b0.
//
// Solidity: function clear(address _oapp, (uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message) returns()
func (_Layerzero *LayerzeroTransactorSession) Clear(_oapp common.Address, _origin Origin, _guid [32]byte, _message []byte) (*types.Transaction, error) {
	return _Layerzero.Contract.Clear(&_Layerzero.TransactOpts, _oapp, _origin, _guid, _message)
}

// LzCompose is a paid mutator transaction binding the contract method 0x91d20fa1.
//
// Solidity: function lzCompose(address _from, address _to, bytes32 _guid, uint16 _index, bytes _message, bytes _extraData) payable returns()
func (_Layerzero *LayerzeroTransactor) LzCompose(opts *bind.TransactOpts, _from common.Address, _to common.Address, _guid [32]byte, _index uint16, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "lzCompose", _from, _to, _guid, _index, _message, _extraData)
}

// LzCompose is a paid mutator transaction binding the contract method 0x91d20fa1.
//
// Solidity: function lzCompose(address _from, address _to, bytes32 _guid, uint16 _index, bytes _message, bytes _extraData) payable returns()
func (_Layerzero *LayerzeroSession) LzCompose(_from common.Address, _to common.Address, _guid [32]byte, _index uint16, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _Layerzero.Contract.LzCompose(&_Layerzero.TransactOpts, _from, _to, _guid, _index, _message, _extraData)
}

// LzCompose is a paid mutator transaction binding the contract method 0x91d20fa1.
//
// Solidity: function lzCompose(address _from, address _to, bytes32 _guid, uint16 _index, bytes _message, bytes _extraData) payable returns()
func (_Layerzero *LayerzeroTransactorSession) LzCompose(_from common.Address, _to common.Address, _guid [32]byte, _index uint16, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _Layerzero.Contract.LzCompose(&_Layerzero.TransactOpts, _from, _to, _guid, _index, _message, _extraData)
}

// LzComposeAlert is a paid mutator transaction binding the contract method 0x697fe6b6.
//
// Solidity: function lzComposeAlert(address _from, address _to, bytes32 _guid, uint16 _index, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_Layerzero *LayerzeroTransactor) LzComposeAlert(opts *bind.TransactOpts, _from common.Address, _to common.Address, _guid [32]byte, _index uint16, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "lzComposeAlert", _from, _to, _guid, _index, _gas, _value, _message, _extraData, _reason)
}

// LzComposeAlert is a paid mutator transaction binding the contract method 0x697fe6b6.
//
// Solidity: function lzComposeAlert(address _from, address _to, bytes32 _guid, uint16 _index, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_Layerzero *LayerzeroSession) LzComposeAlert(_from common.Address, _to common.Address, _guid [32]byte, _index uint16, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _Layerzero.Contract.LzComposeAlert(&_Layerzero.TransactOpts, _from, _to, _guid, _index, _gas, _value, _message, _extraData, _reason)
}

// LzComposeAlert is a paid mutator transaction binding the contract method 0x697fe6b6.
//
// Solidity: function lzComposeAlert(address _from, address _to, bytes32 _guid, uint16 _index, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_Layerzero *LayerzeroTransactorSession) LzComposeAlert(_from common.Address, _to common.Address, _guid [32]byte, _index uint16, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _Layerzero.Contract.LzComposeAlert(&_Layerzero.TransactOpts, _from, _to, _guid, _index, _gas, _value, _message, _extraData, _reason)
}

// LzReceive is a paid mutator transaction binding the contract method 0x0c0c389e.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, bytes _message, bytes _extraData) payable returns()
func (_Layerzero *LayerzeroTransactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _receiver common.Address, _guid [32]byte, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "lzReceive", _origin, _receiver, _guid, _message, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x0c0c389e.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, bytes _message, bytes _extraData) payable returns()
func (_Layerzero *LayerzeroSession) LzReceive(_origin Origin, _receiver common.Address, _guid [32]byte, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _Layerzero.Contract.LzReceive(&_Layerzero.TransactOpts, _origin, _receiver, _guid, _message, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x0c0c389e.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, bytes _message, bytes _extraData) payable returns()
func (_Layerzero *LayerzeroTransactorSession) LzReceive(_origin Origin, _receiver common.Address, _guid [32]byte, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _Layerzero.Contract.LzReceive(&_Layerzero.TransactOpts, _origin, _receiver, _guid, _message, _extraData)
}

// LzReceiveAlert is a paid mutator transaction binding the contract method 0x6bf73fa3.
//
// Solidity: function lzReceiveAlert((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_Layerzero *LayerzeroTransactor) LzReceiveAlert(opts *bind.TransactOpts, _origin Origin, _receiver common.Address, _guid [32]byte, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "lzReceiveAlert", _origin, _receiver, _guid, _gas, _value, _message, _extraData, _reason)
}

// LzReceiveAlert is a paid mutator transaction binding the contract method 0x6bf73fa3.
//
// Solidity: function lzReceiveAlert((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_Layerzero *LayerzeroSession) LzReceiveAlert(_origin Origin, _receiver common.Address, _guid [32]byte, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _Layerzero.Contract.LzReceiveAlert(&_Layerzero.TransactOpts, _origin, _receiver, _guid, _gas, _value, _message, _extraData, _reason)
}

// LzReceiveAlert is a paid mutator transaction binding the contract method 0x6bf73fa3.
//
// Solidity: function lzReceiveAlert((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_Layerzero *LayerzeroTransactorSession) LzReceiveAlert(_origin Origin, _receiver common.Address, _guid [32]byte, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _Layerzero.Contract.LzReceiveAlert(&_Layerzero.TransactOpts, _origin, _receiver, _guid, _gas, _value, _message, _extraData, _reason)
}

// Nilify is a paid mutator transaction binding the contract method 0x2e80fbf3.
//
// Solidity: function nilify(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_Layerzero *LayerzeroTransactor) Nilify(opts *bind.TransactOpts, _oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "nilify", _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Nilify is a paid mutator transaction binding the contract method 0x2e80fbf3.
//
// Solidity: function nilify(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_Layerzero *LayerzeroSession) Nilify(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Layerzero.Contract.Nilify(&_Layerzero.TransactOpts, _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Nilify is a paid mutator transaction binding the contract method 0x2e80fbf3.
//
// Solidity: function nilify(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_Layerzero *LayerzeroTransactorSession) Nilify(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Layerzero.Contract.Nilify(&_Layerzero.TransactOpts, _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xa7229fd9.
//
// Solidity: function recoverToken(address _token, address _to, uint256 _amount) returns()
func (_Layerzero *LayerzeroTransactor) RecoverToken(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "recoverToken", _token, _to, _amount)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xa7229fd9.
//
// Solidity: function recoverToken(address _token, address _to, uint256 _amount) returns()
func (_Layerzero *LayerzeroSession) RecoverToken(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Layerzero.Contract.RecoverToken(&_Layerzero.TransactOpts, _token, _to, _amount)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xa7229fd9.
//
// Solidity: function recoverToken(address _token, address _to, uint256 _amount) returns()
func (_Layerzero *LayerzeroTransactorSession) RecoverToken(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Layerzero.Contract.RecoverToken(&_Layerzero.TransactOpts, _token, _to, _amount)
}

// RegisterLibrary is a paid mutator transaction binding the contract method 0xe8964e81.
//
// Solidity: function registerLibrary(address _lib) returns()
func (_Layerzero *LayerzeroTransactor) RegisterLibrary(opts *bind.TransactOpts, _lib common.Address) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "registerLibrary", _lib)
}

// RegisterLibrary is a paid mutator transaction binding the contract method 0xe8964e81.
//
// Solidity: function registerLibrary(address _lib) returns()
func (_Layerzero *LayerzeroSession) RegisterLibrary(_lib common.Address) (*types.Transaction, error) {
	return _Layerzero.Contract.RegisterLibrary(&_Layerzero.TransactOpts, _lib)
}

// RegisterLibrary is a paid mutator transaction binding the contract method 0xe8964e81.
//
// Solidity: function registerLibrary(address _lib) returns()
func (_Layerzero *LayerzeroTransactorSession) RegisterLibrary(_lib common.Address) (*types.Transaction, error) {
	return _Layerzero.Contract.RegisterLibrary(&_Layerzero.TransactOpts, _lib)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Layerzero *LayerzeroTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Layerzero *LayerzeroSession) RenounceOwnership() (*types.Transaction, error) {
	return _Layerzero.Contract.RenounceOwnership(&_Layerzero.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Layerzero *LayerzeroTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Layerzero.Contract.RenounceOwnership(&_Layerzero.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x2637a450.
//
// Solidity: function send((uint32,bytes32,bytes,bytes,bool) _params, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)))
func (_Layerzero *LayerzeroTransactor) Send(opts *bind.TransactOpts, _params MessagingParams, _refundAddress common.Address) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "send", _params, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0x2637a450.
//
// Solidity: function send((uint32,bytes32,bytes,bytes,bool) _params, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)))
func (_Layerzero *LayerzeroSession) Send(_params MessagingParams, _refundAddress common.Address) (*types.Transaction, error) {
	return _Layerzero.Contract.Send(&_Layerzero.TransactOpts, _params, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0x2637a450.
//
// Solidity: function send((uint32,bytes32,bytes,bytes,bool) _params, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)))
func (_Layerzero *LayerzeroTransactorSession) Send(_params MessagingParams, _refundAddress common.Address) (*types.Transaction, error) {
	return _Layerzero.Contract.Send(&_Layerzero.TransactOpts, _params, _refundAddress)
}

// SendCompose is a paid mutator transaction binding the contract method 0x7cb59012.
//
// Solidity: function sendCompose(address _to, bytes32 _guid, uint16 _index, bytes _message) returns()
func (_Layerzero *LayerzeroTransactor) SendCompose(opts *bind.TransactOpts, _to common.Address, _guid [32]byte, _index uint16, _message []byte) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "sendCompose", _to, _guid, _index, _message)
}

// SendCompose is a paid mutator transaction binding the contract method 0x7cb59012.
//
// Solidity: function sendCompose(address _to, bytes32 _guid, uint16 _index, bytes _message) returns()
func (_Layerzero *LayerzeroSession) SendCompose(_to common.Address, _guid [32]byte, _index uint16, _message []byte) (*types.Transaction, error) {
	return _Layerzero.Contract.SendCompose(&_Layerzero.TransactOpts, _to, _guid, _index, _message)
}

// SendCompose is a paid mutator transaction binding the contract method 0x7cb59012.
//
// Solidity: function sendCompose(address _to, bytes32 _guid, uint16 _index, bytes _message) returns()
func (_Layerzero *LayerzeroTransactorSession) SendCompose(_to common.Address, _guid [32]byte, _index uint16, _message []byte) (*types.Transaction, error) {
	return _Layerzero.Contract.SendCompose(&_Layerzero.TransactOpts, _to, _guid, _index, _message)
}

// SetConfig is a paid mutator transaction binding the contract method 0x6dbd9f90.
//
// Solidity: function setConfig(address _oapp, address _lib, (uint32,uint32,bytes)[] _params) returns()
func (_Layerzero *LayerzeroTransactor) SetConfig(opts *bind.TransactOpts, _oapp common.Address, _lib common.Address, _params []SetConfigParam) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "setConfig", _oapp, _lib, _params)
}

// SetConfig is a paid mutator transaction binding the contract method 0x6dbd9f90.
//
// Solidity: function setConfig(address _oapp, address _lib, (uint32,uint32,bytes)[] _params) returns()
func (_Layerzero *LayerzeroSession) SetConfig(_oapp common.Address, _lib common.Address, _params []SetConfigParam) (*types.Transaction, error) {
	return _Layerzero.Contract.SetConfig(&_Layerzero.TransactOpts, _oapp, _lib, _params)
}

// SetConfig is a paid mutator transaction binding the contract method 0x6dbd9f90.
//
// Solidity: function setConfig(address _oapp, address _lib, (uint32,uint32,bytes)[] _params) returns()
func (_Layerzero *LayerzeroTransactorSession) SetConfig(_oapp common.Address, _lib common.Address, _params []SetConfigParam) (*types.Transaction, error) {
	return _Layerzero.Contract.SetConfig(&_Layerzero.TransactOpts, _oapp, _lib, _params)
}

// SetDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0xa718531b.
//
// Solidity: function setDefaultReceiveLibrary(uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_Layerzero *LayerzeroTransactor) SetDefaultReceiveLibrary(opts *bind.TransactOpts, _eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "setDefaultReceiveLibrary", _eid, _newLib, _gracePeriod)
}

// SetDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0xa718531b.
//
// Solidity: function setDefaultReceiveLibrary(uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_Layerzero *LayerzeroSession) SetDefaultReceiveLibrary(_eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _Layerzero.Contract.SetDefaultReceiveLibrary(&_Layerzero.TransactOpts, _eid, _newLib, _gracePeriod)
}

// SetDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0xa718531b.
//
// Solidity: function setDefaultReceiveLibrary(uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_Layerzero *LayerzeroTransactorSession) SetDefaultReceiveLibrary(_eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _Layerzero.Contract.SetDefaultReceiveLibrary(&_Layerzero.TransactOpts, _eid, _newLib, _gracePeriod)
}

// SetDefaultReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0xd4b4ec8f.
//
// Solidity: function setDefaultReceiveLibraryTimeout(uint32 _eid, address _lib, uint256 _expiry) returns()
func (_Layerzero *LayerzeroTransactor) SetDefaultReceiveLibraryTimeout(opts *bind.TransactOpts, _eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "setDefaultReceiveLibraryTimeout", _eid, _lib, _expiry)
}

// SetDefaultReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0xd4b4ec8f.
//
// Solidity: function setDefaultReceiveLibraryTimeout(uint32 _eid, address _lib, uint256 _expiry) returns()
func (_Layerzero *LayerzeroSession) SetDefaultReceiveLibraryTimeout(_eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _Layerzero.Contract.SetDefaultReceiveLibraryTimeout(&_Layerzero.TransactOpts, _eid, _lib, _expiry)
}

// SetDefaultReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0xd4b4ec8f.
//
// Solidity: function setDefaultReceiveLibraryTimeout(uint32 _eid, address _lib, uint256 _expiry) returns()
func (_Layerzero *LayerzeroTransactorSession) SetDefaultReceiveLibraryTimeout(_eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _Layerzero.Contract.SetDefaultReceiveLibraryTimeout(&_Layerzero.TransactOpts, _eid, _lib, _expiry)
}

// SetDefaultSendLibrary is a paid mutator transaction binding the contract method 0xaafea312.
//
// Solidity: function setDefaultSendLibrary(uint32 _eid, address _newLib) returns()
func (_Layerzero *LayerzeroTransactor) SetDefaultSendLibrary(opts *bind.TransactOpts, _eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "setDefaultSendLibrary", _eid, _newLib)
}

// SetDefaultSendLibrary is a paid mutator transaction binding the contract method 0xaafea312.
//
// Solidity: function setDefaultSendLibrary(uint32 _eid, address _newLib) returns()
func (_Layerzero *LayerzeroSession) SetDefaultSendLibrary(_eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _Layerzero.Contract.SetDefaultSendLibrary(&_Layerzero.TransactOpts, _eid, _newLib)
}

// SetDefaultSendLibrary is a paid mutator transaction binding the contract method 0xaafea312.
//
// Solidity: function setDefaultSendLibrary(uint32 _eid, address _newLib) returns()
func (_Layerzero *LayerzeroTransactorSession) SetDefaultSendLibrary(_eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _Layerzero.Contract.SetDefaultSendLibrary(&_Layerzero.TransactOpts, _eid, _newLib)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Layerzero *LayerzeroTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Layerzero *LayerzeroSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Layerzero.Contract.SetDelegate(&_Layerzero.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Layerzero *LayerzeroTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Layerzero.Contract.SetDelegate(&_Layerzero.TransactOpts, _delegate)
}

// SetLzToken is a paid mutator transaction binding the contract method 0xc28e0eed.
//
// Solidity: function setLzToken(address _lzToken) returns()
func (_Layerzero *LayerzeroTransactor) SetLzToken(opts *bind.TransactOpts, _lzToken common.Address) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "setLzToken", _lzToken)
}

// SetLzToken is a paid mutator transaction binding the contract method 0xc28e0eed.
//
// Solidity: function setLzToken(address _lzToken) returns()
func (_Layerzero *LayerzeroSession) SetLzToken(_lzToken common.Address) (*types.Transaction, error) {
	return _Layerzero.Contract.SetLzToken(&_Layerzero.TransactOpts, _lzToken)
}

// SetLzToken is a paid mutator transaction binding the contract method 0xc28e0eed.
//
// Solidity: function setLzToken(address _lzToken) returns()
func (_Layerzero *LayerzeroTransactorSession) SetLzToken(_lzToken common.Address) (*types.Transaction, error) {
	return _Layerzero.Contract.SetLzToken(&_Layerzero.TransactOpts, _lzToken)
}

// SetReceiveLibrary is a paid mutator transaction binding the contract method 0x6a14d715.
//
// Solidity: function setReceiveLibrary(address _oapp, uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_Layerzero *LayerzeroTransactor) SetReceiveLibrary(opts *bind.TransactOpts, _oapp common.Address, _eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "setReceiveLibrary", _oapp, _eid, _newLib, _gracePeriod)
}

// SetReceiveLibrary is a paid mutator transaction binding the contract method 0x6a14d715.
//
// Solidity: function setReceiveLibrary(address _oapp, uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_Layerzero *LayerzeroSession) SetReceiveLibrary(_oapp common.Address, _eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _Layerzero.Contract.SetReceiveLibrary(&_Layerzero.TransactOpts, _oapp, _eid, _newLib, _gracePeriod)
}

// SetReceiveLibrary is a paid mutator transaction binding the contract method 0x6a14d715.
//
// Solidity: function setReceiveLibrary(address _oapp, uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_Layerzero *LayerzeroTransactorSession) SetReceiveLibrary(_oapp common.Address, _eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _Layerzero.Contract.SetReceiveLibrary(&_Layerzero.TransactOpts, _oapp, _eid, _newLib, _gracePeriod)
}

// SetReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0x183c834f.
//
// Solidity: function setReceiveLibraryTimeout(address _oapp, uint32 _eid, address _lib, uint256 _expiry) returns()
func (_Layerzero *LayerzeroTransactor) SetReceiveLibraryTimeout(opts *bind.TransactOpts, _oapp common.Address, _eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "setReceiveLibraryTimeout", _oapp, _eid, _lib, _expiry)
}

// SetReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0x183c834f.
//
// Solidity: function setReceiveLibraryTimeout(address _oapp, uint32 _eid, address _lib, uint256 _expiry) returns()
func (_Layerzero *LayerzeroSession) SetReceiveLibraryTimeout(_oapp common.Address, _eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _Layerzero.Contract.SetReceiveLibraryTimeout(&_Layerzero.TransactOpts, _oapp, _eid, _lib, _expiry)
}

// SetReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0x183c834f.
//
// Solidity: function setReceiveLibraryTimeout(address _oapp, uint32 _eid, address _lib, uint256 _expiry) returns()
func (_Layerzero *LayerzeroTransactorSession) SetReceiveLibraryTimeout(_oapp common.Address, _eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _Layerzero.Contract.SetReceiveLibraryTimeout(&_Layerzero.TransactOpts, _oapp, _eid, _lib, _expiry)
}

// SetSendLibrary is a paid mutator transaction binding the contract method 0x9535ff30.
//
// Solidity: function setSendLibrary(address _oapp, uint32 _eid, address _newLib) returns()
func (_Layerzero *LayerzeroTransactor) SetSendLibrary(opts *bind.TransactOpts, _oapp common.Address, _eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "setSendLibrary", _oapp, _eid, _newLib)
}

// SetSendLibrary is a paid mutator transaction binding the contract method 0x9535ff30.
//
// Solidity: function setSendLibrary(address _oapp, uint32 _eid, address _newLib) returns()
func (_Layerzero *LayerzeroSession) SetSendLibrary(_oapp common.Address, _eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _Layerzero.Contract.SetSendLibrary(&_Layerzero.TransactOpts, _oapp, _eid, _newLib)
}

// SetSendLibrary is a paid mutator transaction binding the contract method 0x9535ff30.
//
// Solidity: function setSendLibrary(address _oapp, uint32 _eid, address _newLib) returns()
func (_Layerzero *LayerzeroTransactorSession) SetSendLibrary(_oapp common.Address, _eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _Layerzero.Contract.SetSendLibrary(&_Layerzero.TransactOpts, _oapp, _eid, _newLib)
}

// Skip is a paid mutator transaction binding the contract method 0xd70b8902.
//
// Solidity: function skip(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce) returns()
func (_Layerzero *LayerzeroTransactor) Skip(opts *bind.TransactOpts, _oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "skip", _oapp, _srcEid, _sender, _nonce)
}

// Skip is a paid mutator transaction binding the contract method 0xd70b8902.
//
// Solidity: function skip(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce) returns()
func (_Layerzero *LayerzeroSession) Skip(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64) (*types.Transaction, error) {
	return _Layerzero.Contract.Skip(&_Layerzero.TransactOpts, _oapp, _srcEid, _sender, _nonce)
}

// Skip is a paid mutator transaction binding the contract method 0xd70b8902.
//
// Solidity: function skip(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce) returns()
func (_Layerzero *LayerzeroTransactorSession) Skip(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64) (*types.Transaction, error) {
	return _Layerzero.Contract.Skip(&_Layerzero.TransactOpts, _oapp, _srcEid, _sender, _nonce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Layerzero *LayerzeroTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Layerzero *LayerzeroSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Layerzero.Contract.TransferOwnership(&_Layerzero.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Layerzero *LayerzeroTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Layerzero.Contract.TransferOwnership(&_Layerzero.TransactOpts, newOwner)
}

// Verify is a paid mutator transaction binding the contract method 0xa825d747.
//
// Solidity: function verify((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) returns()
func (_Layerzero *LayerzeroTransactor) Verify(opts *bind.TransactOpts, _origin Origin, _receiver common.Address, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Layerzero.contract.Transact(opts, "verify", _origin, _receiver, _payloadHash)
}

// Verify is a paid mutator transaction binding the contract method 0xa825d747.
//
// Solidity: function verify((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) returns()
func (_Layerzero *LayerzeroSession) Verify(_origin Origin, _receiver common.Address, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Layerzero.Contract.Verify(&_Layerzero.TransactOpts, _origin, _receiver, _payloadHash)
}

// Verify is a paid mutator transaction binding the contract method 0xa825d747.
//
// Solidity: function verify((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) returns()
func (_Layerzero *LayerzeroTransactorSession) Verify(_origin Origin, _receiver common.Address, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Layerzero.Contract.Verify(&_Layerzero.TransactOpts, _origin, _receiver, _payloadHash)
}

// LayerzeroComposeDeliveredIterator is returned from FilterComposeDelivered and is used to iterate over the raw logs and unpacked data for ComposeDelivered events raised by the Layerzero contract.
type LayerzeroComposeDeliveredIterator struct {
	Event *LayerzeroComposeDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroComposeDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroComposeDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroComposeDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroComposeDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroComposeDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroComposeDelivered represents a ComposeDelivered event raised by the Layerzero contract.
type LayerzeroComposeDelivered struct {
	From  common.Address
	To    common.Address
	Guid  [32]byte
	Index uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterComposeDelivered is a free log retrieval operation binding the contract event 0x0036c98efcf9e6641dfbc9051f66f405253e8e0c2ab4a24dccda15595b7378c8.
//
// Solidity: event ComposeDelivered(address from, address to, bytes32 guid, uint16 index)
func (_Layerzero *LayerzeroFilterer) FilterComposeDelivered(opts *bind.FilterOpts) (*LayerzeroComposeDeliveredIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "ComposeDelivered")
	if err != nil {
		return nil, err
	}
	return &LayerzeroComposeDeliveredIterator{contract: _Layerzero.contract, event: "ComposeDelivered", logs: logs, sub: sub}, nil
}

// WatchComposeDelivered is a free log subscription operation binding the contract event 0x0036c98efcf9e6641dfbc9051f66f405253e8e0c2ab4a24dccda15595b7378c8.
//
// Solidity: event ComposeDelivered(address from, address to, bytes32 guid, uint16 index)
func (_Layerzero *LayerzeroFilterer) WatchComposeDelivered(opts *bind.WatchOpts, sink chan<- *LayerzeroComposeDelivered) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "ComposeDelivered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroComposeDelivered)
				if err := _Layerzero.contract.UnpackLog(event, "ComposeDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseComposeDelivered is a log parse operation binding the contract event 0x0036c98efcf9e6641dfbc9051f66f405253e8e0c2ab4a24dccda15595b7378c8.
//
// Solidity: event ComposeDelivered(address from, address to, bytes32 guid, uint16 index)
func (_Layerzero *LayerzeroFilterer) ParseComposeDelivered(log types.Log) (*LayerzeroComposeDelivered, error) {
	event := new(LayerzeroComposeDelivered)
	if err := _Layerzero.contract.UnpackLog(event, "ComposeDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroComposeSentIterator is returned from FilterComposeSent and is used to iterate over the raw logs and unpacked data for ComposeSent events raised by the Layerzero contract.
type LayerzeroComposeSentIterator struct {
	Event *LayerzeroComposeSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroComposeSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroComposeSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroComposeSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroComposeSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroComposeSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroComposeSent represents a ComposeSent event raised by the Layerzero contract.
type LayerzeroComposeSent struct {
	From    common.Address
	To      common.Address
	Guid    [32]byte
	Index   uint16
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterComposeSent is a free log retrieval operation binding the contract event 0x3d52ff888d033fd3dd1d8057da59e850c91d91a72c41dfa445b247dfedeb6dc1.
//
// Solidity: event ComposeSent(address from, address to, bytes32 guid, uint16 index, bytes message)
func (_Layerzero *LayerzeroFilterer) FilterComposeSent(opts *bind.FilterOpts) (*LayerzeroComposeSentIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "ComposeSent")
	if err != nil {
		return nil, err
	}
	return &LayerzeroComposeSentIterator{contract: _Layerzero.contract, event: "ComposeSent", logs: logs, sub: sub}, nil
}

// WatchComposeSent is a free log subscription operation binding the contract event 0x3d52ff888d033fd3dd1d8057da59e850c91d91a72c41dfa445b247dfedeb6dc1.
//
// Solidity: event ComposeSent(address from, address to, bytes32 guid, uint16 index, bytes message)
func (_Layerzero *LayerzeroFilterer) WatchComposeSent(opts *bind.WatchOpts, sink chan<- *LayerzeroComposeSent) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "ComposeSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroComposeSent)
				if err := _Layerzero.contract.UnpackLog(event, "ComposeSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseComposeSent is a log parse operation binding the contract event 0x3d52ff888d033fd3dd1d8057da59e850c91d91a72c41dfa445b247dfedeb6dc1.
//
// Solidity: event ComposeSent(address from, address to, bytes32 guid, uint16 index, bytes message)
func (_Layerzero *LayerzeroFilterer) ParseComposeSent(log types.Log) (*LayerzeroComposeSent, error) {
	event := new(LayerzeroComposeSent)
	if err := _Layerzero.contract.UnpackLog(event, "ComposeSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroDefaultReceiveLibrarySetIterator is returned from FilterDefaultReceiveLibrarySet and is used to iterate over the raw logs and unpacked data for DefaultReceiveLibrarySet events raised by the Layerzero contract.
type LayerzeroDefaultReceiveLibrarySetIterator struct {
	Event *LayerzeroDefaultReceiveLibrarySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroDefaultReceiveLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroDefaultReceiveLibrarySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroDefaultReceiveLibrarySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroDefaultReceiveLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroDefaultReceiveLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroDefaultReceiveLibrarySet represents a DefaultReceiveLibrarySet event raised by the Layerzero contract.
type LayerzeroDefaultReceiveLibrarySet struct {
	Eid    uint32
	NewLib common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDefaultReceiveLibrarySet is a free log retrieval operation binding the contract event 0xc16891855cffb4a5ac51ac11864a3f3c96ba816cc45fe686c987ae36277de5ec.
//
// Solidity: event DefaultReceiveLibrarySet(uint32 eid, address newLib)
func (_Layerzero *LayerzeroFilterer) FilterDefaultReceiveLibrarySet(opts *bind.FilterOpts) (*LayerzeroDefaultReceiveLibrarySetIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "DefaultReceiveLibrarySet")
	if err != nil {
		return nil, err
	}
	return &LayerzeroDefaultReceiveLibrarySetIterator{contract: _Layerzero.contract, event: "DefaultReceiveLibrarySet", logs: logs, sub: sub}, nil
}

// WatchDefaultReceiveLibrarySet is a free log subscription operation binding the contract event 0xc16891855cffb4a5ac51ac11864a3f3c96ba816cc45fe686c987ae36277de5ec.
//
// Solidity: event DefaultReceiveLibrarySet(uint32 eid, address newLib)
func (_Layerzero *LayerzeroFilterer) WatchDefaultReceiveLibrarySet(opts *bind.WatchOpts, sink chan<- *LayerzeroDefaultReceiveLibrarySet) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "DefaultReceiveLibrarySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroDefaultReceiveLibrarySet)
				if err := _Layerzero.contract.UnpackLog(event, "DefaultReceiveLibrarySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultReceiveLibrarySet is a log parse operation binding the contract event 0xc16891855cffb4a5ac51ac11864a3f3c96ba816cc45fe686c987ae36277de5ec.
//
// Solidity: event DefaultReceiveLibrarySet(uint32 eid, address newLib)
func (_Layerzero *LayerzeroFilterer) ParseDefaultReceiveLibrarySet(log types.Log) (*LayerzeroDefaultReceiveLibrarySet, error) {
	event := new(LayerzeroDefaultReceiveLibrarySet)
	if err := _Layerzero.contract.UnpackLog(event, "DefaultReceiveLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroDefaultReceiveLibraryTimeoutSetIterator is returned from FilterDefaultReceiveLibraryTimeoutSet and is used to iterate over the raw logs and unpacked data for DefaultReceiveLibraryTimeoutSet events raised by the Layerzero contract.
type LayerzeroDefaultReceiveLibraryTimeoutSetIterator struct {
	Event *LayerzeroDefaultReceiveLibraryTimeoutSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroDefaultReceiveLibraryTimeoutSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroDefaultReceiveLibraryTimeoutSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroDefaultReceiveLibraryTimeoutSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroDefaultReceiveLibraryTimeoutSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroDefaultReceiveLibraryTimeoutSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroDefaultReceiveLibraryTimeoutSet represents a DefaultReceiveLibraryTimeoutSet event raised by the Layerzero contract.
type LayerzeroDefaultReceiveLibraryTimeoutSet struct {
	Eid    uint32
	OldLib common.Address
	Expiry *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDefaultReceiveLibraryTimeoutSet is a free log retrieval operation binding the contract event 0x55b28633cdb29709386f555dfc54418592ad475ce7a65a78ac5928af60ffb8f8.
//
// Solidity: event DefaultReceiveLibraryTimeoutSet(uint32 eid, address oldLib, uint256 expiry)
func (_Layerzero *LayerzeroFilterer) FilterDefaultReceiveLibraryTimeoutSet(opts *bind.FilterOpts) (*LayerzeroDefaultReceiveLibraryTimeoutSetIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "DefaultReceiveLibraryTimeoutSet")
	if err != nil {
		return nil, err
	}
	return &LayerzeroDefaultReceiveLibraryTimeoutSetIterator{contract: _Layerzero.contract, event: "DefaultReceiveLibraryTimeoutSet", logs: logs, sub: sub}, nil
}

// WatchDefaultReceiveLibraryTimeoutSet is a free log subscription operation binding the contract event 0x55b28633cdb29709386f555dfc54418592ad475ce7a65a78ac5928af60ffb8f8.
//
// Solidity: event DefaultReceiveLibraryTimeoutSet(uint32 eid, address oldLib, uint256 expiry)
func (_Layerzero *LayerzeroFilterer) WatchDefaultReceiveLibraryTimeoutSet(opts *bind.WatchOpts, sink chan<- *LayerzeroDefaultReceiveLibraryTimeoutSet) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "DefaultReceiveLibraryTimeoutSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroDefaultReceiveLibraryTimeoutSet)
				if err := _Layerzero.contract.UnpackLog(event, "DefaultReceiveLibraryTimeoutSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultReceiveLibraryTimeoutSet is a log parse operation binding the contract event 0x55b28633cdb29709386f555dfc54418592ad475ce7a65a78ac5928af60ffb8f8.
//
// Solidity: event DefaultReceiveLibraryTimeoutSet(uint32 eid, address oldLib, uint256 expiry)
func (_Layerzero *LayerzeroFilterer) ParseDefaultReceiveLibraryTimeoutSet(log types.Log) (*LayerzeroDefaultReceiveLibraryTimeoutSet, error) {
	event := new(LayerzeroDefaultReceiveLibraryTimeoutSet)
	if err := _Layerzero.contract.UnpackLog(event, "DefaultReceiveLibraryTimeoutSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroDefaultSendLibrarySetIterator is returned from FilterDefaultSendLibrarySet and is used to iterate over the raw logs and unpacked data for DefaultSendLibrarySet events raised by the Layerzero contract.
type LayerzeroDefaultSendLibrarySetIterator struct {
	Event *LayerzeroDefaultSendLibrarySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroDefaultSendLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroDefaultSendLibrarySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroDefaultSendLibrarySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroDefaultSendLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroDefaultSendLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroDefaultSendLibrarySet represents a DefaultSendLibrarySet event raised by the Layerzero contract.
type LayerzeroDefaultSendLibrarySet struct {
	Eid    uint32
	NewLib common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDefaultSendLibrarySet is a free log retrieval operation binding the contract event 0x16aa0f528038ab41019e95bae5b418a50ba8532c5800e3b7ea2f517d3fa625f5.
//
// Solidity: event DefaultSendLibrarySet(uint32 eid, address newLib)
func (_Layerzero *LayerzeroFilterer) FilterDefaultSendLibrarySet(opts *bind.FilterOpts) (*LayerzeroDefaultSendLibrarySetIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "DefaultSendLibrarySet")
	if err != nil {
		return nil, err
	}
	return &LayerzeroDefaultSendLibrarySetIterator{contract: _Layerzero.contract, event: "DefaultSendLibrarySet", logs: logs, sub: sub}, nil
}

// WatchDefaultSendLibrarySet is a free log subscription operation binding the contract event 0x16aa0f528038ab41019e95bae5b418a50ba8532c5800e3b7ea2f517d3fa625f5.
//
// Solidity: event DefaultSendLibrarySet(uint32 eid, address newLib)
func (_Layerzero *LayerzeroFilterer) WatchDefaultSendLibrarySet(opts *bind.WatchOpts, sink chan<- *LayerzeroDefaultSendLibrarySet) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "DefaultSendLibrarySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroDefaultSendLibrarySet)
				if err := _Layerzero.contract.UnpackLog(event, "DefaultSendLibrarySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultSendLibrarySet is a log parse operation binding the contract event 0x16aa0f528038ab41019e95bae5b418a50ba8532c5800e3b7ea2f517d3fa625f5.
//
// Solidity: event DefaultSendLibrarySet(uint32 eid, address newLib)
func (_Layerzero *LayerzeroFilterer) ParseDefaultSendLibrarySet(log types.Log) (*LayerzeroDefaultSendLibrarySet, error) {
	event := new(LayerzeroDefaultSendLibrarySet)
	if err := _Layerzero.contract.UnpackLog(event, "DefaultSendLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroDelegateSetIterator is returned from FilterDelegateSet and is used to iterate over the raw logs and unpacked data for DelegateSet events raised by the Layerzero contract.
type LayerzeroDelegateSetIterator struct {
	Event *LayerzeroDelegateSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroDelegateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroDelegateSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroDelegateSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroDelegateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroDelegateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroDelegateSet represents a DelegateSet event raised by the Layerzero contract.
type LayerzeroDelegateSet struct {
	Sender   common.Address
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegateSet is a free log retrieval operation binding the contract event 0x6ee10e9ed4d6ce9742703a498707862f4b00f1396a87195eb93267b3d7983981.
//
// Solidity: event DelegateSet(address sender, address delegate)
func (_Layerzero *LayerzeroFilterer) FilterDelegateSet(opts *bind.FilterOpts) (*LayerzeroDelegateSetIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "DelegateSet")
	if err != nil {
		return nil, err
	}
	return &LayerzeroDelegateSetIterator{contract: _Layerzero.contract, event: "DelegateSet", logs: logs, sub: sub}, nil
}

// WatchDelegateSet is a free log subscription operation binding the contract event 0x6ee10e9ed4d6ce9742703a498707862f4b00f1396a87195eb93267b3d7983981.
//
// Solidity: event DelegateSet(address sender, address delegate)
func (_Layerzero *LayerzeroFilterer) WatchDelegateSet(opts *bind.WatchOpts, sink chan<- *LayerzeroDelegateSet) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "DelegateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroDelegateSet)
				if err := _Layerzero.contract.UnpackLog(event, "DelegateSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateSet is a log parse operation binding the contract event 0x6ee10e9ed4d6ce9742703a498707862f4b00f1396a87195eb93267b3d7983981.
//
// Solidity: event DelegateSet(address sender, address delegate)
func (_Layerzero *LayerzeroFilterer) ParseDelegateSet(log types.Log) (*LayerzeroDelegateSet, error) {
	event := new(LayerzeroDelegateSet)
	if err := _Layerzero.contract.UnpackLog(event, "DelegateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroInboundNonceSkippedIterator is returned from FilterInboundNonceSkipped and is used to iterate over the raw logs and unpacked data for InboundNonceSkipped events raised by the Layerzero contract.
type LayerzeroInboundNonceSkippedIterator struct {
	Event *LayerzeroInboundNonceSkipped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroInboundNonceSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroInboundNonceSkipped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroInboundNonceSkipped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroInboundNonceSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroInboundNonceSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroInboundNonceSkipped represents a InboundNonceSkipped event raised by the Layerzero contract.
type LayerzeroInboundNonceSkipped struct {
	SrcEid   uint32
	Sender   [32]byte
	Receiver common.Address
	Nonce    uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInboundNonceSkipped is a free log retrieval operation binding the contract event 0x28f40053783033ef755556a0c3315379141f51a33aed8334174ffbadd90bde48.
//
// Solidity: event InboundNonceSkipped(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce)
func (_Layerzero *LayerzeroFilterer) FilterInboundNonceSkipped(opts *bind.FilterOpts) (*LayerzeroInboundNonceSkippedIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "InboundNonceSkipped")
	if err != nil {
		return nil, err
	}
	return &LayerzeroInboundNonceSkippedIterator{contract: _Layerzero.contract, event: "InboundNonceSkipped", logs: logs, sub: sub}, nil
}

// WatchInboundNonceSkipped is a free log subscription operation binding the contract event 0x28f40053783033ef755556a0c3315379141f51a33aed8334174ffbadd90bde48.
//
// Solidity: event InboundNonceSkipped(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce)
func (_Layerzero *LayerzeroFilterer) WatchInboundNonceSkipped(opts *bind.WatchOpts, sink chan<- *LayerzeroInboundNonceSkipped) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "InboundNonceSkipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroInboundNonceSkipped)
				if err := _Layerzero.contract.UnpackLog(event, "InboundNonceSkipped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInboundNonceSkipped is a log parse operation binding the contract event 0x28f40053783033ef755556a0c3315379141f51a33aed8334174ffbadd90bde48.
//
// Solidity: event InboundNonceSkipped(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce)
func (_Layerzero *LayerzeroFilterer) ParseInboundNonceSkipped(log types.Log) (*LayerzeroInboundNonceSkipped, error) {
	event := new(LayerzeroInboundNonceSkipped)
	if err := _Layerzero.contract.UnpackLog(event, "InboundNonceSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroLibraryRegisteredIterator is returned from FilterLibraryRegistered and is used to iterate over the raw logs and unpacked data for LibraryRegistered events raised by the Layerzero contract.
type LayerzeroLibraryRegisteredIterator struct {
	Event *LayerzeroLibraryRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroLibraryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroLibraryRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroLibraryRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroLibraryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroLibraryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroLibraryRegistered represents a LibraryRegistered event raised by the Layerzero contract.
type LayerzeroLibraryRegistered struct {
	NewLib common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLibraryRegistered is a free log retrieval operation binding the contract event 0x6b374d56679ca9463f27c85c6311e2bb7fde69bf201d3da39d53f10bd9d78af5.
//
// Solidity: event LibraryRegistered(address newLib)
func (_Layerzero *LayerzeroFilterer) FilterLibraryRegistered(opts *bind.FilterOpts) (*LayerzeroLibraryRegisteredIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "LibraryRegistered")
	if err != nil {
		return nil, err
	}
	return &LayerzeroLibraryRegisteredIterator{contract: _Layerzero.contract, event: "LibraryRegistered", logs: logs, sub: sub}, nil
}

// WatchLibraryRegistered is a free log subscription operation binding the contract event 0x6b374d56679ca9463f27c85c6311e2bb7fde69bf201d3da39d53f10bd9d78af5.
//
// Solidity: event LibraryRegistered(address newLib)
func (_Layerzero *LayerzeroFilterer) WatchLibraryRegistered(opts *bind.WatchOpts, sink chan<- *LayerzeroLibraryRegistered) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "LibraryRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroLibraryRegistered)
				if err := _Layerzero.contract.UnpackLog(event, "LibraryRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLibraryRegistered is a log parse operation binding the contract event 0x6b374d56679ca9463f27c85c6311e2bb7fde69bf201d3da39d53f10bd9d78af5.
//
// Solidity: event LibraryRegistered(address newLib)
func (_Layerzero *LayerzeroFilterer) ParseLibraryRegistered(log types.Log) (*LayerzeroLibraryRegistered, error) {
	event := new(LayerzeroLibraryRegistered)
	if err := _Layerzero.contract.UnpackLog(event, "LibraryRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroLzComposeAlertIterator is returned from FilterLzComposeAlert and is used to iterate over the raw logs and unpacked data for LzComposeAlert events raised by the Layerzero contract.
type LayerzeroLzComposeAlertIterator struct {
	Event *LayerzeroLzComposeAlert // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroLzComposeAlertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroLzComposeAlert)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroLzComposeAlert)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroLzComposeAlertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroLzComposeAlertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroLzComposeAlert represents a LzComposeAlert event raised by the Layerzero contract.
type LayerzeroLzComposeAlert struct {
	From      common.Address
	To        common.Address
	Executor  common.Address
	Guid      [32]byte
	Index     uint16
	Gas       *big.Int
	Value     *big.Int
	Message   []byte
	ExtraData []byte
	Reason    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLzComposeAlert is a free log retrieval operation binding the contract event 0x8a0b1dce321c5c5fb42349bce46d18087c04140de520917661fb923e44a904b9.
//
// Solidity: event LzComposeAlert(address indexed from, address indexed to, address indexed executor, bytes32 guid, uint16 index, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_Layerzero *LayerzeroFilterer) FilterLzComposeAlert(opts *bind.FilterOpts, from []common.Address, to []common.Address, executor []common.Address) (*LayerzeroLzComposeAlertIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "LzComposeAlert", fromRule, toRule, executorRule)
	if err != nil {
		return nil, err
	}
	return &LayerzeroLzComposeAlertIterator{contract: _Layerzero.contract, event: "LzComposeAlert", logs: logs, sub: sub}, nil
}

// WatchLzComposeAlert is a free log subscription operation binding the contract event 0x8a0b1dce321c5c5fb42349bce46d18087c04140de520917661fb923e44a904b9.
//
// Solidity: event LzComposeAlert(address indexed from, address indexed to, address indexed executor, bytes32 guid, uint16 index, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_Layerzero *LayerzeroFilterer) WatchLzComposeAlert(opts *bind.WatchOpts, sink chan<- *LayerzeroLzComposeAlert, from []common.Address, to []common.Address, executor []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "LzComposeAlert", fromRule, toRule, executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroLzComposeAlert)
				if err := _Layerzero.contract.UnpackLog(event, "LzComposeAlert", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLzComposeAlert is a log parse operation binding the contract event 0x8a0b1dce321c5c5fb42349bce46d18087c04140de520917661fb923e44a904b9.
//
// Solidity: event LzComposeAlert(address indexed from, address indexed to, address indexed executor, bytes32 guid, uint16 index, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_Layerzero *LayerzeroFilterer) ParseLzComposeAlert(log types.Log) (*LayerzeroLzComposeAlert, error) {
	event := new(LayerzeroLzComposeAlert)
	if err := _Layerzero.contract.UnpackLog(event, "LzComposeAlert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroLzReceiveAlertIterator is returned from FilterLzReceiveAlert and is used to iterate over the raw logs and unpacked data for LzReceiveAlert events raised by the Layerzero contract.
type LayerzeroLzReceiveAlertIterator struct {
	Event *LayerzeroLzReceiveAlert // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroLzReceiveAlertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroLzReceiveAlert)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroLzReceiveAlert)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroLzReceiveAlertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroLzReceiveAlertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroLzReceiveAlert represents a LzReceiveAlert event raised by the Layerzero contract.
type LayerzeroLzReceiveAlert struct {
	Receiver  common.Address
	Executor  common.Address
	Origin    Origin
	Guid      [32]byte
	Gas       *big.Int
	Value     *big.Int
	Message   []byte
	ExtraData []byte
	Reason    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLzReceiveAlert is a free log retrieval operation binding the contract event 0x7edfa10fe10193301ad8a8bea7e968c7bcabcc64981f368e3aeada40ce26ae2c.
//
// Solidity: event LzReceiveAlert(address indexed receiver, address indexed executor, (uint32,bytes32,uint64) origin, bytes32 guid, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_Layerzero *LayerzeroFilterer) FilterLzReceiveAlert(opts *bind.FilterOpts, receiver []common.Address, executor []common.Address) (*LayerzeroLzReceiveAlertIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "LzReceiveAlert", receiverRule, executorRule)
	if err != nil {
		return nil, err
	}
	return &LayerzeroLzReceiveAlertIterator{contract: _Layerzero.contract, event: "LzReceiveAlert", logs: logs, sub: sub}, nil
}

// WatchLzReceiveAlert is a free log subscription operation binding the contract event 0x7edfa10fe10193301ad8a8bea7e968c7bcabcc64981f368e3aeada40ce26ae2c.
//
// Solidity: event LzReceiveAlert(address indexed receiver, address indexed executor, (uint32,bytes32,uint64) origin, bytes32 guid, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_Layerzero *LayerzeroFilterer) WatchLzReceiveAlert(opts *bind.WatchOpts, sink chan<- *LayerzeroLzReceiveAlert, receiver []common.Address, executor []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "LzReceiveAlert", receiverRule, executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroLzReceiveAlert)
				if err := _Layerzero.contract.UnpackLog(event, "LzReceiveAlert", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLzReceiveAlert is a log parse operation binding the contract event 0x7edfa10fe10193301ad8a8bea7e968c7bcabcc64981f368e3aeada40ce26ae2c.
//
// Solidity: event LzReceiveAlert(address indexed receiver, address indexed executor, (uint32,bytes32,uint64) origin, bytes32 guid, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_Layerzero *LayerzeroFilterer) ParseLzReceiveAlert(log types.Log) (*LayerzeroLzReceiveAlert, error) {
	event := new(LayerzeroLzReceiveAlert)
	if err := _Layerzero.contract.UnpackLog(event, "LzReceiveAlert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroLzTokenSetIterator is returned from FilterLzTokenSet and is used to iterate over the raw logs and unpacked data for LzTokenSet events raised by the Layerzero contract.
type LayerzeroLzTokenSetIterator struct {
	Event *LayerzeroLzTokenSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroLzTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroLzTokenSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroLzTokenSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroLzTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroLzTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroLzTokenSet represents a LzTokenSet event raised by the Layerzero contract.
type LayerzeroLzTokenSet struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLzTokenSet is a free log retrieval operation binding the contract event 0xd476ec5ec1ac11cec3714d41e7ea49419471aceb9bd0dff1becfc3e363a62396.
//
// Solidity: event LzTokenSet(address token)
func (_Layerzero *LayerzeroFilterer) FilterLzTokenSet(opts *bind.FilterOpts) (*LayerzeroLzTokenSetIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "LzTokenSet")
	if err != nil {
		return nil, err
	}
	return &LayerzeroLzTokenSetIterator{contract: _Layerzero.contract, event: "LzTokenSet", logs: logs, sub: sub}, nil
}

// WatchLzTokenSet is a free log subscription operation binding the contract event 0xd476ec5ec1ac11cec3714d41e7ea49419471aceb9bd0dff1becfc3e363a62396.
//
// Solidity: event LzTokenSet(address token)
func (_Layerzero *LayerzeroFilterer) WatchLzTokenSet(opts *bind.WatchOpts, sink chan<- *LayerzeroLzTokenSet) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "LzTokenSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroLzTokenSet)
				if err := _Layerzero.contract.UnpackLog(event, "LzTokenSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLzTokenSet is a log parse operation binding the contract event 0xd476ec5ec1ac11cec3714d41e7ea49419471aceb9bd0dff1becfc3e363a62396.
//
// Solidity: event LzTokenSet(address token)
func (_Layerzero *LayerzeroFilterer) ParseLzTokenSet(log types.Log) (*LayerzeroLzTokenSet, error) {
	event := new(LayerzeroLzTokenSet)
	if err := _Layerzero.contract.UnpackLog(event, "LzTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Layerzero contract.
type LayerzeroOwnershipTransferredIterator struct {
	Event *LayerzeroOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroOwnershipTransferred represents a OwnershipTransferred event raised by the Layerzero contract.
type LayerzeroOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Layerzero *LayerzeroFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LayerzeroOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LayerzeroOwnershipTransferredIterator{contract: _Layerzero.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Layerzero *LayerzeroFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LayerzeroOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroOwnershipTransferred)
				if err := _Layerzero.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Layerzero *LayerzeroFilterer) ParseOwnershipTransferred(log types.Log) (*LayerzeroOwnershipTransferred, error) {
	event := new(LayerzeroOwnershipTransferred)
	if err := _Layerzero.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroPacketBurntIterator is returned from FilterPacketBurnt and is used to iterate over the raw logs and unpacked data for PacketBurnt events raised by the Layerzero contract.
type LayerzeroPacketBurntIterator struct {
	Event *LayerzeroPacketBurnt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroPacketBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroPacketBurnt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroPacketBurnt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroPacketBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroPacketBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroPacketBurnt represents a PacketBurnt event raised by the Layerzero contract.
type LayerzeroPacketBurnt struct {
	SrcEid      uint32
	Sender      [32]byte
	Receiver    common.Address
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPacketBurnt is a free log retrieval operation binding the contract event 0x7f68a37a6e69a0de35024a234558f9efe4b33b58657753d21eaaa82d51c3510e.
//
// Solidity: event PacketBurnt(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_Layerzero *LayerzeroFilterer) FilterPacketBurnt(opts *bind.FilterOpts) (*LayerzeroPacketBurntIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "PacketBurnt")
	if err != nil {
		return nil, err
	}
	return &LayerzeroPacketBurntIterator{contract: _Layerzero.contract, event: "PacketBurnt", logs: logs, sub: sub}, nil
}

// WatchPacketBurnt is a free log subscription operation binding the contract event 0x7f68a37a6e69a0de35024a234558f9efe4b33b58657753d21eaaa82d51c3510e.
//
// Solidity: event PacketBurnt(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_Layerzero *LayerzeroFilterer) WatchPacketBurnt(opts *bind.WatchOpts, sink chan<- *LayerzeroPacketBurnt) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "PacketBurnt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroPacketBurnt)
				if err := _Layerzero.contract.UnpackLog(event, "PacketBurnt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketBurnt is a log parse operation binding the contract event 0x7f68a37a6e69a0de35024a234558f9efe4b33b58657753d21eaaa82d51c3510e.
//
// Solidity: event PacketBurnt(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_Layerzero *LayerzeroFilterer) ParsePacketBurnt(log types.Log) (*LayerzeroPacketBurnt, error) {
	event := new(LayerzeroPacketBurnt)
	if err := _Layerzero.contract.UnpackLog(event, "PacketBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroPacketDeliveredIterator is returned from FilterPacketDelivered and is used to iterate over the raw logs and unpacked data for PacketDelivered events raised by the Layerzero contract.
type LayerzeroPacketDeliveredIterator struct {
	Event *LayerzeroPacketDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroPacketDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroPacketDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroPacketDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroPacketDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroPacketDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroPacketDelivered represents a PacketDelivered event raised by the Layerzero contract.
type LayerzeroPacketDelivered struct {
	Origin   Origin
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPacketDelivered is a free log retrieval operation binding the contract event 0x3cd5e48f9730b129dc7550f0fcea9c767b7be37837cd10e55eb35f734f4bca04.
//
// Solidity: event PacketDelivered((uint32,bytes32,uint64) origin, address receiver)
func (_Layerzero *LayerzeroFilterer) FilterPacketDelivered(opts *bind.FilterOpts) (*LayerzeroPacketDeliveredIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "PacketDelivered")
	if err != nil {
		return nil, err
	}
	return &LayerzeroPacketDeliveredIterator{contract: _Layerzero.contract, event: "PacketDelivered", logs: logs, sub: sub}, nil
}

// WatchPacketDelivered is a free log subscription operation binding the contract event 0x3cd5e48f9730b129dc7550f0fcea9c767b7be37837cd10e55eb35f734f4bca04.
//
// Solidity: event PacketDelivered((uint32,bytes32,uint64) origin, address receiver)
func (_Layerzero *LayerzeroFilterer) WatchPacketDelivered(opts *bind.WatchOpts, sink chan<- *LayerzeroPacketDelivered) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "PacketDelivered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroPacketDelivered)
				if err := _Layerzero.contract.UnpackLog(event, "PacketDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketDelivered is a log parse operation binding the contract event 0x3cd5e48f9730b129dc7550f0fcea9c767b7be37837cd10e55eb35f734f4bca04.
//
// Solidity: event PacketDelivered((uint32,bytes32,uint64) origin, address receiver)
func (_Layerzero *LayerzeroFilterer) ParsePacketDelivered(log types.Log) (*LayerzeroPacketDelivered, error) {
	event := new(LayerzeroPacketDelivered)
	if err := _Layerzero.contract.UnpackLog(event, "PacketDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroPacketNilifiedIterator is returned from FilterPacketNilified and is used to iterate over the raw logs and unpacked data for PacketNilified events raised by the Layerzero contract.
type LayerzeroPacketNilifiedIterator struct {
	Event *LayerzeroPacketNilified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroPacketNilifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroPacketNilified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroPacketNilified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroPacketNilifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroPacketNilifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroPacketNilified represents a PacketNilified event raised by the Layerzero contract.
type LayerzeroPacketNilified struct {
	SrcEid      uint32
	Sender      [32]byte
	Receiver    common.Address
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPacketNilified is a free log retrieval operation binding the contract event 0xaf0450c392c4f702515a457a362328c8aa21916048ca6d0419e248b30cb55292.
//
// Solidity: event PacketNilified(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_Layerzero *LayerzeroFilterer) FilterPacketNilified(opts *bind.FilterOpts) (*LayerzeroPacketNilifiedIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "PacketNilified")
	if err != nil {
		return nil, err
	}
	return &LayerzeroPacketNilifiedIterator{contract: _Layerzero.contract, event: "PacketNilified", logs: logs, sub: sub}, nil
}

// WatchPacketNilified is a free log subscription operation binding the contract event 0xaf0450c392c4f702515a457a362328c8aa21916048ca6d0419e248b30cb55292.
//
// Solidity: event PacketNilified(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_Layerzero *LayerzeroFilterer) WatchPacketNilified(opts *bind.WatchOpts, sink chan<- *LayerzeroPacketNilified) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "PacketNilified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroPacketNilified)
				if err := _Layerzero.contract.UnpackLog(event, "PacketNilified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketNilified is a log parse operation binding the contract event 0xaf0450c392c4f702515a457a362328c8aa21916048ca6d0419e248b30cb55292.
//
// Solidity: event PacketNilified(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_Layerzero *LayerzeroFilterer) ParsePacketNilified(log types.Log) (*LayerzeroPacketNilified, error) {
	event := new(LayerzeroPacketNilified)
	if err := _Layerzero.contract.UnpackLog(event, "PacketNilified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroPacketSentIterator is returned from FilterPacketSent and is used to iterate over the raw logs and unpacked data for PacketSent events raised by the Layerzero contract.
type LayerzeroPacketSentIterator struct {
	Event *LayerzeroPacketSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroPacketSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroPacketSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroPacketSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroPacketSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroPacketSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroPacketSent represents a PacketSent event raised by the Layerzero contract.
type LayerzeroPacketSent struct {
	EncodedPayload []byte
	Options        []byte
	SendLibrary    common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPacketSent is a free log retrieval operation binding the contract event 0x1ab700d4ced0c005b164c0f789fd09fcbb0156d4c2041b8a3bfbcd961cd1567f.
//
// Solidity: event PacketSent(bytes encodedPayload, bytes options, address sendLibrary)
func (_Layerzero *LayerzeroFilterer) FilterPacketSent(opts *bind.FilterOpts) (*LayerzeroPacketSentIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return &LayerzeroPacketSentIterator{contract: _Layerzero.contract, event: "PacketSent", logs: logs, sub: sub}, nil
}

// WatchPacketSent is a free log subscription operation binding the contract event 0x1ab700d4ced0c005b164c0f789fd09fcbb0156d4c2041b8a3bfbcd961cd1567f.
//
// Solidity: event PacketSent(bytes encodedPayload, bytes options, address sendLibrary)
func (_Layerzero *LayerzeroFilterer) WatchPacketSent(opts *bind.WatchOpts, sink chan<- *LayerzeroPacketSent) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroPacketSent)
				if err := _Layerzero.contract.UnpackLog(event, "PacketSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketSent is a log parse operation binding the contract event 0x1ab700d4ced0c005b164c0f789fd09fcbb0156d4c2041b8a3bfbcd961cd1567f.
//
// Solidity: event PacketSent(bytes encodedPayload, bytes options, address sendLibrary)
func (_Layerzero *LayerzeroFilterer) ParsePacketSent(log types.Log) (*LayerzeroPacketSent, error) {
	event := new(LayerzeroPacketSent)
	if err := _Layerzero.contract.UnpackLog(event, "PacketSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroPacketVerifiedIterator is returned from FilterPacketVerified and is used to iterate over the raw logs and unpacked data for PacketVerified events raised by the Layerzero contract.
type LayerzeroPacketVerifiedIterator struct {
	Event *LayerzeroPacketVerified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroPacketVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroPacketVerified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroPacketVerified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroPacketVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroPacketVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroPacketVerified represents a PacketVerified event raised by the Layerzero contract.
type LayerzeroPacketVerified struct {
	Origin      Origin
	Receiver    common.Address
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPacketVerified is a free log retrieval operation binding the contract event 0x0d87345f3d1c929caba93e1c3821b54ff3512e12b66aa3cfe54b6bcbc17e59b4.
//
// Solidity: event PacketVerified((uint32,bytes32,uint64) origin, address receiver, bytes32 payloadHash)
func (_Layerzero *LayerzeroFilterer) FilterPacketVerified(opts *bind.FilterOpts) (*LayerzeroPacketVerifiedIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "PacketVerified")
	if err != nil {
		return nil, err
	}
	return &LayerzeroPacketVerifiedIterator{contract: _Layerzero.contract, event: "PacketVerified", logs: logs, sub: sub}, nil
}

// WatchPacketVerified is a free log subscription operation binding the contract event 0x0d87345f3d1c929caba93e1c3821b54ff3512e12b66aa3cfe54b6bcbc17e59b4.
//
// Solidity: event PacketVerified((uint32,bytes32,uint64) origin, address receiver, bytes32 payloadHash)
func (_Layerzero *LayerzeroFilterer) WatchPacketVerified(opts *bind.WatchOpts, sink chan<- *LayerzeroPacketVerified) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "PacketVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroPacketVerified)
				if err := _Layerzero.contract.UnpackLog(event, "PacketVerified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketVerified is a log parse operation binding the contract event 0x0d87345f3d1c929caba93e1c3821b54ff3512e12b66aa3cfe54b6bcbc17e59b4.
//
// Solidity: event PacketVerified((uint32,bytes32,uint64) origin, address receiver, bytes32 payloadHash)
func (_Layerzero *LayerzeroFilterer) ParsePacketVerified(log types.Log) (*LayerzeroPacketVerified, error) {
	event := new(LayerzeroPacketVerified)
	if err := _Layerzero.contract.UnpackLog(event, "PacketVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroReceiveLibrarySetIterator is returned from FilterReceiveLibrarySet and is used to iterate over the raw logs and unpacked data for ReceiveLibrarySet events raised by the Layerzero contract.
type LayerzeroReceiveLibrarySetIterator struct {
	Event *LayerzeroReceiveLibrarySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroReceiveLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroReceiveLibrarySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroReceiveLibrarySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroReceiveLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroReceiveLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroReceiveLibrarySet represents a ReceiveLibrarySet event raised by the Layerzero contract.
type LayerzeroReceiveLibrarySet struct {
	Receiver common.Address
	Eid      uint32
	NewLib   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceiveLibrarySet is a free log retrieval operation binding the contract event 0xcd6f92f5ac6185a5acfa02c92090746cec64d777269cbcd0ed031e396657a1c2.
//
// Solidity: event ReceiveLibrarySet(address receiver, uint32 eid, address newLib)
func (_Layerzero *LayerzeroFilterer) FilterReceiveLibrarySet(opts *bind.FilterOpts) (*LayerzeroReceiveLibrarySetIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "ReceiveLibrarySet")
	if err != nil {
		return nil, err
	}
	return &LayerzeroReceiveLibrarySetIterator{contract: _Layerzero.contract, event: "ReceiveLibrarySet", logs: logs, sub: sub}, nil
}

// WatchReceiveLibrarySet is a free log subscription operation binding the contract event 0xcd6f92f5ac6185a5acfa02c92090746cec64d777269cbcd0ed031e396657a1c2.
//
// Solidity: event ReceiveLibrarySet(address receiver, uint32 eid, address newLib)
func (_Layerzero *LayerzeroFilterer) WatchReceiveLibrarySet(opts *bind.WatchOpts, sink chan<- *LayerzeroReceiveLibrarySet) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "ReceiveLibrarySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroReceiveLibrarySet)
				if err := _Layerzero.contract.UnpackLog(event, "ReceiveLibrarySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceiveLibrarySet is a log parse operation binding the contract event 0xcd6f92f5ac6185a5acfa02c92090746cec64d777269cbcd0ed031e396657a1c2.
//
// Solidity: event ReceiveLibrarySet(address receiver, uint32 eid, address newLib)
func (_Layerzero *LayerzeroFilterer) ParseReceiveLibrarySet(log types.Log) (*LayerzeroReceiveLibrarySet, error) {
	event := new(LayerzeroReceiveLibrarySet)
	if err := _Layerzero.contract.UnpackLog(event, "ReceiveLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroReceiveLibraryTimeoutSetIterator is returned from FilterReceiveLibraryTimeoutSet and is used to iterate over the raw logs and unpacked data for ReceiveLibraryTimeoutSet events raised by the Layerzero contract.
type LayerzeroReceiveLibraryTimeoutSetIterator struct {
	Event *LayerzeroReceiveLibraryTimeoutSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroReceiveLibraryTimeoutSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroReceiveLibraryTimeoutSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroReceiveLibraryTimeoutSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroReceiveLibraryTimeoutSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroReceiveLibraryTimeoutSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroReceiveLibraryTimeoutSet represents a ReceiveLibraryTimeoutSet event raised by the Layerzero contract.
type LayerzeroReceiveLibraryTimeoutSet struct {
	Receiver common.Address
	Eid      uint32
	OldLib   common.Address
	Timeout  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceiveLibraryTimeoutSet is a free log retrieval operation binding the contract event 0x4e0a5bbfa0c11a64effb1ada324b5437a17272e1aed9320398715ef71bb20928.
//
// Solidity: event ReceiveLibraryTimeoutSet(address receiver, uint32 eid, address oldLib, uint256 timeout)
func (_Layerzero *LayerzeroFilterer) FilterReceiveLibraryTimeoutSet(opts *bind.FilterOpts) (*LayerzeroReceiveLibraryTimeoutSetIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "ReceiveLibraryTimeoutSet")
	if err != nil {
		return nil, err
	}
	return &LayerzeroReceiveLibraryTimeoutSetIterator{contract: _Layerzero.contract, event: "ReceiveLibraryTimeoutSet", logs: logs, sub: sub}, nil
}

// WatchReceiveLibraryTimeoutSet is a free log subscription operation binding the contract event 0x4e0a5bbfa0c11a64effb1ada324b5437a17272e1aed9320398715ef71bb20928.
//
// Solidity: event ReceiveLibraryTimeoutSet(address receiver, uint32 eid, address oldLib, uint256 timeout)
func (_Layerzero *LayerzeroFilterer) WatchReceiveLibraryTimeoutSet(opts *bind.WatchOpts, sink chan<- *LayerzeroReceiveLibraryTimeoutSet) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "ReceiveLibraryTimeoutSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroReceiveLibraryTimeoutSet)
				if err := _Layerzero.contract.UnpackLog(event, "ReceiveLibraryTimeoutSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceiveLibraryTimeoutSet is a log parse operation binding the contract event 0x4e0a5bbfa0c11a64effb1ada324b5437a17272e1aed9320398715ef71bb20928.
//
// Solidity: event ReceiveLibraryTimeoutSet(address receiver, uint32 eid, address oldLib, uint256 timeout)
func (_Layerzero *LayerzeroFilterer) ParseReceiveLibraryTimeoutSet(log types.Log) (*LayerzeroReceiveLibraryTimeoutSet, error) {
	event := new(LayerzeroReceiveLibraryTimeoutSet)
	if err := _Layerzero.contract.UnpackLog(event, "ReceiveLibraryTimeoutSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroSendLibrarySetIterator is returned from FilterSendLibrarySet and is used to iterate over the raw logs and unpacked data for SendLibrarySet events raised by the Layerzero contract.
type LayerzeroSendLibrarySetIterator struct {
	Event *LayerzeroSendLibrarySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LayerzeroSendLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroSendLibrarySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LayerzeroSendLibrarySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LayerzeroSendLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroSendLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroSendLibrarySet represents a SendLibrarySet event raised by the Layerzero contract.
type LayerzeroSendLibrarySet struct {
	Sender common.Address
	Eid    uint32
	NewLib common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSendLibrarySet is a free log retrieval operation binding the contract event 0x4cff966ebee29a156dcb34cf72c1d06231fb1777f6bdf6e8089819232f002b1c.
//
// Solidity: event SendLibrarySet(address sender, uint32 eid, address newLib)
func (_Layerzero *LayerzeroFilterer) FilterSendLibrarySet(opts *bind.FilterOpts) (*LayerzeroSendLibrarySetIterator, error) {

	logs, sub, err := _Layerzero.contract.FilterLogs(opts, "SendLibrarySet")
	if err != nil {
		return nil, err
	}
	return &LayerzeroSendLibrarySetIterator{contract: _Layerzero.contract, event: "SendLibrarySet", logs: logs, sub: sub}, nil
}

// WatchSendLibrarySet is a free log subscription operation binding the contract event 0x4cff966ebee29a156dcb34cf72c1d06231fb1777f6bdf6e8089819232f002b1c.
//
// Solidity: event SendLibrarySet(address sender, uint32 eid, address newLib)
func (_Layerzero *LayerzeroFilterer) WatchSendLibrarySet(opts *bind.WatchOpts, sink chan<- *LayerzeroSendLibrarySet) (event.Subscription, error) {

	logs, sub, err := _Layerzero.contract.WatchLogs(opts, "SendLibrarySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroSendLibrarySet)
				if err := _Layerzero.contract.UnpackLog(event, "SendLibrarySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSendLibrarySet is a log parse operation binding the contract event 0x4cff966ebee29a156dcb34cf72c1d06231fb1777f6bdf6e8089819232f002b1c.
//
// Solidity: event SendLibrarySet(address sender, uint32 eid, address newLib)
func (_Layerzero *LayerzeroFilterer) ParseSendLibrarySet(log types.Log) (*LayerzeroSendLibrarySet, error) {
	event := new(LayerzeroSendLibrarySet)
	if err := _Layerzero.contract.UnpackLog(event, "SendLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
