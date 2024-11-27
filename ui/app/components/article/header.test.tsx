import { render, screen } from "@testing-library/react";
import Page from "./header";
import "@testing-library/jest-dom";
import useFeed from "@/app/data/feed";

jest.mock("../../data/feed");

const mockUseFeed = useFeed as jest.Mock;

describe("Header Component", () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  test("renders loading state", () => {
    mockUseFeed.mockImplementation(() => ({
      loading: true,
      error: null,
      feed: null,
    }));

    render(<Page id="1" />);
    expect(screen.getByText("loading...")).toBeInTheDocument();
  });

  test("renders error state", () => {
    mockUseFeed.mockImplementation(() => ({
      loading: false,
      error: "Failed to load feed",
      feed: null,
    }));

    render(<Page id="1" />);
    expect(screen.getByText("failed to load feed")).toBeInTheDocument();
  });

  test("renders feed name", () => {
    mockUseFeed.mockImplementation(() => ({
      loading: false,
      error: null,
      feed: { id: "1", name: "Sample Feed" },
    }));

    render(<Page id="1" />);
    expect(screen.getByText("Sample Feed")).toBeInTheDocument();
  });
});
