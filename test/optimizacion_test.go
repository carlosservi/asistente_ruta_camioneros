package test

import (
	"Asistente_Ruta_Camioneros/src/route"
	"testing"
	"time"
)

func TestCalculateOptimalRoute(t *testing.T) {
	filePathRoute := "../data/route.json"
	routeData, _ := route.ReadJsonRoutes(filePathRoute)
	filePathAreas := "../data/rest_area.json"
	restAreas, _ := route.ReadJsonRestAreas(filePathAreas)
	filePathDescansos := "../data/descansos.json"
	descansos, _ := route.ReadJsonDescansos(filePathDescansos)
	t.Run("Cuando la ruta no necesita descansos", func(t *testing.T) {
		arrivalTime := time.Date(2023, 11, 20, 18, 0, 0, 0, time.UTC)
		result, _, _ := route.CalculateOptimalRoute(arrivalTime, routeData[3], restAreas, descansos)
		expected := []string{"No es necesario parar"}
		assertEqual(t, result[0], expected[0], "Test failed")
	})
	t.Run("Cuando la ruta solo necesita un descanso", func(t *testing.T) {
		arrivalTime := time.Date(2023, 11, 20, 18, 0, 0, 0, time.UTC)
		result, _, _ := route.CalculateOptimalRoute(arrivalTime, routeData[0], restAreas, descansos)
		expected := []string{"Jaen"}
		assertEqual(t, result[0], expected[0], "Test failed")
	})

	t.Run("Cuando la ruta necesita mínimo dos descansos, uno corto y uno largo", func(t *testing.T) {
		arrivalTime := time.Date(2023, 11, 20, 21, 0, 0, 0, time.UTC)
		result, _, _ := route.CalculateOptimalRoute(arrivalTime, routeData[1], restAreas, descansos)
		expected := []string{"Toledo", "Madrid", "Zaragoza"}
		assertEqual(t, result[0], expected[0], "Test failed")
		assertEqual(t, result[1], expected[1], "Test failed")
		assertEqual(t, result[2], expected[2], "Test failed")
	})

	t.Run("Cuando la ruta no se puede hacer porque la distancia entre áreas de servicio es mayor que el tiempo de conducción permitido", func(t *testing.T) {
		arrivalTime := time.Date(2023, 11, 20, 21, 0, 0, 0, time.UTC)
		result, _, _ := route.CalculateOptimalRoute(arrivalTime, routeData[2], restAreas, descansos)
		expected := []string{"No se puede hacer la ruta"}
		assertEqual(t, result[0], expected[0], "Test failed")
	})

	t.Run("Cuando la ruta no se puede hacer porque el área está cerrada y excede el tiempo de conducción de un solo conductor", func(t *testing.T) {
		arrivalTime := time.Date(2023, 11, 21, 22, 0, 0, 0, time.UTC)
		result, _, _ := route.CalculateOptimalRoute(arrivalTime, routeData[1], restAreas, descansos)
		expected := []string{"No se puede hacer la ruta"}
		assertEqual(t, result[0], expected[0], "Test failed")
	})
}

func assertEqual(t *testing.T, result, expected interface{}, message string) {
	if result != expected {
		t.Errorf("%s\nGot: %v\nExpected: %v", message, result, expected)
	}
}
