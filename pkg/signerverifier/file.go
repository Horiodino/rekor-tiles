/*
Copyright 2025 The Sigstore Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Copied from https://github.com/sigstore/rekor/blob/c820fcaf3afdc91f0acf6824d55c1ac7df249df1/pkg/signer/file.go

package signerverifier

import (
	"fmt"

	"github.com/sigstore/sigstore/pkg/signature"
	"go.step.sm/crypto/pemutil"
)

// File is a file-based signer/verifier.
type File struct {
	signature.SignerVerifier
}

// NewFileSignerVerifier returns an file-based signer-verifier, used for spinning up local instances.
func NewFileSignerVerifier(keyPath, keyPass string) (*File, error) {
	opaqueKey, err := pemutil.Read(keyPath, pemutil.WithPassword([]byte(keyPass)))
	if err != nil {
		return nil, fmt.Errorf("file: provide a valid signer, %s is not valid: %w", keyPath, err)
	}

	signerVerifier, err := signature.LoadDefaultSignerVerifier(opaqueKey)
	if err != nil {
		return nil, fmt.Errorf(`file: loaded private key from %s can't be used to sign: %w`, keyPath, err)
	}
	return &File{signerVerifier}, nil
}
