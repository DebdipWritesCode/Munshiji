import api from "@/api/axios";
import { Button } from "../ui/button";
import { toast } from "react-toastify";

interface CreateDelegateButtonProps {
  score_sheet_id: number;
}

const CreateDelegateButton: React.FC<CreateDelegateButtonProps> = ({
  score_sheet_id,
}) => {
  const handleCreateDelegate = async () => {
    try {
      const uri = "/create_delegate";
      const createDelegateOptions = {
        score_sheet_id: score_sheet_id,
        name: "New Delegate",
      };

      const response = await api.post(uri, createDelegateOptions);
      if (response.status === 200) {
        const newDelegate = response.data;
        toast.success(`Delegate ${newDelegate.name} created successfully!`);
        setTimeout(() => {
          window.location.reload();
        }, 600);
      } else {
        throw new Error("Unexpected response from server");
      }
    } catch (err: any) {
      const message =
        err.response?.data?.message ||
        err.message ||
        "Failed to create delegate. Please try again.";
      toast.error(message);
    }
  };

  return (
    <Button
      variant="default"
      className="ml-5 mt-2 bg-blue-700"
      type="button"
      onClick={handleCreateDelegate}>
      Create Delegate
    </Button>
  );
};

export default CreateDelegateButton;
