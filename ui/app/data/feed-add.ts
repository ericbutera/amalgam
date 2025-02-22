import { useState } from "react";
import { getGraph } from "../lib/fetch";
import toast from "react-hot-toast";
import { handle } from "../lib/graphErrors";

export default function useAddFeedMutation() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [validationErrors, setValidationErrors] = useState<string[]>([]);

  const addFeed = async (url: string) => {
    setLoading(true);
    setError(null);
    setValidationErrors([]);

    let success = false;
    try {
      const name = "";
      const resp = await getGraph().AddFeed({ name, url });
      console.log("feed result %o", resp.addFeed); // TODO: use job id to poll for status
      toast.success("Feed processing started");
      success = true;
    } catch (err) {
      handle(err, setValidationErrors, setError);
    } finally {
      setLoading(false);
    }
    return success;
  };

  return {
    addFeed,
    loading,
    error,
    validationErrors,
  };
}
