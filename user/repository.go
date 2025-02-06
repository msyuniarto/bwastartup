package user

import "gorm.io/gorm"

// penamaan huruf kapital di depan menandakan package lain dapat mengakses
type Repository interface {
	Save(user User) (User, error)           // parameter user dan balikannya User
	FindByEmail(email string) (User, error) // parameter user dan balikannya User
}

// penamaan huruf kecil di depan menandakan package lain tidak dapat mengakses langsung (private)
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	// create object repository (struct yg sudah dibuat)
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
