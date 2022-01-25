package main

type ContextType int 


const (
	CHAN ContextType = iota
	WG 
	MUTEX
	ALL
)


type Context struct {
	Name    string // The name of the context
	Content string
	Type string 
}

// Generate the full set of models from the context, given a snippet
// dont forget to replace bound with actual value
func generateModels() {

}

func InitContext(file *os.File) Context {
	ctx := Context{
		Name: strings.Replace(file.Name(), ".go", "", -1), 
		Content: string(content)
		Type: extractType(content),
		Bounds: extractBounds(content)
	}

	return ctx
}

func (c Context) Print() {
	fmt.Println(" ==== ======= ===")
	fmt.Println(" NAME : ", context.Name)
	fmt.Println(" Type : ", context.Type)
	fmt.Println(" Bounds : ", context.bounds)
	fmt.Println(" ==== ======= ===")
}


func extractType(content string) ContextType {
	t := ALL

	lines := strings.Split(content, "\n")


	for _, line := range lines {
		if strings.Contains(line, "type =") {
			ty = strings.Trim(strings.Split(line, "=")[1], " ")

			switch ty {
			case "CHAN": 
				t = CHAN
			case "WG": 
				t = WG
			case "MUTEX": 
				t = MUTEX
			default:
				t = ALL
			}
		} 
	}
	return t
}


func extractBounds(content string) []int {
	bounds := []int{}

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if strings.Contains(line, "bounds =") {
			splitted_bounds = strings.Split(strings.Split(line, "=")[1],",")

			for _, bound := range splitted_bounds {
				bound = strings.Trim(bound, " ") // removing outside spaces
				
				b,err := strconv.Atoi(bound)

				if err != nil {
					panic(err)
				}

				bounds = append(bounds, b)
			}
		} 
	}

	return bounds
}

