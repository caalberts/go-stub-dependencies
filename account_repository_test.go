package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type stubRedis struct {
	data map[string]string
	err error
}

func (stub *stubRedis) HMSet(key string, data map[string]interface{}) (string, error) {
	for k, v := range data {
		stub.data[k] = v.(string)
	}
	return "", nil
}

func (stub *stubRedis) HGetAll(key string) (map[string]string, error) {
	return stub.data, nil
}

const id = "abcd1234"

var fakeData = map[string]string{
	"Id":      id,
	"Name":    "John Doe",
	"Balance": "100",
}

func TestFetchAccountWhenAccountExists(t *testing.T) {
	repository := createRepositoryWithData(fakeData)
	account, err := repository.FetchAccount(id)

	assert.Equal(t, id, account.Id)
	assert.Equal(t, "John Doe", account.Name)
	assert.Equal(t, 100, account.Balance)
	assert.Nil(t, err)
}

func TestUpdateAccountUpdatesAccountWithNewData(t *testing.T) {
	repository := createRepositoryWithData(fakeData)

	newBalance := "200"
	updateValue := map[string]interface{}{"Balance": newBalance}

	account, err := repository.UpdateAccount(id, updateValue)

	assert.Equal(t, id, account.Id)
	assert.Equal(t, "John Doe", account.Name)
	assert.Equal(t, 200, account.Balance)
	assert.Nil(t, err)
}

func createRepositoryWithData(fakeData map[string]string) *AccountRepository {
	redis := stubRedis{data: fakeData}
	repository := NewAccountRepository(&redis)
	return repository
}
