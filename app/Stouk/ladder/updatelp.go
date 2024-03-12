package ladder

//cette fonction aura besoin en entré de l'id du joueur
// ,win or loose, le score 
// avec le rang moyen ont le compare a son rang si inferieur 
//il gagne moin et vice versa en le comparant au rang de 
//l'adversaire 
//

import (
	"data"
)
//le rank il va de 0 a autant que le veut mais il doit avoir un max theorique de 4500 a plus on monte plus c'est dur de monté
//data : 0 rank / 1 mmr / 2 nombre  win / 3 nombre lose / 5 rankmoyen
// de base on perd 20 on gagne 20

func updateRankforPlayer(IdPlayer int,win bool,score []int) int{
	data := data.GetForUpdateLadder(IdPlayer)
	if win {

	}else{

	}

}

func GainLP(data []int,score []int){

}

func LooseLP(data []int,score []int){

}