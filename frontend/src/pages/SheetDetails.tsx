import api from "@/api/axios";
import LogoutDialog from "@/components/auth/LogoutDialog";
import ScoreTable from "@/components/table/ScoreTable";
import ToastComponent from "@/components/ToastComponent";
import { setLoading } from "@/slices/allSheetsSlice";
import {
  setError,
  setParameters,
  setScoreSheet,
  setDelegates,
  setScores,
} from "@/slices/sheetDetailsSlice";
import type { RootState } from "@/store/store";
import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useParams } from "react-router-dom";
import { toast } from "react-toastify";

const SheetDetails = () => {
  const loading = useSelector((state: RootState) => state.sheetDetails.loading);
  const error = useSelector((state: RootState) => state.sheetDetails.error);
  const scoreSheet = useSelector(
    (state: RootState) => state.sheetDetails.score_sheet
  );
  const dispatch = useDispatch();
  const params = useParams<{ id: string }>();
  const score_sheet_id = parseInt(params.id ?? "0", 10);

  if (isNaN(score_sheet_id)) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gradient-to-br from-[#fbfcff] via-[#EDEFFF] to-[#D4E0FF]">
        <h1 className="text-3xl">Invalid score sheet ID</h1>
      </div>
    );
  }

  const fetchScoreSheetDetails = async () => {
    const uri = `/score_sheet_details/${score_sheet_id}`;

    dispatch(setLoading(true));
    dispatch(setError(null));
    try {
      const response = await api.get(uri);

      if (response.status === 200) {
        if (response.data.score_sheet) {
          dispatch(setScoreSheet(response.data.score_sheet));
        }
        if (response.data.parameters) {
          dispatch(setParameters(response.data.parameters));
        }
        if (response.data.delegates) {
          dispatch(setDelegates(response.data.delegates));
        }
        if (response.data.scores) {
          dispatch(setScores(response.data.scores));
        }
      } else {
        toast.error("Unexpected response from server");
        dispatch(setError("Unexpected response from server"));
      }
    } catch (err: any) {
      const message =
        err.response?.data?.message ||
        err.message ||
        "An error occurred while fetching score sheet details.";

      toast.error(message);
      dispatch(setError(message));
    } finally {
      dispatch(setLoading(false));
    }
  };

  useEffect(() => {
    if (score_sheet_id && !isNaN(score_sheet_id) && score_sheet_id > 0) {
      fetchScoreSheetDetails();
    }
  }, [score_sheet_id]);

  return (
    <div className="min-h-screen bg-gradient-to-br from-[#fbfcff] via-[#EDEFFF] to-[#D4E0FF] p-6 flex items-center pt-8 flex-col gap-6">
      <div className="flex w-full justify-between items-center h-10">
        <div className="w-[45%]"></div>
        <h1 className="text-4xl text-blue-900 font-normal font-heading text-center">
          Munshiji
        </h1>
        <LogoutDialog />
      </div>

      {loading && <LoadingState />}
      {error && <ErrorState message={error} />}

      {!loading && !error && scoreSheet && (
        <div className="w-full flex items-center justify-between mt-8 px-10">
          <div className="flex-1 space-y-3">
            <div className="flex items-center gap-3">
              <div className="w-1 h-8 bg-gradient-to-b from-blue-500 to-blue-600 rounded-full"></div>
              <h1 className="text-2xl font-bold text-slate-900 dark:text-slate-100">
                {scoreSheet.name}
              </h1>
            </div>

            <div className="ml-7 space-y-2">
              <div className="flex items-center gap-2">
                <div className="w-2 h-2 rounded-full bg-emerald-500"></div>
                <h2 className="text-lg font-medium text-slate-600 dark:text-slate-400">
                  Committee:{" "}
                  <span className="text-emerald-600 dark:text-emerald-400 font-semibold">
                    {scoreSheet.committee_name}
                  </span>
                </h2>
              </div>

              <div className="flex items-center gap-2">
                <div className="w-2 h-2 rounded-full bg-amber-500"></div>
                <h2 className="text-lg font-medium text-slate-600 dark:text-slate-400">
                  Chairs:{" "}
                  <span className="text-amber-600 dark:text-amber-400 font-semibold">
                    {scoreSheet.chair +
                      (scoreSheet.vice_chair
                        ? `, ${scoreSheet.vice_chair}`
                        : "") +
                      (scoreSheet.rapporteur
                        ? `, ${scoreSheet.rapporteur}`
                        : "")}
                  </span>
                </h2>
              </div>
            </div>
          </div>
          <LogoutDialog />
        </div>
      )}

      {!loading && !error && (
        <ScoreTable />
      )}

      <ToastComponent />
    </div>
  );
};

export default SheetDetails;

const LoadingState = () => (
  <div className="flex items-center justify-center min-h-screen">
    <div className="animate-spin rounded-full h-16 w-16 border-b-2 border-blue-500"></div>
  </div>
);

const ErrorState: React.FC<{ message: string }> = ({ message }) => (
  <div className="flex items-center justify-center min-h-screen">
    <div className="text-red-500 text-lg">
      <p>{message}</p>
    </div>
  </div>
);
