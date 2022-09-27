package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float32
}

func NewOrder(id string, price float64, tax float64) (*Order, error){
	order := &Order{
		ID: id,
		Price: price,
		Tax: tax,
	} 

	err := order.IsValid()

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o Order) IsValid() error {
	if o.ID =="" {
		return errors.New("Invalid id")
	}

	if o.Price == 0 {
		return errors.New("Invalid id")
	}

	if o.Tax == 0 {
		return errors.New("Invalid id")
	}
	
	return nil
}

