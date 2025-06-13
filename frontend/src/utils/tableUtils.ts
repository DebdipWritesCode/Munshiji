import type { Delegate, Parameter, Score } from "./getDataInterfaces";
import type {
  DelegateWithScores,
  TableParameter,
  TableScore,
} from "./getTableInterfaceTypes";
import type { ColumnDef } from "@tanstack/react-table";

function generateAccessorKey(name: string): string {
  return name.toLowerCase().replace(/\s+/g, "_");
}

function getCellRenderer(key: string): (info: any) => string {
  return key === "delegates"
    ? (info) => String(info.getValue())
    : (info) => Number(info.getValue()).toFixed(2);
}

function roundTwoDecimals(value: number): number {
  return Math.round(value * 100) / 100;
}

function calculateTotalForAverageRuleType(scores: TableScore[]) {
  const total = scores.reduce((acc, score) => acc + (score.value || 0), 0);
  return roundTwoDecimals(total / scores.length || 0);
}

function calculateTotalForAbsoluteRuleType(scores: TableScore[]) {
  return roundTwoDecimals(scores.reduce((acc, score) => acc + (score.value || 0), 0));
}

function calculateTotalForSpecialRuleType(
  scores: TableScore[],
  special_scores_rule: string | undefined,
  score_weight: number,
  length_weight: number
) {
  // console.log("Scores:", scores);
  // console.log("Special Scores Rule:", special_scores_rule);
  // console.log("Score Weight:", score_weight);
  // console.log("Length Weight:", length_weight);

  let totalValueScore = 0;
  let totalLengthScore = scores.length;

  if (special_scores_rule === "average") {
    totalValueScore = calculateTotalForAverageRuleType(scores);
  } else if (special_scores_rule === "absolute") {
    totalValueScore = calculateTotalForAbsoluteRuleType(scores);
  }

  return roundTwoDecimals(totalValueScore * score_weight + totalLengthScore * length_weight);
}

export function setTableColumns(parameters: Parameter[]): ColumnDef<any>[] {
  const columnNames = ["Delegates", ...parameters.map((p) => p.name), "Total"];

  return columnNames.map((name) => {
    const key = generateAccessorKey(name);

    return {
      header: name,
      accessorKey: key,
      cell: getCellRenderer(key),
      meta: { key },
    } as ColumnDef<any> & { meta: { key: string } };
  });
}

export function prepareTableData(
  scores: Score[],
  parameters: Parameter[],
  delegates: Delegate[]
): DelegateWithScores[] {
  return delegates.map((delegate) => {
    const delegateScores = scores.filter((s) => s.delegate_id === delegate.id);

    const parameterData: TableParameter[] = parameters.map((param) => {
      const matchingScores = delegateScores
        .filter((s) => s.parameter_id === param.id)
        .map((score) => ({
          score_id: score.id,
          value: score.value,
          note: score.note,
        }));

      return {
        parameter_id: param.id,
        name: param.name,
        rule_type: param.rule_type,
        is_special_parameter: param.is_special_parameter,
        special_scores_rule: param.special_scores_rule,
        special_length_rule: param.special_length_rule,
        score_weight: param.score_weight,
        length_weight: param.length_weight,
        scores: matchingScores,
      };
    });

    return {
      delegate_id: delegate.id,
      name: delegate.name,
      parameters: parameterData,
    };
  });
}

export function calculateRuleWiseTotal(parameterData: TableParameter) {
  // console.log("Parameter Data:", parameterData);
  switch (parameterData.rule_type) {
    case "average":
      return calculateTotalForAverageRuleType(parameterData.scores);
    case "absolute":
      return calculateTotalForAbsoluteRuleType(parameterData.scores);
    case "special":
      return calculateTotalForSpecialRuleType(
        parameterData.scores,
        parameterData.special_scores_rule,
        parameterData.score_weight,
        parameterData.length_weight
      );
    default:
      console.warn(`Unknown rule type: ${parameterData.rule_type}`);
      return 0;
  }
}

export function calculateTotalScore(parameterData: TableParameter[]): number {
  return roundTwoDecimals(parameterData.reduce((total, param) => {
    const ruleWiseTotal = calculateRuleWiseTotal(param);
    return total + ruleWiseTotal;
  }, 0));
}