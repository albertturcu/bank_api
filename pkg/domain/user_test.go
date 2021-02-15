package domain

import (
	"os"
	"restAPI/mocks"
	"restAPI/pkg/storage/mysql/entity"
	"strconv"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

var mockDBRepository *mocks.DBRepository
var mockDService DService

func TestMain(m *testing.M) {
	mockDBRepository = &mocks.DBRepository{}
	mockDService = &dService{r: mockDBRepository}
	os.Exit(m.Run())
}

//test for addUser function
func TestAddUser(t *testing.T) {
	assert := assert.New(t)
	testUser := entity.User{}
	err := faker.FakeData(&testUser)

	if err != nil {
		t.Errorf("%v", err)
	}
	mockDBRepository.On("AddUser", testUser).Return(testUser, nil)

	actualUser, err := mockDService.AddUser(testUser)

	if err != nil {
		t.Errorf("%v", err)
	}

	assert.Equal(actualUser, testUser)
	assert.NotEqual(entity.User{}, testUser)
	assert.NotEqual(entity.User{}, actualUser)
}

func TestGetUser(t *testing.T) {
	assert := assert.New(t)
	testUser := entity.User{}
	err := faker.FakeData(&testUser)
	if err != nil {
		t.Errorf("%v", err)
	}
	id := strconv.FormatUint(uint64(testUser.ID), 10)
	mockDBRepository.
		On("GetUser", id).Return(testUser, nil)

	actualUser, err := mockDService.GetUser(id)

	if err != nil {
		t.Errorf("%v", err)
	}

	assert.Equal(actualUser, testUser)
}
