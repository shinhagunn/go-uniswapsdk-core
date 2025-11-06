package utils

import "github.com/shinhagunn/go-uniswapsdk-core/src/entities"

func ComputePriceImpact(midPrice *entities.Price, inputAmount, outputAmount *entities.CurrencyAmount) (*entities.Percent, error) {
	quotedOutputAmount, err := midPrice.Quote(inputAmount)
	if err != nil {
		return nil, err
	}

	subtracted, err := quotedOutputAmount.Subtract(outputAmount)
	if err != nil {
		return nil, err
	}

	priceImpact := subtracted.Divide(quotedOutputAmount.GetFraction())
	return entities.NewPercent(priceImpact.GetNumerator(), priceImpact.GetDenominator()), nil
}
