package runtime

const (
	attackerFlag          = "attacker"
	maxValidatorIndexFlag = "max-validator-index"
)

type updateParam struct {
	attacker          string
	maxValidatorIndex int
}

var (
	params = &updateParam{
		attacker:          "",
		maxValidatorIndex: -1,
	}
)
