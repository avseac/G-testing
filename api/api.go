package api

import (
	t "GinGonic-Tmpl/types"
)

type Group []struct {
	Name     string
	Products []t.Product
}

var DataProductGroupList Group = Group{
	{
		Name: "Primo",
		Products: []t.Product{
			{Id: 1, Name: "Pasta al pomodoro", Price: 8.73, Category: "Primo"},
			{Id: 2, Name: "Strudel salato ", Price: 3.74, Category: "Primo"},
			{Id: 3, Name: "Canederli alle rape rosse", Price: 7.50, Category: "Primo"},
			{Id: 4, Name: "Pasta alla calabrese", Price: 6.78, Category: "Primo"},
			{Id: 5, Name: "Pasta all'olio", Price: 3.91, Category: "Primo"},
			{Id: 6, Name: "Crespella al radicchio", Price: 4.12, Category: "Primo"},
		},
	},
	{
		Name: "Secondo",
		Products: []t.Product{
			{Id: 7, Name: "Scaloppa di tacchino alla bolognese", Price: 2.56, Category: "Secondo"},
			{Id: 8, Name: "Hamburger di patate e formaggio", Price: 8.81, Category: "Secondo"},
			{Id: 9, Name: "Moscardini in umido", Price: 3.53, Category: "Secondo"},
			{Id: 10, Name: "Filetto di sardina croccante all' arancia", Price: 8.98, Category: "Secondo"},
		},
	},
	{
		Name: "Contorno",
		Products: []t.Product{
			{Id: 11, Name: "Insalata", Price: 5.31, Category: "Contorno"},
			{Id: 12, Name: "Purè di patate", Price: 0.72, Category: "Contorno"},
			{Id: 13, Name: "Zucca al forno", Price: 2.76, Category: "Contorno"},
		},
	},
	{
		Name: "Frutta",
		Products: []t.Product{
			{Id: 14, Name: "Macedonia", Price: 2.62, Category: "Frutta"},
			{Id: 15, Name: "Frutta", Price: 2.43, Category: "Frutta"},
		},
	},
	{
		Name: "Dolce",
		Products: []t.Product{
			{Id: 16, Name: "Crema allo zabaione", Price: 8.26, Category: "Dolce"},
		},
	},
}

var DataProductList []t.Product = []t.Product{
	{Id: 1, Name: "Pasta al pomodoro", Price: 8.73, Category: "Primo", Visible: true},
	{Id: 2, Name: "Strudel salato ", Price: 3.74, Category: "Primo", Visible: false},
	{Id: 3, Name: "Canederli alle rape rosse", Price: 7.50, Category: "Primo", Visible: true},
	{Id: 4, Name: "Pasta alla calabrese", Price: 6.78, Category: "Primo", Visible: false},
	{Id: 5, Name: "Pasta all'olio", Price: 3.91, Category: "Primo", Visible: true},
	{Id: 6, Name: "Crespella al radicchio", Price: 4.12, Category: "Primo", Visible: false},
	{Id: 7, Name: "Scaloppa di tacchino alla bolognese", Price: 2.56, Category: "Secondo", Visible: true},
	{Id: 8, Name: "Hamburger di patate e formaggio", Price: 8.81, Category: "Secondo", Visible: false},
	{Id: 9, Name: "Moscardini in umido", Price: 3.53, Category: "Secondo", Visible: false},
	{Id: 10, Name: "Filetto di sardina croccante all' arancia", Price: 8.98, Category: "Secondo", Visible: true},
	{Id: 11, Name: "Insalata", Price: 5.31, Category: "Contorno", Visible: true},
	{Id: 12, Name: "Purè di patate", Price: 0.72, Category: "Contorno", Visible: false},
	{Id: 13, Name: "Zucca al forno", Price: 2.76, Category: "Contorno", Visible: false},
	{Id: 14, Name: "Macedonia", Price: 2.62, Category: "Frutta", Visible: true},
	{Id: 15, Name: "Frutta", Price: 2.43, Category: "Frutta", Visible: true},
	{Id: 16, Name: "Crema allo zabaione", Price: 8.26, Category: "Dolce", Visible: true},
}
