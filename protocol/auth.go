// Package protocol defines core protocol messages
package protocol

import (
	"encoding/json"

	"github.com/Coursant/Core_common"
	"github.com/Coursant/Core_rapidsnark/types"
)

const (

	// AuthorizationRequestMessageType defines auth request type of the communication protocol
	AuthorizationRequestMessageType Core_common.ProtocolMessage = Core_common.Iden3Protocol + "authorization/1.0/request"
	// AuthorizationResponseMessageType defines auth response type of the communication protocol
	AuthorizationResponseMessageType Core_common.ProtocolMessage = Core_common.Iden3Protocol + "authorization/1.0/response"
)

// AuthorizationResponseMessage is struct the represents iden3message authorization response
type AuthorizationResponseMessage struct {
	ID       string                           `json:"id"`
	Typ      Core_common.MediaType            `json:"typ,omitempty"`
	Type     Core_common.ProtocolMessage      `json:"type"`
	ThreadID string                           `json:"thid,omitempty"`
	Body     AuthorizationMessageResponseBody `json:"body,omitempty"`

	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// AuthorizationMessageResponseBody is struct the represents authorization response data
type AuthorizationMessageResponseBody struct {
	DIDDoc  json.RawMessage              `json:"did_doc,omitempty"`
	Message string                       `json:"message,omitempty"`
	Scope   []ZeroKnowledgeProofResponse `json:"scope"`
}

// AuthorizationRequestMessage is struct the represents iden3message authorization request
type AuthorizationRequestMessage struct {
	ID       string                          `json:"id"`
	Typ      Core_common.MediaType           `json:"typ,omitempty"`
	Type     Core_common.ProtocolMessage     `json:"type"`
	ThreadID string                          `json:"thid,omitempty"`
	Body     AuthorizationRequestMessageBody `json:"body,omitempty"`

	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// AuthorizationRequestMessageBody is body for authorization request
type AuthorizationRequestMessageBody struct {
	CallbackURL string                      `json:"callbackUrl"`
	Reason      string                      `json:"reason,omitempty"`
	Message     string                      `json:"message,omitempty"`
	DIDDoc      json.RawMessage             `json:"did_doc,omitempty"`
	Scope       []ZeroKnowledgeProofRequest `json:"scope"`
}

// ZeroKnowledgeProofRequest represents structure of zkp request object
type ZeroKnowledgeProofRequest struct {
	ID        uint32                 `json:"id"` // unique request id
	CircuitID string                 `json:"circuitId"`
	Optional  *bool                  `json:"optional,omitempty"`
	Query     map[string]interface{} `json:"query"`
}

// ZeroKnowledgeProofResponse represents structure of zkp response
type ZeroKnowledgeProofResponse struct {
	ID                     uint32          `json:"id"` // unique id to present unique proof request
	CircuitID              string          `json:"circuitId"`
	VerifiablePresentation json.RawMessage `json:"vp,omitempty"`
	types.ZKProof
}
