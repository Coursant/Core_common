package packers

import (
	"encoding/json"
	"github.com/Coursant/Core_common"
	"github.com/pkg/errors"
)

// MediaTypePlainMessage is media type for plain message
const MediaTypePlainMessage Core_common.MediaType = "application/iden3comm-plain-json"

// PlainMessagePacker is simple packer that doesn't use encryption / encoding
type PlainMessagePacker struct {
}

// PlainPackerParams is params for plain packer
type PlainPackerParams struct {
	Core_common.PackerParams
}

// Pack returns packed message to transport envelope
func (p *PlainMessagePacker) Pack(payload []byte, params Core_common.PackerParams) ([]byte, error) {

	var msgMap map[string]interface{}
	err := json.Unmarshal(payload, &msgMap)
	if err != nil {
		return nil, err
	}
	msgMap["typ"] = MediaTypePlainMessage
	return json.Marshal(msgMap)
}

// Unpack returns unpacked message from transport envelope
func (p *PlainMessagePacker) Unpack(envelope []byte) (*Core_common.BasicMessage, error) {

	var msg Core_common.BasicMessage
	err := json.Unmarshal(envelope, &msg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &msg, err
}

// MediaType for iden3comm
func (p *PlainMessagePacker) MediaType() Core_common.MediaType {
	return MediaTypePlainMessage
}
