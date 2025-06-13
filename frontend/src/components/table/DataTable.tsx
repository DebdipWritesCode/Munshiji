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

const minWidth = "180px"; 

interface DataTableProps<T> {
  columns: CustomColumn<T>[];
  data: T[];
}

const topColors = {
  0: "bg-yellow-50 border-l-4 border-l-yellow-500",
  1: "bg-purple-50 border-l-4 border-l-purple-600",
  2: "bg-cyan-50 border-l-4 border-l-cyan-500",
};

export function DataTable<T>({ columns, data }: DataTableProps<T>) {
  const scores = useSelector((state: RootState) => state.scores.scores);
  const top3DelegatesIds = getTopDelegates(scores, 3);

  return (
    <div className="rounded-lg border border-gray-200 bg-white shadow-sm overflow-x-auto w-[90vw]">
      <div className="inline-block min-w-full align-middle">
        <Table className="min-w-full divide-y divide-gray-200">
          <TableHeader className="bg-gray-50">
            <TableRow>
              {columns.map((column, index) => (
                <TableHead
                  key={column.id ?? index}
                  className="px-6 py-3 text-left text-md font-medium text-gray-500 uppercase tracking-wider"
                  style={{ minWidth: minWidth || "150px" }} // Set minimum width
                >
                  <div className="flex justify-between items-center gap-2">
                    <p className="font-semibold">{column.header}</p>
                    {column.header !== "Delegate" &&
                    column.header !== "Total" ? (
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
                        length_weight={
                          column.parameterProps?.length_weight ?? 0
                        }
                      />
                    ) : null}
                  </div>
                </TableHead>
              ))}
            </TableRow>
          </TableHeader>
          <TableBody className="bg-white divide-y divide-gray-200">
            {data.length > 0 ? (
              data.map((row, rowIndex) => {
                const delegateId = (row as any).delegate_id as number;
                const topIndex = top3DelegatesIds.indexOf(delegateId);
                const rowClassName =
                  topIndex !== -1 ? topColors[topIndex as 0 | 1 | 2] : "";

                return (
                  <TableRow
                    key={rowIndex}
                    className={`hover:bg-gray-50 ${rowClassName}`}>
                    {columns.map((column, colIndex) => {
                      const isFirst = colIndex === 0;
                      const isLast = colIndex === columns.length - 1;

                      const borderClass = isFirst
                        ? "border-l-4"
                        : isLast
                        ? "border-r-4 border-l-0"
                        : "border-none";

                      return (
                        <TableCell
                          key={column.id ?? colIndex}
                          className={`px-6 py-4 whitespace-nowrap text-sm text-gray-900 ${rowClassName} ${borderClass}`}
                          style={{ minWidth: minWidth || "150px" }} // Set minimum width
                        >
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
                      );
                    })}
                  </TableRow>
                );
              })
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columns.length}
                  className="px-6 py-4 text-center text-sm text-gray-500">
                  No results found.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}
