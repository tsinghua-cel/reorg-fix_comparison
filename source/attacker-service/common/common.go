package common

import (
	"encoding/base64"
	"errors"
	ethpb "github.com/prysmaticlabs/prysm/v5/proto/prysm/v1alpha1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

var (
	ErrNilObject              = errors.New("nil object")
	ErrUnsupportedBeaconBlock = errors.New("unsupported beacon block")
)

func Base64ToAttestationData(attestDataBase64 string) (*ethpb.AttestationData, error) {
	attestData, err := base64.StdEncoding.DecodeString(attestDataBase64)
	if err != nil {
		log.WithError(err).Error("base64 decode attest data failed")
		return nil, err
	}
	var attestation = new(ethpb.AttestationData)
	if err := proto.Unmarshal(attestData, attestation); err != nil {
		log.WithError(err).Error("unmarshal attest data failed")
		return nil, err
	}
	return attestation, nil
}

func AttestationDataToBase64(attestation *ethpb.AttestationData) (string, error) {
	data, err := proto.Marshal(attestation)
	if err != nil {
		log.WithError(err).Error("marshal attest data failed")
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func Base64ToSignedAttestation(signedAttestDataBase64 string) (*ethpb.Attestation, error) {
	signedAttestData, err := base64.StdEncoding.DecodeString(signedAttestDataBase64)
	if err != nil {
		log.WithError(err).Error("base64 decode signed attest data failed")
		return nil, err
	}
	var signedAttestation = new(ethpb.Attestation)
	if err := proto.Unmarshal(signedAttestData, signedAttestation); err != nil {
		log.WithError(err).Error("unmarshal signed attest data failed")
		return nil, err
	}
	return signedAttestation, nil
}

func SignedAttestationToBase64(signedAttestation *ethpb.Attestation) (string, error) {
	data, err := proto.Marshal(signedAttestation)
	if err != nil {
		log.WithError(err).Error("marshal signed attest data failed")
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func Base64ToSignedDenebBlock(signedBlockBase64 string) (*ethpb.SignedBeaconBlockDeneb, error) {
	signedBlockData, err := base64.StdEncoding.DecodeString(signedBlockBase64)
	if err != nil {
		log.WithError(err).Error("base64 decode signed block data failed")
		return nil, err
	}
	var signedBlock = new(ethpb.SignedBeaconBlockDeneb)
	if err := proto.Unmarshal(signedBlockData, signedBlock); err != nil {
		log.WithError(err).Error("unmarshal signed block data failed")
		return nil, err
	}
	return signedBlock, nil
}

func SignedDenebBlockToBase64(signedBlock *ethpb.SignedBeaconBlockDeneb) (string, error) {
	data, err := proto.Marshal(signedBlock)
	if err != nil {
		log.WithError(err).Error("marshal signed block data failed")
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func Base64ToGenericSignedBlock(signedBlockBase64 string) (*ethpb.GenericSignedBeaconBlock, error) {
	signedBlockData, err := base64.StdEncoding.DecodeString(signedBlockBase64)
	if err != nil {
		log.WithError(err).Error("base64 decode signed block data failed")
		return nil, err
	}
	var signedBlock = new(ethpb.GenericSignedBeaconBlock)
	if err := proto.Unmarshal(signedBlockData, signedBlock); err != nil {
		log.WithError(err).Error("unmarshal signed block data failed")
		return nil, err
	}
	return signedBlock, nil
}

func GenericSignedBlockToBase64(signedBlock *ethpb.GenericSignedBeaconBlock) (string, error) {
	data, err := proto.Marshal(signedBlock)
	if err != nil {
		log.WithError(err).Error("marshal signed block data failed")
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func GetDenebBlockFromGenericSignedBlock(signedBlock *ethpb.GenericSignedBeaconBlock) (*ethpb.SignedBeaconBlockDeneb, error) {
	if signedBlock == nil {
		return nil, ErrNilObject
	}
	switch b := signedBlock.Block.(type) {
	case nil:
		return nil, ErrNilObject
	case *ethpb.GenericSignedBeaconBlock_Phase0:
		return nil, ErrUnsupportedBeaconBlock
	case *ethpb.GenericSignedBeaconBlock_Altair:
		return nil, ErrUnsupportedBeaconBlock
	case *ethpb.GenericSignedBeaconBlock_Bellatrix:
		return nil, ErrUnsupportedBeaconBlock
	case *ethpb.GenericSignedBeaconBlock_BlindedBellatrix:
		return nil, ErrUnsupportedBeaconBlock
	case *ethpb.GenericSignedBeaconBlock_Capella:
		return nil, ErrUnsupportedBeaconBlock
	case *ethpb.GenericSignedBeaconBlock_BlindedCapella:
		return nil, ErrUnsupportedBeaconBlock
	case *ethpb.GenericSignedBeaconBlock_Deneb:
		return b.Deneb.Block, ErrUnsupportedBeaconBlock
	case *ethpb.GenericSignedBeaconBlock_BlindedDeneb:
		return nil, ErrUnsupportedBeaconBlock
	default:
		log.WithError(ErrUnsupportedBeaconBlock).Errorf("unsupported beacon block from type %T", b)
		return nil, ErrUnsupportedBeaconBlock
	}
}

func GetCapellaBlockFromGenericSignedBlock(signedBlock *ethpb.GenericSignedBeaconBlock) (*ethpb.SignedBeaconBlockCapella, error) {
	if signedBlock == nil {
		return nil, ErrNilObject
	}
	switch b := signedBlock.Block.(type) {
	case nil:
		return nil, ErrNilObject
	case *ethpb.GenericSignedBeaconBlock_Phase0:
		return nil, ErrUnsupportedBeaconBlock
	case *ethpb.GenericSignedBeaconBlock_Altair:
		return nil, ErrUnsupportedBeaconBlock
	case *ethpb.GenericSignedBeaconBlock_Bellatrix:
		return nil, ErrUnsupportedBeaconBlock
	case *ethpb.GenericSignedBeaconBlock_BlindedBellatrix:
		return nil, ErrUnsupportedBeaconBlock
	case *ethpb.GenericSignedBeaconBlock_Capella:
		return b.Capella, nil
	case *ethpb.GenericSignedBeaconBlock_BlindedCapella:
		return nil, ErrUnsupportedBeaconBlock
	case *ethpb.GenericSignedBeaconBlock_Deneb:
		return nil, ErrUnsupportedBeaconBlock
	case *ethpb.GenericSignedBeaconBlock_BlindedDeneb:
		return nil, ErrUnsupportedBeaconBlock
	default:
		log.WithError(ErrUnsupportedBeaconBlock).Errorf("unsupported beacon block from type %T", b)
		return nil, ErrUnsupportedBeaconBlock
	}
}
