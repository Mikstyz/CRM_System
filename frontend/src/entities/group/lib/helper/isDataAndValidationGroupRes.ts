import { Course, Graduates, Group } from "@/entities/group/types";
import { models } from "@wails/go/models.ts";

type InFGroupAndSubject = models.InFGroupAndSubject;

export const convResDataInGroups = (
  ResGroupsData: InFGroupAndSubject[],
): Group[] => {
  const convGroups =
    (ResGroupsData.map((group) => {
      if (isValidResGroupData(group)) {
        return {
          id: group.Id,
          name: `${group.Course}${group.Spesiality}${group.Groduates}-${group.Number}`,
          dateNameGroup: {
            course: String(group.Course) as Course,
            specialty: group.Spesiality,
            graduates: String(group.Groduates) as Graduates,
            groupNumber: group.Number,
          },
          disciplines: {
            "1": group.Subject.OneSemester,
            "2": group.Subject.TwoSemester,
          },
        };
      }
    }).filter((group) => group !== undefined) as Group[]) || undefined;
  return convGroups || [];
};

function isValidResGroupData(group: InFGroupAndSubject) {
  return (
    group.Id &&
    group.Course &&
    group.Spesiality &&
    group.Groduates &&
    group.Number &&
    group.Subject &&
    group.Subject.OneSemester &&
    group.Subject.TwoSemester &&
    (group.Course == 1 ||
      group.Course == 2 ||
      group.Course == 3 ||
      group.Course == 4) &&
    (group.Groduates == 9 || group.Groduates == 11) &&
    group.Spesiality.length >= 2 &&
    group.Number >= 1
  );
}
