'use client'

import Link from 'next/link'
import useArticles from '../../data/articles';

export default function Articles({ id }) {
    const { loading, error, articles, mutate } = useArticles(id);

    if (error) return <div>An error has occurred.</div>
    if (loading) <div>Loading...</div>
    if (!articles || articles?.length === 0) return <div>no articles found</div>

    return (
        <div>
            <ul>
                {articles?.map((article) => (
                    <li key={article.id}>
                        <Link href={`/article/${article.id}`}>{article.title}</Link>
                    </li>
                ))}
            </ul>
        </div>
    );
}