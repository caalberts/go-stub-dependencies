package main

import (
	"strconv"
	"errors"
)

type AccountRepository struct {
	proxy RedisProxy
}

func NewAccountRepository(client RedisProxy) *AccountRepository {
	return &AccountRepository{client}
}

func (r *AccountRepository) FetchAccount(id string) (*Account, error) {
	data, err := r.proxy.HGetAll(id)
	if err != nil {
		return nil, err
	}

	return toAccount(data)
}

func (r *AccountRepository) UpdateAccount(id string, data map[string]interface{}) (*Account, error) {
	_, err := r.proxy.HMSet(id, data)
	if err != nil {
		return nil, err
	} else {
		return r.FetchAccount(id)
	}
}

type Account struct {
	Id      string
	Name    string
	Balance int
}

func toAccount(data map[string]string) (*Account, error) {
	if _, ok := data["Id"]; !ok {
		return nil, errors.New("Missing account ID")
	}

	balance, err := strconv.Atoi(data["Balance"])
	if err != nil {
		return nil, err
	}

	return &Account{
		Id:      data["Id"],
		Name:    data["Name"],
		Balance: balance,
	}, nil
}
