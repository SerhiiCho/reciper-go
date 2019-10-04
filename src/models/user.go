package models

type User struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Status       string  `json:"status"`
	Email        string  `json:"email"`
	Username     string  `json:"username"`
	XP           uint    `json:"xp"`
	StreakDays   uint    `json:"streak_days"`
	Popularity   float64 `json:"popularity"`
	Active       bool    `json:"active"`
	StreakCheck  string  `json:"streak_check"`
	NotifCheck   string  `json:"notif_check"`
	OnlineCheck  string  `json:"online_check"`
	ContactCheck string  `json:"contact_check"`
	CreatedAt    string  `json:"created_at"`
	Photo        string  `json:"photo"`
	password     string  `json:"password"`
	token        string  `json:"token"`
}
