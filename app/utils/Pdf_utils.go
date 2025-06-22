package utils

import (
	"CRM_System/app/api/routes"
	"CRM_System/app/storage/models"
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
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("[utils][pdf] - Не удалось определить рабочую директорию: %v", err)
	}

	exeDir := filepath.Dir(exePath)

	FontDir = filepath.Join(exeDir, "Data", "Fonts")

	log.Printf("[utils][pdf] - Рабочая директория для шрифтов: %s", FontDir)
}

func GenerateFiledPDF(Data models.GeneratePDF) ([]byte, string, error) {
	log.Println("[utils][pdf] - Формирование PDF документа")

	log.Println("[utils][pdf] - Получение предметов группы")
	SubjectArray, Error := routes.Inf_DisciplinesByGroupId(Data.GroupId, Data.Semester)

	if Error != nil {
		log.Fatal("[utils][pdf] - Ошибка при получении предметов группы")
		return nil, "", Error
	}

	log.Printf("[utils][pdf] - Получено предметов группы %d", len(SubjectArray))

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Указываем путь к шрифту
	pdf.SetFontLocation(FontDir)

	// Добавляем шрифт с поддержкой UTF-8
	pdf.AddUTF8Font("DejaVu", "", "DejaVuSans.ttf")

	// Проверяем, что шрифт существует
	fontPath := filepath.Join(FontDir, "DejaVuSans.ttf")
	if _, err := os.Stat(fontPath); os.IsNotExist(err) {
		return nil, "", fmt.Errorf("[utils][pdf] - шрифт не найден по пути: %s", fontPath)
	}

	// Устанавливаем шрифт
	pdf.SetFont("DejaVu", "", 12)

	// Текст с поддержкой кириллицы
	pdf.CellFormat(190, 10, "БЛАНК РАБОТОДАТЕЛЯ", "", 1, "C", false, 0, "")
	pdf.MultiCell(190, 7, "Справка о результатах освоения профессиональной образовательной программы по индивидуальному учебному плану\nна базе трудоустройства обучающегося", "", "C", false)
	pdf.Ln(5)

	pdf.MultiCell(190, 7, fmt.Sprintf("Сотрудник %s работает на предприятии %s с %s по настоящее время в должности %s", Data.StudentName, Data.Enterprise, Data.WorkStartDate, Data.JobTitle), "", "L", false)
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

	// for i := 0; i < 1; i++ {
	// 	pdf.CellFormat(15, 7, fmt.Sprintf("%d.", i+1), "1", 0, "R", false, 0, "")
	// 	pdf.CellFormat(120, 7, "_", "1", 0, "L", false, 0, "")
	// 	pdf.CellFormat(55, 7, "", "1", 0, "C", false, 0, "")
	// 	pdf.Ln(7)
	// }

	for i, subject := range SubjectArray {
		pdf.CellFormat(15, 7, fmt.Sprintf("%d.", i+1), "1", 0, "R", false, 0, "")
		pdf.CellFormat(120, 7, subject, "1", 0, "L", false, 0, "")
		pdf.CellFormat(55, 7, "", "1", 0, "C", false, 0, "")
		pdf.Ln(7)
	}

	pdf.Ln(10)
	pdf.SetFontSize(12)
	pdf.MultiCell(190, 7, "Предприятие предлагает зачесть вышеуказанные виды работ как успешно освоенные на рабочем месте.", "", "L", false)
	pdf.Ln(15)
	pdf.Cell(0, 7, "Должность ответственного лица:")
	pdf.Ln(10)
	pdf.Cell(0, 7, "____________________________________________________________________________")
	pdf.Ln(30)
	pdf.Cell(0, 7, "«___» _____________ 2025 г.")
	pdf.Ln(0)
	pdf.CellFormat(190, 7, "_______________/__________________", "", 0, "R", false, 0, "")
	pdf.Ln(5)
	pdf.CellFormat(190, 7, "подпись-расшифровка подписи", "", 0, "R", false, 0, "")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		log.Printf("[utils][pdf] - Ошибка при формировании документа: %v", err)
		return nil, "", err
	}

	fileName := fmt.Sprintf("[utils][pdf] - Бланк работодателя для студента %s.PDF", Data.StudentName)
	path, err := SavePDFToFile(buf.Bytes(), fileName)
	if err != nil {
		fmt.Printf("[utils][pdf] - Ошибка при сохранении файла: %v", err)
		return nil, "", err
	}

	return buf.Bytes(), path, nil

}

func SavePDFToFile(pdfBytes []byte, fileName string) (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		log.Printf("[utils][pdf] - Не удалось получить путь к исполняемому файлу: %v", err)
		return "", err
	}

	baseDir := filepath.Dir(execPath)
	savePath := filepath.Join(baseDir, "Data", "Pdf_Document", fileName)

	if err := os.MkdirAll(filepath.Dir(savePath), 0755); err != nil {
		log.Printf("[utils][pdf] - Не удалось создать директорию %s: %v", filepath.Dir(savePath), err)
		return "", err
	}

	if err := os.WriteFile(savePath, pdfBytes, 0644); err != nil {
		log.Printf("[utils][pdf] - Ошибка при сохранении PDF: %v", err)
		return "", err
	}

	log.Printf("[utils][pdf] - PDF успешно сохранён как %s", savePath)
	return savePath, nil
}
