package services

type ServiceInterface interface {
	CheckHealthy() (*string, error)
}
