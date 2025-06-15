import ExcelJS from "exceljs";
import { saveAs } from "file-saver";

export async function exportExcel(
  mainSheet: Record<string, string | number>[],
  delegateSheets: Record<
    string,
    { parameter: string; value: number; note?: string }[]
  >,
  filename: string = "scores.xlsx",
  returnBlob?: boolean
): Promise<Blob | void> {
  try {
    const workBook = new ExcelJS.Workbook();

    
    const main = workBook.addWorksheet("Main Sheet");
    if (mainSheet.length > 0) {
      const columns = Object.keys(mainSheet[0]).map((key) => ({
        header: key,
        key: key,
      }));
      main.columns = columns;
      main.addRows(mainSheet);
      
      main.getRow(1).font = { bold: true };
      main.getRow(1).alignment = { horizontal: "center" };
    }

    for (const [delegateName, rows] of Object.entries(delegateSheets)) {
      const sheet = workBook.addWorksheet(delegateName);
      
      if (rows.length > 0) {
        sheet.columns = [
          { header: "Parameter", key: "parameter" },
          { header: "Value", key: "value" },
          { header: "Note", key: "note" },
        ];
        
        sheet.addRows(rows);
        sheet.getRow(1).font = { bold: true };
        sheet.getRow(1).alignment = { horizontal: "center" };
      }
    }
    
    const buffer = await workBook.xlsx.writeBuffer();
    const blob = new Blob([buffer], {
      type:
      "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
    });
    
    
    if (returnBlob) {
      return blob;
    } else {
      saveAs(blob, filename);
    }
  } catch (error) {
    console.error("Failed to export Excel:", error);
    throw new Error("Excel export failed. Please try again.");
  }
}
