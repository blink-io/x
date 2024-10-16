package snowflake

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/snowflake"
)

func TestSnowflake_1(t *testing.T) {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a snowflake UserID.
	id := node.Generate()

	// Print out the UserID in a few different ways.
	fmt.Printf("Int64  UserID: %d\n", id)
	fmt.Printf("String UserID: %s\n", id)
	fmt.Printf("Base2  UserID: %s\n", id.Base2())
	fmt.Printf("Base64 UserID: %s\n", id.Base64())

	// Print out the UserID's timestamp
	fmt.Printf("UserID Time  : %d\n", id.Time())

	// Print out the UserID's node number
	fmt.Printf("UserID Node  : %d\n", id.Node())

	// Print out the UserID's sequence number
	fmt.Printf("UserID Step  : %d\n", id.Step())
}
