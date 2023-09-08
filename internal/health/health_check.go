package health

import (
	"flightcheck/internal/pkg"
	"time"
)

func Run() {
	now := time.Now()
	for {
		if time.Since(now) > 3*time.Hour {
			now = time.Now()
			pkg.SendMessage("من زنده ام")
		}
	}
}
