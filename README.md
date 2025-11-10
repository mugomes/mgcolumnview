# MGColumnView

MGColumnView é um componente customizado para Fyne que exibe dados em formato de tabela, com suporte a cabeçalhos clicáveis, selação por checkboxes, adição e remoção dinâmica de linhas e ordenação por coluna.

## Recursos

- NewColumnView: Cria um novo columnview com cabeçalhos
- AddRow: Adiciona item
- ListSelected: Retorna os itens que foram selecionados
- ListAll: Retorna todos os itens
- RemoveSelected: Remove os itens selecionados
- RemoveAll: Remove todos os itens
- Ordem Crescente: Em todas as colunas, clicando no cabeçalho
- Selecionar Todos: Checkbox no cabeçalho

## Instalação

`go get github.com/mugomes/mgcolumnview`

## Exemplo

```
import "github.com/mugomes/mgcolumnview"

headers := []string{"Nome","Idade"}
cv := columnview.NewColumnView(headers)

cv.AddRow([]string{
    "Maria",
    "39",
})

cv.AddRow([]string{
    "João",
    "49",
})

for _, row := range cv.ListAll() {
	fmt.Println(row)
}
```

## Information

 - [Page MGColumnView](https://github.com/mugomes/mgcolumnview)

## Requirement

 - Go 1.24.6
 - Fyne 2.7.0

## Support

- GitHub: https://github.com/sponsors/mugomes
- More: https://mugomes.github.io/apoie.html

## License

Copyright (c) 2025 Murilo Gomes Julio

Licensed under the [MIT](https://github.com/mugomes/mgcolumnview/blob/main/LICENSE) license.

All contributions to the MGColumnView are subject to this license.