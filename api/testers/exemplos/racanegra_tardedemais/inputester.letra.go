package racanegra_tardedemais

import (
	"github.com/reizzao/composicao/api/entitys/composicao"
	"github.com/reizzao/musicalidade/api/entitys/campoHarmonico"
)

var Racanegra_tardedemais_RequestInputTester = composicao.RequestComposicao{

	Estrofe:           composicao.Estrofe_A,
	GrauMasterNatural: campoHarmonico.D,
	Frases: []composicao.PerguntaResposta{
		{
			FraseNumero: 1,
			SubFrase_1: composicao.SubFraseProps{
				SubFrase_Tipo: composicao.A_ACONTECEU,
				Silabas:       "Olha",
			},
			SubFrase_2: composicao.SubFraseProps{
				SubFrase_Tipo: composicao.OQUE_do_ACONTECEU,
				Silabas:       "só voce",
			},
		}, // PerguntaResposta - []
		{
			FraseNumero: 2,
			SubFrase_1: composicao.SubFraseProps{
				SubFrase_Tipo: composicao.E_AI,
				Silabas:       "depois de",
			},
			SubFrase_2: composicao.SubFraseProps{
				SubFrase_Tipo: composicao.E_AI_RESPOSTA,
				Silabas:       "me perder.",
			},
		}, // PerguntaResposta - []
		{
			FraseNumero: 3,
			SubFrase_1: composicao.SubFraseProps{
				SubFrase_Tipo: composicao.CONTINUA_ACONTECEU,
				Silabas:       "Veja",
			},
			SubFrase_2: composicao.SubFraseProps{
				SubFrase_Tipo: composicao.OQUE_do_CONTINUA_ACONTECEU,
				Silabas:       "só você",
			},
		}, // PerguntaResposta - []
		{
			FraseNumero: 4,
			SubFrase_1: composicao.SubFraseProps{
				SubFrase_Tipo: composicao.RESULTADOFECHA,
				Silabas:       "Que",
			},
			SubFrase_2: composicao.SubFraseProps{
				SubFrase_Tipo: composicao.RESULTADOFECHA_OQUE,
				Silabas:       "pena",
			},
		}, // PerguntaResposta - []

	}, // Frases

}
