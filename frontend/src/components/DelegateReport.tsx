import type { RootState } from "@/store/store";
import { useSelector } from "react-redux";
import { format } from "date-fns";
import { Pie } from "react-chartjs-2";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";

ChartJS.register(ArcElement, Tooltip, Legend);

interface ChartData {
  parameterName: string;
  value: number;
}

interface DelegateReportProps {
  data: ChartData[];
  delegateName: string;
  feedback: string;
}

const COLORS = [
  "#8884d8",
  "#82ca9d",
  "#ffc658",
  "#ff8042",
  "#8dd1e1",
  "#d0ed57",
  "#a4de6c",
  "#d88884",
  "#a28bd4",
  "#ffc0cb",
];

function joinChairs(
  chair: string | undefined,
  viceChair: string | undefined,
  rapporteur: string | undefined
): string {
  const names = [chair, viceChair, rapporteur]
    .filter((name) => name)
    .join(", ");
  return names || "No chairs assigned";
}

const DelegateReport: React.FC<DelegateReportProps> = ({
  data,
  delegateName,
  feedback,
}) => {
  const chartData = {
    labels: data.map((d) => d.parameterName),
    datasets: [
      {
        data: data.map((d) => d.value),
        backgroundColor: COLORS.slice(0, data.length),
        borderWidth: 1,
      },
    ],
  };

  const chartOptions = {
    responsive: false, // important for html2canvas compatibility
    plugins: {
      legend: {
        position: "bottom" as const,
      },
    },
  };

  const sheetName = useSelector(
    (state: RootState) => state.sheetDetails.score_sheet?.name
  );
  const committeeName = useSelector(
    (state: RootState) => state.sheetDetails.score_sheet?.committee_name
  );
  const chairName = useSelector(
    (state: RootState) => state.sheetDetails.score_sheet?.chair
  );
  const viceChairName = useSelector(
    (state: RootState) => state.sheetDetails.score_sheet?.vice_chair
  );
  const rapporteurName = useSelector(
    (state: RootState) => state.sheetDetails.score_sheet?.rapporteur
  );
  const dateCreated = useSelector(
    (state: RootState) => state.sheetDetails.score_sheet?.created_at
  );
  const feedbackGeneratedDate = Date.now();

  const chairs = joinChairs(chairName, viceChairName, rapporteurName);

  const formattedCreatedDate = dateCreated
    ? format(new Date(dateCreated), "MMMM dd, yyyy")
    : "Unknown";

  const formattedFeedbackDate = format(
    new Date(feedbackGeneratedDate),
    "MMMM dd, yyyy"
  );

  return (
    <div
      id={`report-${delegateName}`}
      className="max-w-4xl mx-auto p-8 font-sans"
      style={{ backgroundColor: "white", color: "#2d3748" }}>
      {/* Header Section */}
      <div className="border-b-2 pb-6 mb-8" style={{ borderColor: "#e2e8f0" }}>
        <h1
          className="text-3xl font-bold mb-6 text-center"
          style={{ color: "#1a202c" }}>
          Delegate Performance Report
        </h1>

        <div
          className="grid grid-cols-1 md:grid-cols-2 gap-4 p-6 rounded-lg"
          style={{ backgroundColor: "#f9fafb" }}>
          <div className="space-y-3">
            <div className="flex">
              <span className="font-semibold w-24" style={{ color: "#4a5568" }}>
                Sheet:
              </span>
              <span style={{ color: "#1a202c" }}>{sheetName}</span>
            </div>
            <div className="flex">
              <span className="font-semibold w-24" style={{ color: "#4a5568" }}>
                Committee:
              </span>
              <span style={{ color: "#1a202c" }}>{committeeName}</span>
            </div>
            <div className="flex">
              <span className="font-semibold w-24" style={{ color: "#4a5568" }}>
                Chairs:
              </span>
              <span style={{ color: "#1a202c" }}>{chairs}</span>
            </div>
          </div>

          <div className="space-y-3">
            <div className="flex">
              <span className="font-semibold w-32" style={{ color: "#4a5568" }}>
                Created:
              </span>
              <span style={{ color: "#1a202c" }}>{formattedCreatedDate}</span>
            </div>
            <div className="flex">
              <span className="font-semibold w-32" style={{ color: "#4a5568" }}>
                Generated:
              </span>
              <span style={{ color: "#1a202c" }}>{formattedFeedbackDate}</span>
            </div>
          </div>
        </div>

        <div
          className="mt-6 p-4 rounded-r-lg"
          style={{
            backgroundColor: "#ebf8ff",
            borderLeft: "4px solid #63b3ed",
          }}>
          <h3 className="text-lg font-semibold" style={{ color: "#2c5282" }}>
            Report for: {delegateName}
          </h3>
        </div>
      </div>

      {/* Main Content Section */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
        {/* Feedback Section */}
        <div className="space-y-4">
          <h2
            className="text-xl font-bold border-b pb-2"
            style={{ color: "#1a202c", borderColor: "#e2e8f0" }}>
            Performance Feedback
          </h2>
          <div
            className="p-6 rounded-lg"
            style={{ backgroundColor: "#f9fafb" }}>
            <p
              style={{
                color: "#2d3748",
                lineHeight: "1.625",
                whiteSpace: "pre-wrap",
              }}>
              {feedback}
            </p>
          </div>
        </div>

        {/* Chart Section */}
        <div className="space-y-4">
          <h2
            className="text-xl font-bold border-b pb-2"
            style={{ color: "#1a202c", borderColor: "#e2e8f0" }}>
            Performance Metrics
          </h2>
          <div
            className="border rounded-lg p-4"
            style={{ backgroundColor: "white", borderColor: "#e2e8f0" }}>
            <div
              id={`chart-${delegateName}`}
              className="h-96 w-full flex items-center justify-center"
              style={{
                width: "400px", // fixed size
                height: "400px", // fixed size
              }}>
              <Pie data={chartData} options={chartOptions} width={300} height={350} />
            </div>
          </div>
        </div>
      </div>

      {/* Disclaimer Section */}
      <div className="border-t-2 pt-6" style={{ borderColor: "#e2e8f0" }}>
        <h3 className="text-lg font-semibold mb-4" style={{ color: "#1a202c" }}>
          Important Notice
        </h3>
        <div
          className="border rounded-lg p-4 space-y-2"
          style={{ backgroundColor: "#fefcbf", borderColor: "#f6e05e" }}>
          <p className="text-sm" style={{ color: "#4a5568" }}>
            • This feedback is auto-generated by an AI model based on numerical
            scores.
          </p>
          <p className="text-sm" style={{ color: "#4a5568" }}>
            • The feedback is derived from scores provided by the chairs and
            does not reflect personal opinions or biases.
          </p>
          <p className="text-sm" style={{ color: "#4a5568" }}>
            • For personalized feedback and detailed evaluation, please contact
            the committee chairs directly.
          </p>
        </div>
      </div>
    </div>
  );
};

export default DelegateReport;
