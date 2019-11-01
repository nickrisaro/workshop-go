package auto_test

import (
	"testing"

	"github.com/nickrisaro/workshop-go/auto"
)

func TestUnAutoTieneCilindros(t *testing.T) {
	fitito := auto.Auto{Motor: Motor{Cilindros: 4}, Marca: "Fiat"}

	if fitito.Cilindros() != 4 {
		t.Errorf("Esperaba 4 cilindros pero encontré %d", fitito.Cilindros())
		return
	}
}
