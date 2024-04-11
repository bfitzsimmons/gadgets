package gadgets

import (
	"log"
	"net/http"
	"time"
)

// TimerMiddleware times the request/response time.
func TimerMiddleware(next http.Handler) http.Handler {
	timerFunc := func(w http.ResponseWriter, r *http.Request) {
		/*----------------------------------------------------------------------
		Request side.
		----------------------------------------------------------------------*/
		// Start the timer.
		startTime := time.Now()

		// This will print the execution time when the function returns for
		// *any* reason.
		defer func() {
			log.Println(time.Since(startTime))
		}()

		// Run the other middleware/handlers.
		next.ServeHTTP(w, r)

		/*----------------------------------------------------------------------
		Response side.
		----------------------------------------------------------------------*/
	}

	return http.HandlerFunc(timerFunc)
}
