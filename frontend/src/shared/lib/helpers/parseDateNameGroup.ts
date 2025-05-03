import { DateNameGroup } from "@/entities/group/types";

type ParseSuccess = { status: "fulfilled"; dateNameGroup: DateNameGroup };
type ParseError = { status: "rejected"; errors: string[] };
type ParseResult = ParseSuccess | ParseError;

const NAME_RE = /^([1-4])(\p{L}{2,})(9|11)-(\d+)$/u;

/**
 * Парсит строку формата `${course}${specialty}${graduates}-${groupNumber}`
 * и возвращает объект DateNameGroup.
 *
 * При некорректном формате бросает исключение – легче отлавливать ошибку при
 * использовании (можно заменить на `null`, если требуется «тихий» возврат).
 */
export function parseDateNameGroup(name: string): ParseResult {
  const match = NAME_RE.exec(name.trim());

  if (match) {
    const [, course, specialty, graduates, groupNumber] = match;
    return {
      status: "fulfilled",
      dateNameGroup: {
        course: course as DateNameGroup["course"],
        specialty,
        graduates: graduates as DateNameGroup["graduates"],
        groupNumber: Number(groupNumber),
      },
    };
  }
  console.log("match", match);
  console.log("name", name);

  // TODO Добвыть обработку ошибок: - неизменять, если ошибка, - показавать ошибки.

  const errors: string[] = [];
  // 1. курс
  if (!/^[1-4]/.test(name)) {
    errors.push("course должен быть числом 1-4 и стоять первым символом");
  }

  // 2. наличие «-»
  if (!/-/.test(name)) {
    errors.push("отсутствует дефис «-» перед номером группы");
  }

  // 3. выпускники
  const gradsMatch = name.match(/(9|11)(?=-\d+$)/);
  if (!gradsMatch) {
    errors.push(
      "graduates должен быть «9» или «11» непосредственно перед дефисом",
    );
  }

  // 4. specialty
  const specMatch = name.match(/^[1-4]([A-Za-z]{2,})/);
  if (!specMatch) {
    errors.push("specialty должен состоять минимум из двух букв после course");
  }

  // 5. groupNumber
  if (!/-\d+$/.test(name)) {
    errors.push("groupNumber должен быть положительным числом после дефиса");
  }

  return { status: "rejected", errors };
}
