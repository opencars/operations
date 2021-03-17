package mockstore

//go:generate mockgen -destination=./store.go -package=mockstore github.com/opencars/operations/pkg/domain OperationRepository,ResourceRepository
