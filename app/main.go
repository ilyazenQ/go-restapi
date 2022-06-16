package main

import (
	"payment_gateway/controller"
	"payment_gateway/model"
)

func main() {
	model.Init()
	controller.Start()
}
