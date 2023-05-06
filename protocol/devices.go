package protocol

import "github.com/Coursant/Core_common"

const (
	// DeviceRegistrationRequestMessageType defines device registration request type of the communication protocol
	DeviceRegistrationRequestMessageType Core_common.ProtocolMessage = Core_common.Iden3Protocol + "devices/1.0/registration"
)

// DeviceRegistrationRequestMessage represent Iden3message for register device request
type DeviceRegistrationRequestMessage struct {
	ID       string                      `json:"id"`
	Typ      Core_common.MediaType       `json:"typ,omitempty"`
	Type     Core_common.ProtocolMessage `json:"type"`
	ThreadID string                      `json:"thid,omitempty"`

	Body DeviceRegistrationRequestMessageBody `json:"body,omitempty"`

	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// DeviceRegistrationRequestMessageBody is struct the represents body for register device request request
type DeviceRegistrationRequestMessageBody struct {
	AppID     string `json:"app_id"`
	PushToken string `json:"push_token"`
}
