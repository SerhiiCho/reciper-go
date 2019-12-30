package sqlstore

import (
	"database/sql"

	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/store"
)

// UserRepo database repository
type UserRepo struct {
	store *Store
}

// CreateUser creates new user in database
func (repo *UserRepo) CreateUser(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	result, err := repo.store.db.Exec(`INSERT INTO users (email, password) 
		VALUES (?, ?)`, user.Email, user.HashedPassword)

	if err != nil {
		return err
	}

	userID, lastIDErr := result.LastInsertId()

	if lastIDErr != nil {
		return lastIDErr
	}

	user.ID = uint(userID)

	return nil
}

// FindByEmail returns user from database with given email
func (repo *UserRepo) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	row := repo.store.db.QueryRow(`
		SELECT id, 'name', 'status', email, xp, streak_days, streak_check, password,
			popularity, active, notif_check, online_check, updated_at, created_at
		FROM users WHERE email = ?
	`, email)

	if scanErr := row.Scan(
		&user.ID,
		&user.Name,
		&user.Status,
		&user.Email,
		&user.XP,
		&user.StreakDays,
		&user.StreakCheck,
		&user.HashedPassword,
		&user.Popularity,
		&user.Active,
		&user.NotifCheck,
		&user.OnlineCheck,
		&user.UpdatedAt,
		&user.CreatedAt,
	); scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, scanErr
	}

	return user, nil
}

// FindUser returns user from database with given id
func (repo *UserRepo) FindUser(id uint) (*model.User, error) {
	user := &model.User{}

	row := repo.store.db.QueryRow(`
		SELECT id, 'name', 'status', email, xp, streak_days, streak_check, password,
			popularity, active, notif_check, online_check, updated_at, created_at
		FROM users WHERE id = ?
	`, id)

	if scanErr := row.Scan(
		&user.ID,
		&user.Name,
		&user.Status,
		&user.Email,
		&user.XP,
		&user.StreakDays,
		&user.StreakCheck,
		&user.HashedPassword,
		&user.Popularity,
		&user.Active,
		&user.NotifCheck,
		&user.OnlineCheck,
		&user.UpdatedAt,
		&user.CreatedAt,
	); scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, scanErr
	}

	return user, nil
}
