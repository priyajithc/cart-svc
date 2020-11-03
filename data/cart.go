package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Product defines the structure for an API product
type Cart struct {
	ID          int           `json:"id"`
	CartDetails []CartProduct `json:"cartDetails"`
}

type CartProduct struct {
	ProdID string `json:"productID"`
	SKU    string `json:"sku"`
	Qty    int    `json:"qty"`
}

var cartMap map[int]Cart

func AddToCart(id int, cps []CartProduct) error {
	fmt.Println("Inside AddToCart", id)
	cart := Cart{}
	if val, ok := cartMap[id]; ok {
		cart = val
		cart.CartDetails = cps
	} else {
		cart.ID = id
		cart.CartDetails = cps
	}
	if cartMap == nil {
		cartMap = make(map[int]Cart)
	}
	cartMap[id] = cart
	fmt.Println(cartMap)
	return nil
}

func GetCart(id int) (Cart, error) {
	fmt.Println("Inside GetCart Data", id)
	fmt.Println(cartMap)
	if val, ok := cartMap[id]; ok {
		return val, nil
	}
	return Cart{}, errors.New("Cart not found")
}

func (c Cart) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}
