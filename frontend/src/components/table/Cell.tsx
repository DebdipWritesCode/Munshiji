import type { TableScore } from "@/utils/getTableInterfaceTypes"
import ScoreDropdown from "./ScoreDropdown";

interface CellProps {
  scores: TableScore[];
  parameter_id: number;
  delegate_id: number;
  valueToDisplay: number;
}

const Cell: React.FC<CellProps> = ({
  scores,
  parameter_id,
  delegate_id,
  valueToDisplay,
}) => {
  return (
    <div className="">
      <p>{valueToDisplay}</p>
      <ScoreDropdown
        scores={scores}
        parameter_id={parameter_id}
        delegate_id={delegate_id}
      />
    </div>
  )
}

export default Cell