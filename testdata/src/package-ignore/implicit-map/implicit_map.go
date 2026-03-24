package implicitmap

// Package-level ignore: all map literals in the package should be skipped
// in non-explicit (implicit) mode, without per-map //exhaustive:ignore.

func _ignored() {
	// this should not report: package-level ignore is active.
	_ = map[Direction]string{
		N: "north",
		S: "south",
		W: "west",
	}
}

func _enforceOverride() {
	// this should report: per-map enforce overrides package-level ignore.
	//exhaustive:enforce
	_ = map[Direction]string{ // want "^missing keys in map of key type implicitmap.Direction: implicitmap.E, implicitmap.directionInvalid$"
		N: "north",
		S: "south",
		W: "west",
	}
}

func _exhaustive() {
	// this should not report: all members are listed (and package-level ignore).
	_ = map[Direction]string{
		N:                "north",
		E:                "east",
		S:                "south",
		W:                "west",
		directionInvalid: "invalid",
	}
}
