export interface Delegate {
  id: number;
  score_sheet_id: number;
  name: string;
}

export interface Parameter {
  id: number;
  score_sheet_id: number;
  name: string;
  rule_type: string;
  is_special_parameter: boolean;
  special_scores_rule?: string;
  special_length_rule?: string;
  score_weight: number;
  length_weight: number;
}

export interface Score {
  id: number;
  delegate_id: number;
  parameter_id: number;
  value: number;
  note?: string;
}

export interface ScoreSheet {
  id: number;
  name: string;
  committee_name: string;
  created_at: string;
  updated_at: string;
  created_by: number;
  chair: string;
  vice_chair?: string;
  rapporteur?: string;
}