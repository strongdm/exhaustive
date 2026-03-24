package explicitmap

// Package-level enforce: all map literals in the package should be checked
// even in explicit mode, without per-map //exhaustive:enforce.

func _enforced() {
	// this should report: package-level enforce is active.
	_ = map[Direction]string{ // want "^missing keys in map of key type explicitmap.Direction: explicitmap.E, explicitmap.directionInvalid$"
		N: "north",
		S: "south",
		W: "west",
	}
}

func _ignoredOverride() {
	// this should not report: per-map ignore overrides package-level enforce.
	//exhaustive:ignore
	_ = map[Direction]string{
		N: "north",
		S: "south",
		W: "west",
	}
}

func _exhaustive() {
	// this should not report: all members are listed.
	_ = map[Direction]string{
		N:                "north",
		E:                "east",
		S:                "south",
		W:                "west",
		directionInvalid: "invalid",
	}
}
