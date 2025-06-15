import html2Canvas from "html2canvas";
import jsPDF from "jspdf";
import jsZip from "jszip";
import { saveAs } from "file-saver";

async function renderComponentToPDFBlob(element: HTMLElement): Promise<Blob> {
  try {
    
    const canvas = await html2Canvas(element, {
      scale: 2,
      useCORS: true,
    });

    const imgData = canvas.toDataURL("image/png");
    const pdf = new jsPDF({
      orientation: "portrait",
      unit: "px",
      format: [canvas.width, canvas.height],
    });
    pdf.addImage(imgData, "PNG", 0, 0, canvas.width, canvas.height);
    return pdf.output("blob");
  } catch (err) {
    console.error("Error generating PDF for element:", element, err);
    throw err;
  }
}


export async function generateAndZipPDFsWithExcel(
  delegateNames: string[],
  excelBlob: Blob | void
) {
  const zip = new jsZip();

  for (const name of delegateNames) {
    const element = document.getElementById(`report-${name}`);
    if (!element) continue;

    await new Promise((resolve) => setTimeout(resolve, 100));
    const blob = await renderComponentToPDFBlob(element);
    zip.file(`${name}.pdf`, blob);
  }

  if (excelBlob) {
    zip.file("scoresheet.xlsx", excelBlob);
  }

  const content = await zip.generateAsync({ type: "blob" });
  saveAs(content, "delegate-reports.zip");
}
