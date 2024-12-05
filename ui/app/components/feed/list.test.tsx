import { render, screen } from "@testing-library/react";
import Feeds from "./list";
import "@testing-library/jest-dom";
import useFeeds from "../../data/feeds";

jest.mock("../../data/feeds");

describe("Feeds Component", () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  test("renders loading state", () => {
    (useFeeds as jest.Mock).mockImplementation(() => ({
      loading: true,
      error: null,
      feeds: [],
    }));

    render(<Feeds />);
    expect(screen.getByText("loading...")).toBeInTheDocument();
  });

  test("renders error state", () => {
    (useFeeds as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: "Failed to load feeds",
      feeds: [],
    }));

    render(<Feeds />);
    expect(screen.getByText("failed to load feeds")).toBeInTheDocument();
  });

  test("renders no feeds found state", () => {
    (useFeeds as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: null,
      feeds: [],
    }));

    render(<Feeds />);
    expect(screen.getByText("no feeds found")).toBeInTheDocument();
  });

  test("renders list of feeds", () => {
    (useFeeds as jest.Mock).mockImplementation(() => ({
      loading: false,
      error: null,
      feeds: [{ id: "1", name: "Feed 1", url: "http://example.com/feed1" }],
    }));

    render(<Feeds />);
    expect(screen.getByText("Feed 1")).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Feed 1" })).toHaveAttribute(
      "href",
      "/feeds/1/articles",
    );
  });
});
