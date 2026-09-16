package order

import (
	"github.com/Tolan4ik/sprint1-demo/internal/user"
)

type Order struct {
	ID    string
	Owner user.User
}
