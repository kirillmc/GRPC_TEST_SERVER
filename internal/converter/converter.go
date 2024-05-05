package converter

import (
	model "github.com/kirillmc/data_filler/pkg/program_v1"
	desc "github.com/kirillmc/grpc_test_server/pkg/program_v3"
)

func ToResponseProgramsFromRepo(programs *model.TrainPrograms) *desc.TrainPrograms {
	return &desc.TrainPrograms{TrainPrograms: toResponseProgramFromRepo(programs)}
}

func toResponseProgramFromRepo(programsSet *model.TrainPrograms) []*desc.TrainProgram {
	var trainProgramsSet []*desc.TrainProgram
	for _, program := range programsSet.TrainPrograms {
		trainProgramsSet = append(trainProgramsSet, &desc.TrainProgram{
			Id:          program.Id,
			ProgramName: program.ProgramName,
			Description: program.Description,
			Status:      program.Status,
			TrainDays:   toResponseTrainDaysFromRepo(program.TrainDays)})
	}

	return trainProgramsSet
}
func toResponseTrainDaysFromRepo(trainDays []*model.TrainDay) []*desc.TrainDay {
	var trainDaysSet []*desc.TrainDay
	for _, trainDay := range trainDays {
		trainDaysSet = append(trainDaysSet, &desc.TrainDay{
			Id:          trainDay.Id,
			DayName:     trainDay.DayName,
			Description: trainDay.Description,
			Exercises:   toResponseExercisesFromRepo(trainDay.Exercises)})
	}

	return trainDaysSet
}
func toResponseExercisesFromRepo(exercises []*model.Exercise) []*desc.Exercise {
	var exercisesSet []*desc.Exercise
	for _, exercise := range exercises {
		exercisesSet = append(exercisesSet, &desc.Exercise{Id: exercise.Id,
			ExerciseName: exercise.ExerciseName,
			Pictures:     exercise.Pictures,
			Description:  exercise.Description,
			Sets:         toResponseSetsFromRepo(exercise.Sets)})
	}

	return exercisesSet
}
func toResponseSetsFromRepo(sets []*model.Set) []*desc.Set {
	var setsSet []*desc.Set
	for _, set := range sets {
		setsSet = append(
			setsSet,
			&desc.Set{Id: set.Id,
				Quantity: set.Quantity,
				Weight:   set.Weight,
			})
	}
	return setsSet
}
