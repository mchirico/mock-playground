package idea

import (
	"github.com/golang/mock/gomock"
	Mocks "github.com/mchirico/mock-playground/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestS_Foo(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	MockInterface := Mocks.NewMockIDEA(controller)

	first := MockInterface.EXPECT().Start("stuff", false, MockInterface).Return(false)
	MockInterface.EXPECT().Start(gomock.Any(), false, MockInterface).After(first)
	MockInterface.EXPECT().Running(gomock.Any(), false, MockInterface).Return(false)
	MockInterface.EXPECT().Running(gomock.Any(), false, MockInterface).Return(false)

	result, err := (&S{}).Foo("stuff", false, MockInterface)
	result, err = (&S{}).Foo("stuff", false, MockInterface)
	assert.Contains(t, result["stuff"], "no")
	assert.Nil(t, err)

}

func Test_Foo2(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	MockInterface := Mocks.NewMockIDEA(controller)

	MockInterface.EXPECT().F("a").Return(map[string]string{"a": "b"}, nil)
	MockInterface.EXPECT().Start(gomock.Any(), false, MockInterface).Return(true)

	result, err := (&S{}).Foo("a", false, MockInterface)

	assert.Contains(t, result["a"], "b")
	assert.Nil(t, err)

}

func Test_idea(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	MockInterface := Mocks.NewMockIDEA(controller)

	MockInterface.EXPECT().Start(gomock.Any(), false, MockInterface).Return(false)
	MockInterface.EXPECT().Running(gomock.Any(), false, MockInterface).Return(true)

	MockInterface.EXPECT().IdeaCheck(MockInterface, gomock.Any(), gomock.Any()).DoAndReturn(func(arg0, arg1, arg2 interface{}) {
		arg1.(chan map[string]string) <- map[string]string{"mock": "test"}
		arg2.(chan error) <- nil
	})

	result, err := (&S{}).Foo("a", false, MockInterface)

	assert.Contains(t, result["mock"], "test")
	assert.Nil(t, err)

}
