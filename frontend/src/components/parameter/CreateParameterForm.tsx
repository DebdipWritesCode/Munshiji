import { useForm, useWatch } from "react-hook-form";
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
import { useEffect, useState } from "react";
import { toast } from "react-toastify";
import api from "@/api/axios";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

const onlyAlphaAndSpaces = /^[A-Za-z\s]*$/;

interface CreateParameterFormProps {
  isCreate?: boolean;
  id?: number;
  score_sheet_id?: number;
  name?: string;
  rule_type?: string;
  is_special_parameter?: string;
  special_scores_rule?: string;
  special_length_rule?: string;
  score_weight?: number;
  length_weight?: number;
}

const createParameterFormSchema = z.object({
  name: z
    .string()
    .min(1, "Name is required")
    .max(30)
    .regex(onlyAlphaAndSpaces, "Name must contain only letters and spaces"),
  rule_type: z.enum(["average", "absolute", "special"], {
    errorMap: () => ({
      message: "Rule Type must be either 'average', 'absolute', or 'special'",
    }),
  }),
  is_special_parameter: z
    .union([z.literal("true"), z.literal("false")])
    .optional(),
  special_scores_rule: z
    .enum(["average", "absolute"], {
      errorMap: () => ({
        message: "Special Score Rule must be average or absolute",
      }),
    })
    .optional(),
  special_length_rule: z
    .enum(["average", "absolute"], {
      errorMap: () => ({
        message: "Special Length Rule must be average or absolute",
      }),
    })
    .optional(),
  score_weight: z
    .string()
    .refine(
      (value) =>
        !isNaN(parseFloat(value)) &&
        parseFloat(value) >= 0.0 &&
        parseFloat(value) <= 1.0,
      {
        message: "Score Weight must be between 0.0 and 1.0, inclusive",
      }
    )
    .optional(),
  length_weight: z
    .string()
    .refine(
      (value) =>
        !isNaN(parseFloat(value)) &&
        parseFloat(value) >= 0.0 &&
        parseFloat(value) <= 1.0,
      {
        message: "Length Weight must be between 0.0 and 1.0, inclusive",
      }
    )
    .optional(),
  score_sheet_id: z.number().nullable(),
});

const CreateParameterForm: React.FC<CreateParameterFormProps> = ({
  id,
  name,
  score_sheet_id,
  rule_type,
  is_special_parameter,
  special_scores_rule,
  special_length_rule,
  score_weight,
  length_weight,
  isCreate = true,
}) => {
  const [loading, setLoading] = useState(false);

  const form = useForm<z.infer<typeof createParameterFormSchema>>({
    resolver: zodResolver(createParameterFormSchema),
    defaultValues: {
      name: name || "",
      rule_type:
        rule_type === "average" ||
        rule_type === "absolute" ||
        rule_type === "special"
          ? rule_type
          : "average",
      is_special_parameter: is_special_parameter === "true" ? "true" : "false",
      special_scores_rule:
        special_scores_rule === "average" || special_scores_rule === "absolute"
          ? special_scores_rule
          : "average",
      special_length_rule:
        special_length_rule === "average" || special_length_rule === "absolute"
          ? special_length_rule
          : "average",
      score_weight: String(score_weight ?? "1.0"),
      length_weight: String(length_weight ?? "1.0"),
      score_sheet_id: score_sheet_id ?? null,
    },
  });

  const onSubmitCreate = async (
    values: z.infer<typeof createParameterFormSchema>
  ) => {
    setLoading(true);
    try {
      const { is_special_parameter, ...rest } = values;

      const createParameterData = {
        ...rest,
        is_special_parameter: is_special_parameter === "true",
      };

      const response = await api.post("/create_parameter", createParameterData);

      if (response.status === 200) {
        toast.success("Parameter created successfully!");
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
          toast.error("Parameter creation failed. Please try again.");
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
    values: z.infer<typeof createParameterFormSchema>
  ) => {
    setLoading(true);
    try {
      const parameter_id = id;
      const { is_special_parameter, ...rest } = values;
      const updateParameterData = {
        ...rest,
        is_special_parameter: is_special_parameter === "true",
        parameter_id,
      };

      const response = await api.patch(
        "/update_parameter",
        updateParameterData
      );

      if (response.status === 200) {
        toast.success("Parameter updated successfully!");
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
          toast.error("Parameter updation failed. Please try again.");
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

  const ruleType = useWatch({ control: form.control, name: "rule_type" });

  useEffect(() => {
    if (ruleType === "special") {
      form.setValue("is_special_parameter", "true");
    } else {
      form.setValue("is_special_parameter", "false");
    }
  }, [ruleType, form]);

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
                <Input placeholder="My Parameter" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <div className="flex gap-4 items-center">
          <FormField
            control={form.control}
            name="rule_type"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Rule Type</FormLabel>
                <FormControl>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={field.value}>
                    <SelectTrigger className="w-[180px]">
                      <SelectValue placeholder="Select rule type" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectGroup>
                        <SelectLabel>Rule Types</SelectLabel>
                        <SelectItem value="average">Average</SelectItem>
                        <SelectItem value="absolute">Absolute</SelectItem>
                        <SelectItem value="special">Special</SelectItem>
                      </SelectGroup>
                    </SelectContent>
                  </Select>
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        {form.watch("is_special_parameter") === "true" && (
          <>
            <div className="flex gap-4 items-center justify-between">
              <FormField
                control={form.control}
                name="special_scores_rule"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Special Score Rule</FormLabel>
                    <FormControl>
                      <Select
                        onValueChange={field.onChange}
                        defaultValue={field.value}>
                        <SelectTrigger className="w-[180px]">
                          <SelectValue placeholder="Select special score rule" />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectGroup>
                            <SelectLabel>Special Score Rule</SelectLabel>
                            <SelectItem value="average">Average</SelectItem>
                            <SelectItem value="absolute">Absolute</SelectItem>
                          </SelectGroup>
                        </SelectContent>
                      </Select>
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="special_length_rule"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Special Length Rule</FormLabel>
                    <FormControl>
                      <Select
                        onValueChange={field.onChange}
                        defaultValue={field.value}>
                        <SelectTrigger className="w-[180px]">
                          <SelectValue placeholder="Select special length rule" />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectGroup>
                            <SelectLabel>Special Length Rule</SelectLabel>
                            <SelectItem value="absolute">Absolute</SelectItem>
                          </SelectGroup>
                        </SelectContent>
                      </Select>
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>

            <div className="flex gap-4 items-center justify-between">
              <FormField
                control={form.control}
                name="score_weight"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Score Weight</FormLabel>
                    <FormControl>
                      <Input
                        type="number"
                        step="0.01"
                        placeholder="Enter score weight"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="length_weight"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Length Weight</FormLabel>
                    <FormControl>
                      <Input
                        type="number"
                        step="0.01"
                        placeholder="Enter length weight"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>
          </>
        )}

        <Button type="submit" className="w-full" disabled={loading}>
          {loading
            ? isCreate
              ? "Creating..."
              : "Updating..."
            : isCreate
            ? "Create Parameter"
            : "Update Parameter"}
        </Button>
      </form>

    </Form>
  );
};

export default CreateParameterForm;
