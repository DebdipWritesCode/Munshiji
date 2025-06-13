import type { ScoreState } from "@/slices/scoresSlice";
import { getSortedTotalScores } from "./scoresUtils";
import type { Delegate, Parameter, Score } from "./getDataInterfaces";
import { sortById } from "./tableUtils";

export function mainSheetData(
  scores: ScoreState[],
  delegates: Delegate[],
  parameters: Parameter[]
): Record<string, string | number>[] {
  const totals = getSortedTotalScores(scores);

  const result: Record<string, string | number>[] = [];

  for (const item of totals) {
    const { delegate_id, totalScore } = item;
    const delegate = delegates.find((d) => d.id === delegate_id);
    if (!delegate) continue;

    const row: Record<string, string | number> = {
      Name: delegate.name,
    };

    for (const param of parameters) {
      const score = scores.find(
        (s) => s.delegate_id == delegate_id && s.parameter_id === param.id
      );
      row[param.name] = score ? score.valueToDisplay : "";
    }

    row["Total"] = totalScore;

    result.push(row);
  }
  return result;
}

export function delegateSheetsData(
  scores: Score[],
  delegates: Delegate[],
  parameters: Parameter[]
): Record<string, { parameter: string; value: number; note?: string }[]> {
  const sortedScores = sortById(scores);

  const sheets: Record<string, { parameter: string; value: number; note?: string }[]> = {};

  for (const delegate of delegates) {
    const delegateScores = sortedScores.filter(
      (score) => score.delegate_id === delegate.id
    );

    const parameterMap: Record<
      number,
      { value: number; note?: string }[]
    > = {};

    for (const score of delegateScores) {
      if (!parameterMap[score.parameter_id]) {
        parameterMap[score.parameter_id] = [];
      }
      parameterMap[score.parameter_id].push({
        value: score.value,
        note: score.note,
      });
    }

    const sheetData: { parameter: string; value: number; note?: string }[] = [];

    for (const [paramIdStr, scoresArr] of Object.entries(parameterMap)) {
      const paramId = parseInt(paramIdStr);
      const param = parameters.find((p) => p.id === paramId);
      const paramName = param ? param.name : `Param ${paramId}`;

      for (const s of scoresArr) {
        sheetData.push({
          parameter: paramName,
          value: s.value,
          note: s.note,
        });
      }
    }

    sheets[delegate.name] = sheetData;
  }

  return sheets;
}
