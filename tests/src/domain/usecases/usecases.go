package usecases

import usecases "github.com/braejan/learning-golang-poo/tests/src/domain/usecases/util"

type GeneralUsecase struct {
	usecases.UtilUsecases
}

func NewGeneralUsecases() (generalUsecases *GeneralUsecase) {
	generalUsecases = &GeneralUsecase{
		usecases.NewUtilUsecase(),
	}
	return
}
