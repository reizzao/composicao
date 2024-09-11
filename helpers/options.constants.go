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
type EmocaoOptions = string

const (
	// TRISTE
	EMOCAO_TRISTE_LAMENTANDO EmocaoOptions = "EMOCAO_TRISTE_LAMENTANDO"
	EMOCAO_TRISTE_REFLEXAO   EmocaoOptions = "EMOCAO_TRISTE_REFLEXAO"
	EMOCAO_TRISTE_TRISTEZA   EmocaoOptions = "EMOCAO_TRISTE_TRISTEZA"
	EMOCAO_TRISTE_MEDO       EmocaoOptions = "EMOCAO_TRISTE_MEDO"

	// NAO_TRISTE
	EMOCAO_NAOTRISTE_ALEGRE   EmocaoOptions = "EMOCAO_NAOTRISTE_ALEGRE"
	EMOCAO_NAOTRISTE_REFLEXAO EmocaoOptions = "EMOCAO_NAOTRISTE_REFLEXAO"
	EMOCAO_NAOTRISTE_VITORIA  EmocaoOptions = "EMOCAO_NAOTRISTE_VITORIA"
	EMOCAO_NAOTRISTE_PODER    EmocaoOptions = "EMOCAO_NAOTRISTE_PODER"
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

type EstrofeOptions = string

const (
	Estrofe_A               EstrofeOptions = "Estrofe_A"
	Estrofe_Opcional_B      EstrofeOptions = "Estrofe_Opcional_B"
	Estrofe_C               EstrofeOptions = "Estrofe_C"
	Estrofe_Ponte_Refrao    EstrofeOptions = "Estrofe_Ponte_Refrao"
	Estrofe_Refrao          EstrofeOptions = "Estrofe_Refrao"
	Estrofe_RefraoConclusao EstrofeOptions = "Estrofe_RefraoConclusao"
	Estrofe_Opcional_Coro   EstrofeOptions = "Estrofe_Opcional_Coro"
)

// POR FRASE
type Gatilhos_Versos_Options = string

const (
	GATILHOS_1A Gatilhos_Versos_Options = "PelaEmocaoEscolhida, Diria, Fazeria, Apontaria, ObserveiQue, Quando_"

	GATILHOS_2_CONSEQUENCIAS Gatilhos_Versos_Options = "Como_Assim, EAi, "

	GATILHOS_FECHA_SACADAS Gatilhos_Versos_Options = "SACADAS, NA_VIDA?, DITADOS, "
)

type TempoVerbal_Options = string

const (
	TEMPO_PASSADO  TempoVerbal_Options = "TEMPO_PASSADO"
	TEMPO_PRESENTE TempoVerbal_Options = "TEMPO_PRESENTE"
	TEMPO_FUTURO   TempoVerbal_Options = "TEMPO_FUTURO"
)

type Personagem_Options = string

const (
	Person_EU       TempoVerbal_Options = "Person_EU"
	Person_ELA      TempoVerbal_Options = "Person_ELA"
	Person_NOS_DOIS TempoVerbal_Options = "Person_NOS_DOIS"
)
