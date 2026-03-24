//exhaustive:enforce
package packageconflict // want "^failed to parse directives"

type Direction int // want Direction:"^N,E,S,W$"

const (
	N Direction = 1
	E Direction = 2
	S Direction = 3
	W Direction = 4
)
