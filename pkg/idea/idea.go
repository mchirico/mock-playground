package idea

import (
	"github.com/mchirico/mock-playground/pkg/idea/imp"
)

type S struct {
	M map[string]interface{}
	f func(string) (map[string]string, error)
	v bool
}

func (s *S) Foo(input string, b bool, ideaInterface imp.IDEA) (map[string]string, error) {

	if ideaInterface.Start(input, b, ideaInterface) {
		return ideaInterface.F(input)
	}

	if ideaInterface.Running(input, b, ideaInterface) {
		var ideaMap = make(chan map[string]string)
		var chanErr = make(chan error)
		go ideaInterface.IdeaCheck(ideaInterface, ideaMap, chanErr)
		return <-ideaMap, <-chanErr

	}

	m := map[string]string{}
	m[input] = "no"

	return m, nil
}
