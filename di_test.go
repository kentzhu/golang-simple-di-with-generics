package simpledi

import (
	"reflect"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

func Test_Inject(t *testing.T) {
	c := NewContainer()

	Put[iDemoAService](c, &demoAService{})
	PutWithName[iDemoBService](c, &demoBService{}, "HB")

	ctl := demoController{}
	Inject(c, &ctl)
	assert.Equal(t, ctl.GetA(), "A")
	assert.Equal(t, ctl.GetB(), "B")
}

func Test_Put_And_Set(t *testing.T) {
	c := NewContainer()

	assert.Panics(t, func() {
		Get[iDemoAService](c)
	})

	Put[iDemoAService](c, &demoAService{})
	assert.NotNil(t, Get[iDemoAService](c))
	assert.Equal(t, Get[iDemoAService](c).demoA(), "A")

	assert.Panics(t, func() {
		Get[iDemoBService](c)
	})

	Put[iDemoBService](c, &demoBService{})
	assert.NotNil(t, Get[iDemoBService](c))
	assert.Equal(t, Get[iDemoBService](c).demoB(), "B")
}

func Test_GetWithName_And_SetWithName(t *testing.T) {
	c := NewContainer()

	PutWithName[iDemoAService](c, &demoAService{}, "HA")
	assert.Panics(t, func() {
		Get[iDemoAService](c)
	})
	assert.NotNil(t, GetWithName[iDemoAService](c, "HA"))
	assert.Equal(t, GetWithName[iDemoAService](c, "HA").(iDemoAService).demoA(), "A")
}

func Test_GetInterfaceKey(t *testing.T) {
	correctKey := "testing/quick.Generator"
	key := getInterfaceKey(reflect.TypeOf((*quick.Generator)(nil)).Elem())
	assert.Equal(t, correctKey, key)
}

// ---

type demoController struct {
	ASvc iDemoAService `inject:""`
	BSvc iDemoBService `inject:"HB"`
}

func (m demoController) GetA() string {
	return m.ASvc.demoA()
}

func (m demoController) GetB() string {
	return m.BSvc.demoB()
}

// ---

type iDemoAService interface {
	demoA() string
}

// ---

type demoAService struct {
}

func (d *demoAService) demoA() string {
	return "A"
}

// ---

type iDemoBService interface {
	demoB() string
}

// ---

type demoBService struct {
}

func (d *demoBService) demoB() string {
	return "B"
}
