import { useSelector } from "react-redux";
import { Button } from "./ui/button";
import type { RootState } from "@/store/store";
import { delegateSheetsData, mainSheetData } from "@/utils/excelUtils";
import { exportExcel } from "@/lib/excel/exportExcel";

const ExportButton = () => {
  const scores = useSelector((state: RootState) => state.scores.scores);
  const delegates = useSelector(
    (state: RootState) => state.sheetDetails.delegates
  );
  const parameters = useSelector(
    (state: RootState) => state.sheetDetails.parameters
  );
  const scoresFromDetails = useSelector(
    (state: RootState) => state.sheetDetails.scores
  );
  const scoreSheetName = useSelector(
    (state: RootState) => state.sheetDetails.score_sheet?.name
  );
  const scoreSheetCommitteeName = useSelector(
    (state: RootState) => state.sheetDetails.score_sheet?.committee_name
  );

  const handleClick = async () => {
    const mainSheet = mainSheetData(scores, delegates, parameters);
    const otherSheets = delegateSheetsData(
      scoresFromDetails,
      delegates,
      parameters
    );

    await exportExcel(
      mainSheet,
      otherSheets,
      `${scoreSheetName || "scores"}${
        scoreSheetCommitteeName ? `_${scoreSheetCommitteeName}` : ""
      }_scoresheet.xlsx`
    );
  };

  return (
    <Button type="button" className="h-8 bg-purple-500" onClick={handleClick}>
      Export
    </Button>
  );
};

export default ExportButton;
