package asymmetric

import (
	"github.com/suyuan32/knife/cryptox/asymmetric/rsa"
)

// NewRSA returns an RSA struct.
func NewRSA() *rsa.RSA {
	return &rsa.RSA{
		Standard: rsa.PKCS1,
	}
}
