package utils

//code to unmarshal json
import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, v interface{}) error {
	if body, err := io.ReadAll(r.Body); err != nil {
		return err
	} else {
		err = json.Unmarshal(body, v)
		if err != nil {
			return err
		}
	}
	return nil

}
