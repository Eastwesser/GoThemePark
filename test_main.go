package GoThemePark

import (
	"testing"
)

func TestCalculateRevenue(t *testing.T) {
	revenue := calculateRevenue(10, 20.5)
	if revenue != 205.0 {
		t.Errorf("Ожидалось 205.0, но получено %.2f", revenue)
	}
}

func TestPromoteAttraction(t *testing.T) {
	attraction := "Колесо обозрения"
	promoteAttraction(&attraction)
	if attraction != "Колесо обозрения - САМЫЙ ПОПУЛЯРНЫЙ!" {
		t.Errorf("Ожидалось 'Колесо обозрения - САМЫЙ ПОПУЛЯРНЫЙ!', но получено '%s'", attraction)
	}
}

func TestCheckCapacity(t *testing.T) {
	err := checkCapacity(1200)
	if err == nil {
		t.Error("Ожидалась ошибка при превышении вместимости, но её не было")
	}
}
