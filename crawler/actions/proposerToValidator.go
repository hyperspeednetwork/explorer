package actions

import (
	"github.com/wongyinlong/hsnNet/models"
	"github.com/wongyinlong/hsnNet/models/utils"
	"go.uber.org/zap"
)

func SetValidatorHashAddress(operaAddress string, pubkey string,log zap.Logger) {
	/*
		Get the validators public key .
		Creat hash mapping between public key and validator and hash values.
	*/
	var validatorAddressAndKey models.ValidatorAddressAndKey
	Sign,_:= validatorAddressAndKey.CheckValidator(pubkey)
	if Sign == 0 {
		validatorAddressAndKey.ConsensusPubkey = pubkey
		validatorAddressAndKey.OperatorAddress = operaAddress
		validatorAddressAndKey.ProposerHash = utils.GenHexAddrFromPubKey(pubkey)
		validatorAddressAndKey.SetInfo(log)
	}
}
