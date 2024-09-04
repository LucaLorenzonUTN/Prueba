package lib

type BodyRequest struct {
	Name   string `json:"name"`
	Pedido string `json:"pedido"`
}

type BodyResponse struct {
	Mensaje string `json:"message"`
}
