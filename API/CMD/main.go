package main




func main() {
ns := GetNutritionalScore(
	NutritionalData{
		Energy:EnergyFromcal(),
		Sugars: SugarGram(),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium: SodiumMilligram(),
		Fruits: FruitsPercent(),
		Fibre: FibreGram(),
		Protein: ProteinGram(),
	},Food)
}