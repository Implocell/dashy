import { Card, CardContent, CardMedia, Typography } from "@suid/material";
import { createEffect, createResource, createSignal, Suspense } from "solid-js";
import { getPoem } from "../../api/getPoem";
import styles from "./MemeCard.module.css";
import loadingImg from "../../assets/noose.jpg";
import errorImg from "../../assets/depressed.png";

export const MemeCard = () => {
  const [poem, setPoem] = createSignal(
    "Hang around while OpenAI Davinci writes poetry..."
  );
  const [img, setImg] = createSignal(loadingImg);

  createEffect(() => {
    getPoem("SomeString")
      .then((data) => {
        setPoem(`"${data}" - OpenAI Davinci`);
      })
      .catch((err) => {
        setPoem("OpenAI Davinci encountered an existential crisis!");
        setImg(errorImg);
      });
  });

  return (
    <Card class={styles["memecard"]}>
      <CardMedia component="img" height={256} width={256} image={img()} />
      <CardContent class={styles["poem"]}>
        <Typography align={"center"} variant={"subtitle1"}>
          <i>{poem}</i>
        </Typography>
      </CardContent>
    </Card>
  );
};
