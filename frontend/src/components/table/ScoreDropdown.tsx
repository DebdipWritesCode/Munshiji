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
import { useState } from "react";

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
  const [tempValues, setTempValues] = useState<Record<number, string>>({});
  const [tempNotes, setTempNotes] = useState<Record<number, string>>({});

  const dispatch = useDispatch();

  const scoreValueChangeHandler = async (id: number, value: number) => {
    const uri = `/update_score`;
    const updateScoreOptions = {
      score_id: id,
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
      score_id: id,
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
        const newScore = {
          ...response.data.score,
          score_id: response.data.id, 
        };
        dispatch(addNewScore(newScore));
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
      <DropdownMenuContent className="w-64 max-h-72 overflow-y-auto bg-slate-100 rounded-lg shadow-lg border border-gray-200 p-2 space-y-2">
        {scores.map((score) => (
          <DropdownMenuItem
            key={score.score_id}
            onSelect={(e) => {
              e.preventDefault();
            }}>
            <div className="flex">
              <Input
                type="number"
                step={0.01}
                value={tempValues[score.score_id] ?? String(score.value)}
                onChange={(e) => {
                  const val = e.target.value;
                  setTempValues((prev) => ({ ...prev, [score.score_id]: val }));

                  const num = parseFloat(val);
                  if (!isNaN(num)) {
                    scoreValueChangeHandler(score.score_id, num);
                  }
                }}
                onBlur={(e) => {
                  if (e.target.value === "") {
                    setTempValues((prev) => ({
                      ...prev,
                      [score.score_id]: "0",
                    }));
                    scoreValueChangeHandler(score.score_id, 0);
                  }
                }}
              />
              <Input
                type="text"
                maxLength={5}
                value={tempNotes[score.score_id] ?? score.note ?? ""}
                placeholder="Note"
                onChange={(e) => {
                  const newNote = e.target.value.trim();
                  setTempNotes((prev) => ({
                    ...prev,
                    [score.score_id]: newNote,
                  }));
                  scoreNoteChangeHandler(score.score_id, newNote);
                }}
              />
            </div>
            <Button
              type="button"
              variant="destructive"
              onClick={() => deleteScoreHandler(score.score_id)}>
              <Trash color="white" />
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
