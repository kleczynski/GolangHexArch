package api

import "GoHexArchTutorial/internal/ports"

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort // How Application layer connect to core layer
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

func (apia Adapter) GetAddition(a int32, b int32) (int32, error) {
	answer, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetSubtraction(a int32, b int32) (int32, error) {
	answer, err := apia.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "Subtraction")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetMultiplication(a int32, b int32) (int32, error) {
	answer, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "Multiplication")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetDivision(a int32, b int32) (int32, error) {
	answer, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "Division")
	if err != nil {
		return 0, err
	}
	return answer, nil
}
