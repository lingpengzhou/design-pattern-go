package Template

import "testing"

func Test_Template(t *testing.T) {
	cr := Cricket{}
	game := setGameTemplate(&cr)
	game.play()
	ft := FootBall{}
	game = setGameTemplate(&ft)
	game.play()

}
