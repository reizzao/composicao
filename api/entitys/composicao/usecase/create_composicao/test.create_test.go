package create_composicao

import (
	"testing"

	itt "github.com/reizzao/composicao/api/entitys/composicao/literals/inputtester/itt_exemplo_racanegra_tardedemais"
	mdc "github.com/reizzao/composicao/api/entitys/composicao/modelcomposicao"
	lib "github.com/reizzao/composicao/api/lib"
)

type ResSut = mdc.ComposicaoModel

var sut ResSut = CreateNew(itt.InputTestRequest_TardeDemais, itt.InputTestComputed_TardeDemais)
var inputTestRequest mdc.ResRequest = itt.InputTestRequest_TardeDemais
var inputTestComputed mdc.ResRequest = itt.InputTestComputed_TardeDemais

func Test_Entity(t *testing.T) {
	expect_request := sut
	compare_request := ResSut{
		Request: inputTestRequest,
		// Computed: inputTestComputed,
		Computed: "FOO",
	}

	/* -- Suites -- */

	// SUITE :: TARGET: Test Request - TITULO: deve retornar um novo caso de uso {create}
	if sut != compare_request {
		t.Error(lib.MessageErrorTest(expect_request, compare_request))
	}

}
