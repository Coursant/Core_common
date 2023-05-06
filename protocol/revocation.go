package protocol

import (
	"github.com/Coursant/Core_common"
	"github.com/Coursant/Core_common/schema-processor/verifiable"
)

const (
	// RevocationStatusRequestMessageType is type for request of revocation status
	RevocationStatusRequestMessageType Core_common.ProtocolMessage = Core_common.Iden3Protocol + "revocation/1.0/request-status"
	// RevocationStatusResponseMessageType is type for response with a revocation status
	RevocationStatusResponseMessageType Core_common.ProtocolMessage = Core_common.Iden3Protocol + "revocation/1.0/status"
)

// RevocationStatusRequestMessage is struct the represents body for proof generation request
type RevocationStatusRequestMessage struct {
	ID       string                             `json:"id"`
	Typ      Core_common.MediaType              `json:"typ,omitempty"`
	Type     Core_common.ProtocolMessage        `json:"type"`
	ThreadID string                             `json:"thid,omitempty"`
	Body     RevocationStatusRequestMessageBody `json:"body,omitempty"`

	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// RevocationStatusRequestMessageBody is struct the represents request for revocation status
type RevocationStatusRequestMessageBody struct {
	RevocationNonce uint64 `json:"revocation_nonce"`
}

// RevocationStatusResponseMessage is struct the represents body for proof generation request
type RevocationStatusResponseMessage struct {
	ID       string                      `json:"id"`
	Typ      Core_common.MediaType       `json:"typ,omitempty"`
	Type     Core_common.ProtocolMessage `json:"type"`
	ThreadID string                      `json:"thid,omitempty"`

	Body RevocationStatusResponseMessageBody `json:"body,omitempty"`

	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// RevocationStatusResponseMessageBody is struct the represents request for revocation status
type RevocationStatusResponseMessageBody struct {
	verifiable.RevocationStatus
}
