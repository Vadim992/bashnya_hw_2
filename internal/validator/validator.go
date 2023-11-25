package validator

import "AppCLI/customerr"

type Validator struct{}

// CheckFlagsCDU checks that we dont use flags -c, -d, -u in the same time

func (v *Validator) CheckFlagsCDU(m map[string]struct{}) error {

	if len(m) > 1 {
		return customerr.ErrUseFlag
	}

	return nil
}
