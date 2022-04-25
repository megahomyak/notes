package workers

import (
	"notes/src/models"
	"notes/src/logging"
	"time"
)

func DeleteExpiredTokensPeriodically() {
	ticker := time.NewTicker(time.Hour * 24)
	for {
		if err := models.DB.Delete(&models.AccessToken{}, "expires_in < ?", time.Now()).Error; err != nil {
			logging.LogError(err)
		}
		<- ticker.C
	}
}
