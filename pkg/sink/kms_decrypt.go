package sink

import (
	"encoding/base64"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/piotrjaromin/map-processor/pkg/awsutil"
)

type KmsDecrypt struct {
	Region         *string
	DecryptContext *string
	params         map[string]string
}

func NewKmsDecrypt(region string, decryptContext string) (*KmsDecrypt, error) {

	return &KmsDecrypt{
		Region:         &region,
		DecryptContext: &decryptContext,
	}, nil

}

func (ts *KmsDecrypt) Fill(params map[string]string) error {
	result := map[string]string{}

	sess, err := awsutil.GetSession(ts.Region)
	if err != nil {
		return err
	}

	svc := kms.New(sess)
	for key, val := range params {

		blob, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			result[key] = err.Error()
		}

		input := &kms.DecryptInput{
			CiphertextBlob:    []byte(blob),
			EncryptionContext: map[string]*string{"service": ts.DecryptContext},
		}

		decryptedVal, err := svc.Decrypt(input)
		if err != nil {
			result[key] = err.Error()
		} else {
			result[key] = string(decryptedVal.Plaintext)
		}
	}

	ts.params = result
	return nil
}

func (ts *KmsDecrypt) Get() (map[string]string, error) {
	return ts.params, nil
}
