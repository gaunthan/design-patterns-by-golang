package decorator

func Example_one() {
	var logger Logger = NewScreenLogger()

	logger.Log("hello")
	logger.Log("world")

	// Output:
	// helloworld
}

func Example_two() {
	var logger Logger = NewScreenLogger()
	logger = NewDateLoggerDecorator(logger)
	logger = NewLineLoggerDecorator(logger)

	logger.Log("hello")
	logger.Log("world")

	// Output:
	// [June 18, 2021] hello
	// [June 18, 2021] world
}
