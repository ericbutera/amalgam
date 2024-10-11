'use client'

import Link from 'next/link'
import useArticles from '../data/articles';

export default function Articles({ id }) {
    const { loading, error, articles, mutate } = useArticles(id);

    if (error) return <div>An error has occurred.</div>
    if (loading) <div>Loading...</div>

    return (
        <div>
            <ul>
                {articles.map((article) => (
                    <li key={article.ID}>
                        <Link href={`/article/${article.ID}`}>{article.Title}</Link>
                    </li>
                ))}
            </ul>
        </div>
    );
}