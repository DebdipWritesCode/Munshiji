import type { RootState } from "@/store/store";
import { useSelector } from "react-redux";
import { format } from "date-fns";
import {
  PieChart,
  Pie,
  Cell,
  Tooltip,
  ResponsiveContainer,
  Legend,
} from "recharts";

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
    <div className="max-w-4xl mx-auto p-8 bg-white text-gray-800 font-sans">
      {/* Header Section */}
      <div className="border-b-2 border-gray-200 pb-6 mb-8">
        <h1 className="text-3xl font-bold text-gray-900 mb-6 text-center">
          Delegate Performance Report
        </h1>
        
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 bg-gray-50 p-6 rounded-lg">
          <div className="space-y-3">
            <div className="flex">
              <span className="font-semibold text-gray-700 w-24">Sheet:</span>
              <span className="text-gray-900">{sheetName}</span>
            </div>
            <div className="flex">
              <span className="font-semibold text-gray-700 w-24">Committee:</span>
              <span className="text-gray-900">{committeeName}</span>
            </div>
            <div className="flex">
              <span className="font-semibold text-gray-700 w-24">Chairs:</span>
              <span className="text-gray-900">{chairs}</span>
            </div>
          </div>
          
          <div className="space-y-3">
            <div className="flex">
              <span className="font-semibold text-gray-700 w-32">Created:</span>
              <span className="text-gray-900">{formattedCreatedDate}</span>
            </div>
            <div className="flex">
              <span className="font-semibold text-gray-700 w-32">Generated:</span>
              <span className="text-gray-900">{formattedFeedbackDate}</span>
            </div>
          </div>
        </div>
        
        <div className="mt-6 p-4 bg-blue-50 border-l-4 border-blue-400 rounded-r-lg">
          <h3 className="text-lg font-semibold text-blue-900">
            Report for: {delegateName}
          </h3>
        </div>
      </div>

      {/* Main Content Section */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
        {/* Feedback Section */}
        <div className="space-y-4">
          <h2 className="text-xl font-bold text-gray-900 border-b border-gray-300 pb-2">
            Performance Feedback
          </h2>
          <div className="bg-gray-50 p-6 rounded-lg">
            <p className="text-gray-800 leading-relaxed whitespace-pre-wrap">
              {feedback}
            </p>
          </div>
        </div>

        {/* Chart Section */}
        <div className="space-y-4">
          <h2 className="text-xl font-bold text-gray-900 border-b border-gray-300 pb-2">
            Performance Metrics
          </h2>
          <div className="bg-white border border-gray-200 rounded-lg p-4">
            <div className="h-96 w-full">
              <ResponsiveContainer>
                <PieChart>
                  <Pie
                    data={data}
                    dataKey="value"
                    nameKey="parameterName"
                    cx="50%"
                    cy="50%"
                    outerRadius={120}
                    label>
                    {data.map((_, index) => (
                      <Cell
                        key={`cell-${index}`}
                        fill={COLORS[index % COLORS.length]}
                      />
                    ))}
                  </Pie>
                  <Tooltip />
                  <Legend />
                </PieChart>
              </ResponsiveContainer>
            </div>
          </div>
        </div>
      </div>

      {/* Disclaimer Section */}
      <div className="border-t-2 border-gray-200 pt-6">
        <h3 className="text-lg font-semibold text-gray-900 mb-4">Important Notice</h3>
        <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-4 space-y-2">
          <p className="text-sm text-gray-700">
            • This feedback is auto-generated by an AI model based on numerical scores.
          </p>
          <p className="text-sm text-gray-700">
            • The feedback is derived from scores provided by the chairs and does not reflect personal opinions or biases.
          </p>
          <p className="text-sm text-gray-700">
            • For personalized feedback and detailed evaluation, please contact the committee chairs directly.
          </p>
        </div>
      </div>
    </div>
  );
};

export default DelegateReport;