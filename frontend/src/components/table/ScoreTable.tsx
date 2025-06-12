import { useSelector } from "react-redux";
import type { RootState } from "@/store/store";
import { calculateRuleWiseTotal, calculateTotalScore, prepareTableData } from "@/utils/tableUtils";
import type { DelegateWithScores, TableParameter } from "@/utils/getTableInterfaceTypes";
import { DataTable, type CustomColumn } from "./DataTable";
import Cell from "./Cell";

const ScoreTable = () => {
  const parameters = useSelector((state: RootState) => state.sheetDetails.parameters);
  const delegates = useSelector((state: RootState) => state.sheetDetails.delegates);
  const scores = useSelector((state: RootState) => state.sheetDetails.scores);

  const delegateData: DelegateWithScores[] = prepareTableData(scores, parameters, delegates);

  const columns: CustomColumn<DelegateWithScores>[] = [
    {
      id: "delegate_name",
      header: "Delegate",
      cell: (row) => <p>{row.name}</p>,
    },
    ...parameters.map((param) => ({
      id: `param_${param.id}`,
      header: param.name,
      cell: (row: any) => {
        const currentParam = row.parameters.find(
          (p: TableParameter) => p.parameter_id === param.id
        );

        if (!currentParam) return <p>N/A</p>;

        const valueToDisplay = calculateRuleWiseTotal(currentParam);

        return (
          <Cell
            parameter_id={param.id}
            delegate_id={row.delegate_id}
            scores={currentParam.scores}
            valueToDisplay={valueToDisplay}
          />
        );
      },
    })),
    {
      id: "total",
      header: "Total",
      cell: (row) => {
        const total = calculateTotalScore(row.parameters);
        return <p className="font-semibold">{total}</p>;
      },
    },
  ];

  return (
    <div className="p-4">
      <DataTable columns={columns} data={delegateData} />
    </div>
  );
};

export default ScoreTable;
