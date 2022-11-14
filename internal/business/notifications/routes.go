package notifications

import "net/http"

func Routes(nh *NotificationHandler) {
	http.HandleFunc("/notifications", nh.CreateNotification)
}
