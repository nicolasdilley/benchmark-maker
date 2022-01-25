// type = all
package main

func main() {
	// CP
	applyOnce(func() {
		// CS
	})
}

func applyOnce(f func()) {
	f()
}
