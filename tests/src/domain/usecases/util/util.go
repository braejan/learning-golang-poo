package usecases

import "errors"

type UtilUsecase struct{}

type UtilUsecases interface {
	ReverseString(text string) (reverse string, err error)
	// IsPalindrome(text string) (result bool)
}

func NewUtilUsecase() (utilUsecase *UtilUsecase) {
	utilUsecase = &UtilUsecase{}
	return
}

func (uu *UtilUsecase) ReverseString(text string) (reverse string, err error) {
	if text == "" {
		err = errors.New("reverse string required a valid input string")
		return
	}
	toByteArray := []byte(text)
	reverseArray := []byte(text)
	length := len(toByteArray)
	for i := length; i > 0; i-- {
		reverseArray[length-i] = toByteArray[i-1]
	}
	reverse = string(reverseArray)
	return
}
