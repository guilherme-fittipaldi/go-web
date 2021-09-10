package models

type Produto struct {
	Nome, Descricao string
	Preco float64
	Quantidade, Id int
}

func BuscaTodosOsProdutos() []Produto {
	db := conectComBandoDeDados()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p:= Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next(){
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	return produtos
	defer db.Close()
}