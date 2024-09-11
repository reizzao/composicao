package composicao

import (
	"testing"
	help "github.com/reizzao/music/helpers"
)

type ResSut = ComposicaoModel

var (
	sut ResSut = CreateNew(
		InputTestRequest_TardeDemais,
		InputTestComputed_TardeDemais,
	)
	inputTestRequest    ResRequest                   = InputTestRequest_TardeDemais
	compare_entrada_voz help.EntradaPadraoVozOptions = help.ENTRA_em_1
)

func Test_Entity(t *testing.T) {
	// useVars
	realExpehelled_entradaVoz := sut.Request.DadosMusica.EntradaPadraoVoz
	tryTest_entradaVoz := compare_entrada_voz

	realExpehelled_computed := sut.Computed
	tryTest_computed := "foo"

	/* -- Suites -- */

	// SUITE ::  Test [EntradaPadraoVozOptions] - TITULO: deve retornar o valor de { EntradaPadraoVozOptions }
	if realExpehelled_entradaVoz != tryTest_entradaVoz {
		t.Error(help.MessageErrorTest(realExpehelled_entradaVoz, tryTest_entradaVoz))
	}

	// SUITE ::  Test [Computed] - TITULO: deve retornar o valor de { Computed }
	if realExpehelled_computed != tryTest_computed {
		t.Error(help.MessageErrorTest(realExpehelled_computed, tryTest_computed))
	}

}
