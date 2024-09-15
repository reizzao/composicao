package intester_recomecar

import (
	comp "github.com/reizzao/music/composicao"
	help "github.com/reizzao/music/helpers"
)

// todo

var Reinventar_RequestIN comp.Request = comp.Request{
	DadosMusica: comp.DadosMusicaProps{
		EntradaPadraoVoz: help.ENTRA_em_1, Grau_EscalaNatural_Define_TOM: help.GRAU_3,
	},
	Motivacoes: comp.MotivacoesProps{
		Sentimento_Triste:            false,
		Emocao_Central:               help.EMOCAO_NAOTRISTE_REFLEXAO,
		Quem_eo_Personagem_Principal: help.Person_NOS_DOIS,
		RefraoCabeca_OqueMaisPrecisa_OPersonagemPrincipal: "Recomecar",
		RefraoFecha_Declaracao_O_Jeito_e:                  "EU SEMPRE FUI APAIXONADO POR VOCÊ",
	},
	Estrofe_A1: comp.Estrofe_AS_Props{

		Frase_1: comp.Verso_Props{
			FraseNumero:     1,
			Marca_Rima:      true,
			Gatilhos_Verso:  help.GATILHOS_V1_FATO,
			TempoVerbal_Def: help.TEMPO_PASSADO,
			Personagem_Def:  help.Person_ELA,
			Pergunta: comp.Duas_PalavraPoetica{
				PalavraPoetica1: "DIZ",
				PalavraPoetica2: "QUE",
			},
			Resposta: comp.Duas_PalavraPoetica{
				PalavraPoetica1: "JÁ É",
				PalavraPoetica2: "TARDE",
			},
		},

		Frase_2: comp.Verso_Props{
			FraseNumero:     2,
			Marca_Rima:      true,
			Gatilhos_Verso:  help.GATILHOS_V1_FATO,
			TempoVerbal_Def: help.TEMPO_PRESENTE,
			Personagem_Def:  help.Person_EU,
			Pergunta: comp.Duas_PalavraPoetica{
				PalavraPoetica1: "QUE",
				PalavraPoetica2: "NÃO",
			},
			Resposta: comp.Duas_PalavraPoetica{
				PalavraPoetica1: "TEM",
				PalavraPoetica2: "MAIS JEITO",
			},
		},

		Frase_3: comp.Verso_Props{
			FraseNumero:     3,
			Marca_Rima:      true,
			Gatilhos_Verso:  help.GATILHOS_V1_FATO,
			TempoVerbal_Def: help.TEMPO_PRESENTE,
			Personagem_Def:  help.Person_EU,
			Pergunta: comp.Duas_PalavraPoetica{
				PalavraPoetica1: "MAS EU",
				PalavraPoetica2: "NÃO ACEITO",
			},
			Resposta: comp.Duas_PalavraPoetica{
				PalavraPoetica1: "A DECISÃO",
				PalavraPoetica2: "DO TEU CORAÇÃO EM PARTIR",
			},
		},

		Estrofe: help.Estrofe_A1,
	},
}
