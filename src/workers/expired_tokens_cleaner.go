package workers

import (
	"notes/src/models"
	"time"
)

func DeleteExpiredTokensPeriodically() {
	ticker := time.NewTicker(time.Hour * 24)
	for {
		models.DB.Delete(&models.AccessToken{}, "expires_in < ?", time.Now())
		<- ticker.C
	}
}
