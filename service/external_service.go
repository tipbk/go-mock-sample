package service

type ExternalServiceInterface interface {
	GetHelloWorld(name string) string
}
type ExternalService struct{}

func (e ExternalService) GetHelloWorld(name string) string {
	return "Hello " + name
}
