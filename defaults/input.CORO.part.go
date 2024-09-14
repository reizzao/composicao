package defaults

var Estrofe_CORO_part = Estrofe_CORO_Props{

	Estrofe_Props: Estrofe_Remodel{
		RazaoDaEstrofe:        RAZAO_A2__CONSEQUENCIA_DA_ESTROFE_ANTERIOR_SUAVE,
		ClimaEstrofe:          SUAVE,
		Compasso:              COMPASSADO_LENTO,
		Numero_Notas_PorVerso: UMA_NOTA,
	},

	Versos: VersosCoro_Remodel_2{

		Verso1_Comum: Verso_Remodel{
			Numero:         1,
			Pergunto:       E_FATO_QUE,
			TempoVerbal:    TEMPO_PRESENTE,
			PersonagemTema: PERSON_EU,
		},

		Verso2_MenosComum: Verso_Remodel{
			Numero:         2,
			Pergunto:       OOUTRO_COMOASSIM_EAI,
			TempoVerbal:    TEMPO_PASSADO,
			PersonagemTema: PERSON_AOUTRA,
		},
	}, // subModel

}
