package composicao

import (
	"testing"

	help "github.com/reizzao/music/helpers"
)

type ResSut = ComposicaoModel

var (
	sut ResSut = CreateNew(
		TardeDemais_RequestIN,
		TardeDemais_ComputedIN,
	)
	inputTestRequest    Request                      = TardeDemais_RequestIN
	compare_entrada_voz help.EntradaPadraoVozOptions = help.ENTRA_em_1
)

func Test_Entity(t *testing.T) {
	// useVars
	sut_entradaVoz := sut.Request.DadosMusica.EntradaPadraoVoz
	try_entradaVoz := compare_entrada_voz

	sut_computed := sut.Computed
	try_computed := "foo"

	/* -- Suites -- */

	// SUITE ::  Test [EntradaPadraoVozOptions] - TITULO: deve retornar o valor de { EntradaPadraoVozOptions }
	if sut_entradaVoz != try_entradaVoz {
		t.Error(help.MessageErrorTest(sut_entradaVoz, try_entradaVoz))
	}

	// SUITE ::  Test [Computed] - TITULO: deve retornar o valor de { Computed }
	if sut_computed != try_computed {
		t.Error(help.MessageErrorTest(sut_computed, try_computed))
	}

}
