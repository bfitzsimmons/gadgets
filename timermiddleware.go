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

		// Run the other middleware/handlers.
		next.ServeHTTP(w, r)

		// Print the execution time.
		go func() {
			log.Println(time.Since(startTime))
		}()
	}

	return http.HandlerFunc(timerFunc)
}
