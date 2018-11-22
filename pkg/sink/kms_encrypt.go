package sink

import (
	"encoding/base64"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/piotrjaromin/map-processor/pkg/awsutil"
)

type KmsEncrypt struct {
	Region         *string
	EncryptContext *string
	KeyId          *string
	params         map[string]string
}

func NewKmsEncrypt(region string, EncryptContext string, KeyId string) (*KmsEncrypt, error) {

	return &KmsEncrypt{
		Region:         &region,
		EncryptContext: &EncryptContext,
		KeyId:          &KeyId,
	}, nil

}

func (ts *KmsEncrypt) Fill(params map[string]string) error {
	result := map[string]string{}

	sess, err := awsutil.GetSession(ts.Region)
	if err != nil {
		return err
	}

	svc := kms.New(sess)
	for key, val := range params {

		input := &kms.EncryptInput{
			Plaintext:         []byte(val),
			EncryptionContext: map[string]*string{"service": ts.EncryptContext},
			KeyId:             ts.KeyId,
		}

		EncryptedVal, err := svc.Encrypt(input)
		if err != nil {
			result[key] = err.Error()
		} else {

			blob := base64.StdEncoding.EncodeToString(EncryptedVal.CiphertextBlob)
			result[key] = string(blob)
		}
	}

	ts.params = result
	return nil
}

func (ts *KmsEncrypt) Get() (map[string]string, error) {
	return ts.params, nil
}
