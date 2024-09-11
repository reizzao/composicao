package composicao

// Request
func createRequestPrepare(d Request) Request {
	request := Request{
		DadosMusica: DadosMusicaProps{
			Grau_EscalaNatural_Define_TOM: d.DadosMusica.Grau_EscalaNatural_Define_TOM,
			EntradaPadraoVoz:              d.DadosMusica.EntradaPadraoVoz,
		},
		Motivacoes: MotivacoesProps{
			Sentimento_Triste: d.Motivacoes.Sentimento_Triste,
			Emocao_Central:    d.Motivacoes.Emocao_Central,
			RefraoCabeca_OqueMaisPrecisa_OPersonagemPrincipal: d.Motivacoes.RefraoCabeca_OqueMaisPrecisa_OPersonagemPrincipal,
			RefraoFecha_O_Jeito_e:                             d.Motivacoes.RefraoFecha_O_Jeito_e,
		},

		Estrofe_A1: d.ICreateFrase_FN(
			Verso_Props{
				Estrofe:         d.Estrofe_A1.Estrofe,
				FraseNumero:     d.Estrofe_A1.FraseNumero,
				Gatilhos_Verso:  d.Estrofe_A1.Gatilhos_Verso,
				TempoVerbal_Def: d.Estrofe_A1.TempoVerbal_Def,
				Personagem_Def:  d.Estrofe_A1.Personagem_Def,
				Pergunta:        d.Estrofe_A1.Pergunta,
				Resposta:        d.Estrofe_A1.Resposta,
			},
		),

		Estrofe_A2_Opcional: d.ICreateFrase_FN(
			Verso_Props{
				Estrofe:         d.Estrofe_A2_Opcional.Estrofe,
				FraseNumero:     d.Estrofe_A2_Opcional.FraseNumero,
				Gatilhos_Verso:  d.Estrofe_A2_Opcional.Gatilhos_Verso,
				TempoVerbal_Def: d.Estrofe_A2_Opcional.TempoVerbal_Def,
				Personagem_Def:  d.Estrofe_A2_Opcional.Personagem_Def,
				Pergunta:        d.Estrofe_A2_Opcional.Pergunta,
				Resposta:        d.Estrofe_A2_Opcional.Resposta,
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
