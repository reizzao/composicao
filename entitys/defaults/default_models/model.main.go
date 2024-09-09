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
	Estrofe_A             Estrofe_A_Props
	Estrofe_B_Opcional    Estrofe_B_Props
	Estrofe_C             Estrofe_C_Props
	Estrofe_PONTE         Estrofe_PONTE_Props
	Estrofe_REFRAO        Estrofe_REFRAO_Props
	Estrofe_CORO_Opcional Estrofe_CORO_Props
}

type Estrofe_Props struct {
	RazaoDaEstrofe        RazaoDaEstrofeOptions
	ClimaEstrofe          ClimaEstrofeOptions
	Compasso              CompassoOptions
	Numero_Notas_PorVerso Numero_Notas_PorVerso_Options
	// Versos                []Verso
}

type Verso_Props struct {
	Numero         int
	RazaoDoVerso   RazaoDoVersoOptions
	TempoVerbal    TempoVerbalOptions
	PersonagemTema PersonagemTemaOptions
}

type Estrofe_A_Props struct {
	Estrofe_Props Estrofe_Props
	Versos        VersosEstrofe_A_Props
}
type Estrofe_B_Props struct {
	Estrofe_Props Estrofe_Props
	Versos        VersosEstrofe_B_Props
}
type Estrofe_C_Props struct {
	Estrofe_Props Estrofe_Props
	Versos        VersosEstrofe_C_Props
}
type Estrofe_PONTE_Props struct {
	Estrofe_Props Estrofe_Props
	Versos        VersosEstrofe_PONTE_Props
}
type Estrofe_REFRAO_Props struct {
	Estrofe_Props Estrofe_Props
	Versos        VersosEstrofe_REFRAO_Props
}
type Estrofe_CORO_Props struct {
	Estrofe_Props Estrofe_Props
	Versos        VersosEstrofe_CORO_Props
}

type VersosEstrofe_A_Props struct {
	Verso_1 Verso_Props
	Verso_2 Verso_Props
	Verso_3 Verso_Props
}

type VersosEstrofe_B_Props struct {
	Verso_1 Verso_Props
	Verso_2 Verso_Props
	Verso_3 Verso_Props
}
type VersosEstrofe_C_Props struct {
	Verso_1 Verso_Props
	Verso_2 Verso_Props
}
type VersosEstrofe_PONTE_Props struct {
	Verso_1 Verso_Props
	Verso_2 Verso_Props
}
type VersosEstrofe_REFRAO_Props struct {
	Verso_1 Verso_Props
	Verso_2 Verso_Props
	Verso_3 Verso_Props
	Verso_4 Verso_Props
}
type VersosEstrofe_CORO_Props struct {
	Verso_1 Verso_Props
	Verso_2 Verso_Props
}
