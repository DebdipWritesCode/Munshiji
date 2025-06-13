import {
  Dialog,
  DialogFooter,
  DialogTrigger,
  DialogContent,
  DialogTitle,
  DialogDescription,
  DialogHeader,
  DialogClose,
} from "../ui/dialog";
import { Button } from "../ui/button";
import api from "@/api/axios";
import { useState } from "react";
import { toast } from "react-toastify";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Pen } from "lucide-react";
import { Input } from "../ui/input";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../ui/form";
import DeleteDialog from "../DeleteDialog";

interface DelegateNameCellProps {
  delegateName: string;
  delegate_id: number;
}

const onlyAlphaAndSpaces = /^[A-Za-z\s]*$/;

const updateDelegateNameSchema = z.object({
  name: z
    .string()
    .min(1, "Delegate name is required")
    .max(30, "Delegate name must be at most 30 characters long")
    .regex(
      onlyAlphaAndSpaces,
      "Delegate name must contain only letters and spaces"
    ),
  delegate_id: z
    .number()
    .int()
    .positive("Delegate ID must be a positive integer"),
});

const DelegateNameCell: React.FC<DelegateNameCellProps> = ({
  delegateName,
  delegate_id,
}) => {
  const [loading, setLoading] = useState<boolean>(false);

  const form = useForm<z.infer<typeof updateDelegateNameSchema>>({
    resolver: zodResolver(updateDelegateNameSchema),
    defaultValues: {
      name: delegateName,
      delegate_id: delegate_id,
    },
  });

  const handleEditDelegateName = async (
    values: z.infer<typeof updateDelegateNameSchema>
  ) => {
    const uri = "update_delegate_name_by_id";

    setLoading(true);
    try {
      const response = await api.patch(`/${uri}`, values);
      if (response.status === 200) {
        toast.success("Delegate name updated successfully!");
        setTimeout(() => {
          window.location.reload();
        }, 1000);
      } else {
        throw new Error("Unexpected response from server");
      }
    } catch (err: any) {
      if (err.response?.data?.message) {
        toast.error(err.response.data.message);
      } else if (err.request) {
        toast.error("No response from server. Please check your connection.");
      } else {
        toast.error("An error occurred: " + err.message);
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <Dialog>
        <DialogTrigger asChild>
          <Button variant="ghost" className="h-8 flex justify-between w-full">
            {delegateName}
            <Pen />
          </Button>
        </DialogTrigger>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Edit Delegate</DialogTitle>
            <DialogDescription>
              Edit the name of the delegate.
            </DialogDescription>
          </DialogHeader>

          <Form {...form}>
            <form
              onSubmit={form.handleSubmit(handleEditDelegateName)}
              className="space-y-4 px-1">
              <FormField
                control={form.control}
                name="name"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Name</FormLabel>
                    <FormControl>
                      <Input placeholder="John Doe" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <DialogFooter className="flex justify-between">
                <div className="flex gap-2">
                  <Button type="submit" disabled={loading}>
                    {loading ? "Updating..." : "Confirm"}
                  </Button>
                  <DialogClose asChild>
                    <Button type="button" variant="outline">
                      Cancel
                    </Button>
                  </DialogClose>
                </div>

                <DeleteDialog
                  id={delegate_id}
                  uri="delete_delegate"
                  deleteItem="delegate"
                />
              </DialogFooter>
            </form>
          </Form>
        </DialogContent>
      </Dialog>
    </>
  );
};

export default DelegateNameCell;
