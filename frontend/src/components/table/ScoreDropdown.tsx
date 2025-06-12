import type { TableScore } from "@/utils/getTableInterfaceTypes";
import { DropdownMenu, DropdownMenuItem } from "../ui/dropdown-menu";
import {
  DropdownMenuContent,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@radix-ui/react-dropdown-menu";
import { ChevronDown, Trash } from "lucide-react";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import api from "@/api/axios";
import { useDispatch } from "react-redux";
import { toast } from "react-toastify";
import {
  addNewScore,
  deleteScore,
  updateScore,
  updateScoreNote,
} from "@/slices/sheetDetailsSlice";

interface ScoreDropdownProps {
  scores: TableScore[];
  parameter_id: number;
  delegate_id: number;
}

const ScoreDropdown: React.FC<ScoreDropdownProps> = ({
  scores,
  parameter_id,
  delegate_id,
}) => {
  const dispatch = useDispatch();

  const scoreValueChangeHandler = async (id: number, value: number) => {
    const uri = `/update_score`;
    const updateScoreOptions = {
      id: id,
      value: value,
    };

    try {
      const response = await api.patch(uri, updateScoreOptions);
      if (response.status === 200) {
        dispatch(updateScore({ id, value }));
      } else {
        toast.error("Unexpected response from server");
      }
    } catch (err: any) {
      const message =
        err.response?.data?.message ||
        err.message ||
        "Failed to update score. Please try again.";
      toast.error(message);
    }
  };

  const scoreNoteChangeHandler = async (id: number, note: string) => {
    const uri = `/update_score`;
    const updateScoreOptions = {
      id: id,
      note: note,
    };

    try {
      const response = await api.patch(uri, updateScoreOptions);
      if (response.status === 200) {
        dispatch(updateScoreNote({ id, note }));
      } else {
        toast.error("Unexpected response from server");
      }
    } catch (err: any) {
      const message =
        err.response?.data?.message ||
        err.message ||
        "Failed to update score. Please try again.";
      toast.error(message);
    }
  };

  const addScoreHandler = async () => {
    const uri = "create_score";
    const newScoreOptions = {
      value: 0,
      note: "",
      parameter_id: parameter_id,
      delegate_id: delegate_id,
    };

    try {
      const response = await api.post(uri, newScoreOptions);
      if (response.status === 200) {
        const newScore = response.data;
        dispatch(addNewScore(newScore));
        toast.success("Score added!");
      } else {
        toast.error("Unexpected response from server");
      }
    } catch (err: any) {
      const message =
        err.response?.data?.message ||
        err.message ||
        "Failed to add new score. Please try again.";
      toast.error(message);
    }
  };

  const deleteScoreHandler = async (id: number) => {
    const uri = `/delete_score/${id}`;

    try {
      const response = await api.delete(uri);
      if (response.status === 200) {
        dispatch(deleteScore(id));
        toast.success("Score deleted successfully!");
      } else {
        toast.error("Unexpected response from server");
      }
    } catch (err: any) {
      const message =
        err.response?.data?.message ||
        err.message ||
        "Failed to delete score. Please try again.";
      toast.error(message);
    }
  };

  return (
    <DropdownMenu>
      <DropdownMenuTrigger>
        <ChevronDown className=" cursor-pointer" />
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56">
        {scores.map((score) => (
          <DropdownMenuItem key={score.score_id} className="">
            <div className="">
              <Input
                type="number"
                value={score.value ?? ""}
                onChange={(e) => {
                  const newValue = parseFloat(e.target.value);
                  if (!isNaN(newValue)) {
                    scoreValueChangeHandler(score.score_id, newValue);
                  }
                }}
              />
              <Input
                type="text"
                value={score.note ?? ""}
                placeholder="Note"
                onChange={(e) => {
                  const newNote = e.target.value;
                  scoreNoteChangeHandler(score.score_id, newNote);
                }}
              />
            </div>
            <Button
              type="button"
              variant="destructive"
              onClick={() => deleteScoreHandler(score.score_id)}>
              <Trash />
            </Button>
          </DropdownMenuItem>
        ))}
        <DropdownMenuSeparator />
        <Button type="button" variant="ghost" onClick={addScoreHandler}>
          Add
        </Button>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default ScoreDropdown;
