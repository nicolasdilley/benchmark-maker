package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

type Verifier interface {
	Verify(Model) Verification
}

type Verification struct {
	ToolName  string
	ModelName string
	Result    VerificationResult
	Time      int64
	Type      int
	Feedback  string
	Content   string
}

type VerificationResult string

const (
	NO_BUG     VerificationResult = "No bug reported"
	BUG        VerificationResult = "Bug reported"
	CRASH      VerificationResult = "Crashed"
	NO_SUPPORT VerificationResult = "not supported"
)

func verifyModels(models map[string]map[string][]Model) {

	verifiers := []Verifier{
		GomelaVerifier{},
		GCatchVerifier{},
		Godel2Verifier{},
	}

	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	keys := make([]string, 0, len(models))
	for key := range models {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, name := range keys {
		names := make([]string, 0, len(models[name]))
		for key := range models[name] {
			names = append(names, key)
		}
		sort.Strings(names)
		names = putMinimalInFrontOfQueue(names)

		fmt.Println("===== Verifying ", name, " =====")
		f.WriteString("\n===== Verifying " + name + " =====\n")

		for _, modelName := range names {

			for _, model := range models[name][modelName] {
				if model.Bound != "" {
					fmt.Print("& \\textit{"+modelName, "(", model.Bound, ")} & ")
				} else {
					fmt.Print("& \\textit{"+modelName, "} & ")
				}
				f.WriteString("& \\textit{" + modelName + "} & ")

				for i, verifier := range verifiers {
					ver := verifier.Verify(model)

					to_print := ""

					switch ver.Result {
					case BUG:
						to_print = " \\cmark "
					case NO_SUPPORT:
						to_print = " \\nosupport "
					case NO_BUG:
						to_print = " \\xmark "
					case CRASH:
						to_print = " \\crash "
					}

					to_print += fmt.Sprintf(" %d", ver.Time)

					fmt.Print(to_print)
					f.WriteString(to_print)

					if i < len(verifiers)-1 {
						fmt.Print(" & ")
						f.WriteString(" & ")
					}
				}

				fmt.Println("\\\\")
				f.WriteString("\\\\ \n")
			}
			fmt.Println()
			f.WriteString("\n")
		}
	}

}

func putMinimalInFrontOfQueue(names []string) []string {
	for i, name := range names {
		if name == "minimal" {
			return append([]string{"minimal"}, append(names[:i], names[i+1:]...)...)
		}
	}
	return names
}