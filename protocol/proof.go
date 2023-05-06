package protocol

import (
	"github.com/Coursant/Core_common"
)

const (
	// ProofGenerationRequestMessageType is type for request of proof generation
	ProofGenerationRequestMessageType Core_common.ProtocolMessage = Core_common.Iden3Protocol + "proofs/1.0/request"
	// ProofGenerationResponseMessageType is type for response of proof generation
	ProofGenerationResponseMessageType Core_common.ProtocolMessage = Core_common.Iden3Protocol + "proofs/1.0/response"
)

// ProofGenerationRequestMessage is struct the represents body for proof generation request
type ProofGenerationRequestMessage struct {
	ID       string                      `json:"id"`
	Typ      Core_common.MediaType       `json:"typ,omitempty"`
	Type     Core_common.ProtocolMessage `json:"type"`
	ThreadID string                      `json:"thid,omitempty"`

	Body ProofGenerationRequestMessageBody `json:"body,omitempty"`

	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// ProofGenerationRequestMessageBody is struct the represents body for proof generation request
type ProofGenerationRequestMessageBody struct {
	Scope []ZeroKnowledgeProofRequest `json:"scope"`
}

// ProofGenerationResponseMessage is struct the represents body for proof generation request
type ProofGenerationResponseMessage struct {
	ID       string                      `json:"id"`
	Typ      Core_common.MediaType       `json:"typ,omitempty"`
	Type     Core_common.ProtocolMessage `json:"type"`
	ThreadID string                      `json:"thid,omitempty"`

	Body ResponseMessageBody `json:"body,omitempty"`

	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// ResponseMessageBody is struct the represents request for revocation status
type ResponseMessageBody struct {
	Scope []ZeroKnowledgeProofResponse `json:"scope"`
}
