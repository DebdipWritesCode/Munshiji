import { useSelector } from "react-redux";
import type { RootState } from "@/store/store";
import {
  prepareTableData,
  setTableColumns,
} from "@/utils/tableUtils";
import type { DelegateWithScores } from "@/utils/getTableInterfaceTypes";

const ScoreTable = () => {
  const parameters = useSelector(
    (state: RootState) => state.sheetDetails.parameters
  );
  const delegates = useSelector(
    (state: RootState) => state.sheetDetails.delegates
  );
  const scores = useSelector((state: RootState) => state.sheetDetails.scores);

  const tableColumns = setTableColumns(parameters);
  const tableData: DelegateWithScores[] = prepareTableData(
    scores,
    parameters,
    delegates
  );

  return <div>hfnofnodwiv</div>;
};

export default ScoreTable;
