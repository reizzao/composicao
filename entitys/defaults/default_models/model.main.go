package default_models

type Default struct {
	Guias    GuiasProps
	Estrofes EstrofesDefault
}

type GuiasProps struct {
	EstiloMusicalGravar        string
	EstiloMusicalBatida        string
	EstiloMusicalGravarExemplo string
}

type EstrofesDefault struct {
	Estrofe_A             Estrofe_Props
	Estrofe_B             Estrofe_Props
	Estrofe_C             Estrofe_Props
	Estrofe_Ponte         Estrofe_Props
	Estrofe_Refrao        Estrofe_Props
	Estrofe_Coro_Opcional Estrofe_Props
}

type Estrofe_Props struct {
	RazaoDaEstrofe        RazaoDaEstrofeOptions
	ClimaEstrofe          ClimaEstrofeOptions
	Compasso              CompassoOptions
	Numero_Notas_PorVerso Numero_Notas_PorVerso_Options
	Versos                []Verso
}

type Verso struct {
	Props Verso_Props
}

type Verso_Props struct {
	Numero         int
	RazaoDoVerso   RazaoDoVersoOptions
	TempoVerbal    TempoVerbalOptions
	PersonagemTema PersonagemTemaOptions
}
