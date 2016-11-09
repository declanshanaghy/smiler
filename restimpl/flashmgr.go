package restimpl

import (
	"github.com/declanshanaghy/smiler/models"
)

type FlashMgr struct {}

func NewFlashMgr() *FlashMgr {
	return &FlashMgr{}
}

func (f *FlashMgr) GetFlash() (models.FlashState, error) {
	x := int64(10)
	state := models.FlashState{}
	state.Freq = &x
	return state, nil
}

func (f *FlashMgr) SetFlash(state models.FlashState) (models.FlashState, error) {
	return state, nil
}