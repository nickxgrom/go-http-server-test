package store

import "main/internal/app/model"

type UserRepository struct {
	Store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	err := u.BeforeCreate()
	if err != nil {
		return nil, err
	}

	if err := r.Store.db.QueryRow(
		"insert into users(email, encrypted_password) values($1, $2) returning id",
		&u.Email,
		&u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.Store.db.QueryRow(
		"select id, email, encrypted_password from users where email = $1",
		email,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		return nil, err
	}

	return u, nil
}
