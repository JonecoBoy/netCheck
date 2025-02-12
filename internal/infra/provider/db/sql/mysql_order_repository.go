package db

import (
	"database/sql"
	entity2 "github.com/jonecoboy/netCheck/internal/core/domain/entity"
	"github.com/jonecoboy/netCheck/internal/core/domain/entity/order"
	"github.com/jonecoboy/netCheck/internal/core/domain/repository"
	"github.com/jonecoboy/netCheck/pkg/entity"
)

type MySQLOrderRepository struct {
	Db *sql.DB
}

func NewMySQLOrderRepository(db *sql.DB) repository.OrderRepository {
	return &MySQLOrderRepository{Db: db}
}

func (r *MySQLOrderRepository) Save(order *order.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	return err
}

func (r *MySQLOrderRepository) FindByID(id entity.ID) (*entity2.Order, error) {
	row := r.Db.QueryRow("SELECT id, price, tax, final_price FROM orders WHERE id = ?", id)
	var order entity2.Order
	err := row.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
