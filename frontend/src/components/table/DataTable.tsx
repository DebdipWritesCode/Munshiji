import {
  Table,
  TableHeader,
  TableBody,
  TableRow,
  TableHead,
  TableCell,
} from "@/components/ui/table";
import React from "react";
import CreateParameterDialog from "../parameter/CreateParameterDialog";
import DelegateNameCell from "./DelegateNameCell";
import { useSelector } from "react-redux";
import type { RootState } from "@/store/store";
import { getTopDelegates } from "@/utils/scoresUtils";

export interface CustomColumn<T> {
  id?: string;
  header: string | React.ReactNode;
  cell: (row: T) => React.ReactNode;
  parameterProps?: {
    id?: number;
    name?: string;
    rule_type?: string;
    is_special_parameter?: boolean;
    special_scores_rule?: string;
    special_length_rule?: string;
    score_weight?: number;
    length_weight?: number;
  };
}

interface DataTableProps<T> {
  columns: CustomColumn<T>[];
  data: T[];
}

const topColors = {
  0: "border-purple-500 border-3 border-dashed",
  1: "border-green-500 border-3 border-dashed",
  2: "border-pink-500 border-3 border-dashed",
};

export function DataTable<T>({ columns, data }: DataTableProps<T>) {
  const scores = useSelector((state: RootState) => state.scores.scores);
  const top3DelegatesIds = getTopDelegates(scores, 3);

  return (
    <div className="rounded-md border overflow-x-auto">
      <Table>
        <TableHeader>
          <TableRow>
            {columns.map((column, index) => (
              <TableHead key={column.id ?? index}>
                <div className="flex justify-between items-center gap-2">
                  <p>{column.header}</p>
                  {column.header !== "Delegate" && column.header !== "Total" ? (
                    <CreateParameterDialog
                      isCreate={false}
                      btn_Variant="ghost"
                      id={column.parameterProps?.id ?? undefined}
                      name={column.parameterProps?.name ?? ""}
                      rule_type={column.parameterProps?.rule_type ?? ""}
                      is_special_parameter={String(
                        column.parameterProps?.is_special_parameter ?? "false"
                      )}
                      special_scores_rule={
                        column.parameterProps?.special_scores_rule ?? ""
                      }
                      special_length_rule={
                        column.parameterProps?.special_length_rule ?? ""
                      }
                      score_weight={column.parameterProps?.score_weight ?? 0}
                      length_weight={column.parameterProps?.length_weight ?? 0}
                    />
                  ) : null}
                </div>
              </TableHead>
            ))}
          </TableRow>
        </TableHeader>
        <TableBody>
          {data.length > 0 ? (
            data.map((row, rowIndex) => {
              const delegateId = (row as any).delegate_id as number;
              const topIndex = top3DelegatesIds.indexOf(delegateId);
              const rowClassName =
                topIndex !== -1 ? topColors[topIndex as 0 | 1 | 2] : "";

              return (
                <TableRow key={rowIndex}>
                  {columns.map((column, colIndex) => (
                    <TableCell key={column.id ?? colIndex} className={rowClassName}>
                      {column.header === "Delegate" ? (
                        <DelegateNameCell
                          delegateName={
                            column.cell(row) &&
                            (column.cell(row) as any).props &&
                            (column.cell(row) as any).props.children
                              ? ((column.cell(row) as any).props
                                  .children as string)
                              : ""
                          }
                          delegate_id={(row as any).delegate_id as number}
                        />
                      ) : (
                        column.cell(row)
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              );
            })
          ) : (
            <TableRow>
              <TableCell colSpan={columns.length} className="text-center h-24">
                No results.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}
