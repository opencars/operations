package mocks

//go:generate mockgen -destination=./service.go -package=mocks github.com/opencars/operations/pkg/domain CustomerService
//go:generate mockgen -destination=./store.go -package=mocks github.com/opencars/operations/pkg/domain OperationRepository,ResourceRepository
//go:generate mockgen -destination=./producer.go -package=mocks github.com/opencars/schema Producer
