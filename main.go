package main

import (
	"io/ioutil"
	"os"
	"strings"
)

const (
	CONTEXT_PATH = "./contexts"
	SNIPPET_PATH = "./snippets"
)

func main() {

	// Takes a list of contexts and a list of snippets
	// Each contexts and snippets are in a folder (contexts and snippets)
	// The contexts have 2 meta variables:
	//    ==== Type : which specify which snippets it accepts (ch, wg, mutex or * (all))
	//    ==== Bounds : what value to give to 'bound' ([0,10,10000])
	// The snippets have one meta variable :
	//    ==== Type : which specify what the type of the snippets is (ch, wg or mutex)
	// The program used the command design pattern to support any number of tools
	// 				(the command is given the file to verify (the full model) and verify it with the specific tools,
	//				it returns true or false based on the result of the verification) (true meaning there is a bug)

	contexts := []Context{}

	files, err := ioutil.ReadDir(CONTEXT_PATH)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(CONTEXT_PATH + "/" + file.Name())

		if err != nil {
			panic(err)
		}

		contexts = append(contexts, InitContext(content))

		for _, context := range contexts {
			context.Print()
			fmt.Println()
			fmt.Println()
		}
	}

	// files, err = ioutil.ReadDir(haddock_programs_path)

	// for _, file := range files {
	// 	content, err := ioutil.ReadFile(haddock_programs_path + "/" + file.Name())

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	haddock_programs = append(haddock_programs, Program{Name: strings.Replace(file.Name(), ".go", "", -1), Content: string(content)})
	// }

	// // Produce the cartesian products of the benchmarks
	// os.Mkdir("benchmarks", 0777)

	// for _, buggy_program := range buggy_programs {
	// 	benchmark := `package main
	// 	func main() {
	// 		go haddock()
	// 		bug()
	// 	}

	// 	`

	// 	benchmark += buggy_program.Content

	// 	for _, haddock_program := range haddock_programs {

	// 		new_benchmark := benchmark + "\n\n" + haddock_program.Content

	// 		if strings.Contains(haddock_program.Content, "bound") {
	// 			// Generate small bound
	// 			small_bound_program := strings.Replace(new_benchmark, "*", small_bound, -1)

	// 			ioutil.WriteFile("./benchmarks/"+buggy_program.Name+"_small_"+haddock_program.Name+".go", []byte(small_bound_program), 0666)

	// 			// Generate big bound

	// 			large_bound_program := strings.Replace(new_benchmark, "*", large_bound, -1)

	// 			ioutil.WriteFile("./benchmarks/"+buggy_program.Name+"_large_"+haddock_program.Name+".go", []byte(large_bound_program), 0666)

	// 			// Generate unbounded bound
	// 			unknown_bound_program := new_benchmark[:12] + "\n import \"runtime\" \n"

	// 			unknown_bound_program += strings.Replace(new_benchmark[12:], "*", unknown_bound, -1)

	// 			ioutil.WriteFile("./benchmarks/"+buggy_program.Name+"_unbounded_"+haddock_program.Name+".go", []byte(unknown_bound_program), 0666)
	// 		} else {

	// 			// need to add the import "timeout" at the top of declaration
	// 			if strings.Contains(haddock_program.Name, "timeout") {
	// 				new_benchmark = new_benchmark[:12] + "\n import \"time\" \n" + new_benchmark[12:]
	// 			}
	// 			err := ioutil.WriteFile("./benchmarks/"+buggy_program.Name+"_"+haddock_program.Name+".go", []byte(new_benchmark), 0666)

	// 			if err != nil {
	// 				panic(err)
	// 			}
	// 		}
	// 	}

	// }

}
