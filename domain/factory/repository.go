package factory

import "github.com/wenealves10/microservice-creditcard-processing-golang/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
