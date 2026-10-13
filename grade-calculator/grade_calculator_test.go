package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	if actual_value := gradeCalculator.GetFinalGrade(); actual_value != "C" {
		t.Errorf("Expected GetGrade to return 'C'; got '%s' instead", actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	if actual_value := gradeCalculator.GetFinalGrade(); actual_value != "D" {
		t.Errorf("Expected GetGrade to return 'D'; got '%s' instead", actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" || Exam.String() != "exam" || Essay.String() != "essay" {
		t.Errorf("Unexpected GradeType names")
	}
}

func TestPassFailPass(t *testing.T) {
	gradeCalculator := NewPassFailGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	if actual_value := gradeCalculator.GetFinalGrade(); actual_value != "Pass" {
		t.Errorf("Expected GetGrade to return 'Pass'; got '%s' instead", actual_value)
	}
}

func TestPassFailFail(t *testing.T) {
	gradeCalculator := NewPassFailGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 69, Assignment)
	gradeCalculator.AddGrade("exam 1", 69, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 69, Essay)

	if actual_value := gradeCalculator.GetFinalGrade(); actual_value != "Fail" {
		t.Errorf("Expected GetGrade to return 'Fail'; got '%s' instead", actual_value)
	}
}
