import api from "@/api/axios";
import CreateSheetDialog from "@/components/CreateSheetDialog";
import SheetCard from "@/components/SheetCard";
import ToastComponent from "@/components/ToastComponent";
import { setAllSheets, setLoading } from "@/slices/allSheetsSlice";
import type { RootState } from "@/store/store";
import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { toast } from "react-toastify";

interface ScoreSheetState {
  id: number | null;
  name: string | null;
  committee_name: string | null;
  chair: string | null;
  vice_chair?: string | null;
  rapporteur?: string | null;
  created_by: number | null;
  created_at: string | null;
  updated_at: string | null;
}

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

const SeeSheets = () => {
  const [error, setError] = useState<string | null>(null);
  const userId = useSelector((state: RootState) => state.auth.user_id);
  const dispatch = useDispatch();
  const allSheets: ScoreSheetState[] = useSelector(
    (state: RootState) => state.allSheets.sheets
  );
  const loading = useSelector((state: RootState) => state.allSheets.loading);

  const fetchSheets = async () => {
    dispatch(setLoading(true));
    setError(null);
    try {
      const uri = `/get_score_sheet_by_user_id/${userId}`;
      const response = await api.get(uri);

      if (response.status === 200) {
        if (response.data.score_sheets) {
          dispatch(
            setAllSheets(response.data.score_sheets as ScoreSheetState[])
          );
        }
      } else {
        toast.error("Unexpected response from server");
      }
    } catch (err: any) {
      const message =
        err.response?.data?.message ||
        (err.response?.status === 404
          ? "No score sheets found for this user."
          : "Failed to fetch score sheets. Please try again.");

      toast.error(message);
      setError(message);
    } finally {
      dispatch(setLoading(false));
    }
  };

  useEffect(() => {
    if (userId) {
      fetchSheets();
    }
  }, [userId]);

  return (
    <div className="min-h-screen bg-gradient-to-br from-[#fbfcff] via-[#EDEFFF] to-[#D4E0FF] p-6">
      {loading && <LoadingState />}
      {error && <ErrorState message={error} />}
      <div className="flex mt-30 justify-end mr-6 mb-4">
        <CreateSheetDialog isCreate={true} />
      </div>

      {!loading && !error && allSheets.length === 0 && (
        <div className="flex items-center justify-center min-h-[50vh]">
          <p className="text-gray-500 text-lg">No score sheets found.</p>
        </div>
      )}
      {!loading && !error && allSheets.length > 0 && (
        <div className="container mx-auto px-4 py-6">
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-8">
            {allSheets.map((sheet: ScoreSheetState) => (
              <SheetCard
                key={sheet.id}
                id={sheet.id}
                name={sheet.name}
                committee_name={sheet.committee_name}
                chair={sheet.chair}
                vice_chair={sheet.vice_chair}
                rapporteur={sheet.rapporteur}
                created_by={sheet.created_by}
                created_at={sheet.created_at}
                updated_at={sheet.updated_at}
              />
            ))}
          </div>
        </div>
      )}

      <ToastComponent />
    </div>
  );
};

export default SeeSheets;
