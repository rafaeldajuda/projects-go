package main

import (
	"estoque/internal/models"
	"estoque/internal/services"
	"fmt"
)

func main() {
	fmt.Println("Sistema de Estoque\n")
	estoque := services.NewEstoque()

	itens := []models.Item{
		{ID: 1, Name: "Fone", Quantity: 10, Price: 100},
		{ID: 2, Name: "Camiseta", Quantity: 1, Price: 55.99},
		{ID: 3, Name: "Mouse", Quantity: 2, Price: 12.99},
	}

	for _, item := range itens {
		err := estoque.AddItem(item, "Rafael")
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	for _, item := range estoque.ListItems() {
		fmt.Printf("ID: %d | item: %s | Quantidade: %d | Preço: %.2f\n", item.ID, item.Name, item.Quantity, item.Price)
	}

	fmt.Println("Valor total do estoque R$:", estoque.CalculateTotalCost())

	err := estoque.RemoveItem(1, 3, "Rafael")
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	err = estoque.RemoveItem(2, 10, "Rafael")
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	err = estoque.RemoveItem(3, -2, "")
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	err = estoque.RemoveItem(4, 1, "")
	if err != nil {
		fmt.Println("err:", err.Error())
	}

	for _, log := range estoque.ViewAuditLog() {
		fmt.Printf("[%s] Ação: %s - Usuário: %s - Item ID: %d - Quantidade: %d - Motivo: %s\n",
			log.Timestamp.Format("01/02 15:04:05"), log.Action, log.User, log.ItemID, log.Quantity, log.Reason)
	}

	for _, item := range estoque.ListItems() {
		fmt.Printf("ID: %d | item: %s | Quantidade: %d | Preço: %.2f\n", item.ID, item.Name, item.Quantity, item.Price)
	}

}
