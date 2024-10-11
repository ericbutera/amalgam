'use client'

import useSWR from "swr";
import Link from 'next/link'

const fetcher = (url) => fetch(url).then((res) => res.json());

export default function Articles({ id }) {
    const { data, error, isLoading } = useSWR(
        `http://localhost:8080/feed/${id}/articles`,
        fetcher
    );

    if (error) return "An error has occurred.";
    if (isLoading) return "Loading...";

    return (
        <div>
            <ul>
                {data.articles.map((article) => (
                    <li key={article.ID}>
                        <Link href={`/article/${article.ID}`}>{article.Title}</Link>
                    </li>
                ))}
            </ul>
        </div>
    );
}