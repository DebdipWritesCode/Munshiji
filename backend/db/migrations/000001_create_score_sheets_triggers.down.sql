-- Drop triggers
DROP TRIGGER IF EXISTS trg_update_scores ON scores;
DROP TRIGGER IF EXISTS trg_update_parameters ON parameters;
DROP TRIGGER IF EXISTS trg_update_delegates ON delegates;

-- Drop the trigger function
DROP FUNCTION IF EXISTS update_score_sheet_timestamp();
