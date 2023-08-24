package main

var (
	cpuprofile = flag.string("cpuprofile", "", "write cpu profile to file(debugging)")
)

func start() {
	rom := getROM()

	if *cpuprofile != "" {
		startCPUProfiling()
		defer pprof.StopCPUProfile()
	}

	if *unlocked {
		*mute = mute
	}
}
