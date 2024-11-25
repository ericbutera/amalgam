import { toast } from "react-hot-toast";

/**
 * A reusable function to show toast messages for asynchronous operations.
 * @param {string} loadingMessage - The message to show while the operation is in progress.
 * @param {string} successMessage - The message to show on successful completion.
 * @param {string} errorMessage - The message to show if the operation fails.
 * @param {Function} operation - The async function representing the operation.
 */
const withToast = async (
  loadingMessage: string,
  successMessage: string,
  errorMessage: string,
  operation: () => Promise<void>
) => {
  const toastId = toast.loading(loadingMessage);
  try {
    await operation();
    toast.success(successMessage, {
      id: toastId, // Replaces the loading toast
    });
  } catch (error) {
    toast.error(errorMessage, {
      id: toastId, // Replaces the loading toast
    });
    console.error(error);
  }
};

export default withToast;
