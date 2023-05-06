// Package packers defines core 3 protocol packers: anoncrypt, plain and zkp
package packers

import (
	"encoding/json"

	"github.com/Coursant/Core_common"
	"github.com/pkg/errors"

	jose "gopkg.in/square/go-jose.v2"
)

// MediaTypeEncryptedMessage is media type for ecnrypted message
const MediaTypeEncryptedMessage Core_common.MediaType = "application/Core_common-encrypted-json"

// AnoncryptPacker is  packer for anon encryption / decryption
type AnoncryptPacker struct {
	kr KeyResolverHandlerFunc
}

// AnoncryptPackerParams is params for anoncrypt packer
type AnoncryptPackerParams struct {
	RecipientKey *jose.JSONWebKey
	Core_common.PackerParams
}

// NewAnoncryptPacker returns new anon packers
func NewAnoncryptPacker(kr KeyResolverHandlerFunc) *AnoncryptPacker {
	return &AnoncryptPacker{kr: kr}
}

// KeyResolverHandlerFunc resolve private key by key id
type KeyResolverHandlerFunc func(keyID string) (key interface{}, err error)

// Resolve function is responsible to call provided handler for recipient private key resolve
func (kr KeyResolverHandlerFunc) Resolve(keyID string) (key interface{}, err error) {
	return kr(keyID)
}

// Pack returns packed message to transport envelope
func (p *AnoncryptPacker) Pack(payload []byte, params Core_common.PackerParams) ([]byte, error) {

	packerParams, ok := params.(AnoncryptPackerParams)
	if !ok {
		return nil, errors.New("can't cast params to anoncrypt packer params")
	}

	encryptor, err := jose.NewEncrypter(jose.A256CBC_HS512, jose.Recipient{
		Algorithm: jose.ECDH_ES_A256KW,
		Key:       packerParams.RecipientKey,
		KeyID:     packerParams.RecipientKey.KeyID,
	}, new(jose.EncrypterOptions).WithType(jose.ContentType(MediaTypeEncryptedMessage)))
	if err != nil {
		return nil, err
	}
	jwe, err := encryptor.Encrypt(payload)
	if err != nil {
		return nil, err
	}
	jweString, err := jwe.CompactSerialize()
	if err != nil {
		return nil, err
	}
	return []byte(jweString), nil

}

// Unpack returns unpacked message from transport envelope
func (p *AnoncryptPacker) Unpack(envelope []byte) (*Core_common.BasicMessage, error) {

	jwe, err := jose.ParseEncrypted(string(envelope))
	if err != nil {
		return nil, err
	}
	decryptionKey, err := p.kr.Resolve(jwe.Header.KeyID)
	if err != nil {
		return nil, err
	}
	payload, err := jwe.Decrypt(decryptionKey)
	if err != nil {
		return nil, err
	}
	var msg Core_common.BasicMessage
	err = json.Unmarshal(payload, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, err
}

// MediaType for Core_common
func (p *AnoncryptPacker) MediaType() Core_common.MediaType {
	return MediaTypeEncryptedMessage
}
