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
	gradeCalculator.AddGrade("exam 1", 40, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 58, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
    expected_value := "C"

    gradeCalculator := NewGradeCalculator()

    gradeCalculator.AddGrade("assignment 1", 70, Assignment)
    gradeCalculator.AddGrade("exam 1", 78, Exam)
    gradeCalculator.AddGrade("essay 1", 73, Essay)

    actual_value := gradeCalculator.GetFinalGrade()

    if expected_value != actual_value {
        t.Errorf("Expected C; got %s instead", actual_value)
    }
}

func TestGetGradeD(t *testing.T) {
    expected_value := "D"

    gradeCalculator := NewGradeCalculator()

    gradeCalculator.AddGrade("assignment 1", 60, Assignment)
    gradeCalculator.AddGrade("exam 1", 68, Exam)
    gradeCalculator.AddGrade("essay 1", 63, Essay)

    actual_value := gradeCalculator.GetFinalGrade()

    if expected_value != actual_value {
        t.Errorf("Expected D; got %s instead", actual_value)
    }
}

func TestGradeTypeString(t *testing.T) {
    expected_value := "assignment"
    actual_value := Assignment.String()

    if expected_value != actual_value {
        t.Errorf("Expected %s; got %s instead", expected_value, actual_value)
    }
}
