package tenniskata

type tennisGame1 struct {
	scores      map[string]int
	player1Name string
	player2Name string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	game := &tennisGame1{
		scores: map[string]int{
			player1Name: 0,
			player2Name: 0,
		},
		player1Name: player1Name,
		player2Name: player2Name}

	return game
}

func (game *tennisGame1) WonPoint(playerName string) {
	game.scores[playerName] += 1
}

func (game *tennisGame1) GetScore() string {
	score := ""
	tempScore := 0
	if game.scores[game.player1Name] == game.scores[game.player2Name] {
		switch game.scores[game.player1Name] {
		case 0:
			score = "Love-All"
		case 1:
			score = "Fifteen-All"
		case 2:
			score = "Thirty-All"
		default:
			score = "Deuce"
		}
	} else if game.scores[game.player1Name] >= 4 || game.scores[game.player2Name] >= 4 {
		minusResult := game.scores[game.player1Name] - game.scores[game.player2Name]
		if minusResult == 1 {
			score = "Advantage player1"
		} else if minusResult == -1 {
			score = "Advantage player2"
		} else if minusResult >= 2 {
			score = "Win for player1"
		} else {
			score = "Win for player2"
		}
	} else {
		for i := 1; i < 3; i++ {
			if i == 1 {
				tempScore = game.scores[game.player1Name]
			} else {
				score += "-"
				tempScore = game.scores[game.player2Name]
			}
			switch tempScore {
			case 0:
				score += "Love"
			case 1:
				score += "Fifteen"
			case 2:
				score += "Thirty"
			case 3:
				score += "Forty"
			}
		}
	}
	return score
}
