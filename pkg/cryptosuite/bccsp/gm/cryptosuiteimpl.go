/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package gm

import (
	"github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	"github.com/hyperledger/fabric-sdk-go/api/apicryptosuite"
	"github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/bccsp"
	bccspSw "github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/bccsp/factory/gm"
	"github.com/hyperledger/fabric-sdk-go/pkg/cryptosuite/bccsp/wrapper"
	"github.com/hyperledger/fabric-sdk-go/pkg/logging"
	"github.com/pkg/errors"
)

var logger = logging.NewLogger("fabric_sdk_go")

//GetSuiteByConfig returns cryptosuite adaptor for bccsp loaded according to given config
func GetSuiteByConfig(config apiconfig.Config) (apicryptosuite.CryptoSuite, error) {
	// TODO: delete this check?
	if config.SecurityProvider() != "GM" {
		return nil, errors.Errorf("Unsupported BCCSP Provider: %s", config.SecurityProvider())
	}

	opts := getOptsByConfig(config)
	bccsp, err := getBCCSPFromOpts(opts)
	if err != nil {
		return nil, err
	}
	return wrapper.NewCryptoSuite(bccsp), nil
}

//GetSuiteWithDefaultEphemeral returns cryptosuite adaptor for bccsp with default ephemeral options (intended to aid testing)
func GetSuiteWithDefaultEphemeral() (apicryptosuite.CryptoSuite, error) {
	opts := getEphemeralOpts()

	bccsp, err := getBCCSPFromOpts(opts)
	if err != nil {
		return nil, err
	}
	return wrapper.NewCryptoSuite(bccsp), nil
}

func getBCCSPFromOpts(config *bccspSw.GmOpts) (bccsp.BCCSP, error) {
	f := &bccspSw.GMFactory{}

	csp, err := f.Get(config)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not initialize BCCSP %s", f.Name())
	}
	return csp, nil
}

//GetOptsByConfig Returns Factory opts for given SDK config
func getOptsByConfig(c apiconfig.Config) *bccspSw.GmOpts {
	opts := &bccspSw.GmOpts{
		HashFamily: c.SecurityAlgorithm(),
		SecLevel:   c.SecurityLevel(),
		FileKeystore: &bccspSw.FileKeystoreOpts{
			KeyStorePath: c.KeyStorePath(),
		},
		Ephemeral: c.Ephemeral(),
	}
	logger.Debug("Initialized SW cryptosuite11111, %v",c.SecurityAlgorithm())

	return opts
}

func getEphemeralOpts() *bccspSw.GmOpts {
	opts := &bccspSw.GmOpts{
		HashFamily: "SHA256",
		SecLevel:   256,
		Ephemeral:  true,
	}
	logger.Debug("Initialized ephemeral SW cryptosuite with default opts")

	return opts
}