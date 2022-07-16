package storage

import (
	"errors"
	

	"github.com/MelvinRB27/server-user/models"
)

func Migrate() {
	DB().AutoMigrate(&models.User{})
}

var (
	ErrPasswordIncorrect = errors.New("password incorrect")
	ErrUserNoFound = errors.New("user not found")
	ErrCantCreate = errors.New("user already exists already")
	ErrCantUpdate = errors.New("data could not be updated")
)

//Create is used for create a record of the database
func Register(m *models.User) error {
	user := models.User{
		Name:        m.Name,
		LastName: m.LastName,
		UserName: m.UserName,
		Gender: m.Gender,
		Rol: m.Rol,
		Password: m.Password,
		PasswordConfirm: m.PasswordConfirm,
	}
	 
	result := DB().Create(&user)

	if err := result.Error; err != nil {
		return ErrCantCreate
	} 

	return nil
}

func Login(userName, password string) (models.User, error) {
	userDta := models.User{}

	user := make([]models.User, 0)

	//search for the user by email
	if  err := DB().Where("user_name = ?", userName).First(&user).Error; err != nil {
		//if no found, return UserNoFound
		return userDta, ErrUserNoFound
	}
	
	for _, person := range user {
		//validate if password match
		if password != person.Password{
			return userDta, ErrPasswordIncorrect
		}
		//if user exist, return user data
		userDta = models.User{
			Name: person.Name,
			LastName: person.LastName,
			UserName: person.UserName,
			Gender: person.Gender,
			Rol: person.Rol,
		}
	}

	return userDta, nil
}

func Update(m *models.User) (models.User, error ){
	user := models.User{}

	user.UserName = m.UserName
	err := DB().Model(&user).Where(models.User{UserName: m.UserName}).Updates(models.User{Name: m.Name, LastName: m.LastName, UserName: m.UserName,
		Rol: m.Rol, Gender: m.Gender,Password: m.Password, PasswordConfirm: m.PasswordConfirm})
		DB().Save(err)
	if err.Error != nil {
		return user, ErrCantUpdate
	}

	//if user exist, return user data
	user = models.User{
		Name: m.Name,
		LastName: m.LastName,
		UserName: m.UserName,
		Gender: m.Gender,
		Rol: m.Rol,
	}

	return user, nil
}