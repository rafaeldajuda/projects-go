package services

import (
	"estoque/internal/models"
	"fmt"
	"strconv"
	"time"
)

type Estoque struct {
	items map[string]models.Item
	logs  []models.Log
}

func NewEstoque() *Estoque {
	return &Estoque{
		items: make(map[string]models.Item),
		logs:  []models.Log{},
	}
}

func (e *Estoque) AddItem(item models.Item, user string) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("erro ao adicionar item: [ID:%d] possui uma quantidade inválida (zero ou negativa)", item.ID)
	}

	existingItem, exists := e.items[strconv.Itoa(item.ID)]
	if exists {
		item.Quantity += existingItem.Quantity
	}

	e.logs = append(e.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Entrada de estoque",
		User:      user,
		ItemID:    item.ID,
		Quantity:  item.Quantity,
		Reason:    "Adicionando novos itens no estoque",
	})

	e.items[strconv.Itoa(item.ID)] = item
	return nil
}

func (e *Estoque) RemoveItem(itemID int, quantity int, user string) error {
	// Valida a existêcnia do item no estoque
	existingItem, exist := e.items[strconv.Itoa(itemID)]
	if !exist {
		return fmt.Errorf("erro ao remover item: [ID:%d] não existe no estoque", itemID)
	}

	// Checa a quantidade a ser removida é válida
	if quantity <= 0 {
		return fmt.Errorf("erro ao remover item: quantidade inválida (zero ou negativa) para [ID:%d]", itemID)
	}

	// Checa a quantidade disponível no estoque é suficiente
	if existingItem.Quantity < quantity {
		return fmt.Errorf("erro ao remover item: estoque insuficiente para [ID:%d]. Disponível: %d, Solicitado: %d", itemID, existingItem.Quantity, quantity)
	}

	// Atualize o estoque removendo a quantidade informada
	existingItem.Quantity -= quantity
	if existingItem.Quantity == 0 {
		// Remova o item completamente se a quantidade for zero
		delete(e.items, strconv.Itoa(itemID))
	} else {
		// Atualize a quantidade do item no estoque
		e.items[strconv.Itoa(itemID)] = existingItem
	}

	// Registre a operação em um log
	e.logs = append(e.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Saída de estoque",
		User:      user,
		ItemID:    itemID,
		Quantity:  quantity,
		Reason:    "Removendo itens do estoque",
	})

	return nil
}

func (e *Estoque) ListItems() []models.Item {
	var itemList []models.Item
	for _, item := range e.items {
		itemList = append(itemList, item)
	}

	return itemList
}

func (e *Estoque) ViewAuditLog() []models.Log {
	return e.logs
}

func (e *Estoque) CalculateTotalCost() float64 {
	var totalCost float64
	for _, item := range e.items {
		totalCost += float64(item.Quantity) * item.Price
	}
	return totalCost
}

// MINHAS FUNÇÕES	/////////////////////////

func (e *Estoque) ListLogs() []string {
	logs := []string{}
	for _, log := range e.logs {
		logs = append(logs, fmt.Sprintf("timestamp: %s | action: %s | user: %s | itemID: %d | quantity: %d | reason: %s",
			log.Timestamp, log.Action, log.User, log.ItemID, log.Quantity, log.Reason))
	}
	return logs
}

// func (e *Estoque) RemoveItem(itemID int, quantity int, user string) bool {
// 	if user == "" {
// 		user = "usuário desconhecido"
// 	}

// 	for index, item := range e.items {
// 		if item.ID == itemID {
// 			if quantity < 0 {
// 				e.logs = append(e.logs, models.Log{
// 					Timestamp: time.Now(),
// 					Action:    "Saída de estoque",
// 					User:      user,
// 					ItemID:    itemID,
// 					Quantity:  quantity,
// 					Reason:    "Valor da quantidade inválida",
// 				})
// 				return false
// 			}
// 			if item.Quantity < quantity {
// 				e.logs = append(e.logs, models.Log{
// 					Timestamp: time.Now(),
// 					Action:    "Saída de estoque",
// 					User:      user,
// 					ItemID:    itemID,
// 					Quantity:  quantity,
// 					Reason:    "Quantidade insuficiente para a remoção do item",
// 				})
// 				return false
// 			}
// 			newQuantity := item.Quantity - quantity
// 			newItem := item
// 			newItem.Quantity = newQuantity
// 			e.items[index] = newItem

// 			e.logs = append(e.logs, models.Log{
// 				Timestamp: time.Now(),
// 				Action:    "Saída de estoque",
// 				User:      user,
// 				ItemID:    itemID,
// 				Quantity:  newQuantity,
// 				Reason:    "Itens removidos do estoque",
// 			})

// 			return true
// 		}
// 	}

// 	e.logs = append(e.logs, models.Log{
// 		Timestamp: time.Now(),
// 		Action:    "Saída de estoque",
// 		User:      user,
// 		ItemID:    itemID,
// 		Quantity:  quantity,
// 		Reason:    "Item não encontrado no estoque",
// 	})
// 	return false
// }
