package composicao

// Request
func createRequestPrepare(d ResRequest) ResRequest {
	request := ResRequest{
		DadosMusica: DadosMusicaProps{
			EntradaPadraoVoz: d.DadosMusica.EntradaPadraoVoz,
		},
		Emocao:            d.Emocao,
		GrauMasterNatural: d.GrauMasterNatural,

		Estrofe_A: d.ICreateFrase_FN(
			FraseProps{
				Estrofe:         d.Estrofe_A.Estrofe,
				FraseNumero:     d.Estrofe_A.FraseNumero,
				MetaFrase:       d.Estrofe_A.MetaFrase,
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
				MetaFrase:       d.Estrofe_B.MetaFrase,
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

func createComputedPrepare(c ResComputed) ResComputed {
	return c
}
