package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	//Déclaration de variable
	var listFile []string
	var solution []string

	file, err := os.Open("File.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//Itération sur le ficher + ajout de chaque ligne dans le tableau listeFile
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		listFile = append(listFile, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	solution = append(solution, listFile[0])
	solution = append(solution, listFile[len(listFile)-1])

	for j, mot := range listFile {
		if mot == "before" {
			v, _ := strconv.Atoi(listFile[j+1]) //v = int du mot après before
			// v-1 car listeFile commence a 0 pas 1
			s := listFile[v-1]
			solution = append(solution, s)

			for k, lettre := range mot {
				if k == 1 {
					x := int(lettre) / len(mot)              //Division de la valeur ASCII de la 2ème lettre de before par longueur de before
					solution = append(solution, listFile[x]) //Ajout du x-ème mot de la liste
				}
			}
		}
	}

	fmt.Println(solution)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())
}
