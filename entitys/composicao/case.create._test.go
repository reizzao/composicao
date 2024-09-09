package composicao

import (
	"testing"

	ct "github.com/reizzao/music/helpers/constants"
	lib "github.com/reizzao/music/helpers/lib"
)

type ResSut = ComposicaoModel

var (
	sut ResSut = CreateNew(
		InputTestRequest_TardeDemais,
		InputTestComputed_TardeDemais,
	)
	inputTestRequest    ResRequest            = InputTestRequest_TardeDemais
	compare_entrada_voz ct.EntradaPadraoVozOptions = ct.ENTRA_em_1
)

func Test_Entity(t *testing.T) {
	// useVars
	realExpected_entradaVoz := sut.Request.DadosMusica.EntradaPadraoVoz
	tryTest_entradaVoz := compare_entrada_voz

	realExpected_computed := sut.Computed
	tryTest_computed := "foo"

	/* -- Suites -- */

	// SUITE ::  Test [EntradaPadraoVozOptions] - TITULO: deve retornar o valor de { EntradaPadraoVozOptions }
	if realExpected_entradaVoz != tryTest_entradaVoz {
		t.Error(lib.MessageErrorTest(realExpected_entradaVoz, tryTest_entradaVoz))
	}

	// SUITE ::  Test [Computed] - TITULO: deve retornar o valor de { Computed }
	if realExpected_computed != tryTest_computed {
		t.Error(lib.MessageErrorTest(realExpected_computed, tryTest_computed))
	}

}
