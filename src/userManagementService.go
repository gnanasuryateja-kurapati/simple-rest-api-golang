package src

import (
	"errors"

	"github.com/gnanasuryateja-kurapati/simple-rest-api-golang/datastore"
	"github.com/gnanasuryateja-kurapati/simple-rest-api-golang/model"
)

func GetUsers() map[string]model.UserData {
	return datastore.Users
}

func GetUserById(userId string) (*model.UserData, error) {
	userData, ok := datastore.Users[userId]
	if ok {
		return &userData, nil
	}
	return nil, errors.New("user not found with given userId")
}

func generateUserId(name, email string) string {
	return email + "_" + name
}

func CreateNewUser(userData model.UserDataInput) (*string, error) {
	userId := generateUserId(userData.Email, userData.Name)
	_, err := GetUserById(userId)
	if err != nil {
		newUser := model.UserData{
			UserID:  userId,
			Name:    userData.Name,
			Age:     userData.Age,
			Email:   userData.Email,
			Address: userData.Address,
		}
		datastore.Users[userId] = newUser
		ackMsg := "new user created with userId: " + newUser.UserID
		return &ackMsg, nil
	}
	return nil, errors.New("user already exists with the given information")
}

func UpdateUser(userData model.UserDataInput) (*string, error) {
	userId := generateUserId(userData.Email, userData.Name)
	user, err := GetUserById(userId)
	if err != nil {
		return nil, err
	}
	user.Age = userData.Age
	user.Address = userData.Address
	datastore.Users[userId] = *user
	acgMsg := "data is updated for user with id:" + userId
	return &acgMsg, nil
}

func DeleteUser(userId string) (*string, error) {
	_, err := GetUserById(userId)
	if err != nil {
		return nil, err
	}
	delete(datastore.Users, userId)
	ackMsg := "user with id:" + userId + "has been deleted"
	return &ackMsg, nil
}
