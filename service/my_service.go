package service

type MyService struct {
	externalService ExternalServiceInterface
}

func NewMyService(externalService ExternalServiceInterface) MyService {
	return MyService{externalService: externalService}
}

func (s MyService) CallExternalService(name string) string {
	result := s.externalService.GetHelloWorld(name)
	return result
}
