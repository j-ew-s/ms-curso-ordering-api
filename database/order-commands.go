package database

import (
	"fmt"

	"github.com/j-ew-s/ms-curso-ordering-api/ordering-serivces/models"
	"gorm.io/gorm/clause"
)

func (sqlCommand SQLCommand) InsertOrder(order *models.Order) error {

	db, err := sqlCommand.OpenConnection()
	if err != nil {
		fmt.Println(fmt.Printf("[ERROR] :: InsertOrder ::: Could not Open Connection :::   %+v ::", err))
		return err
	}

	err = db.Table("orders").Create(&order).Error
	if err != nil {
		fmt.Println(fmt.Printf("[ERROR] :: InsertOrder:::  %+v ::", err))
		return err
	}

	return err
}

func (sqlCommand SQLCommand) InsertItem(item *models.Item, orderId int32) error {

	item.OrderId = orderId

	db, err := sqlCommand.OpenConnection()
	if err != nil {
		fmt.Println(fmt.Printf("[ERROR] ::  InsertItem ::::  %+v ::", err))
		return err
	}

	err = db.Table("orders_items").Create(&item).Error
	if err != nil {
		fmt.Println(fmt.Printf("[ERROR] :: InsertItem ::::  %+v ::", err))
		return err
	}

	return err
}

func (sqlCommand SQLCommand) GetOrdersById(orderId string) ([]*models.Order, error) {

	db, err := sqlCommand.OpenConnection()
	if err != nil {
		fmt.Println(fmt.Printf("[ERROR] :: GetOrdersById ::: Could not Open Connection :::   %+v ::", err))
		return nil, err
	}
	orders := make([]*models.Order, 0)

	db.
		Preload(clause.Associations).
		Where("id = ?", orderId).
		Find(&orders)

	if err != nil {
		fmt.Println(fmt.Printf("[ERROR] :: GetOrdersById :::  %+v ::", err))
		return nil, err
	}

	return orders, err
}

func (sqlCommand SQLCommand) GetOrdersByUserId(orderId string) ([]*models.Order, error) {

	db, err := sqlCommand.OpenConnection()
	if err != nil {
		fmt.Println(fmt.Printf("[ERROR] :: GetOrdersByUserId  ::: Could not Open Connection :::   %+v ::", err))
		return nil, err
	}
	orders := make([]*models.Order, 0)

	db.
		Preload(clause.Associations).
		Where("user_id = ?", orderId).
		Find(&orders)

	if err != nil {
		fmt.Println(fmt.Printf("[ERROR] :: GetOrdersByUserId :::  %+v ::", err))
		return nil, err
	}

	return orders, err
}
