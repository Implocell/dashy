import { CircularProgress } from "@suid/material";
import { createEffect, createResource, onCleanup, Suspense } from "solid-js"
import { getMemes } from "../api/getMemes"
import { MemeCard } from "../components/MemeCard";
import styles from './memes.module.scss'


export const Memes = () => {
    const [memes, {refetch}] = createResource(getMemes);


    return (
        <Suspense fallback={<CircularProgress />}>

<div class={styles["slider"]}>
                <div class={styles["slide-track"]}>
            {memes()?.map(meme => <MemeCard meme={meme} />)}
                </div>
                </div>

        </Suspense>
    )

}