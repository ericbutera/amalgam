import { useState } from "react";
import { getGraph } from "../lib/fetch";

export default function useAddFeedMutation() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [validationErrors, setValidationErrors] = useState<string[]>([]);

  const addFeed = async (name: string, url: string) => {
    setLoading(true);
    setError(null);
    setValidationErrors([]);

    /*
    withToast("Saving feed", "Saved", `Failed to save`, async () => {
      try {
        return await getGraph().AddFeed({ name, url });
      } catch (err: any) {
        if (err?.response?.errors) {
          const errors: string[] = [];
          err.response.errors.forEach((error: any) => {
            if (error.message.includes("feed.url")) {
              errors.push("Please provide a valid URL.");
            } else if (error.message.includes("exists")) {
              errors.push("A feed with this URL already exists.");
            } else {
              errors.push(error.message);
            }
          });
          setValidationErrors(errors);
          setError("Failed to add feed due to validation errors.");
          return false;
        }
      } finally {
        setLoading(false);
      }
    });
    */

    try {
      await getGraph().AddFeed({ name, url });
    } catch (err: any) {
      setError("Error " + err);
    } finally {
      setLoading(false);
    }
  };

  return {
    addFeed,
    loading,
    error,
    validationErrors,
  };
}
