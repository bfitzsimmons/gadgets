package gadgets

import (
	"fmt"
	"net/http"
)

// HTTPError provides a more concise way to deal with http errors.
func HTTPError(w http.ResponseWriter, msg string, err error, code int) {
	http.Error(w, fmt.Sprintf("%s --> %s", msg, err.Error()), code)
}
