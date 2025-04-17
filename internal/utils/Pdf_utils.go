package utils

import (
	"CRM_System/internal/modeles"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
)

// Генерация PDF
var FontDir string

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Не удалось определить рабочую директорию: %v", err)
	}

	FontDir = filepath.Join(wd, "internal", "Config", "ttf")
	log.Printf("Рабочая директория для шрифтов: %s", FontDir)
}

func GenerateFilledPDF(Data modeles.PdfDoc) ([]byte, error) {
	log.Println("Формирование PDF документа")

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Указываем путь к шрифту
	pdf.SetFontLocation(FontDir)

	// Добавляем шрифт с поддержкой UTF-8
	pdf.AddUTF8Font("DejaVu", "", "DejaVuSans.ttf")

	// Проверяем, что шрифт существует
	fontPath := filepath.Join(FontDir, "DejaVuSans.ttf")
	if _, err := os.Stat(fontPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("шрифт не найден по пути: %s", fontPath)
	}

	// Устанавливаем шрифт
	pdf.SetFont("DejaVu", "", 12)

	// Текст с поддержкой кириллицы
	pdf.CellFormat(190, 10, "БЛАНК РАБОТОДАТЕЛЯ", "", 1, "C", false, 0, "")
	pdf.MultiCell(190, 7, "Справка о результатах освоения профессиональной образовательной программы по индивидуальному учебному плану\nна базе трудоустройства обучающегося", "", "C", false)
	pdf.Ln(5)

	pdf.MultiCell(190, 7, fmt.Sprintf("Сотрудник %s работает на предприятии %s с %s по настоящее время в должности %s", Data.Name, Data.Enterprise, Data.WorkStartDate, Data.JobTitle), "", "L", false)
	pdf.Ln(5)

	pdf.Cell(0, 7, "за указанный период выполнил следующие виды работ:")
	pdf.Ln(10)

	pdf.SetFontSize(10)
	pdf.SetFillColor(200, 200, 200)
	pdf.CellFormat(15, 10, "№ п/п", "1", 0, "C", true, 0, "")
	pdf.CellFormat(120, 10, "Вид работ", "1", 0, "C", true, 0, "")
	pdf.CellFormat(55, 10, "Оценка (оценка прописью)", "1", 0, "C", true, 0, "")
	pdf.Ln(10)

	pdf.SetFillColor(255, 255, 255)
	for i, subject := range Data.SubjectArray {
		pdf.CellFormat(15, 7, fmt.Sprintf("%d.", i+1), "1", 0, "R", false, 0, "")
		pdf.CellFormat(120, 7, subject, "1", 0, "L", false, 0, "")
		pdf.CellFormat(55, 7, "", "1", 0, "C", false, 0, "")
		pdf.Ln(7)
	}

	pdf.Ln(10)
	pdf.SetFontSize(12)
	pdf.MultiCell(190, 7, "Предприятие предлагает зачесть вышеуказанные виды работ как успешно освоенные на рабочем месте.", "", "L", false)
	pdf.Ln(5)
	pdf.Cell(0, 7, "Должность ответственного лица:")
	pdf.Ln(20)
	pdf.Cell(0, 7, "____________________________________________________________________________________________________________")
	pdf.Ln(30)
	pdf.Cell(0, 7, "«___» _____________ 2025 г.")
	pdf.Ln(0)
	pdf.CellFormat(190, 7, "_______________/__________________", "", 0, "R", false, 0, "")
	pdf.Ln(5)
	pdf.CellFormat(190, 7, "подпись-расшифровка подписи", "", 0, "R", false, 0, "")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		log.Printf("Ошибка при формировании документа: %v", err)
		return nil, err
	}

	fileName := fmt.Sprintf("Бланк работодателя для студента %s.PDF", Data.Name)
	err = SavePDFToFile(buf.Bytes(), fileName)
	if err != nil {
		fmt.Printf("Ошибка при сохранении файла: %v", err)
	}

	return buf.Bytes(), nil
}

func SavePDFToFile(pdfBytes []byte, fileName string) error {
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Не удалось получить рабочую директорию: %v", err)
		return err
	}

	savePath := filepath.Join(wd, "Data", "Excel", fileName)

	dir := filepath.Dir(savePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Printf("Не удалось создать директорию %s: %v", dir, err)
			return err
		}
	}

	err = os.WriteFile(savePath, pdfBytes, 0644)
	if err != nil {
		log.Printf("Ошибка при сохранении PDF: %v", err)
		return err
	}

	log.Printf("PDF успешно сохранён как %s", savePath)
	return nil
}
