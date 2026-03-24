package nodirective

// No package-level enforce: in explicit mode, switches without
// per-switch enforce should not be checked.

func _notChecked() {
	var d Direction

	// this should not report: no enforce directive anywhere.
	switch d {
	case N:
	case S:
	case W:
	default:
	}
}

func _enforced() {
	var d Direction

	// this should report: per-switch enforce.
	//exhaustive:enforce
	switch d { // want "^missing cases in switch of type nodirective.Direction: nodirective.E, nodirective.directionInvalid$"
	case N:
	case S:
	case W:
	default:
	}
}
