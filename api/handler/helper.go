package handler 
import (
	"net/http"
	"github.com/gin-gonic/gin"
    "github.com/syndtr/goleveldb/leveldb"
	"fmt"
	"github.com/airchains-studio/settlement_layer_calls_api/model"
)

func formatErrorMessage(err error) string {
    if err == nil {
        return "No error"
    }
    return fmt.Sprintf(err.Error())
}

// respondWithError sends a JSON error response
func respondWithError(c *gin.Context,error_msg string) {
	response := model.ResponseBody{
		Status : false,
		Data : "",
		Description : error_msg,
	}
	c.JSON(http.StatusBadRequest, response)
	return
}

// respondWithJSON sends a JSON response
func respondWithSuccess(c *gin.Context, data string, description string) {
	response := model.ResponseBody{
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
