package database

import (
	"fmt"

	"github.com/j-ew-s/ms-curso-ordering-api/ordering-serivces/models"
)

func (sqlCommand SQLCommand) InsertOrder(order *models.Order) error {

	db, err := sqlCommand.OpenConnection()
	if err != nil {
		fmt.Println(fmt.Printf("[ERROR] :: Inserting ORDER ::: Could not Open Connection :::   %+v ::", err))
		return err
	}

	err = db.Table("orders").Create(&order).Error
	if err != nil {
		fmt.Println(fmt.Printf("[ERROR] :: Inserting ORDER :::  %+v ::", err))
		return err
	}

	return err
}

func (sqlCommand SQLCommand) InsertItem(item *models.Item, orderId int32) error {

	item.OrderId = orderId

	db, err := sqlCommand.OpenConnection()
	if err != nil {
		fmt.Println(fmt.Printf("[ERROR] ::  Inserting ITEM ::::  %+v ::", err))
		return err
	}

	err = db.Table("orders_items").Create(&item).Error
	if err != nil {
		fmt.Println(fmt.Printf("[ERROR] :: Inserting ITEM ::::  %+v ::", err))
		return err
	}

	return err
}

func (sqlCommand SQLCommand) GetAllOrders() ([]*models.Orders, error) {

	return nil, nil
}

func (sqlCommand SQLCommand) GetOrderByIDS(id string) (order *models.Orders, err error) {

	return nil, nil
}
