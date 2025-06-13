import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { format } from "date-fns";
import { Link } from "react-router";
import { User, Calendar, Crown, Users, FileText } from "lucide-react";
import CreateSheetDialog from "./CreateSheetDialog";
import DeleteDialog from "../DeleteDialog";

interface SheetCardProps {
  id: number | null;
  name: string | null;
  committee_name: string | null;
  chair: string | null;
  vice_chair?: string | null;
  rapporteur?: string | null;
  created_by: number | null;
  created_at: string | null;
  updated_at: string | null;
}

const SheetCard: React.FC<SheetCardProps> = ({
  id,
  name,
  committee_name,
  chair,
  vice_chair,
  rapporteur,
  created_at,
  updated_at,
}) => {
  const formattedCreatedAt = created_at
    ? format(new Date(created_at), "MMM d, yyyy")
    : "Unknown";

  const formattedUpdatedAt = updated_at
    ? format(new Date(updated_at), "MMM d, yyyy")
    : "Never";

  return (
    <Card className="w-full max-w-sm flex flex-col justify-between hover:shadow-lg border-0 shadow-md bg-gradient-to-br from-white to-slate-50 dark:from-slate-900 dark:to-slate-800 group">
      <div className="flex-1">
        <CardHeader className="pb-3">
          <div className="flex items-start justify-between gap-2">
            <div className="flex-1 min-w-0">
              <CardTitle className="text-lg font-bold truncate text-slate-900 dark:text-slate-100 group-hover:text-blue-700 dark:group-hover:text-blue-400 transition-colors">
                {name || "Untitled Score Sheet"}
              </CardTitle>
              <CardDescription className="text-sm truncate mt-1 flex items-center gap-1.5 text-slate-600 dark:text-slate-400">
                <Users className="w-3.5 h-3.5 flex-shrink-0" />
                {committee_name || "No committee specified"}
              </CardDescription>
            </div>
            <Badge variant="secondary" className="bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200 hover:bg-blue-200 dark:hover:bg-blue-800 transition-colors">
              <FileText className="w-3 h-3 mr-1" />
              Sheet
            </Badge>
          </div>
        </CardHeader>

        <CardContent className="space-y-4 pb-4 h-[80%] flex flex-col justify-between">
          <div className="space-y-3">
            <div className="flex items-center gap-2 p-2.5 rounded-lg bg-amber-50 dark:bg-amber-950/20 border border-amber-200 dark:border-amber-800/30">
              <Crown className="w-4 h-4 text-amber-600 dark:text-amber-400 flex-shrink-0" />
              <div className="min-w-0 flex-1">
                <p className="text-xs font-medium text-amber-800 dark:text-amber-300">Chair</p>
                <p className="font-semibold text-sm text-amber-900 dark:text-amber-100 truncate">
                  {chair || "Not specified"}
                </p>
              </div>
            </div>

            {vice_chair && (
              <div className="flex items-center gap-2 p-2.5 rounded-lg bg-emerald-50 dark:bg-emerald-950/20 border border-emerald-200 dark:border-emerald-800/30">
                <User className="w-4 h-4 text-emerald-600 dark:text-emerald-400 flex-shrink-0" />
                <div className="min-w-0 flex-1">
                  <p className="text-xs font-medium text-emerald-800 dark:text-emerald-300">Vice Chair</p>
                  <p className="font-semibold text-sm text-emerald-900 dark:text-emerald-100 truncate">
                    {vice_chair}
                  </p>
                </div>
              </div>
            )}

            {rapporteur && (
              <div className="flex items-center gap-2 p-2.5 rounded-lg bg-purple-50 dark:bg-purple-950/20 border border-purple-200 dark:border-purple-800/30">
                <FileText className="w-4 h-4 text-purple-600 dark:text-purple-400 flex-shrink-0" />
                <div className="min-w-0 flex-1">
                  <p className="text-xs font-medium text-purple-800 dark:text-purple-300">Rapporteur</p>
                  <p className="font-semibold text-sm text-purple-900 dark:text-purple-100 truncate">
                    {rapporteur}
                  </p>
                </div>
              </div>
            )}
          </div>

          <div className="pt-2 border-t border-slate-200 dark:border-slate-700">
            <div className="flex items-center justify-between text-xs text-slate-500 dark:text-slate-400">
              <div className="flex items-center gap-1">
                <Calendar className="w-3 h-3" />
                <span>Created {formattedCreatedAt}</span>
              </div>
              <div className="text-right">
                <span>Updated {formattedUpdatedAt}</span>
              </div>
            </div>
          </div>
        </CardContent>
      </div>

      <CardFooter className="flex flex-wrap justify-end gap-2 pt-4 border-t border-slate-100 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900/50 rounded-b-lg">
        <Button asChild variant="outline" size="sm" className="border-blue-200 text-blue-700 hover:bg-blue-50 hover:text-blue-800 dark:border-blue-800 dark:text-blue-400 dark:hover:bg-blue-950 dark:hover:text-blue-300">
          <Link to={`/sheet/${id}`}>See Details</Link>
        </Button>
        <CreateSheetDialog
          isCreate={false}
          btn_ClassName="bg-slate-200 text-slate-700 hover:bg-slate-300 hover:text-slate-800 dark:bg-slate-700 dark:text-slate-300 dark:hover:bg-slate-600 dark:hover:text-slate-200 h-8"
          btn_Variant="secondary"
          id={id || undefined}
          name={name || ""}
          committee_name={committee_name || ""}
          chair={chair || ""}
          vice_chair={vice_chair || ""}
          rapporteur={rapporteur || ""}
        />
        <DeleteDialog id={id} uri="delete_score_sheet" deleteItem="score sheet" />
      </CardFooter>
    </Card>
  );
};

export default SheetCard;