package composicao

// Request
func createRequestPrepare(d Request) Request {
	request := Request{
		DadosMusica: DadosMusicaProps{
			Tom_via_Grau_EscalaNatural: d.DadosMusica.Tom_via_Grau_EscalaNatural,
			EntradaPadraoVoz:           d.DadosMusica.EntradaPadraoVoz,
		},
		Motivacoes: MotivacoesProps{
			Sentimento_Triste:           d.Motivacoes.Sentimento_Triste,
			Emocao:                      d.Motivacoes.Emocao,
			CabecaoRefrao_O_Que_Resolve: d.Motivacoes.CabecaoRefrao_O_Que_Resolve,
			FraseFinal_O_Jeito_e:        d.Motivacoes.FraseFinal_O_Jeito_e,
		},

		Estrofe_A: d.ICreateFrase_FN(
			FraseProps{
				Estrofe:         d.Estrofe_A.Estrofe,
				FraseNumero:     d.Estrofe_A.FraseNumero,
				Gatilhos_Verso:  d.Estrofe_A.Gatilhos_Verso,
				TempoVerbal_Def: d.Estrofe_A.TempoVerbal_Def,
				Personagem_Def:  d.Estrofe_A.Personagem_Def,
				Pergunta:        d.Estrofe_A.Pergunta,
				Resposta:        d.Estrofe_A.Resposta,
			},
		),

		Estrofe_B: d.ICreateFrase_FN(
			FraseProps{
				Estrofe:         d.Estrofe_B.Estrofe,
				FraseNumero:     d.Estrofe_B.FraseNumero,
				Gatilhos_Verso:  d.Estrofe_B.Gatilhos_Verso,
				TempoVerbal_Def: d.Estrofe_B.TempoVerbal_Def,
				Personagem_Def:  d.Estrofe_B.Personagem_Def,
				Pergunta:        d.Estrofe_B.Pergunta,
				Resposta:        d.Estrofe_B.Resposta,
			},
		),
	}
	res := request

	return res
}

// Computed

func createComputedPrepare(c Computed) Computed {
	return c
}
