package middleware

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"log"
	"wb-l0/internal"
)

func Validate(body []byte) (string, error) {
	var order internal.Order

	err := json.Unmarshal(body, &order)
	if err != nil {
		return "", &internal.Error{
			Message: err.Error(),
			Code:    500,
		}
	}
	if err = validator.Validate(order); err != nil {
		errs := err.(validator.ErrorMap)
		log.Println(errs)
		return "", &internal.Error{
			Message: "Bad data",
			Code:    500,
		}
	}
	return order.OrderUid, nil
}
