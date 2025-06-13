import type { ScoreState } from "@/slices/scoresSlice";
import type { Delegate } from "@/utils/getDataInterfaces";

export function calculateTotalScores(
  scores: ScoreState[]
): Record<number, number> {
  const totals: Record<number, number> = {};

  for (const score of scores) {
    const { delegate_id, valueToDisplay } = score;

    if (!totals[delegate_id]) {
      totals[delegate_id] = 0;
    }

    totals[delegate_id] += valueToDisplay;
  }

  return totals;
}

export function getSortedTotalScores(
  scores: ScoreState[]
): { delegate_id: number; totalScore: number }[] {
  const totals = calculateTotalScores(scores);

  return Object.entries(totals)
    .map(([delegate_id, totalScore]) => ({
      delegate_id: Number(delegate_id),
      totalScore,
    }))
    .sort((a, b) => b.totalScore - a.totalScore);
}

export function getTopDelegates(
  scores: ScoreState[],
  topN: number
): { delegate_id: number; totalScore: number }[] {
  const sortedScores = getSortedTotalScores(scores);
  return sortedScores.slice(0, topN);
}

export function sortDelegatesByScore(scores: ScoreState[], delegates: Delegate[]): Delegate[] {
  const totalScores = getSortedTotalScores(scores); // Returns: { delegate_id, totalScore }[]

  const delegateMap = new Map(delegates.map((d) => [d.id, d]));

  return totalScores
    .map(({ delegate_id }) => delegateMap.get(delegate_id))
    .filter((d): d is Delegate => d !== undefined); // filters out unmatched delegate_ids, if any
}