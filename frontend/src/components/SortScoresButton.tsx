import { useDispatch, useSelector } from "react-redux";
import { Button } from "./ui/button";
import type { RootState } from "@/store/store";
import { sortDelegatesByScore  } from "@/utils/scoresUtils";
import { setDelegates } from "@/slices/sheetDetailsSlice";

const SortScoresButton = () => {
  const dispatch = useDispatch();

  const scores = useSelector((state: RootState) => state.scores.scores);
  const delegates = useSelector((state: RootState) => state.sheetDetails.delegates);

  const handleSorting = () => {
    const sortedDelegates = sortDelegatesByScore(
      scores,
      delegates
    );

    dispatch(setDelegates(sortedDelegates));
  };

  return (
    <>
      <Button onClick={handleSorting} className="bg-blue-600">Sort Scores</Button>
    </>
  );
};

export default SortScoresButton;
