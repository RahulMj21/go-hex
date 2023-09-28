package api

import "github.com/RahulMj21/go-hex/internal/ports"

type Adapter struct {
	arith ports.ArithmeticPort
}

func NewAdapter(arith ports.ArithmeticPort) *Adapter {
	return &Adapter{arith: arith}
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}
func (apia Adapter) GetSubstraction(a, b int32) (int32, error) {
	answer, err := apia.arith.Substraction(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}
func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}
func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}
