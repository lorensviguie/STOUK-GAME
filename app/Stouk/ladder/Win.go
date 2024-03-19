package ladder

import "math"

func calculateMultiplierCoefficientForWin(rank, mmr, medium_rank int) float64 {
	Final_coef := 0.0
	diff := float64(rank - mmr)
	baseCoefficient := 1.0
	influenceFactor := 0.005
	Mmr_Coefficient := baseCoefficient + math.Abs(diff)*influenceFactor
	Medium_Rank_Coefficient := math.Abs(float64(rank-medium_rank)) * influenceFactor
	if rank > mmr {
		Final_coef = 1 - (Mmr_Coefficient - 1 + Medium_Rank_Coefficient)
		return Final_coef
	}
	Final_coef = Mmr_Coefficient + Medium_Rank_Coefficient
	return Final_coef
}
