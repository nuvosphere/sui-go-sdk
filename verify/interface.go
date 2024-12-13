package verify

import "github.com/machinebox/graphql"

type IPublicKey interface {
	ToSuiAddress() string
	VerifyPersonalMessage(message []byte, signature string, client *graphql.Client) (bool, error)
}
