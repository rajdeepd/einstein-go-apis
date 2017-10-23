package common

type DataSet struct {
	ID int `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	LabelSummary struct {
		Labels []struct {
			ID int `json:"id"`
			DatasetID int `json:"datasetId"`
			Name string `json:"name"`
			NumExamples int `json:"numExamples"`
		} `json:"labels"`
	} `json:"labelSummary"`
	TotalExamples int `json:"totalExamples"`
	TotalLabels int `json:"totalLabels"`
	Available bool `json:"available"`
	Object string `json:"object"`
	StatusMsg string `json:"statusMsg,omitempty"`
}

type DataSetList struct {
	Object string `json:"object"`
	DataSet []struct {
		ID int `json:"id"`
		Name string `json:"name"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
		LabelSummary struct {
			Labels []struct {
				ID int `json:"id"`
				DatasetID int `json:"datasetId"`
				Name string `json:"name"`
				NumExamples int `json:"numExamples"`
			} `json:"labels"`
		} `json:"labelSummary"`
		TotalExamples int `json:"totalExamples"`
		TotalLabels int `json:"totalLabels"`
		Available bool `json:"available"`
		Object string `json:"object"`
		StatusMsg string `json:"statusMsg,omitempty"`
	} `json:"data"`
}


type TrainDataSetResponse struct {
	CreatedAt        string      `json:"createdAt"`
	DatasetID        int         `json:"datasetId"`
	DatasetVersionID int         `json:"datasetVersionId"`
	Epochs           int         `json:"epochs"`
	LearningRate     float64     `json:"learningRate"`
	ModelID          string      `json:"modelId"`
	ModelType        string      `json:"modelType"`
	Name             string      `json:"name"`
	Object           string      `json:"object"`
	Progress         int         `json:"progress"`
	QueuePosition    int         `json:"queuePosition"`
	Status           string      `json:"status"`
	TrainParams      interface{} `json:"trainParams"`
	TrainStats       interface{} `json:"trainStats"`
	UpdatedAt        string      `json:"updatedAt"`
}

type Model struct {
	CreatedAt string `json:"createdAt"`
	MetricsData struct { F1 []float64 `json:"f1"`
	Labels []string `json:"labels"`
	TestAccuracy float64 `json:"testAccuracy"`
	TrainingLoss float64 `json:"trainingLoss"`
	ConfusionMatrix [][]int `json:"confusionMatrix"`
	TrainingAccuracy float64 `json:"trainingAccuracy"`
	PrecisionRecallCurve struct { F1 []float64 `json:"f1"`
		Recall []float64 `json:"recall"`
		Precision []float64 `json:"precision"`
		Threshold []float64 `json:"threshold"` }`json:"precisionRecallCurve"`
	} `json:"metricsData"`
	ID string `json:"id"`
	Object string `json:"object"`
	}
type PredictResponse struct {
	Probabilities []struct {
		Label string `json:"label"`
		Probability float64 `json:"probability"`
	} `json:"probabilities"`
	Object string `json:"object"`
}
