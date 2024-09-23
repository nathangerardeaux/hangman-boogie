package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func random() string {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error:", err)
		//	return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	words := []string{}
	for scanner.Scan() {
		words = append(words, strings.Split(scanner.Text(), " ")...)
	}
	randomNumber := rand.Intn(len(words))
	mot := words[randomNumber]
	fmt.Println(mot)
	return mot
}
func motcache(mot string) string {
	cache := []rune{}
	for range mot {
		cache = append(cache, '_')
	}
	println(string(cache))

	// Transforme le mot en liste de lettres
	lettres := []rune{}
	for _, lettre := range mot {
		lettres = append(lettres, rune(lettre))
	}

	long := len(mot)
	if long < 5 {
		for i := 0; i < 2; i++ {
			rev := rand.Intn(long - 1)
			cache[rev] = lettres[rev]
		}
	} else if long >= 5 {
		for i := 0; i < 4; i++ {
			rev := rand.Intn(long - 1)
			cache[rev] = lettres[rev]
		}
	}
	return string(cache)
}

func main() {
	mot := random()
	cache := motcache(mot)
	fmt.Println("Mot caché:", cache)
}
