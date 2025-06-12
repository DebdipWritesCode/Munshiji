export interface TableScore {
  score_id: number;
  value: number;
  note?: string;
}

export interface TableParameter {
  parameter_id: number;
  name: string;
  rule_type: string;
  is_special_parameter: boolean;
  special_scores_rule?: string;
  special_length_rule?: string;
  score_weight: number;
  length_weight: number;
  scores: TableScore[];
}

export interface DelegateWithScores {
  delegate_id: number;
  name: string;
  parameters: TableParameter[];
}
