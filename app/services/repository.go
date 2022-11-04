package services

type RepositoryInterface interface {
	CheckHealthy() (*string, error)
}
