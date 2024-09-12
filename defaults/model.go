package defaults

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
	Estrofe_A1            Estrofe_A1_Props
	Estrofe_A2_Opcional   Estrofe_A2_Opcional_Props
	Estrofe_B             Estrofe_B_Props
	Estrofe_PONTE         Estrofe_PONTE_Props
	Estrofe_REFRAO        Estrofe_REFRAO_Props
	Estrofe_CORO_Opcional Estrofe_CORO_Props
}

type Estrofe_Props struct {
	RazaoDaEstrofe        RazaoDaEstrofeOptions
	ClimaEstrofe          ClimaEstrofeOptions
	Compasso              CompassoOptions
	Numero_Notas_PorVerso Numero_Notas_PorVerso_Options
}

type Verso_Props struct {
	Numero            int
	PerguntaQueMeFaco PerguntaQueMeFaco_byVerso_Options
	TempoVerbal       TempoVerbalOptions
	PersonagemTema    PersonagemTemaOptions
}

type Estrofe_A1_Props struct {
	Estrofe_Props Estrofe_Props
	Versos        VersosEstrofe_A1_Props
}
type Estrofe_A2_Opcional_Props struct {
	Estrofe_Props Estrofe_Props
	Versos        VersosEstrofe_A2_Opcional_Props
}
type Estrofe_B_Props struct {
	Estrofe_Props Estrofe_Props
	Versos        VersosEstrofe_B_Props
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

// Versos by Estrofe
type VersosEstrofe_A1_Props struct {
	Verso_1 Verso_Props
	Verso_2 Verso_Props
	Verso_3 Verso_Props
}

// Props Versos by Estrofe
type VersosEstrofe_A2_Opcional_Props struct {
	Verso_1 Verso_Props
	Verso_2 Verso_Props
	Verso_3 Verso_Props
}
type VersosEstrofe_B_Props struct {
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
