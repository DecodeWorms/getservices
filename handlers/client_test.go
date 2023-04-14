package handlers

import (
	"testing"

	"github.com/DecodeWorms/getservices/mocks"
	"github.com/DecodeWorms/getservices/models"
	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	//the ctrl.Finish is deprecated because an instance of *testing.T
	//is already passed to gomock.NewController(t)
	defer ctrl.Finish()

	mockClientServices := mocks.NewMockClientServices(ctrl)

	cl := models.Client{
		FirstName:          "Abdulhameed",
		LastName:           "Musa",
		PhoneNumber:        "0800000000",
		Email:              "lekzy.csharp@gmail.com",
		Password:           "harvest600",
		IsFullyOnboarded:   true,
		IsAccountConfirmed: true,
		Pin:                "099745",
	}

	mockClientServices.EXPECT().Create(cl).Return(nil)
	if err := mockClientServices.Create(cl); err != nil {
		t.Errorf("Error creating client: %v", err)
	}
}
