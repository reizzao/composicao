package composicao

// INTERFACES
type IEstrofes interface {
	ICreateFrase_FN(d FraseProps) FraseProps
}

// METHODS
func (t Request) ICreateFrase_FN(d FraseProps) FraseProps {

	request := FraseProps{
		FraseNumero:     d.FraseNumero,
		Gatilhos_Verso:  d.Gatilhos_Verso,
		TempoVerbal_Def: d.TempoVerbal_Def,
		Personagem_Def:  d.Personagem_Def,
		Pergunta:        d.Pergunta,
		Resposta:        d.Resposta,
	}

	res := request

	return res

}
