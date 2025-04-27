package Test

import (
	"CRM_System/internal/models"
	"CRM_System/internal/routes"
	"errors"
	"log"
)

// Тест генерации PDF
func Test_GenerateFilledPDF() (int, int, error) {
	var Ok, Bad int

	// Шаг 1: Тестируем успешную генерацию PDF
	log.Println("[INFO] Тестируем генерацию PDF с валидными данными...")
	dataPdf := models.GeneratePDF{
		StudentName:   "Иван Иванов",
		Enterprise:    "ООО Ромашка",
		WorkStartDate: "2025-04-27",
		JobTitle:      "Программист",
		Course:        1,
		Speciality:    "1",
		Groduates:     1,
		Number:        1,
	}

	pdfBytes, err := routes.GenerateFilledPDF(dataPdf)
	if err != nil {
		log.Printf("[ERROR] Ошибка при генерации PDF: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if len(pdfBytes) == 0 {
		log.Printf("[ERROR] PDF должен быть непустым, получено %d байт", len(pdfBytes))
		Bad++
		return Ok, Bad, errors.New("PDF пустой")
	}
	log.Printf("[SUCCESS] PDF успешно сгенерирован, размер: %d байт", len(pdfBytes))
	Ok++

	// Шаг 2: Тестируем обработку ошибки (опционально, если utils.GenerateFilledPDF может вернуть ошибку)
	// Примечание: Без моков сложно симулировать ошибку. Если есть способ, уточни!
	// Здесь можно добавить тест с некорректными данными, если utils.GenerateFilledPDF их проверяет.
	log.Println("[INFO] Тестируем генерацию PDF с некорректными данными...")
	dataPdfInvalid := models.GeneratePDF{
		StudentName:   "", // Пустое имя, предположим, что это вызовет ошибку
		Enterprise:    "ООО Ромашка",
		WorkStartDate: "2025-04-27",
		JobTitle:      "Программист",
	}

	pdfBytes, err = routes.GenerateFilledPDF(dataPdfInvalid)
	if err == nil {
		log.Printf("[ERROR] Ожидалась ошибка при пустом имени, но PDF сгенерирован, размер: %d байт", len(pdfBytes))
		Bad++
		return Ok, Bad, errors.New("ожидалась ошибка при пустом имени")
	}
	log.Printf("[SUCCESS] Ошибка при генерации PDF с пустым именем обработана: %v", err)
	Ok++

	return Ok, Bad, nil
}
