package GoThemePark

import (
	"GolangThemePark/GoThemePark"
	"testing"
)

func TestCalculateRevenue(t *testing.T) {
	revenue := GoThemePark.calculateRevenue(10, 20.5)
	if revenue != 205.0 {
		t.Errorf("Ожидалось 205.0, но получено %.2f", revenue)
	}
}

func TestPromoteAttraction(t *testing.T) {
	attraction := "Колесо обозрения"
	GoThemePark.promoteAttraction(&attraction)
	if attraction != "Колесо обозрения - САМЫЙ ПОПУЛЯРНЫЙ!" {
		t.Errorf("Ожидалось 'Колесо обозрения - САМЫЙ ПОПУЛЯРНЫЙ!', но получено '%s'", attraction)
	}
}

func TestHandleErrors(t *testing.T) {
	GoThemePark.handleErrors(0, 0) // Протестируем с нулём
}
