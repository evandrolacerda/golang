package main

import (
	"fmt"
	"sort"
)

func main() {
	namesAndAges := make(map[string]int)

	namesAndAges["evandro"] = 38
	namesAndAges["leandro"] = 33
	namesAndAges["luzia"] = 37
	namesAndAges["helena"] = 12
	namesAndAges["henrique"] = 7
	namesAndAges["alice"] = 87

	var names []string

	for name, _ := range namesAndAges {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s tem %d anos \n", name, namesAndAges[name])
	}

	booksAndLanguages := map[string]string{
		"Arquitetura Limpa":           "Java",
		"Web Interativa":              "PHP",
		"Java Para Iniciantes":        "Java",
		"Programação Java para a WEB": "Java",
		"PHP Moderno":                 "PHP",
	}

	ocurrences := make(map[string]int)

	for _, language := range booksAndLanguages {
		ocurrences[language]++
	}

	for language, ocurrence := range ocurrences {
		fmt.Printf("Há %d livros sobre a linguagem %s \n", ocurrence, language)
	}
	nameMarcelo, indiceExiste := namesAndAges["marcelo"]

	if !indiceExiste {
		fmt.Printf("%s não existe no mapa \n", "Marcelo")
	}

	fmt.Println(nameMarcelo, indiceExiste)

}
