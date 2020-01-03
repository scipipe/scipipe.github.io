package main

import (
	"github.com/scipipe/scipipe"
	"github.com/scipipe/scipipe/components"
)

func main() {
	wf := scipipe.NewWorkflow("wf", 4)

	letterGlobber := components.NewFileGlobber(wf, "letter_globber", "letterfile_*.txt")
	numberGlobber := components.NewFileGlobber(wf, "number_globber", "numberfile_*.txt")

	fileCombiner := components.NewFileCombinator(wf, "file_combiner")
	fileCombiner.In("letters").From(letterGlobber.Out())
	fileCombiner.In("numbers").From(numberGlobber.Out())

	catenator := wf.NewProc("catenator", "cat {i:letters} {i:numbers} > {o:combined}")
	catenator.In("letters").From(fileCombiner.Out("letters"))
	catenator.In("numbers").From(fileCombiner.Out("numbers"))
	catenator.SetOut("combined", "{i:letters|basename|%.txt}.{i:numbers|basename|%.txt}.combined.txt")

	wf.Run()
}
