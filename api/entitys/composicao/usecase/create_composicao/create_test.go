package create_composicao

import (
	"testing"

	mdc "github.com/reizzao/composicao/api/entitys/composicao/modelcomposicao"
	"github.com/reizzao/composicao/api/helpers/constants/lib"
	itt "github.com/reizzao/composicao/api/testers/itt_exemplo_racanegra_tardedemais"
)

type ResSut = mdc.ComposicaoModel

var sut ResSut = CreateNew(itt.InputTester_ExemploRacanegra_TardeDemais_ByUseCaseCreate)

func Test_Entity(t *testing.T) {
	expect_request := sut
	compare_request := ResSut{
		Request:  "tardeDemais",
		Computed: "",
	}

	/* -- Suites -- */

	// SUITE :: TARGET: Test Request - TITULO: deve retornar tardeDemais no campo Request
	if sut != compare_request {
		t.Error(lib.MessageErrorTest(), expect_request, compare_request)
	}

}
