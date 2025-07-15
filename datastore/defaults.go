package datastore

import "github.com/gnanasuryateja-kurapati/simple-rest-api-golang/model"

var Users map[string]model.UserData
var defaultAge = 30

func LoadSampleUsers() {
	Users = make(map[string]model.UserData)
	sampleData1 := model.UserData{
		UserID:  "user_1",
		Name:    "SampleName1",
		Age:     &defaultAge,
		Email:   "sampleEmail1@usermanagement.com",
		Address: "House No, Road No, Locality, Area, City, State - Pincode",
	}
	Users[sampleData1.UserID] = sampleData1
	sampleData2 := model.UserData{
		UserID:  "user_2",
		Name:    "SampleName2",
		Age:     &defaultAge,
		Email:   "sampleEmail2@usermanagement.com",
		Address: "House No, Road No, Locality, Area, City, State - Pincode",
	}
	Users[sampleData2.UserID] = sampleData2
}
