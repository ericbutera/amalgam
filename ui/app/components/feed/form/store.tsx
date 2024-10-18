import { create } from "zustand";

export const useFormStore = create((set) => ({
  formData: { name: "", url: "" },
  setFormData: (field, value) =>
    set((state) => ({
      formData: { ...state.formData, [field]: value },
    })),
  resetForm: () => set({ formData: { name: "", url: "" } }),
}));
