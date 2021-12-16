package main

import (
	"io/ioutil"
	"os"
	"strings"
)

var (
	small_bound   string = "4"
	large_bound   string = "1000"
	unknown_bound string = "runtime.NumCPU()"
)

type Program struct {
	Name    string
	Content string
}

func main() {

	// Take the buggy programs
	buggy_programs_path := os.Args[1]

	buggy_programs := []Program{}

	// Take the haddock programs
	haddock_programs_path := os.Args[2]
	haddock_programs := []Program{}

	files, err := ioutil.ReadDir(buggy_programs_path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(buggy_programs_path + "/" + file.Name())

		if err != nil {
			panic(err)
		}

		buggy_programs = append(buggy_programs, Program{Name: strings.Replace(file.Name(), ".go", "", -1), Content: string(content)})
	}

	files, err = ioutil.ReadDir(haddock_programs_path)

	for _, file := range files {
		content, err := ioutil.ReadFile(haddock_programs_path + "/" + file.Name())

		if err != nil {
			panic(err)
		}

		haddock_programs = append(haddock_programs, Program{Name: strings.Replace(file.Name(), ".go", "", -1), Content: string(content)})
	}

	// Produce the cartesian products of the benchmarks
	os.Mkdir("benchmarks", 0777)

	for _, buggy_program := range buggy_programs {
		benchmark := `package main
		func main() {
			go haddock()
			bug()
		}

		`

		benchmark += buggy_program.Content

		for _, haddock_program := range haddock_programs {

			new_benchmark := benchmark + "\n\n" + haddock_program.Content

			if strings.Contains(haddock_program.Content, "bound") {
				// Generate small bound
				small_bound_program := strings.Replace(new_benchmark, "*", small_bound, -1)

				ioutil.WriteFile("./benchmarks/"+buggy_program.Name+"_small_"+haddock_program.Name+".go", []byte(small_bound_program), 0666)

				// Generate big bound

				large_bound_program := strings.Replace(new_benchmark, "*", large_bound, -1)

				ioutil.WriteFile("./benchmarks/"+buggy_program.Name+"_large_"+haddock_program.Name+".go", []byte(large_bound_program), 0666)

				// Generate unbounded bound
				unknown_bound_program := new_benchmark[:12] + "\n import \"runtime\" \n"

				unknown_bound_program += strings.Replace(new_benchmark[12:], "*", unknown_bound, -1)

				ioutil.WriteFile("./benchmarks/"+buggy_program.Name+"_unbounded_"+haddock_program.Name+".go", []byte(unknown_bound_program), 0666)
			} else {

				// need to add the import "timeout" at the top of declaration
				if strings.Contains(haddock_program.Name, "timeout") {
					new_benchmark = new_benchmark[:12] + "\n import \"time\" \n" + new_benchmark[12:]
				}
				err := ioutil.WriteFile("./benchmarks/"+buggy_program.Name+"_"+haddock_program.Name+".go", []byte(new_benchmark), 0666)

				if err != nil {
					panic(err)
				}
			}
		}

	}

}
