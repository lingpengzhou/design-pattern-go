package Template

import "log"

type Game interface {
	initialize()
	startGame()
	endPay()
}

type GameTemplate struct {
	Game
}

func setGameTemplate(game Game) *GameTemplate {
	return &GameTemplate{
		game,
	}
}

func (gameTemplate *GameTemplate) play() {
	gameTemplate.Game.initialize()
	gameTemplate.Game.startGame()
	gameTemplate.Game.endPay()
}

type Cricket struct {
}

func (cricket *Cricket) initialize() {
	log.Println("Cricket Game Initialized! Start playing.")
}

func (cricket *Cricket) startGame() {
	log.Println("Cricket Game Started. Enjoy the game!")
}

func (cricket *Cricket) endPay() {
	log.Println("Cricket Game Finished!")
}

type FootBall struct {
}

func (footBall *FootBall) initialize() {
	log.Println("FootBall Game Initialized! Start playing.")
}

func (footBall *FootBall) startGame() {
	log.Println("FootBall Game Started. Enjoy the game!")
}

func (footBall *FootBall) endPay() {
	log.Println("FootBall Game Finished!")
}
