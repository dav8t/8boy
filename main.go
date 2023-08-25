package main

import (
	"flag"
	"github.com/faiface/mainthread"
	"github.com/sqweek/dialog"
	"runtime/pprof"
)

var (
	mute       = flag.Bool("mute", false, "mute sound output")
	dmgMode    = flag.Bool("dmg", false, "set to force dmg mode")
	unlocked   = flag.Bool("unlocked", false, "if to unlock the cpu speed(debugging)")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file(debugging)")
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

	fmt.println(fmt.Sprintf(logo, version))
	fmt.Printf("APU: %v\nCGB: %v\nROM: %v\n", !*mute, !*dmgMode, rom)
}

func getROM() string {
	rom := flag.Arg(0)
	if rom == "" {
		mainthread.Call(func() {
			var err error
			rom, err = dialog.File().
				Filter("GameBoy ROM File").Load()
		})
	}
	return rom
}
