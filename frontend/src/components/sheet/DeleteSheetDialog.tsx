import {
  Dialog,
  DialogHeader,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import ToastComponent from "../ToastComponent";
import { Button } from "../ui/button";
import { useState } from "react";
import api from "@/api/axios";
import { toast } from "react-toastify";

interface DeleteSheetDialogProps {
  id: number | null;
}

const DeleteSheetDialog: React.FC<DeleteSheetDialogProps> = ({ id }) => {
  const [_, setLoading] = useState(false);

  const handleDelete = async (id: number | null) => {
    setLoading(true);
    try {
      if (!id) {
        toast.error("Invalid score sheet ID");
        return;
      }

      const response = await api.delete(`/delete_score_sheet/${id}`);
      if (response.status === 200) {
        toast.success("Score sheet deleted successfully!");
        // Refresh the page
        setTimeout(() => {
          window.location.reload();
        }, 1000);
      } else {
        throw new Error("Unexpected response from server");
      }
    } catch (err: any) {
      if (err.response) {
        if (err.response.data?.message) {
          toast.error(err.response.data.message);
        } else {
          toast.error("Failed to delete score sheet. Please try again.");
        }
      }
    } finally {
      setLoading(false);
    }
  }

  return (
    <>
      <Dialog>
        <DialogTrigger asChild>
          <Button variant="destructive" className=" h-8">Delete</Button>
        </DialogTrigger>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Delete Score Sheet</DialogTitle>
            <DialogDescription>
              Are you sure you want to delete this score sheet?
            </DialogDescription>
          </DialogHeader>

          <DialogFooter>
            <Button variant="destructive" onClick={() => handleDelete(id)}>
              Confirm
            </Button>
            <DialogClose>
              <Button variant="outline">Cancel</Button>
            </DialogClose>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      <ToastComponent />
    </>
  );
};

export default DeleteSheetDialog;
