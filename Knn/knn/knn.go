package knn

import(
	"Knn/point"
	"sort"
)

type Point = point.Point

type KNeighborsClassifier struct {
	Nneighbors int
	TrainData   []Point
	Labels     []float64
}

func sortDistance(distances []float64) []float64 {
	sort.Float64s(distances)
	return distances
}

func pickFirstKNearestIndexes(distances []float64, k int) []int {
	FirstKNearestIndexes := []int{}
	mapStore := make(map[float64]int)

	for i, _ := range distances {
		mapStore[distances[i]] = i
	}

	sortedDistance := sortDistance(distances)

	for i, v := range sortedDistance {
		if i == k {
			break
		}
		FirstKNearestIndexes = append(FirstKNearestIndexes ,mapStore[v])
	}

	return FirstKNearestIndexes
}

func countKNereastNumber(FirstKNearestIndexes []int, y []float64) map[float64]int {
	mapStore := make(map[float64]int)

	for _,v := range FirstKNearestIndexes {
		map_key := y[v]
		if _, exist := mapStore[map_key]; exist {
			mapStore[map_key] += 1
		} else {
			mapStore[map_key] = 1
		}
	}
	return mapStore
}

func (knnP *KNeighborsClassifier) Fit(trainData []Point, labels []float64) {
	knnP.TrainData = trainData
	knnP.Labels = labels
}

func findMapMaxNumberKey(mapStore map[float64]int) float64 {
	var maxVal int
	var labelStore float64

	for k,v := range mapStore {
		if v>maxVal {
			labelStore = k
			maxVal = v
		}
	}
	return labelStore
}

func countMapProbability(mapStore map[float64]int) map[float64]float64 {
	mapProba := make(map[float64]float64)
	sum := 0

	for _,v := range mapStore {
		sum += v
	}

	for k,v := range mapStore {
		mapProba[k] = float64(v)/float64(sum)
	}
	return mapProba
}

func (knnP *KNeighborsClassifier) Predict(testData Point) float64 {
	distances := testData.Distance(knnP.TrainData)
	firstKNearestIndexes := pickFirstKNearestIndexes(distances, knnP.Nneighbors)
	mapStore := countKNereastNumber(firstKNearestIndexes, knnP.Labels)
	predictLable := findMapMaxNumberKey(mapStore)
	
	return predictLable
}

func (knnP *KNeighborsClassifier) PredictProb(testData Point) map[float64]float64 {
	distances := testData.Distance(knnP.TrainData)
	firstKNearestIndexes := pickFirstKNearestIndexes(distances, knnP.Nneighbors)
	mapStore := countKNereastNumber(firstKNearestIndexes, knnP.Labels)
	mapProba := countMapProbability(mapStore)

	return mapProba
}