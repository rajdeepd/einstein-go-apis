package vision

const BASE_URL_V1 string = "https://api.einstein.ai/v1/vision"
const BASE_URL string = "https://api.einstein.ai/v2/vision"
const PREDICT_URL string = BASE_URL + "/predict"
const DETECT_URL string = BASE_URL + "/detect"
const DATASET_URL string = BASE_URL + "/datasets"
const DATASET_UPLOAD_SYNC_URL = DATASET_URL + "/upload/sync"

const DATASET_UPLOAD_SYNC_URL_V1 = BASE_URL_V1 + "/datasets/upload/sync"
const TRAIN_URL string = BASE_URL + "/train"
