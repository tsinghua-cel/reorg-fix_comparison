package library

import (
	"github.com/tsinghua-cel/strategy-gen/utils"
	"strconv"
)

func CheckDuties(maxValidatorIndex int, duties []utils.ProposerDuty) ([]utils.ProposerDuty, bool) {
	maxLength := 0
	maxIndex := 0
	length := 0
	index := 0
	for i, duty := range duties {
		valIdx, _ := strconv.Atoi(duty.ValidatorIndex)

		if valIdx <= maxValidatorIndex {
			length++
			if length > maxLength {
				maxLength = length
				maxIndex = index
			}
		} else {
			length = 0
			index = i + 1
		}
	}
	if maxLength > 2 {
		return duties[maxIndex : maxIndex+maxLength], true
	}
	return nil, false
}
