package implicitswitch

// Package-level enforce in non-explicit (implicit) mode:
// switches are checked by default, ignore still works.

func _checked() {
	var d Direction

	// this should report: normal implicit checking.
	switch d { // want "^missing cases in switch of type implicitswitch.Direction: implicitswitch.E, implicitswitch.directionInvalid$"
	case N:
	case S:
	case W:
	default:
	}
}

func _ignored() {
	var d Direction

	// this should not report: per-switch ignore.
	//exhaustive:ignore
	switch d {
	case N:
	case S:
	case W:
	default:
	}
}
