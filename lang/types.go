package lang

type PredictIntentResponse struct {
	Probabilities []struct
	{
		Label string `json:"label"`
		Probability float64 `json:"probability"`
	} `json:"probabilities"`
	Object string `json:"object"`
}