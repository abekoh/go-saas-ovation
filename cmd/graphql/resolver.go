package main

import (
	"github.com/abekoh/go-saas-ovation/usecase/command"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	backlogsUsecase command.Commands
}
