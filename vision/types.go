package vision

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
