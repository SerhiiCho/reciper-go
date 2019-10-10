package store

import (
	"github.com/SerhiiCho/reciper/backend/model"
)

// UserRepo database repository
type UserRepo struct {
	store *Store
}

// Create creates new user in database
func (repo *UserRepo) Create(user *model.User) (*model.User, error) {
	result, err := repo.store.db.Exec(`INSERT INTO users (email, password) 
		VALUES (?, ?)`, user.Email, user.Password)

	if err != nil {
		return nil, err
	}

	userID, lastIdErr := result.LastInsertId()

	if lastIdErr != nil {
		return nil, lastIdErr
	}

	user.ID = uint(userID)

	return user, nil
}

// FindByEmail returns user from database with given email
func (repo *UserRepo) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	row := repo.store.db.QueryRow(`
		SELECT id, 'name', 'status', email, xp, streak_days, streak_check, password, popularity,
			active, notif_check, online_check, updated_at, created_at
				FROM users WHERE email = ?
	`, email)

	scanErr := row.Scan(
		&user.ID,
		&user.Name,
		&user.Status,
		&user.Email,
		&user.XP,
		&user.StreakDays,
		&user.StreakCheck,
		&user.Password,
		&user.Popularity,
		&user.Active,
		&user.NotifCheck,
		&user.OnlineCheck,
		&user.UpdatedAt,
		&user.CreatedAt,
	)

	if scanErr != nil {
		return nil, scanErr
	}

	return user, nil
}
