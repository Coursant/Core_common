package protocol

import (
	"encoding/json"

	"github.com/Coursant/Core_common"
)

const (
	// ContractInvokeRequestMessageType defines contract invoke request type of the communication protocol
	ContractInvokeRequestMessageType Core_common.ProtocolMessage = Core_common.Iden3Protocol + "proofs/1.0/contract-invoke-request"
)

// ContractInvokeRequestMessage is struct the represents iden3message contract invoke request
type ContractInvokeRequestMessage struct {
	ID       string                           `json:"id"`
	Typ      Core_common.MediaType            `json:"typ,omitempty"`
	Type     Core_common.ProtocolMessage      `json:"type"`
	ThreadID string                           `json:"thid,omitempty"`
	Body     ContractInvokeRequestMessageBody `json:"body,omitempty"`

	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// ContractInvokeRequestMessageBody is body for contract invoke request
type ContractInvokeRequestMessageBody struct {
	Reason          string                      `json:"reason,omitempty"`
	Message         string                      `json:"message,omitempty"`
	TransactionData TransactionData             `json:"transaction_data"`
	DIDDoc          json.RawMessage             `json:"did_doc,omitempty"`
	Scope           []ZeroKnowledgeProofRequest `json:"scope"`
}

// TransactionData represents structure for on chain verification
type TransactionData struct {
	ContractAddress string `json:"contract_address"`
	MethodID        string `json:"method_id"`
	ChainID         int    `json:"chain_id"`
	Network         string `json:"network"`
}
