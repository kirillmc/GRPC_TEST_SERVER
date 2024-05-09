package filler_pb

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	desc "github.com/kirillmc/grpc_test_server/pkg/program_v3"
)

const (
	oneCount    = 1
	smallCount  = 5
	normalCount = 11
	bigCount    = 31

	picturesCount = 5
)

func FindByteSizeOfProto(programs *desc.TrainPrograms) (int64, error) {
	numOfSets, err := proto.Marshal(programs)
	if err != nil {
		return 0.0, fmt.Errorf("fail to get json: %v", err)
	}

	return int64(len(numOfSets)), nil
}

func CreateEmptySetOfPrograms() *desc.TrainPrograms {
	return &desc.TrainPrograms{}
}

func CreateSetOfOneProgram() *desc.TrainPrograms {
	return &desc.TrainPrograms{
		TrainPrograms: fillTrainPrograms(oneCount),
	}
}

func CreateSmallSetOfPrograms() *desc.TrainPrograms {
	return &desc.TrainPrograms{
		TrainPrograms: fillTrainPrograms(smallCount),
	}
}

func CreateNormalSetOfPrograms() *desc.TrainPrograms {
	return &desc.TrainPrograms{
		TrainPrograms: fillTrainPrograms(normalCount),
	}
}

func CreateBigSetOfPrograms() *desc.TrainPrograms {
	return &desc.TrainPrograms{
		TrainPrograms: fillTrainPrograms(bigCount),
	}
}

func CreateOwnSetOfPrograms(count int) *desc.TrainPrograms {
	return &desc.TrainPrograms{
		TrainPrograms: fillTrainPrograms(count),
	}
}

func fillTrainPrograms(count int) []*desc.TrainProgram {
	var setOfTrainPrograms []*desc.TrainProgram
	for i := 1; i < count+1; i++ {
		trainPrograms := &desc.TrainProgram{
			Id:          int64(i),
			ProgramName: fmt.Sprintf("Название программы тренировок №%d", i),
			Description: fmt.Sprintf("Описание программы тренирово №%d", i),
			Status:      "На модерации",
			TrainDays:   fillTrainDays(count),
		}

		setOfTrainPrograms = append(setOfTrainPrograms, trainPrograms)
	}

	return setOfTrainPrograms
}

func fillTrainDays(count int) []*desc.TrainDay {
	var setOfTrainDays []*desc.TrainDay
	for i := 1; i < count+1; i++ {
		trainDay := &desc.TrainDay{
			Id:          int64(i),
			DayName:     fmt.Sprintf("Название тренировочного дня №%d", i),
			Description: fmt.Sprintf("Описание тренировочного дня №%d", i),
			Exercises:   fillExercises(count),
		}

		setOfTrainDays = append(setOfTrainDays, trainDay)
	}

	return setOfTrainDays
}

func fillExercises(count int) []*desc.Exercise {
	var setOfExercises []*desc.Exercise
	for i := 1; i < count+1; i++ {
		exercise := &desc.Exercise{
			Id:           int64(i),
			ExerciseName: fmt.Sprintf("Название упражнения №%d", i),
			Pictures:     fillPictures(picturesCount),
			Description:  fmt.Sprintf("Описание упражнения №%d", i),
			Sets:         fillSets(count),
		}

		setOfExercises = append(setOfExercises, exercise)
	}

	return setOfExercises
}

func fillPictures(count int) []string {
	var setOfPictures []string
	for i := 1; i <= count; i++ {
		setOfPictures = append(setOfPictures, fmt.Sprintf("http://picture_%d.png", i))
	}

	return setOfPictures
}

func fillSets(count int) []*desc.Set {
	var setOfSets []*desc.Set
	for i := 1; i < count+1; i++ {
		set := &desc.Set{
			Id:       int64(i),
			Quantity: 11,
			Weight:   55.5,
		}

		setOfSets = append(setOfSets, set)
	}

	return setOfSets
}
