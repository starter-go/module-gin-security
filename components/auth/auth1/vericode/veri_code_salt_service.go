package vericode

import (
	"crypto/rand"

	"github.com/starter-go/base/lang"
)

type saltGeneratingService struct {
	value []byte
}

func (inst *saltGeneratingService) init() error {
	buffer := make([]byte, 128)
	rand.Read(buffer)
	inst.value = buffer
	return nil
}

func (inst *saltGeneratingService) salt(t1 lang.Time) lang.Base64 {
	return lang.Base64FromBytes(inst.value)
}
