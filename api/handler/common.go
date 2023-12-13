package handler
import (
	"net/http"
	"github.com/gin-gonic/gin"
    "github.com/syndtr/goleveldb/leveldb"
	"fmt"
)

// ResponseBody is the structure for the JSON response
type ResponseBody struct {
	Status bool   `json:"status"`
	Data   string `json:"data"`
	Description string `json:"description"`
}

func formatErrorMessage(err error) string {
    if err == nil {
        return "No error"
    }
    return fmt.Sprintf("Failed to add execution layer: %v", err.Error())
}

// respondWithError sends a JSON error response
func respondWithError(c *gin.Context,error_msg string) {
	response := ResponseBody{
		Status : false,
		Data : "",
		Description : error_msg,
	}
	c.JSON(http.StatusBadRequest, response)
	return
}

// respondWithJSON sends a JSON response
func respondWithSuccess(c *gin.Context, data string, description string) {
	response := ResponseBody{
		Status : true,
		Data : data,
		Description : description,
	}
	c.JSON(http.StatusOK, response)
	return
}

func storeData(db *leveldb.DB, key, value string) error {
	// Store the key-value pair in LevelDB
	err := db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		return err
	}
	return nil
}

func retrieveData(db *leveldb.DB, key string) (string, error) {
	// Retrieve the value for the given key from LevelDB
	data, err := db.Get([]byte(key), nil)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
