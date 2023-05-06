package protocol

import (
	"encoding/json"
	"github.com/Coursant/Core_common"

	"github.com/Coursant/Core_common/schema-processor/verifiable"
)

const (

	// CredentialIssuanceRequestMessageType accepts request for credential creation
	CredentialIssuanceRequestMessageType Core_common.ProtocolMessage = Core_common.Iden3Protocol + "credentials/1.0/issuance-request"

	// CredentialFetchRequestMessageType is type for request of credential generation
	CredentialFetchRequestMessageType Core_common.ProtocolMessage = Core_common.Iden3Protocol + "credentials/1.0/fetch-request"

	// CredentialOfferMessageType is type of message with credential offering
	CredentialOfferMessageType Core_common.ProtocolMessage = Core_common.Iden3Protocol + "credentials/1.0/offer"

	// CredentialIssuanceResponseMessageType is type for message with a credential issuance
	CredentialIssuanceResponseMessageType Core_common.ProtocolMessage = Core_common.Iden3Protocol + "credentials/1.0/issuance-response"
)

// CredentialIssuanceRequestMessage represent Iden3message for credential request
type CredentialIssuanceRequestMessage struct {
	ID       string                      `json:"id"`
	Typ      Core_common.MediaType       `json:"typ,omitempty"`
	Type     Core_common.ProtocolMessage `json:"type"`
	ThreadID string                      `json:"thid,omitempty"`

	Body CredentialIssuanceRequestMessageBody `json:"body,omitempty"`

	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// CredentialIssuanceRequestMessageBody represents data for credential issuance request
type CredentialIssuanceRequestMessageBody struct {
	Schema     Schema          `json:"schema"`
	Data       json.RawMessage `json:"data"`
	Expiration int64           `json:"expiration"`
}

// CredentialsOfferMessage represent Iden3message for credential offer
type CredentialsOfferMessage struct {
	ID       string                      `json:"id"`
	Typ      Core_common.MediaType       `json:"typ,omitempty"`
	Type     Core_common.ProtocolMessage `json:"type"`
	ThreadID string                      `json:"thid,omitempty"`

	Body CredentialsOfferMessageBody `json:"body,omitempty"`

	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// CredentialsOfferMessageBody is struct the represents offer message
type CredentialsOfferMessageBody struct {
	URL         string            `json:"url"`
	Credentials []CredentialOffer `json:"credentials"`
}

// CredentialOffer is structure to fetch credential
type CredentialOffer struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

// CredentialIssuanceMessage represent Iden3message for credential issuance
type CredentialIssuanceMessage struct {
	ID       string                      `json:"id"`
	Typ      Core_common.MediaType       `json:"typ,omitempty"`
	Type     Core_common.ProtocolMessage `json:"type"`
	ThreadID string                      `json:"thid,omitempty"`

	Body IssuanceMessageBody `json:"body,omitempty"`

	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// IssuanceMessageBody is struct the represents message when credential is issued
type IssuanceMessageBody struct {
	Credential verifiable.W3CCredential `json:"credential"`
}

// CredentialFetchRequestMessage represent Iden3message for credential fetch request
type CredentialFetchRequestMessage struct {
	ID       string                            `json:"id"`
	Typ      Core_common.MediaType             `json:"typ,omitempty"`
	Type     Core_common.ProtocolMessage       `json:"type"`
	ThreadID string                            `json:"thid,omitempty"`
	Body     CredentialFetchRequestMessageBody `json:"body,omitempty"`

	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// CredentialFetchRequestMessageBody is msg body for fetch request
type CredentialFetchRequestMessageBody struct {
	ID string `json:"id"`
}

// Schema represents location and type where it's stored
type Schema struct {
	Hash string `json:"hash,omitempty"`
	URL  string `json:"url"`
	Type string `json:"type"`
}
