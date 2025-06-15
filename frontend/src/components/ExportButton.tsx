import { useSelector } from "react-redux";
import { Button } from "./ui/button";
import type { RootState } from "@/store/store";
import { delegateSheetsData, mainSheetData } from "@/utils/excelUtils";
import { exportExcel } from "@/lib/excel/exportExcel";
import { prepareLLMData } from "@/utils/llmUtils";
import { useState } from "react";
import api from "@/api/axios";
import { toast } from "react-toastify";
import DelegateReport from "./DelegateReport";
import { getDelegateScoresByName } from "@/utils/scoresUtils";
import { generateAndZipPDFsWithExcel } from "@/utils/generateAndZipPDFs";

const ExportButton = () => {
  const [loading, setLoading] = useState(false);
  const [feedbackData, setFeedbackData] = useState<any>(null);

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

    const llmData = prepareLLMData(scoresFromDetails, parameters, delegates);

    try {
      setLoading(true);
      const uri = "/get_feedback_by_llm";

      const payload = {
        user_id,
        delegates: llmData,
      };

      const response = await api.post(uri, payload);
      if (response.status === 200) {
        toast.success("Feedback generated successfully!");
        setFeedbackData(response.data.feedbacks);

        const excelBlob = await exportExcel(
          mainSheet,
          otherSheets,
          `${scoreSheetName || "scores"}${
            scoreSheetCommitteeName ? `_${scoreSheetCommitteeName}` : ""
          }_scoresheet.xlsx`,
          true
        );

        const delegateNames = response.data.feedbacks.map((f: any) => f.delegate_name);
        
        await generateAndZipPDFsWithExcel(delegateNames, excelBlob);

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
    <>
      <Button
        type="button"
        className="h-8 bg-purple-500"
        onClick={handleClick}
        disabled={loading}>
        {loading ? "Exporting..." : "Export"}
      </Button>

      {feedbackData && (
        <div className="absolute left-[-9999px] top-0">
          {feedbackData.map((entry: any, index: number) => (
            <DelegateReport
              key={index}
              delegateName={entry.delegate_name}
              feedback={entry.feedback_text}
              data={getDelegateScoresByName(
                scores,
                entry.delegate_name,
                delegates,
                parameters
              )}
            />
          ))}
        </div>
      )}
    </>
  );
};

export default ExportButton;
