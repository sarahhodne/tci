package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/henrikhodne/tci/travis"
)

func init() {
	cmds["encrypt"] = cmd{encrypt, "str", "encrypt variables for .travis.yml"}
	cmdHelp["encrypt"] = `Encrypts environment variables.

$ tci encrypt FOO=bar
Please add the following to your .travis.yml file:

  secure: "gSly+Kvzd5uSul15CVaEV91ALwsGSU7yJLHSK0vk+oqjmLm0jp05iiKfs08j/Wo0DG8l4O9WT0mCEnMoMBwX4GiK4mUmGdKt0R2/2IAea+M44kBoKsiRM7R3+62xEl0q9Wzt8Aw3GCDY4XnoCyirO49DpCH6a9JEAfILY/n6qF8="
`
}

func encrypt(str string) {
	client := travis.NewClient()
	key, err := client.GetRepositoryKey(detectSlug())
	if err != nil {
		fmt.Printf("Error retrieving repository key: %+v", err)
		return
	}

	pemData := []byte(key.Key)
	block, _ := pem.Decode(pemData)
	maybePub, err := x509.ParsePKIXPublicKey(block.Bytes)
	pub := maybePub.(*rsa.PublicKey)
	out, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(str))
	outStr := base64.StdEncoding.EncodeToString(out)
	fmt.Printf("secure: %q\n", outStr)
}
