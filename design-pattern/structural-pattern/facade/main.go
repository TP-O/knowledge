package facade

func Run() {
	wb := NewWomboCombo()

	println("Pefect:")
	wb.PerfectWomboCombo()

	println("\nFailed:")
	wb.FailedWomboCombo()
}
