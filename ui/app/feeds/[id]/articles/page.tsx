'use client'

import Articles from '../../../components/article/list'

export default function Page({ params }: { params: { id: string } }) {
    return (
        <div>
            <h1>viewing articles by feed {params.id}!</h1>
            <Articles id={params.id} />
        </div>
    );
}