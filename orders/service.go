package main

import "context"

type service struct {
	store OrderStore
}

// Constructor
func NewService(store OrderStore) *service {
	return &service{store: store}
}

func (s *service) CreateOrder(context.Context) error {
	return nil
}
