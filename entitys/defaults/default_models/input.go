package default_models

var Default_Input = Default{

	Guias: GuiasProps{
		EstiloMusicalGravar:        "NEUTRO",
		EstiloMusicalBatida:        "POP",
		EstiloMusicalGravarExemplo: "https://youtu.be/Os42dqBG-Yk?si=XNRg5fxvoIs8vgah",
	},

	Estrofes: EstrofesDefault{

		Estrofe_A: Estrofe_Props{
			RazaoDaEstrofe:        CONSEQUENCIA_DA_EMOCAO,
			ClimaEstrofe:          SUAVE,
			Compasso:              LENTO,
			Numero_Notas_PorVerso: UMA_NOTA,
			Versos: []Verso{
				{Props: Verso_Props{
					Numero: 1, RazaoDoVerso: UM_FATO,
					TempoVerbal: TEMPO_PRESENTE, PersonagemTema: PERSON_EU}},
				{Props: Verso_Props{
					Numero: 2, RazaoDoVerso: COMO_ASSIM_EAI,
					TempoVerbal: TEMPO_PASSADO, PersonagemTema: PERSON_AOUTRA}},
				{Props: Verso_Props{
					Numero: 3, RazaoDoVerso: NAVIDA_SACADA_POESIA,
					TempoVerbal: TEMPO_FUTURO, PersonagemTema: PERSON_TODOSNOS}},
			}}, // Estrofe_

	},
}
