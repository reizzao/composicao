package conceito_composicao

type Conceito_Composicao struct {
	Padroes  PadroesProps
	Guias    GuiasProps
	Estrofes EstrofesDefault
}

type PadroesProps struct {
	Quant_PalavrasPoeticas_por_Pergunta Quant_PalavrasPoeticas_por_Pergunta_Options
	Quant_PalavrasPoeticas_por_Resposta Quant_PalavrasPoeticas_por_Pergunta_Options
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
	Estrofe_REFRAO        Estrofe_REFRAO_Props
	Estrofe_CORO_Opcional Estrofe_CORO_Props
}

// ReModel
type Estrofe_Remodel struct {
	RazaoDaEstrofe        RazaoDaEstrofeOptions
	ClimaEstrofe          ClimaEstrofeOptions
	Compasso              CompassoOptions
	Numero_Notas_PorVerso Numero_Notas_PorVerso_Options
}

type Verso_Remodel struct {
	Numero         int
	Pergunto       Pergunto_byVerso_Options
	TempoVerbal    TempoVerbalOptions
	PersonagemTema PersonagemTemaOptions
}

type Versos_Remodel_Suave4 struct {
	Verso1_Comum                  Verso_Remodel
	Verso2_ComumPlus              Verso_Remodel
	Verso3_MenosComum_ApontaFecha Verso_Remodel
	Verso4_MenosComumPlus         Verso_Remodel
}

type Versos_Remodel_Tenso4 struct {
	Verso1_Comum          Verso_Remodel
	Verso2_MenosComum     Verso_Remodel
	Verso3_ComumPlus      Verso_Remodel
	Verso4_MenosComumPlus Verso_Remodel
}

type VersosRefrao_Remodel_4 struct {
	Verso_1_Clima_GritoouSuave_Objetivo        Verso_Remodel
	Verso_2_ClimaIgualAnterior_EAI             Verso_Remodel
	Verso_3_ClimaOposto_Explica_Fecha          Verso_Remodel
	Verso_4_ClimaOposto_Explica_Fecha_Opcional Verso_Remodel
}

type VersosCoro_Remodel_2 struct {
	Verso1_Comum      Verso_Remodel
	Verso2_MenosComum Verso_Remodel
}

// Estrofes
type Estrofe_A1_Props struct {
	Estrofe_Props Estrofe_Remodel
	Versos        VersosEstrofe_A1_Props
}
type Estrofe_A2_Opcional_Props struct {
	Estrofe_Props Estrofe_Remodel
	Versos        VersosEstrofe_A2_Opcional_Props
}

type Estrofe_B_Props struct {
	Estrofe_Props Estrofe_Remodel
	Versos        VersosEstrofe_B_Props
}

type Estrofe_REFRAO_Props struct {
	Estrofe_Props Estrofe_Remodel
	Versos        VersosEstrofe_REFRAO_Props
}
type Estrofe_CORO_Props struct {
	Estrofe_Props Estrofe_Remodel
	Versos        VersosEstrofe_CORO_Props
}

// Versos
type VersosEstrofe_A1_Props = Versos_Remodel_Suave4

type VersosEstrofe_A2_Opcional_Props = Versos_Remodel_Suave4

type VersosEstrofe_B_Props = Versos_Remodel_Tenso4

type VersosEstrofe_REFRAO_Props = VersosRefrao_Remodel_4

type VersosEstrofe_CORO_Props = VersosCoro_Remodel_2
