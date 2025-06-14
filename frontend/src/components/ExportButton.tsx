import { useSelector } from "react-redux";
import { Button } from "./ui/button";
import type { RootState } from "@/store/store";
import { delegateSheetsData, mainSheetData } from "@/utils/excelUtils";
import { exportExcel } from "@/lib/excel/exportExcel";
import { prepareLLMData } from "@/utils/llmUtils";
import { useState } from "react";
import api from "@/api/axios";
import { toast } from "react-toastify";

const ExportButton = () => {
  const [loading, setLoading] = useState(false);

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

  const user_id = useSelector((state: RootState) => state.auth.user_id);

  const handleClick = async () => {
    const mainSheet = mainSheetData(scores, delegates, parameters);
    const otherSheets = delegateSheetsData(
      scoresFromDetails,
      delegates,
      parameters
    );

    // await exportExcel(
    //   mainSheet,
    //   otherSheets,
    //   `${scoreSheetName || "scores"}${
    //     scoreSheetCommitteeName ? `_${scoreSheetCommitteeName}` : ""
    //   }_scoresheet.xlsx`
    // );

    const llmData = prepareLLMData(scoresFromDetails, parameters, delegates);

    try {
      setLoading(true);
      const uri = "/get_feedback_by_llm";

      const payload = {
        user_id,
        delegates: llmData
      }

      const response = await api.post(uri, payload);
      if (response.status === 200) {
        toast.success("Feedback generated successfully!");
        console.log("Data: ", response.data);
      } else {
        throw new Error("Unexpected response from server");
      }
    } catch (err: any) {
      if (err.response) {
        if (err.response.data?.message) {
          toast.error(err.response.data.message);
        } else {
          toast.error("Failed to generate feedback. Please try again.");
        }
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <Button
      type="button"
      className="h-8 bg-purple-500"
      onClick={handleClick}
      disabled={loading}>
      {loading ? "Exporting..." : "Export"}
    </Button>
  );
};

export default ExportButton;
