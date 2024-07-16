package order

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/rohithrajasekharan/go-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	var id int
	err := s.db.QueryRow("INSERT INTO orders (userId, total, status, address) VALUES ($1, $2, $3, $4) RETURNING id",
		order.UserID, order.Total, order.Status, order.Address).Scan(&id)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return 0, fmt.Errorf("failed to create order: %s (Code: %s)", pgErr.Message, pgErr.Code)
		}
		return 0, fmt.Errorf("failed to create order: %w", err)
	}
	return id, nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (orderId, productId, quantity, price) VALUES ($1, $2, $3, $4)",
		orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}
