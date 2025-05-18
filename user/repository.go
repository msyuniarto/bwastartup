package user

import "gorm.io/gorm"

// penamaan huruf kapital di depan menandakan package lain dapat mengakses
type Repository interface {
	Save(user User) (User, error)           // parameter user dan balikannya User
	FindByEmail(email string) (User, error) // parameter email dan balikannya User
	FindByID(ID int) (User, error)          // parameter ID dan balikannya User
	Update(user User) (User, error)         // parameter user dan balikannya User
	FindAll() ([]User, error)
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

func (r *repository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindAll() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}
