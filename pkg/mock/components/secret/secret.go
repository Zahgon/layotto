package secret

import (
	"github.com/dapr/components-contrib/secretstores"
)

type FakeSecretStore struct{}

func (c FakeSecretStore) GetSecret(req secretstores.GetSecretRequest) (secretstores.GetSecretResponse, error) {
	_ = "STUB: not implemented"
	return *new(secretstores.GetSecretResponse), nil
}

func (c FakeSecretStore) BulkGetSecret(req secretstores.BulkGetSecretRequest) (secretstores.BulkGetSecretResponse, error) {
	_ = "STUB: not implemented"
	return *new(secretstores.BulkGetSecretResponse), nil
}

func (c FakeSecretStore) Init(metadata secretstores.Metadata) error {
	_ = "STUB: not implemented"
	return nil
}

func (c FakeSecretStore) Close() error { _ = "STUB: not implemented"; return nil }
