package mock

import "fmt"

type Person interface {
	Run()
	Login(id int) (bool, error)
}

func PersonCheck(p Person) error {
	id := 10
	if ok, err := p.Login(id); !ok {
		return fmt.Errorf("this is not %v\n", id)
	} else if err != nil {
		return err
	}
	p.Run()
	return nil
}
