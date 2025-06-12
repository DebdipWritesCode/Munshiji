import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useState } from "react";
import { toast } from "react-toastify";
import api from "@/api/axios";
import ToastComponent from "../ToastComponent";
import { useSelector } from "react-redux";
import type { RootState } from "@/store/store";

const onlyAlphaAndSpaces = /^[A-Za-z\s]*$/;

interface createSheetFormProps {
  isCreate?: boolean;
  name?: string;
  id?: number;
  committee_name?: string;
  chair?: string;
  vice_chair?: string;
  rapporteur?: string;
}

const createSheetFormSchema = z.object({
  name: z
    .string()
    .min(1, "Name is required")
    .max(30)
    .regex(onlyAlphaAndSpaces, "Name must contain only letters and spaces"),
  committee_name: z
    .string()
    .min(1, "Committee Name is required")
    .max(30)
    .regex(
      onlyAlphaAndSpaces,
      "Committee Name must contain only letters and spaces"
    ),
  chair: z
    .string()
    .min(1, "Chair is required")
    .max(30)
    .regex(onlyAlphaAndSpaces, "Chair must contain only letters and spaces"),
  vice_chair: z
    .string()
    .max(30)
    .regex(
      onlyAlphaAndSpaces,
      "Vice Chair must contain only letters and spaces"
    )
    .optional(),
  rapporteur: z
    .string()
    .max(30)
    .regex(
      onlyAlphaAndSpaces,
      "Rapporteur must contain only letters and spaces"
    )
    .optional(),
  created_by: z.number().nullable(),
});

const CreateSheetForm: React.FC<createSheetFormProps> = ({
  id,
  name,
  committee_name,
  chair,
  vice_chair,
  rapporteur,
  isCreate = true,
}) => {
  const [loading, setLoading] = useState(false);
  const userId = useSelector((state: RootState) => state.auth.user_id);

  const form = useForm<z.infer<typeof createSheetFormSchema>>({
    resolver: zodResolver(createSheetFormSchema),
    defaultValues: {
      name: name || "",
      committee_name: committee_name || "",
      chair: chair || "",
      vice_chair: vice_chair || "",
      rapporteur: rapporteur || "",
      created_by: userId || null,
    },
  });

  const onSubmitCreate = async (
    values: z.infer<typeof createSheetFormSchema>
  ) => {
    setLoading(true);
    try {
      const { ...createSheetData } = values;

      const response = await api.post("/create_score_sheet", createSheetData);

      if (response.status === 200) {
        toast.success("Score sheet created successfully!");
        // Refresh the page
        setTimeout(() => {
          window.location.reload();
        }, 1000);
      } else {
        throw new Error("Unexpected response from server");
      }
    } catch (err: any) {
      if (err.response) {
        if (err.response.data?.message) {
          toast.error(err.response.data.message);
        } else {
          toast.error("Score sheet creation failed. Please try again.");
        }
      } else if (err.request) {
        toast.error("No response from server. Please check your connection.");
      } else {
        toast.error("An error occurred: " + err.message);
      }
    } finally {
      setLoading(false);
    }
  };

  const onSubmitEdit = async (
    values: z.infer<typeof createSheetFormSchema>
  ) => {
    setLoading(true);
    try {
      const score_sheet_id = id;
      const updateSheetData = { ...values, score_sheet_id };

      const response = await api.patch("/update_score_sheet", updateSheetData);

      if (response.status === 200) {
        toast.success("Score sheet updated successfully!");
        // Refresh the page
        setTimeout(() => {
          window.location.reload();
        }, 1000);
      } else {
        throw new Error("Unexpected response from server");
      }
    } catch (err: any) {
      if (err.response) {
        if (err.response.data?.message) {
          toast.error(err.response.data.message);
        } else {
          toast.error("Score sheet updation failed. Please try again.");
        }
      } else if (err.request) {
        toast.error("No response from server. Please check your connection.");
      } else {
        toast.error("An error occurred: " + err.message);
      }
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = isCreate ? onSubmitCreate : onSubmitEdit;

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(handleSubmit)}
        className="max-w-lg px-3">
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Name</FormLabel>
              <FormControl>
                <Input placeholder="My Score Sheet" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="committee_name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Committee Name</FormLabel>
              <FormControl>
                <Input placeholder="UNSC" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="chair"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Chair</FormLabel>
              <FormControl>
                <Input placeholder="John Doe" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="vice_chair"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Vice Chair(Optional)</FormLabel>
              <FormControl>
                <Input placeholder="Potato Tomato" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="rapporteur"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Rapporteur(Optional)</FormLabel>
              <FormControl>
                <Input placeholder="Potato Tomato" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <Button type="submit" className="w-full" disabled={loading}>
          {loading
            ? isCreate
              ? "Creating..."
              : "Updating..."
            : isCreate
            ? "Create Score Sheet"
            : "Update Score Sheet"}
        </Button>
      </form>

      <ToastComponent />
    </Form>
  );
};

export default CreateSheetForm;
