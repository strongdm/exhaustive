package implicitswitch

// Package-level ignore: all switches in the package should be skipped
// in non-explicit (implicit) mode, without per-switch //exhaustive:ignore.

func _ignored() {
	var d Direction

	// this should not report: package-level ignore is active.
	switch d {
	case N:
	case S:
	case W:
	default:
	}
}

func _enforceOverride() {
	var d Direction

	// this should report: per-switch enforce overrides package-level ignore.
	//exhaustive:enforce
	switch d { // want "^missing cases in switch of type implicitswitch.Direction: implicitswitch.E, implicitswitch.directionInvalid$"
	case N:
	case S:
	case W:
	default:
	}
}

func _exhaustive() {
	var d Direction

	// this should not report: all members are listed (and package-level ignore).
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

	// outer should not report: package-level ignore.
	switch d {
	case N:
	case S:
	case W:
	default:
		// inner should also not report: package-level ignore applies to nested switches.
		switch d {
		case N:
		default:
		}
	}
}

func _nestedEnforced() {
	var d Direction

	// outer should not report: package-level ignore.
	switch d {
	case N:
	case S:
	case W:
	default:
		// inner should report: per-switch enforce overrides.
		//exhaustive:enforce
		switch d { // want "^missing cases in switch of type implicitswitch.Direction: implicitswitch.E, implicitswitch.directionInvalid$"
		case N:
		case S:
		case W:
		default:
		}
	}
}
