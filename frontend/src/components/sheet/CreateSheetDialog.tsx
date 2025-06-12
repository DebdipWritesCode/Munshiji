import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import ToastComponent from "../ToastComponent";
import { Button } from "../ui/button";
import { DialogHeader } from "../ui/dialog";
import CreateSheetForm from "./CreateSheetForm";

interface CreateSheetDialogProps {
  isCreate?: boolean;
  id?: number;
  name?: string;
  committee_name?: string;
  chair?: string;
  vice_chair?: string;
  rapporteur?: string;
  btn_ClassName?: string;
  btn_Variant?: "default" | "outline" | "ghost" | "secondary" | "destructive";
}

const CreateSheetDialog: React.FC<CreateSheetDialogProps> = ({
  isCreate = true,
  id = undefined,
  name = "",
  committee_name = "",
  chair = "",
  vice_chair = "",
  rapporteur = "",
  btn_ClassName = "",
  btn_Variant = "default",
}) => {
  return (
    <>
      <Dialog>
        <DialogTrigger asChild>
          <Button variant={btn_Variant} className={btn_ClassName}>
            {isCreate ? "Create Score Sheet" : "Edit"}
          </Button>
        </DialogTrigger>
        <DialogContent className="sm:max-w-[425px]">
            <DialogHeader>
              <DialogTitle>
                {isCreate ? "Create Score Sheet" : "Edit Score Sheet"}
              </DialogTitle>
              <DialogDescription>
                {isCreate
                  ? "Fill in the details to create a new score sheet."
                  : "Edit the details of the selected score sheet."}
              </DialogDescription>
            </DialogHeader>

            <CreateSheetForm
              isCreate={isCreate}
              id={id}
              name={name}
              committee_name={committee_name}
              chair={chair}
              vice_chair={vice_chair}
              rapporteur={rapporteur}
            />
        </DialogContent>
      </Dialog>

      <ToastComponent />
    </>
  );
};

export default CreateSheetDialog;
