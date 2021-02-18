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
	type TestCase struct {
		ID    string
		User  entity.User
		Error error
	}

	testUser := entity.User{}
	err := faker.FakeData(&testUser)
	if err != nil {
		t.Errorf("%v", err)
	}
	id := strconv.FormatUint(uint64(testUser.ID), 10)

	testCases := map[string]TestCase{
		"zeroValue":     {ID: "0", User: entity.User{}, Error: nil},
		"success":       {ID: id, User: testUser, Error: nil},
		"negativeValue": {ID: "-32", User: entity.User{}, Error: nil},
	}

	for name, tc := range testCases {
		mockDBRepository.On("GetUser", tc.ID).Return(tc.User, tc.Error)
		output, err := mockDService.GetUser(tc.ID)
		if err != nil {
			t.Errorf("%v", err)
		}
		if !assert.Equal(t, output, tc.User) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.ID, tc.User)
		}
	}
}
