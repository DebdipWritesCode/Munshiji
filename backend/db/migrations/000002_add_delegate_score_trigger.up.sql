-- Create trigger function
CREATE OR REPLACE FUNCTION insert_default_scores_for_delegate()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO scores (delegate_id, parameter_id, value)
  SELECT NEW.id, p.id, 0.0
  FROM parameters p
  WHERE p.score_sheet_id = NEW.score_sheet_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger
CREATE TRIGGER trg_insert_scores_on_delegate
AFTER INSERT ON delegates
FOR EACH ROW
EXECUTE FUNCTION insert_default_scores_for_delegate();
