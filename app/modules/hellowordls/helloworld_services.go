package hellowordls

type HelloWorldServices struct {
	repo *HelloWorldRepositories
}

func NewHelloWorldServices(repo *HelloWorldRepositories) *HelloWorldServices {
	return &HelloWorldServices{repo: repo}
}
