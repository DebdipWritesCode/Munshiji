-- Create trigger function
CREATE OR REPLACE FUNCTION insert_default_scores_for_parameter()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO scores (delegate_id, parameter_id, value)
  SELECT d.id, NEW.id, 0.0
  FROM delegates d
  WHERE d.score_sheet_id = NEW.score_sheet_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger
CREATE TRIGGER trg_insert_scores_on_parameter
AFTER INSERT ON parameters
FOR EACH ROW
EXECUTE FUNCTION insert_default_scores_for_parameter();
