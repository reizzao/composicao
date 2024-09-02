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
			FraseNumero:     1,
			MetaFrase:       composicao.META_A_INFORMACAO_DO_QUEQUERODIZER,
			TempoVerbal_Def: composicao.TEMPO_PASSADO,
			Personagem_Def:  composicao.Person_EU,
			FormaFrase_1: composicao.FormaFraseProps{
				FormaFrase_Tipo: composicao.A_ACONTECEU,
				Silabas:         "Olha",
			},
			FormaFrase_2: composicao.FormaFraseProps{
				FormaFrase_Tipo: composicao.OQUE_do_ACONTECEU,
				Silabas:         "só voce",
			},
		}, // PerguntaResposta - []
		{
			FraseNumero:     2,
			MetaFrase:       composicao.META_A_INFORMACAO_DO_QUEQUERODIZER,
			TempoVerbal_Def: composicao.TEMPO_PASSADO,
			Personagem_Def:  composicao.Person_EU,
			FormaFrase_1: composicao.FormaFraseProps{
				FormaFrase_Tipo: composicao.E_AI,
				Silabas:         "depois de",
			},
			FormaFrase_2: composicao.FormaFraseProps{
				FormaFrase_Tipo: composicao.E_AI_RESPOSTA,
				Silabas:         "me perder.",
			},
		}, // PerguntaResposta - []
		{
			FraseNumero:     3,
			MetaFrase:       composicao.META_A_INFORMACAO_DO_QUEQUERODIZER,
			TempoVerbal_Def: composicao.TEMPO_PRESENTE,
			Personagem_Def:  composicao.Person_ELA,
			FormaFrase_1: composicao.FormaFraseProps{
				FormaFrase_Tipo: composicao.CONTINUA_ACONTECEU,
				Silabas:         "Veja",
			},
			FormaFrase_2: composicao.FormaFraseProps{
				FormaFrase_Tipo: composicao.OQUE_do_CONTINUA_ACONTECEU,
				Silabas:         "só você",
			},
		}, // PerguntaResposta - []
		{
			FraseNumero:     4,
			MetaFrase:       composicao.META_A_INFORMACAO_DO_QUEQUERODIZER,
			TempoVerbal_Def: composicao.TEMPO_FUTURO,
			Personagem_Def:  composicao.Person_NOS_DOIS,
			FormaFrase_1: composicao.FormaFraseProps{
				FormaFrase_Tipo: composicao.RESULTADOFECHA,
				Silabas:         "Que",
			},
			FormaFrase_2: composicao.FormaFraseProps{
				FormaFrase_Tipo: composicao.RESULTADOFECHA_OQUE,
				Silabas:         "pena",
			},
		}, // PerguntaResposta - []

	}, // Frases

}
