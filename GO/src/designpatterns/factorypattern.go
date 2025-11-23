package designpatterns

import "fmt"

type (
	CryptoGraphic interface {
		Encrypt(key string, payload string) string
		Decrypt(key string, payload string) string
	}

	RSA struct {
	}

	EC struct {
	}
)

func (r *RSA) Encrypt(key string, payload string) string {

	return "RSA encryption completed"

}

func (r *RSA) Decrypt(key string, payload string) string {

	return "RSA Decryption completed"

}

func (r *EC) Encrypt(key string, payload string) string {

	return "RSA encryption completed"

}

func (r *EC) Decrypt(key string, payload string) string {

	return "RSA Decryption completed"

}

func CryptoFactory(algo string) CryptoGraphic {
	switch algo {
	case "RSA":
		return &RSA{}
	case "EC":
		return &EC{}

	default:
		fmt.Println("NO suitable algo found")
		return nil
	}
}
