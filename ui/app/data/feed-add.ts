import { useState } from "react";
import { getGraph } from "../lib/fetch";
import toast from "react-hot-toast";

interface ErrorResponse {
  response: {
    errors: {
      message: string;
      path: string[];
    }[];
  };
}

export default function useAddFeedMutation() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [validationErrors, setValidationErrors] = useState<string[]>([]);

  const addFeed = async (name: string, url: string) => {
    setLoading(true);
    setError(null);
    setValidationErrors([]);

    try {
      await getGraph().AddFeed({ name, url });
      toast.success("Feed added");
      // TODO: clear form
    } catch (err) {
      // TODO: multiple error types
      // - protobuf validation errors
      // - service layer validation errors
      //
      // UI shouldn't have to care about these concerns. Messages should be consistent. However, protobuf errors are in the middleware and there
      // isn't a clean way to modify them without rewriting the middleware. Revisit later.
      if ((err as ErrorResponse).response?.errors) {
        const errs: string[] = [];
        for (const error of (err as ErrorResponse).response.errors) {
          errs.push(error.message);
        }
        setValidationErrors(errs);
      } else {
        setError(`Error: ${err}`);
      }
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
