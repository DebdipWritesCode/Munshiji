import { useSelector } from "react-redux";
import { Button } from "./ui/button";
import type { RootState } from "@/store/store";
import { delegateSheetsData, mainSheetData } from "@/utils/excelUtils";

const ExportButton = () => {
  const scores = useSelector((state: RootState) => state.scores.scores);
  const delegates = useSelector(
    (state: RootState) => state.sheetDetails.delegates
  );
  const parameters = useSelector(
    (state: RootState) => state.sheetDetails.parameters
  );
  const scoresFromDetails = useSelector((state: RootState) => state.sheetDetails.scores);

  const handleClick = () => {
    const mainSheet = mainSheetData(scores, delegates, parameters);
    const otherSheets = delegateSheetsData(scoresFromDetails, delegates, parameters);

    console.log("Main Sheet Data:", mainSheet);
    console.log("Delegate Sheets Data:", otherSheets);
  };

  return (
    <Button type="button" className="h-8 bg-purple-500" onClick={handleClick}>
      Export
    </Button>
  );
};

export default ExportButton;
