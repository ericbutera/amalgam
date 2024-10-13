'use client'

import Link from 'next/link';
import useArticle from '../../data/article';

export default function Page({ params }: { params: { id: number } }) {
    const { loading, error, article, mutate } = useArticle(params.id);

    if (error) return <div>An error has occurred.</div>
    if (loading) <div>Loading...</div>
    if (!article) return <div>Article not found.</div>

    return (
        <div>
            <h1>viewing article by id {params.id}!</h1>
            <div>
                <h2>{article.title}</h2>
                {article.url && <Link href={article.url}>open</Link>}
                <p>{article.content}</p>
            </div>
        </div>
    );
}