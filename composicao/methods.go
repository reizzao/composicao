package composicao

// INTERFACES
type IEstrofes interface {
	ICreateFrase_FN(d Verso_Props) Verso_Props
}

// METHODS
func (t Request) ICreateFrase_FN(d Verso_Props) Verso_Props {

	request := Verso_Props{
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
