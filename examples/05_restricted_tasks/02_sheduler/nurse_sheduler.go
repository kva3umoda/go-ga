package main

import "fmt"

var (
	// список медсетер
	nursers = []string{"A", "B", "C", "D", "E", "F", "G", "H"}

	// пожелания медсестер к сменам – утренняя, дневная, ночная
	shiftPreferences = map[string][]float64{
		"A": {1, 0, 0},
		"B": {1, 1, 0},
		"C": {0, 0, 1},
		"D": {0, 1, 0},
		"E": {0, 0, 1},
		"F": {1, 1, 1},
		"G": {0, 1, 1},
		"H": {1, 1, 1}}
	// минимальное и максимальное количество медсестер в сменах – утренней, дневной, ночной
	shiftMin = []float64{2, 2, 1}
	shiftMax = []float64{3, 4, 2}
	// максимальное количество смен, которые разрешено отработать медсестре в неделю
	maxShiftsPerWeek = 5.0

	weeks = 1 // количество недель в рассписание

	shiftPerDay   = len(shiftMin)   // количество смен в день
	shiftsPerWeek = 7 * shiftPerDay // количество смен в неделю

	HARD_CONSTRAINT_PENALTY = 10.0 // штраф за нарушение жесткий правил
)

// расчет количество смен в рассписание
func getNumberShiftsShedule() int {
	return len(nursers) * shiftsPerWeek * weeks
}

// Преобразует все расписание в словарь с отдельным расписанием для каждой медсестры
func getNurseShifts(schedule []float64) map[string][]float64 {
	shiftsPerNurse := getNumberShiftsShedule() / len(nursers)
	nurseShiftsDict := make(map[string][]float64)
	shiftIndex := 0
	for _, nurse := range nursers {
		nurseShiftsDict[nurse] = schedule[shiftIndex : shiftIndex+shiftsPerNurse]
		shiftIndex += shiftsPerNurse
	}
	return nurseShiftsDict
}

// Рассчитывает общую стоимость различных нарушений в заданном графике.
func getCost(schedule []float64) float64 {
	// convert entire schedule into a dictionary with a separate schedule for each nurse:
	nurseShiftsDict := getNurseShifts(schedule)

	// count the various violations:
	consecutiveShiftViolations := countConsecutiveShiftViolations(nurseShiftsDict)
	_, shiftsPerWeekViolations := countShiftsPerWeekViolations(nurseShiftsDict)
	_, nursesPerShiftViolations := countNursesPerShiftViolations(nurseShiftsDict)
	shiftPreferenceViolations := countShiftPreferenceViolations(nurseShiftsDict)

	// жестике нарушения
	hardContstraintViolations := consecutiveShiftViolations + nursesPerShiftViolations + shiftsPerWeekViolations
	// допустимые нарушения
	softContstraintViolations := shiftPreferenceViolations

	return HARD_CONSTRAINT_PENALTY*hardContstraintViolations + softContstraintViolations
}

// Подсчитывает последовательные нарушения смен в расписании
func countConsecutiveShiftViolations(nurseShiftsDict map[string][]float64) float64 {
	violations := 0.0
	for _, nurseShifts := range nurseShiftsDict {
		for i := 0; i < len(nurseShifts)-1; i++ {
			if nurseShifts[i] == 1 && nurseShifts[i+1] == 1 {
				violations += 1
			}
		}
	}
	return violations
}

// Подсчитывает нарушения максимального количества смен в неделю в расписании
func countShiftsPerWeekViolations(nurseShiftsDict map[string][]float64) ([]float64, float64) {
	weeklyShiftsList := []float64{}
	violations := 0.0
	for _, nurseShifts := range nurseShiftsDict {
		for i := 0; i < weeks*shiftsPerWeek; i += shiftsPerWeek {
			// count all the '1's over the week:
			weeklyShifts := sum(nurseShifts[i : i+shiftsPerWeek])
			weeklyShiftsList = append(weeklyShiftsList, weeklyShifts)
			if weeklyShifts > maxShiftsPerWeek {
				violations += weeklyShifts - maxShiftsPerWeek
			}
		}
	}
	return weeklyShiftsList, violations
}

// Подсчитывает нарушения количества медсестер в смену в расписании
func countNursesPerShiftViolations(nurseShiftsDict map[string][]float64) ([]float64, float64) {
	totalPerShiftList := make([]float64, shiftsPerWeek)
	violations := 0.0
	for _, nurseShifts := range nurseShiftsDict {
		for i, v := range nurseShifts {
			if v > 0.0 {
				totalPerShiftList[i]++
			}
		}
	}

	for i := range totalPerShiftList {
		dailyShiftIndex := i % shiftPerDay
		switch {
		case totalPerShiftList[i] > shiftMax[dailyShiftIndex]: // нарушения максимального
			violations += totalPerShiftList[i] - shiftMax[dailyShiftIndex]
		case totalPerShiftList[i] < shiftMin[dailyShiftIndex]: // нарушени минимального числа
			violations += shiftMin[dailyShiftIndex] - totalPerShiftList[i]
		}
	}

	return totalPerShiftList, violations
}

// Считает нарушения предпочтений медсестры в расписании
func countShiftPreferenceViolations(nurseShiftsDict map[string][]float64) float64 {
	violations := 0.0

	for nurse, nurseShifts := range nurseShiftsDict {
		shiftPreference := shiftPreferences[nurse]
		for i, shift := range nurseShifts {
			pref := shiftPreference[i%len(shiftPreference)]
			if shift == 1 && pref == 0 {
				violations += 1
			}
		}
	}
	return violations
}

// печать детально расписания
func printScheduleInfo(schedule []float64) {
	nurseShiftsDict := getNurseShifts(schedule)

	fmt.Println("Графики медсестер:")
	for nurse := range nurseShiftsDict {
		fmt.Printf("%s:%+v\n", nurse, nurseShiftsDict[nurse])
	}

	fmt.Printf("нарушений ограничения на две смены подряд = %v\n\n", countConsecutiveShiftViolations(nurseShiftsDict))

	weeklyShiftsList, violations := countShiftsPerWeekViolations(nurseShiftsDict)
	fmt.Printf("смен в неделю = %v\n", weeklyShiftsList)
	fmt.Printf("нарушений ограничения на число смен в неделю = %v\n\n", violations)

	totalPerShiftList, violations := countNursesPerShiftViolations(nurseShiftsDict)
	fmt.Printf("сестер в смене = %v\n", totalPerShiftList)
	fmt.Printf("нарушений ограничения на число сестер в смене = %v\n\n", violations)

	shiftPreferenceViolations := countShiftPreferenceViolations(nurseShiftsDict)
	fmt.Printf("нарушений пожеланий сестер = %v\n\n", shiftPreferenceViolations)
}

func sum(arr []float64) float64 {
	s := 0.0
	for _, v := range arr {
		s += v
	}
	return s
}
