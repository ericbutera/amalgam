import { useState } from "react";
import { getGraph } from "../lib/fetch";
//import { useRouter } from "next/navigation";
import toast from "react-hot-toast";
import { handle } from "../lib/graphErrors";

export default function useAddFeedMutation() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [validationErrors, setValidationErrors] = useState<string[]>([]);

  //const router = useRouter();

  const addFeed = async (url: string) => {
    setLoading(true);
    setError(null);
    setValidationErrors([]);

    try {
      const name = "";
      const resp = await getGraph().AddFeed({ name, url });
      if (resp.addFeed.jobId) {
        toast.success("Feed added");
        // TODO: must wait for jobID to finish before getting feedID
        //router.push(`/feeds/${resp.addFeed.id}/articles`);
        return;
      }
      throw new Error("Failed to add feed");
    } catch (err) {
      handle(err, setValidationErrors, setError);
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
