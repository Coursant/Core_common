package protocol

import "github.com/Coursant/Core_common"

const (
	// MessageFetchRequestMessageType defines message fetch request type of the communication protocol.
	MessageFetchRequestMessageType Core_common.ProtocolMessage = Core_common.Iden3Protocol + "messages/1.0/fetch"
)

// MessageFetchRequestMessage represent Iden3message for message fetch request.
type MessageFetchRequestMessage struct {
	ID       string                      `json:"id"`
	Typ      Core_common.MediaType       `json:"typ,omitempty"`
	Type     Core_common.ProtocolMessage `json:"type"`
	ThreadID string                      `json:"thid,omitempty"`

	Body MessageFetchRequestMessageBody `json:"body,omitempty"`

	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// MessageFetchRequestMessageBody is struct the represents body for message fetch request.
type MessageFetchRequestMessageBody struct {
	ID string `json:"id"`
}
