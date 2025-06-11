import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogTitle,
  DialogTrigger,
} from "./ui/dialog";
import ToastComponent from "./ToastComponent";
import { Button } from "./ui/button";
import { DialogHeader } from "./ui/dialog";
import CreateSheetForm from "./CreateSheetForm";
import { Card } from "./ui/card";

interface CreateSheetDialogProps {
  isCreate?: boolean;
  name?: string;
  committee_name?: string;
  chair?: string;
  vice_chair?: string;
  rapporteur?: string;
}

const CreateSheetDialog: React.FC<CreateSheetDialogProps> = ({
  isCreate = true,
  name = "",
  committee_name = "",
  chair = "",
  vice_chair = "",
  rapporteur = "",
}) => {
  return (
    <>
      <Dialog>
        <DialogTrigger asChild>
          <Button variant="default">Create Score Sheet</Button>
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
