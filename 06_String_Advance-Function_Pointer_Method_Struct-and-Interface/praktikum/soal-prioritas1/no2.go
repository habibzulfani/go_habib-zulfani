package main

import "fmt"

type Student struct {
	name  []string
	score []float64
}

func (s Student) nilaiStudent() (float64, float64) {
	var max float64 = 0
	var min float64 = 0

	for _, i := range s.score {
		if min == 0 || i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	return min, max

}

func (s Student) nilaiAverage() float64 {
	var totalScore float64 = 0
	var average float64

	for _, i := range s.score {
		totalScore += i
	}

	n := float64(len(s.score))

	average = totalScore / n

	return average
}

func (s Student) merge() map[string]float64 {
	studentMap := make(map[string]float64)
	
	for i := 0; i < len(s.name); i++ {
		
		studentMap[s.name[i]] = s.score[i]
	}
	return studentMap
}

func (s Student) findNameByScore(score float64) string {
	studentMap := s.merge()
	for name, s := range studentMap {
		if s == score {
			return name
		}
	}
	return ""
}


func main() {
	student := Student{
		name:  []string{"Rizky", "Kobar", "Ismail", "Umam", "Yopan"},
		score: []float64{80, 75, 70, 75, 60},
	}

	minScore, maxScore := student.nilaiStudent()
	averageScore := student.nilaiAverage()
	minName := student.findNameByScore(minScore)
	maxName := student.findNameByScore(maxScore)

	fmt.Printf("Nilai minimum: %s(%d)\n", minName ,int(minScore))
	fmt.Printf("Nilai maksimum: %s(%d)\n", maxName, int(maxScore))
	fmt.Printf("Nilai rata - rata: %.2f\n", averageScore)


}
