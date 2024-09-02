package spc_interfone

import (
	"github.com/reizzao/composicao/api/entitys/composicao"
	"github.com/reizzao/musicalidade/api/entitys/campoHarmonico"
)

var Spc_Interfone_RequestInputTester = composicao.RequestComposicao{

	Estrofe:           composicao.Estrofe_A,
	GrauMasterNatural: campoHarmonico.D,
	Frases: []composicao.PerguntaResposta{
		{
			FraseNumero: 1,
			SubFrase_1: composicao.SubFraseProps{
				SubFrase_Tipo: composicao.A_ACONTECEU,
				Silabas:       "A cidade",
			},
			SubFrase_2: composicao.SubFraseProps{
				SubFrase_Tipo: composicao.OQUE_do_ACONTECEU,
				Silabas:       "ja pegou no sono",
			},
		}, // PerguntaResposta - []
		{
			FraseNumero: 2,
			SubFrase_1: composicao.SubFraseProps{
				SubFrase_Tipo: composicao.A_ACONTECEU,
				Silabas:       "Eu já li de novo",
			},
			SubFrase_2: composicao.SubFraseProps{
				SubFrase_Tipo: composicao.OQUE_do_ACONTECEU,
				Silabas:       "as velhas revistas.",
			},
		}, // PerguntaResposta - []
		{
			FraseNumero: 3,
			SubFrase_1: composicao.SubFraseProps{
				SubFrase_Tipo: composicao.A_ACONTECEU,
				Silabas:       "O Jô",
			},
			SubFrase_2: composicao.SubFraseProps{
				SubFrase_Tipo: composicao.OQUE_do_ACONTECEU,
				Silabas:       "já fez a última entrevista.",
			},
		}, // PerguntaResposta - []

	}, // Frases

}
