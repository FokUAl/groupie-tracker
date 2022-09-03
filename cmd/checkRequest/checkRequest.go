package checkRequest

import (
	"fmt"
	"net/http"
)

func CheckStatus(status int) string {
	switch status {
	case http.StatusNotFound:
		return fmt.Sprintf("%d %s", http.StatusNotFound, http.StatusText(http.StatusNotFound))
	case http.StatusInternalServerError:
		return fmt.Sprintf("%d %s", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	return ""
}
