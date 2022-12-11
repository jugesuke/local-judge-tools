package usecase

import (
	"errors"
)

func (u *Usecase) Run(question string) error {
	if err := u.domain.Build(question); err != nil {
		return errors.New("build was failed")
	}
	if err := u.domain.Run(question); err != nil {
		return errors.New("execute was failed")
	}
	return nil
}
