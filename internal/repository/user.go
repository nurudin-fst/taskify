package repository

import "gorm.io/gorm"

type User struct {
	Id       int    `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (*User) TableName() string {
	return "users"
}

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (r *UserRepository) GetUserByEmail(email string) (User, error) {
	var user User
	err := r.Db.Select(`id, name, email, password`).
		Table("users").
		Take(&user, "email = ?", email).Error

	return user, err
}

func (r *UserRepository) Insert(user User) error {
	err := r.Db.Create(&user).Error
	return err
}
