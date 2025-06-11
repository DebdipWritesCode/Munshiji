CREATE OR REPLACE FUNCTION update_score_sheet_timestamp()
RETURNS TRIGGER AS $$
DECLARE
  sheet_id INTEGER;  
BEGIN
  IF TG_TABLE_NAME = 'scores' THEN
    IF TG_OP = 'DELETE' THEN
      SELECT p.score_sheet_id INTO sheet_id
      FROM parameters p
      WHERE p.id = OLD.parameter_id;
    ELSE
      SELECT p.score_sheet_id INTO sheet_id
      FROM parameters p
      WHERE p.id = NEW.parameter_id;
    END IF;
  ELSIF TG_TABLE_NAME = 'parameters' THEN
    sheet_id := COALESCE(NEW.score_sheet_id, OLD.score_sheet_id);
  ELSIF TG_TABLE_NAME = 'delegates' THEN
    sheet_id := COALESCE(NEW.score_sheet_id, OLD.score_sheet_id);
  END IF;

  IF sheet_id IS NOT NULL THEN
    UPDATE score_sheets
    SET updated_at = NOW()
    WHERE id = sheet_id;
  END IF;

  RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Triggers
CREATE TRIGGER trg_update_scores
AFTER INSERT OR UPDATE OR DELETE ON scores
FOR EACH ROW
EXECUTE FUNCTION update_score_sheet_timestamp();

CREATE TRIGGER trg_update_parameters
AFTER INSERT OR UPDATE OR DELETE ON parameters
FOR EACH ROW
EXECUTE FUNCTION update_score_sheet_timestamp();

CREATE TRIGGER trg_update_delegates
AFTER INSERT OR UPDATE OR DELETE ON delegates
FOR EACH ROW
EXECUTE FUNCTION update_score_sheet_timestamp();
