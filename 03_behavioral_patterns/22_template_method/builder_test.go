package templatemethod

func Example_one() {
	builders := []Builder{
		NewCPP11Builder(),
		NewCPP20Builder(),
	}

	for _, builder := range builders {
		builder.Build()
	}

	// Output:
	// checking code syntax by C++11 standard
	// compiling code by GCC 4.8
	// linking code by GNU linker
	// checking code syntax by C++20 standard
	// compiling code by GCC 8
	// linking code by GNU linker
}
