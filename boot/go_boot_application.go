package boot

import "fmt"

type Application struct {

}

func NewBootApplication() *Application {
	return &Application{}
}

func (*Application) Run() {
	fmt.Println("test")
}




