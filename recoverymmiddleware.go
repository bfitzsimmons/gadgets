package gadgets

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// RecoveryMiddleware recovers from any panics experienced during the handling of web requests.
func RecoveryMiddleware(next http.Handler) http.Handler {
	recoveryFunc := func(w http.ResponseWriter, r *http.Request) {
		/*----------------------------------------------------------------------------
		Request side.
		----------------------------------------------------------------------------*/
		// This will execute at the end of the function.
		defer func() {
			if rval := recover(); rval != nil {
				// Log the error and return a 500.
				HTTPError(w, "Something went very wrong", fmt.Errorf("Panic recovery :: %+v", rval),
					http.StatusInternalServerError)

				// Print the stacktrace.
				debug.PrintStack()
			}
		}()

		// Send the request on to the next handler.
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(recoveryFunc)
}
