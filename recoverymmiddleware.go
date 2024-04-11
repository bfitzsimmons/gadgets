package gadgets

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// RecoveryMiddleware recovers from any panics experienced during the handling of web requests.
func RecoveryMiddleware(next http.Handler) http.Handler {
	recoveryFunc := func(w http.ResponseWriter, r *http.Request) {
		// This will execute when the function returns for *any* reason.
		defer func() {
			recovered := recover()
			if recovered != nil {
				// Log the error and return a 500.
				HTTPError(
					w,
					"Something went very wrong",
					fmt.Errorf("panic recovery :: %+v", recovered),
					http.StatusInternalServerError,
				)

				// Print the stacktrace.
				debug.PrintStack()
			}
		}()

		// Send the request on to the next handler.
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(recoveryFunc)
}
