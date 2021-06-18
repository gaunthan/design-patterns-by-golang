package adapter

func Example_one() {
	video := VideoOutputter{}
	video.Start()

	screen := Screen{}

	// VDI port is incompatible with HDMI port, so the
	//  following statement would cause a compile error:
	//
	//		screen.SetInput(video.GetOutput())
	//
	// To solve this problem, we need to use an adapter.
	adapter := HDMI2DVIAdapter{}
	adapter.SetInput(video.GetOutput())

	screen.SetInput(adapter.GetOutput())
	screen.Play()

	// Output:
	// 0123456789
}
