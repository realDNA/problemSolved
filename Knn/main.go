package main

import (
	"Knn/knn"
	"Knn/point"
	"fmt"
)

type Point = point.Point


func main() {
	// test
	X := []Point{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {10,11}, {11,12},{12,13}}
	y := []float64{0, 0, 0, 0, 1,1,1}
	Knn := knn.KNeighborsClassifier{Nneighbors: 5}
	Knn.Fit(X, y)
	testData := Point{15, 11}
	predict := Knn.Predict(testData)
	fmt.Println(predict)
	predictProbability := Knn.PredictProb(testData)
	fmt.Println(predictProbability)
}
