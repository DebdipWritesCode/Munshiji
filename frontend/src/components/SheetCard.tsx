import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { format } from "date-fns";
import { Link } from "react-router";

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
    ? format(new Date(created_at), "PPpp")
    : "Unknown creation date";

  const formattedUpdatedAt = updated_at
    ? format(new Date(updated_at), "PPpp")
    : "Never updated";

  return (
    <Card className="w-[350px] hover:shadow-lg transition-shadow flex flex-col">
      <div className="flex-1">
        <CardHeader>
          <CardTitle className="text-xl">
            {name || "Untitled Score Sheet"}
          </CardTitle>
          <CardDescription>
            {committee_name || "No committee specified"}
          </CardDescription>
        </CardHeader>
        <CardContent className="grid gap-4">
          <div className="grid grid-cols-2 gap-4">
            <div>
              <p className="text-sm text-muted-foreground">Chair</p>
              <p className="font-medium">{chair || "Not specified"}</p>
            </div>
            {vice_chair && (
              <div>
                <p className="text-sm text-muted-foreground">Vice Chair</p>
                <p className="font-medium">{vice_chair}</p>
              </div>
            )}
          </div>

          {rapporteur && (
            <div>
              <p className="text-sm text-muted-foreground">Rapporteur</p>
              <p className="font-medium">{rapporteur}</p>
            </div>
          )}

          <div className="text-sm text-muted-foreground space-y-1">
            <p>Created: {formattedCreatedAt}</p>
            <p>Last updated: {formattedUpdatedAt}</p>
          </div>
        </CardContent>
      </div>

      <CardFooter className="flex justify-end">
        <Button asChild variant="outline">
          <Link to={`/score-sheets/${id}`}>See Details</Link>
        </Button>
      </CardFooter>
    </Card>
  );
};

export default SheetCard;
