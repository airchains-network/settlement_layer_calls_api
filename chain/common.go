package chain 

import (
	"fmt"
)

func formatErrorMessage(err error) string {
    if err == nil {
        return "No error"
    }
    return fmt.Sprintf("Failed to add execution layer: %v", err.Error())
}
