package main

import (
	"os"
	"mycode"
)

func main() {
	// get input
	file_input := os.Args[1]
	file_output := os.Args[2]
	text, err := os.ReadFile(file_input)
	if err != nil {
		panic(err)
	}
	newString := mycode.Split(string(text)) // split words
	// applay funchen of words
	newString = mycode.UpLowCap_numpers(newString)
	newString = mycode.UpLowCap(newString)
	newString = mycode.SpichelCarecter(newString)
	newString = mycode.MultChar(newString)
	newString = mycode.Hex_Bin(newString)
	newString = mycode.ChangAtoAn(newString)
	newString = mycode.Apotrof(newString)
	// join out put
	sentence := mycode.Join(newString)
	// Create a new file and save out put in file
	file, err := os.Create(file_output)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(sentence)
	if err != nil {
		panic(err)
	}
}
