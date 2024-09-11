package helpers

type EntradaPadraoVozOptions = string

const (
	ENTRA_em_1  EntradaPadraoVozOptions = "ENTRA_em_1"
	ENTRA_em_1e EntradaPadraoVozOptions = "ENTRA_em_1e"
	ENTRA_em_2  EntradaPadraoVozOptions = "ENTRA_em_2"
	ENTRA_em_2e EntradaPadraoVozOptions = "ENTRA_em_2e"
	ENTRA_em_3  EntradaPadraoVozOptions = "ENTRA_em_3"
	ENTRA_em_3e EntradaPadraoVozOptions = "ENTRA_em_3e"
	ENTRA_em_4  EntradaPadraoVozOptions = "ENTRA_em_4"
	ENTRA_em_4e EntradaPadraoVozOptions = "ENTRA_em_4e"
)

// -- Emocao
type Emocao_Central_Options = string

const (
	// TRISTE
	EMOCAO_TRISTE_LAMENTANDO Emocao_Central_Options = "EMOCAO_TRISTE_LAMENTANDO"
	EMOCAO_TRISTE_REFLEXAO   Emocao_Central_Options = "EMOCAO_TRISTE_REFLEXAO"
	EMOCAO_TRISTE_TRISTEZA   Emocao_Central_Options = "EMOCAO_TRISTE_TRISTEZA"
	EMOCAO_TRISTE_MEDO       Emocao_Central_Options = "EMOCAO_TRISTE_MEDO"

	// NAO_TRISTE
	EMOCAO_NAOTRISTE_ALEGRE   Emocao_Central_Options = "EMOCAO_NAOTRISTE_ALEGRE"
	EMOCAO_NAOTRISTE_REFLEXAO Emocao_Central_Options = "EMOCAO_NAOTRISTE_REFLEXAO"
	EMOCAO_NAOTRISTE_VITORIA  Emocao_Central_Options = "EMOCAO_NAOTRISTE_VITORIA"
	EMOCAO_NAOTRISTE_PODER    Emocao_Central_Options = "EMOCAO_NAOTRISTE_PODER"
)

type Grau_EscalaNatural_Options = string

const (
	C Grau_EscalaNatural_Options = "C"
	D Grau_EscalaNatural_Options = "D"
	E Grau_EscalaNatural_Options = "E"
	F Grau_EscalaNatural_Options = "F"
	G Grau_EscalaNatural_Options = "G"
	A Grau_EscalaNatural_Options = "A"
	B Grau_EscalaNatural_Options = "B"
)

type Estrofes_Options = string

const (
	Estrofe_A1              Estrofes_Options = "Estrofe_A1"
	Estrofe_A2_Opcional     Estrofes_Options = "Estrofe_A2_Opcional"
	Estrofe_B               Estrofes_Options = "Estrofe_B"
	Estrofe_Ponte_Refrao    Estrofes_Options = "Estrofe_Ponte_Refrao"
	Estrofe_Refrao          Estrofes_Options = "Estrofe_Refrao"
	Estrofe_RefraoConclusao Estrofes_Options = "Estrofe_RefraoConclusao"
	Estrofe_Opcional_Coro   Estrofes_Options = "Estrofe_Opcional_Coro"
)

// POR FRASE
type Gatilhos_Versos_Options = string

const (
	GATILHOS_A1 Gatilhos_Versos_Options = "Sacadas_FrasesDeComando que levam a EmocaoCentralEscolhida ex:, Diria, Fazeria, Apontaria, ObserveiQue, Quando_"

	GATILHOS_2_CONSEQUENCIAS Gatilhos_Versos_Options = "Como_Assim, EAi, "

	GATILHOS_FECHA_SACADAS Gatilhos_Versos_Options = "SACADAS, NA_VIDA?, DITADOS, "
)

type TempoVerbal_Options = string

const (
	TEMPO_PASSADO  TempoVerbal_Options = "TEMPO_PASSADO"
	TEMPO_PRESENTE TempoVerbal_Options = "TEMPO_PRESENTE"
	TEMPO_FUTURO   TempoVerbal_Options = "TEMPO_FUTURO"
)

type Personagens_Options = string

const (
	Person_EU       TempoVerbal_Options = "Person_EU"
	Person_ELA      TempoVerbal_Options = "Person_ELA"
	Person_NOS_DOIS TempoVerbal_Options = "Person_NOS_DOIS"
)
