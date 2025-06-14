import type { Delegate, Parameter, Score } from "./getDataInterfaces";

type PreparedParameterData = {
  parameter_id: number;
  name: string;
  highestScore: number;
};

type PreparedDelegateData = {
  delegate_id: number;
  parameter_id: number;
  delegate_name: string;
  parameter_name: string;
  highestScore: number;
};

type FinalLLMData = {
  delegate_name: string;
  parameters: {
    parameter_name: string;
    received: number; // delegate's highest score
    highest: number; // highest score for this parameter overall
  }[];
};

export function prepareParameterData(
  scores: Score[],
  parameters: Parameter[]
) {
  const parameterMap = new Map<number, PreparedParameterData>();

  parameters.forEach((parameter) => {
    parameterMap.set(parameter.id, {
      parameter_id: parameter.id,
      name: parameter.name,
      highestScore: 0,
    });
  });

  scores.forEach((score) => {
    if (score.value == null) return;

    const existing = parameterMap.get(score.parameter_id);
    if (existing) {
      existing.highestScore = Math.max(
        existing.highestScore,
        score.value
      );
    }
  });

  return parameterMap;
}

export function prepareDelegateData(
  scores: Score[],
  delegates: Delegate[],
  parameters: Parameter[]
) {
  const resultMap = new Map<string, PreparedDelegateData>();

  const paramMap = new Map(parameters.map((p) => [p.id, p.name]));
  const delegateMap = new Map(delegates.map((d) => [d.id, d.name]));

  for (const score of scores) {
    if (score.value == null) continue;

    const key = `${score.delegate_id}-${score.parameter_id}`;
    const existing = resultMap.get(key);

    if (!existing) {
      resultMap.set(key, {
        delegate_id: score.delegate_id,
        parameter_id: score.parameter_id,
        delegate_name: delegateMap.get(score.delegate_id) ?? "Unknown",
        parameter_name: paramMap.get(score.parameter_id) ?? "Unknown",
        highestScore: score.value,
      });
    } else {
      existing.highestScore = Math.max(
        existing.highestScore,
        score.value
      );
    }
  }

  return resultMap;
}

export function prepareLLMData(
  scores: Score[],
  parameters: Parameter[],
  delegates: Delegate[]
) {
  const parameterMap = prepareParameterData(scores, parameters);
  const delegateMap = prepareDelegateData(scores, delegates, parameters);

  const delegateGroupMap = new Map<string, FinalLLMData>();

  for (const delegateData of delegateMap.values()) {
    const { delegate_name, parameter_name, parameter_id, highestScore } =
      delegateData;

    if (!delegateGroupMap.has(delegate_name)) {
      delegateGroupMap.set(delegate_name, {
        delegate_name,
        parameters: [],
      });
    }

    const globalParametersData = parameterMap.get(parameter_id);
    const highest = globalParametersData
      ? globalParametersData.highestScore
      : 0;

    delegateGroupMap.get(delegate_name)!.parameters.push({
      parameter_name,
      received: highestScore,
      highest,
    });
  }

  return Array.from(delegateGroupMap.values());
}
