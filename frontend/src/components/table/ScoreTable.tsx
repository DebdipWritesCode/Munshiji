import { useDispatch, useSelector } from "react-redux";
import type { RootState } from "@/store/store";
import {
  calculateRuleWiseTotal,
  calculateTotalScore,
  prepareTableData,
} from "@/utils/tableUtils";
import type {
  DelegateWithScores,
  TableParameter,
} from "@/utils/getTableInterfaceTypes";
import { DataTable, type CustomColumn } from "./DataTable";
import Cell from "./Cell";
import { useEffect, useMemo } from "react";
import { setScores, type ScoreState } from "@/slices/scoresSlice";

const ScoreTable = () => {
  const parameters = useSelector(
    (state: RootState) => state.sheetDetails.parameters
  );

  const delegates = useSelector(
    (state: RootState) => state.sheetDetails.delegates
  );
  const scores = useSelector((state: RootState) => state.sheetDetails.scores);

  const dispatch = useDispatch();

  const delegateData: DelegateWithScores[] = useMemo(() => {
    return prepareTableData(scores, parameters, delegates);
  }, [scores, parameters, delegates]);

  useEffect(() => {
    const updatedScores: ScoreState[] = [];

    delegateData.forEach((delegate) => {
      delegate.parameters.forEach((param) => {
        const value = calculateRuleWiseTotal(param);
        updatedScores.push({
          delegate_id: delegate.delegate_id,
          parameter_id: param.parameter_id,
          valueToDisplay: value,
        });
      });
    });

    dispatch(setScores(updatedScores));
  }, [delegateData, dispatch]);

  const columns: CustomColumn<DelegateWithScores>[] = [
    {
      id: "delegate_name",
      header: "Delegate",
      cell: (row) => <p>{row.name}</p>,
      parameterProps: {
        id: undefined,
        name: "",
        rule_type: "",
        is_special_parameter: false,
        special_scores_rule: "",
        special_length_rule: "",
        score_weight: 0,
        length_weight: 0,
      },
    },
    ...parameters.map((param) => ({
      id: `param_${param.id}`,
      header: param.name,
      cell: (row: DelegateWithScores) => {
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
      parameterProps: {
        id: param.id,
        name: param.name,
        rule_type: param.rule_type,
        is_special_parameter: param.is_special_parameter,
        special_scores_rule: param.special_scores_rule,
        special_length_rule: param.special_length_rule,
        score_weight: param.score_weight,
        length_weight: param.length_weight,
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
    <div className="p-4 w-full">
      <DataTable columns={columns} data={delegateData} />
    </div>
  );
};

export default ScoreTable;
