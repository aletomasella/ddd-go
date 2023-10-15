package entities

// THIS IS VALUE OBJECT - IT CANNOT CHANGE
import "time"

type Transaction struct {
	id int
	// The users IDs that made the transaction
	from     int
	to       int
	createAt time.Time
}
