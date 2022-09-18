package redisRepo

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAll(t *testing.T) {
	tests := []struct {
		testName    string
		mockPrepare func(mock redismock.ClientMock)
		key         string
		expectedErr error
	}{
		{
			testName: "All ok",
			mockPrepare: func(mock redismock.ClientMock) {
				mock.ExpectZRangeWithScores("hackers", 0, -1).SetVal([]redis.Z{{5, "Alex"}, {4, "Den"}})
			},
			key:         "hackers",
			expectedErr: nil,
		},
		{
			testName: "Not found value",
			mockPrepare: func(mock redismock.ClientMock) {
				mock.ExpectZRangeWithScores("hackers", 0, -1).SetVal([]redis.Z{})
			},
			key:         "hackers",
			expectedErr: fiber.ErrNotFound,
		},
		{
			testName: "Internal error",
			mockPrepare: func(mock redismock.ClientMock) {
				mock.ExpectZRangeWithScores("hackers", 0, -1).SetErr(redis.ErrClosed)
			},
			key:         "hackers",
			expectedErr: fiber.ErrInternalServerError,
		},
	}
	for _, test := range tests {
		client, mock := redismock.NewClientMock()
		test.mockPrepare(mock)
		redisRepo := CreateHackersRedis(client)

		_, err := redisRepo.GetAll(test.key)
		err1 := mock.ExpectationsWereMet()
		if !errors.Is(err, test.expectedErr) {
			t.Errorf("Expected error: '%v', got '%v'", test.expectedErr, err)
		}
		assert.Equal(t, nil, err1)
	}

}
