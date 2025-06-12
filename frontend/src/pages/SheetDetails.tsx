import LogoutDialog from "@/components/LogoutDialog";
import { useParams } from "react-router-dom";

const SheetDetails = () => {
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
    
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-[#fbfcff] via-[#EDEFFF] to-[#D4E0FF] p-6 flex justify-center pt-8">
      <div className="flex w-full justify-between items-center h-10">
        <div className="w-[45%]"></div>
        <h1 className="text-4xl text-blue-900 font-normal font-heading text-center">
          Munshiji
        </h1>
        <LogoutDialog />
      </div>
    </div>
  );
};

export default SheetDetails;
