package explicitswitch

// Package-level enforce: all switches in the package should be checked
// even in explicit mode, without per-switch //exhaustive:enforce.

func _enforced() {
	var d Direction

	// this should report: package-level enforce is active.
	switch d { // want "^missing cases in switch of type explicitswitch.Direction: explicitswitch.E, explicitswitch.directionInvalid$"
	case N:
	case S:
	case W:
	default:
	}
}

func _ignoredOverride() {
	var d Direction

	// this should not report: per-switch ignore overrides package-level enforce.
	//exhaustive:ignore
	switch d {
	case N:
	case S:
	case W:
	default:
	}
}

func _exhaustive() {
	var d Direction

	// this should not report: all members are listed.
	switch d {
	case N:
	case E:
	case S:
	case W:
	case directionInvalid:
	}
}

func _nested() {
	var d Direction

	// outer should report.
	switch d { // want "^missing cases in switch of type explicitswitch.Direction: explicitswitch.E, explicitswitch.directionInvalid$"
	case N:
	case S:
	case W:
	default:
		// inner should also report: package-level enforce applies to nested switches.
		switch d { // want "^missing cases in switch of type explicitswitch.Direction: explicitswitch.directionInvalid$"
		case N:
		case E:
		case S:
		case W:
		default:
		}
	}
}

func _nestedIgnored() {
	var d Direction

	// outer should report.
	switch d { // want "^missing cases in switch of type explicitswitch.Direction: explicitswitch.E, explicitswitch.directionInvalid$"
	case N:
	case S:
	case W:
	default:
		// inner should not report: per-switch ignore.
		//exhaustive:ignore
		switch d {
		case N:
		default:
		}
	}
}
