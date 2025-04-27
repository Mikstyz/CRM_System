package Test

import "log"

func TestAll() {
	var Ok, Bad int

	// Запускаем методы
	log.Println("[INFO] Запуск теста EmployersALL...")
	ok, bad := TestEmployersALL()
	Ok += ok
	Bad += bad
	log.Printf("[SUCCESS] Тест EmployersALL: %d успешных, %d неудачных", ok, bad)

	log.Println("[INFO] Запуск теста GroupALL...")
	ok, bad = TestGroupALL()
	Ok += ok
	Bad += bad
	log.Printf("[SUCCESS] Тест GroupALL: %d успешных, %d неудачных", ok, bad)

	log.Println("[INFO] Запуск теста StudentALL...")
	ok, bad = TestStudentALL()
	Ok += ok
	Bad += bad
	log.Printf("[SUCCESS] Тест StudentALL: %d успешных, %d неудачных", ok, bad)

	log.Println("[INFO] Запуск теста SubjectsALL...")
	ok, bad = TestSubjectsALL()
	Ok += ok
	Bad += bad
	log.Printf("[SUCCESS] Тест SubjectsALL: %d успешных, %d неудачных", ok, bad)

	// Выводим итоги
	log.Println("")
	log.Println("-------------------------------------------------------")
	log.Println("-------------------------------------------------------")
	log.Println("-------------------------------------------------------")
	log.Printf("[SUMMARY] Итоговые результаты: TEST [%d/%d]\nOk: %d\nBad: %d\n", Ok, Ok+Bad, Ok, Bad)
	log.Println("-------------------------------------------------------")
	log.Println("-------------------------------------------------------")
	log.Println("-------------------------------------------------------")
}
